// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

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

func TestAlertPolicyDeleteConfirmsAbsenceAfterServerError(t *testing.T) {
	t.Parallel()

	var deleteCalls, getCalls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/api/config/namespaces/system/alert_policys/fixture" {
			t.Errorf("request path = %s", request.URL.Path)
		}
		switch request.Method {
		case http.MethodDelete:
			deleteCalls++
			http.Error(w, "server error", http.StatusInternalServerError)
		case http.MethodGet:
			getCalls++
			http.Error(w, "not found", http.StatusNotFound)
		default:
			t.Errorf("request method = %s", request.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))
	defer server.Close()

	apiClient := client.NewClient(server.URL, "fixture-token", client.WithMaxRetries(0))
	implementation := &AlertPolicyResource{client: apiClient}
	ctx := context.Background()
	schemaResponse := &resource.SchemaResponse{}
	implementation.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	timeoutTypes := map[string]attr.Type{
		"create": types.StringType,
		"read":   types.StringType,
		"update": types.StringType,
		"delete": types.StringType,
	}
	model := AlertPolicyResourceModel{
		Name:        types.StringValue("fixture"),
		Namespace:   types.StringValue("system"),
		Annotations: types.MapNull(types.StringType),
		Description: types.StringNull(),
		Disable:     types.BoolNull(),
		Labels:      types.MapNull(types.StringType),
		ID:          types.StringValue("fixture"),
		Timeouts:    frameworktimeouts.Value{Object: types.ObjectNull(timeoutTypes)},
		Receivers:   types.ListNull(types.ObjectType{AttrTypes: AlertPolicyReceiversModelAttrTypes}),
		Routes:      types.ListNull(types.ObjectType{AttrTypes: AlertPolicyRoutesModelAttrTypes}),
	}
	state := tfsdk.State{Schema: schemaResponse.Schema}
	if diagnostics := state.Set(ctx, &model); diagnostics.HasError() {
		t.Fatalf("encode AlertPolicy state: %v", diagnostics)
	}

	response := &resource.DeleteResponse{}
	implementation.Delete(ctx, resource.DeleteRequest{State: state}, response)
	if response.Diagnostics.HasError() {
		t.Fatalf("verified idempotent delete returned diagnostics: %v", response.Diagnostics)
	}
	if deleteCalls != 1 || getCalls != 1 {
		t.Fatalf("requests = DELETE:%d GET:%d, want one each", deleteCalls, getCalls)
	}
}

func TestAlertPolicyDeletePreservesServerErrorWhenObjectExists(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodDelete {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"metadata":{"name":"fixture","namespace":"system"},"spec":{}}`))
	}))
	defer server.Close()

	apiClient := client.NewClient(server.URL, "fixture-token", client.WithMaxRetries(0))
	implementation := &AlertPolicyResource{client: apiClient}
	ctx := context.Background()
	schemaResponse := &resource.SchemaResponse{}
	implementation.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	timeoutTypes := map[string]attr.Type{
		"create": types.StringType,
		"read":   types.StringType,
		"update": types.StringType,
		"delete": types.StringType,
	}
	model := AlertPolicyResourceModel{
		Name:        types.StringValue("fixture"),
		Namespace:   types.StringValue("system"),
		Annotations: types.MapNull(types.StringType),
		Description: types.StringNull(),
		Disable:     types.BoolNull(),
		Labels:      types.MapNull(types.StringType),
		ID:          types.StringValue("fixture"),
		Timeouts:    frameworktimeouts.Value{Object: types.ObjectNull(timeoutTypes)},
		Receivers:   types.ListNull(types.ObjectType{AttrTypes: AlertPolicyReceiversModelAttrTypes}),
		Routes:      types.ListNull(types.ObjectType{AttrTypes: AlertPolicyRoutesModelAttrTypes}),
	}
	state := tfsdk.State{Schema: schemaResponse.Schema}
	if diagnostics := state.Set(ctx, &model); diagnostics.HasError() {
		t.Fatalf("encode AlertPolicy state: %v", diagnostics)
	}

	response := &resource.DeleteResponse{}
	implementation.Delete(ctx, resource.DeleteRequest{State: state}, response)
	if !response.Diagnostics.HasError() {
		t.Fatal("persistent DELETE 500 was treated as success while GET still found the object")
	}
}
