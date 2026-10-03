package planmodifiers

import (
	"context"
	"testing"

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
