package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	frameworktimeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

func nullResourceTimeouts() frameworktimeouts.Value {
	return frameworktimeouts.Value{Object: types.ObjectNull(map[string]attr.Type{
		"create": types.StringType,
		"read":   types.StringType,
		"update": types.StringType,
		"delete": types.StringType,
	})}
}

func TestRestoredNamespaceDeleteUsesCascadeEndpoint(t *testing.T) {
	t.Parallel()

	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requests++
		if request.Method != http.MethodPost {
			t.Errorf("request method = %s, want POST", request.Method)
		}
		if request.URL.Path != "/api/web/namespaces/fixture/cascade_delete" {
			t.Errorf("request path = %s", request.URL.Path)
		}
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{}`))
	}))
	defer server.Close()

	implementation := &NamespaceResource{client: client.NewClient(server.URL, "fixture-token", client.WithMaxRetries(0))}
	schemaResponse := &resource.SchemaResponse{}
	implementation.Schema(context.Background(), resource.SchemaRequest{}, schemaResponse)
	state := tfsdk.State{Schema: schemaResponse.Schema}
	model := NamespaceResourceModel{
		Name:        types.StringValue("fixture"),
		Namespace:   types.StringNull(),
		Annotations: types.MapNull(types.StringType),
		Description: types.StringNull(),
		Disable:     types.BoolNull(),
		Labels:      types.MapNull(types.StringType),
		ID:          types.StringValue("fixture"),
		Timeouts:    nullResourceTimeouts(),
	}
	if diagnostics := state.Set(context.Background(), &model); diagnostics.HasError() {
		t.Fatalf("encode namespace state: %v", diagnostics)
	}
	response := &resource.DeleteResponse{}
	implementation.Delete(context.Background(), resource.DeleteRequest{State: state}, response)
	if response.Diagnostics.HasError() {
		t.Fatalf("cascade delete returned diagnostics: %v", response.Diagnostics)
	}
	if requests != 1 {
		t.Fatalf("cascade delete made %d requests, want 1", requests)
	}
}

func TestRestoredBotInfrastructureUnsupportedDeleteRetainsOwnership(t *testing.T) {
	t.Parallel()

	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requests++
		if request.Method != http.MethodDelete {
			t.Errorf("request method = %s, want DELETE", request.Method)
		}
		if request.URL.Path != "/api/shape/bot/namespaces/system/bot_infrastructures/fixture" {
			t.Errorf("request path = %s", request.URL.Path)
		}
		http.Error(writer, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}))
	defer server.Close()

	implementation := &BotInfrastructureResource{client: client.NewClient(server.URL, "fixture-token", client.WithMaxRetries(0))}
	schemaResponse := &resource.SchemaResponse{}
	implementation.Schema(context.Background(), resource.SchemaRequest{}, schemaResponse)
	state := tfsdk.State{Schema: schemaResponse.Schema}
	model := BotInfrastructureResourceModel{
		Name:              types.StringValue("fixture"),
		Namespace:         types.StringValue("system"),
		Annotations:       types.MapNull(types.StringType),
		Description:       types.StringNull(),
		Disable:           types.BoolNull(),
		Labels:            types.MapNull(types.StringType),
		ID:                types.StringValue("fixture"),
		TrafficType:       types.StringNull(),
		Timeouts:          nullResourceTimeouts(),
		CreateCloudHosted: nil,
	}
	if diagnostics := state.Set(context.Background(), &model); diagnostics.HasError() {
		t.Fatalf("encode bot infrastructure state: %v", diagnostics)
	}
	response := &resource.DeleteResponse{}
	implementation.Delete(context.Background(), resource.DeleteRequest{State: state}, response)
	if !response.Diagnostics.HasError() {
		t.Fatal("unsupported DELETE was treated as success; Terraform would drop ownership")
	}
	if requests != 1 {
		t.Fatalf("unsupported DELETE made %d requests, want 1", requests)
	}
}
