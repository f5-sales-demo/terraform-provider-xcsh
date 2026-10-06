package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	frameworktimeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func protectedDomainState(t *testing.T, r *ProtectedDomainResource) tfsdk.State {
	t.Helper()
	ctx := context.Background()
	schema := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schema)
	timeoutTypes := map[string]attr.Type{"create": types.StringType, "read": types.StringType, "update": types.StringType, "delete": types.StringType}
	data := ProtectedDomainResourceModel{
		ID: types.StringValue("owned"), Name: types.StringValue("owned"), Namespace: types.StringValue("demo"),
		ProtectedDomain: types.StringValue("example.com"), Description: types.StringValue("configured"),
		Labels:      types.MapValueMust(types.StringType, map[string]attr.Value{"owner": types.StringValue("demo")}),
		Annotations: types.MapNull(types.StringType), Timeouts: frameworktimeouts.Value{Object: types.ObjectNull(timeoutTypes)},
	}
	state := tfsdk.State{Schema: schema.Schema}
	if d := state.Set(ctx, &data); d.HasError() {
		t.Fatal(d)
	}
	return state
}
func allocatePrivate(response interface{}) {
	private := reflect.ValueOf(response).Elem().FieldByName("Private")
	private.Set(reflect.New(private.Type().Elem()))
}

func TestProtectedDomainResourceRefresh(t *testing.T) {
	for _, tc := range []struct {
		name, body        string
		status            int
		wantError, absent bool
	}{
		{"verified", "{\"items\":[{\"get_spec\":{\"protected_domain\":\"example.com\"}}]}", 501, false, false},
		{"absent", "{\"items\":[]}", 501, false, true},
		{"placeholder", "{\"items\":[{}]}", 501, true, false},
		{"deceptive", "404 501", 500, true, false},
		{"supported_absence", "", 404, false, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.Method != "GET" {
					t.Errorf("refresh mutation %s", req.Method)
				}
				if req.URL.RawQuery == "" {
					w.WriteHeader(tc.status)
					fmt.Fprint(w, "404 501")
					return
				}
				fmt.Fprint(w, tc.body)
			}))
			defer server.Close()
			r := &ProtectedDomainResource{client: client.NewClient(server.URL, "fixture-token", client.WithMaxRetries(0))}
			state := protectedDomainState(t, r)
			resp := &resource.ReadResponse{State: state}
			allocatePrivate(resp)
			r.Read(context.Background(), resource.ReadRequest{State: state, Private: resp.Private}, resp)
			if resp.Diagnostics.HasError() != tc.wantError {
				t.Fatalf("diagnostics: %v", resp.Diagnostics)
			}
			if resp.State.Raw.IsNull() != tc.absent {
				t.Fatalf("removed=%v", resp.State.Raw.IsNull())
			}
			if !tc.absent {
				for key, want := range map[string]string{"id": "owned", "name": "owned", "namespace": "demo", "protected_domain": "example.com", "description": "configured"} {
					var got string
					if d := resp.State.GetAttribute(context.Background(), path.Root(key), &got); d.HasError() || got != want {
						t.Fatalf("%s=%q: %v", key, got, d)
					}
				}
			}
		})
	}
}
func TestProtectedDomainAuthoritativeImport(t *testing.T) {
	for _, tc := range []struct {
		id, body  string
		wantError bool
	}{
		{"demo/owned/example.com", "{\"items\":[{\"get_spec\":{\"protected_domain\":\"example.com\"}}]}", false},
		{"demo/owned/other.example", "{\"items\":[{\"get_spec\":{\"protected_domain\":\"example.com\"}}]}", true},
		{"demo/owned/", "{\"items\":[]}", true},
		{"demo/owned/example.com", "{\"items\":[{}]}", true},
	} {
		t.Run(tc.id+tc.body, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.Method != "GET" {
					t.Errorf("import mutation")
				}
				if req.URL.RawQuery == "" {
					w.WriteHeader(501)
					return
				}
				fmt.Fprint(w, tc.body)
			}))
			defer server.Close()
			r := &ProtectedDomainResource{client: client.NewClient(server.URL, "fixture-token", client.WithMaxRetries(0))}
			schema := &resource.SchemaResponse{}
			r.Schema(context.Background(), resource.SchemaRequest{}, schema)
			resp := &resource.ImportStateResponse{State: tfsdk.State{Schema: schema.Schema, Raw: tftypes.NewValue(schema.Schema.Type().TerraformType(context.Background()), nil)}}
			allocatePrivate(resp)
			r.ImportState(context.Background(), resource.ImportStateRequest{ID: tc.id}, resp)
			if resp.Diagnostics.HasError() != tc.wantError {
				t.Fatalf("diagnostics=%v", resp.Diagnostics)
			}
		})
	}
}

func TestProtectedDomainCreateOwnership(t *testing.T) {
	for _, tc := range []struct {
		name       string
		post       int
		collection string
		wantError  bool
	}{
		{"verified", 200, "{\"items\":[{\"get_spec\":{\"protected_domain\":\"example.com\"}}]}", false},
		{"unavailable", 200, "{\"items\":[{}]}", true},
		{"conflict", 409, "", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			posts := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.Method == "POST" {
					posts++
					w.WriteHeader(tc.post)
					fmt.Fprint(w, "{}")
					return
				}
				if req.URL.RawQuery == "" {
					w.WriteHeader(501)
					return
				}
				fmt.Fprint(w, tc.collection)
			}))
			defer server.Close()
			r := &ProtectedDomainResource{client: client.NewClient(server.URL, "fixture-token", client.WithMaxRetries(0))}
			state := protectedDomainState(t, r)
			resp := &resource.CreateResponse{State: state}
			allocatePrivate(resp)
			r.Create(context.Background(), resource.CreateRequest{Plan: tfsdk.Plan(state)}, resp)
			if resp.Diagnostics.HasError() != tc.wantError {
				t.Fatalf("diagnostics=%v", resp.Diagnostics)
			}
			if posts != 1 {
				t.Fatalf("POST count=%d", posts)
			}
			if tc.post == 200 {
				var id string
				if d := resp.State.GetAttribute(context.Background(), path.Root("id"), &id); d.HasError() || id != "owned" {
					t.Fatalf("ownership lost: %s %v", id, d)
				}
			}
		})
	}
}

func TestProtectedDomainLookupRoot(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			t.Error("lookup mutation")
		}
		if req.URL.RawQuery == "" {
			w.WriteHeader(501)
			return
		}
		fmt.Fprint(w, "{\"items\":[{\"get_spec\":{\"protected_domain\":\"example.com\"}}]}")
	}))
	defer server.Close()
	ctx := context.Background()
	d := &ProtectedDomainDataSource{client: client.NewClient(server.URL, "fixture-token", client.WithMaxRetries(0))}
	schema := &datasource.SchemaResponse{}
	d.Schema(ctx, datasource.SchemaRequest{}, schema)
	model := ProtectedDomainDataSourceModel{ID: types.StringNull(), Name: types.StringValue("owned"), Namespace: types.StringValue("demo"), ProtectedDomain: types.StringValue("example.com"), Description: types.StringNull(), Labels: types.MapNull(types.StringType), Annotations: types.MapNull(types.StringType)}
	state := tfsdk.State{Schema: schema.Schema}
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatal(diags)
	}
	resp := &datasource.ReadResponse{State: state}
	d.Read(ctx, datasource.ReadRequest{Config: tfsdk.Config{Schema: schema.Schema, Raw: state.Raw}}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatal(resp.Diagnostics)
	}
	var root string
	if diags := resp.State.GetAttribute(ctx, path.Root("protected_domain"), &root); diags.HasError() || root != "example.com" {
		t.Fatal(root, diags)
	}
}

func TestProtectedDomainDeleteTypedStatus(t *testing.T) {
	for _, status := range []int{404, 501, 500} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.Method == "GET" {
					if req.URL.RawQuery == "" {
						w.WriteHeader(501)
						return
					}
					fmt.Fprint(w, "{\"items\":[]}")
					return
				}
				w.WriteHeader(status)
				fmt.Fprint(w, "message 404 501")
			}))
			defer server.Close()
			r := &ProtectedDomainResource{client: client.NewClient(server.URL, "fixture-token", client.WithMaxRetries(0))}
			resp := &resource.DeleteResponse{}
			r.Delete(context.Background(), resource.DeleteRequest{State: protectedDomainState(t, r)}, resp)
			if resp.Diagnostics.HasError() != (status != 404) {
				t.Fatal(resp.Diagnostics)
			}
		})
	}
}
