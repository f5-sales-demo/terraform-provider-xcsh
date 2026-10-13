// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"reflect"
	"testing"
)

func TestFixtureValidationRejectsSchemaAndReferences(t *testing.T) {
	schema := &FixtureSchema{Resources: map[string]fixtureBlockSchema{"xcsh_" + "probe": {Attributes: map[string]struct {
		Required bool `json:"required"`
		Computed bool `json:"computed"`
		Optional bool `json:"optional"`
	}{
		"name": {Required: true}, "namespace": {Required: true},
	}}}}
	diagnostics := schema.Validate("probe", `resource "xcsh_probe" "test" {
 namespace = xcsh_namespace.missing.name
 unsupported = true
 labels {}
}`)
	categories := map[string]bool{}
	for _, d := range diagnostics {
		categories[d.Category] = true
	}
	for _, want := range []string{"missing_required_argument", "unsupported_argument", "unsupported_block", "undeclared_reference"} {
		if !categories[want] {
			t.Fatalf("missing %s: %#v", want, diagnostics)
		}
	}
}
func TestFixtureRequirementsUseExactRegisteredTest(t *testing.T) {
	RegisterFixture("probe", func() string {
		return `resource "xcsh_namespace" "test" { name="tf-acc-test-fixture" }
resource "xcsh_origin_pool" "test" {
 name="tf-acc-test-fixture-pool"
 namespace=xcsh_namespace.test.name
}`
	})
	RegisterTestFixtures("TestAccAPISynthetic_basic", []string{"probe"})
	requirements, err := TestOperationRequirements("TestAccAPISynthetic_basic")
	if err != nil || len(requirements) != 2 || requirements[0].Path != "/api/web/namespaces" || requirements[1].Namespace != "disposable" {
		t.Fatalf("requirements: %v %v", requirements, err)
	}
	_, err = TestOperationRequirements("TestAccApiSynthetic_basic")
	if err == nil {
		t.Fatal("capitalization prefix guessed an unregistered test")
	}
	requirements2, err := TestOperationRequirements("TestAccAPISynthetic_basic")
	if err != nil || !reflect.DeepEqual(requirements, requirements2) {
		t.Fatal("inventory not deterministic")
	}
}
