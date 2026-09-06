// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestSMSv2AddressingFieldsReachEverySupportedPlatform(t *testing.T) {
	var response resource.SchemaResponse
	(&SecuremeshSiteV2Resource{}).Schema(context.Background(), resource.SchemaRequest{}, &response)
	for _, platform := range []string{"aws", "azure", "baremetal", "equinix", "gcp", "kvm", "nutanix", "oci", "openshift_virtualization", "openstack", "vmware"} {
		t.Run(platform, func(t *testing.T) {
			branch := response.Schema.Blocks[platform].(schema.SingleNestedBlock)
			nodes := branch.Blocks["not_managed"].(schema.SingleNestedBlock).Blocks["node_list"].(schema.ListNestedBlock)
			interfaces := nodes.NestedObject.Blocks["interface_list"].(schema.ListNestedBlock)
			static := interfaces.NestedObject.Blocks["static_ip"].(schema.SingleNestedBlock)
			if _, ok := static.Attributes["dns_server"]; !ok {
				t.Error("static DNS retained by the current API is absent from Terraform")
			}
			dhcp, ok := interfaces.NestedObject.Blocks["dhcp_server"].(schema.SingleNestedBlock)
			if !ok {
				t.Fatal("DHCP server retained by the current API is absent from Terraform")
			}
			if _, ok := dhcp.Attributes["dhcp_option82_tag"]; !ok {
				t.Error("option 82 is absent")
			}
			networks := dhcp.Blocks["dhcp_networks"].(schema.ListNestedBlock)
			pools := networks.NestedObject.Blocks["pools"].(schema.ListNestedBlock)
			if _, ok := pools.NestedObject.Attributes["exclude"]; !ok {
				t.Error("excluded DHCP pool is absent")
			}
		})
	}
	if _, exists := response.Schema.Blocks["rseries"]; exists {
		t.Error("unsupported rSeries must not be reintroduced")
	}
}
