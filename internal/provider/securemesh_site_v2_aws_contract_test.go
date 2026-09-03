// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestValidateSecuremeshSiteV2AWSContract(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name string
		data SecuremeshSiteV2ResourceModel
	}{
		{
			name: "requires system namespace",
			data: SecuremeshSiteV2ResourceModel{
				Namespace: types.StringValue("tenant-a"),
				AWS:       &SecuremeshSiteV2AWSModel{},
			},
		},
		{
			name: "requires nodes",
			data: SecuremeshSiteV2ResourceModel{
				Namespace: types.StringValue("system"),
				AWS: &SecuremeshSiteV2AWSModel{
					NotManaged: &SecuremeshSiteV2AWSNotManagedModel{NodeList: types.ListNull(types.StringType)},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var resp resource.ValidateConfigResponse
			validateSecuremeshSiteV2AWSContract(ctx, test.data, &resp)
			if !resp.Diagnostics.HasError() {
				t.Fatal("expected AWS SMSv2 contract validation error")
			}
		})
	}
}

func TestValidateSecuremeshSiteV2AWSCreateCapability(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		data         SecuremeshSiteV2ResourceModel
		capabilities map[string]string
		wantError    string
	}{
		{
			name:         "non AWS site does not require AWS capability",
			data:         SecuremeshSiteV2ResourceModel{},
			capabilities: map[string]string{"aws_ce_create": "unavailable"},
		},
		{
			name:         "available contract permits AWS create",
			data:         SecuremeshSiteV2ResourceModel{AWS: &SecuremeshSiteV2AWSModel{}},
			capabilities: map[string]string{"aws_ce_create": "available"},
		},
		{
			name:         "unavailable contract fails closed at create",
			data:         SecuremeshSiteV2ResourceModel{AWS: &SecuremeshSiteV2AWSModel{}},
			capabilities: map[string]string{"aws_ce_create": "unavailable"},
			wantError:    "aws_ce_create",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateSecuremeshSiteV2AWSCreateCapability(test.data, test.capabilities, "v5.0.1")
			if test.wantError == "" && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if test.wantError != "" && (err == nil || !strings.Contains(err.Error(), test.wantError)) {
				t.Fatalf("error = %v, want containing %q", err, test.wantError)
			}
		})
	}
}

func TestValidateSecuremeshSiteV2AWSContractDefersUnknownValues(t *testing.T) {
	t.Parallel()
	tests := []SecuremeshSiteV2ResourceModel{
		{Namespace: types.StringUnknown(), AWS: &SecuremeshSiteV2AWSModel{NotManaged: &SecuremeshSiteV2AWSNotManagedModel{NodeList: types.ListUnknown(types.ObjectType{AttrTypes: SecuremeshSiteV2AWSNotManagedNodeListModelAttrTypes})}}},
		awsSMSv2ContractFixture(t, []contractInterface{{mac: "unknown", role: "slo"}, {mac: "02:00:00:00:00:02", role: "sli"}}),
		awsSMSv2ContractFixture(t, []contractInterface{{mac: "02:00:00:00:00:01", role: "slo", device: "unknown"}, {mac: "02:00:00:00:00:02", role: "sli"}}),
	}
	for _, data := range tests {
		var resp resource.ValidateConfigResponse
		validateSecuremeshSiteV2AWSContract(context.Background(), data, &resp)
		if resp.Diagnostics.HasError() {
			t.Fatalf("unknown values must defer, got %v", resp.Diagnostics)
		}
	}
}

func TestValidateSecuremeshSiteV2AWSContractKnownValues(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		interfaces []contractInterface
		want       string
	}{
		{name: "valid", interfaces: []contractInterface{{mac: "02:00:00:00:00:01", role: "slo"}, {mac: "02:00:00:00:00:02", role: "sli"}}},
		{name: "null mac", interfaces: []contractInterface{{mac: "null", role: "slo"}, {mac: "02:00:00:00:00:02", role: "sli"}}, want: "MAC Is Required"},
		{name: "empty mac", interfaces: []contractInterface{{mac: "  ", role: "slo"}, {mac: "02:00:00:00:00:02", role: "sli"}}, want: "MAC Is Required"},
		{name: "malformed mac", interfaces: []contractInterface{{mac: "not-a-mac", role: "slo"}, {mac: "02:00:00:00:00:02", role: "sli"}}, want: "MAC Is Invalid"},
		{name: "normalized duplicate mac", interfaces: []contractInterface{{mac: "02-00-00-00-00-01", role: "slo"}, {mac: "02:00:00:00:00:01", role: "sli"}}, want: "MAC Is Duplicate"},
		{name: "missing device", interfaces: []contractInterface{{mac: "02:00:00:00:00:01", role: "slo", device: "null"}, {mac: "02:00:00:00:00:02", role: "sli"}}, want: "Device Is Required"},
		{name: "empty device", interfaces: []contractInterface{{mac: "02:00:00:00:00:01", role: "slo", device: "empty"}, {mac: "02:00:00:00:00:02", role: "sli"}}, want: "Device Is Required"},
		{name: "device role mismatch", interfaces: []contractInterface{{mac: "02:00:00:00:00:01", role: "slo", device: "eth1"}, {mac: "02:00:00:00:00:02", role: "sli"}}, want: "Device Does Not Match Role"},
		{name: "missing sli", interfaces: []contractInterface{{mac: "02:00:00:00:00:01", role: "slo"}}, want: "SLI Is Required"},
		{name: "duplicate role", interfaces: []contractInterface{{mac: "02:00:00:00:00:01", role: "slo"}, {mac: "02:00:00:00:00:02", role: "slo"}}, want: "Role Is Duplicate"},
		{name: "ambiguous role", interfaces: []contractInterface{{mac: "02:00:00:00:00:01", role: "ambiguous"}, {mac: "02:00:00:00:00:02", role: "sli"}}, want: "Role Is Ambiguous"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var resp resource.ValidateConfigResponse
			validateSecuremeshSiteV2AWSContract(context.Background(), awsSMSv2ContractFixture(t, test.interfaces), &resp)
			got := resp.Diagnostics.Errors()
			if test.want == "" && len(got) != 0 {
				t.Fatalf("unexpected diagnostics: %v", got)
			}
			if test.want != "" && !strings.Contains(gotSummary(got), test.want) {
				t.Fatalf("diagnostics %q do not contain %q", gotSummary(got), test.want)
			}
		})
	}
}

type contractInterface struct{ mac, role, device string }

func awsSMSv2ContractFixture(t *testing.T, fixtures []contractInterface) SecuremeshSiteV2ResourceModel {
	t.Helper()
	interfaces := make([]SecuremeshSiteV2AWSNotManagedNodeListInterfaceListModel, 0, len(fixtures))
	for _, fixture := range fixtures {
		mac := types.StringValue(fixture.mac)
		if fixture.mac == "null" {
			mac = types.StringNull()
		}
		if fixture.mac == "unknown" {
			mac = types.StringUnknown()
		}
		deviceName := map[string]string{"slo": "eth0", "sli": "eth1", "ambiguous": "eth0"}[fixture.role]
		if deviceName == "" {
			deviceName = "eth0"
		}
		device := types.StringValue(deviceName)
		if fixture.device != "" && fixture.device != "null" && fixture.device != "empty" && fixture.device != "unknown" {
			device = types.StringValue(fixture.device)
		}
		if fixture.device == "null" {
			device = types.StringNull()
		}
		if fixture.device == "empty" {
			device = types.StringValue("  ")
		}
		if fixture.device == "unknown" {
			device = types.StringUnknown()
		}
		role := &SecuremeshSiteV2AWSNotManagedNodeListInterfaceListNetworkOptionModel{}
		switch fixture.role {
		case "slo":
			role.SiteLocalNetwork = &SecuremeshSiteV2EmptyModel{}
		case "sli":
			role.SiteLocalInsideNetwork = &SecuremeshSiteV2EmptyModel{}
		case "ambiguous":
			role.SiteLocalNetwork, role.SiteLocalInsideNetwork = &SecuremeshSiteV2EmptyModel{}, &SecuremeshSiteV2EmptyModel{}
		}
		interfaces = append(interfaces, SecuremeshSiteV2AWSNotManagedNodeListInterfaceListModel{
			Labels:            types.MapNull(types.StringType),
			EthernetInterface: &SecuremeshSiteV2AWSNotManagedNodeListInterfaceListEthernetInterfaceModel{Mac: mac, Device: device},
			NetworkOption:     role,
		})
	}
	interfaceValues, diagnostics := types.ListValueFrom(context.Background(), types.ObjectType{AttrTypes: SecuremeshSiteV2AWSNotManagedNodeListInterfaceListModelAttrTypes}, interfaces)
	if diagnostics.HasError() {
		t.Fatal(diagnostics)
	}
	nodes := []SecuremeshSiteV2AWSNotManagedNodeListModel{{Hostname: types.StringValue("ce-1"), InterfaceList: interfaceValues}}
	nodeValues, diagnostics := types.ListValueFrom(context.Background(), types.ObjectType{AttrTypes: SecuremeshSiteV2AWSNotManagedNodeListModelAttrTypes}, nodes)
	if diagnostics.HasError() {
		t.Fatal(diagnostics)
	}
	return SecuremeshSiteV2ResourceModel{Namespace: types.StringValue("system"), AWS: &SecuremeshSiteV2AWSModel{NotManaged: &SecuremeshSiteV2AWSNotManagedModel{NodeList: nodeValues}}}
}

func gotSummary(diagnostics diag.Diagnostics) string {
	var summaries []string
	for _, diagnostic := range diagnostics {
		summaries = append(summaries, diagnostic.Summary())
	}
	return strings.Join(summaries, "; ")
}
