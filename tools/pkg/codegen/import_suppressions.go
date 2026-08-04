// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/suppress"
)

// Import-default suppression: per resource (title-case model prefix), the empty
// marker blocks the F5 XC API ALWAYS returns as a server default. Two classes:
//  1. oneof base members the API echoes for their group (e.g. disable_waf,
//     any_client, round_robin);
//  2. plain optional message blocks the API materializes as present-but-empty on
//     every element — including inside list elements — even though they carry no
//     meaning empty (e.g. origin_servers[].labels {}, default_route_pools[].
//     endpoint_subsets {}; see #1103).
//
// On `terraform import` there is no prior config to preserve, so the flatten would
// otherwise populate every such marker and the next plan would show spurious drift.
// Suppressing the marker on import is semantically safe: omitting it means the
// server re-applies the same default. Non-default and user-intent markers (e.g.
// app_firewall, advertise_on_public_default_vip) are NOT listed and still import
// normally. Matched by leaf name at any depth (see isImportDefaultSuppressed), so
// one entry per resource covers every nesting/list depth it appears at.
//
// Canonical measured data lives in tools/import-default-suppressions.json and is
// updated from live-tenant observations by tools/emit-import-suppressions.go.
// See tracking issue #1006.

var (
	suppressOnce sync.Once
	suppressMap  map[string]map[string]bool
)

// loadImportSuppressions loads the canonical measured suppression data. A missing
// or malformed data file is a generator defect and must stop generation visibly.
func loadImportSuppressions() {
	suppressMap = map[string]map[string]bool{}
	add := func(resource string, members []string) {
		if suppressMap[resource] == nil {
			suppressMap[resource] = map[string]bool{}
		}
		for _, m := range members {
			suppressMap[resource][m] = true
		}
	}
	repositoryRoot, err := findRepositoryRoot()
	if err != nil {
		panic(err)
	}
	jsonPath := filepath.Join(repositoryRoot, "tools", "import-default-suppressions.json")
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		panic(fmt.Sprintf("read canonical import suppressions %s: %v", jsonPath, err))
	}
	parsed, err := parseSuppressionsJSON(data)
	if err != nil {
		panic(fmt.Sprintf("parse canonical import suppressions %s: %v", jsonPath, err))
	}
	for resourceName, members := range parsed {
		add(resourceName, members)
	}
}

// findRepositoryRoot locates the checked-out module without depending on source
// file paths. runtime.Caller paths are module-relative in -trimpath builds and
// therefore cannot locate repository data files.
func findRepositoryRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("locate repository root: get working directory: %w", err)
	}
	return findRepositoryRootFrom(dir)
}

func findRepositoryRootFrom(dir string) (string, error) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return "", fmt.Errorf("locate repository root from %s: %w", dir, err)
	}
	for {
		goMod := filepath.Join(dir, "go.mod")
		dataFile := filepath.Join(dir, "tools", "import-default-suppressions.json")
		if _, goModErr := os.Stat(goMod); goModErr == nil {
			if _, dataErr := os.Stat(dataFile); dataErr == nil {
				return dir, nil
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("locate repository root from %s: go.mod and tools/import-default-suppressions.json not found", dir)
		}
		dir = parent
	}
}

func parseSuppressionsJSON(data []byte) (map[string][]string, error) {
	parsed, _, err := suppress.ParseCanonicalSuppressions(data, "tools/import-default-suppressions.json")
	return parsed, err
}

// isImportDefaultSuppressed reports whether the given member of the given resource
// is a server-default oneof member that must not be populated from the API
// response on the import path.
func isImportDefaultSuppressed(resourceTitleCase, jsonName string) bool {
	suppressOnce.Do(loadImportSuppressions)
	members, ok := suppressMap[resourceTitleCase]
	if !ok {
		return false
	}
	return members[jsonName]
}

// isSystemManagedFilteredList reports whether the given list member of the given
// resource holds F5 XC system-managed elements that must be dropped from Terraform
// state on EVERY refresh path (Read/Create/Update, not just import), so a config
// that does not declare them does not plan a forbidden delete.
//
// DNSZone rr_set_group is the case: when a primary block sets
// allow_http_lb_managed_records = true, F5 XC auto-creates a reserved
// "x-ves-io-managed" rr_set_group holding load-balancer-owned records. The caller
// may not modify or delete it (403 FORBIDDEN), so the generated flatten filters it
// via filterSystemManagedRrSetGroups (internal/provider/provider_helpers.go).
func isSystemManagedFilteredList(resourceTitleCase, jsonName string) bool {
	return resourceTitleCase == "DNSZone" && jsonName == "rr_set_group"
}

// suppressionRootOnly scopes a suppressed leaf to the resource ROOT only. A few leaves are a
// server-default oneof member at the top level (must suppress so a bare resource imports clean)
// AND a legitimately user-DECLARED oneof arm when nested. Because isImportDefaultSuppressed
// matches by leaf name at any depth, suppressing such a leaf everywhere strips the nested
// declared value on import, drifting the next plan.
//
// http_loadbalancer disable_waf is the case (#1145): the LB-level WAF oneof default vs. the
// per-route routes[].{simple_route}.advanced_options "disable WAF for this route" choice. Root
// single blocks render via renderUnmarshalTopLevelSingle (which keeps suppressing these leaves);
// nested single blocks — including single blocks inside list elements like routes[] — render via
// renderUnmarshalSingleChild, which skips suppression for a root-only leaf so the declared nested
// marker reads back and round-trips. Keep this list tight: only a leaf a live round-trip proves
// is a DECLARED arm when nested (never a server-default-on-omit there) belongs here.
var suppressionRootOnly = map[string][]string{
	"HTTPLoadBalancer": {"disable_waf"},
}

// isSuppressionRootOnly reports whether the given suppressed leaf must be suppressed only at the
// resource root (and therefore read back — not suppressed — when nested).
func isSuppressionRootOnly(resourceTitleCase, jsonName string) bool {
	for _, leaf := range suppressionRootOnly[resourceTitleCase] {
		if leaf == jsonName {
			return true
		}
	}
	return false
}
