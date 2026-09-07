// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func runtimePhysicalLinkStatus() client.SMSv2Observation {
	return client.SMSv2Observation{
		"metadata": map[string]interface{}{"name": "lab-site", "namespace": "system"},
		"status": []interface{}{map[string]interface{}{
			"metadata": map[string]interface{}{"creator_class": "ver", "creator_id": "master-0", "vtrp_stale": false, "publish": "STATUS_PUBLISH"},
			"ver_status": map[string]interface{}{"ver_instance_name": "opaque-observed-instance", "intf_status": []interface{}{
				map[string]interface{}{"name": "eth0", "mac": "02:aa:bb:cc:dd:01", "link_state": true, "link_type": "LINK_TYPE_ETHERNET", "active_state": "STATE_ACTIVE", "network_type": "VIRTUAL_NETWORK_SITE_LOCAL"},
			}},
		}},
	}
}

func TestSMSv2RuntimeRequiresPhysicalLinkHealth(t *testing.T) {
	withSMSv2Capabilities(t, map[string]string{"runtime_status": "available"})
	for _, tc := range []struct {
		name    string
		change  func(client.SMSv2Observation)
		healthy bool
	}{
		{"up", func(client.SMSv2Observation) {}, true},
		{"down", func(v client.SMSv2Observation) { physicalFixtureInterface(v)["link_state"] = false }, false},
		{"missing link state", func(v client.SMSv2Observation) { delete(physicalFixtureInterface(v), "link_state") }, false},
		{"wrong MAC", func(v client.SMSv2Observation) { physicalFixtureInterface(v)["mac"] = "02:aa:bb:cc:dd:99" }, false},
		{"wrong device", func(v client.SMSv2Observation) { physicalFixtureInterface(v)["name"] = "other-device" }, false},
		{"stale", func(v client.SMSv2Observation) { physicalFixtureMetadata(v)["vtrp_stale"] = true }, false},
		{"foreign publisher", func(v client.SMSv2Observation) { physicalFixtureMetadata(v)["creator_id"] = "foreign-node" }, false},
		{"foreign site", func(v client.SMSv2Observation) { v["metadata"].(map[string]interface{})["name"] = "other-site" }, false},
		{"duplicate node status", func(v client.SMSv2Observation) {
			v["status"] = append(v["status"].([]interface{}), v["status"].([]interface{})[0])
		}, false},
		{"unpublished", func(v client.SMSv2Observation) { physicalFixtureMetadata(v)["publish"] = "STATUS_UNPUBLISH" }, false},
		{"unknown staleness", func(v client.SMSv2Observation) { delete(physicalFixtureMetadata(v), "vtrp_stale") }, false},
		{"wrong publisher class", func(v client.SMSv2Observation) { physicalFixtureMetadata(v)["creator_class"] = "other" }, false},
		{"wrong network", func(v client.SMSv2Observation) {
			physicalFixtureInterface(v)["network_type"] = "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE"
		}, false},
		{"wrong link type", func(v client.SMSv2Observation) { physicalFixtureInterface(v)["link_type"] = "LINK_TYPE_UNKNOWN" }, false},
		{"malformed link state", func(v client.SMSv2Observation) { physicalFixtureInterface(v)["link_state"] = "true" }, false},
		{"foreign namespace", func(v client.SMSv2Observation) { v["metadata"].(map[string]interface{})["namespace"] = "other" }, false},
		{"other status document", func(v client.SMSv2Observation) {
			v["status"] = append(v["status"].([]interface{}), map[string]interface{}{"metadata": physicalFixtureMetadata(v), "ver_status": nil})
		}, true},
		{"backup link up", func(v client.SMSv2Observation) { physicalFixtureInterface(v)["active_state"] = "STATE_BACKUP" }, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			status := runtimePhysicalLinkStatus()
			tc.change(status)
			reads := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var value interface{}
				switch r.URL.Path {
				case "/api/config/namespaces/system/securemesh_site_v2s/lab-site":
					value = runtimeConfiguration()
				case "/api/config/namespaces/system/network_interfaces":
					value = runtimeInterfaceObjects(runtimeConfiguration())
				case "/api/operate/namespaces/system/sites/lab-site/vpm/debug/global/health":
					value = map[string]interface{}{"hostname": "master-0", "state": "PROVISIONED"}
				case "/api/config/namespaces/system/sites/lab-site":
					reads++
					value = status
				default:
					t.Errorf("unexpected path %s", r.URL.Path)
					http.NotFound(w, r)
					return
				}
				_ = json.NewEncoder(w).Encode(value)
			}))
			defer server.Close()
			now := time.Now()
			d := &Smsv2AWSRuntimeDataSource{client: client.NewClient(server.URL, "test-token", client.WithMaxRetries(0)), now: func() time.Time { return now }, wait: func(_ context.Context, delay time.Duration) error { now = now.Add(delay); return nil }}
			ctx := context.Background()
			schema := &datasource.SchemaResponse{}
			d.Schema(ctx, datasource.SchemaRequest{}, schema)
			req := runtimeDataSourceConfig(t, schema, runtimeBindings(t))
			var config Smsv2AWSRuntimeDataSourceModel
			if diags := req.Config.Get(ctx, &config); diags.HasError() {
				t.Fatal(diags)
			}
			config.TimeoutSeconds = types.Int64Value(1)
			config.PollIntervalSeconds = types.Int64Value(1)
			req.Config.Raw = responseOperationRaw(t, config, schema.Schema.Type())
			resp := datasource.ReadResponse{State: tfsdk.State{Schema: schema.Schema}}
			d.Read(ctx, req, &resp)
			if resp.Diagnostics.HasError() == tc.healthy {
				t.Fatalf("physical health=%v diagnostics=%v", tc.healthy, resp.Diagnostics)
			}
			if reads < 1 || reads > 2 {
				t.Fatalf("expected bounded physical link reads, got %d", reads)
			}
			if tc.healthy {
				var got Smsv2AWSRuntimeDataSourceModel
				if diags := resp.State.Get(ctx, &got); diags.HasError() || !got.Healthy.ValueBool() {
					t.Fatalf("healthy physical link was not retained in state: %v", diags)
				}
			}
		})
	}
}

func physicalFixtureMetadata(v client.SMSv2Observation) map[string]interface{} {
	return v["status"].([]interface{})[0].(map[string]interface{})["metadata"].(map[string]interface{})
}
func physicalFixtureInterface(v client.SMSv2Observation) map[string]interface{} {
	return v["status"].([]interface{})[0].(map[string]interface{})["ver_status"].(map[string]interface{})["intf_status"].([]interface{})[0].(map[string]interface{})
}

func TestSMSv2PhysicalHealthRequiresEveryRequestedNode(t *testing.T) {
	configuration := multiNodeRuntimeConfiguration()
	configured, err := resolvedRuntimeFixture(configuration)
	if err != nil {
		t.Fatal(err)
	}
	bindings := map[string]smsv2BindingModel{}
	observation := runtimePhysicalLinkStatus()
	observation["status"] = []interface{}{}
	byNode := map[string]map[string]interface{}{}
	for _, iface := range configured {
		bindings[iface.Node+"/"+iface.Role] = smsv2BindingModel{Node: types.StringValue(iface.Node), Role: types.StringValue(iface.Role), MAC: types.StringValue(iface.MAC)}
		status, exists := byNode[iface.Node]
		if !exists {
			status = runtimePhysicalLinkStatus()["status"].([]interface{})[0].(map[string]interface{})
			status["metadata"].(map[string]interface{})["creator_id"] = iface.Node
			status["ver_status"].(map[string]interface{})["intf_status"] = []interface{}{}
			byNode[iface.Node] = status
			observation["status"] = append(observation["status"].([]interface{}), status)
		}
		link := physicalFixtureInterface(runtimePhysicalLinkStatus())
		link["name"] = iface.Device
		link["mac"] = iface.MAC
		if iface.Role == "sli" {
			link["network_type"] = "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE"
		}
		ver := status["ver_status"].(map[string]interface{})
		ver["intf_status"] = append(ver["intf_status"].([]interface{}), link)
	}
	if err := validateSMSv2PhysicalLinks(configuration, observation, configured, bindings); err != nil {
		t.Fatal(err)
	}
	all := observation["status"].([]interface{})
	observation["status"] = all[1:]
	if err := validateSMSv2PhysicalLinks(configuration, observation, configured, bindings); err == nil {
		t.Fatal("one healthy node cannot establish health for a missing node")
	}
	observation["status"] = all
	links := byNode["master-1"]["ver_status"].(map[string]interface{})["intf_status"].([]interface{})
	links[1].(map[string]interface{})["link_state"] = false
	if err := validateSMSv2PhysicalLinks(configuration, observation, configured, bindings); err == nil {
		t.Fatal("other healthy links masked one failed SLI link")
	}
	links[1].(map[string]interface{})["link_state"] = true
	byNode["master-1"]["ver_status"].(map[string]interface{})["intf_status"] = append(links, links[0])
	if err := validateSMSv2PhysicalLinks(configuration, observation, configured, bindings); err == nil {
		t.Fatal("duplicate physical device observations were accepted")
	}
}
