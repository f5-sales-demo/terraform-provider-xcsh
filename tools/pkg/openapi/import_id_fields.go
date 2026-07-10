// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

// Some resources have create-only spec fields the API cannot return on read
// (e.g. the CSD protected_domain: GET-by-name is 501 and the LIST projection is
// empty). Such a field cannot be recovered on import from state, so its value is
// supplied as a trailing segment of the import ID. tools/import-id-fields.json
// (keyed by resource TitleCase) declares those fields; the codegen emits an
// ImportState that parses and sets them. See RequirementPreflight for the sibling
// data-driven codegen pattern.

var (
	importIDFieldsOnce sync.Once
	importIDFieldsMap  map[string][]string
)

func loadImportIDFields() {
	importIDFieldsMap = map[string][]string{}
	if _, file, _, ok := runtime.Caller(0); ok {
		jsonPath := filepath.Join(filepath.Dir(file), "..", "..", "import-id-fields.json")
		if data, err := os.ReadFile(jsonPath); err == nil {
			importIDFieldsMap = parseImportIDFieldsJSON(data)
		}
	}
}

// parseImportIDFieldsJSON decodes the data file into resource -> field names,
// skipping the string "_comment" documentation key.
func parseImportIDFieldsJSON(data []byte) map[string][]string {
	out := map[string][]string{}
	var raw map[string]json.RawMessage
	if json.Unmarshal(data, &raw) != nil {
		return out
	}
	for resource, rawFields := range raw {
		if resource == "_comment" {
			continue
		}
		var fields []string
		if json.Unmarshal(rawFields, &fields) == nil {
			out[resource] = fields
		}
	}
	return out
}

// LoadImportIDFields returns the create-only fields carried in a resource's import
// ID (by TitleCase), or nil if none.
func LoadImportIDFields(resourceTitleCase string) []string {
	importIDFieldsOnce.Do(loadImportIDFields)
	fields := importIDFieldsMap[resourceTitleCase]
	out := make([]string, len(fields))
	copy(out, fields)
	return out
}
