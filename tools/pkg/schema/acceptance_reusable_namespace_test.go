// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package schema

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"testing"
)

func TestFixtureMayReuseVerifiedSharedNamespace(t *testing.T) {
	spec := &openapi.DomainSpec{}
	spec.Components.Schemas = map[string]openapi.Schema{"viewsprobeCreateSpecType": {XF5XCNamespaceProfile: &openapi.NamespaceProfileSpec{Constraint: &openapi.NamespaceConstraint{Allowed: []string{"custom", "default", "shared"}, Enforced: true}}}}
	namespace, err := AcceptanceReusableFixtureNamespace(spec, "probe")
	if err != nil || namespace != "shared" {
		t.Fatal("verified shared namespace not reusable")
	}
	spec.Components.Schemas["viewsprobeCreateSpecType"] = openapi.Schema{XF5XCNamespaceProfile: &openapi.NamespaceProfileSpec{Constraint: &openapi.NamespaceConstraint{Allowed: []string{"custom"}, Enforced: true}}}
	namespace, err = AcceptanceReusableFixtureNamespace(spec, "probe")
	if err != nil || namespace != "" {
		t.Fatal("custom-only resource assigned shared namespace")
	}
}
