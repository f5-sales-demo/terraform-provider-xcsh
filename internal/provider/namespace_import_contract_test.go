// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestNamespaceExplicitImportContract(t *testing.T) {
	t.Parallel()
	for _, identifier := range []string{"example-namespace", "system/example-namespace"} {
		t.Run(identifier, func(t *testing.T) {
			var mutex sync.Mutex
			var addresses []string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
				mutex.Lock()
				addresses = append(addresses, request.URL.Path)
				mutex.Unlock()
				if request.Method != http.MethodGet {
					t.Errorf("adoption issued mutation: %s %s", request.Method, request.URL.Path)
					http.Error(w, "mutations forbidden", http.StatusMethodNotAllowed)
					return
				}
				if request.URL.Path != "/api/web/namespaces/example-namespace" {
					http.NotFound(w, request)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"metadata":{"name":"example-namespace","namespace":""},"spec":{},"resource_version":"fixture-version"}`))
			}))
			defer server.Close()
			implementation := &NamespaceResource{client: client.NewClient(server.URL, "fixture-token", client.WithMaxRetries(0))}
			ctx := context.Background()
			schema := &resource.SchemaResponse{}
			implementation.Schema(ctx, resource.SchemaRequest{}, schema)
			imported := &resource.ImportStateResponse{State: tfsdk.State{
				Schema: schema.Schema,
				Raw:    tftypes.NewValue(schema.Schema.Type().TerraformType(ctx), nil),
			}}
			// Private's type is framework-internal. Allocate its exported response field
			// without importing or modifying framework internals.
			private := reflect.ValueOf(imported).Elem().FieldByName("Private")
			private.Set(reflect.New(private.Type().Elem()))
			implementation.ImportState(ctx, resource.ImportStateRequest{ID: identifier}, imported)
			if imported.Diagnostics.HasError() {
				t.Fatalf("import: %v", imported.Diagnostics)
			}
			for key, want := range map[string]string{"name": identifier, "id": identifier, "namespace": ""} {
				var got string
				if diagnostics := imported.State.GetAttribute(ctx, path.Root(key), &got); diagnostics.HasError() || got != want {
					t.Fatalf("import %s = %q, want %q: %v", key, got, want, diagnostics)
				}
			}
			refreshed := &resource.ReadResponse{State: imported.State, Private: imported.Private}
			implementation.Read(ctx, resource.ReadRequest{State: imported.State, Private: imported.Private}, refreshed)
			if refreshed.Diagnostics.HasError() {
				t.Fatalf("read: %v", refreshed.Diagnostics)
			}
			if identifier == "example-namespace" {
				for key, want := range map[string]string{"name": identifier, "id": identifier, "namespace": ""} {
					var got string
					if diagnostics := refreshed.State.GetAttribute(ctx, path.Root(key), &got); diagnostics.HasError() || got != want {
						t.Fatalf("refreshed %s = %q, want %q: %v", key, got, want, diagnostics)
					}
				}
			} else if !refreshed.State.Raw.IsNull() {
				t.Fatal("prefixed import unexpectedly resolved the existing namespace")
			}
			mutex.Lock()
			defer mutex.Unlock()
			if len(addresses) != 1 || addresses[0] != "/api/web/namespaces/"+identifier {
				t.Fatalf("read addresses = %v", addresses)
			}
		})
	}
}

func TestNamespaceDataSourceNamespaceOptional(t *testing.T) {
	t.Parallel()
	response := &datasource.SchemaResponse{}
	(&NamespaceDataSource{}).Schema(context.Background(), datasource.SchemaRequest{}, response)
	attribute := response.Schema.Attributes["namespace"]
	if attribute.IsRequired() || !attribute.IsOptional() {
		t.Fatal("namespace data-source namespace must allow omission")
	}
}
