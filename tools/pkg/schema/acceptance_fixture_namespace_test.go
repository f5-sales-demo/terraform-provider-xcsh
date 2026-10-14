// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package schema

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"testing"
)

func TestAcceptanceFixtureNamespaceRequiresVerifiedConstraint(t *testing.T) {
	for _, tc := range []struct {
		name     string
		allowed  []string
		enforced bool
		want     string
	}{
		{"shared", []string{"shared"}, true, "shared"},
		{"system", []string{"system"}, true, "system"},
		{"unverified", []string{"system"}, false, ""},
		{"multi", []string{"system", "custom"}, true, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			spec := &openapi.DomainSpec{}
			spec.Components.Schemas = map[string]openapi.Schema{"viewsprobeCreateSpecType": {XF5XCNamespaceProfile: &openapi.NamespaceProfileSpec{Constraint: &openapi.NamespaceConstraint{Allowed: tc.allowed, Enforced: tc.enforced}}}}
			got, err := AcceptanceFixtureNamespace(spec, "probe")
			if err != nil || got != tc.want {
				t.Fatalf("namespace=%q error=%v want=%q", got, err, tc.want)
			}
		})
	}
}
