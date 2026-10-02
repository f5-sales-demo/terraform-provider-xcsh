package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestPublicIPBindingPreservesIdentityAndUnrelatedFields(t *testing.T) {
	object := map[string]interface{}{
		"metadata":         map[string]interface{}{"name": "ip-test", "namespace": "shared", "labels": map[string]interface{}{"keep": "yes"}},
		"resource_version": "7",
		"spec":             map[string]interface{}{"ip": "192.0.2.10", "ipv6": "", "virtual_sites": []interface{}{map[string]interface{}{"name": "all", "namespace": "shared", "tenant": "ves-io"}}, "other": "preserve"},
	}
	writes := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/config/namespaces/shared/public_ips/ip-test" {
			t.Errorf("wrong path %s", r.URL.Path)
		}
		if r.Method == http.MethodPut {
			writes++
			var body map[string]interface{}
			_ = json.NewDecoder(r.Body).Decode(&body)
			if body["resource_version"] != "7" {
				t.Errorf("missing current revision")
			}
			spec := body["spec"].(map[string]interface{})
			if spec["other"] != "preserve" || spec["ip"] != "192.0.2.10" {
				t.Errorf("unrelated identity changed")
			}
			object = body
		} else if r.Method != http.MethodGet {
			t.Errorf("unexpected method %s", r.Method)
		}
		_ = json.NewEncoder(w).Encode(object)
	}))
	defer server.Close()
	c := client.NewClient(server.URL, "test-token")
	ctx := context.Background()
	got, err := readPublicIPBinding(ctx, c, "shared", "ip-test", "192.0.2.10")
	if err != nil {
		t.Fatal(err)
	}
	original := got.Bindings
	desired := []interface{}{map[string]interface{}{"name": "canada", "namespace": "demo"}}
	if err = writePublicIPBinding(ctx, c, got, desired); err != nil {
		t.Fatal(err)
	}
	current, err := readPublicIPBinding(ctx, c, "shared", "ip-test", "192.0.2.10")
	if err != nil {
		t.Fatal(err)
	}
	if !samePublicIPBindings(current.Bindings, desired) {
		t.Fatal("new binding missing")
	}
	if err = writePublicIPBinding(ctx, c, current, desired); err != nil {
		t.Fatal(err)
	}
	if writes != 1 {
		t.Fatalf("repeat wrote %d times", writes)
	}
	if err = writePublicIPBinding(ctx, c, current, original); err != nil {
		t.Fatal(err)
	}
	if writes != 2 {
		t.Fatal("original binding not restored")
	}
	if _, err = readPublicIPBinding(ctx, c, "shared", "ip-test", "192.0.2.99"); err == nil {
		t.Fatal("wrong identity accepted")
	}
}

func TestPublicIPBindingReconcilesAmbiguousWrite(t *testing.T) {
	for _, applied := range []bool{true, false} {
		t.Run(map[bool]string{true: "applied", false: "rejected"}[applied], func(t *testing.T) {
			object := map[string]interface{}{"metadata": map[string]interface{}{"name": "ip-test", "namespace": "shared"}, "resource_version": "1", "spec": map[string]interface{}{"ip": "192.0.2.10", "virtual_sites": []interface{}{map[string]interface{}{"name": "all", "namespace": "shared"}}}}
			puts := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodPut {
					puts++
					var body map[string]interface{}
					_ = json.NewDecoder(r.Body).Decode(&body)
					if applied {
						object = body
					}
					w.WriteHeader(http.StatusBadGateway)
					return
				}
				_ = json.NewEncoder(w).Encode(object)
			}))
			defer server.Close()
			c := client.NewClient(server.URL, "test-token")
			current, err := readPublicIPBinding(context.Background(), c, "shared", "ip-test", "192.0.2.10")
			if err != nil {
				t.Fatal(err)
			}
			err = writePublicIPBinding(context.Background(), c, current, []interface{}{map[string]interface{}{"name": "canada", "namespace": "demo"}})
			if (err == nil) != applied {
				t.Fatalf("reconcile applied=%v err=%v", applied, err)
			}
			if puts != 1 {
				t.Fatalf("ambiguous PUT retried: %d", puts)
			}
		})
	}
}

func TestPublicIPBindingResourceLifecycleAndForeignDeleteGuard(t *testing.T) {
	object := map[string]interface{}{
		"metadata":         map[string]interface{}{"name": "ip-test", "namespace": "shared"},
		"resource_version": "7",
		"spec":             map[string]interface{}{"ip": "192.0.2.10", "virtual_sites": []interface{}{map[string]interface{}{"name": "all", "namespace": "shared", "tenant": "ves-io"}}},
	}
	original := normalizedPublicIPBindings(object["spec"].(map[string]interface{})["virtual_sites"].([]interface{}))
	puts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/config/namespaces/demo/virtual_sites/canada" {
			_ = json.NewEncoder(w).Encode(map[string]interface{}{"metadata": map[string]interface{}{"name": "canada", "namespace": "demo"}, "spec": map[string]interface{}{"site_type": "REGIONAL_EDGE"}})
			return
		}
		if r.Method == http.MethodPut {
			puts++
			_ = json.NewDecoder(r.Body).Decode(&object)
		}
		_ = json.NewEncoder(w).Encode(object)
	}))
	defer server.Close()
	resourceUnderTest := &PublicIPBindingResource{client: client.NewClient(server.URL, "test-token")}
	ctx := context.Background()
	schemaResponse := resource.SchemaResponse{}
	resourceUnderTest.Schema(ctx, resource.SchemaRequest{}, &schemaResponse)
	model := publicIPBindingModel{Name: types.StringValue("ip-test"), Namespace: types.StringValue("shared"), ExpectedIP: types.StringValue("192.0.2.10"), VirtualSite: types.StringValue("canada"), VirtualSiteNamespace: types.StringValue("demo"), ID: types.StringNull(), OriginalBindings: types.StringNull(), ManagedBindings: types.StringNull()}
	create := resource.CreateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	resourceUnderTest.Create(ctx, resource.CreateRequest{Plan: tfsdk.Plan{Schema: schemaResponse.Schema, Raw: responseOperationRaw(t, model, schemaResponse.Schema.Type())}}, &create)
	if create.Diagnostics.HasError() {
		t.Fatalf("create: %v", create.Diagnostics)
	}
	var saved publicIPBindingModel
	create.State.Get(ctx, &saved)
	if saved.ID.ValueString() != "shared/ip-test" || puts != 1 {
		t.Fatal("adoption state missing")
	}
	// Refresh a foreign change, then verify deletion still compares against last
	// managed state instead of treating the refreshed foreign value as owned.
	object["spec"].(map[string]interface{})["virtual_sites"] = []interface{}{map[string]interface{}{"name": "foreign", "namespace": "demo"}}
	read := resource.ReadResponse{State: create.State}
	resourceUnderTest.Read(ctx, resource.ReadRequest{State: create.State}, &read)
	if read.Diagnostics.HasError() {
		t.Fatalf("read: %v", read.Diagnostics)
	}
	var drifted publicIPBindingModel
	read.State.Get(ctx, &drifted)
	if drifted.VirtualSite.ValueString() != "foreign" {
		t.Fatal("drift not observed")
	}
	deleted := resource.DeleteResponse{State: read.State}
	resourceUnderTest.Delete(ctx, resource.DeleteRequest{State: read.State}, &deleted)
	if !deleted.Diagnostics.HasError() || puts != 1 {
		t.Fatal("delete overwrote foreign binding")
	}
	// The ordinary update repairs the drift using the configured Canadian target.
	update := resource.UpdateResponse{State: read.State}
	resourceUnderTest.Update(ctx, resource.UpdateRequest{State: read.State, Plan: tfsdk.Plan{Schema: schemaResponse.Schema, Raw: responseOperationRaw(t, saved, schemaResponse.Schema.Type())}}, &update)
	if update.Diagnostics.HasError() {
		t.Fatalf("update: %v", update.Diagnostics)
	}
	deleted = resource.DeleteResponse{State: update.State}
	resourceUnderTest.Delete(ctx, resource.DeleteRequest{State: update.State}, &deleted)
	if deleted.Diagnostics.HasError() {
		t.Fatalf("delete: %v", deleted.Diagnostics)
	}
	got := normalizedPublicIPBindings(object["spec"].(map[string]interface{})["virtual_sites"].([]interface{}))
	if got != original || puts != 3 {
		t.Fatal("original binding not restored")
	}
}

func TestPublicIPBindingSchemaKeepsAllocatedIdentityImmutable(t *testing.T) {
	r := &PublicIPBindingResource{}
	s := resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, &s)
	for _, name := range []string{"name", "namespace", "expected_ip"} {
		field, ok := s.Schema.Attributes[name].(schema.StringAttribute)
		if !ok || !field.Required || len(field.PlanModifiers) != 1 {
			t.Fatalf("%s identity is mutable", name)
		}
	}
	c := resource.ConfigureResponse{}
	r.Configure(context.Background(), resource.ConfigureRequest{ProviderData: "wrong"}, &c)
	if !c.Diagnostics.HasError() {
		t.Fatal("invalid client accepted")
	}
}

func TestPublicIPBindingReadRejectsMalformedIdentityAndMissingRevision(t *testing.T) {
	for _, fixture := range []struct {
		name     string
		metadata map[string]interface{}
		revision string
		expected string
	}{
		{"wrong-name", map[string]interface{}{"name": "foreign", "namespace": "shared"}, "1", "192.0.2.10"},
		{"invalid-address", map[string]interface{}{"name": "ip-test", "namespace": "shared"}, "1", "not-an-ip"},
		{"missing-revision", map[string]interface{}{"name": "ip-test", "namespace": "shared"}, "", "192.0.2.10"},
	} {
		t.Run(fixture.name, func(t *testing.T) {
			object := map[string]interface{}{"metadata": fixture.metadata, "resource_version": fixture.revision, "spec": map[string]interface{}{"ip": "192.0.2.10", "virtual_sites": []interface{}{map[string]interface{}{"name": "all", "namespace": "shared"}}}}
			puts := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.Method == http.MethodPut {
					puts++
				}
				_ = json.NewEncoder(w).Encode(object)
			}))
			defer server.Close()
			c := client.NewClient(server.URL, "test-token")
			current, err := readPublicIPBinding(context.Background(), c, "shared", "ip-test", fixture.expected)
			if fixture.name == "missing-revision" {
				if err != nil {
					t.Fatal(err)
				}
				err = writePublicIPBinding(context.Background(), c, current, []interface{}{map[string]interface{}{"name": "canada", "namespace": "demo"}})
			}
			if err == nil || puts != 0 {
				t.Fatal("unsafe identity or revision was accepted")
			}
		})
	}
}
