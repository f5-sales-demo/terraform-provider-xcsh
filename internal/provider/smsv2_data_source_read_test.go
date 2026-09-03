// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

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
	response := datasource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	dataSource.Read(ctx, datasource.ReadRequest{}, &response)
	if response.Diagnostics.HasError() {
		t.Fatalf("contract Read diagnostics: %v", response.Diagnostics)
	}
	var state Smsv2ContractDataSourceModel
	response.Diagnostics.Append(response.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() || state.ContractID.ValueString() != "f5xc-ce-automation/v3" ||
		state.ContractVersion.ValueString() != "6.0.0" || state.TelemetrySchema.ValueString() != "f5xc-smsv2-aws-tgw-telemetry/v2" {
		t.Fatalf("unexpected contract state: %#v diagnostics=%v", state, response.Diagnostics)
	}
}

func runtimeDataSourceConfig(t *testing.T, schemaResponse *datasource.SchemaResponse, nodes types.Map) datasource.ReadRequest {
	t.Helper()
	model := Smsv2AWSRuntimeDataSourceModel{
		ID: types.StringNull(), Namespace: types.StringValue("system"), Site: types.StringValue("lab-site"),
		Nodes: nodes, Interfaces: types.MapNull(types.ObjectType{AttrTypes: smsv2RuntimeInterfaceAttrTypes}), Healthy: types.BoolNull(),
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
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if fail {
			http.Error(w, "fixture failure", http.StatusBadGateway)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		switch request.URL.Path {
		case "/api/config/namespaces/system/securemesh_site_v2s/lab-site":
			_ = json.NewEncoder(w).Encode(runtimeConfiguration())
		case "/api/operate/namespaces/system/sites/lab-site/vpm/debug/global/health":
			_, _ = w.Write([]byte("{\"hostname\":\"master-0\",\"state\":\"PROVISIONED\"}"))
		default:
			http.NotFound(w, request)
		}
	}))
	defer server.Close()
	dataSource := &Smsv2AWSRuntimeDataSource{client: client.NewClient(server.URL, "test-token", client.WithMaxRetries(0))}
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

	fail = true
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
			"peer_address": types.StringType, "expected_routes": types.SetType{ElemType: types.StringType},
		}}),
		TimeoutSeconds: types.Int64Null(), PollIntervalSeconds: types.Int64Null(),
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
