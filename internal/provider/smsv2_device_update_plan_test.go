// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"testing"
)

func TestSMSv2DeviceEditPlansUpdateForSingleNodeNonHA(t *testing.T) {
	ctx := context.Background()
	before := awsSMSv2ContractFixture(t, []contractInterface{{mac: "02:00:00:00:00:01", role: "slo"}, {mac: "02:00:00:00:00:02", role: "sli"}})
	after := awsSMSv2ContractFixture(t, []contractInterface{{mac: "02:00:00:00:00:01", role: "slo"}, {mac: "02:00:00:00:00:02", role: "sli", device: "ens6"}})
	before.DisableHA = &SecuremeshSiteV2EmptyModel{}
	after.DisableHA = &SecuremeshSiteV2EmptyModel{}
	r := &SecuremeshSiteV2Resource{}
	schema := resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, &schema)
	raw := func(model SecuremeshSiteV2ResourceModel) tftypes.Value {
		rootType := schema.Schema.Type().(types.ObjectType)
		attrs := map[string]attr.Value{}
		for key, typ := range rootType.AttrTypes {
			value, err := typ.ValueFromTerraform(ctx, tftypes.NewValue(typ.TerraformType(ctx), nil))
			if err != nil {
				t.Fatal(err)
			}
			attrs[key] = value
		}
		var aws types.Object
		if diags := tfsdk.ValueFrom(ctx, model.AWS, rootType.AttrTypes["aws"], &aws); diags.HasError() {
			t.Fatal(diags)
		}
		attrs["aws"] = aws
		attrs["disable_ha"] = types.ObjectValueMust(map[string]attr.Type{}, map[string]attr.Value{})
		value, err := types.ObjectValueMust(rootType.AttrTypes, attrs).ToTerraformValue(ctx)
		if err != nil {
			t.Fatal(err)
		}
		return value
	}
	req := resource.ModifyPlanRequest{
		Plan:  tfsdk.Plan{Schema: schema.Schema, Raw: raw(after)},
		State: tfsdk.State{Schema: schema.Schema, Raw: raw(before)},
	}
	resp := resource.ModifyPlanResponse{}
	r.ModifyPlan(ctx, req, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatal(resp.Diagnostics)
	}
	if len(resp.RequiresReplace) != 0 {
		t.Fatalf("device-only edit would replace site: %v", resp.RequiresReplace)
	}
}
