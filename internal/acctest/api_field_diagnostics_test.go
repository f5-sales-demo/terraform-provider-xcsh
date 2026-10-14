// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import "testing"

func TestAPIFieldDiagnosticsUseOnlyInstalledSchemaFields(t *testing.T) {
	schema := &FixtureSchema{Resources: map[string]fixtureBlockSchema{"xcsh_" + "probe": {Blocks: map[string]struct {
		Block    fixtureBlockSchema `json:"block"`
		MinItems int                `json:"min_items"`
	}{"http": {}}}}}
	fields := diagnosticSchemaFields(schema, "must configure http; SECRET_SENTINEL invalid")
	if len(fields) != 1 || fields[0] != "http" {
		t.Fatal("required schema field unavailable or raw input leaked")
	}
}
