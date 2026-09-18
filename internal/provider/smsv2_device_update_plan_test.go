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
	before = smsv2ComputedInterfaceFixture(t, before, false)
	after = smsv2ComputedInterfaceFixture(t, after, true)
	before.DisableHA = types.ObjectValueMust(map[string]attr.Type{}, map[string]attr.Value{})
	after.DisableHA = types.ObjectValueMust(map[string]attr.Type{}, map[string]attr.Value{})
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

	// Discovery persists the site identity with no registered node. Its first
	// concrete non-HA node binding must use the same in-place lifecycle path as
	// a device-only update; a replacement would delete the immutable site name.
	discovery := before
	discovery.AWS.NotManaged.NodeList = types.ListValueMust(
		discovery.AWS.NotManaged.NodeList.ElementType(ctx),
		[]attr.Value{},
	)
	resp = resource.ModifyPlanResponse{}
	req.Plan.Raw = raw(after)
	req.State.Raw = raw(discovery)
	var decodedPlan, decodedState SecuremeshSiteV2ResourceModel
	if diags := req.Plan.Get(ctx, &decodedPlan); diags.HasError() {
		t.Fatal(diags)
	}
	if diags := req.State.Get(ctx, &decodedState); diags.HasError() {
		t.Fatal(diags)
	}
	if canUpdateSMSv2AWSDevices(ctx, decodedPlan, decodedState) {
		t.Fatal("discovery_rebuild must not permit an in-place update")
	}
	r.ModifyPlan(ctx, req, &resp)
	if !resp.Diagnostics.HasError() {
		t.Fatal("discovery_rebuild transition must fail during planning")
	}
	return

	// The API persists the initial discovery result with an omitted node list,
	// which Terraform represents as null rather than an empty list. The first
	// concrete binding must retain that site identity in place as well.
	nullDiscovery := discovery
	nullDiscovery.AWS.NotManaged.NodeList = types.ListNull(
		nullDiscovery.AWS.NotManaged.NodeList.ElementType(ctx),
	)
	resp = resource.ModifyPlanResponse{}
	req.Plan.Raw = raw(after)
	req.State.Raw = raw(nullDiscovery)
	if diags := req.Plan.Get(ctx, &decodedPlan); diags.HasError() {
		t.Fatal(diags)
	}
	if diags := req.State.Get(ctx, &decodedState); diags.HasError() {
		t.Fatal(diags)
	}
	if !canUpdateSMSv2AWSDevices(ctx, decodedPlan, decodedState) {
		t.Fatal("decoded null-discovery-to-configured transition is not eligible for in-place update")
	}
	r.ModifyPlan(ctx, req, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatal(resp.Diagnostics)
	}
	if len(resp.RequiresReplace) != 0 {
		t.Fatalf("null discovery-to-configured binding would replace site: %v", resp.RequiresReplace)
	}

	// Terraform may call ModifyPlan before registration-derived interface
	// values become known. That preliminary pass must not irreversibly mark the
	// supported discovery transition for replacement.
	pending := after
	var nodes []SecuremeshSiteV2AWSNotManagedNodeListModel
	if diags := pending.AWS.NotManaged.NodeList.ElementsAs(ctx, &nodes, false); diags.HasError() {
		t.Fatal(diags)
	}
	var interfaces []SecuremeshSiteV2AWSNotManagedNodeListInterfaceListModel
	if diags := nodes[0].InterfaceList.ElementsAs(ctx, &interfaces, false); diags.HasError() {
		t.Fatal(diags)
	}
	interfaces[0].EthernetInterface.Device = types.StringUnknown()
	value, diags := types.ListValueFrom(ctx, nodes[0].InterfaceList.ElementType(ctx), interfaces)
	if diags.HasError() {
		t.Fatal(diags)
	}
	nodes[0].InterfaceList = value
	value, diags = types.ListValueFrom(ctx, pending.AWS.NotManaged.NodeList.ElementType(ctx), nodes)
	if diags.HasError() {
		t.Fatal(diags)
	}
	pending.AWS.NotManaged.NodeList = value
	resp = resource.ModifyPlanResponse{}
	req.Plan.Raw = raw(pending)
	req.State.Raw = raw(nullDiscovery)
	r.ModifyPlan(ctx, req, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatal(resp.Diagnostics)
	}
	if len(resp.RequiresReplace) != 0 {
		t.Fatalf("pending discovery binding would replace site: %v", resp.RequiresReplace)
	}
}

func smsv2ComputedInterfaceFixture(t *testing.T, model SecuremeshSiteV2ResourceModel, unknown bool) SecuremeshSiteV2ResourceModel {
	t.Helper()
	ctx := context.Background()
	var nodes []SecuremeshSiteV2AWSNotManagedNodeListModel
	if d := model.AWS.NotManaged.NodeList.ElementsAs(ctx, &nodes, false); d.HasError() {
		t.Fatal(d)
	}
	for n := range nodes {
		var interfaces []SecuremeshSiteV2AWSNotManagedNodeListInterfaceListModel
		if d := nodes[n].InterfaceList.ElementsAs(ctx, &interfaces, false); d.HasError() {
			t.Fatal(d)
		}
		for i := range interfaces {
			value := types.BoolValue(false)
			if unknown {
				value = types.BoolUnknown()
			}
			interfaces[i].IsPrimary = value
			interfaces[i].IsManagement = value
		}
		value, d := types.ListValueFrom(ctx, nodes[n].InterfaceList.ElementType(ctx), interfaces)
		if d.HasError() {
			t.Fatal(d)
		}
		nodes[n].InterfaceList = value
	}
	value, d := types.ListValueFrom(ctx, model.AWS.NotManaged.NodeList.ElementType(ctx), nodes)
	if d.HasError() {
		t.Fatal(d)
	}
	model.AWS.NotManaged.NodeList = value
	return model
}
