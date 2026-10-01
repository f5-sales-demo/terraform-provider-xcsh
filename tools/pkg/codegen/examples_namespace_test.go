// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestExampleNamespaceUsesSingleSchemaConstraint(t *testing.T) {
	rt := &openapi.ResourceTemplate{Attributes: []openapi.TerraformAttribute{
		{TfsdkTag: "namespace", EnumValues: []string{"system"}},
	}}
	if got := ExampleNamespace(rt, "securemesh_site_v2"); got != "system" {
		t.Fatalf("ExampleNamespace() = %q, want system", got)
	}
}

func TestDNSZoneExamplesEnforceManagedRecordPrerequisite(t *testing.T) {
	resourceExample := RenderResourceExampleHCL(&openapi.ResourceTemplate{Description: "DNS zone."}, "dns_zone", "system")
	for _, want := range []string{"primary {", "allow_http_lb_managed_records = true"} {
		if !strings.Contains(resourceExample, want) {
			t.Errorf("stack-owned DNS-zone example missing %q:\n%s", want, resourceExample)
		}
	}

	dataSourceExample := RenderDataSourceExampleHCL("dns_zone", "system")
	for _, want := range []string{
		`resource "terraform_data" "require_managed_records"`,
		`data.xcsh_dns_zone.example.primary.allow_http_lb_managed_records`,
		`error_message = "The selected DNS zone must enable HTTP LB managed records."`,
	} {
		if !strings.Contains(dataSourceExample, want) {
			t.Errorf("external DNS-zone example missing %q:\n%s", want, dataSourceExample)
		}
	}
}

func TestExampleNamespaceRetainsUnclassifiedFallback(t *testing.T) {
	if got := ExampleNamespace(nil, "unregistered_query_only_data_source"); got != "staging" {
		t.Fatalf("ExampleNamespace() = %q, want staging", got)
	}
}

func TestNamespaceExamplesOmitNamespaceAndConfigureProvider(t *testing.T) {
	for _, example := range []string{
		RenderResourceExampleHCL(&openapi.ResourceTemplate{}, "namespace", "staging"),
		RenderDataSourceExampleHCL("namespace", "staging"),
	} {
		if strings.Contains(example, "namespace =") {
			t.Errorf("tenant-level example sets namespace: %s", example)
		}
		if !strings.Contains(example, `provider "xcsh" {}`) {
			t.Errorf("standalone example lacks provider: %s", example)
		}
	}
}

func TestNamespaceDataSourceSchemaAllowsOmittedNamespace(t *testing.T) {
	dir := t.TempDir()
	if err := GenerateDataSource(&openapi.ResourceTemplate{Name: "namespace", TitleCase: "Namespace"}, dir); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "namespace_data_source.go"))
	if err != nil {
		t.Fatal(err)
	}
	source := string(data)
	start := strings.Index(source, `"namespace": schema.StringAttribute{`)
	end := strings.Index(source[start:], "},")
	block := source[start : start+end]
	if strings.Contains(block, "Required:") || !strings.Contains(block, "Optional:") {
		t.Fatalf("namespace must be optional: %s", block)
	}
}
