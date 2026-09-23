// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

func kvmRuntimeFixture() (client.SMSv2Observation, *client.RegistrationListResponse, client.SMSv2Observation) {
	site := client.SMSv2Observation{
		"metadata":        map[string]interface{}{"name": "onprem-nuc-kvm", "namespace": "system"},
		"system_metadata": map[string]interface{}{"uid": "site-uid-current"},
	}
	registrations := &client.RegistrationListResponse{Items: []client.RegistrationListItem{{
		Name:   "r-current",
		Object: client.RegistrationObject{Status: client.RegistrationStatus{CurrentState: "PENDING"}},
		GetSpec: client.RegistrationGetSpec{
			Passport: client.RegistrationPassport{ClusterName: "onprem-nuc-kvm", ClusterSize: 1},
			Infra: client.RegistrationInfra{
				Provider: "KVM", Hostname: "onprem-ce-01-674f7",
				HWInfo: client.RegistrationHWInfo{Network: []client.RegistrationNetwork{{Name: "ens3", MACAddress: "52:54:00:10:00:11"}}},
			},
		},
	}}}
	interfaces := client.SMSv2Observation{"items": []interface{}{
		map[string]interface{}{
			"name": "stale-interface", "namespace": "system",
			"owner_view": map[string]interface{}{"kind": "securemesh_site_v2", "name": "onprem-nuc-kvm", "namespace": "system", "uid": "site-uid-stale"},
			"get_spec":   map[string]interface{}{"ethernet_interface": map[string]interface{}{"node": "onprem-ce-01-674f7", "device": "ens3", "site_local_network": map[string]interface{}{}}},
		},
		map[string]interface{}{
			"name": "ves-io-owned-runtime-interface", "namespace": "system",
			"owner_view": map[string]interface{}{"kind": "securemesh_site_v2", "name": "onprem-nuc-kvm", "namespace": "system", "uid": "site-uid-current"},
			"get_spec":   map[string]interface{}{"ethernet_interface": map[string]interface{}{"node": "onprem-ce-01-674f7", "device": "ens3", "site_local_network": map[string]interface{}{}}},
		},
	}}
	return site, registrations, interfaces
}

func TestResolveSMSv2KVMRuntimeUsesOwnerRegistrationDeviceAndMAC(t *testing.T) {
	site, registrations, interfaces := kvmRuntimeFixture()
	got, err := resolveSMSv2KVMRuntime(site, registrations, interfaces, "52-54-00-10-00-11")
	if err != nil {
		t.Fatal(err)
	}
	if got.InterfaceName != "ves-io-owned-runtime-interface" || got.Hostname != "onprem-ce-01-674f7" || got.Device != "ens3" || got.MAC != "52:54:00:10:00:11" || got.RegistrationState != "PENDING" {
		t.Fatalf("unexpected KVM runtime identity: %#v", got)
	}
}

func TestResolveSMSv2KVMRuntimeRejectsAmbiguityAndDrift(t *testing.T) {
	for name, mutate := range map[string]func(client.SMSv2Observation, *client.RegistrationListResponse, client.SMSv2Observation){
		"missing site uid": func(site client.SMSv2Observation, _ *client.RegistrationListResponse, _ client.SMSv2Observation) {
			delete(site, "system_metadata")
		},
		"duplicate registration mac": func(_ client.SMSv2Observation, registrations *client.RegistrationListResponse, _ client.SMSv2Observation) {
			registrations.Items = append(registrations.Items, registrations.Items[0])
		},
		"foreign registration site": func(_ client.SMSv2Observation, registrations *client.RegistrationListResponse, _ client.SMSv2Observation) {
			registrations.Items[0].GetSpec.Passport.ClusterName = "other-site"
		},
		"wrong provider": func(_ client.SMSv2Observation, registrations *client.RegistrationListResponse, _ client.SMSv2Observation) {
			registrations.Items[0].GetSpec.Infra.Provider = "AWS"
		},
		"stale owner uid": func(_ client.SMSv2Observation, _ *client.RegistrationListResponse, interfaces client.SMSv2Observation) {
			interfaces["items"] = interfaces["items"].([]interface{})[:1]
		},
		"inside role": func(_ client.SMSv2Observation, _ *client.RegistrationListResponse, interfaces client.SMSv2Observation) {
			items := interfaces["items"].([]interface{})
			ethernet := items[1].(map[string]interface{})["get_spec"].(map[string]interface{})["ethernet_interface"].(map[string]interface{})
			delete(ethernet, "site_local_network")
			ethernet["site_local_inside_network"] = map[string]interface{}{}
		},
		"duplicate interface": func(_ client.SMSv2Observation, _ *client.RegistrationListResponse, interfaces client.SMSv2Observation) {
			items := interfaces["items"].([]interface{})
			interfaces["items"] = append(items, items[1])
		},
	} {
		t.Run(name, func(t *testing.T) {
			site, registrations, interfaces := kvmRuntimeFixture()
			mutate(site, registrations, interfaces)
			if _, err := resolveSMSv2KVMRuntime(site, registrations, interfaces, "52:54:00:10:00:11"); err == nil {
				t.Fatal("invalid runtime identity was accepted")
			}
		})
	}
}

func TestResolveSMSv2KVMRuntimeRejectsPartialAPIErrorsAndTerminalRegistrations(t *testing.T) {
	site, registrations, interfaces := kvmRuntimeFixture()
	registrations.Errors = []client.APIErrorDetail{{Code: "PARTIAL", Message: "redacted"}}
	if _, err := resolveSMSv2KVMRuntime(site, registrations, interfaces, "52:54:00:10:00:11"); err == nil || !strings.Contains(err.Error(), "partial errors") {
		t.Fatalf("registration partial error = %v", err)
	}
	registrations.Errors = nil
	registrations.Items[0].Object.Status.CurrentState = "FAILED_INACTIVE"
	if _, err := resolveSMSv2KVMRuntime(site, registrations, interfaces, "52:54:00:10:00:11"); err == nil || !strings.Contains(err.Error(), "0 live") {
		t.Fatalf("terminal registration error = %v", err)
	}
	registrations.Items[0].Object.Status.CurrentState = "PENDING"
	interfaces["errors"] = []interface{}{map[string]interface{}{"code": "PARTIAL"}}
	if _, err := resolveSMSv2KVMRuntime(site, registrations, interfaces, "52:54:00:10:00:11"); err == nil || !strings.Contains(err.Error(), "partial errors") {
		t.Fatalf("interface partial error = %v", err)
	}
}
