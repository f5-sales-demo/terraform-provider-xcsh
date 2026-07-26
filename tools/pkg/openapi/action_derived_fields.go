// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

// An action request body can require a field that is a fact about the object
// being acted on rather than user input — the F5 XC registration approve is
// rejected with "Validation approval: Passport is required" unless it carries
// the registration's own passport, and the API accepts only that exact object
// echoed back (#1355). tools/action-derived-fields.json (keyed by resource
// TitleCase) declares those fields and where to read them from; the codegen
// emits a Create that reads the sibling object and fills them in. See
// LoadImportIDFields / LoadExposeUID for the sibling data-driven codegen
// patterns.
//
// Unlike those siblings this loader FAILS LOUDLY. A missing or malformed data
// file silently degrading to "no derived fields" would regenerate exactly the
// broken resource #1355 is about, and the resulting 500 surfaces only against a
// live tenant.

var (
	actionDerivedOnce sync.Once
	actionDerivedMap  map[string][]ActionDerivedField
	actionDerivedErr  error
)

// actionDerivedFieldsPath resolves tools/action-derived-fields.json relative to
// THIS source file, so it is found no matter which package's directory `go test`
// or the generator happens to run in.
func actionDerivedFieldsPath() (string, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("cannot resolve the action-derived-fields.json location: runtime.Caller failed")
	}
	return filepath.Join(filepath.Dir(file), "..", "..", "action-derived-fields.json"), nil
}

func loadActionDerivedFields() {
	actionDerivedMap = map[string][]ActionDerivedField{}

	jsonPath, err := actionDerivedFieldsPath()
	if err != nil {
		actionDerivedErr = err
		return
	}
	data, err := os.ReadFile(jsonPath) // #nosec G304 -- path derived from this source file, not user input
	if err != nil {
		actionDerivedErr = fmt.Errorf("reading %s: %w", jsonPath, err)
		return
	}
	parsed, err := parseActionDerivedFieldsJSON(data)
	if err != nil {
		actionDerivedErr = fmt.Errorf("parsing %s: %w", jsonPath, err)
		return
	}
	actionDerivedMap = parsed
}

// parseActionDerivedFieldsJSON decodes the data file into resource TitleCase ->
// derived fields, skipping the string "_comment" documentation key, and derives
// each field's Go and wire names. Every entry is validated: a typo that produced
// an empty field name or a pathless source would generate code that cannot
// supply the value the API requires.
func parseActionDerivedFieldsJSON(data []byte) (map[string][]ActionDerivedField, error) {
	out := map[string][]ActionDerivedField{}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	for resource, rawFields := range raw {
		if resource == "_comment" {
			continue
		}
		var fields []ActionDerivedField
		if err := json.Unmarshal(rawFields, &fields); err != nil {
			return nil, fmt.Errorf("resource %s: %w", resource, err)
		}
		for i := range fields {
			if fields[i].Field == "" {
				return nil, fmt.Errorf("resource %s: entry %d has an empty \"field\"", resource, i)
			}
			if len(fields[i].Sources) == 0 {
				return nil, fmt.Errorf("resource %s: field %q declares no \"sources\" to read it from", resource, fields[i].Field)
			}
			for _, src := range fields[i].Sources {
				if strings.TrimSpace(src) == "" {
					return nil, fmt.Errorf("resource %s: field %q has an empty source path", resource, fields[i].Field)
				}
			}
			fields[i].GoName = toGoFieldName(fields[i].Field)
			fields[i].JSONName = fields[i].Field
		}
		out[resource] = fields
	}
	return out, nil
}

// toGoFieldName converts a snake_case wire property name to an exported Go field
// name (passport -> Passport, long_term_token -> LongTermToken).
func toGoFieldName(name string) string {
	parts := strings.Split(name, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]))
		b.WriteString(p[1:])
	}
	return b.String()
}

// LoadActionDerivedFields returns the server-derived request-body fields declared
// for the action resource (by TitleCase), or nil if it declares none. An
// unreadable or malformed data file is an error, never an empty result.
func LoadActionDerivedFields(resourceTitleCase string) ([]ActionDerivedField, error) {
	actionDerivedOnce.Do(loadActionDerivedFields)
	if actionDerivedErr != nil {
		return nil, actionDerivedErr
	}
	return actionDerivedMap[resourceTitleCase], nil
}
