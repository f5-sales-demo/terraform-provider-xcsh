package provider

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/blindfold"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// The adapter leaves generated API models unchanged. Provider-only values are projected
// away before delegating serialization; only normalized public PEM and encrypted keys cross it.
type blindfoldResource struct {
	inner            resource.Resource
	client           *client.Client
	collection, item string
	named            bool
}

func newBlindfoldResource(inner resource.Resource, collection, item string, named bool) resource.Resource {
	return &blindfoldResource{inner: inner, collection: collection, item: item, named: named}
}
func (r *blindfoldResource) Metadata(c context.Context, q resource.MetadataRequest, s *resource.MetadataResponse) {
	r.inner.Metadata(c, q, s)
}
func (r *blindfoldResource) Configure(c context.Context, q resource.ConfigureRequest, s *resource.ConfigureResponse) {
	if v, ok := q.ProviderData.(*client.Client); ok {
		r.client = v
	}
	r.inner.(resource.ResourceWithConfigure).Configure(c, q, s)
}
func (r *blindfoldResource) base(c context.Context) schema.Schema {
	var s resource.SchemaResponse
	r.inner.Schema(c, resource.SchemaRequest{}, &s)
	return s.Schema
}
func (r *blindfoldResource) Schema(c context.Context, q resource.SchemaRequest, s *resource.SchemaResponse) {
	r.inner.Schema(c, q, s)
	augmentBlindfold(s.Schema.Attributes, s.Schema.Blocks)
}
func blindfoldSchema() schema.SingleNestedAttribute {
	a := map[string]schema.Attribute{}
	for _, k := range []string{"id", "certificate_file", "private_key_file", "pkcs12_file", "certificate_pem", "passphrase_env", "policy", "material_version"} {
		a[k] = schema.StringAttribute{Optional: true}
	}
	for _, k := range []string{"private_key_wo", "pkcs12_wo", "passphrase_wo"} {
		a[k] = schema.StringAttribute{Optional: true, Sensitive: true, WriteOnly: true}
	}
	for _, k := range []string{"fingerprint", "expires_at", "context_digest", "chain_identity", "spki_identity", "algorithm", "prepared_identity"} {
		a[k] = schema.StringAttribute{Computed: true}
	}
	a["encrypted_location"] = schema.StringAttribute{Computed: true, Sensitive: true}
	return schema.SingleNestedAttribute{Optional: true, Attributes: a, MarkdownDescription: "Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with material_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates require unique IDs. Private inputs are never stored."}
}
func augmentBlindfold(a map[string]schema.Attribute, b map[string]schema.Block) {
	_, cert := a["certificate_url"]
	_, key := b["private_key"]
	if cert && key {
		a["blindfold"] = blindfoldSchema()
		v := a["certificate_url"].(schema.StringAttribute)
		v.Required = false
		v.Optional = true
		v.Computed = true
		a["certificate_url"] = v
	}
	for k, v := range b {
		switch n := v.(type) {
		case schema.SingleNestedBlock:
			augmentBlindfold(n.Attributes, n.Blocks)
			b[k] = n
		case schema.ListNestedBlock:
			augmentBlindfold(n.NestedObject.Attributes, n.NestedObject.Blocks)
			b[k] = n
		case schema.SetNestedBlock:
			augmentBlindfold(n.NestedObject.Attributes, n.NestedObject.Blocks)
			b[k] = n
		}
	}
}

// projectValue changes only shape, retaining exact scalar and unknown values.
func projectValue(v tftypes.Value, t tftypes.Type) tftypes.Value {
	if v.Type() == nil || v.IsNull() {
		return tftypes.NewValue(t, nil)
	}
	if !v.IsKnown() {
		return tftypes.NewValue(t, tftypes.UnknownValue)
	}
	switch target := t.(type) {
	case tftypes.Object:
		source := map[string]tftypes.Value{}
		if v.As(&source) != nil {
			return tftypes.NewValue(t, nil)
		}
		out := map[string]tftypes.Value{}
		for k, typ := range target.AttributeTypes {
			out[k] = projectValue(source[k], typ)
		}
		return tftypes.NewValue(t, out)
	case tftypes.List:
		var source []tftypes.Value
		if v.As(&source) != nil {
			return tftypes.NewValue(t, nil)
		}
		out := make([]tftypes.Value, len(source))
		for i, x := range source {
			out[i] = projectValue(x, target.ElementType)
		}
		return tftypes.NewValue(t, out)
	case tftypes.Set:
		var source []tftypes.Value
		if v.As(&source) != nil {
			return tftypes.NewValue(t, nil)
		}
		out := make([]tftypes.Value, len(source))
		for i, x := range source {
			out[i] = projectValue(x, target.ElementType)
		}
		return tftypes.NewValue(t, out)
	}
	return v
}
func objectValue(v tftypes.Value) map[string]tftypes.Value {
	m := map[string]tftypes.Value{}
	if v.Type() != nil && v.IsKnown() && !v.IsNull() {
		_ = v.As(&m)
	}
	return m
}
func textValue(v tftypes.Value) string {
	var s string
	if v.Type() != nil && v.IsKnown() && !v.IsNull() {
		_ = v.As(&s)
	}
	return s
}
func setString(m map[string]tftypes.Value, k, v string) { m[k] = tftypes.NewValue(tftypes.String, v) }

type nativeNode struct {
	path               []any
	id                 string
	config, plan       map[string]tftypes.Value
	material           *blindfold.Material
	context            blindfold.Context
	location, identity string
	rotate             bool
}

func walkNative(v tftypes.Value, p []any, fn func([]any, map[string]tftypes.Value)) {
	if v.Type() == nil || !v.IsKnown() || v.IsNull() {
		return
	}
	switch v.Type().(type) {
	case tftypes.Object:
		m := objectValue(v)
		if x, ok := m["blindfold"]; ok && x.IsKnown() && !x.IsNull() {
			fn(p, m)
		}
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			if k != "blindfold" {
				walkNative(m[k], append(append([]any{}, p...), k), fn)
			}
		}
	case tftypes.List, tftypes.Set:
		var a []tftypes.Value
		_ = v.As(&a)
		for i, x := range a {
			walkNative(x, append(append([]any{}, p...), i), fn)
		}
	}
}
func atValue(v tftypes.Value, p []any) tftypes.Value {
	for _, x := range p {
		switch k := x.(type) {
		case string:
			v = objectValue(v)[k]
		case int:
			var a []tftypes.Value
			if v.Type() == nil || v.As(&a) != nil || k >= len(a) {
				return tftypes.Value{}
			}
			v = a[k]
		}
	}
	return v
}
func editValue(v tftypes.Value, p []any, fn func(map[string]tftypes.Value)) tftypes.Value {
	if len(p) == 0 {
		m := objectValue(v)
		fn(m)
		return tftypes.NewValue(v.Type(), m)
	}
	switch k := p[0].(type) {
	case string:
		m := objectValue(v)
		m[k] = editValue(m[k], p[1:], fn)
		return tftypes.NewValue(v.Type(), m)
	case int:
		var a []tftypes.Value
		_ = v.As(&a)
		a[k] = editValue(a[k], p[1:], fn)
		return tftypes.NewValue(v.Type(), a)
	}
	return v
}
func (r *blindfoldResource) context(c context.Context, policy string) (blindfold.Context, error) {
	if r.client == nil {
		return blindfold.Context{}, errors.New("Blindfold provider is not configured")
	}
	if policy == "" {
		policy = "shared/ves-io-allow-volterra"
	}
	p := strings.Split(policy, "/")
	if len(p) != 2 {
		return blindfold.Context{}, errors.New("policy must be namespace/name")
	}
	for _, v := range p {
		if !blindfoldLabel(v) {
			return blindfold.Context{}, errors.New("invalid policy identity")
		}
	}
	var pub, pol json.RawMessage
	if r.client.Get(c, "/api/secret_management/get_public_key", &pub) != nil || r.client.Get(c, fmt.Sprintf("/api/secret_management/namespaces/%s/secret_policys/%s/get_policy_document", p[0], p[1]), &pol) != nil {
		return blindfold.Context{}, errors.New("cannot refresh Blindfold encryption context")
	}
	return blindfold.ParseContext(pub, pol)
}
func blindfoldLabel(s string) bool {
	if s == "" || len(s) > 63 {
		return false
	}
	for i, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || (c == '-' && i > 0 && i < len(s)-1) {
			continue
		}
		return false
	}
	return true
}
func inputMaterial(m map[string]tftypes.Value) (*blindfold.Material, bool, error) {
	var in blindfold.Input
	defer func() { clear(in.PrivateKey); clear(in.PKCS12); clear(in.Passphrase) }()
	file := func(k string) ([]byte, error) {
		if s := textValue(m[k]); s != "" {
			return blindfold.ReadFile(s)
		}
		return nil, nil
	}
	cf, kf, pf := textValue(m["certificate_file"]), textValue(m["private_key_file"]), textValue(m["pkcs12_file"])
	cp, kw, pw := textValue(m["certificate_pem"]), textValue(m["private_key_wo"]), textValue(m["pkcs12_wo"])
	if (kf != "" && kw != "") || (pf != "" && pw != "") || (cf != "" && cp != "") || ((pf != "" || pw != "") && (cf != "" || kf != "" || cp != "" || kw != "")) {
		return nil, false, errors.New("conflicting Blindfold certificate sources")
	}
	version := textValue(m["material_version"])
	wo := kw != "" || pw != "" || (kf == "" && pf == "")
	if wo && version == "" {
		return nil, false, errors.New("write-only private inputs require material_version")
	}
	env, pass := textValue(m["passphrase_env"]), textValue(m["passphrase_wo"])
	if env != "" && pass != "" {
		return nil, false, errors.New("passphrase sources conflict")
	}
	if env != "" {
		if !validEnv(env) {
			return nil, false, errors.New("invalid passphrase environment name")
		}
		var ok bool
		pass, ok = os.LookupEnv(env)
		if !ok {
			return nil, false, errors.New("passphrase environment variable is unavailable")
		}
	}
	in.Passphrase = []byte(pass)
	var err error
	in.CertificatePEM, err = file("certificate_file")
	if err != nil {
		return nil, wo, err
	}
	if cp != "" {
		in.CertificatePEM = []byte(cp)
	}
	in.PrivateKey, err = file("private_key_file")
	if err != nil {
		return nil, wo, err
	}
	if kw != "" {
		in.PrivateKey = []byte(kw)
	}
	in.PKCS12, err = file("pkcs12_file")
	if err != nil {
		return nil, wo, err
	}
	if pw != "" {
		in.PKCS12, err = base64.StdEncoding.Strict().DecodeString(pw)
		if err != nil {
			return nil, wo, errors.New("invalid write-only PKCS12 base64")
		}
	}
	if wo && len(in.PrivateKey) == 0 && len(in.PKCS12) == 0 {
		return nil, true, nil
	}
	material, err := blindfold.Normalize(in)
	return material, wo, err
}
func validEnv(s string) bool {
	for i, c := range s {
		if c == '_' || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (i > 0 && c >= '0' && c <= '9') {
			continue
		}
		return false
	}
	return s != ""
}
func (r *blindfoldResource) remote(c context.Context, raw tftypes.Value) (map[string]any, error) {
	m := objectValue(raw)
	ns, name := textValue(m["namespace"]), textValue(m["name"])
	if !blindfoldLabel(ns) || !blindfoldLabel(name) {
		return nil, errors.New("resource identity must be known")
	}
	route := fmt.Sprintf(r.item, ns, name)
	var body map[string]any
	route += "?response_format=2"
	if err := r.client.Get(c, route, &body); err != nil {
		return nil, err
	}
	if f, ok := body["replace_form"].(map[string]any); ok {
		if token, ok := body["resource_version"]; ok {
			f["resource_version"] = token
		}
		return f, nil
	}
	return body, nil
}
func remoteNode(body map[string]any, p []any) map[string]any {
	v := body["spec"]
	for _, step := range p {
		switch k := step.(type) {
		case string:
			m, ok := v.(map[string]any)
			if !ok {
				return nil
			}
			v = m[k]
		case int:
			a, ok := v.([]any)
			if !ok || k >= len(a) {
				return nil
			}
			v = a[k]
		}
	}
	m, _ := v.(map[string]any)
	return m
}
func remoteLocation(n map[string]any) string {
	k, _ := n["private_key"].(map[string]any)
	b, _ := k["blindfold_secret_info"].(map[string]any)
	s, _ := b["location"].(string)
	return s
}
func provenance(body map[string]any) blindfold.Provenance {
	p := blindfold.Provenance{}
	meta, _ := body["metadata"].(map[string]any)
	a, _ := meta["annotations"].(map[string]any)
	s, _ := a[blindfold.Annotation].(string)
	if len(s) > blindfold.MaxEncoded {
		return p
	}
	_ = json.Unmarshal([]byte(s), &p)
	if p.Version != 1 {
		return blindfold.Provenance{}
	}
	return p
}
func (r *blindfoldResource) nodes(c context.Context, config, plan, state tftypes.Value, apply bool) ([]*nativeNode, error) {
	if err := blindfold.ValidateContract(); err != nil {
		return nil, err
	}
	var nodes []*nativeNode
	var failure error
	ids := map[string]bool{}
	contexts := map[string]blindfold.Context{}
	var remote map[string]any
	if !state.IsNull() && state.Type() != nil {
		var err error
		remote, err = r.remote(c, plan)
		if err != nil && !client.HasHTTPStatus(err, 404) {
			return nil, errors.New("cannot inspect remote certificate configuration")
		}
	}
	prov := provenance(remote)
	walkNative(config, nil, func(p []any, parent map[string]tftypes.Value) {
		if failure != nil {
			return
		}
		input := objectValue(parent["blindfold"])
		id := textValue(input["id"])
		if id == "" && r.named {
			id = "default"
		}
		if !blindfoldLabel(id) || ids[id] {
			failure = errors.New("inline Blindfold IDs must be explicit and unique")
			return
		}
		ids[id] = true
		if v := parent["certificate_url"]; v.IsKnown() && !v.IsNull() {
			failure = errors.New("Blindfold conflicts with supplied certificate_url")
			return
		}
		if v := parent["private_key"]; v.Type() != nil && v.IsKnown() && !v.IsNull() {
			failure = errors.New("Blindfold conflicts with supplied private_key")
			return
		}
		policy := textValue(input["policy"])
		context, ok := contexts[policy]
		if !ok {
			var err error
			context, err = r.context(c, policy)
			if err != nil {
				failure = err
				return
			}
			contexts[policy] = context
		}
		material, wo, err := inputMaterial(input)
		if err != nil {
			failure = err
			return
		}
		plannedParent := objectValue(atValue(plan, p))
		planned := objectValue(plannedParent["blindfold"])
		oldParent := objectValue(atValue(state, p))
		old := objectValue(oldParent["blindfold"])
		walkNative(state, nil, func(_ []any, candidate map[string]tftypes.Value) {
			entry := objectValue(candidate["blindfold"])
			previousID := textValue(entry["id"])
			if previousID == "" && r.named {
				previousID = "default"
			}
			if previousID == id {
				old = entry
			}
		})
		node := &nativeNode{path: p, id: id, config: input, plan: planned, material: material, context: context}
		nodes = append(nodes, node)
		remoteN := remoteNode(remote, p)
		if stable := remoteByEntry(remote, prov.Entries[id]); stable != nil {
			remoteN = stable
		}
		location := remoteLocation(remoteN)
		cert, _ := remoteN["certificate_url"].(string)
		if material != nil {
			node.identity = blindfold.Hash([]byte(material.ChainIdentity + "|" + material.SPKIIdentity + "|" + context.Digest + "|" + textValue(input["material_version"])))
			node.rotate = !blindfold.Matches(prov.Entries[id], material, context.Digest, cert, location)
			if previousVersion := textValue(old["material_version"]); previousVersion != "" && previousVersion != textValue(input["material_version"]) {
				node.rotate = true
			}
			// A real remote key drift repairs the retained encrypted result. Provenance
			// itself must match that retained result and normalized public chain/context.
			retained := textValue(old["encrypted_location"])
			if node.rotate && retained != "" && textValue(old["chain_identity"]) == material.ChainIdentity && textValue(old["context_digest"]) == context.Digest && textValue(old["material_version"]) == textValue(input["material_version"]) && blindfold.Matches(prov.Entries[id], material, context.Digest, material.CertificateURL, retained) {
				node.rotate = false
				location = retained
			}
			if !node.rotate {
				node.location = location
			}
			previous := prov.Entries[id].Algorithm
			if previous == "" {
				previous = textValue(old["algorithm"])
			}
			if previous == "" && cert != "" {
				previous = publicAlgorithm(cert)
			}
			if previous != "" && previous != material.Algorithm {
				failure = errors.New("RSA/EC algorithm changes require a new certificate and reference cutover")
				return
			}
		} else {
			node.identity = blindfold.Hash([]byte(textValue(old["chain_identity"]) + "|" + textValue(old["spki_identity"]) + "|" + context.Digest + "|" + textValue(input["material_version"])))
			node.rotate = textValue(old["prepared_identity"]) != node.identity || prov.Entries[id].Ciphertext != blindfold.Hash([]byte(location)) || prov.Entries[id].Chain != textValue(old["chain_identity"])
			if !node.rotate {
				node.location = location
			}
			if node.rotate && location != "" && prov.Entries[id].Ciphertext == blindfold.Hash([]byte(location)) && textValue(old["material_version"]) == textValue(input["material_version"]) && textValue(old["context_digest"]) == context.Digest {
				node.rotate = false
				node.location = location
			}
		}
		if apply {
			if expectedContext := textValue(planned["context_digest"]); expectedContext != "" && expectedContext != context.Digest {
				failure = errors.New("Blindfold encryption context changed after planning; create a new plan")
				return
			}
			expected := textValue(planned["prepared_identity"])
			if expected != "" && expected != node.identity && (!wo || textValue(planned["chain_identity"]) != "") {
				failure = errors.New("Blindfold material or encryption context changed after planning; create a new plan")
				return
			}
			if node.rotate {
				if material == nil {
					failure = errors.New("write-only private material is required for apply")
					return
				}
				node.location, err = blindfold.Encrypt(material.KeyPEM, context)
				if err != nil {
					failure = err
					return
				}
			}
		}
		_ = wo
	})
	if failure != nil {
		for _, n := range nodes {
			if n.material != nil {
				n.material.Close()
			}
		}
		return nil, failure
	}
	return nodes, nil
}
func closeNodes(n []*nativeNode) {
	for _, v := range n {
		if v.material != nil {
			v.material.Close()
		}
	}
}
func nativeOutputs(raw tftypes.Value, nodes []*nativeNode, applying bool) tftypes.Value {
	for _, n := range nodes {
		raw = editValue(raw, n.path, func(parent map[string]tftypes.Value) {
			v := parent["blindfold"]
			m := objectValue(v)
			for _, k := range []string{"private_key_wo", "pkcs12_wo", "passphrase_wo"} {
				m[k] = tftypes.NewValue(tftypes.String, nil)
			}
			setString(m, "context_digest", n.context.Digest)
			if n.material == nil && textValue(m["chain_identity"]) == "" && !applying {
				m["prepared_identity"] = tftypes.NewValue(tftypes.String, tftypes.UnknownValue)
			} else {
				setString(m, "prepared_identity", n.identity)
			}
			if n.material != nil {
				for k, s := range map[string]string{"fingerprint": n.material.Fingerprint, "expires_at": n.material.ExpiresAt, "chain_identity": n.material.ChainIdentity, "spki_identity": n.material.SPKIIdentity, "algorithm": n.material.Algorithm} {
					setString(m, k, s)
				}
				setString(parent, "certificate_url", n.material.CertificateURL)
			}
			if n.rotate && !applying {
				m["encrypted_location"] = tftypes.NewValue(tftypes.String, tftypes.UnknownValue)
			} else {
				setString(m, "encrypted_location", n.location)
			}
			parent["blindfold"] = tftypes.NewValue(v.Type(), m)
		})
	}
	return raw
}
func injectNative(raw tftypes.Value, nodes []*nativeNode) tftypes.Value {
	for _, n := range nodes {
		raw = editValue(raw, n.path, func(m map[string]tftypes.Value) {
			if n.material != nil {
				setString(m, "certificate_url", n.material.CertificateURL)
			}
			typ := m["private_key"].Type()
			m["private_key"] = valueFromJSON(typ, map[string]any{"blindfold_secret_info": map[string]any{"location": n.location}})
		})
	}
	return raw
}
func valueFromJSON(t tftypes.Type, v any) tftypes.Value {
	if v == nil {
		return tftypes.NewValue(t, nil)
	}
	switch typ := t.(type) {
	case tftypes.Object:
		m, _ := v.(map[string]any)
		out := map[string]tftypes.Value{}
		for k, child := range typ.AttributeTypes {
			out[k] = valueFromJSON(child, m[k])
		}
		return tftypes.NewValue(t, out)
	case tftypes.List:
		a, _ := v.([]any)
		out := make([]tftypes.Value, len(a))
		for i, x := range a {
			out[i] = valueFromJSON(typ.ElementType, x)
		}
		return tftypes.NewValue(t, out)
	case tftypes.Map:
		m, _ := v.(map[string]any)
		out := map[string]tftypes.Value{}
		for k, x := range m {
			out[k] = valueFromJSON(typ.ElementType, x)
		}
		return tftypes.NewValue(t, out)
	}
	return tftypes.NewValue(t, v)
}
func mergeNative(base, original tftypes.Value) tftypes.Value {
	if original.Type() == nil || original.IsNull() {
		return projectValue(base, original.Type())
	}
	if base.Type() == nil || base.IsNull() || !base.IsKnown() {
		return projectValue(base, original.Type())
	}
	switch original.Type().(type) {
	case tftypes.Object:
		old := objectValue(original)
		m := objectValue(base)
		out := map[string]tftypes.Value{}
		for k, v := range old {
			if k == "blindfold" {
				out[k] = v
			} else {
				out[k] = mergeNative(m[k], v)
			}
		}
		if v, ok := old["blindfold"]; ok && !v.IsNull() && old["private_key"].Type() != nil {
			out["private_key"] = tftypes.NewValue(old["private_key"].Type(), nil)
		}
		return tftypes.NewValue(original.Type(), out)
	case tftypes.List:
		var a, b []tftypes.Value
		_ = base.As(&a)
		_ = original.As(&b)
		out := make([]tftypes.Value, len(a))
		for i, v := range a {
			if i < len(b) {
				out[i] = mergeNative(v, b[i])
			} else {
				out[i] = v
			}
		}
		return tftypes.NewValue(original.Type(), out)
	}
	return base
}
func annotationRaw(raw tftypes.Value, nodes []*nativeNode, remote map[string]any) tftypes.Value {
	m := objectValue(raw)
	v := m["annotations"]
	a := map[string]tftypes.Value{}
	if meta, ok := remote["metadata"].(map[string]any); ok {
		if existing, ok := meta["annotations"].(map[string]any); ok {
			for k, value := range existing {
				if text, ok := value.(string); ok {
					a[k] = tftypes.NewValue(tftypes.String, text)
				}
			}
		}
	}
	if v.Type() != nil && !v.IsNull() && v.IsKnown() {
		declared := map[string]tftypes.Value{}
		_ = v.As(&declared)
		for k, value := range declared {
			a[k] = value
		}
	}
	p := provenance(remote)
	if p.Entries == nil {
		p = blindfold.Provenance{Version: 1, Entries: map[string]blindfold.Entry{}}
	}
	for _, n := range nodes {
		if n.material != nil {
			p.Entries[n.id] = n.material.Entry(n.context.Digest, n.location)
		}
	}
	b, _ := json.Marshal(p)
	a[blindfold.Annotation] = tftypes.NewValue(tftypes.String, string(b))
	m["annotations"] = tftypes.NewValue(v.Type(), a)
	return tftypes.NewValue(raw.Type(), m)
}
func nativeError(d *diag.Diagnostics, err error) {
	d.AddError("Blindfold certificate preparation", err.Error())
}

func (r *blindfoldResource) ValidateConfig(c context.Context, q resource.ValidateConfigRequest, s *resource.ValidateConfigResponse) {
	b := r.base(c)
	validateNativeSources(q.Config.Raw, &s.Diagnostics, r.named)
	q.Config = tfsdk.Config{Raw: projectValue(q.Config.Raw, b.Type().TerraformType(c)), Schema: b}
	r.inner.(resource.ResourceWithValidateConfig).ValidateConfig(c, q, s)
}
func (r *blindfoldResource) ModifyPlan(c context.Context, q resource.ModifyPlanRequest, s *resource.ModifyPlanResponse) {
	if q.Plan.Raw.IsNull() {
		return
	}
	b := r.base(c)
	original := q.Plan
	inner := q
	inner.Config = tfsdk.Config{Raw: projectValue(q.Config.Raw, b.Type().TerraformType(c)), Schema: b}
	inner.Plan = tfsdk.Plan{Raw: projectValue(q.Plan.Raw, b.Type().TerraformType(c)), Schema: b}
	inner.State = tfsdk.State{Raw: projectValue(q.State.Raw, b.Type().TerraformType(c)), Schema: b}
	reply := *s
	reply.Plan = inner.Plan
	r.inner.(resource.ResourceWithModifyPlan).ModifyPlan(c, inner, &reply)
	s.Diagnostics.Append(reply.Diagnostics...)
	s.RequiresReplace = reply.RequiresReplace
	if s.Diagnostics.HasError() {
		return
	}
	s.Plan = original
	s.Plan.Raw = mergeNative(reply.Plan.Raw, original.Raw)
	nodes, err := r.nodes(c, q.Config.Raw, s.Plan.Raw, q.State.Raw, false)
	if err != nil {
		nativeError(&s.Diagnostics, err)
		return
	}
	defer closeNodes(nodes)
	s.Plan.Raw = nativeOutputs(s.Plan.Raw, nodes, false)
	for _, n := range nodes {
		previous := objectValue(objectValue(atValue(q.State.Raw, n.path))["blindfold"])
		if textValue(previous["prepared_identity"]) == "remote-drift" {
			s.Plan.Raw = editValue(s.Plan.Raw, n.path, func(parent map[string]tftypes.Value) {
				v := parent["blindfold"]
				m := objectValue(v)
				m["encrypted_location"] = tftypes.NewValue(tftypes.String, tftypes.UnknownValue)
				parent["blindfold"] = tftypes.NewValue(v.Type(), m)
			})
		}
	}
	// Preserve output identities by stable ID when list elements are reordered.
	for _, n := range nodes {
		if n.material == nil {
			walkNative(q.State.Raw, nil, func(_ []any, parent map[string]tftypes.Value) {
				m := objectValue(parent["blindfold"])
				id := textValue(m["id"])
				if id == "" && r.named {
					id = "default"
				}
				if id == n.id {
					s.Plan.Raw = editValue(s.Plan.Raw, n.path, func(p map[string]tftypes.Value) {
						v := p["blindfold"]
						out := objectValue(v)
						for _, k := range []string{"fingerprint", "expires_at", "chain_identity", "spki_identity", "algorithm"} {
							out[k] = m[k]
						}
						p["blindfold"] = tftypes.NewValue(v.Type(), out)
					})
				}
			})
		}
	}
}
func (r *blindfoldResource) Create(c context.Context, q resource.CreateRequest, s *resource.CreateResponse) {
	nodes, err := r.nodes(c, q.Config.Raw, q.Plan.Raw, tftypes.Value{}, true)
	if err != nil {
		nativeError(&s.Diagnostics, err)
		return
	}
	defer closeNodes(nodes)
	b := r.base(c)
	out := nativeOutputs(q.Plan.Raw, nodes, true)
	raw := projectValue(out, b.Type().TerraformType(c))
	raw = injectNative(raw, nodes)
	if len(nodes) > 0 {
		raw = annotationRaw(raw, nodes, nil)
	}
	q.Config = tfsdk.Config{Raw: projectValue(q.Config.Raw, b.Type().TerraformType(c)), Schema: b}
	q.Plan = tfsdk.Plan{Raw: raw, Schema: b}
	s.State = tfsdk.State{Raw: raw, Schema: b}
	r.inner.Create(c, q, s)
	if s.Diagnostics.HasError() {
		return
	}
	if s.State.Raw.IsNull() {
		return
	}
	var enhanced resource.SchemaResponse
	r.Schema(c, resource.SchemaRequest{}, &enhanced)
	s.State = tfsdk.State{Raw: mergeNative(s.State.Raw, out), Schema: enhanced.Schema}
	if len(nodes) > 0 {
		s.State.Raw = restoreAnnotations(s.State.Raw, out)
	}
}
func (r *blindfoldResource) Update(c context.Context, q resource.UpdateRequest, s *resource.UpdateResponse) {
	nodes, err := r.nodes(c, q.Config.Raw, q.Plan.Raw, q.State.Raw, true)
	if err != nil {
		nativeError(&s.Diagnostics, err)
		return
	}
	defer closeNodes(nodes)
	b := r.base(c)
	out := nativeOutputs(q.Plan.Raw, nodes, true)
	raw := injectNative(projectValue(out, b.Type().TerraformType(c)), nodes)
	if len(nodes) > 0 {
		remote, err := r.remote(c, q.Plan.Raw)
		if err != nil {
			nativeError(&s.Diagnostics, errors.New("cannot inspect remote metadata"))
			return
		}
		raw = annotationRaw(raw, nodes, remote)
	}
	q.Config = tfsdk.Config{Raw: projectValue(q.Config.Raw, b.Type().TerraformType(c)), Schema: b}
	q.State = tfsdk.State{Raw: projectValue(q.State.Raw, b.Type().TerraformType(c)), Schema: b}
	q.Plan = tfsdk.Plan{Raw: raw, Schema: b}
	s.State = tfsdk.State{Raw: raw, Schema: b}
	r.inner.Update(c, q, s)
	if s.Diagnostics.HasError() {
		return
	}
	var enhanced resource.SchemaResponse
	r.Schema(c, resource.SchemaRequest{}, &enhanced)
	s.State = tfsdk.State{Raw: mergeNative(s.State.Raw, out), Schema: enhanced.Schema}
	if len(nodes) > 0 {
		s.State.Raw = restoreAnnotations(s.State.Raw, out)
	}
}
func (r *blindfoldResource) Read(c context.Context, q resource.ReadRequest, s *resource.ReadResponse) {
	original := q.State
	b := r.base(c)
	q.State = tfsdk.State{Raw: projectValue(original.Raw, b.Type().TerraformType(c)), Schema: b}
	s.State = q.State
	r.inner.Read(c, q, s)
	if s.State.Raw.IsNull() {
		return
	}
	raw := mergeNative(s.State.Raw, original.Raw)
	var remote map[string]any
	walkNative(original.Raw, nil, func(p []any, parent map[string]tftypes.Value) {
		if remote == nil {
			var err error
			remote, err = r.remote(c, original.Raw)
			if err != nil {
				nativeError(&s.Diagnostics, errors.New("cannot observe remote Blindfold material"))
				return
			}
		}
		n := remoteNode(remote, p)
		raw = editValue(raw, p, func(m map[string]tftypes.Value) {
			v := m["blindfold"]
			out := objectValue(v)
			retained := textValue(out["encrypted_location"])
			// Keep the managed encrypted result for repair; mark the planned output
			// unknown on drift so Terraform invokes Update.
			if retained != remoteLocation(n) || textValue(parent["certificate_url"]) != fmt.Sprint(n["certificate_url"]) {
				setString(out, "prepared_identity", "remote-drift")
			}
			m["blindfold"] = tftypes.NewValue(v.Type(), out)
		})
	})
	native := false
	walkNative(original.Raw, nil, func(_ []any, _ map[string]tftypes.Value) { native = true })
	if native {
		raw = restoreAnnotations(raw, original.Raw)
	}
	s.State = tfsdk.State{Raw: raw, Schema: original.Schema}
}
func (r *blindfoldResource) Delete(c context.Context, q resource.DeleteRequest, s *resource.DeleteResponse) {
	b := r.base(c)
	q.State = tfsdk.State{Raw: projectValue(q.State.Raw, b.Type().TerraformType(c)), Schema: b}
	s.State = q.State
	r.inner.Delete(c, q, s)
}
func (r *blindfoldResource) ImportState(c context.Context, q resource.ImportStateRequest, s *resource.ImportStateResponse) {
	r.inner.(resource.ResourceWithImportState).ImportState(c, q, s)
}

var _ resource.ResourceWithModifyPlan = (*blindfoldResource)(nil)

func restoreAnnotations(raw, source tftypes.Value) tftypes.Value {
	m := objectValue(raw)
	original := objectValue(source)
	if v, ok := original["annotations"]; ok {
		m["annotations"] = v
	}
	return tftypes.NewValue(raw.Type(), m)
}
func publicAlgorithm(location string) string {
	if !strings.HasPrefix(location, "string:///") {
		return ""
	}
	b, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(location, "string:///"))
	if err != nil {
		return ""
	}
	block, _ := pem.Decode(b)
	if block == nil {
		return ""
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return ""
	}
	switch cert.PublicKeyAlgorithm {
	case x509.RSA:
		return "RSA"
	case x509.ECDSA:
		return "EC"
	}
	return ""
}

// Provenance IDs remain stable across list insertion/reordering. Locate the
// actual ciphertext by its public metadata before reusing it at a new pointer.
func remoteByEntry(body map[string]any, entry blindfold.Entry) map[string]any {
	var found map[string]any
	var visit func(any)
	visit = func(value any) {
		switch v := value.(type) {
		case map[string]any:
			if loc := remoteLocation(v); loc != "" && blindfold.Hash([]byte(loc)) == entry.Ciphertext {
				if cert, ok := v["certificate_url"].(string); ok && strings.HasPrefix(cert, "string:///") {
					chain, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(cert, "string:///"))
					if err == nil && blindfold.Hash(chain) == entry.Chain {
						found = v
					}
				}
			}
			for _, child := range v {
				visit(child)
			}
		case []any:
			for _, child := range v {
				visit(child)
			}
		}
	}
	visit(body["spec"])
	return found
}

func validateNativeSources(raw tftypes.Value, diagnostics *diag.Diagnostics, named bool) {
	var walk func(tftypes.Value)
	walk = func(v tftypes.Value) {
		if v.Type() == nil || v.IsNull() || !v.IsKnown() {
			return
		}
		switch v.Type().(type) {
		case tftypes.Object:
			m := objectValue(v)
			if _, node := m["blindfold"]; node {
				native := m["blindfold"]
				cert := m["certificate_url"]
				key := m["private_key"]
				if native.IsNull() && cert.IsNull() {
					diagnostics.AddError("Missing certificate source", "Configure blindfold or certificate_url at every certificate node.")
				}
				if !native.IsNull() && (!cert.IsNull() || (key.Type() != nil && !key.IsNull())) {
					diagnostics.AddError("Conflicting certificate sources", "Native blindfold conflicts with certificate_url and private_key.")
				}
				if !native.IsNull() && native.IsKnown() {
					inputs := objectValue(native)
					if !named && textValue(inputs["id"]) == "" {
						diagnostics.AddError("Missing inline certificate ID", "Native inline certificates require a stable unique id.")
					}
					wo := inputs["private_key_wo"].Type() != nil && !inputs["private_key_wo"].IsNull() || inputs["pkcs12_wo"].Type() != nil && !inputs["pkcs12_wo"].IsNull()
					if wo && textValue(inputs["material_version"]) == "" {
						diagnostics.AddError("Missing material version", "Write-only certificate material requires material_version.")
					}
				}
			}
			for k, child := range m {
				if k != "blindfold" {
					walk(child)
				}
			}
		case tftypes.List, tftypes.Set:
			var values []tftypes.Value
			_ = v.As(&values)
			for _, child := range values {
				walk(child)
			}
		}
	}
	walk(raw)
}
