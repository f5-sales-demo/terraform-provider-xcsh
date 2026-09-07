// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

func bgpExpectedPeers(t *testing.T) types.Map {
	t.Helper()
	value, diagnostics := types.MapValueFrom(context.Background(), types.ObjectType{AttrTypes: map[string]attr.Type{
		"node": types.StringType, "role": types.StringType, "mac": types.StringType,
		"peer_address": types.StringType, "expected_routes": types.SetType{ElemType: types.StringType},
	}}, map[string]smsv2ExpectedPeerModel{
		"outside": {
			Node: types.StringValue("master-0"), Role: types.StringValue("slo"),
			MAC: types.StringValue("02-AA-BB-CC-DD-01"), PeerAddress: types.StringValue("169.254.10.1"),
			ExpectedRoutes: types.SetValueMust(types.StringType, []attr.Value{types.StringValue("10.10.0.0/16")}),
		},
	})
	if diagnostics.HasError() {
		t.Fatalf("encode expected peers: %v", diagnostics)
	}
	return value
}

func bgpReadRequest(t *testing.T, schemaResponse *datasource.SchemaResponse, expected types.Map, timeout, poll int64) datasource.ReadRequest {
	t.Helper()
	model := SiteBGPStatusDataSourceModel{
		ID: types.StringNull(), Namespace: types.StringValue("system"), Site: types.StringValue("lab-site"),
		ExpectedPeers: expected, TimeoutSeconds: types.Int64Value(timeout), PollIntervalSeconds: types.Int64Value(poll),
		Peers:         types.MapNull(types.ObjectType{AttrTypes: smsv2PeerStatusAttrTypes}),
		BGPRoutesJSON: types.StringNull(), SLORoutesJSON: types.StringNull(), SLIRoutesJSON: types.StringNull(), Converged: types.BoolNull(),
	}
	return datasource.ReadRequest{Config: tfsdk.Config{
		Schema: schemaResponse.Schema,
		Raw:    responseOperationRaw(t, model, schemaResponse.Schema.Type()),
	}}
}

func bgpPeerFixture(state string) map[string]interface{} {
	return map[string]interface{}{"ver": []interface{}{map[string]interface{}{
		"name": "master-0.example.internal", "peer": []interface{}{map[string]interface{}{
			"interface_name": "tap-bgp-observation", "peer_address": map[string]interface{}{"ipv4": map[string]interface{}{"addr": "169.254.10.1"}},
			"protocol_status": state, "up_down_timestamp": "2026-09-03T00:00:00Z",
			"received_prefix_count": float64(1), "advertised_prefix_count": float64(1),
		}},
	}}}
}

func TestPeerAddressLiveShape(t *testing.T) {
	for name, fixture := range map[string]interface{}{
		"ipv4 object":          map[string]interface{}{"ipv4": map[string]interface{}{"addr": "169.254.10.1"}},
		"ipv6 object":          map[string]interface{}{"ipv6": map[string]interface{}{"addr": "2001:db8::1"}},
		"legacy direct string": map[string]interface{}{"ipv4": "169.254.10.1"},
	} {
		t.Run(name, func(t *testing.T) {
			if got := peerAddress(fixture); got == "" {
				t.Fatal("peer address was not extracted")
			}
		})
	}
}

func bgpRoutesFixture() map[string]interface{} {
	return map[string]interface{}{"ver": []interface{}{map[string]interface{}{
		"name": "master-0.example.internal", "ri_table": []interface{}{map[string]interface{}{
			"rt_table": []interface{}{map[string]interface{}{
				"imported": []interface{}{map[string]interface{}{"subnet": "10.10.0.0/16", "path": []interface{}{map[string]interface{}{"peer": "169.254.10.1"}, map[string]interface{}{"peer": "169.254.10.2"}}}},
				"exported": []interface{}{map[string]interface{}{"subnet": "10.20.0.0/16"}},
			}},
		}},
	}}}
}

func simplifiedRoutesFixture() map[string]interface{} {
	return map[string]interface{}{"ver_routes": []interface{}{map[string]interface{}{
		"node": "master-0.example.internal", "route": []interface{}{map[string]interface{}{"prefix": "10.20.0.0/16"}},
	}}}
}

func bgpFixtureServer(t *testing.T, state *string, fail *bool) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if *fail {
			http.Error(w, "fixture failure", http.StatusBadGateway)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		var payload interface{}
		switch request.URL.Path {
		case "/api/config/namespaces/system/network_interfaces":
			_ = json.NewEncoder(w).Encode(runtimeInterfaceObjects(runtimeConfiguration()))
			return
		case "/api/config/namespaces/system/securemesh_site_v2s/lab-site":
			payload = runtimeConfiguration()
		case "/api/operate/namespaces/system/sites/lab-site/ver/bgp_peers":
			payload = bgpPeerFixture(*state)
		case "/api/operate/namespaces/system/sites/lab-site/ver/bgp_routes":
			payload = bgpRoutesFixture()
		case "/api/operate/namespaces/system/sites/lab-site/ver/simplified_routes":
			if request.Method != http.MethodPost {
				t.Errorf("simplified route method = %s", request.Method)
			}
			var body map[string]interface{}
			if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
				t.Errorf("decode simplified route body: %v", err)
			}
			if _, slo := body["slo"]; !slo {
				if _, sli := body["sli"]; !sli {
					t.Errorf("missing route role selector: %#v", body)
				}
			}
			payload = simplifiedRoutesFixture()
		default:
			http.NotFound(w, request)
			return
		}
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			t.Errorf("encode fixture: %v", err)
		}
	}))
}

func TestSiteBGPStatusDataSourceRead(t *testing.T) {
	withSMSv2Capabilities(t, map[string]string{"runtime_status": "available", "tgw_connect": "available"})
	state, fail := "Established", false
	server := bgpFixtureServer(t, &state, &fail)
	defer server.Close()

	ctx := context.Background()
	dataSource := &SiteBGPStatusDataSource{client: client.NewClient(server.URL, "test-token", client.WithMaxRetries(0))}
	schemaResponse := &datasource.SchemaResponse{}
	dataSource.Schema(ctx, datasource.SchemaRequest{}, schemaResponse)
	response := datasource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	dataSource.Read(ctx, bgpReadRequest(t, schemaResponse, bgpExpectedPeers(t), 5, 1), &response)
	if response.Diagnostics.HasError() {
		t.Fatalf("BGP Read diagnostics: %v", response.Diagnostics)
	}
	var got SiteBGPStatusDataSourceModel
	response.Diagnostics.Append(response.State.Get(ctx, &got)...)
	if response.Diagnostics.HasError() || !got.Converged.ValueBool() || got.ID.ValueString() != "system/lab-site" {
		t.Fatalf("unexpected BGP state: %#v diagnostics=%v", got, response.Diagnostics)
	}
	peers := map[string]smsv2PeerStatusModel{}
	response.Diagnostics.Append(got.Peers.ElementsAs(ctx, &peers, false)...)
	if response.Diagnostics.HasError() || peers["outside"].StateChangedAt.ValueString() != "2026-09-03T00:00:00Z" {
		t.Fatalf("unexpected peer state: %#v diagnostics=%v", peers, response.Diagnostics)
	}
}

func TestSiteBGPStatusPollingBoundsAndCancellation(t *testing.T) {
	withSMSv2Capabilities(t, map[string]string{"runtime_status": "available", "tgw_connect": "available"})
	state, fail := "Idle", false
	server := bgpFixtureServer(t, &state, &fail)
	defer server.Close()
	ctx := context.Background()

	t.Run("caps final sleep and times out exactly", func(t *testing.T) {
		now := time.Date(2026, 9, 3, 0, 0, 0, 0, time.UTC)
		var delays []time.Duration
		dataSource := &SiteBGPStatusDataSource{
			client: client.NewClient(server.URL, "test-token", client.WithMaxRetries(0)),
			now:    func() time.Time { return now },
			wait: func(_ context.Context, delay time.Duration) error {
				delays = append(delays, delay)
				now = now.Add(delay)
				return nil
			},
		}
		schemaResponse := &datasource.SchemaResponse{}
		dataSource.Schema(ctx, datasource.SchemaRequest{}, schemaResponse)
		response := datasource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
		dataSource.Read(ctx, bgpReadRequest(t, schemaResponse, bgpExpectedPeers(t), 5, 3), &response)
		if len(delays) != 2 || delays[0] != 3*time.Second || delays[1] != 2*time.Second || !now.Equal(time.Date(2026, 9, 3, 0, 0, 5, 0, time.UTC)) {
			t.Fatalf("delays=%v now=%s", delays, now)
		}
		if !response.Diagnostics.HasError() || !strings.Contains(fmt.Sprint(response.Diagnostics), "Idle") {
			t.Fatalf("timeout lost convergence reason: %v", response.Diagnostics)
		}
	})

	t.Run("cancellation", func(t *testing.T) {
		var delay time.Duration
		dataSource := &SiteBGPStatusDataSource{
			client: client.NewClient(server.URL, "test-token", client.WithMaxRetries(0)),
			now:    func() time.Time { return time.Date(2026, 9, 3, 0, 0, 0, 0, time.UTC) },
			wait: func(_ context.Context, got time.Duration) error {
				delay = got
				return context.Canceled
			},
		}
		schemaResponse := &datasource.SchemaResponse{}
		dataSource.Schema(ctx, datasource.SchemaRequest{}, schemaResponse)
		response := datasource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
		dataSource.Read(ctx, bgpReadRequest(t, schemaResponse, bgpExpectedPeers(t), 5, 3), &response)
		if !response.Diagnostics.HasError() || delay != 3*time.Second || !strings.Contains(fmt.Sprint(response.Diagnostics), "context canceled") {
			t.Fatalf("cancellation diagnostics=%v delay=%s", response.Diagnostics, delay)
		}
	})

	t.Run("HTTP failure is bounded", func(t *testing.T) {
		fail = true
		t.Cleanup(func() { fail = false })
		now := time.Date(2026, 9, 3, 0, 0, 0, 0, time.UTC)
		dataSource := &SiteBGPStatusDataSource{
			client: client.NewClient(server.URL, "test-token", client.WithMaxRetries(0)),
			now:    func() time.Time { return now },
			wait: func(_ context.Context, delay time.Duration) error {
				now = now.Add(delay)
				return nil
			},
		}
		schemaResponse := &datasource.SchemaResponse{}
		dataSource.Schema(ctx, datasource.SchemaRequest{}, schemaResponse)
		response := datasource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
		dataSource.Read(ctx, bgpReadRequest(t, schemaResponse, bgpExpectedPeers(t), 1, 1), &response)
		if !response.Diagnostics.HasError() || !strings.Contains(fmt.Sprint(response.Diagnostics), "502") {
			t.Fatalf("HTTP failure diagnostics=%v", response.Diagnostics)
		}
	})
}

func TestSiteBGPStatusSessionIdentity(t *testing.T) {
	withSMSv2Capabilities(t, map[string]string{"runtime_status": "available", "tgw_connect": "available"})
	for _, tc := range []struct {
		name, address, mac string
		duplicate          bool
	}{
		{"two sessions on one interface", "169.254.10.2", "02:aa:bb:cc:dd:01", false},
		{"duplicate session", "169.254.10.1", "02:aa:bb:cc:dd:01", true},
		{"duplicate session with different MAC", "169.254.10.1", "02:aa:bb:cc:dd:02", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			expected := bgpExpectedPeers(t)
			values := map[string]smsv2ExpectedPeerModel{}
			if diags := expected.ElementsAs(ctx, &values, false); diags.HasError() {
				t.Fatal(diags)
			}
			second := values["outside"]
			second.PeerAddress = types.StringValue(tc.address)
			second.MAC = types.StringValue(tc.mac)
			values["second"] = second
			expected, diags := types.MapValueFrom(ctx, expected.ElementType(ctx), values)
			if diags.HasError() {
				t.Fatal(diags)
			}
			requests := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requests++
				var payload interface{}
				switch r.URL.Path {
				case "/api/config/namespaces/system/network_interfaces":
					_ = json.NewEncoder(w).Encode(runtimeInterfaceObjects(runtimeConfiguration()))
					return
				case "/api/config/namespaces/system/securemesh_site_v2s/lab-site":
					payload = runtimeConfiguration()
				case "/api/operate/namespaces/system/sites/lab-site/ver/bgp_peers":
					payload = map[string]interface{}{"ver": []interface{}{map[string]interface{}{
						"name": "master-0.example.internal", "peer": []interface{}{
							map[string]interface{}{"peer_address": "169.254.10.1", "protocol_status": "Established"},
							map[string]interface{}{"peer_address": "169.254.10.2", "protocol_status": "Established"},
						},
					}}}
				case "/api/operate/namespaces/system/sites/lab-site/ver/bgp_routes":
					payload = bgpRoutesFixture()
				case "/api/operate/namespaces/system/sites/lab-site/ver/simplified_routes":
					payload = simplifiedRoutesFixture()
				default:
					t.Errorf("unexpected path %s", r.URL.Path)
					http.NotFound(w, r)
					return
				}
				if err := json.NewEncoder(w).Encode(payload); err != nil {
					t.Error(err)
				}
			}))
			defer server.Close()
			source := &SiteBGPStatusDataSource{client: client.NewClient(server.URL, "test-token", client.WithMaxRetries(0))}
			schemaResponse := &datasource.SchemaResponse{}
			source.Schema(ctx, datasource.SchemaRequest{}, schemaResponse)
			response := datasource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
			source.Read(ctx, bgpReadRequest(t, schemaResponse, expected, 1, 1), &response)
			if tc.duplicate {
				if !response.Diagnostics.HasError() || !strings.Contains(fmt.Sprint(response.Diagnostics), "Duplicate BGP Session") || requests != 0 {
					t.Fatalf("duplicate must fail before HTTP: requests=%d diagnostics=%v", requests, response.Diagnostics)
				}
				return
			}
			if response.Diagnostics.HasError() {
				t.Fatal(response.Diagnostics)
			}
			var result SiteBGPStatusDataSourceModel
			if diags := response.State.Get(ctx, &result); diags.HasError() {
				t.Fatal(diags)
			}
			if !result.Converged.ValueBool() || len(result.Peers.Elements()) != 2 {
				t.Fatalf("expected two converged sessions: %#v", result)
			}
		})
	}
}

func TestSMSv2PeerAddressIdentity(t *testing.T) {
	for _, tc := range []struct {
		left, right string
		equal       bool
	}{
		{"2001:0db8:0:0:0:0:0:1", "2001:db8::1", true},
		{"192.0.2.1", "::ffff:192.0.2.1", true},
		{"192.0.2.1", "192.0.2.2", false},
		{"", "", false},
		{"invalid", "invalid", false},
	} {
		if got := sameSMSv2PeerAddress(tc.left, tc.right); got != tc.equal {
			t.Errorf("address identity (%q, %q) = %t", tc.left, tc.right, got)
		}
	}
}
