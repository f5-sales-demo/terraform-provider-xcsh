// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"strings"
	"testing"
)

func csdResolver(resource string) (string, bool) {
	if resource == "protected_domain" {
		return "/api/shape/csd/namespaces/%s/protected_domains", true
	}
	return "", false
}

func schemaWithRequires(prop string, entries []RequiresEntry) Schema {
	return Schema{
		Properties: map[string]Schema{
			prop: {XF5XCRequires: entries},
		},
	}
}

func TestDeriveRequirementPreflights_SameNamespaceResource(t *testing.T) {
	createSpec := schemaWithRequires("client_side_defense", []RequiresEntry{
		{
			Field:            "client_side_defense",
			RequiresResource: &RequiresResource{Resource: "protected_domain", Scope: "same_namespace"},
			Reason:           "CSD needs a protected_domain in the same namespace.",
		},
	})

	got := DeriveRequirementPreflights(createSpec, csdResolver)
	if len(got) != 1 {
		t.Fatalf("expected 1 derived preflight, got %d", len(got))
	}
	p := got[0]
	if p.WhenField != "client_side_defense" {
		t.Errorf("WhenField = %q, want client_side_defense", p.WhenField)
	}
	if p.ListPath != "/api/shape/csd/namespaces/%s/protected_domains" {
		t.Errorf("ListPath = %q", p.ListPath)
	}
	// Well-formed invariants (mirror TestLoadPreflights_AllEntriesWellFormed).
	if strings.Count(p.ListPath, "%s") != 1 {
		t.Errorf("ListPath must contain exactly one %%s: %q", p.ListPath)
	}
	if strings.Count(p.ErrorDetail, "%s") != 1 {
		t.Errorf("ErrorDetail must contain exactly one %%s: %q", p.ErrorDetail)
	}
	if p.ErrorTitle == "" || p.Requires == "" {
		t.Errorf("ErrorTitle/Requires must be non-empty: %+v", p)
	}
}

func TestDeriveRequirementPreflights_SiblingFieldIgnored(t *testing.T) {
	createSpec := schemaWithRequires("foo", []RequiresEntry{
		{Field: "foo", RequiresField: "spec.enable_waf"},
	})
	if got := DeriveRequirementPreflights(createSpec, csdResolver); len(got) != 0 {
		t.Errorf("sibling-field requires must not derive a preflight, got %d", len(got))
	}
}

func TestDeriveRequirementPreflights_UnresolvableResourceSkipped(t *testing.T) {
	createSpec := schemaWithRequires("thing", []RequiresEntry{
		{Field: "thing", RequiresResource: &RequiresResource{Resource: "nonexistent", Scope: "same_namespace"}},
	})
	// resolver returns ok=false for unknown resources -> never emit a broken list_path.
	if got := DeriveRequirementPreflights(createSpec, csdResolver); len(got) != 0 {
		t.Errorf("unresolvable target must be skipped, got %d", len(got))
	}
}

func TestMergePreflights_JSONOverridesByWhenField(t *testing.T) {
	derived := []RequirementPreflight{
		{WhenField: "client_side_defense", ListPath: "derived/%s/x", ErrorDetail: "d %s"},
		{WhenField: "only_derived", ListPath: "d/%s", ErrorDetail: "e %s"},
	}
	override := []RequirementPreflight{
		{WhenField: "client_side_defense", ListPath: "json/%s/protected_domains", ErrorDetail: "j %s"},
		{WhenField: "only_json", ListPath: "j/%s", ErrorDetail: "k %s"},
	}
	merged := MergePreflights(derived, override)
	byField := map[string]RequirementPreflight{}
	for _, p := range merged {
		byField[p.WhenField] = p
	}
	if len(merged) != 3 {
		t.Fatalf("expected 3 merged (cds override + only_derived + only_json), got %d: %+v", len(merged), merged)
	}
	if byField["client_side_defense"].ListPath != "json/%s/protected_domains" {
		t.Errorf("JSON must override derived for client_side_defense, got %q", byField["client_side_defense"].ListPath)
	}
	if _, ok := byField["only_derived"]; !ok {
		t.Errorf("derived-only entry must survive")
	}
	if _, ok := byField["only_json"]; !ok {
		t.Errorf("json-only entry must survive")
	}
}
