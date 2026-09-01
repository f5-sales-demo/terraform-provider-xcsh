// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// validateSecuremeshSiteV2AWSContract enforces the MAC-bound configuration
// portion of the published AWS SMSv2 v2 contract. Unknown Terraform values are
// deferred until they are known; known invalid values fail immediately.
func validateSecuremeshSiteV2AWSContract(
	ctx context.Context,
	data SecuremeshSiteV2ResourceModel,
	resp *resource.ValidateConfigResponse,
) {
	if data.AWS == nil {
		return
	}
	if !data.Namespace.IsUnknown() && !data.Namespace.IsNull() && data.Namespace.ValueString() != "system" {
		resp.Diagnostics.AddAttributeError(
			path.Root("namespace"),
			"AWS SMSv2 Requires the System Namespace",
			"The verified AWS Secure Mesh Site v2 contract is limited to namespace \"system\".",
		)
		return
	}
	if data.AWS.NotManaged == nil || data.AWS.NotManaged.NodeList.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("aws").AtName("not_managed").AtName("node_list"),
			"AWS SMSv2 Nodes Are Required",
			"AWS CE configuration requires an ordered non-empty node_list.",
		)
		return
	}
	if data.AWS.NotManaged.NodeList.IsUnknown() {
		return
	}
	var nodes []SecuremeshSiteV2AWSNotManagedNodeListModel
	resp.Diagnostics.Append(data.AWS.NotManaged.NodeList.ElementsAs(ctx, &nodes, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if len(nodes) == 0 {
		resp.Diagnostics.AddAttributeError(
			path.Root("aws").AtName("not_managed").AtName("node_list"),
			"AWS SMSv2 Nodes Are Required",
			"AWS CE configuration requires an ordered non-empty node_list.",
		)
		return
	}
	for nodeIndex, node := range nodes {
		if node.InterfaceList.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("aws").AtName("not_managed").AtName("node_list").AtListIndex(nodeIndex).AtName("interface_list"),
				"AWS SMSv2 Interfaces Are Required",
				"Every AWS CE node requires an ordered MAC-bound interface_list with exactly one SLO and one SLI.",
			)
			continue
		}
		if node.InterfaceList.IsUnknown() {
			continue
		}
		var interfaces []SecuremeshSiteV2AWSNotManagedNodeListInterfaceListModel
		interfaceDiags := node.InterfaceList.ElementsAs(ctx, &interfaces, false)
		resp.Diagnostics.Append(interfaceDiags...)
		if interfaceDiags.HasError() {
			continue
		}
		macs := map[string]bool{}
		roles := map[string]bool{}
		unknownInterface := false
		for interfaceIndex, iface := range interfaces {
			interfacePath := path.Root("aws").AtName("not_managed").AtName("node_list").AtListIndex(nodeIndex).AtName("interface_list").AtListIndex(interfaceIndex)
			if iface.EthernetInterface == nil || iface.EthernetInterface.Mac.IsNull() {
				resp.Diagnostics.AddAttributeError(interfacePath, "AWS SMSv2 Interface MAC Is Required", "Each AWS CE interface must be bound to its ENI MAC address.")
				continue
			}
			if iface.EthernetInterface.Mac.IsUnknown() {
				unknownInterface = true
				continue
			}
			if strings.TrimSpace(iface.EthernetInterface.Mac.ValueString()) == "" {
				resp.Diagnostics.AddAttributeError(interfacePath.AtName("ethernet_interface").AtName("mac"), "AWS SMSv2 Interface MAC Is Required", "Each AWS CE interface must be bound to its ENI MAC address.")
				continue
			}
			if !iface.EthernetInterface.Device.IsNull() && !iface.EthernetInterface.Device.IsUnknown() {
				resp.Diagnostics.AddAttributeError(interfacePath.AtName("ethernet_interface").AtName("device"), "Guest Interface Inference Is Unsupported", "Do not configure a guest interface name; the verified contract binds interfaces by MAC only.")
			}
			parsedMAC, err := net.ParseMAC(strings.TrimSpace(iface.EthernetInterface.Mac.ValueString()))
			if err != nil || len(parsedMAC) != 6 {
				resp.Diagnostics.AddAttributeError(interfacePath.AtName("ethernet_interface").AtName("mac"), "AWS SMSv2 Interface MAC Is Invalid", "Each interface MAC must be a six-octet IEEE 802 MAC address.")
				continue
			}
			mac := strings.ToLower(parsedMAC.String())
			if macs[mac] {
				resp.Diagnostics.AddAttributeError(interfacePath.AtName("ethernet_interface").AtName("mac"), "AWS SMSv2 Interface MAC Is Duplicate", "Each interface MAC must be unique within its CE node.")
			}
			macs[mac] = true
			role := ""
			if iface.NetworkOption != nil && iface.NetworkOption.SiteLocalNetwork != nil {
				role = "slo"
			}
			if iface.NetworkOption != nil && iface.NetworkOption.SiteLocalInsideNetwork != nil {
				if role != "" {
					resp.Diagnostics.AddAttributeError(interfacePath.AtName("network_option"), "AWS SMSv2 Interface Role Is Ambiguous", "An interface may have exactly one role.")
					continue
				}
				role = "sli"
			}
			if role == "" {
				resp.Diagnostics.AddAttributeError(interfacePath.AtName("network_option"), "AWS SMSv2 Interface Role Is Required", "Each interface must declare an explicit SLO or SLI role; guest-interface inference is unsupported.")
				continue
			}
			if roles[role] {
				resp.Diagnostics.AddAttributeError(interfacePath.AtName("network_option"), "AWS SMSv2 Interface Role Is Duplicate", fmt.Sprintf("Role %q may appear only once per CE node.", role))
			}
			roles[role] = true
		}
		if !roles["slo"] && !unknownInterface {
			resp.Diagnostics.AddAttributeError(
				path.Root("aws").AtName("not_managed").AtName("node_list").AtListIndex(nodeIndex).AtName("interface_list"),
				"AWS SMSv2 SLO Is Required",
				"Every AWS CE node must declare exactly one MAC-bound SLO interface.",
			)
		}
		if !roles["sli"] && !unknownInterface {
			resp.Diagnostics.AddAttributeError(
				path.Root("aws").AtName("not_managed").AtName("node_list").AtListIndex(nodeIndex).AtName("interface_list"),
				"AWS SMSv2 SLI Is Required",
				"Every AWS CE node must declare exactly one MAC-bound SLI interface.",
			)
		}
	}
}
