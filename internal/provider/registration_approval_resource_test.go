// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestRegistrationApprovalResourceSchema(t *testing.T) {
	t.Parallel()
	r := NewRegistrationApprovalResource()
	req := resource.SchemaRequest{}
	resp := &resource.SchemaResponse{}
	r.Schema(context.Background(), req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("Schema returned unexpected errors: %v", resp.Diagnostics)
	}

	if _, ok := resp.Schema.Attributes["name"]; !ok {
		t.Errorf("Expected 'name' attribute in schema")
	}
	if _, ok := resp.Schema.Attributes["namespace"]; !ok {
		t.Errorf("Expected 'namespace' attribute in schema")
	}
	if _, ok := resp.Schema.Attributes["state"]; !ok {
		t.Errorf("Expected 'state' attribute in schema")
	}
}

func TestRegistrationApprovalIdempotency(t *testing.T) {
	// Issue #1278: Ensure registration approval is idempotent
	// Server responds with 400 'not in NEW state' on Post approve, but GET returns 'APPROVED' state.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{
				"object": {
					"spec": {
						"gc_spec": {
							"passport": "dummy-passport"
						}
					},
					"status": {
						"state": "APPROVED"
					}
				}
			}`))
		case "POST":
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"code": "BAD_REQUEST", "message": "Registration is not in NEW state"}`))
		}
	}))
	defer server.Close()

	c := client.NewClient(server.URL, "test-token")
	r := &RegistrationApprovalResource{client: c}

	schemaReq := resource.SchemaRequest{}
	schemaResp := &resource.SchemaResponse{}
	r.Schema(context.Background(), schemaReq, schemaResp)

	planVal := tftypes.NewValue(tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"backup_connected_region": tftypes.String,
			"connected_region":        tftypes.String,
			"id":                      tftypes.String,
			"name":                    tftypes.String,
			"namespace":               tftypes.String,
			"preferred_active_re":     tftypes.String,
			"state":                   tftypes.String,
		},
	}, map[string]tftypes.Value{
		"backup_connected_region": tftypes.NewValue(tftypes.String, nil),
		"connected_region":        tftypes.NewValue(tftypes.String, nil),
		"id":                      tftypes.NewValue(tftypes.String, nil),
		"name":                    tftypes.NewValue(tftypes.String, "test-site"),
		"namespace":               tftypes.NewValue(tftypes.String, "system"),
		"preferred_active_re":     tftypes.NewValue(tftypes.String, nil),
		"state":                   tftypes.NewValue(tftypes.String, nil),
	})

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Schema: schemaResp.Schema,
			Raw:    planVal,
		},
	}
	resp := resource.CreateResponse{
		State: tfsdk.State{
			Schema: schemaResp.Schema,
		},
	}

	r.Create(context.Background(), req, &resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("Create returned unexpected diagnostics for idempotent 400: %v", resp.Diagnostics)
	}
}
