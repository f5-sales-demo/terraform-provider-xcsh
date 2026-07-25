// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

// Regression guard for #1257.
//
// CI's "Fix generated content for lint compliance" step used to apply
// tools/codespell-corrections.json to generated Go with a blind
// `txt.replace(wrong, right)`. Generated Go embeds every OpenAPI property name
// verbatim as a JSON wire key (`"blocked_sevice"`) and again inside derived
// identifiers (`BlockedSevice`, `schemaFleetBlockedSevicesListType`), so any
// correction key that occurs inside a real property name silently renames an
// API field. That is exactly what happened: `sevice`->`service` turned the
// upstream-misspelled field `blocked_sevice` into `blocked_service`, which the
// F5 XC API silently ignores, and `checkin`->`checking` broke the bot_defense
// flow-label `checkin` property the same way.
//
// The spelling pass no longer touches internal/provider/, but the corrections
// file is still applied to prose that is generated from these same property
// names (docs/, examples/). This test keeps the data itself honest so the class
// of defect cannot come back through another entry.

// codespellCorrections loads the misspelling -> correction map that CI applies.
func codespellCorrections(t *testing.T, root string) map[string]string {
	t.Helper()

	path := filepath.Join(root, "tools", "codespell-corrections.json")
	data, err := os.ReadFile(path) // #nosec G304 -- fixed repo-relative path
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}

	var file struct {
		Corrections map[string]string `json:"corrections"`
	}
	if err := json.Unmarshal(data, &file); err != nil {
		t.Fatalf("parsing %s: %v", path, err)
	}
	if len(file.Corrections) == 0 {
		t.Fatalf("%s declares no corrections; the guard would be vacuous", path)
	}
	return file.Corrections
}

// specPropertyNames returns every property name declared anywhere in the
// released spec artifact the generator consumes, discovered exactly the way
// tools/generate-all-schemas.go discovers it (openapi.FindDomainSpecFiles over
// the --spec-dir layout). Properties are collected recursively: nested objects,
// array items and map values all contribute names that reach generated code.
// Returns nil when the artifact is not present locally.
func specPropertyNames(t *testing.T, specDir string) map[string]struct{} {
	t.Helper()

	if !openapi.IsV2SpecDirectory(specDir) {
		return nil
	}

	files, err := openapi.FindDomainSpecFiles(specDir)
	if err != nil {
		t.Fatalf("discovering domain specs in %s: %v", specDir, err)
	}
	if len(files) == 0 {
		t.Fatalf("no domain specs found in %s", specDir)
	}

	names := make(map[string]struct{})
	for _, file := range files {
		data, err := os.ReadFile(file) // #nosec G304 -- path from spec discovery
		if err != nil {
			t.Fatalf("reading %s: %v", file, err)
		}
		var doc any
		if err := json.Unmarshal(data, &doc); err != nil {
			t.Fatalf("parsing %s: %v", file, err)
		}
		collectPropertyNames(doc, names)
	}
	if len(names) == 0 {
		t.Fatalf("collected no property names from %s", specDir)
	}
	return names
}

// collectPropertyNames walks a decoded JSON document and records the keys of
// every "properties" object it encounters, at any depth.
func collectPropertyNames(node any, out map[string]struct{}) {
	switch typed := node.(type) {
	case map[string]any:
		if props, ok := typed["properties"].(map[string]any); ok {
			for name := range props {
				out[name] = struct{}{}
			}
		}
		for _, child := range typed {
			collectPropertyNames(child, out)
		}
	case []any:
		for _, child := range typed {
			collectPropertyNames(child, out)
		}
	}
}

// TestCodespellCorrectionsDoNotCollideWithSpecPropertyNames asserts that no
// spelling-correction key can rewrite a real API property name.
//
// Matching rule (deliberately mirrors what the CI replacement actually does):
// a key collides with a property when the key occurs as a SUBSTRING of the
// property name, compared case-insensitively.
//
//   - Substring, not equality, because `txt.replace` is a substring operation:
//     `checkin` -> `checking` corrupts both the property `checkin` and the
//     property `msg_hdr_checking_disabled` (-> `msg_hdr_checkingg_disabled`).
//   - Case-insensitive, because generated code carries each property name both
//     verbatim (the wire key) and in CamelCase (Go identifiers), so `Sevice` is
//     just as dangerous as `sevice`.
//   - Keys are compared EXACTLY AS WRITTEN, with no trimming. Space-padded keys
//     such as `" ADN "` or `" defin "` cannot occur inside a property name, so
//     they are correctly reported as safe; trimming them would invent false
//     positives (` defin ` would "collide" with `api_definition`).
func TestCodespellCorrectionsDoNotCollideWithSpecPropertyNames(t *testing.T) {
	root := repoRootFromTest(t)

	specDir := os.Getenv("XCSH_SPEC_DIR")
	if specDir == "" {
		specDir = filepath.Join(root, "docs", "specifications", "api")
	}

	properties := specPropertyNames(t, specDir)
	if properties == nil {
		t.Skipf("spec artifact not present at %s; set XCSH_SPEC_DIR to the released spec bundle to run this guard", specDir)
	}

	corrections := codespellCorrections(t, root)

	keys := make([]string, 0, len(corrections))
	for key := range corrections {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var failures []string
	for _, key := range keys {
		lowerKey := strings.ToLower(key)
		var hits []string
		for property := range properties {
			if strings.Contains(strings.ToLower(property), lowerKey) {
				hits = append(hits, property)
			}
		}
		if len(hits) == 0 {
			continue
		}
		sort.Strings(hits)
		failures = append(failures, fmt.Sprintf(
			"  %q -> %q collides with spec propert%s: %s",
			key, corrections[key], plural(len(hits)), strings.Join(hits, ", "),
		))
	}

	if len(failures) > 0 {
		t.Errorf("tools/codespell-corrections.json contains %d key(s) that rewrite real API property names (#1257).\n%s\n\n"+
			"Every such key silently renames a wire field wherever the corrections are applied. "+
			"Remove the key, or narrow it (e.g. pad it with spaces) so it can only match prose.",
			len(failures), strings.Join(failures, "\n"))
	}
}

func plural(n int) string {
	if n == 1 {
		return "y"
	}
	return "ies"
}
