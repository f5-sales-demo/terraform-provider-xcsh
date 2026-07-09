// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

// Apply-time requirement pre-flights are declared in tools/preflight-requirements.json
// (source of truth: x-f5xc-requires in api-specs-enriched) and compiled into each
// resource's Create/Update by the codegen. This encodes the prerequisite in the
// shipped provider binary, so every remote workstation enforces it identically
// instead of relying on out-of-band knowledge. See RequirementPreflight (types.go).

var (
	preflightOnce sync.Once
	preflightMap  map[string][]RequirementPreflight
)

// loadPreflights reads tools/preflight-requirements.json (relative to this source
// file) into the resource-keyed map. A missing or malformed file yields an empty
// map — resources then generate with no pre-flight, exactly as before.
func loadPreflights() {
	preflightMap = map[string][]RequirementPreflight{}
	if _, file, _, ok := runtime.Caller(0); ok {
		jsonPath := filepath.Join(filepath.Dir(file), "..", "..", "preflight-requirements.json")
		if data, err := os.ReadFile(jsonPath); err == nil {
			preflightMap = parsePreflightsJSON(data)
		}
	}
}

// parsePreflightsJSON parses the data file into resource TitleCase -> preflights.
// It decodes via RawMessage so the string "_comment" documentation field does not
// break unmarshalling of the []RequirementPreflight resource entries.
func parsePreflightsJSON(data []byte) map[string][]RequirementPreflight {
	out := map[string][]RequirementPreflight{}
	var raw map[string]json.RawMessage
	if json.Unmarshal(data, &raw) != nil {
		return out
	}
	for resource, rawEntries := range raw {
		if resource == "_comment" {
			continue
		}
		var entries []RequirementPreflight
		if json.Unmarshal(rawEntries, &entries) == nil {
			out[resource] = entries
		}
	}
	return out
}

// LoadPreflights returns the declared apply-time pre-flights for a resource
// (by TitleCase), or an empty slice if none. WhenGoField is left unresolved here;
// the schema extractor fills it from the resource's attributes.
func LoadPreflights(resourceTitleCase string) []RequirementPreflight {
	preflightOnce.Do(loadPreflights)
	entries := preflightMap[resourceTitleCase]
	out := make([]RequirementPreflight, len(entries))
	copy(out, entries)
	return out
}
