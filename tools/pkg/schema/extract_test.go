// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package schema

import (
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/namespace"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

// findAttr returns the attribute with the given tfsdk tag, or nil.
func findAttr(attrs []openapi.TerraformAttribute, tag string) *openapi.TerraformAttribute {
	for i := range attrs {
		if attrs[i].TfsdkTag == tag {
			return &attrs[i]
		}
	}
	return nil
}

// systemOnlySpec builds a minimal spec + namespaced path for a resource.
func systemOnlySpec(resourceName string) (*openapi.Spec, func(*openapi.Spec, string) (string, string, bool)) {
	spec := &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				resourceName + "CreateSpecType": {
					Type:       "object",
					Properties: map[string]openapi.Schema{"port": {Type: "integer"}},
				},
			},
		},
	}
	extractAPIPath := func(_ *openapi.Spec, _ string) (string, string, bool) {
		return "/api/config/dns/namespaces/%s/" + resourceName + "s", "/api/config/dns/namespaces/%s/" + resourceName + "s/%s", true
	}
	return spec, extractAPIPath
}

// A resource whose spec profile restricts it to a single namespace ("system")
// must emit namespace as Optional+Computed with a static default and a OneOf
// validator — not Required — so it can be omitted and cannot be set wrong.
func TestExtractResourceSchema_SingleAllowedNamespaceIsDefaulted(t *testing.T) {
	namespace.ClearProfiles()
	namespace.SetProfile("sys_res", namespace.Profile{Allowed: []namespace.NamespaceType{namespace.System}, Enforced: true})
	defer namespace.ClearProfiles()

	spec, extractAPIPath := systemOnlySpec("sys_res")
	result, err := ExtractResourceSchema(spec, "sys_res", extractAPIPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ns := findAttr(result.Attributes, "namespace")
	if ns == nil {
		t.Fatal("namespace attribute missing")
	}
	if ns.Required {
		t.Error("namespace must not be Required for a single-allowed-namespace resource")
	}
	if !ns.Optional || !ns.Computed {
		t.Errorf("namespace should be Optional+Computed, got Optional=%v Computed=%v", ns.Optional, ns.Computed)
	}
	if ns.StringDefault != "system" {
		t.Errorf("StringDefault = %q, want %q", ns.StringDefault, "system")
	}
	if len(ns.EnumValues) != 1 || ns.EnumValues[0] != "system" {
		t.Errorf("EnumValues = %v, want [system]", ns.EnumValues)
	}
	if !result.HasStringDefaults {
		t.Error("ResourceTemplate.HasStringDefaults should be true")
	}
}

// A resource whose spec profile allows multiple namespaces keeps namespace
// Required with no default — the user must choose.
func TestExtractResourceSchema_MultiAllowedNamespaceStaysRequired(t *testing.T) {
	namespace.ClearProfiles()
	namespace.SetProfile("app_res", namespace.Profile{Allowed: []namespace.NamespaceType{namespace.Custom, namespace.Default, namespace.Shared}})
	defer namespace.ClearProfiles()

	spec, extractAPIPath := systemOnlySpec("app_res")
	result, err := ExtractResourceSchema(spec, "app_res", extractAPIPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ns := findAttr(result.Attributes, "namespace")
	if ns == nil {
		t.Fatal("namespace attribute missing")
	}
	if !ns.Required {
		t.Error("namespace must stay Required for a multi-allowed-namespace resource")
	}
	if ns.StringDefault != "" {
		t.Errorf("StringDefault = %q, want empty", ns.StringDefault)
	}
	if len(ns.EnumValues) != 0 {
		t.Errorf("EnumValues = %v, want empty", ns.EnumValues)
	}
}

// With no registered profile, namespace behaviour is unchanged (Required when the
// API path is namespaced).
func TestExtractResourceSchema_NoProfileStaysRequired(t *testing.T) {
	namespace.ClearProfiles()
	spec, extractAPIPath := systemOnlySpec("unprofiled_res")
	result, err := ExtractResourceSchema(spec, "unprofiled_res", extractAPIPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ns := findAttr(result.Attributes, "namespace")
	if ns == nil {
		t.Fatal("namespace attribute missing")
	}
	if !ns.Required || ns.StringDefault != "" {
		t.Errorf("unprofiled namespace should be Required with no default; got Required=%v StringDefault=%q", ns.Required, ns.StringDefault)
	}
}

func TestExtractOneOfGroups_Empty(t *testing.T) {
	spec := &openapi.Spec{}
	result := ExtractOneOfGroups(spec, "NonExistent")
	if len(result) != 0 {
		t.Errorf("Expected empty map, got %d groups", len(result))
	}
}

func TestExtractOneOfGroups_FromCache(t *testing.T) {
	spec := &openapi.Spec{}

	// Put raw schema data into cache
	RawSpecCache["TestType"] = map[string]interface{}{
		"x-ves-oneof-field-choice": []interface{}{"field_a", "field_b"},
		"type":                     "object",
	}
	defer delete(RawSpecCache, "TestType")

	result := ExtractOneOfGroups(spec, "TestType")
	if len(result) != 1 {
		t.Fatalf("Expected 1 group, got %d", len(result))
	}
	fields, ok := result["choice"]
	if !ok {
		t.Fatal("Expected group 'choice'")
	}
	if len(fields) != 2 {
		t.Fatalf("Expected 2 fields in group, got %d", len(fields))
	}
	if fields[0] != "field_a" || fields[1] != "field_b" {
		t.Errorf("Unexpected fields: %v", fields)
	}
}

func TestExtractOneOfGroups_StringFormat(t *testing.T) {
	spec := &openapi.Spec{}

	RawSpecCache["StringType"] = map[string]interface{}{
		"x-ves-oneof-field-mode": `["fast","slow"]`,
	}
	defer delete(RawSpecCache, "StringType")

	result := ExtractOneOfGroups(spec, "StringType")
	if len(result) != 1 {
		t.Fatalf("Expected 1 group, got %d", len(result))
	}
	fields := result["mode"]
	if len(fields) != 2 {
		t.Fatalf("Expected 2 fields, got %d", len(fields))
	}
}

func TestExtractResourceSchema_NoCreateSpecType(t *testing.T) {
	spec := &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				"SomeOtherType": {Type: "object"},
			},
		},
	}
	extractAPIPath := func(spec *openapi.Spec, resourceName string) (string, string, bool) {
		return "/api/test", "/api/test/{name}", true
	}
	_, err := ExtractResourceSchema(spec, "test_resource", extractAPIPath)
	if err == nil {
		t.Error("Expected error when no CreateSpecType found")
	}
}

func TestExtractResourceSchema_SchemaPrefixMatch(t *testing.T) {
	// Specs like fast_acl use "schemafast_aclCreateSpecType" (with "schema" prefix).
	// The generator must match this before a shorter name like "fast_acl_ruleCreateSpecType".
	spec := &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				"fast_acl_ruleCreateSpecType": {
					Type:       "object",
					Properties: map[string]openapi.Schema{"action": {Type: "string"}},
				},
				"schemafast_aclCreateSpecType": {
					Type:       "object",
					Properties: map[string]openapi.Schema{"re_acl": {Type: "object"}, "site_acl": {Type: "object"}},
				},
			},
		},
	}
	extractAPIPath := func(spec *openapi.Spec, resourceName string) (string, string, bool) {
		return "/api/config/namespaces/%s/fast_acls", "/api/config/namespaces/%s/fast_acls/%s", true
	}
	result, err := ExtractResourceSchema(spec, "fast_acl", extractAPIPath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	foundReAcl := false
	for _, attr := range result.Attributes {
		if attr.Name == "re_acl" {
			foundReAcl = true
		}
	}
	if !foundReAcl {
		t.Error("Expected 're_acl' attribute (from schemafast_aclCreateSpecType), got fast_acl_rule schema instead")
	}
}

func TestExtractResourceSchema_Basic(t *testing.T) {
	spec := &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				"ves.io.schema.my_resource.CreateSpecType": {
					Type:        "object",
					Description: "A test resource",
					Properties: map[string]openapi.Schema{
						"port": {Type: "integer", Description: "Port number"},
					},
				},
			},
		},
	}
	extractAPIPath := func(spec *openapi.Spec, resourceName string) (string, string, bool) {
		return "/api/config/namespaces/%s/my_resources", "/api/config/namespaces/%s/my_resources/%s", true
	}
	result, err := ExtractResourceSchema(spec, "my_resource", extractAPIPath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.Name != "my_resource" {
		t.Errorf("Name = %q, want %q", result.Name, "my_resource")
	}
	if result.HasNamespaceInPath != true {
		t.Error("HasNamespaceInPath should be true")
	}
	// Should have standard attrs (name, namespace, annotations, description, disable, labels, id) + port
	foundPort := false
	for _, attr := range result.Attributes {
		if attr.Name == "port" {
			foundPort = true
		}
	}
	if !foundPort {
		t.Error("Expected 'port' attribute in result")
	}
}
