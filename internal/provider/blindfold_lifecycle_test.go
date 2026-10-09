package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
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
