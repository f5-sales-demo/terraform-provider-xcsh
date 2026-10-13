package planmodifiers

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestDefaultOneOfMarker(t *testing.T) {
	empty := tftypes.Object{AttributeTypes: map[string]tftypes.Type{}}
	root := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"default": empty, "custom": empty}}
	for _, test := range []struct {
		name            string
		marker, sibling interface{}
		clear           bool
	}{
		{"explicit_sibling", nil, map[string]tftypes.Value{}, true},
		{"configured_default", map[string]tftypes.Value{}, nil, false},
		{"omitted_both", nil, nil, false},
		{"unknown_sibling", nil, tftypes.UnknownValue, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			config := tftypes.NewValue(root, map[string]tftypes.Value{"default": tftypes.NewValue(empty, test.marker), "custom": tftypes.NewValue(empty, test.sibling)})
			configValue := types.ObjectNull(nil)
			if test.marker != nil {
				configValue = types.ObjectValueMust(nil, nil)
			}
			plan := types.ObjectUnknown(nil)
			req := planmodifier.ObjectRequest{Config: tfsdk.Config{Raw: config}, ConfigValue: configValue, PlanValue: plan}
			resp := planmodifier.ObjectResponse{PlanValue: plan}
			DefaultOneOfMarker("default", []string{"default", "custom"}).PlanModifyObject(context.Background(), req, &resp)
			if resp.Diagnostics.HasError() {
				t.Fatal(resp.Diagnostics)
			}
			if resp.PlanValue.IsNull() != test.clear {
				t.Fatalf("clear=%v want %v", resp.PlanValue.IsNull(), test.clear)
			}
		})
	}
}

func TestDefaultOneOfMarkerNestedSiblingCutover(t *testing.T) {
	empty := tftypes.Object{AttributeTypes: map[string]tftypes.Type{}}
	settings := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"default": empty, "custom": empty}}
	root := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"settings": settings}}
	values := map[string]tftypes.Value{"default": tftypes.NewValue(empty, nil), "custom": tftypes.NewValue(empty, map[string]tftypes.Value{})}
	config := tftypes.NewValue(root, map[string]tftypes.Value{"settings": tftypes.NewValue(settings, values)})
	prior := types.ObjectValueMust(nil, nil)
	req := planmodifier.ObjectRequest{Path: path.Root("settings").AtName("default"), Config: tfsdk.Config{Raw: config}, ConfigValue: types.ObjectNull(nil), PlanValue: prior, StateValue: prior}
	resp := planmodifier.ObjectResponse{PlanValue: prior}
	DefaultOneOfMarker("default", []string{"default", "custom"}).PlanModifyObject(context.Background(), req, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatal(resp.Diagnostics)
	}
	if !resp.PlanValue.IsNull() {
		t.Fatal("prior nested default survived explicit sibling")
	}
}

func TestDefaultOneOfMarkerListNestedCutover(t *testing.T) {
	empty := tftypes.Object{AttributeTypes: map[string]tftypes.Type{}}
	settings := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"default": empty, "custom": empty}}
	root := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"settings": tftypes.List{ElementType: settings}}}
	for _, unknown := range []bool{false, true} {
		var sibling interface{} = map[string]tftypes.Value{}
		if unknown {
			sibling = tftypes.UnknownValue
		}
		values := map[string]tftypes.Value{"default": tftypes.NewValue(empty, nil), "custom": tftypes.NewValue(empty, sibling)}
		config := tftypes.NewValue(root, map[string]tftypes.Value{"settings": tftypes.NewValue(tftypes.List{ElementType: settings}, []tftypes.Value{tftypes.NewValue(settings, values)})})
		prior := types.ObjectValueMust(nil, nil)
		req := planmodifier.ObjectRequest{Path: path.Root("settings").AtListIndex(0).AtName("default"), Config: tfsdk.Config{Raw: config}, ConfigValue: types.ObjectNull(nil), PlanValue: prior, StateValue: prior}
		resp := planmodifier.ObjectResponse{PlanValue: prior}
		DefaultOneOfMarker("default", []string{"default", "custom"}).PlanModifyObject(context.Background(), req, &resp)
		if resp.Diagnostics.HasError() {
			t.Fatal(resp.Diagnostics)
		}
		if resp.PlanValue.IsNull() == unknown {
			t.Fatalf("unknown sibling=%v must only clear known selection", unknown)
		}
	}
}
