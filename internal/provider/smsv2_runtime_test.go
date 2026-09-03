// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func runtimeConfiguration() client.SMSv2Observation {
	interfaces := []interface{}{
		map[string]interface{}{"name": "slo", "mtu": float64(1500), "ethernet_interface": map[string]interface{}{"device": "eth0", "mac": "02:AA:BB:CC:DD:01"}, "network_option": map[string]interface{}{"site_local_network": map[string]interface{}{}}},
		map[string]interface{}{"name": "sli", "mtu": float64(1500), "ethernet_interface": map[string]interface{}{"device": "eth1", "mac": "02:aa:bb:cc:dd:02"}, "network_option": map[string]interface{}{"site_local_inside_network": map[string]interface{}{}}},
	}
	nodes := []interface{}{map[string]interface{}{"hostname": "master-0", "interface_list": interfaces}}
	return client.SMSv2Observation{
		"metadata": map[string]interface{}{"name": "lab-site", "namespace": "system"},
		"spec": map[string]interface{}{
			"aws": map[string]interface{}{
				"not_managed": map[string]interface{}{"node_list": nodes},
			},
		},
	}
}

func multiNodeRuntimeConfiguration() client.SMSv2Observation {
	nodes := []interface{}{}
	for index, hostname := range []string{"master-0", "master-1", "master-2"} {
		nodes = append(nodes, map[string]interface{}{
			"hostname": hostname,
			"interface_list": []interface{}{
				map[string]interface{}{"name": "slo", "mtu": float64(1500), "ethernet_interface": map[string]interface{}{"device": "eth0", "mac": fmt.Sprintf("02:aa:bb:cc:%02x:01", index)}, "network_option": map[string]interface{}{"site_local_network": map[string]interface{}{}}},
				map[string]interface{}{"name": "sli", "mtu": float64(1500), "ethernet_interface": map[string]interface{}{"device": "eth1", "mac": fmt.Sprintf("02:aa:bb:cc:%02x:02", index)}, "network_option": map[string]interface{}{"site_local_inside_network": map[string]interface{}{}}},
			},
		})
	}
	return client.SMSv2Observation{"metadata": map[string]interface{}{"name": "lab-site", "namespace": "system"}, "spec": map[string]interface{}{"aws": map[string]interface{}{"not_managed": map[string]interface{}{"node_list": nodes}}}}
}

func TestSMSv2CapabilitiesFailClosedBeforeRuntimeReads(t *testing.T) {
	t.Parallel()
	capabilities := map[string]string{
		"aws_ce_create":  "unavailable",
		"runtime_status": "unavailable",
		"tgw_connect":    "unavailable",
	}
	err := validateSMSv2Capabilities(capabilities, "v5.0.1", "runtime_status", "tgw_connect")
	if err == nil || !strings.Contains(err.Error(), "runtime_status, tgw_connect") {
		t.Fatalf("fail-closed capability error = %v", err)
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
	if len(got) != 2 || got["outside"].InterfaceName.ValueString() != "ves-io-securemesh-site-v2-lab-site-network-master-0-eth0-0" || !got["inside"].Healthy.ValueBool() {
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

func TestSMSv2RuntimeCorrelatesMultiNodeSiteGlobalHealth(t *testing.T) {
	configured, err := extractSMSv2ConfiguredInterfaces(multiNodeRuntimeConfiguration())
	if err != nil {
		t.Fatal(err)
	}
	bindings := map[string]smsv2BindingModel{}
	for index, hostname := range []string{"master-0", "master-1", "master-2"} {
		bindings[fmt.Sprintf("node_%d_slo", index)] = smsv2BindingModel{Node: types.StringValue(hostname), Role: types.StringValue("slo"), MAC: types.StringValue(fmt.Sprintf("02:aa:bb:cc:%02x:01", index))}
		bindings[fmt.Sprintf("node_%d_sli", index)] = smsv2BindingModel{Node: types.StringValue(hostname), Role: types.StringValue("sli"), MAC: types.StringValue(fmt.Sprintf("02:aa:bb:cc:%02x:02", index))}
	}
	bindings, err = validateSMSv2Bindings(bindings)
	if err != nil {
		t.Fatal(err)
	}

	got, err := correlateSMSv2Runtime(bindings, configured, client.SMSv2Observation{"hostname": "master-1.us-east-2.compute.internal", "state": "PROVISIONED"})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 6 {
		t.Fatalf("got %d interfaces, want 6", len(got))
	}
	for key, iface := range got {
		if !iface.Healthy.ValueBool() {
			t.Fatalf("interface %q unexpectedly unhealthy for provisioned global site health", key)
		}
	}
}

func TestSMSv2RuntimeRejectsForeignOrIncompleteSiteHealth(t *testing.T) {
	configured, err := extractSMSv2ConfiguredInterfaces(multiNodeRuntimeConfiguration())
	if err != nil {
		t.Fatal(err)
	}
	bindings := map[string]smsv2BindingModel{
		"outside": {Node: types.StringValue("master-0"), Role: types.StringValue("slo"), MAC: types.StringValue("02:aa:bb:cc:00:01")},
	}
	bindings, err = validateSMSv2Bindings(bindings)
	if err != nil {
		t.Fatal(err)
	}

	for name, health := range map[string]client.SMSv2Observation{
		"foreign node":  {"hostname": "other-site-node", "state": "PROVISIONED"},
		"missing node":  {"state": "PROVISIONED"},
		"missing state": {"hostname": "master-0"},
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := correlateSMSv2Runtime(bindings, configured, health); err == nil {
				t.Fatal("expected invalid site health error")
			}
		})
	}
}

func TestSMSv2RuntimePropagatesUnhealthySiteState(t *testing.T) {
	configured, err := extractSMSv2ConfiguredInterfaces(multiNodeRuntimeConfiguration())
	if err != nil {
		t.Fatal(err)
	}
	bindings := map[string]smsv2BindingModel{
		"outside": {Node: types.StringValue("master-0"), Role: types.StringValue("slo"), MAC: types.StringValue("02:aa:bb:cc:00:01")},
		"inside":  {Node: types.StringValue("master-1"), Role: types.StringValue("sli"), MAC: types.StringValue("02:aa:bb:cc:01:02")},
	}
	bindings, err = validateSMSv2Bindings(bindings)
	if err != nil {
		t.Fatal(err)
	}

	got, err := correlateSMSv2Runtime(bindings, configured, client.SMSv2Observation{"hostname": "master-2", "state": "WAITING_FOR_CONFIG"})
	if err != nil {
		t.Fatal(err)
	}
	for key, iface := range got {
		if iface.Healthy.ValueBool() {
			t.Fatalf("interface %q unexpectedly healthy: %#v", key, iface)
		}
	}
}

func TestSMSv2BGPConvergenceUsesExactNodeScopedSchemas(t *testing.T) {
	configured, err := extractSMSv2ConfiguredInterfaces(runtimeConfiguration())
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]smsv2ExpectedPeerModel{"peer-a": {
		Node: types.StringValue("master-0"), Role: types.StringValue("slo"),
		MAC: types.StringValue("02:aa:bb:cc:dd:01"), PeerAddress: types.StringValue("169.254.10.1"),
		ExpectedRoutes: types.SetValueMust(types.StringType, []attr.Value{types.StringValue("10.10.0.0/16")}),
	}}
	peers := []observedBGPPeer{{
		Node: "master-0", InterfaceName: "ves-io-securemesh-site-v2-lab-site-network-master-0-eth0-0", PeerAddress: "169.254.10.1",
		State: "Established", StateChangedAt: "2026-09-03T00:00:00Z",
		ReceivedPrefixes: 1, AdvertisedPrefixes: 1,
	}}
	bgpRoutes := client.SMSv2Observation{"ver": []interface{}{map[string]interface{}{
		"name": "master-0", "ri_table": []interface{}{map[string]interface{}{
			"rt_table": []interface{}{map[string]interface{}{
				"imported": []interface{}{map[string]interface{}{"subnet": "10.10.0.0/16"}},
				"exported": []interface{}{map[string]interface{}{"subnet": "10.20.0.0/16"}},
			}},
		}},
	}}}
	simplified := client.SMSv2Observation{"ver_routes": []interface{}{map[string]interface{}{
		"node": "master-0", "route": []interface{}{map[string]interface{}{"prefix": "10.20.0.0/16"}},
	}}}
	status, converged, reason := convergeSMSv2BGP(expected, configured, peers, bgpRoutes, simplified, simplified)
	if !converged || reason != "" || !status["peer-a"].Established.ValueBool() {
		t.Fatalf("converged=%v reason=%q status=%#v", converged, reason, status)
	}
	if got := status["peer-a"].StateChangedAt.ValueString(); got != "2026-09-03T00:00:00Z" {
		t.Fatalf("state_changed_at = %q", got)
	}
	peers[0].StateChangedAt = "not-a-freshness-claim"
	if _, converged, reason = convergeSMSv2BGP(expected, configured, peers, bgpRoutes, simplified, simplified); !converged || reason != "" {
		t.Fatalf("state change timestamp was incorrectly treated as freshness: converged=%v reason=%q", converged, reason)
	}
}

func TestRoutePrefixesForNodeRejectsAmbiguousFQDNObservations(t *testing.T) {
	prefixes := map[string]map[string]struct{}{
		"master-0.example.internal": {"10.10.0.0/16": {}},
		"master-0.other.internal":   {"10.20.0.0/16": {}},
	}
	if _, err := routePrefixesForNode(prefixes, "master-0", "route"); err == nil || !strings.Contains(err.Error(), "2 observations") {
		t.Fatalf("ambiguous node observations were not rejected: %v", err)
	}
}

func TestSMSv2BindingIdentityValidation(t *testing.T) {
	sharedMAC := "02:aa:bb:cc:dd:01"
	accepted := map[string]smsv2BindingModel{
		"node0": {Node: types.StringValue("master-0"), Role: types.StringValue("slo"), MAC: types.StringValue(sharedMAC)},
		"node1": {Node: types.StringValue("master-1"), Role: types.StringValue("sli"), MAC: types.StringValue(sharedMAC)},
	}
	if _, err := validateSMSv2Bindings(accepted); err != nil {
		t.Fatalf("same MAC on distinct nodes was rejected: %v", err)
	}

	cases := map[string]smsv2BindingModel{
		"null node": {
			Node: types.StringNull(), Role: types.StringValue("slo"), MAC: types.StringValue(sharedMAC),
		},
		"unknown role": {
			Node: types.StringValue("master-0"), Role: types.StringUnknown(), MAC: types.StringValue(sharedMAC),
		},
		"empty node": {
			Node: types.StringValue(" "), Role: types.StringValue("slo"), MAC: types.StringValue(sharedMAC),
		},
		"malformed MAC": {
			Node: types.StringValue("master-0"), Role: types.StringValue("slo"), MAC: types.StringValue("not-a-mac"),
		},
		"guest device role": {
			Node: types.StringValue("master-0"), Role: types.StringValue("eth0"), MAC: types.StringValue(sharedMAC),
		},
	}
	for name, binding := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := validateSMSv2Bindings(map[string]smsv2BindingModel{"candidate": binding}); err == nil {
				t.Fatal("invalid binding was accepted")
			}
		})
	}

	duplicate := map[string]smsv2BindingModel{
		"one": {Node: types.StringValue("master-0"), Role: types.StringValue("slo"), MAC: types.StringValue(sharedMAC)},
		"two": {Node: types.StringValue("master-0"), Role: types.StringValue("sli"), MAC: types.StringValue("02-AA-BB-CC-DD-01")},
	}
	if _, err := validateSMSv2Bindings(duplicate); err == nil || !strings.Contains(err.Error(), "within node") {
		t.Fatalf("within-node duplicate error = %v", err)
	}
}
