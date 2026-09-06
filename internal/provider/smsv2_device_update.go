// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// canUpdateSMSv2AWSDevices recognizes the narrow in-place update verified against
// the API: ethernet device edits on an explicitly non-HA, single-node site. It
// preserves replacement behavior for node identities/counts and other AWS edits.
func canUpdateSMSv2AWSDevices(ctx context.Context, plan, state SecuremeshSiteV2ResourceModel) bool {
	if plan.DisableHA == nil || state.DisableHA == nil || plan.EnableHA != nil || state.EnableHA != nil || plan.AWS == nil || state.AWS == nil || plan.AWS.NotManaged == nil || state.AWS.NotManaged == nil {
		return false
	}
	planned, previous := plan.AWS.NotManaged.NodeList, state.AWS.NotManaged.NodeList
	if planned.IsNull() || planned.IsUnknown() || previous.IsNull() || previous.IsUnknown() {
		return false
	}
	var nodes, oldNodes []SecuremeshSiteV2AWSNotManagedNodeListModel
	if planned.ElementsAs(ctx, &nodes, false).HasError() || previous.ElementsAs(ctx, &oldNodes, false).HasError() || len(nodes) != 1 || len(oldNodes) != 1 {
		return false
	}
	interfaces, oldInterfaces := nodes[0].InterfaceList, oldNodes[0].InterfaceList
	if interfaces.IsNull() || interfaces.IsUnknown() || oldInterfaces.IsNull() || oldInterfaces.IsUnknown() {
		return false
	}
	var current, old []SecuremeshSiteV2AWSNotManagedNodeListInterfaceListModel
	if interfaces.ElementsAs(ctx, &current, false).HasError() || oldInterfaces.ElementsAs(ctx, &old, false).HasError() || len(current) == 0 || len(current) != len(old) {
		return false
	}
	changed := false
	for i := range current {
		if current[i].EthernetInterface == nil || old[i].EthernetInterface == nil {
			return false
		}
		a, b := current[i].EthernetInterface.Device, old[i].EthernetInterface.Device
		if a.IsNull() || a.IsUnknown() || b.IsNull() || b.IsUnknown() {
			return false
		}
		changed = changed || !a.Equal(b)
		current[i].EthernetInterface.Device = b
	}
	if !changed {
		return false
	}
	normalizedInterfaces, diags := types.ListValueFrom(ctx, interfaces.ElementType(ctx), current)
	if diags.HasError() {
		return false
	}
	nodes[0].InterfaceList = normalizedInterfaces
	normalizedNodes, diags := types.ListValueFrom(ctx, planned.ElementType(ctx), nodes)
	if diags.HasError() {
		return false
	}
	normalizedAWS := *plan.AWS
	normalizedNotManaged := *plan.AWS.NotManaged
	normalizedNotManaged.NodeList = normalizedNodes
	normalizedAWS.NotManaged = &normalizedNotManaged
	return reflect.DeepEqual(&normalizedAWS, state.AWS)
}
