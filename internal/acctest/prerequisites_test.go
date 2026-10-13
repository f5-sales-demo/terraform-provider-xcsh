// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"encoding/json"
	"testing"
)

func TestExternalPrerequisiteDoesNotLeakCredential(t *testing.T) {
	t.Setenv("XCSH_ACC_CLOUD_CREDENTIAL", "")
	t.Run("missing", func(t *testing.T) { RequireExternalCredential(t, "cloud") })
}
func TestFixtureLiteralTypes(t *testing.T) {
	schema := &FixtureSchema{Resources: map[string]fixtureBlockSchema{"xcsh_" + "probe": {Attributes: map[string]struct {
		Required bool            `json:"required"`
		Computed bool            `json:"computed"`
		Optional bool            `json:"optional"`
		Type     json.RawMessage `json:"type,omitempty"`
	}{
		"marker": {Optional: true, Type: json.RawMessage(`["object",{}]`)},
	}}}}
	diagnostics := schema.Validate("typed", `resource "xcsh_probe" "test" { marker = true }`)
	if len(diagnostics) != 1 || diagnostics[0].Category != "incorrect_attribute_type" {
		t.Fatal(diagnostics)
	}
}
