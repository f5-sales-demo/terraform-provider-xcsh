// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	frameworktimeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

func tokenDeleteState(t *testing.T, implementation *TokenResource) tfsdk.State {
	t.Helper()
	ctx := context.Background()
	schemaResponse := &resource.SchemaResponse{}
	implementation.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	timeoutTypes := map[string]attr.Type{
		"create": types.StringType,
		"read":   types.StringType,
		"update": types.StringType,
		"delete": types.StringType,
	}
	model := TokenResourceModel{
		Name:        types.StringValue("fixture"),
		Namespace:   types.StringValue("system"),
		Annotations: types.MapNull(types.StringType),
		Description: types.StringNull(),
		Disable:     types.BoolNull(),
		Labels:      types.MapNull(types.StringType),
		ID:          types.StringValue("fixture"),
		Uid:         types.StringNull(),
		Content:     types.StringNull(),
		SiteName:    types.StringNull(),
		Type:        types.Int64Null(),
		Timeouts:    frameworktimeouts.Value{Object: types.ObjectNull(timeoutTypes)},
	}
	state := tfsdk.State{Schema: schemaResponse.Schema}
	if diagnostics := state.Set(ctx, &model); diagnostics.HasError() {
		t.Fatalf("encode Token state: %v", diagnostics)
	}
	return state
}

func TestTokenDeleteOnlyTreatsNotFoundAsIdempotentSuccess(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
		name       string
		statusCode int
		wantError  bool
	}{
		{name: "not_found", statusCode: http.StatusNotFound, wantError: false},
		{name: "not_implemented", statusCode: http.StatusNotImplemented, wantError: true},
		{name: "server_error", statusCode: http.StatusInternalServerError, wantError: true},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			requests := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
				requests++
				if request.Method != http.MethodDelete {
					t.Errorf("request method = %s, want DELETE", request.Method)
				}
				if request.URL.Path != "/api/register/namespaces/system/tokens/fixture" {
					t.Errorf("request path = %s", request.URL.Path)
				}
				http.Error(w, http.StatusText(testCase.statusCode), testCase.statusCode)
			}))
			defer server.Close()

			apiClient := client.NewClient(server.URL, "fixture-token", client.WithMaxRetries(0))
			implementation := &TokenResource{client: apiClient}
			response := &resource.DeleteResponse{}
			implementation.Delete(context.Background(), resource.DeleteRequest{State: tokenDeleteState(t, implementation)}, response)

			if response.Diagnostics.HasError() != testCase.wantError {
				t.Fatalf("DELETE %d diagnostics error = %t, want %t: %v", testCase.statusCode, response.Diagnostics.HasError(), testCase.wantError, response.Diagnostics)
			}
			if requests != 1 {
				t.Fatalf("DELETE %d made %d requests, want 1", testCase.statusCode, requests)
			}
		})
	}
}

func TestTokenDeleteTransientBadRequestHonorsContextDeadline(t *testing.T) {
	t.Parallel()

	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		http.Error(w, "bad request", http.StatusBadRequest)
	}))
	defer server.Close()

	apiClient := client.NewClient(server.URL, "fixture-token", client.WithMaxRetries(0))
	implementation := &TokenResource{client: apiClient}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	started := time.Now()
	response := &resource.DeleteResponse{}
	implementation.Delete(ctx, resource.DeleteRequest{State: tokenDeleteState(t, implementation)}, response)

	if !response.Diagnostics.HasError() {
		t.Fatal("transient DELETE 400 timeout was treated as success")
	}
	if requests != 1 {
		t.Fatalf("DELETE 400 made %d requests before context cancellation, want 1", requests)
	}
	if elapsed := time.Since(started); elapsed > time.Second {
		t.Fatalf("DELETE 400 did not honor context deadline: %s", elapsed)
	}
}
