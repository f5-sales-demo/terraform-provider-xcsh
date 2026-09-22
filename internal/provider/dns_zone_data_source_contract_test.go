// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

func TestDNSZoneDataSourceSchemaExposesManagedRecordPrerequisite(t *testing.T) {
	ctx := context.Background()
	response := &datasource.SchemaResponse{}
	(&DNSZoneDataSource{}).Schema(ctx, datasource.SchemaRequest{}, response)
	if response.Diagnostics.HasError() {
		t.Fatalf("schema diagnostics: %v", response.Diagnostics)
	}

	namespace, ok := response.Schema.Attributes["namespace"].(datasourceschema.StringAttribute)
	if !ok || !namespace.Optional || !namespace.Computed || namespace.Required {
		t.Fatalf("namespace schema = %#v, want Optional+Computed with the resource's system default", response.Schema.Attributes["namespace"])
	}
	primary, ok := response.Schema.Attributes["primary"].(datasourceschema.SingleNestedAttribute)
	if !ok || !primary.Computed {
		t.Fatalf("primary schema = %#v, want computed single nested attribute", response.Schema.Attributes["primary"])
	}
	managed, ok := primary.Attributes["allow_http_lb_managed_records"].(datasourceschema.BoolAttribute)
	if !ok || !managed.Computed || managed.Optional || managed.Required {
		t.Fatalf("managed-record schema = %#v, want computed bool", primary.Attributes["allow_http_lb_managed_records"])
	}
	if secondary, ok := response.Schema.Attributes["secondary"].(datasourceschema.SingleNestedAttribute); !ok || !secondary.Computed {
		t.Fatalf("secondary schema = %#v, want computed single nested attribute", response.Schema.Attributes["secondary"])
	}
}

func TestDNSZoneDataSourceReadPreservesManagedRecordTriState(t *testing.T) {
	for _, test := range []struct {
		name     string
		primary  map[string]any
		wantNull bool
		want     bool
	}{
		{name: "true", primary: map[string]any{"allow_http_lb_managed_records": true}, want: true},
		{name: "false", primary: map[string]any{"allow_http_lb_managed_records": false}, want: false},
		{name: "absent", primary: map[string]any{}, wantNull: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			var requestPath string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
				requestPath = request.URL.Path
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(map[string]any{
					"metadata": map[string]any{"name": "example.com", "namespace": "system"},
					"spec":     map[string]any{"primary": test.primary},
				})
			}))
			defer server.Close()

			ctx := context.Background()
			dataSource := &DNSZoneDataSource{client: client.NewClient(server.URL, "test-token")}
			schemaResponse := &datasource.SchemaResponse{}
			dataSource.Schema(ctx, datasource.SchemaRequest{}, schemaResponse)
			config := DNSZoneDataSourceModel{
				ID: types.StringNull(), Name: types.StringValue("example.com"), Namespace: types.StringNull(),
				Description: types.StringNull(), Labels: types.MapNull(types.StringType), Annotations: types.MapNull(types.StringType),
			}
			response := datasource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
			dataSource.Read(ctx, datasource.ReadRequest{Config: tfsdk.Config{
				Schema: schemaResponse.Schema,
				Raw:    responseOperationRaw(t, config, schemaResponse.Schema.Type()),
			}}, &response)
			if response.Diagnostics.HasError() {
				t.Fatalf("read diagnostics: %v", response.Diagnostics)
			}
			var state DNSZoneDataSourceModel
			response.Diagnostics.Append(response.State.Get(ctx, &state)...)
			if response.Diagnostics.HasError() {
				t.Fatalf("state diagnostics: %v", response.Diagnostics)
			}
			if requestPath != "/api/config/dns/namespaces/system/dns_zones/example.com" {
				t.Fatalf("request path = %q, want system namespace default", requestPath)
			}
			if state.Primary == nil {
				t.Fatal("primary object is null")
			}
			got := state.Primary.AllowHTTPLBManagedRecords
			if got.IsNull() != test.wantNull {
				t.Fatalf("managed-record null = %v, want %v", got.IsNull(), test.wantNull)
			}
			if !test.wantNull && got.ValueBool() != test.want {
				t.Fatalf("managed-record value = %v, want %v", got.ValueBool(), test.want)
			}
		})
	}
}

func TestDNSZoneManagedRecordsPrecondition(t *testing.T) {
	for _, test := range []struct {
		name        string
		enabled     bool
		expectError *regexp.Regexp
	}{
		{name: "enabled", enabled: true},
		{name: "disabled", enabled: false, expectError: regexp.MustCompile("The selected DNS zone must enable HTTP LB managed records")},
	} {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(map[string]any{
					"metadata": map[string]any{"name": "example.com", "namespace": "system"},
					"spec": map[string]any{"primary": map[string]any{
						"allow_http_lb_managed_records": test.enabled,
					}},
				})
			}))
			defer server.Close()

			resource.UnitTest(t, resource.TestCase{
				ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
					"xcsh": providerserver.NewProtocol6WithError(New("test")()),
				},
				Steps: []resource.TestStep{{
					Config: fmt.Sprintf(`
provider "xcsh" {
  api_url   = %q
  api_token = "test-token"
}

data "xcsh_dns_zone" "selected" {
  name = "example.com"
}

resource "terraform_data" "guard" {
  lifecycle {
    precondition {
      condition = try(
        data.xcsh_dns_zone.selected.primary.allow_http_lb_managed_records,
        false
      )
      error_message = "The selected DNS zone must enable HTTP LB managed records."
    }
  }
}
`, server.URL),
					PlanOnly:           true,
					ExpectError:        test.expectError,
					ExpectNonEmptyPlan: test.expectError == nil,
				}},
			})
		})
	}
}
