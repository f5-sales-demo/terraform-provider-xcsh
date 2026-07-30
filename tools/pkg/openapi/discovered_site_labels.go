// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

// F5 XC decorates a node-backed site object with six labels discovered from the
// hardware and OS it booted on — domain, host-os-version, hw-model,
// hw-serial-number, hw-vendor, hw-version. They carry no reserved prefix, so a
// prefix test cannot recognise them, and until #1391 they reached Terraform state
// and were proposed for deletion on every plan.
//
// The filter for those six is opt-in per resource
// (tools/discovered-site-labels.json, keyed by TitleCase) rather than global,
// because the names are ordinary enough that a user could legitimately own a label
// called `domain` on an unrelated object. Only the resources F5 XC actually
// decorates filter them. The two platform PREFIXES (ves.io/, internal.ves.io/) are
// filtered everywhere and are not driven by this file.
//
// See LoadExposeUID / LoadImportIDFields for the sibling data-driven codegen
// patterns.

var (
	discoveredSiteLabelsOnce sync.Once
	discoveredSiteLabelsMap  map[string]bool
)

func loadDiscoveredSiteLabels() {
	discoveredSiteLabelsMap = map[string]bool{}
	if _, file, _, ok := runtime.Caller(0); ok {
		jsonPath := filepath.Join(filepath.Dir(file), "..", "..", "discovered-site-labels.json")
		if data, err := os.ReadFile(jsonPath); err == nil {
			discoveredSiteLabelsMap = parseDiscoveredSiteLabelsJSON(data)
		}
	}
}

// parseDiscoveredSiteLabelsJSON decodes the data file into resource -> bool,
// skipping the string "_comment" documentation key.
func parseDiscoveredSiteLabelsJSON(data []byte) map[string]bool {
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

// LoadDiscoveredSiteLabels reports whether the resource (by TitleCase) is one F5 XC
// decorates with hardware/OS discovery labels, and so must filter them out of the
// read-back.
func LoadDiscoveredSiteLabels(resourceTitleCase string) bool {
	discoveredSiteLabelsOnce.Do(loadDiscoveredSiteLabels)
	return discoveredSiteLabelsMap[resourceTitleCase]
}
