// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

func withSMSv2Capabilities(t *testing.T, values map[string]string) {
	t.Helper()
	previous := smsv2ContractCapabilities
	smsv2ContractCapabilities = values
	t.Cleanup(func() { smsv2ContractCapabilities = previous })
}

func TestSMSv2ContractDataSourceRead(t *testing.T) {
	ctx := context.Background()
	dataSource := &Smsv2ContractDataSource{}
	schemaResponse := &datasource.SchemaResponse{}
	dataSource.Schema(ctx, datasource.SchemaRequest{}, schemaResponse)
	config := Smsv2ContractDataSourceModel{
		ID: types.StringNull(), RequiredCapabilities: types.SetNull(types.StringType),
		ContractID: types.StringNull(), ContractVersion: types.StringNull(), APIReleaseTag: types.StringNull(),
		APIReleaseCommit: types.StringNull(), TelemetrySchema: types.StringNull(),
		Capabilities: types.MapNull(types.StringType), F5XCAuthorities: types.ListNull(types.StringType),
		AWSAuthorities:               types.ListNull(types.StringType),
		AzureRouteServerEBGPMultihop: types.ObjectNull(smsv2CapabilityBoundaryAttrTypes),
	}
	response := datasource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	dataSource.Read(ctx, datasource.ReadRequest{Config: tfsdk.Config{
		Schema: schemaResponse.Schema,
		Raw:    responseOperationRaw(t, config, schemaResponse.Schema.Type()),
	}}, &response)
	if response.Diagnostics.HasError() {
		t.Fatalf("contract Read diagnostics: %v", response.Diagnostics)
	}
	var state Smsv2ContractDataSourceModel
	response.Diagnostics.Append(response.State.Get(ctx, &state)...)
	siteUpgrade, hasSiteUpgrade := state.Capabilities.Elements()["site_upgrade"].(types.String)
	azure := state.AzureRouteServerEBGPMultihop.Attributes()
	if response.Diagnostics.HasError() || state.ContractID.ValueString() != "f5xc-smsv2-api/v1" ||
		state.ContractVersion.ValueString() != "7.0.0" || state.TelemetrySchema.ValueString() != "f5xc-smsv2-aws-tgw-telemetry/v2" ||
		!hasSiteUpgrade || siteUpgrade.ValueString() != "available" {
		t.Fatalf("unexpected contract state: %#v diagnostics=%v", state, response.Diagnostics)
	}
	if azure["availability"].(types.String).ValueString() != "unavailable" ||
		azure["enforcement"].(types.String).ValueString() != "reject_before_mutation" ||
		azure["reason"].(types.String).ValueString() != "no_schema_valid_ebgp_multihop_request_control" {
		t.Fatalf("unexpected Azure multihop contract: %#v", azure)
	}
}

func TestSMSv2ContractRequiredCapabilityRejectsDuringValidation(t *testing.T) {
	ctx := context.Background()
	dataSource := &Smsv2ContractDataSource{}
	schemaResponse := &datasource.SchemaResponse{}
	dataSource.Schema(ctx, datasource.SchemaRequest{}, schemaResponse)
	config := Smsv2ContractDataSourceModel{
		ID: types.StringNull(),
		RequiredCapabilities: types.SetValueMust(types.StringType, []attr.Value{
			types.StringValue("azure_route_server_ebgp_multihop"),
		}),
		ContractID: types.StringNull(), ContractVersion: types.StringNull(), APIReleaseTag: types.StringNull(),
		APIReleaseCommit: types.StringNull(), TelemetrySchema: types.StringNull(),
		Capabilities: types.MapNull(types.StringType), F5XCAuthorities: types.ListNull(types.StringType),
		AWSAuthorities:               types.ListNull(types.StringType),
		AzureRouteServerEBGPMultihop: types.ObjectNull(smsv2CapabilityBoundaryAttrTypes),
	}
	response := datasource.ValidateConfigResponse{}
	dataSource.ValidateConfig(ctx, datasource.ValidateConfigRequest{Config: tfsdk.Config{
		Schema: schemaResponse.Schema,
		Raw:    responseOperationRaw(t, config, schemaResponse.Schema.Type()),
	}}, &response)
	if !response.Diagnostics.HasError() {
		t.Fatal("unavailable Azure Route Server eBGP multihop capability was accepted")
	}
	detail := response.Diagnostics.Errors()[0].Detail()
	for _, want := range []string{
		"azure_route_server_ebgp_multihop", "unavailable", "reject_before_mutation",
		"no_schema_valid_ebgp_multihop_request_control", smsv2APIReleaseTag,
		smsv2AzureRouteServerEBGPMultihop.Source.Commit,
	} {
		if !strings.Contains(detail, want) {
			t.Fatalf("planning diagnostic %q does not contain %q", detail, want)
		}
	}
}

func TestSMSv2ContractRequiredCapabilityRejectsPlanWithoutNetworkRequest(t *testing.T) {
	var requests atomic.Int64
	server := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		requests.Add(1)
	}))
	defer server.Close()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"xcsh": providerserver.NewProtocol6WithError(New("test")()),
		},
		Steps: []resource.TestStep{{
			Config: fmt.Sprintf(`
provider "xcsh" {
  api_url   = %q
  api_token = "test-token"
}

data "xcsh_smsv2_contract" "current" {
  required_capabilities = ["azure_route_server_ebgp_multihop"]
}
`, server.URL),
			PlanOnly:    true,
			ExpectError: regexp.MustCompile(`(?s)Required SMSv2 Capability Unavailable.*no_schema_valid_ebgp_multihop_request_control`),
		}},
	})
	if got := requests.Load(); got != 0 {
		t.Fatalf("unavailable capability plan made %d F5 API requests, want 0", got)
	}
}

func runtimeDataSourceConfig(t *testing.T, schemaResponse *datasource.SchemaResponse, nodes types.Map) datasource.ReadRequest {
	t.Helper()
	model := Smsv2AWSRuntimeDataSourceModel{
		ID: types.StringNull(), Namespace: types.StringValue("system"), Site: types.StringValue("lab-site"),
		Nodes: nodes, TimeoutSeconds: types.Int64Value(2), PollIntervalSeconds: types.Int64Value(1),
		Interfaces: types.MapNull(types.ObjectType{AttrTypes: smsv2RuntimeInterfaceAttrTypes}), Healthy: types.BoolNull(),
	}
	return datasource.ReadRequest{Config: tfsdk.Config{
		Schema: schemaResponse.Schema,
		Raw:    responseOperationRaw(t, model, schemaResponse.Schema.Type()),
	}}
}

func runtimeBindings(t *testing.T) types.Map {
	t.Helper()
	value, diagnostics := types.MapValueFrom(context.Background(), types.ObjectType{AttrTypes: map[string]attr.Type{
		"node": types.StringType, "role": types.StringType, "mac": types.StringType,
	}}, map[string]smsv2BindingModel{
		"outside": {Node: types.StringValue("master-0"), Role: types.StringValue("slo"), MAC: types.StringValue("02-AA-BB-CC-DD-01")},
	})
	if diagnostics.HasError() {
		t.Fatalf("encode runtime bindings: %v", diagnostics)
	}
	return value
}

func TestSMSv2AWSRuntimeDataSourceReadAndHTTPFailure(t *testing.T) {
	withSMSv2Capabilities(t, map[string]string{"aws_ce_create": "unavailable", "runtime_status": "available", "tgw_connect": "unavailable"})
	ctx := context.Background()
	fail := false
	transientHealthFailures := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if fail {
			http.Error(w, "fixture failure", http.StatusBadGateway)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		switch request.URL.Path {
		case "/api/config/namespaces/system/sites/lab-site":
			_ = json.NewEncoder(w).Encode(runtimePhysicalLinkStatus())
		case "/api/config/namespaces/system/network_interfaces":
			_ = json.NewEncoder(w).Encode(runtimeInterfaceObjects(runtimeConfiguration()))
		case "/api/config/namespaces/system/securemesh_site_v2s/lab-site":
			_ = json.NewEncoder(w).Encode(runtimeConfiguration())
		case "/api/operate/namespaces/system/sites/lab-site/vpm/debug/global/health":
			if transientHealthFailures > 0 {
				transientHealthFailures--
				http.Error(w, "unresolved DNS", http.StatusBadGateway)
				return
			}
			_, _ = w.Write([]byte("{\"hostname\":\"master-0\",\"state\":\"PROVISIONED\"}"))
		default:
			http.NotFound(w, request)
		}
	}))
	defer server.Close()
	dataSource := &Smsv2AWSRuntimeDataSource{
		client: client.NewClient(server.URL, "test-token", client.WithMaxRetries(0)),
		wait:   func(context.Context, time.Duration) error { return nil },
	}
	schemaResponse := &datasource.SchemaResponse{}
	dataSource.Schema(ctx, datasource.SchemaRequest{}, schemaResponse)
	request := runtimeDataSourceConfig(t, schemaResponse, runtimeBindings(t))
	response := datasource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	dataSource.Read(ctx, request, &response)
	if response.Diagnostics.HasError() {
		t.Fatalf("runtime Read diagnostics: %v", response.Diagnostics)
	}
	var state Smsv2AWSRuntimeDataSourceModel
	response.Diagnostics.Append(response.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() || !state.Healthy.ValueBool() {
		t.Fatalf("unexpected runtime state: %#v diagnostics=%v", state, response.Diagnostics)
	}

	transientHealthFailures = 1
	response = datasource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	dataSource.Read(ctx, request, &response)
	if response.Diagnostics.HasError() || transientHealthFailures != 0 {
		t.Fatalf("runtime Read did not recover from transient health failure: diagnostics=%v", response.Diagnostics)
	}

	fail = true
	nowCalls := 0
	dataSource.now = func() time.Time {
		nowCalls++
		return time.Unix(0, 0).Add(time.Duration(nowCalls) * time.Second)
	}
	response = datasource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	dataSource.Read(ctx, request, &response)
	if !response.Diagnostics.HasError() {
		t.Fatal("runtime Read accepted an HTTP failure")
	}
}

func TestSMSv2DataSourcesDeferNestedUnknowns(t *testing.T) {
	ctx := context.Background()
	runtimeSource := &Smsv2AWSRuntimeDataSource{}
	runtimeSchema := &datasource.SchemaResponse{}
	runtimeSource.Schema(ctx, datasource.SchemaRequest{}, runtimeSchema)
	runtimeResponse := datasource.ReadResponse{State: tfsdk.State{Schema: runtimeSchema.Schema}}
	runtimeRequest := runtimeDataSourceConfig(t, runtimeSchema, types.MapUnknown(types.ObjectType{AttrTypes: map[string]attr.Type{
		"node": types.StringType, "role": types.StringType, "mac": types.StringType,
	}}))
	runtimeRequest.ClientCapabilities.DeferralAllowed = true
	runtimeSource.Read(ctx, runtimeRequest, &runtimeResponse)
	if runtimeResponse.Deferred == nil || runtimeResponse.Diagnostics.HasError() {
		t.Fatalf("runtime unknown was not deferred: deferred=%v diagnostics=%v", runtimeResponse.Deferred, runtimeResponse.Diagnostics)
	}

	bgpSource := &SiteBGPStatusDataSource{}
	bgpSchema := &datasource.SchemaResponse{}
	bgpSource.Schema(ctx, datasource.SchemaRequest{}, bgpSchema)
	model := SiteBGPStatusDataSourceModel{
		ID: types.StringNull(), Namespace: types.StringValue("system"), Site: types.StringValue("lab-site"),
		ExpectedPeers: types.MapUnknown(types.ObjectType{AttrTypes: map[string]attr.Type{
			"node": types.StringType, "role": types.StringType, "mac": types.StringType,
			"peer_address": types.StringType, "expected_imported_routes": types.SetType{ElemType: types.StringType},
		}}),
		ExpectedExportedRoutes: types.SetUnknown(types.StringType),
		TimeoutSeconds:         types.Int64Null(), PollIntervalSeconds: types.Int64Null(),
		Peers:         types.MapNull(types.ObjectType{AttrTypes: smsv2PeerStatusAttrTypes}),
		BGPRoutesJSON: types.StringNull(), SLORoutesJSON: types.StringNull(), SLIRoutesJSON: types.StringNull(), Converged: types.BoolNull(),
	}
	response := datasource.ReadResponse{State: tfsdk.State{Schema: bgpSchema.Schema}}
	bgpSource.Read(ctx, datasource.ReadRequest{
		Config:             tfsdk.Config{Schema: bgpSchema.Schema, Raw: responseOperationRaw(t, model, bgpSchema.Schema.Type())},
		ClientCapabilities: datasource.ReadClientCapabilities{DeferralAllowed: true},
	}, &response)
	if response.Deferred == nil || response.Diagnostics.HasError() {
		t.Fatalf("BGP unknown was not deferred: deferred=%v diagnostics=%v", response.Deferred, response.Diagnostics)
	}
}
