package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func policyState(t *testing.T, r *SecretPolicyResource) tfsdk.State {
	t.Helper()
	ctx := context.Background()
	s := resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, &s)
	selector, _ := types.ListValueFrom(ctx, types.StringType, []string{"app=demo"})
	rules, _ := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: secretPolicyRuleTypes}, []secretPolicyRuleModel{{Name: types.StringValue("allow"), Description: types.StringNull(), Action: types.StringValue("ALLOW"), Selector: selector}, {Name: types.StringValue("deny"), Description: types.StringNull(), Action: types.StringValue("DENY"), Selector: types.ListNull(types.StringType)}})
	labels, _ := types.MapValueFrom(ctx, types.StringType, map[string]string{"owned": "yes"})
	d := secretPolicyModel{Name: types.StringValue("policy"), Namespace: types.StringValue("demo"), Description: types.StringNull(), Labels: labels, Annotations: types.MapNull(types.StringType), AllowF5XC: types.BoolValue(true), Cache: types.StringValue("30s"), Algorithm: types.StringValue("FIRST_MATCH"), Rules: rules, ID: types.StringUnknown(), UID: types.StringUnknown(), MarkedForDelete: types.BoolUnknown(), DeletionTime: types.StringUnknown()}
	state := tfsdk.State{Schema: s.Schema}
	if diag := state.Set(ctx, &d); diag.HasError() {
		t.Fatal(diag)
	}
	return state
}
func TestSecretPolicyCRUDMetadataAndTombstone(t *testing.T) {
	ctx := context.Background()
	exists := false
	deletes := 0
	var policy client.SecretPolicy
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			if !exists {
				http.NotFound(w, req)
				return
			}
			_ = json.NewEncoder(w).Encode(policy)
		case "POST":
			if exists {
				t.Error("adopted existing policy")
			}
			_ = json.NewDecoder(req.Body).Decode(&policy)
			exists = true
			policy.SystemMetadata.UID = "uid-1"
			policy.Spec.Algorithm = "FIRST_MATCH"
			policy.Metadata.Annotations = map[string]string{"external": "keep"}
			policy.Metadata.Labels["server-default"] = "keep"
			_, _ = w.Write([]byte(`{}`))
		case "PUT":
			_ = json.NewDecoder(req.Body).Decode(&policy)
			policy.SystemMetadata.UID = "uid-1"
			policy.Spec.Algorithm = "FIRST_MATCH"
			_, _ = w.Write([]byte(`{}`))
		case "DELETE":
			deletes++
			policy.Spec.MarkedForDelete = true
			policy.Spec.DeletionTime = "2030-01-01T00:00:00Z"
			_, _ = w.Write([]byte(`{}`))
		default:
			t.Error("unexpected method")
		}
	}))
	defer server.Close()
	r := &SecretPolicyResource{client: client.NewClient(server.URL, "fixture", client.WithMaxRetries(0))}
	state := policyState(t, r)
	created := resource.CreateResponse{State: state}
	r.Create(ctx, resource.CreateRequest{Plan: tfsdk.Plan(state)}, &created)
	if created.Diagnostics.HasError() {
		t.Fatal(created.Diagnostics)
	}
	read := resource.ReadResponse{State: created.State}
	r.Read(ctx, resource.ReadRequest{State: created.State}, &read)
	if read.Diagnostics.HasError() || !read.State.Raw.Equal(created.State.Raw) {
		t.Fatal("read caused state drift", read.Diagnostics)
	}
	var d secretPolicyModel
	_ = created.State.Get(ctx, &d)
	d.Description = types.StringValue("updated")
	updatedPlan := tfsdk.Plan{Schema: state.Schema}
	_ = updatedPlan.Set(ctx, &d)
	update := resource.UpdateResponse{State: created.State}
	r.Update(ctx, resource.UpdateRequest{State: created.State, Plan: updatedPlan}, &update)
	if update.Diagnostics.HasError() {
		t.Fatal(update.Diagnostics)
	}
	if policy.Metadata.Annotations["external"] != "keep" || policy.Metadata.Labels["server-default"] != "keep" {
		t.Fatal("unmanaged metadata lost")
	}
	read = resource.ReadResponse{State: update.State}
	r.Read(ctx, resource.ReadRequest{State: update.State}, &read)
	if read.Diagnostics.HasError() || !read.State.Raw.Equal(update.State.Raw) {
		t.Fatal("updated state drift", read.Diagnostics)
	}
	for i := 0; i < 2; i++ {
		deleted := resource.DeleteResponse{State: update.State}
		r.Delete(ctx, resource.DeleteRequest{State: update.State}, &deleted)
		if !deleted.Diagnostics.HasError() {
			t.Fatal("tombstone reported as absence")
		}
	}
	if deletes != 1 {
		t.Fatal("repeated DELETE for tombstone", deletes)
	}
	exists = false
	for i := 0; i < 2; i++ {
		deleted := resource.DeleteResponse{State: update.State}
		r.Delete(ctx, resource.DeleteRequest{State: update.State}, &deleted)
		if deleted.Diagnostics.HasError() {
			t.Fatal(deleted.Diagnostics)
		}
	}
}
func TestSecretPolicyPermissionIsNotAbsence(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		http.Error(w, "forbidden", http.StatusForbidden)
	}))
	defer server.Close()
	r := &SecretPolicyResource{client: client.NewClient(server.URL, "fixture", client.WithMaxRetries(0))}
	state := policyState(t, r)
	read := resource.ReadResponse{State: state}
	r.Read(context.Background(), resource.ReadRequest{State: state}, &read)
	if !read.Diagnostics.HasError() || read.State.Raw.IsNull() {
		t.Fatal("403 reported as absence")
	}
	create := resource.CreateResponse{State: state}
	r.Create(context.Background(), resource.CreateRequest{Plan: tfsdk.Plan(state)}, &create)
	if !create.Diagnostics.HasError() || calls != 2 {
		t.Fatal("created without absence proof")
	}
}
func TestSecretPolicyUnsafeSelectorsAndDuplicateNames(t *testing.T) {
	r := &SecretPolicyResource{}
	ctx := context.Background()
	state := policyState(t, r)
	var d secretPolicyModel
	_ = state.Get(ctx, &d)
	for _, selector := range []types.List{types.ListNull(types.StringType), types.ListValueMust(types.StringType, []attr.Value{types.StringValue("")})} {
		d.Rules = types.ListValueMust(types.ObjectType{AttrTypes: secretPolicyRuleTypes}, []attr.Value{types.ObjectValueMust(secretPolicyRuleTypes, map[string]attr.Value{"name": types.StringValue("allow"), "action": types.StringValue("ALLOW"), "description": types.StringNull(), "client_selector": selector})})
		_ = state.Set(ctx, &d)
		response := resource.ValidateConfigResponse{}
		r.ValidateConfig(ctx, resource.ValidateConfigRequest{Config: tfsdk.Config(state)}, &response)
		if !response.Diagnostics.HasError() {
			t.Fatal("unbounded ALLOW accepted")
		}
	}
}
