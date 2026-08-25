// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package schema

import (
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/namespace"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

// findAttr returns the attribute with the given tfsdk tag, or nil.
func findAttr(attrs []openapi.TerraformAttribute, tag string) *openapi.TerraformAttribute {
	for i := range attrs {
		if attrs[i].TfsdkTag == tag {
			return &attrs[i]
		}
	}
	return nil
}

// systemOnlySpec builds a minimal spec + namespaced path for a resource.
func systemOnlySpec(resourceName string) (*openapi.Spec, func(*openapi.Spec, string) (string, string, bool)) {
	spec := &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				resourceName + "CreateSpecType": {
					Type:       "object",
					Properties: map[string]openapi.Schema{"port": {Type: "integer"}},
				},
			},
		},
	}
	extractAPIPath := func(_ *openapi.Spec, _ string) (string, string, bool) {
		return "/api/config/dns/namespaces/%s/" + resourceName + "s", "/api/config/dns/namespaces/%s/" + resourceName + "s/%s", true
	}
	return spec, extractAPIPath
}

// A resource whose spec profile restricts it to a single namespace ("system")
// must emit namespace as Optional+Computed with a static default and a OneOf
// validator — not Required — so it can be omitted and cannot be set wrong.
func TestExtractResourceSchema_SingleAllowedNamespaceIsDefaulted(t *testing.T) {
	namespace.ClearProfiles()
	namespace.SetProfile("sys_res", namespace.Profile{Allowed: []namespace.NamespaceType{namespace.System}, Enforced: true})
	defer namespace.ClearProfiles()

	spec, extractAPIPath := systemOnlySpec("sys_res")
	result, err := ExtractResourceSchema(spec, "sys_res", extractAPIPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ns := findAttr(result.Attributes, "namespace")
	if ns == nil {
		t.Fatal("namespace attribute missing")
	}
	if ns.Required {
		t.Error("namespace must not be Required for a single-allowed-namespace resource")
	}
	if !ns.Optional || !ns.Computed {
		t.Errorf("namespace should be Optional+Computed, got Optional=%v Computed=%v", ns.Optional, ns.Computed)
	}
	if ns.StringDefault != "system" {
		t.Errorf("StringDefault = %q, want %q", ns.StringDefault, "system")
	}
	if len(ns.EnumValues) != 1 || ns.EnumValues[0] != "system" {
		t.Errorf("EnumValues = %v, want [system]", ns.EnumValues)
	}
	if !result.HasStringDefaults {
		t.Error("ResourceTemplate.HasStringDefaults should be true")
	}
}

// A resource whose spec profile allows multiple namespaces keeps namespace
// Required with no default — the user must choose.
func TestExtractResourceSchema_MultiAllowedNamespaceStaysRequired(t *testing.T) {
	namespace.ClearProfiles()
	namespace.SetProfile("app_res", namespace.Profile{Allowed: []namespace.NamespaceType{namespace.Custom, namespace.Default, namespace.Shared}})
	defer namespace.ClearProfiles()

	spec, extractAPIPath := systemOnlySpec("app_res")
	result, err := ExtractResourceSchema(spec, "app_res", extractAPIPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ns := findAttr(result.Attributes, "namespace")
	if ns == nil {
		t.Fatal("namespace attribute missing")
	}
	if !ns.Required {
		t.Error("namespace must stay Required for a multi-allowed-namespace resource")
	}
	if ns.StringDefault != "" {
		t.Errorf("StringDefault = %q, want empty", ns.StringDefault)
	}
	if len(ns.EnumValues) != 0 {
		t.Errorf("EnumValues = %v, want empty", ns.EnumValues)
	}
}

// With no registered profile, namespace behaviour is unchanged (Required when the
// API path is namespaced).
func TestExtractResourceSchema_NoProfileStaysRequired(t *testing.T) {
	namespace.ClearProfiles()
	spec, extractAPIPath := systemOnlySpec("unprofiled_res")
	result, err := ExtractResourceSchema(spec, "unprofiled_res", extractAPIPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ns := findAttr(result.Attributes, "namespace")
	if ns == nil {
		t.Fatal("namespace attribute missing")
	}
	if !ns.Required || ns.StringDefault != "" {
		t.Errorf("unprofiled namespace should be Required with no default; got Required=%v StringDefault=%q", ns.Required, ns.StringDefault)
	}
}

func TestExtractResourceSchema_MinimumConfigurationDoesNotPromoteOptional(t *testing.T) {
	spec, extractAPIPath := systemOnlySpec("min_config_probe")
	create := spec.Components.Schemas["min_config_probeCreateSpecType"]
	create.XF5XCMinimumConfiguration = map[string]interface{}{
		"required_fields": []interface{}{"spec.port"},
	}
	spec.Components.Schemas["min_config_probeCreateSpecType"] = create

	result, err := ExtractResourceSchema(spec, "min_config_probe", extractAPIPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	port := findAttr(result.Attributes, "port")
	if port == nil {
		t.Fatal("port attribute missing")
	}
	if port.Required || !port.Optional {
		t.Fatalf("minimum-configuration guidance must not change Terraform requiredness; got Required=%v Optional=%v", port.Required, port.Optional)
	}
}

func concurrencyTokenSchema(serverAssigned bool, operations ...string) openapi.Schema {
	return openapi.Schema{
		Type: "string",
		XF5XCConcurrencyToken: &openapi.ConcurrencyToken{
			ServerAssigned:   serverAssigned,
			EchoOnOperations: operations,
		},
	}
}

func concurrencyTokenSpec() (*openapi.Spec, func(*openapi.Spec, string) (string, string, bool)) {
	spec, extractAPIPath := systemOnlySpec("token_probe")
	spec.Components.Schemas["token_probeGetResponse"] = openapi.Schema{
		Type: "object",
		Properties: map[string]openapi.Schema{
			"resource_version": concurrencyTokenSchema(true, "replace"),
		},
	}
	spec.Components.Schemas["token_probeReplaceRequest"] = openapi.Schema{
		Type: "object",
		Properties: map[string]openapi.Schema{
			"resource_version": concurrencyTokenSchema(true, "replace"),
		},
	}
	return spec, extractAPIPath
}

func TestExtractResourceSchema_ConcurrencyTokenContract(t *testing.T) {
	spec, extractAPIPath := concurrencyTokenSpec()
	result, err := ExtractResourceSchema(spec, "token_probe", extractAPIPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.HasConcurrencyToken {
		t.Fatal("valid GET/replace concurrency contract was not detected")
	}
	if result.ConcurrencyTokenJSONName != "resource_version" || result.ConcurrencyTokenGoName != "ResourceVersion" {
		t.Fatalf("unexpected token names: JSON=%q Go=%q", result.ConcurrencyTokenJSONName, result.ConcurrencyTokenGoName)
	}
	if findAttr(result.Attributes, "resource_version") != nil {
		t.Fatal("concurrency token must remain client-only and absent from Terraform attributes")
	}
}

func TestExtractResourceSchema_RejectsInvalidConcurrencyTokenContract(t *testing.T) {
	tests := map[string]func(*openapi.Spec){
		"missing replace token": func(spec *openapi.Spec) {
			replace := spec.Components.Schemas["token_probeReplaceRequest"]
			delete(replace.Properties, "resource_version")
			spec.Components.Schemas["token_probeReplaceRequest"] = replace
		},
		"mismatched property": func(spec *openapi.Spec) {
			replace := spec.Components.Schemas["token_probeReplaceRequest"]
			delete(replace.Properties, "resource_version")
			replace.Properties["etag"] = concurrencyTokenSchema(true, "replace")
			spec.Components.Schemas["token_probeReplaceRequest"] = replace
		},
		"non string": func(spec *openapi.Spec) {
			get := spec.Components.Schemas["token_probeGetResponse"]
			field := get.Properties["resource_version"]
			field.Type = "integer"
			get.Properties["resource_version"] = field
			spec.Components.Schemas["token_probeGetResponse"] = get
		},
		"not server assigned": func(spec *openapi.Spec) {
			get := spec.Components.Schemas["token_probeGetResponse"]
			get.Properties["resource_version"] = concurrencyTokenSchema(false, "replace")
			spec.Components.Schemas["token_probeGetResponse"] = get
		},
		"replace not echoed": func(spec *openapi.Spec) {
			replace := spec.Components.Schemas["token_probeReplaceRequest"]
			replace.Properties["resource_version"] = concurrencyTokenSchema(true, "update")
			spec.Components.Schemas["token_probeReplaceRequest"] = replace
		},
		"create declares token": func(spec *openapi.Spec) {
			create := spec.Components.Schemas["token_probeCreateRequest"]
			create.Type = "object"
			create.Properties = map[string]openapi.Schema{"resource_version": concurrencyTokenSchema(true, "replace")}
			spec.Components.Schemas["token_probeCreateRequest"] = create
		},
	}

	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			spec, extractAPIPath := concurrencyTokenSpec()
			mutate(spec)
			if _, err := ExtractResourceSchema(spec, "token_probe", extractAPIPath); err == nil {
				t.Fatal("expected invalid concurrency-token contract to fail extraction")
			}
		})
	}
}

// A single-allowed profile that is NOT enforced (unverified classification) must not
// be defaulted/locked — namespace stays Required so we don't over-restrict on a guess.
func TestExtractResourceSchema_UnverifiedSingleAllowedStaysRequired(t *testing.T) {
	namespace.ClearProfiles()
	namespace.SetProfile("unverified_sys", namespace.Profile{
		Allowed:  []namespace.NamespaceType{namespace.System},
		Enforced: false,
	})
	defer namespace.ClearProfiles()

	spec, extractAPIPath := systemOnlySpec("unverified_sys")
	result, err := ExtractResourceSchema(spec, "unverified_sys", extractAPIPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ns := findAttr(result.Attributes, "namespace")
	if ns == nil {
		t.Fatal("namespace attribute missing")
	}
	if !ns.Required {
		t.Error("unverified (enforced=false) single-allowed namespace must stay Required")
	}
	if ns.StringDefault != "" || len(ns.EnumValues) != 0 {
		t.Errorf("unverified namespace must not be defaulted/locked; got StringDefault=%q EnumValues=%v", ns.StringDefault, ns.EnumValues)
	}
}

// tokenLikeSpec builds a spec shaped like the F5 XC token domain: an (empty)
// CreateSpecType plus a GetResponse envelope whose system_metadata is an
// allOf-$ref to a shared system-metadata type carrying uid.
func tokenLikeSpec(resourceName string, withSystemMetadataUID bool) (*openapi.Spec, func(*openapi.Spec, string) (string, string, bool)) {
	schemas := map[string]openapi.Schema{
		resourceName + "CreateSpecType": {
			Type:       "object",
			Properties: map[string]openapi.Schema{},
		},
		resourceName + "GetSpecType": {
			Type:       "object",
			Properties: map[string]openapi.Schema{},
		},
		"schemaSystemObjectGetMetaType": {
			Type: "object",
			Properties: map[string]openapi.Schema{
				"uid":    {Type: "string"},
				"tenant": {Type: "string"},
			},
		},
	}
	getResp := openapi.Schema{Type: "object", Properties: map[string]openapi.Schema{
		"metadata": {Type: "object"},
		"spec":     {Type: "object"},
	}}
	if withSystemMetadataUID {
		getResp.Properties["system_metadata"] = openapi.Schema{
			AllOf: []openapi.Schema{{Ref: "#/components/schemas/schemaSystemObjectGetMetaType"}},
		}
	}
	schemas[resourceName+"GetResponse"] = getResp
	spec := &openapi.Spec{Components: openapi.Components{Schemas: schemas}}
	extractAPIPath := func(_ *openapi.Spec, _ string) (string, string, bool) {
		return "/api/register/namespaces/%s/" + resourceName + "s",
			"/api/register/namespaces/%s/" + resourceName + "s/%s", true
	}
	return spec, extractAPIPath
}

// ResponseHasSystemMetadataUID resolves the allOf-wrapped system_metadata ref and
// reports whether it carries uid.
func TestResponseHasSystemMetadataUID(t *testing.T) {
	specYes, _ := tokenLikeSpec("token", true)
	if !ResponseHasSystemMetadataUID(specYes, "token") {
		t.Error("expected true when GetResponse.system_metadata carries uid")
	}
	specNo, _ := tokenLikeSpec("token", false)
	if ResponseHasSystemMetadataUID(specNo, "token") {
		t.Error("expected false when the response envelope has no system_metadata")
	}
}

// The token resource (opted in via tools/expose-uid.json AND schema-verified)
// gains a Computed, read-only `uid` attribute and ExposeUID=true; the uid is not
// a spec field. A resource NOT opted in is unchanged (no uid, ExposeUID=false).
func TestExtractResourceSchema_ExposesUIDForToken(t *testing.T) {
	namespace.ClearProfiles()
	defer namespace.ClearProfiles()

	spec, extractAPIPath := tokenLikeSpec("token", true)
	result, err := ExtractResourceSchema(spec, "token", extractAPIPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.ExposeUID {
		t.Error("ResourceTemplate.ExposeUID should be true for token")
	}
	uid := findAttr(result.Attributes, "uid")
	if uid == nil {
		t.Fatal("uid attribute missing")
	}
	if !uid.Computed || uid.Required || uid.Optional {
		t.Errorf("uid must be Computed-only, got Computed=%v Required=%v Optional=%v", uid.Computed, uid.Required, uid.Optional)
	}
	if uid.IsSpecField {
		t.Error("uid must not be a spec field (excluded from spec marshal/unmarshal)")
	}
	if uid.GoName != "Uid" || uid.Type != "string" {
		t.Errorf("uid GoName/Type = %q/%q, want Uid/string", uid.GoName, uid.Type)
	}
	if !uid.Sensitive {
		t.Error("uid must be marked Sensitive for token")
	}
	if !strings.Contains(uid.Description, "plain text in the Terraform state file") {
		t.Errorf("uid description must warn about state file storage, got %q", uid.Description)
	}
}

func TestExtractReadOnlyResourceSchema_ExposesUIDForToken(t *testing.T) {
	namespace.ClearProfiles()
	defer namespace.ClearProfiles()

	spec, extractAPIPath := tokenLikeSpec("token", true)
	result, err := ExtractReadOnlyResourceSchema(spec, "token", extractAPIPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.ExposeUID {
		t.Error("DataSourceTemplate.ExposeUID should be true for token")
	}
	uid := findAttr(result.Attributes, "uid")
	if uid == nil {
		t.Fatal("uid attribute missing")
	}
	if !uid.Computed || uid.Required || uid.Optional {
		t.Errorf("uid must be Computed-only, got Computed=%v Required=%v Optional=%v", uid.Computed, uid.Required, uid.Optional)
	}
	if !uid.Sensitive {
		t.Error("uid must be marked Sensitive for token data source")
	}
	if !strings.Contains(uid.Description, "plain text in the Terraform state file") {
		t.Errorf("uid description must warn about state file storage, got %q", uid.Description)
	}
}

// A resource that is NOT opted in gets no uid attribute even if its response
// schema carries system_metadata.uid — scoping is opt-in, not schema-wide.
func TestExtractResourceSchema_NoUIDWhenNotOptedIn(t *testing.T) {
	namespace.ClearProfiles()
	defer namespace.ClearProfiles()

	spec, extractAPIPath := tokenLikeSpec("not_opted_in_res", true)
	result, err := ExtractResourceSchema(spec, "not_opted_in_res", extractAPIPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ExposeUID {
		t.Error("ExposeUID must be false for a resource not in expose-uid.json")
	}
	if findAttr(result.Attributes, "uid") != nil {
		t.Error("non-opted-in resource must not gain a uid attribute")
	}
}

// actionApproveSpec builds a spec shaped like the F5 XC registration-approval
// action: a request-body component schema carrying a schema-level x-f5xc-action
// marker, the raw POST action path that $refs it, and a sibling plural GET path
// for the object being acted on. ResourcePath carries the exact operation-catalog
// resolution supplied to the schema extractor.
func actionApproveSpec() (*openapi.Spec, openapi.ResourcePath) {
	spec := &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				// Mirror the real spec: the request-body component is camelCase
				// ("registrationApprovalReq") and its properties mix scalar strings
				// with object props (annotations, labels) and non-enum $ref props
				// (passport, tunnel_type). `namespace` is a PATH parameter, not a
				// body property, so it is intentionally absent here. `state` is a
				// $ref to a string enum.
				"registrationApprovalReq": {
					Type:        "object",
					XF5xcAction: "approve",
					Properties: map[string]openapi.Schema{
						"name":                    {Type: "string"},
						"backup_connected_region": {Type: "string"},
						"connected_region":        {Type: "string"},
						"preferred_active_re":     {Type: "string"},
						"annotations":             {Type: "object"},
						"labels":                  {Type: "object"},
						"passport":                {AllOf: []openapi.Schema{{Ref: "#/components/schemas/registrationPassport"}}},
						"tunnel_type":             {AllOf: []openapi.Schema{{Ref: "#/components/schemas/schemaSiteToSiteTunnelType"}}},
						"state":                   {Ref: "#/components/schemas/registrationObjectState"},
					},
				},
			},
		},
		Paths: map[string]interface{}{
			"/api/register/namespaces/{namespace}/registration/{name}/approve": map[string]interface{}{
				"post": map[string]interface{}{
					"requestBody": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/registrationApprovalReq",
								},
							},
						},
					},
				},
			},
			"/api/register/namespaces/{namespace}/registrations/{name}": map[string]interface{}{
				"get": map[string]interface{}{},
			},
		},
	}
	action := openapi.ResourcePath{
		ResourceName:   "registration_approval",
		SchemaName:     "registrationApprovalReq",
		ActionValue:    "approve",
		ActionPath:     "/api/register/namespaces/%s/registration/%s/approve",
		ReadObjectPath: "/api/register/namespaces/%s/registrations/%s",
	}
	return spec, action
}

// A schema-level x-f5xc-action marker drives an action-style resource: attributes
// derive from the request body, the singular action POST path and the pluralized
// sibling GET path are captured, state constant-defaults to APPROVED, and every
// user-settable attribute forces replace (there is no in-place update).
func TestActionResourceApprove(t *testing.T) {
	spec, action := actionApproveSpec()

	result, err := ExtractActionResourceSchema(spec, action)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsAction {
		t.Error("IsAction should be true")
	}
	if result.ActionPath != "/api/register/namespaces/%s/registration/%s/approve" {
		t.Errorf("ActionPath = %q, want the singular approve path", result.ActionPath)
	}
	if result.ReadObjectPath != "/api/register/namespaces/%s/registrations/%s" {
		t.Errorf("ReadObjectPath = %q, want the pluralized sibling GET path", result.ReadObjectPath)
	}
	if result.ActionState != "APPROVED" {
		t.Errorf("ActionState = %q, want APPROVED", result.ActionState)
	}
	if len(result.Attributes) == 0 {
		t.Fatal("expected attributes derived from the request body")
	}
	for _, a := range result.Attributes {
		if a.PlanModifier != "RequiresReplace" {
			t.Errorf("attr %q: PlanModifier = %q, want RequiresReplace", a.TfsdkTag, a.PlanModifier)
		}
	}
	if st := findAttr(result.Attributes, "state"); st == nil {
		t.Error("state attribute missing")
	} else if st.StringDefault != "APPROVED" {
		t.Errorf("state StringDefault = %q, want APPROVED", st.StringDefault)
	} else if !(st.Optional && st.Computed) {
		t.Error("state must be Optional+Computed")
	}

	// namespace is a path parameter (not a body property); it must be injected
	// as a Required attribute so the generated model carries a Namespace field.
	if ns := findAttr(result.Attributes, "namespace"); ns == nil {
		t.Error("namespace attribute missing (must be injected from the path param)")
	} else if !ns.Required {
		t.Error("namespace must be Required")
	}

	// name is a path parameter; the request schema's required list is null, so it
	// must be forced Required.
	if nm := findAttr(result.Attributes, "name"); nm == nil {
		t.Error("name attribute missing")
	} else if !nm.Required {
		t.Error("name must be Required")
	}

	// Scalar string metadata props are emitted as optional attributes.
	if bc := findAttr(result.Attributes, "backup_connected_region"); bc == nil {
		t.Error("backup_connected_region attribute missing")
	} else if !bc.Optional {
		t.Error("backup_connected_region must be Optional")
	}

	// Object props (annotations, labels) and non-enum $ref props (passport,
	// tunnel_type) are not representable as plain string attributes and must be
	// excluded entirely.
	for _, excluded := range []string{"annotations", "labels", "passport", "tunnel_type"} {
		if a := findAttr(result.Attributes, excluded); a != nil {
			t.Errorf("attribute %q must be excluded (object/$ref prop), but it was emitted", excluded)
		}
	}

	// The full expected attribute set is exactly these four.
	wantAttrs := map[string]bool{
		"namespace": true, "name": true, "state": true,
		"backup_connected_region": true, "connected_region": true, "preferred_active_re": true,
	}
	for _, a := range result.Attributes {
		if !wantAttrs[a.TfsdkTag] {
			t.Errorf("unexpected attribute %q in action model", a.TfsdkTag)
		}
	}
}

// The approve API rejects a request that omits the registration's passport
// ("Validation approval: Passport is required", HTTP 500 — #1355). The passport
// is a fact about the registration, not user input: F5 accepts only the value it
// already holds. So it must be extracted as a SERVER-DERIVED field — carried on
// the request body and read back off the sibling object — and must NOT appear as
// a user-settable Terraform attribute.
func TestActionResourceDerivesPassportFromTheRegistration(t *testing.T) {
	spec, action := actionApproveSpec()

	result, err := ExtractActionResourceSchema(spec, action)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var passport *openapi.ActionDerivedField
	for i := range result.ActionDerivedFields {
		if result.ActionDerivedFields[i].Field == "passport" {
			passport = &result.ActionDerivedFields[i]
		}
	}
	if passport == nil {
		t.Fatalf("passport must be a server-derived action field; got %+v", result.ActionDerivedFields)
	}
	if passport.GoName != "Passport" {
		t.Errorf("passport GoName = %q, want Passport", passport.GoName)
	}
	if passport.JSONName != "passport" {
		t.Errorf("passport JSONName = %q, want passport", passport.JSONName)
	}
	if len(passport.Sources) == 0 {
		t.Fatal("passport must declare at least one lookup path into the sibling registration object")
	}
	// Every declared source must address the passport of the registration read.
	for _, src := range passport.Sources {
		if !strings.HasSuffix(src, ".passport") {
			t.Errorf("passport source %q must address the object's passport field", src)
		}
	}

	// Server-derived means NOT user input: no Terraform attribute, so no
	// practitioner can supply a passport the API would reject.
	if a := findAttr(result.Attributes, "passport"); a != nil {
		t.Errorf("passport must not be a user-settable attribute; got %+v", a)
	}
}

// A property the action request REQUIRES must never be silently discarded by the
// attribute extractor. Dropping one is precisely how #1355 shipped: the filter
// keyed on the Go type ("not a string, skip") rather than on whether the API
// needs the field, so a required object vanished and every approve POSTed a 500.
// The extractor now fails generation instead — loudly, at build time — unless the
// required property is either a Terraform attribute or a declared server-derived
// field.
func TestActionResourceRequiredPropertyIsNeverSilentlyDropped(t *testing.T) {
	// A required non-string property that is neither an attribute nor declared
	// server-derived: generation must fail.
	spec := actionGuardSpec(map[string]openapi.Schema{
		"name": {Type: "string"},
		"credentials": {
			AllOf: []openapi.Schema{{Ref: "#/components/schemas/someCredentials"}},
		},
	}, []string{"name", "credentials"})

	action := actionGuardPath()
	_, err := ExtractActionResourceSchema(spec, action)
	if err == nil {
		t.Fatal("expected an error: the required object property 'credentials' is not representable as an attribute and is not declared server-derived")
	}
	if !strings.Contains(err.Error(), "credentials") {
		t.Errorf("error must name the dropped required property; got: %v", err)
	}

	// Positive control: with the same shape but the property NOT required, the
	// existing skip is correct and generation succeeds. Without this the guard
	// could be a blanket failure and the test above would be unfalsifiable.
	ok := actionGuardSpec(map[string]openapi.Schema{
		"name": {Type: "string"},
		"credentials": {
			AllOf: []openapi.Schema{{Ref: "#/components/schemas/someCredentials"}},
		},
	}, []string{"name"})
	if _, err := ExtractActionResourceSchema(ok, action); err != nil {
		t.Fatalf("a non-required object property must still be skipped without error; got: %v", err)
	}

	// The spec also marks required with x-f5xc-required-for.create and
	// x-ves-required rather than a `required` list, so both must trip the guard.
	forCreate := actionGuardSpec(map[string]openapi.Schema{
		"name": {Type: "string"},
		"credentials": {
			AllOf:            []openapi.Schema{{Ref: "#/components/schemas/someCredentials"}},
			XF5XCRequiredFor: openapi.RequiredFor{Create: true},
		},
	}, nil)
	if _, err := ExtractActionResourceSchema(forCreate, action); err == nil {
		t.Error("x-f5xc-required-for.create=true must trip the dropped-required-property guard")
	}

	vesRequired := actionGuardSpec(map[string]openapi.Schema{
		"name": {Type: "string"},
		"credentials": {
			AllOf:        []openapi.Schema{{Ref: "#/components/schemas/someCredentials"}},
			XVesRequired: "true",
		},
	}, nil)
	if _, err := ExtractActionResourceSchema(vesRequired, action); err == nil {
		t.Error("x-ves-required=true must trip the dropped-required-property guard")
	}
}

// The real registration approve request marks nothing required, so the guard must
// not fire on it — the passport is covered by its server-derived declaration.
func TestActionResourceGuardAcceptsTheRealApproveSchema(t *testing.T) {
	spec, action := actionApproveSpec()
	if _, err := ExtractActionResourceSchema(spec, action); err != nil {
		t.Fatalf("the real approve schema must extract cleanly; got: %v", err)
	}
}

// actionGuardSpec builds a minimal action spec with the supplied request-body
// properties and `required` list, under a synthetic resource name that can never
// collide with a real one.
func actionGuardSpec(props map[string]openapi.Schema, required []string) *openapi.Spec {
	return &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				"zzGuardProbeReq": {
					Type:        "object",
					XF5xcAction: "approve",
					Required:    required,
					Properties:  props,
				},
			},
		},
		Paths: map[string]interface{}{
			"/api/register/namespaces/{namespace}/guard_probe/{name}/approve": map[string]interface{}{
				"post": map[string]interface{}{
					"requestBody": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/zzGuardProbeReq",
								},
							},
						},
					},
				},
			},
		},
	}
}

func actionGuardPath() openapi.ResourcePath {
	return openapi.ResourcePath{
		ResourceName:   "zz_guard_probe",
		SchemaName:     "zzGuardProbeReq",
		ActionValue:    "approve",
		ActionPath:     "/api/register/namespaces/%s/guard_probe/%s/approve",
		ReadObjectPath: "/api/register/namespaces/%s/guard_probes/%s",
	}
}

func TestExtractOneOfGroups_Empty(t *testing.T) {
	spec := &openapi.Spec{}
	result := ExtractOneOfGroups(spec, "NonExistent")
	if len(result) != 0 {
		t.Errorf("Expected empty map, got %d groups", len(result))
	}
}

func TestExtractOneOfGroups_FromCache(t *testing.T) {
	spec := &openapi.Spec{}

	// Put raw schema data into cache
	RawSpecCache["TestType"] = map[string]interface{}{
		"x-ves-oneof-field-choice": []interface{}{"field_a", "field_b"},
		"type":                     "object",
	}
	defer delete(RawSpecCache, "TestType")

	result := ExtractOneOfGroups(spec, "TestType")
	if len(result) != 1 {
		t.Fatalf("Expected 1 group, got %d", len(result))
	}
	fields, ok := result["choice"]
	if !ok {
		t.Fatal("Expected group 'choice'")
	}
	if len(fields) != 2 {
		t.Fatalf("Expected 2 fields in group, got %d", len(fields))
	}
	if fields[0] != "field_a" || fields[1] != "field_b" {
		t.Errorf("Unexpected fields: %v", fields)
	}
}

func TestExtractOneOfGroups_StringFormat(t *testing.T) {
	spec := &openapi.Spec{}

	RawSpecCache["StringType"] = map[string]interface{}{
		"x-ves-oneof-field-mode": `["fast","slow"]`,
	}
	defer delete(RawSpecCache, "StringType")

	result := ExtractOneOfGroups(spec, "StringType")
	if len(result) != 1 {
		t.Fatalf("Expected 1 group, got %d", len(result))
	}
	fields := result["mode"]
	if len(fields) != 2 {
		t.Fatalf("Expected 2 fields, got %d", len(fields))
	}
}

func TestExtractResourceSchema_NoCreateSpecType(t *testing.T) {
	spec := &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				"SomeOtherType": {Type: "object"},
			},
		},
	}
	extractAPIPath := func(spec *openapi.Spec, resourceName string) (string, string, bool) {
		return "/api/test", "/api/test/{name}", true
	}
	_, err := ExtractResourceSchema(spec, "test_resource", extractAPIPath)
	if err == nil {
		t.Error("Expected error when no CreateSpecType found")
	}
}

func TestExtractResourceSchema_SchemaPrefixMatch(t *testing.T) {
	// Specs like fast_acl use "schemafast_aclCreateSpecType" (with "schema" prefix).
	// The generator must match this before a shorter name like "fast_acl_ruleCreateSpecType".
	spec := &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				"fast_acl_ruleCreateSpecType": {
					Type:       "object",
					Properties: map[string]openapi.Schema{"action": {Type: "string"}},
				},
				"schemafast_aclCreateSpecType": {
					Type:       "object",
					Properties: map[string]openapi.Schema{"re_acl": {Type: "object"}, "site_acl": {Type: "object"}},
				},
			},
		},
	}
	extractAPIPath := func(spec *openapi.Spec, resourceName string) (string, string, bool) {
		return "/api/config/namespaces/%s/fast_acls", "/api/config/namespaces/%s/fast_acls/%s", true
	}
	result, err := ExtractResourceSchema(spec, "fast_acl", extractAPIPath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	foundReAcl := false
	for _, attr := range result.Attributes {
		if attr.Name == "re_acl" {
			foundReAcl = true
		}
	}
	if !foundReAcl {
		t.Error("Expected 're_acl' attribute (from schemafast_aclCreateSpecType), got fast_acl_rule schema instead")
	}
}

func TestExtractResourceSchema_Basic(t *testing.T) {
	spec := &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				"ves.io.schema.my_resource.CreateSpecType": {
					Type:        "object",
					Description: "A test resource",
					Properties: map[string]openapi.Schema{
						"port": {Type: "integer", Description: "Port number"},
					},
				},
			},
		},
	}
	extractAPIPath := func(spec *openapi.Spec, resourceName string) (string, string, bool) {
		return "/api/config/namespaces/%s/my_resources", "/api/config/namespaces/%s/my_resources/%s", true
	}
	result, err := ExtractResourceSchema(spec, "my_resource", extractAPIPath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.Name != "my_resource" {
		t.Errorf("Name = %q, want %q", result.Name, "my_resource")
	}
	if result.HasNamespaceInPath != true {
		t.Error("HasNamespaceInPath should be true")
	}
	// Should have standard attrs (name, namespace, annotations, description, disable, labels, id) + port
	foundPort := false
	for _, attr := range result.Attributes {
		if attr.Name == "port" {
			foundPort = true
		}
	}
	if !foundPort {
		t.Error("Expected 'port' attribute in result")
	}
}

// A resource whose CreateSpecType property declares a structured cross-resource
// requirement (x-f5xc-requires.requires_resource) auto-derives an apply-time
// preflight: the target resource's LIST path is resolved via extractAPIPath and
// the trigger field is bound to its Go model field. This is the end-to-end wiring
// that makes preflight-requirements.json an override rather than the sole source.
func TestExtractResourceSchema_AutoDerivesPreflightFromRequires(t *testing.T) {
	namespace.ClearProfiles()
	defer namespace.ClearProfiles()

	spec := &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				"app_lbCreateSpecType": {
					Type: "object",
					Properties: map[string]openapi.Schema{
						"client_side_defense": {
							Type: "boolean",
							XF5XCRequires: []openapi.RequiresEntry{{
								Field:            "client_side_defense",
								RequiresResource: &openapi.RequiresResource{Resource: "protected_domain", Scope: "same_namespace"},
								Reason:           "CSD needs a protected_domain in the same namespace.",
							}},
						},
					},
				},
			},
		},
	}
	extractAPIPath := func(_ *openapi.Spec, resource string) (string, string, bool) {
		if resource == "protected_domain" {
			return "/api/shape/csd/namespaces/%s/protected_domains",
				"/api/shape/csd/namespaces/%s/protected_domains/%s", true
		}
		return "/api/config/namespaces/%s/app_lbs", "/api/config/namespaces/%s/app_lbs/%s", true
	}

	result, err := ExtractResourceSchema(spec, "app_lb", extractAPIPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var found *openapi.RequirementPreflight
	for i := range result.Preflights {
		if result.Preflights[i].WhenField == "client_side_defense" {
			found = &result.Preflights[i]
		}
	}
	if found == nil {
		t.Fatalf("expected a derived preflight for client_side_defense, got %+v", result.Preflights)
	}
	if found.ListPath != "/api/shape/csd/namespaces/%s/protected_domains" {
		t.Errorf("ListPath = %q, want the resolved protected_domains shape path", found.ListPath)
	}
	if found.WhenGoField == "" {
		t.Error("WhenGoField must be resolved to the Go model field, got empty")
	}
}

// TestForceReplaceForCreateDeleteOnly verifies that create/delete-only resources (no PUT/update
// endpoint on the F5 XC API) get RequiresReplace on every user-settable field, so any change
// forces delete+create instead of a phantom in-place update that 404s.
func TestForceReplaceForCreateDeleteOnly(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{Name: "name", TfsdkTag: "name", Required: true, PlanModifier: "RequiresReplace"},
		{Name: "namespace", TfsdkTag: "namespace", Required: true, PlanModifier: "RequiresReplace"},
		{Name: "mitigated_domain", TfsdkTag: "mitigated_domain", Type: "string", Required: true},
		{Name: "description", TfsdkTag: "description", Type: "string", Optional: true},
		{Name: "disable", TfsdkTag: "disable", Type: "bool", Optional: true},
		{Name: "labels", TfsdkTag: "labels", Type: "map", Optional: true},
		{Name: "annotations", TfsdkTag: "annotations", Type: "map", Optional: true},
		{Name: "id", TfsdkTag: "id", Type: "string", Computed: true, PlanModifier: "UseStateForUnknown"},
	}
	ForceReplaceForCreateDeleteOnly(attrs)

	for _, tag := range []string{"mitigated_domain", "description", "disable", "labels", "annotations", "name", "namespace"} {
		a := findAttr(attrs, tag)
		if a == nil {
			t.Fatalf("attr %q missing", tag)
		}
		if a.PlanModifier != "RequiresReplace" {
			t.Errorf("attr %q: PlanModifier = %q, want RequiresReplace", tag, a.PlanModifier)
		}
	}
	// Computed id must NOT be forced to replace.
	if id := findAttr(attrs, "id"); id == nil || id.PlanModifier != "UseStateForUnknown" {
		t.Errorf("computed id must keep UseStateForUnknown, got %v", id)
	}
}
