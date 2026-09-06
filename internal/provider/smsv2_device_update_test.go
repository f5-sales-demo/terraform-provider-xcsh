// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"testing"
)

func TestSMSv2DeviceUpdateRetainsTopologyReplacementGuards(t *testing.T) {
	for _, scenario := range []string{"non-HA device edit", "HA unspecified", "HA enabled", "MAC edit", "unknown device", "empty node list", "two nodes"} {
		t.Run(scenario, func(t *testing.T) {
			fixtures := []contractInterface{{mac: "02:00:00:00:00:01", role: "slo"}, {mac: "02:00:00:00:00:02", role: "sli"}}
			before := awsSMSv2ContractFixture(t, fixtures)
			fixtures[1].device = "ens6"
			if scenario == "MAC edit" {
				fixtures[1].mac = "02:00:00:00:00:03"
			}
			if scenario == "unknown device" {
				fixtures[1].device = "unknown"
			}
			after := awsSMSv2ContractFixture(t, fixtures)
			before.DisableHA = &SecuremeshSiteV2EmptyModel{}
			after.DisableHA = &SecuremeshSiteV2EmptyModel{}
			switch scenario {
			case "HA unspecified":
				after.DisableHA = nil
			case "HA enabled":
				after.EnableHA = &SecuremeshSiteV2EmptyModel{}
			case "empty node list":
				after.AWS.NotManaged.NodeList = types.ListValueMust(after.AWS.NotManaged.NodeList.ElementType(context.Background()), []attr.Value{})
			case "two nodes":
				nodes := after.AWS.NotManaged.NodeList.Elements()
				after.AWS.NotManaged.NodeList = types.ListValueMust(after.AWS.NotManaged.NodeList.ElementType(context.Background()), append(nodes, nodes[0]))
			}
			if got := canUpdateSMSv2AWSDevices(context.Background(), after, before); got != (scenario == "non-HA device edit") {
				t.Fatalf("update eligibility=%v", got)
			}
			if scenario == "non-HA device edit" {
				if !after.AWS.NotManaged.NodeList.Elements()[0].(types.Object).Attributes()["interface_list"].(types.List).Elements()[1].(types.Object).Attributes()["ethernet_interface"].(types.Object).Attributes()["device"].Equal(types.StringValue("ens6")) {
					t.Fatal("guard mutated the plan")
				}
			}
		})
	}
}
