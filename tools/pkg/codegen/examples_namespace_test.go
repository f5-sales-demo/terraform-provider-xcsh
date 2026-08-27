// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
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

func TestExampleNamespaceRetainsUnclassifiedFallback(t *testing.T) {
	if got := ExampleNamespace(nil, "unregistered_query_only_data_source"); got != "staging" {
		t.Fatalf("ExampleNamespace() = %q, want staging", got)
	}
}
