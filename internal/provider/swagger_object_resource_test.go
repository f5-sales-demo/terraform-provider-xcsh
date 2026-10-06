package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const swaggerFixture = `{"openapi":"3.0.3","info":{"title":"Synthetic fixture","version":"1"},"paths":{}}`

func swaggerState(t *testing.T, r *SwaggerObjectResource, version, content string) tfsdk.State {
	t.Helper()
	ctx := context.Background()
	sr := resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, &sr)
	d := SwaggerObjectResourceModel{Name: types.StringValue("fixture"), Namespace: types.StringValue("demo"), Content: types.StringValue(content), Version: types.StringUnknown(), ID: types.StringUnknown(), Path: types.StringUnknown(), SHA256: types.StringUnknown()}
	if version != "" {
		swaggerModel(&d, version, content)
	}
	s := tfsdk.State{Schema: sr.Schema}
	if diags := s.Set(ctx, &d); diags.HasError() {
		t.Fatal(diags)
	}
	return s
}
func TestSwaggerLifecycleExactVersion(t *testing.T) {
	calls := []string{}
	exists := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		calls = append(calls, req.Method+" "+req.URL.Path)
		switch {
		case req.Method == "GET" && req.URL.Path == "/api/object_store/namespaces/demo/stored_objects/swagger":
			if req.URL.Query().Get("latest_version_only") != "false" || req.URL.Query().Get("query_type") != "EXACT_MATCH" {
				t.Error("unbounded listing")
			}
			fmt.Fprint(w, `{"items":[]}`)
		case req.Method == "PUT":
			if exists {
				t.Error("repeated version issuance")
			}
			exists = true
			var body map[string]string
			_ = json.NewDecoder(req.Body).Decode(&body)
			if body["string_value"] != swaggerFixture {
				t.Error("content bytes changed")
			}
			fmt.Fprint(w, `{"metadata":{"namespace":"demo","name":"fixture","version":"v1"}}`)
		case req.Method == "GET" && strings.HasSuffix(req.URL.Path, "/fixture/v1"):
			if !exists {
				http.NotFound(w, req)
				return
			}
			_ = json.NewEncoder(w).Encode(map[string]any{"metadata": map[string]string{"namespace": "demo", "name": "fixture", "version": "v1"}, "string_value": swaggerFixture})
		case req.Method == "DELETE" && strings.HasSuffix(req.URL.Path, "/fixture/v1"):
			exists = false
			fmt.Fprint(w, "{}")
		default:
			t.Errorf("unexpected route %s", req.URL)
			http.Error(w, "wrong route", 400)
		}
	}))
	defer server.Close()
	r := &SwaggerObjectResource{client: client.NewClient(server.URL, "fixture", client.WithMaxRetries(0))}
	state := swaggerState(t, r, "", swaggerFixture)
	create := resource.CreateResponse{State: state}
	r.Create(context.Background(), resource.CreateRequest{Plan: tfsdk.Plan(state)}, &create)
	if create.Diagnostics.HasError() {
		t.Fatal(create.Diagnostics)
	}
	var d SwaggerObjectResourceModel
	_ = create.State.Get(context.Background(), &d)
	if d.Version.ValueString() != "v1" || d.ID.ValueString() != "demo/fixture/v1" {
		t.Fatal(d)
	}
	read := resource.ReadResponse{State: create.State}
	r.Read(context.Background(), resource.ReadRequest{State: create.State}, &read)
	if read.Diagnostics.HasError() || !read.State.Raw.Equal(create.State.Raw) {
		t.Fatal("read did not converge", read.Diagnostics)
	}
	imported := resource.ImportStateResponse{State: tfsdk.State{Schema: state.Schema, Raw: state.Raw}}
	r.ImportState(context.Background(), resource.ImportStateRequest{ID: "demo/fixture/v1"}, &imported)
	// Import starts without an owned digest. Read materializes the complete verified document.
	_ = imported.State.SetAttribute(context.Background(), path.Root("sha256"), types.StringNull())
	rr := resource.ReadResponse{State: imported.State}
	r.Read(context.Background(), resource.ReadRequest{State: imported.State}, &rr)
	if rr.Diagnostics.HasError() || !rr.State.Raw.Equal(create.State.Raw) {
		t.Fatal("import round trip differs", rr.Diagnostics)
	}
	del := resource.DeleteResponse{}
	r.Delete(context.Background(), resource.DeleteRequest{State: create.State}, &del)
	if del.Diagnostics.HasError() {
		t.Fatal(del.Diagnostics)
	}
	for _, c := range calls {
		if strings.Contains(c, "latest") {
			t.Fatal(c)
		}
	}
	absent := resource.ReadResponse{State: create.State}
	r.Read(context.Background(), resource.ReadRequest{State: create.State}, &absent)
	if absent.Diagnostics.HasError() || !absent.State.Raw.IsNull() {
		t.Fatal("missing owned version not removed", absent.Diagnostics)
	}
}
func TestSwaggerImportRejectsUnpinnedIdentities(t *testing.T) {
	r := &SwaggerObjectResource{}
	for _, id := range []string{"demo/fixture/latest", "demo/fixture/LATEST", "demo/fixture", "demo/fixture/v1/other", "../fixture/v1", "demo/fixture/"} {
		s := swaggerState(t, r, "", swaggerFixture)
		resp := resource.ImportStateResponse{State: s}
		r.ImportState(context.Background(), resource.ImportStateRequest{ID: id}, &resp)
		if !resp.Diagnostics.HasError() {
			t.Errorf("accepted %q", id)
		}
	}
}
func TestSwaggerDriftRejectsDeletion(t *testing.T) {
	deletes := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "DELETE" {
			deletes++
			fmt.Fprint(w, "{}")
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"metadata": map[string]string{"namespace": "demo", "name": "fixture", "version": "v1"}, "string_value": swaggerFixture + " "})
	}))
	defer server.Close()
	r := &SwaggerObjectResource{client: client.NewClient(server.URL, "fixture", client.WithMaxRetries(0))}
	s := swaggerState(t, r, "v1", swaggerFixture)
	read := resource.ReadResponse{State: s}
	r.Read(context.Background(), resource.ReadRequest{State: s}, &read)
	if !read.Diagnostics.HasError() {
		t.Fatal("immutable drift accepted")
	}
	del := resource.DeleteResponse{}
	r.Delete(context.Background(), resource.DeleteRequest{State: s}, &del)
	if !del.Diagnostics.HasError() || deletes != 0 {
		t.Fatal("foreign content deleted")
	}
}
func TestSwaggerFailedUploadNotReplayed(t *testing.T) {
	puts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			fmt.Fprint(w, `{"items":[]}`)
			return
		}
		puts++
		http.Error(w, "upstream", http.StatusServiceUnavailable)
	}))
	defer server.Close()
	r := &SwaggerObjectResource{client: client.NewClient(server.URL, "fixture", client.WithMaxRetries(3))}
	s := swaggerState(t, r, "", swaggerFixture)
	resp := resource.CreateResponse{State: s}
	r.Create(context.Background(), resource.CreateRequest{Plan: tfsdk.Plan(s)}, &resp)
	if !resp.Diagnostics.HasError() || puts != 1 {
		t.Fatalf("ambiguous PUT calls=%d diagnostics=%v", puts, resp.Diagnostics)
	}
}
