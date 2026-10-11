package provider

import (
	"context"
	"encoding/json"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/blindfold"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"
	"testing"
)

func countNative(a map[string]schema.Attribute, b map[string]schema.Block) int {
	n := 0
	if _, ok := a["blindfold"]; ok {
		n++
	}
	for _, v := range b {
		switch x := v.(type) {
		case schema.SingleNestedBlock:
			n += countNative(x.Attributes, x.Blocks)
		case schema.ListNestedBlock:
			n += countNative(x.NestedObject.Attributes, x.NestedObject.Blocks)
		}
	}
	return n
}
func TestBlindfoldInventorySchema(t *testing.T) {
	c := context.Background()
	p := New("test")().(*XCSHProvider)
	nodes, resources := 0, 0
	for _, factory := range p.Resources(c) {
		r := factory()
		var s resource.SchemaResponse
		r.Schema(c, resource.SchemaRequest{}, &s)
		if ds := s.Schema.ValidateImplementation(c); ds.HasError() {
			t.Fatalf("invalid schema: %v", ds)
		}
		if n := countNative(s.Schema.Attributes, s.Schema.Blocks); n > 0 {
			nodes += n
			resources++
		}
	}
	if nodes != 25 || resources != 14 {
		t.Fatalf("inventory: %d nodes across %d resources", nodes, resources)
	}
}
func TestBlindfoldProjection(t *testing.T) {
	s := map[string]tftypes.Type{"certificate_url": tftypes.String, "blindfold": tftypes.Object{AttributeTypes: map[string]tftypes.Type{"private_key_wo": tftypes.String}}}
	raw := tftypes.NewValue(tftypes.Object{AttributeTypes: s}, map[string]tftypes.Value{"certificate_url": tftypes.NewValue(tftypes.String, "public"), "blindfold": tftypes.NewValue(s["blindfold"], map[string]tftypes.Value{"private_key_wo": tftypes.NewValue(tftypes.String, "private-marker")})})
	base := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"certificate_url": tftypes.String}}
	projected := projectValue(raw, base)
	if _, ok := objectValue(projected)["blindfold"]; ok {
		t.Fatal("provider input escaped projection")
	}
	merged := mergeNative(projected, raw)
	if textValue(objectValue(merged)["certificate_url"]) != "public" {
		t.Fatal("public data lost")
	}
}

func TestBlindfoldTLSProtocolDefaults(t *testing.T) {
	ctx := context.Background()
	var response resource.SchemaResponse
	r := NewClusterResource()
	r.Schema(ctx, resource.SchemaRequest{}, &response)
	tls := response.Schema.Blocks["tls_parameters"].(schema.SingleNestedBlock)
	common := tls.Blocks["common_params"].(schema.SingleNestedBlock)
	for _, key := range []string{"minimum_protocol_version", "maximum_protocol_version"} {
		v := common.Attributes[key].(schema.StringAttribute)
		if !v.Optional || !v.Computed {
			t.Fatalf("%s must accept observed API default", key)
		}
	}
}

func TestBlindfoldAnnotationRemovesObsoleteIDs(t *testing.T) {
	annotationType := tftypes.Map{ElementType: tftypes.String}
	rawType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"annotations": annotationType}}
	raw := tftypes.NewValue(rawType, map[string]tftypes.Value{"annotations": tftypes.NewValue(annotationType, map[string]tftypes.Value{"example.test/keep": tftypes.NewValue(tftypes.String, "configured")})})
	provenance := blindfold.Provenance{Version: 1, Entries: map[string]blindfold.Entry{
		"retained": {Chain: strings.Repeat("a", 64), SPKI: strings.Repeat("b", 64), Context: strings.Repeat("c", 64), Ciphertext: strings.Repeat("d", 64), Algorithm: "RSA"}, "removed": {Chain: strings.Repeat("a", 64), SPKI: strings.Repeat("b", 64), Context: strings.Repeat("c", 64), Ciphertext: strings.Repeat("e", 64), Algorithm: "RSA"},
	}}
	serialized, _ := json.Marshal(provenance)
	remote := map[string]any{"metadata": map[string]any{"annotations": map[string]any{blindfold.Annotation: string(serialized), "example.test/remote": "remote"}}}
	updated := annotationRaw(raw, []*nativeNode{{id: "retained"}}, remote)
	attrs := map[string]tftypes.Value{}
	if err := objectValue(updated)["annotations"].As(&attrs); err != nil {
		t.Fatal(err)
	}
	var actual blindfold.Provenance
	if err := json.Unmarshal([]byte(textValue(attrs[blindfold.Annotation])), &actual); err != nil {
		t.Fatal(err)
	}
	if len(actual.Entries) != 1 || actual.Entries["retained"] != provenance.Entries["retained"] {
		t.Fatalf("obsolete ID retained or unchanged write-only entry lost: %+v", actual)
	}
	if textValue(attrs["example.test/keep"]) != "configured" || textValue(attrs["example.test/remote"]) != "remote" {
		t.Fatal("unrelated annotations changed")
	}
}
