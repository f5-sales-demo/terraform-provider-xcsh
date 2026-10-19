package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func TestTokenCreateRetainsOriginalReceiptWhenReadFails(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if request.Method == http.MethodPost {
			_, _ = w.Write([]byte(`{"metadata":{"name":"fixture","namespace":"system"},"system_metadata":{"uid":"original-uid"}}`))
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()
	implementation := &TokenResource{client: client.NewClient(server.URL, "fixture-token", client.WithMaxRetries(0))}
	state := tokenDeleteState(t, implementation)
	response := &resource.CreateResponse{State: tfsdk.State{Schema: state.Schema}}
	ctx := context.Background()
	implementation.Create(ctx, resource.CreateRequest{Plan: tfsdk.Plan(state)}, response)
	if !response.Diagnostics.HasError() {
		t.Fatal("failed read was not reported")
	}
	var retained TokenResourceModel
	if diagnostics := response.State.Get(ctx, &retained); diagnostics.HasError() {
		t.Fatal(diagnostics)
	}
	if retained.ID.ValueString() != "fixture" || !retained.CreationIdentitySHA256.Equal(creationIdentityReceipt("original-uid")) {
		t.Fatal("successful POST receipt discarded after failed read")
	}
}

func TestTokenOriginalCreationIdentityGuardsDelete(t *testing.T) {
	for _, test := range []struct {
		name        string
		body        string
		status      int
		wantError   bool
		wantDeletes int
	}{
		{"original", `{"metadata":{"name":"fixture","namespace":"system"},"system_metadata":{"uid":"original-uid"}}`, 200, false, 1},
		{"replacement", `{"system_metadata":{"uid":"replacement-uid"}}`, 200, true, 0},
		{"missing_uid", `{}`, 200, true, 0},
		{"absent", `{}`, 404, false, 0},
		{"unavailable", `{}`, 503, true, 0},
	} {
		t.Run(test.name, func(t *testing.T) {
			deletes := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
				if request.Method == http.MethodDelete {
					deletes++
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(`{}`))
					return
				}
				if request.Method != http.MethodGet {
					t.Errorf("unexpected method %s", request.Method)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			}))
			defer server.Close()
			implementation := &TokenResource{client: client.NewClient(server.URL, "fixture-token", client.WithMaxRetries(0))}
			state := tokenDeleteState(t, implementation)
			ctx := context.Background()
			if diagnostics := state.SetAttribute(ctx, path.Root("creation_identity_sha256"), creationIdentityReceipt("original-uid")); diagnostics.HasError() {
				t.Fatal(diagnostics)
			}
			response := &resource.DeleteResponse{}
			implementation.Delete(ctx, resource.DeleteRequest{State: state}, response)
			if response.Diagnostics.HasError() != test.wantError || deletes != test.wantDeletes {
				t.Fatalf("error=%t deletes=%d, want error=%t deletes=%d: %v", response.Diagnostics.HasError(), deletes, test.wantError, test.wantDeletes, response.Diagnostics)
			}
		})
	}
}
