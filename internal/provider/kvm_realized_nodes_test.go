// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

func realizedKVMNodeFixture() map[string]interface{} {
	return map[string]interface{}{
		"hostname": "example-node", "type": "Control", "platform_field": map[string]interface{}{"preserve": true},
		"interface_list": []interface{}{
			map[string]interface{}{"name": "slo", "is_primary": true, "ethernet_interface": map[string]interface{}{"device": "ens3", "mac": "52:54:00:10:00:11"}, "dhcp_client": map[string]interface{}{}, "labels": map[string]interface{}{"role": "outside"}},
			map[string]interface{}{"name": "sli", "ethernet_interface": map[string]interface{}{"device": "ens4", "mac": "52:54:00:10:01:11"}, "static_ip": map[string]interface{}{"node_static_ip": map[string]interface{}{"ipv4_address": "192.0.2.11", "prefix_length": float64(24)}}},
		},
	}
}

func realizedKVMResources() (*client.SecuremeshSiteV2, *client.SecuremeshSiteV2) {
	current := &client.SecuremeshSiteV2{Metadata: client.Metadata{Name: "example-kvm", Namespace: "system"}, ResourceVersion: "rv-reviewed", Spec: map[string]interface{}{"kvm": map[string]interface{}{"not_managed": map[string]interface{}{"node_list": []interface{}{realizedKVMNodeFixture()}, "platform_extension": true}}}}
	desired := &client.SecuremeshSiteV2{Metadata: client.Metadata{Name: "example-kvm", Namespace: "system", Labels: map[string]string{"source": "new"}}, ResourceVersion: "rv-reviewed", Spec: map[string]interface{}{"kvm": map[string]interface{}{"not_managed": map[string]interface{}{}}, "block_all_services": map[string]interface{}{}}}
	return current, desired
}

func TestKVMMetadataUpdatePreservesRealizedNodesOnWire(t *testing.T) {
	current, desired := realizedKVMResources()
	original := deepCopySMSv2Map(current.Spec)
	if err := preserveRealizedKVMNodes(current, desired); err != nil {
		t.Fatal(err)
	}
	var writes int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writes++
		var body client.SecuremeshSiteV2
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Error(err)
			w.WriteHeader(400)
			return
		}
		nodes, _ := nestedMap(body.Spec, "kvm", "not_managed")
		want, _ := nestedMap(original, "kvm", "not_managed")
		if !reflect.DeepEqual(nodes["node_list"], want["node_list"]) {
			w.WriteHeader(400)
			_, _ = w.Write([]byte(`{"message":"nodes configuration error: number of nodes cannot be changed"}`))
			return
		}
		if body.ResourceVersion != "rv-reviewed" || body.Metadata.Labels["source"] != "new" {
			t.Error("reviewed token or new labels lost")
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(body)
	}))
	defer server.Close()
	xc := client.NewClient(server.URL, "test-token", client.WithMaxRetries(0))
	if _, err := xc.UpdateSecuremeshSiteV2(context.Background(), desired); err != nil {
		t.Fatal(err)
	}
	if writes != 1 {
		t.Fatalf("writes=%d", writes)
	}
	if !reflect.DeepEqual(current.Spec, original) {
		t.Fatal("fresh read mutated")
	}
	if _, ok := desired.Spec["block_all_services"]; !ok {
		t.Fatal("unrelated desired spec lost")
	}
}

func TestKVMRealizedNodesPreservationRejectsStaleOrForeignRead(t *testing.T) {
	for _, mutation := range []func(*client.SecuremeshSiteV2){
		func(c *client.SecuremeshSiteV2) { c.ResourceVersion = "rv-new" },
		func(c *client.SecuremeshSiteV2) { c.Metadata.Name = "foreign" },
		func(c *client.SecuremeshSiteV2) { c.Metadata.Namespace = "foreign" },
		func(c *client.SecuremeshSiteV2) { delete(c.Spec, "kvm") },
	} {
		current, desired := realizedKVMResources()
		mutation(current)
		if err := preserveRealizedKVMNodes(current, desired); err == nil {
			t.Fatal("stale or foreign read accepted")
		}
	}
}

func TestKVMRealizedNodesExplicitConfigurationAndOtherPlatformsStayUntouched(t *testing.T) {
	current, desired := realizedKVMResources()
	nodes, _ := nestedMap(desired.Spec, "kvm", "not_managed")
	explicit := []interface{}{map[string]interface{}{"hostname": "configured"}}
	nodes["node_list"] = explicit
	if err := preserveRealizedKVMNodes(current, desired); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(nodes["node_list"], explicit) {
		t.Fatal("explicit node configuration overwritten")
	}
	delete(desired.Spec, "kvm")
	desired.Spec["azure"] = map[string]interface{}{}
	before := deepCopySMSv2Map(desired.Spec)
	if err := preserveRealizedKVMNodes(nil, desired); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(before, desired.Spec) {
		t.Fatal("unrelated platform changed")
	}
}
