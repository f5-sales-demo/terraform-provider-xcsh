package provider

import (
	"context"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/blindfold"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"
	"testing"
)

// Populate every certificate-bearing schema branch, including list ancestors,
// then verify projection and encrypted material at each of the discovered nodes.
func TestEveryBlindfoldNodeSerialization(t *testing.T) {
	c := context.Background()
	p := New("test")().(*XCSHProvider)
	total := 0
	for _, factory := range p.Resources(c) {
		r := factory()
		adapter, ok := r.(*blindfoldResource)
		if !ok {
			continue
		}
		var s resource.SchemaResponse
		r.Schema(c, resource.SchemaRequest{}, &s)
		var seed func(tftypes.Type) tftypes.Value
		seed = func(typ tftypes.Type) tftypes.Value {
			switch typ := typ.(type) {
			case tftypes.Object:
				values := map[string]tftypes.Value{}
				for k, child := range typ.AttributeTypes {
					values[k] = tftypes.NewValue(child, nil)
				}
				if bt, ok := typ.AttributeTypes["blindfold"]; ok {
					bv := map[string]tftypes.Value{}
					for k, t := range bt.(tftypes.Object).AttributeTypes {
						bv[k] = tftypes.NewValue(t, nil)
					}
					setString(bv, "id", "synthetic")
					values["blindfold"] = tftypes.NewValue(bt, bv)
				}
				for k, child := range typ.AttributeTypes {
					if k == "blindfold" {
						continue
					}
					switch child.(type) {
					case tftypes.Object, tftypes.List:
						values[k] = seed(child)
					}
				}
				return tftypes.NewValue(typ, values)
			case tftypes.List:
				return tftypes.NewValue(typ, []tftypes.Value{seed(typ.ElementType)})
			}
			return tftypes.NewValue(typ, nil)
		}
		raw := seed(s.Schema.Type().TerraformType(c))
		nodes := []*nativeNode{}
		walkNative(raw, nil, func(path []any, parent map[string]tftypes.Value) {
			nodes = append(nodes, &nativeNode{path: path, id: "synthetic", location: "string:///encrypted", material: &blindfold.Material{CertificateURL: "string:///public", ChainIdentity: strings.Repeat("a", 64), SPKIIdentity: strings.Repeat("b", 64)}, context: blindfold.Context{Digest: strings.Repeat("c", 64)}, identity: "public-plan-identity"})
		})
		total += len(nodes)
		out := nativeOutputs(raw, nodes, true)
		base := adapter.base(c)
		wire := injectNative(projectValue(out, base.Type().TerraformType(c)), nodes)
		walkNative(wire, nil, func(_ []any, _ map[string]tftypes.Value) { t.Fatal("provider input escaped API projection") })
		for _, n := range nodes {
			m := objectValue(atValue(wire, n.path))
			if textValue(m["certificate_url"]) != "string:///public" {
				t.Fatal("public location missing")
			}
			k := objectValue(m["private_key"])
			b := objectValue(k["blindfold_secret_info"])
			if textValue(b["location"]) != "string:///encrypted" {
				t.Fatal("encrypted location missing")
			}
			native := objectValue(objectValue(atValue(out, n.path))["blindfold"])
			for _, secret := range []string{"private_key_wo", "pkcs12_wo", "passphrase_wo"} {
				if !native[secret].IsNull() {
					t.Fatal("write-only value retained")
				}
			}
		}
	}
	if total != 25 {
		t.Fatalf("serialization covered %d nodes", total)
	}
}
