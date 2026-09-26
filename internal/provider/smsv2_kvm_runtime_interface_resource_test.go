// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	frameworkresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const longKVMRuntimeInterfaceName = "ves-io-securemesh-site-v2-mcn-ce-ha-smsv2-current-kvm-network-onprem-ce-01-90607-ens4-0"

func kvmRuntimeInterfaceFixture() (client.SMSv2Observation, *client.RegistrationListResponse, client.SMSv2Observation) {
	site := client.SMSv2Observation{
		"metadata":        map[string]interface{}{"name": "mcn-ce-ha-smsv2-current-kvm", "namespace": "system"},
		"system_metadata": map[string]interface{}{"uid": "site-uid-current"},
	}
	registrations := &client.RegistrationListResponse{Items: []client.RegistrationListItem{{
		Name:   "registration-current",
		Object: client.RegistrationObject{Status: client.RegistrationStatus{CurrentState: "ONLINE"}},
		GetSpec: client.RegistrationGetSpec{
			Passport: client.RegistrationPassport{ClusterName: "mcn-ce-ha-smsv2-current-kvm", ClusterSize: 1},
			Infra: client.RegistrationInfra{
				Provider: "KVM", Hostname: "onprem-ce-01-90607",
				HWInfo: client.RegistrationHWInfo{Network: []client.RegistrationNetwork{
					{Name: "ens3", MACAddress: "52:54:00:10:00:11"},
					{Name: "ens4", MACAddress: "52:54:00:20:00:11"},
				}},
			},
		},
	}}}
	interfaces := client.SMSv2Observation{"items": []interface{}{
		map[string]interface{}{
			"name": longKVMRuntimeInterfaceName, "namespace": "system", "resource_version": "rv-7",
			"owner_view": map[string]interface{}{"kind": "securemesh_site_v2", "name": "mcn-ce-ha-smsv2-current-kvm", "namespace": "system", "uid": "site-uid-current"},
			"get_spec": map[string]interface{}{"ethernet_interface": map[string]interface{}{
				"node": "onprem-ce-01-90607", "device": "ens4", "mtu": float64(1500),
				"site_local_inside_network": map[string]interface{}{}, "dhcp_client": map[string]interface{}{},
				"platform_owned_extension": map[string]interface{}{"keep": true},
			}},
		},
	}}
	return site, registrations, interfaces
}

func kvmRuntimeInterfaceTargetFixture() smsv2KVMRuntimeInterfaceTarget {
	return smsv2KVMRuntimeInterfaceTarget{
		Namespace: "system", Site: "mcn-ce-ha-smsv2-current-kvm", Name: longKVMRuntimeInterfaceName,
		ExpectedMAC: "52-54-00-20-00-11", Hostname: "onprem-ce-01-90607", Device: "ens4",
	}
}

func TestSelectSMSv2KVMRuntimeInterfaceAcceptsLongOwnedSLIName(t *testing.T) {
	site, registrations, interfaces := kvmRuntimeInterfaceFixture()
	target := kvmRuntimeInterfaceTargetFixture()
	target.Name, target.Hostname, target.Device = "", "", ""
	got, err := selectSMSv2KVMRuntimeInterface(site, registrations, interfaces, target)
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Name) <= 64 || got.Name != longKVMRuntimeInterfaceName || got.ResourceVersion != "rv-7" {
		t.Fatalf("unexpected runtime interface: %#v", got)
	}
	if got.Hostname != "onprem-ce-01-90607" || got.Device != "ens4" || got.MAC != "52:54:00:20:00:11" || got.OwnerUID != "site-uid-current" {
		t.Fatalf("identity was not preserved: %#v", got)
	}
}

func TestSelectSMSv2KVMRuntimeInterfaceIgnoresUnrelatedInterfaces(t *testing.T) {
	site, registrations, interfaces := kvmRuntimeInterfaceFixture()
	target := kvmRuntimeInterfaceTargetFixture()
	target.Name, target.Hostname, target.Device = "", "", ""
	wanted := interfaces["items"].([]interface{})[0].(map[string]interface{})

	otherSite := deepCopySMSv2Map(wanted)
	otherSite["name"] = "other-site-ens4-0"
	otherSite["owner_view"].(map[string]interface{})["name"] = "other-site"
	otherSite["owner_view"].(map[string]interface{})["uid"] = "other-site-uid"

	slo := deepCopySMSv2Map(wanted)
	slo["name"] = "same-site-slo-ens4-0"
	ethernet := slo["get_spec"].(map[string]interface{})["ethernet_interface"].(map[string]interface{})
	delete(ethernet, "site_local_inside_network")
	ethernet["site_local_network"] = map[string]interface{}{}

	interfaces["items"] = []interface{}{otherSite, slo, wanted}
	got, err := selectSMSv2KVMRuntimeInterface(site, registrations, interfaces, target)
	if err != nil {
		t.Fatal(err)
	}
	if got.Name != longKVMRuntimeInterfaceName {
		t.Fatalf("selected runtime interface %q", got.Name)
	}
}

func TestSelectSMSv2KVMRuntimeInterfaceRejectsIdentityAndOwnershipDrift(t *testing.T) {
	for name, mutate := range map[string]func(client.SMSv2Observation, *client.RegistrationListResponse, client.SMSv2Observation, *smsv2KVMRuntimeInterfaceTarget){
		"malformed name": func(_ client.SMSv2Observation, _ *client.RegistrationListResponse, _ client.SMSv2Observation, target *smsv2KVMRuntimeInterfaceTarget) {
			target.Name = "invalid/name"
		},
		"wrong namespace": func(_ client.SMSv2Observation, _ *client.RegistrationListResponse, _ client.SMSv2Observation, target *smsv2KVMRuntimeInterfaceTarget) {
			target.Namespace = "other"
		},
		"wrong site": func(_ client.SMSv2Observation, _ *client.RegistrationListResponse, _ client.SMSv2Observation, target *smsv2KVMRuntimeInterfaceTarget) {
			target.Site = "other"
		},
		"wrong hostname": func(_ client.SMSv2Observation, _ *client.RegistrationListResponse, _ client.SMSv2Observation, target *smsv2KVMRuntimeInterfaceTarget) {
			target.Hostname = "other"
		},
		"wrong device": func(_ client.SMSv2Observation, _ *client.RegistrationListResponse, _ client.SMSv2Observation, target *smsv2KVMRuntimeInterfaceTarget) {
			target.Device = "ens3"
		},
		"wrong mac": func(_ client.SMSv2Observation, _ *client.RegistrationListResponse, _ client.SMSv2Observation, target *smsv2KVMRuntimeInterfaceTarget) {
			target.ExpectedMAC = "52:54:00:20:00:12"
		},
		"stale owner": func(_ client.SMSv2Observation, _ *client.RegistrationListResponse, interfaces client.SMSv2Observation, _ *smsv2KVMRuntimeInterfaceTarget) {
			interfaces["items"].([]interface{})[0].(map[string]interface{})["owner_view"].(map[string]interface{})["uid"] = "stale"
		},
		"wrong role": func(_ client.SMSv2Observation, _ *client.RegistrationListResponse, interfaces client.SMSv2Observation, _ *smsv2KVMRuntimeInterfaceTarget) {
			ethernet := interfaces["items"].([]interface{})[0].(map[string]interface{})["get_spec"].(map[string]interface{})["ethernet_interface"].(map[string]interface{})
			delete(ethernet, "site_local_inside_network")
			ethernet["site_local_network"] = map[string]interface{}{}
		},
		"duplicate": func(_ client.SMSv2Observation, _ *client.RegistrationListResponse, interfaces client.SMSv2Observation, _ *smsv2KVMRuntimeInterfaceTarget) {
			interfaces["items"] = append(interfaces["items"].([]interface{}), interfaces["items"].([]interface{})[0])
		},
	} {
		t.Run(name, func(t *testing.T) {
			site, registrations, interfaces := kvmRuntimeInterfaceFixture()
			target := kvmRuntimeInterfaceTargetFixture()
			mutate(site, registrations, interfaces, &target)
			if _, err := selectSMSv2KVMRuntimeInterface(site, registrations, interfaces, target); err == nil {
				t.Fatal("invalid runtime interface was accepted")
			}
		})
	}
}

func TestBuildSMSv2KVMParentUpdatePreservesUnrelatedFields(t *testing.T) {
	api := newKVMRuntimeInterfaceAPIFixture(t)
	original := deepCopySMSv2Map(api.configuration["spec"].(map[string]interface{}))
	target := kvmRuntimeInterfaceTargetFixture()
	target.OwnerUID = "site-uid-current"
	request, err := buildSMSv2KVMParentUpdate(api.configuration, target, "10.201.0.11/24", false)
	if err != nil {
		t.Fatal(err)
	}
	if request["resource_version"] != "site-rv-7" {
		t.Fatalf("parent resource version = %#v", request["resource_version"])
	}
	metadata := request["metadata"].(map[string]interface{})
	if metadata["name"] != target.Site || metadata["namespace"] != target.Namespace {
		t.Fatalf("parent metadata = %#v", metadata)
	}
	spec := request["spec"].(map[string]interface{})
	if !reflect.DeepEqual(spec["platform_settings"], map[string]interface{}{"keep": true}) {
		t.Fatal("unrelated parent settings changed")
	}
	nodes := spec["kvm"].(map[string]interface{})["not_managed"].(map[string]interface{})["node_list"].([]interface{})
	interfaces := nodes[0].(map[string]interface{})["interface_list"].([]interface{})
	slo := interfaces[0].(map[string]interface{})
	sli := interfaces[1].(map[string]interface{})
	if _, ok := slo["dhcp_client"]; !ok {
		t.Fatal("primary SLO lost DHCP")
	}
	if _, ok := sli["dhcp_client"]; ok {
		t.Fatal("SLI retained DHCP")
	}
	if got := sli["static_ip"].(map[string]interface{})["ip_address"]; got != "10.201.0.11/24" {
		t.Fatalf("static CIDR = %#v", got)
	}
	if !reflect.DeepEqual(api.configuration["spec"], original) {
		t.Fatal("source observation was mutated")
	}
}

func TestBuildSMSv2KVMParentUpdateRestoresDHCP(t *testing.T) {
	api := newKVMRuntimeInterfaceAPIFixture(t)
	nodes := api.configuration["spec"].(map[string]interface{})["kvm"].(map[string]interface{})["not_managed"].(map[string]interface{})["node_list"].([]interface{})
	interfaces := nodes[0].(map[string]interface{})["interface_list"].([]interface{})
	sli := interfaces[1].(map[string]interface{})
	delete(sli, "dhcp_client")
	sli["static_ip"] = map[string]interface{}{"ip_address": "10.201.0.11/24"}
	target := kvmRuntimeInterfaceTargetFixture()
	target.OwnerUID = "site-uid-current"
	request, err := buildSMSv2KVMParentUpdate(api.configuration, target, "", true)
	if err != nil {
		t.Fatal(err)
	}
	updatedNodes := request["spec"].(map[string]interface{})["kvm"].(map[string]interface{})["not_managed"].(map[string]interface{})["node_list"].([]interface{})
	updated := updatedNodes[0].(map[string]interface{})["interface_list"].([]interface{})[1].(map[string]interface{})
	if _, ok := updated["static_ip"]; ok {
		t.Fatal("static IP remained in DHCP restore request")
	}
	if got, ok := updated["dhcp_client"].(map[string]interface{}); !ok || len(got) != 0 {
		t.Fatalf("DHCP marker = %#v", updated["dhcp_client"])
	}
}

func TestBuildSMSv2KVMParentUpdateRejectsIdentityAndRoleDrift(t *testing.T) {
	for name, mutate := range map[string]func(*kvmRuntimeInterfaceAPIFixture, *smsv2KVMRuntimeInterfaceTarget){
		"missing version": func(api *kvmRuntimeInterfaceAPIFixture, _ *smsv2KVMRuntimeInterfaceTarget) {
			delete(api.configuration, "resource_version")
		},
		"stale owner": func(_ *kvmRuntimeInterfaceAPIFixture, target *smsv2KVMRuntimeInterfaceTarget) {
			target.OwnerUID = "other"
		},
		"wrong device": func(_ *kvmRuntimeInterfaceAPIFixture, target *smsv2KVMRuntimeInterfaceTarget) { target.Device = "ens3" },
		"wrong MAC": func(_ *kvmRuntimeInterfaceAPIFixture, target *smsv2KVMRuntimeInterfaceTarget) {
			target.ExpectedMAC = "52:54:00:20:00:12"
		},
		"wrong role": func(api *kvmRuntimeInterfaceAPIFixture, _ *smsv2KVMRuntimeInterfaceTarget) {
			nodes := api.configuration["spec"].(map[string]interface{})["kvm"].(map[string]interface{})["not_managed"].(map[string]interface{})["node_list"].([]interface{})
			sli := nodes[0].(map[string]interface{})["interface_list"].([]interface{})[1].(map[string]interface{})
			sli["network_option"] = map[string]interface{}{"site_local_network": map[string]interface{}{}}
		},
	} {
		t.Run(name, func(t *testing.T) {
			api := newKVMRuntimeInterfaceAPIFixture(t)
			target := kvmRuntimeInterfaceTargetFixture()
			target.OwnerUID = "site-uid-current"
			mutate(api, &target)
			if _, err := buildSMSv2KVMParentUpdate(api.configuration, target, "10.201.0.11/24", false); err == nil {
				t.Fatal("unsafe parent update accepted")
			}
		})
	}
}

func TestKVMRuntimeInterfaceCIDRValidation(t *testing.T) {
	for _, invalid := range []string{"", "10.201.0.11", "10.201.0.11/33", "2001:db8::1/64", "10.201.0.0/24"} {
		if _, err := canonicalKVMRuntimeInterfaceCIDR(invalid); err == nil {
			t.Fatalf("invalid CIDR %q accepted", invalid)
		}
	}
	if got, err := canonicalKVMRuntimeInterfaceCIDR("10.201.0.11/24"); err != nil || got != "10.201.0.11/24" {
		t.Fatalf("canonical CIDR = %q, %v", got, err)
	}
}

func TestKVMRuntimeInterfaceDiagnosticsAreSecretFree(t *testing.T) {
	site, registrations, interfaces := kvmRuntimeInterfaceFixture()
	target := kvmRuntimeInterfaceTargetFixture()
	interfaces["errors"] = []interface{}{map[string]interface{}{"message": "APIToken super-secret"}}
	_, err := selectSMSv2KVMRuntimeInterface(site, registrations, interfaces, target)
	if err == nil || strings.Contains(err.Error(), "super-secret") {
		t.Fatalf("unsanitized error: %v", err)
	}
}

type kvmRuntimeInterfaceAPIFixture struct {
	t               *testing.T
	mu              sync.Mutex
	object          map[string]interface{}
	missing         bool
	methods         []string
	lastPut         map[string]interface{}
	putCount        int
	parentPutCount  int
	rejectChildPUT  bool
	rejectParentPUT bool
	postCount       int
	deleteCount     int
	registrations   *client.RegistrationListResponse
	configuration   client.SMSv2Observation
}

func newKVMRuntimeInterfaceAPIFixture(t *testing.T) *kvmRuntimeInterfaceAPIFixture {
	site, registrations, interfaces := kvmRuntimeInterfaceFixture()
	site["resource_version"] = "site-rv-7"
	site["spec"] = map[string]interface{}{
		"platform_settings": map[string]interface{}{"keep": true},
		"kvm": map[string]interface{}{"not_managed": map[string]interface{}{"node_list": []interface{}{
			map[string]interface{}{"hostname": "onprem-ce-01-90607", "interface_list": []interface{}{
				map[string]interface{}{"name": "ens3", "ethernet_interface": map[string]interface{}{"device": "ens3", "mac": "52:54:00:10:00:11"}, "network_option": map[string]interface{}{"site_local_network": map[string]interface{}{}}, "dhcp_client": map[string]interface{}{}, "is_primary": true},
				map[string]interface{}{"name": "ens4", "ethernet_interface": map[string]interface{}{"device": "ens4", "mac": "52:54:00:20:00:11"}, "network_option": map[string]interface{}{"site_local_inside_network": map[string]interface{}{}}, "dhcp_client": map[string]interface{}{}, "is_primary": false},
			}},
		}}},
	}
	item := interfaces["items"].([]interface{})[0].(map[string]interface{})
	return &kvmRuntimeInterfaceAPIFixture{
		t: t, configuration: site, registrations: registrations,
		object: map[string]interface{}{
			"metadata":        map[string]interface{}{"name": item["name"], "namespace": item["namespace"], "labels": map[string]interface{}{"platform": "keep"}},
			"system_metadata": map[string]interface{}{"owner_view": item["owner_view"]},
			"spec":            item["get_spec"], "resource_version": item["resource_version"],
		},
	}
}

func (f *kvmRuntimeInterfaceAPIFixture) listObject() map[string]interface{} {
	return map[string]interface{}{
		"name":       f.object["metadata"].(map[string]interface{})["name"],
		"namespace":  f.object["metadata"].(map[string]interface{})["namespace"],
		"owner_view": f.object["system_metadata"].(map[string]interface{})["owner_view"],
		"get_spec":   f.object["spec"], "resource_version": f.object["resource_version"],
	}
}

func (f *kvmRuntimeInterfaceAPIFixture) handler(w http.ResponseWriter, request *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.methods = append(f.methods, request.Method+" "+request.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	switch request.Method {
	case http.MethodPost:
		f.postCount++
	case http.MethodDelete:
		f.deleteCount++
	}
	switch {
	case request.Method == http.MethodGet && request.URL.Path == "/api/config/namespaces/system/securemesh_site_v2s/mcn-ce-ha-smsv2-current-kvm":
		_ = json.NewEncoder(w).Encode(f.configuration)
	case request.Method == http.MethodPut && request.URL.Path == "/api/config/namespaces/system/securemesh_site_v2s/mcn-ce-ha-smsv2-current-kvm":
		var body map[string]interface{}
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			f.t.Fatalf("decode parent PUT: %v", err)
		}
		f.putCount++
		f.parentPutCount++
		f.lastPut = deepCopySMSv2Map(body)
		if f.rejectParentPUT {
			http.Error(w, "stale parent version", http.StatusConflict)
			return
		}
		if body["resource_version"] != f.configuration["resource_version"] {
			http.Error(w, "stale parent version", http.StatusConflict)
			return
		}
		nodes := body["spec"].(map[string]interface{})["kvm"].(map[string]interface{})["not_managed"].(map[string]interface{})["node_list"].([]interface{})
		interfaces := nodes[0].(map[string]interface{})["interface_list"].([]interface{})
		sli := interfaces[1].(map[string]interface{})
		ethernet := f.object["spec"].(map[string]interface{})["ethernet_interface"].(map[string]interface{})
		if static, ok := sli["static_ip"]; ok {
			delete(ethernet, "dhcp_client")
			ethernet["static_ip"] = map[string]interface{}{"node_static_ip": map[string]interface{}{"ip_address": static.(map[string]interface{})["ip_address"]}}
		} else {
			delete(ethernet, "static_ip")
			ethernet["dhcp_client"] = map[string]interface{}{}
		}
		f.configuration["spec"] = body["spec"]
		f.configuration["resource_version"] = "site-rv-next"
		f.object["resource_version"] = "rv-next"
		_ = json.NewEncoder(w).Encode(f.configuration)
	case request.Method == http.MethodGet && request.URL.Path == "/api/register/namespaces/system/registrations_by_site/mcn-ce-ha-smsv2-current-kvm":
		_ = json.NewEncoder(w).Encode(f.registrations)
	case request.Method == http.MethodGet && request.URL.Path == "/api/config/namespaces/system/network_interfaces":
		if f.missing {
			_ = json.NewEncoder(w).Encode(map[string]interface{}{"items": []interface{}{}})
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"items": []interface{}{f.listObject()}})
	case request.Method == http.MethodGet && request.URL.Path == "/api/config/namespaces/system/network_interfaces/"+longKVMRuntimeInterfaceName:
		if f.missing {
			http.NotFound(w, request)
			return
		}
		_ = json.NewEncoder(w).Encode(f.object)
	case request.Method == http.MethodPut && request.URL.Path == "/api/config/namespaces/system/network_interfaces/"+longKVMRuntimeInterfaceName:
		if f.rejectChildPUT {
			http.Error(w, "child object cannot be updated directly", http.StatusForbidden)
			return
		}
		var body map[string]interface{}
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			f.t.Fatalf("decode PUT: %v", err)
		}
		f.putCount++
		f.lastPut = deepCopySMSv2Map(body)
		f.object["metadata"] = body["metadata"]
		f.object["spec"] = body["spec"]
		f.object["resource_version"] = "rv-next"
		_ = json.NewEncoder(w).Encode(f.object)
	default:
		http.NotFound(w, request)
	}
}

func TestSMSv2KVMRuntimeInterfaceWritesThroughOwningSite(t *testing.T) {
	api := newKVMRuntimeInterfaceAPIFixture(t)
	originalNodes := api.configuration["spec"].(map[string]interface{})["kvm"].(map[string]interface{})["not_managed"].(map[string]interface{})["node_list"].([]interface{})
	originalInterfaces := originalNodes[0].(map[string]interface{})["interface_list"].([]interface{})
	originalSLO := deepCopySMSv2Map(originalInterfaces[0].(map[string]interface{}))
	api.rejectChildPUT = true
	server := httptest.NewServer(http.HandlerFunc(api.handler))
	defer server.Close()
	providerResource := &Smsv2KVMRuntimeInterfaceResource{client: client.NewClient(server.URL, "test-token", client.WithMaxRetries(0))}
	ctx := context.Background()
	schemaResponse := &frameworkresource.SchemaResponse{}
	providerResource.Schema(ctx, frameworkresource.SchemaRequest{}, schemaResponse)
	model := kvmRuntimeInterfaceModel()
	model.InterfaceName = types.StringNull()
	model.Hostname = types.StringNull()
	model.Device = types.StringNull()
	response := frameworkresource.CreateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	providerResource.Create(ctx, frameworkresource.CreateRequest{Plan: tfsdk.Plan{
		Schema: schemaResponse.Schema, Raw: responseOperationRaw(t, model, schemaResponse.Schema.Type()),
	}}, &response)
	if response.Diagnostics.HasError() {
		t.Fatalf("owner update failed: %v", response.Diagnostics)
	}
	if api.parentPutCount != 1 || api.putCount != 1 {
		t.Fatalf("expected one parent PUT and no child PUT: parent=%d total=%d", api.parentPutCount, api.putCount)
	}
	spec := api.lastPut["spec"].(map[string]interface{})
	if !reflect.DeepEqual(spec["platform_settings"], map[string]interface{}{"keep": true}) {
		t.Fatal("owner update discarded unrelated site settings")
	}
	nodes := spec["kvm"].(map[string]interface{})["not_managed"].(map[string]interface{})["node_list"].([]interface{})
	interfaces := nodes[0].(map[string]interface{})["interface_list"].([]interface{})
	slo := interfaces[0].(map[string]interface{})
	sli := interfaces[1].(map[string]interface{})
	if !reflect.DeepEqual(slo, originalSLO) {
		t.Fatal("owner update changed the primary SLO")
	}
	if _, ok := sli["dhcp_client"]; ok {
		t.Fatal("owner update left SLI on DHCP")
	}
	if sli["static_ip"].(map[string]interface{})["ip_address"] != "10.201.0.11/24" {
		t.Fatal("owner update used the wrong SLI address")
	}
}

func TestSMSv2KVMRuntimeInterfaceRejectsStaleParentVersionWithoutRetry(t *testing.T) {
	api := newKVMRuntimeInterfaceAPIFixture(t)
	api.rejectChildPUT = true
	api.rejectParentPUT = true
	server := httptest.NewServer(http.HandlerFunc(api.handler))
	defer server.Close()
	providerResource := &Smsv2KVMRuntimeInterfaceResource{client: client.NewClient(server.URL, "test-token", client.WithMaxRetries(0))}
	ctx := context.Background()
	schemaResponse := &frameworkresource.SchemaResponse{}
	providerResource.Schema(ctx, frameworkresource.SchemaRequest{}, schemaResponse)
	model := kvmRuntimeInterfaceModel()
	model.InterfaceName, model.Hostname, model.Device = types.StringNull(), types.StringNull(), types.StringNull()
	response := frameworkresource.CreateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	providerResource.Create(ctx, frameworkresource.CreateRequest{Plan: tfsdk.Plan{
		Schema: schemaResponse.Schema, Raw: responseOperationRaw(t, model, schemaResponse.Schema.Type()),
	}}, &response)
	if !response.Diagnostics.HasError() {
		t.Fatal("stale parent version was accepted")
	}
	if api.parentPutCount != 1 || api.putCount != 1 {
		t.Fatalf("stale version retried or child written: parent=%d total=%d", api.parentPutCount, api.putCount)
	}
}

type firstPutEOFTransport struct {
	base http.RoundTripper
	mu   sync.Mutex
	done bool
}

func (t *firstPutEOFTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	response, err := t.base.RoundTrip(request)
	if err != nil || request.Method != http.MethodPut {
		return response, err
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.done {
		return response, nil
	}
	t.done = true
	_, _ = io.Copy(io.Discard, response.Body)
	_ = response.Body.Close()
	return nil, io.ErrUnexpectedEOF
}

func kvmRuntimeInterfaceModel() Smsv2KVMRuntimeInterfaceResourceModel {
	return Smsv2KVMRuntimeInterfaceResourceModel{
		ID: types.StringNull(), Namespace: types.StringValue("system"), Site: types.StringValue("mcn-ce-ha-smsv2-current-kvm"),
		InterfaceName: types.StringValue(longKVMRuntimeInterfaceName), ExpectedMAC: types.StringValue("52:54:00:20:00:11"),
		Hostname: types.StringValue("onprem-ce-01-90607"), Device: types.StringValue("ens4"), IPv4CIDR: types.StringValue("10.201.0.11/24"),
		OwnerUID: types.StringNull(), ResourceVersion: types.StringNull(), Configured: types.BoolNull(),
	}
}

func TestSMSv2KVMRuntimeInterfaceAdoptsWithPUTReconcilesEOFAndRestoresDHCP(t *testing.T) {
	api := newKVMRuntimeInterfaceAPIFixture(t)
	server := httptest.NewServer(http.HandlerFunc(api.handler))
	defer server.Close()
	httpClient := server.Client()
	httpClient.Timeout = 5 * time.Second
	httpClient.Transport = &firstPutEOFTransport{base: httpClient.Transport}
	providerResource := &Smsv2KVMRuntimeInterfaceResource{client: client.NewClient(server.URL, "test-token", client.WithHTTPClient(httpClient))}
	ctx := context.Background()
	schemaResponse := &frameworkresource.SchemaResponse{}
	providerResource.Schema(ctx, frameworkresource.SchemaRequest{}, schemaResponse)
	model := kvmRuntimeInterfaceModel()
	model.InterfaceName = types.StringNull()
	model.Hostname = types.StringNull()
	model.Device = types.StringNull()
	createResponse := frameworkresource.CreateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	providerResource.Create(ctx, frameworkresource.CreateRequest{Plan: tfsdk.Plan{
		Schema: schemaResponse.Schema, Raw: responseOperationRaw(t, model, schemaResponse.Schema.Type()),
	}}, &createResponse)
	if createResponse.Diagnostics.HasError() {
		t.Fatalf("Create diagnostics: %v", createResponse.Diagnostics)
	}
	var state Smsv2KVMRuntimeInterfaceResourceModel
	createResponse.Diagnostics.Append(createResponse.State.Get(ctx, &state)...)
	if createResponse.Diagnostics.HasError() || !state.Configured.ValueBool() || state.IPv4CIDR.ValueString() != "10.201.0.11/24" || state.InterfaceName.ValueString() != longKVMRuntimeInterfaceName || state.OwnerUID.ValueString() != "site-uid-current" {
		t.Fatalf("unexpected state: %#v diagnostics=%v", state, createResponse.Diagnostics)
	}
	api.mu.Lock()
	if api.putCount != 1 || api.postCount != 0 || api.deleteCount != 0 {
		t.Fatalf("mutation counts after adoption: PUT=%d POST=%d DELETE=%d", api.putCount, api.postCount, api.deleteCount)
	}
	if api.parentPutCount != 1 {
		t.Fatalf("expected parent PUT, got %d", api.parentPutCount)
	}
	api.mu.Unlock()

	deleteResponse := &frameworkresource.DeleteResponse{}
	providerResource.Delete(ctx, frameworkresource.DeleteRequest{State: createResponse.State}, deleteResponse)
	if deleteResponse.Diagnostics.HasError() {
		t.Fatalf("Delete diagnostics: %v", deleteResponse.Diagnostics)
	}
	api.mu.Lock()
	defer api.mu.Unlock()
	if api.putCount != 2 || api.postCount != 0 || api.deleteCount != 0 {
		t.Fatalf("mutation counts after destroy: PUT=%d POST=%d DELETE=%d", api.putCount, api.postCount, api.deleteCount)
	}
	ethernet := api.object["spec"].(map[string]interface{})["ethernet_interface"].(map[string]interface{})
	if _, static := ethernet["static_ip"]; static {
		t.Fatal("destroy left static_ip configured")
	}
	if _, dhcp := ethernet["dhcp_client"]; !dhcp {
		t.Fatal("destroy did not restore DHCP")
	}
}

func TestSMSv2KVMRuntimeInterfaceDestroyRefusesOwnershipDrift(t *testing.T) {
	api := newKVMRuntimeInterfaceAPIFixture(t)
	ethernet := api.object["spec"].(map[string]interface{})["ethernet_interface"].(map[string]interface{})
	delete(ethernet, "dhcp_client")
	ethernet["static_ip"] = map[string]interface{}{"node_static_ip": map[string]interface{}{"ip_address": "10.201.0.11/24"}}
	api.object["system_metadata"].(map[string]interface{})["owner_view"].(map[string]interface{})["uid"] = "stale-owner"
	server := httptest.NewServer(http.HandlerFunc(api.handler))
	defer server.Close()
	providerResource := &Smsv2KVMRuntimeInterfaceResource{client: client.NewClient(server.URL, "test-token", client.WithMaxRetries(0))}
	ctx := context.Background()
	schemaResponse := &frameworkresource.SchemaResponse{}
	providerResource.Schema(ctx, frameworkresource.SchemaRequest{}, schemaResponse)
	model := kvmRuntimeInterfaceModel()
	model.ID = types.StringValue("system/" + longKVMRuntimeInterfaceName)
	model.ResourceVersion = types.StringValue("rv-7")
	model.OwnerUID = types.StringValue("site-uid-current")
	model.Configured = types.BoolValue(true)
	state := tfsdk.State{Schema: schemaResponse.Schema, Raw: responseOperationRaw(t, model, schemaResponse.Schema.Type())}
	response := &frameworkresource.DeleteResponse{}
	providerResource.Delete(ctx, frameworkresource.DeleteRequest{State: state}, response)
	if !response.Diagnostics.HasError() {
		t.Fatal("destroy accepted a stale owner")
	}
	api.mu.Lock()
	defer api.mu.Unlock()
	if api.putCount != 0 || api.postCount != 0 || api.deleteCount != 0 {
		t.Fatalf("ownership failure mutated the API: PUT=%d POST=%d DELETE=%d", api.putCount, api.postCount, api.deleteCount)
	}
}

func TestSMSv2KVMRuntimeInterfaceDestroyAcceptsAlreadyAbsentChild(t *testing.T) {
	api := newKVMRuntimeInterfaceAPIFixture(t)
	api.missing = true
	server := httptest.NewServer(http.HandlerFunc(api.handler))
	defer server.Close()
	providerResource := &Smsv2KVMRuntimeInterfaceResource{client: client.NewClient(server.URL, "test-token", client.WithMaxRetries(0))}
	ctx := context.Background()
	schemaResponse := &frameworkresource.SchemaResponse{}
	providerResource.Schema(ctx, frameworkresource.SchemaRequest{}, schemaResponse)
	model := kvmRuntimeInterfaceModel()
	model.ID = types.StringValue("system/" + longKVMRuntimeInterfaceName)
	model.ResourceVersion = types.StringValue("rv-7")
	model.OwnerUID = types.StringValue("site-uid-current")
	model.Configured = types.BoolValue(true)
	state := tfsdk.State{Schema: schemaResponse.Schema, Raw: responseOperationRaw(t, model, schemaResponse.Schema.Type())}
	response := &frameworkresource.DeleteResponse{}
	providerResource.Delete(ctx, frameworkresource.DeleteRequest{State: state}, response)
	if response.Diagnostics.HasError() {
		t.Fatalf("Delete diagnostics: %v", response.Diagnostics)
	}
	api.mu.Lock()
	defer api.mu.Unlock()
	if api.putCount != 0 || api.postCount != 0 || api.deleteCount != 0 {
		t.Fatalf("absent destroy mutated the API: PUT=%d POST=%d DELETE=%d", api.putCount, api.postCount, api.deleteCount)
	}
}
