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
	if !emptyObjectMarkerConfigured(plan.DisableHA) || !emptyObjectMarkerConfigured(state.DisableHA) ||
		emptyObjectMarkerConfigured(plan.EnableHA) || emptyObjectMarkerConfigured(state.EnableHA) ||
		plan.AWS == nil || state.AWS == nil || plan.AWS.NotManaged == nil || state.AWS.NotManaged == nil {
		return false
	}
	planned, previous := plan.AWS.NotManaged.NodeList, state.AWS.NotManaged.NodeList
	if planned.IsNull() || planned.IsUnknown() || previous.IsUnknown() {
		return false
	}
	var nodes, oldNodes []SecuremeshSiteV2AWSNotManagedNodeListModel
	if planned.ElementsAs(ctx, &nodes, false).HasError() || len(nodes) != 1 {
		return false
	}
	if !previous.IsNull() && previous.ElementsAs(ctx, &oldNodes, false).HasError() {
		return false
	}
	if previous.IsNull() || len(oldNodes) == 0 {
		// Discovery creates the immutable site identity without a node. The first
		// validated non-HA node binding is an in-place configuration update.
		if !configuredSMSv2AWSNodeIsKnown(ctx, nodes[0]) {
			return false
		}
		normalizedAWS := *plan.AWS
		normalizedNotManaged := *plan.AWS.NotManaged
		if previous.IsNull() {
			normalizedNotManaged.NodeList = types.ListNull(planned.ElementType(ctx))
		} else {
			normalizedNotManaged.NodeList = previous
		}
		normalizedAWS.NotManaged = &normalizedNotManaged
		return sameSMSv2AWSInputs(ctx, &normalizedAWS, state.AWS)
	}
	if len(oldNodes) != 1 {
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
	return sameSMSv2AWSInputs(ctx, &normalizedAWS, state.AWS)
}

// pendingSMSv2AWSDiscoveryBinding reports the one transition whose concrete
// interface identity is often produced by a data source during planning. An
// initial ModifyPlan pass can therefore see the single node but not its
// device/MAC values. Terraform invokes ModifyPlan again after those values are
// known; recording replacement in the preliminary pass is irreversible even
// when the later pass proves this is the supported in-place transition.
func pendingSMSv2AWSDiscoveryBinding(ctx context.Context, plan, state SecuremeshSiteV2ResourceModel) bool {
	if !emptyObjectMarkerConfigured(plan.DisableHA) || !emptyObjectMarkerConfigured(state.DisableHA) ||
		emptyObjectMarkerConfigured(plan.EnableHA) || emptyObjectMarkerConfigured(state.EnableHA) ||
		plan.AWS == nil || state.AWS == nil || plan.AWS.NotManaged == nil || state.AWS.NotManaged == nil {
		return false
	}
	planned, previous := plan.AWS.NotManaged.NodeList, state.AWS.NotManaged.NodeList
	if previous.IsUnknown() {
		return false
	}
	var oldNodes []SecuremeshSiteV2AWSNotManagedNodeListModel
	if !previous.IsNull() && (previous.ElementsAs(ctx, &oldNodes, false).HasError() || len(oldNodes) != 0) {
		return false
	}
	if planned.IsUnknown() {
		return true
	}
	if planned.IsNull() {
		return false
	}
	var nodes []SecuremeshSiteV2AWSNotManagedNodeListModel
	if planned.ElementsAs(ctx, &nodes, false).HasError() || len(nodes) != 1 {
		return false
	}
	return !configuredSMSv2AWSNodeIsKnown(ctx, nodes[0])
}

// configuredSMSv2AWSNodeIsKnown prevents an unknown binding from silently
// changing the lifecycle decision after planning. Full schema validation stays
// authoritative for configuration semantics; this guard only establishes that
// Terraform has a concrete identity-bearing binding to update in place.
func configuredSMSv2AWSNodeIsKnown(ctx context.Context, node SecuremeshSiteV2AWSNotManagedNodeListModel) bool {
	if node.InterfaceList.IsNull() || node.InterfaceList.IsUnknown() {
		return false
	}
	var interfaces []SecuremeshSiteV2AWSNotManagedNodeListInterfaceListModel
	if node.InterfaceList.ElementsAs(ctx, &interfaces, false).HasError() || len(interfaces) == 0 {
		return false
	}
	for _, entry := range interfaces {
		if entry.EthernetInterface == nil ||
			entry.EthernetInterface.Mac.IsNull() || entry.EthernetInterface.Mac.IsUnknown() ||
			entry.EthernetInterface.Device.IsNull() || entry.EthernetInterface.Device.IsUnknown() {
			return false
		}
	}
	return true
}

func emptyObjectMarkerConfigured(value types.Object) bool {
	return !value.IsNull() && !value.IsUnknown()
}

// sameSMSv2AWSInputs excludes read-only interface observations from topology
// comparisons. Framework marks computed attributes unknown during an update;
// those output values do not represent a requested AWS configuration change.
func sameSMSv2AWSInputs(ctx context.Context, a, b *SecuremeshSiteV2AWSModel) bool {
	left, ok := smsv2AWSInputs(ctx, a)
	if !ok {
		return false
	}
	right, ok := smsv2AWSInputs(ctx, b)
	return ok && reflect.DeepEqual(left, right)
}

func smsv2AWSInputs(ctx context.Context, value *SecuremeshSiteV2AWSModel) (*SecuremeshSiteV2AWSModel, bool) {
	if value == nil {
		return nil, true
	}
	result := *value
	if value.NotManaged == nil {
		return &result, true
	}
	notManaged := *value.NotManaged
	result.NotManaged = &notManaged
	list := notManaged.NodeList
	if list.IsNull() {
		return &result, true
	}
	if list.IsUnknown() {
		return nil, false
	}
	var nodes []SecuremeshSiteV2AWSNotManagedNodeListModel
	if list.ElementsAs(ctx, &nodes, false).HasError() {
		return nil, false
	}
	for n := range nodes {
		interfaces := nodes[n].InterfaceList
		if interfaces.IsNull() {
			continue
		}
		if interfaces.IsUnknown() {
			return nil, false
		}
		var entries []SecuremeshSiteV2AWSNotManagedNodeListInterfaceListModel
		if interfaces.ElementsAs(ctx, &entries, false).HasError() {
			return nil, false
		}
		for i := range entries {
			entries[i].IsPrimary = types.BoolNull()
			entries[i].IsManagement = types.BoolNull()
		}
		normalized, d := types.ListValueFrom(ctx, interfaces.ElementType(ctx), entries)
		if d.HasError() {
			return nil, false
		}
		nodes[n].InterfaceList = normalized
	}
	normalized, d := types.ListValueFrom(ctx, list.ElementType(ctx), nodes)
	if d.HasError() {
		return nil, false
	}
	notManaged.NodeList = normalized
	return &result, true
}
