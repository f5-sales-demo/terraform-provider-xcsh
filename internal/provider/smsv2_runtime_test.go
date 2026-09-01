// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"strings"
	"testing"
	"time"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func runtimeConfiguration() client.SMSv2Observation {
	interfaces := []interface{}{
		map[string]interface{}{"name": "eth0", "mtu": float64(1500), "ethernet_interface": map[string]interface{}{"mac": "02:AA:BB:CC:DD:01"}, "network_option": map[string]interface{}{"site_local_network": map[string]interface{}{}}},
		map[string]interface{}{"name": "eth1", "mtu": float64(1500), "ethernet_interface": map[string]interface{}{"mac": "02:aa:bb:cc:dd:02"}, "network_option": map[string]interface{}{"site_local_inside_network": map[string]interface{}{}}},
	}
	nodes := []interface{}{map[string]interface{}{"hostname": "master-0", "interface_list": interfaces}}
	return client.SMSv2Observation{
		"spec": map[string]interface{}{
			"aws": map[string]interface{}{
				"not_managed": map[string]interface{}{"node_list": nodes},
			},
		},
	}
}

func TestExtractAndCorrelateSMSv2Runtime(t *testing.T) {
	configured, err := extractSMSv2ConfiguredInterfaces(runtimeConfiguration())
	if err != nil {
		t.Fatal(err)
	}
	bindings := map[string]smsv2BindingModel{
		"outside": {Node: types.StringValue("master-0"), Role: types.StringValue("slo"), MAC: types.StringValue("02-AA-BB-CC-DD-01")},
		"inside":  {Node: types.StringValue("master-0"), Role: types.StringValue("sli"), MAC: types.StringValue("02:aa:bb:cc:dd:02")},
	}
	bindings, err = validateSMSv2Bindings(bindings)
	if err != nil {
		t.Fatal(err)
	}
	got, err := correlateSMSv2Runtime(bindings, configured, client.SMSv2Observation{"hostname": "master-0", "state": "PROVISIONED"})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got["outside"].InterfaceName.ValueString() != "eth0" || !got["inside"].Healthy.ValueBool() {
		t.Fatalf("unexpected correlation: %#v", got)
	}
}

func TestSMSv2RuntimeRejectsDuplicateAndMismatch(t *testing.T) {
	bindings := map[string]smsv2BindingModel{
		"one": {Node: types.StringValue("master-0"), Role: types.StringValue("slo"), MAC: types.StringValue("02:aa:bb:cc:dd:01")},
		"two": {Node: types.StringValue("master-0"), Role: types.StringValue("sli"), MAC: types.StringValue("02-AA-BB-CC-DD-01")},
	}
	if _, err := validateSMSv2Bindings(bindings); err == nil || !strings.Contains(err.Error(), "duplicate MAC") {
		t.Fatalf("duplicate error = %v", err)
	}
	configured, err := extractSMSv2ConfiguredInterfaces(runtimeConfiguration())
	if err != nil {
		t.Fatal(err)
	}
	bindings = map[string]smsv2BindingModel{"outside": {Node: types.StringValue("wrong-node"), Role: types.StringValue("slo"), MAC: types.StringValue("02:aa:bb:cc:dd:01")}}
	bindings, _ = validateSMSv2Bindings(bindings)
	if _, err := correlateSMSv2Runtime(bindings, configured, client.SMSv2Observation{"hostname": "master-0", "state": "PROVISIONED"}); err == nil {
		t.Fatal("expected node mismatch")
	}
}

func TestSMSv2BGPConvergenceAndStaleness(t *testing.T) {
	configured, err := extractSMSv2ConfiguredInterfaces(runtimeConfiguration())
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]smsv2ExpectedPeerModel{"peer-a": {Node: types.StringValue("master-0"), Role: types.StringValue("slo"), MAC: types.StringValue("02:aa:bb:cc:dd:01"), PeerAddress: types.StringValue("169.254.10.1"), ExpectedRoutes: types.SetValueMust(types.StringType, []attr.Value{types.StringValue("10.10.0.0/16")})}}
	now := time.Now().UTC().Truncate(time.Second)
	peers := []observedBGPPeer{{Node: "master-0", InterfaceName: "eth0", PeerAddress: "169.254.10.1", State: "Established", ObservedAt: now.Format(time.RFC3339), ReceivedPrefixes: 1, AdvertisedPrefixes: 1}}
	routes := client.SMSv2Observation{"routes": []interface{}{map[string]interface{}{"subnet": "10.10.0.0/16"}}}
	status, converged, reason := convergeSMSv2BGP(expected, configured, peers, routes, client.SMSv2Observation{"ver_routes": []interface{}{map[string]interface{}{"route": []interface{}{map[string]interface{}{"prefix": "10.10.0.0/16"}}}}}, client.SMSv2Observation{}, time.Minute, now)
	if !converged || reason != "" || !status["peer-a"].Established.ValueBool() {
		t.Fatalf("converged=%v reason=%q status=%#v", converged, reason, status)
	}
	peers[0].ObservedAt = now.Add(-2 * time.Minute).Format(time.RFC3339)
	if _, converged, reason = convergeSMSv2BGP(expected, configured, peers, routes, routes, routes, time.Minute, now); converged || !strings.Contains(reason, "stale") {
		t.Fatalf("converged=%v reason=%q", converged, reason)
	}
}
