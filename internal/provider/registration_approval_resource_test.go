// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	if _, ok := resp.Schema.Attributes["cluster_size"]; !ok {
		t.Errorf("Expected 'cluster_size' attribute in schema")
	}
}

func TestRegistrationApprovalClusterSizeContract(t *testing.T) {
	validate := int64validator.OneOf(1, 3)
	for _, size := range []int64{1, 3} {
		var response validator.Int64Response
		validate.ValidateInt64(context.Background(), validator.Int64Request{ConfigValue: types.Int64Value(size)}, &response)
		if response.Diagnostics.HasError() {
			t.Fatalf("cluster_size %d was rejected: %v", size, response.Diagnostics)
		}
	}
	var response validator.Int64Response
	validate.ValidateInt64(context.Background(), validator.Int64Request{ConfigValue: types.Int64Value(2)}, &response)
	if !response.Diagnostics.HasError() {
		t.Fatal("cluster_size 2 was accepted")
	}
}

func TestRegistrationApprovalOverridesPendingClusterSize(t *testing.T) {
	postedClusterSize := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{
				"object": {
					"spec": {
						"gc_spec": {
							"passport": {"cluster_name":"test-site","cluster_size":1}
						}
					},
					"status": {
						"current_state": "PENDING"
					}
				}
			}`))
		case "POST":
			var body map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&body); err == nil {
				passport, _ := body["passport"].(map[string]interface{})
				postedClusterSize = passport["cluster_size"] == float64(3)
			}
			w.WriteHeader(http.StatusOK)
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
			"cluster_size":            tftypes.Number,
			"connected_region":        tftypes.String,
			"id":                      tftypes.String,
			"name":                    tftypes.String,
			"namespace":               tftypes.String,
			"preferred_active_re":     tftypes.String,
			"state":                   tftypes.String,
		},
	}, map[string]tftypes.Value{
		"backup_connected_region": tftypes.NewValue(tftypes.String, nil),
		"cluster_size":            tftypes.NewValue(tftypes.Number, big.NewFloat(3)),
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
		t.Fatalf("Create returned unexpected diagnostics: %v", resp.Diagnostics)
	}
	if !postedClusterSize {
		t.Fatal("Create did not override the server-derived passport cluster_size with the configured value")
	}
}

func TestRegistrationApprovalAlreadySatisfiedRequiresMatchingClusterSize(t *testing.T) {
	source := map[string]interface{}{
		"object": map[string]interface{}{
			"spec":   map[string]interface{}{"gc_spec": map[string]interface{}{"passport": map[string]interface{}{"cluster_size": float64(1)}}},
			"status": map[string]interface{}{"current_state": "ONLINE"},
		},
	}

	satisfied, err := registrationApprovalAlreadySatisfied(source, 1)
	if err != nil || !satisfied {
		t.Fatalf("matching progressed registration: satisfied=%v err=%v", satisfied, err)
	}
	if satisfied, err = registrationApprovalAlreadySatisfied(source, 3); err == nil || satisfied {
		t.Fatalf("mismatched progressed registration: satisfied=%v err=%v", satisfied, err)
	}
	source["object"].(map[string]interface{})["status"] = map[string]interface{}{"current_state": "PENDING"}
	if satisfied, err = registrationApprovalAlreadySatisfied(source, 3); err != nil || satisfied {
		t.Fatalf("pending registration should still be actionable: satisfied=%v err=%v", satisfied, err)
	}
}
