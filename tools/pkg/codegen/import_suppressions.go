// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

// Import-default suppression: per resource (title-case model prefix), the oneof
// members the F5 XC API ALWAYS returns as the server default for their group. On
// `terraform import` there is no prior config to preserve, so the flatten would
// otherwise populate every such default member and the next plan would show
// spurious drift. Suppressing the DEFAULT member on import is semantically safe:
// omitting it means the server re-applies the same default. Non-default and
// user-intent markers (e.g. app_firewall, advertise_on_public_default_vip) are
// NOT listed and still import normally.
//
// Data lives in tools/import-default-suppressions.json (auto-populated by
// tools/discover-defaults.go against a live tenant). The seed below is the
// built-in fallback used when that file is absent; the JSON is merged over it.
// See tracking issue #1006.
var importDefaultSuppressionsSeed = map[string][]string{
	"Healthcheck": {"headers"},
	"HTTPLoadBalancer": {
		"dns_volterra_managed",
		"default_sensitive_data_policy",
		"disable_api_definition",
		"disable_api_discovery",
		"disable_api_testing",
		"disable_malicious_user_detection",
		"disable_malware_protection",
		"disable_rate_limit",
		"disable_threat_mesh",
		"disable_trust_client_ip_headers",
		"disable_waf",
		"l7_ddos_protection",
		"no_challenge",
		"round_robin",
		"service_policies_from_namespace",
		"user_id_client_ip",
	},
}

var (
	suppressOnce sync.Once
	suppressMap  map[string]map[string]bool
)

// loadImportSuppressions builds the effective suppression map: the built-in seed
// overlaid with tools/import-default-suppressions.json when present.
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
	for r, members := range importDefaultSuppressionsSeed {
		add(r, members)
	}
	if _, file, _, ok := runtime.Caller(0); ok {
		jsonPath := filepath.Join(filepath.Dir(file), "..", "..", "import-default-suppressions.json")
		if data, err := os.ReadFile(jsonPath); err == nil {
			var fromJSON map[string][]string
			if json.Unmarshal(data, &fromJSON) == nil {
				for r, members := range fromJSON {
					if r == "_comment" {
						continue
					}
					add(r, members)
				}
			}
		}
	}
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
