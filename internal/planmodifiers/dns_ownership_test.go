package planmodifiers

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestDNSOwnershipRequiresReplaceTransitions(t *testing.T) {
	ctx := context.Background()
	s := schema.Schema{Attributes: map[string]schema.Attribute{"dns_volterra_managed": schema.BoolAttribute{}}}
	raw := tftypes.NewValue(s.Type().TerraformType(ctx), map[string]tftypes.Value{"dns_volterra_managed": tftypes.NewValue(tftypes.Bool, false)})
	for _, tc := range []struct {
		name          string
		before, after types.Bool
		replace       bool
	}{
		{"delegated to tenant", types.BoolValue(true), types.BoolValue(false), true},
		{"tenant to delegated", types.BoolValue(false), types.BoolValue(true), true},
		{"unchanged", types.BoolValue(false), types.BoolValue(false), false},
		{"both null", types.BoolNull(), types.BoolNull(), false},
		{"unknown differs", types.BoolValue(false), types.BoolUnknown(), true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req := planmodifier.BoolRequest{State: tfsdk.State{Schema: s, Raw: raw}, Plan: tfsdk.Plan{Schema: s, Raw: raw}, StateValue: tc.before, PlanValue: tc.after}
			resp := &planmodifier.BoolResponse{}
			boolplanmodifier.RequiresReplace().PlanModifyBool(ctx, req, resp)
			if resp.RequiresReplace != tc.replace {
				t.Fatalf("RequiresReplace=%v want %v", resp.RequiresReplace, tc.replace)
			}
		})
	}
}
