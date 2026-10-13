// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package schema

import "github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"

// AcceptanceFixtureNamespace shares the provider's exact-envelope resolution and
// uses only an enforced singleton namespace. An empty result keeps a disposable
// application namespace; recommendations do not prove a constraint.
func AcceptanceFixtureNamespace(spec *openapi.DomainSpec, resourceName string) (string, error) {
	resolved, _, found, err := ResolveNamespaceProfileSchema(spec.Components.Schemas, resourceName, true)
	if err != nil {
		return "", err
	}
	var profile *openapi.NamespaceProfileSpec
	if found {
		profile = resolved.XF5XCNamespaceProfile
	}
	if profile == nil {
		profile = spec.Info.XF5XCNamespaceProfile
	}
	if profile == nil {
		profile = spec.XF5XCNamespaceProfile
	}
	if profile == nil || profile.Constraint == nil || !profile.Constraint.Enforced || len(profile.Constraint.Allowed) != 1 {
		return "", nil
	}
	value := profile.Constraint.Allowed[0]
	if value == "system" || value == "shared" {
		return value, nil
	}
	return "", nil
}

// AcceptanceReusableFixtureNamespace chooses an existing permitted tenant namespace
// to avoid making unrelated namespace creation a prerequisite for every fixture.
func AcceptanceReusableFixtureNamespace(spec *openapi.DomainSpec, resourceName string) (string, error) {
	resolved, _, found, err := ResolveNamespaceProfileSchema(spec.Components.Schemas, resourceName, true)
	if err != nil {
		return "", err
	}
	var profile *openapi.NamespaceProfileSpec
	if found {
		profile = resolved.XF5XCNamespaceProfile
	}
	if profile == nil {
		profile = spec.Info.XF5XCNamespaceProfile
	}
	if profile == nil {
		profile = spec.XF5XCNamespaceProfile
	}
	if profile == nil || profile.Constraint == nil || !profile.Constraint.Enforced {
		return "", nil
	}
	for _, candidate := range []string{"shared", "system", "default"} {
		for _, allowed := range profile.Constraint.Allowed {
			if candidate == allowed {
				return candidate, nil
			}
		}
	}
	return "", nil
}
