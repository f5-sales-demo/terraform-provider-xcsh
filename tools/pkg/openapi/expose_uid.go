// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

// Most F5 XC objects carry a server-generated system_metadata.uid, but the
// generator surfaces it as a Terraform attribute only for the handful of
// resources where that uid is the value a consumer actually needs (e.g. a CE
// registration token: the token VALUE is system_metadata.uid, not the object's
// name). tools/expose-uid.json (keyed by resource TitleCase) opts those
// resources in; the codegen additionally verifies the response schema really
// carries system_metadata.uid (schema.ResponseHasSystemMetadataUID) before
// emitting anything, so the opt-in list stays schema-driven and surgically
// scoped. See RequirementPreflight / LoadImportIDFields for the sibling
// data-driven codegen patterns.

var (
	exposeUIDOnce sync.Once
	exposeUIDMap  map[string]bool
)

func loadExposeUID() {
	exposeUIDMap = map[string]bool{}
	if _, file, _, ok := runtime.Caller(0); ok {
		jsonPath := filepath.Join(filepath.Dir(file), "..", "..", "expose-uid.json")
		if data, err := os.ReadFile(jsonPath); err == nil {
			exposeUIDMap = parseExposeUIDJSON(data)
		}
	}
}

// parseExposeUIDJSON decodes the data file into resource -> bool, skipping the
// string "_comment" documentation key.
func parseExposeUIDJSON(data []byte) map[string]bool {
	out := map[string]bool{}
	var raw map[string]json.RawMessage
	if json.Unmarshal(data, &raw) != nil {
		return out
	}
	for resource, rawVal := range raw {
		if resource == "_comment" {
			continue
		}
		var enabled bool
		if json.Unmarshal(rawVal, &enabled) == nil {
			out[resource] = enabled
		}
	}
	return out
}

// LoadExposeUID reports whether the resource (by TitleCase) is opted in to
// surfacing system_metadata.uid as a Computed `uid` attribute.
func LoadExposeUID(resourceTitleCase string) bool {
	exposeUIDOnce.Do(loadExposeUID)
	return exposeUIDMap[resourceTitleCase]
}
