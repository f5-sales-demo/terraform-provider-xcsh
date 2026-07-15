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
		"disable_client_side_defense",
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
		// enable_api_discovery inner oneof defaults: the server materializes these
		// whenever discovery is enabled, but discover-defaults.go can't observe them
		// (its probe LB never enables discovery). Hand-seeded so a bare
		// enable_api_discovery {} is import-clean and consistent on first apply.
		// Matched by leaf name at any depth (see isImportDefaultSuppressed).
		"default_api_auth_discovery",
		"disable_learn_from_redirect_traffic",
		// SP3 API Protection (#41) server-default oneof base markers: the API echoes
		// any_client {} alongside a concrete client_matcher arm (ip_prefix_list /
		// ip_threat_category_list), and skip_response_validation {} on every
		// open_api_validation_rules entry's validation_mode. The module never declares
		// either (it omits client_matcher for "match any"), so both must be suppressed
		// on import to keep api_protection_rules / validation_custom_list round-trip
		// clean. discover-defaults.go can't observe them (its probe LB configures no
		// api_protection_rules / validation_custom_list). Verified live (f5-sales-demo
		// webapp-api-protection SP3 matrix variants 004, 006). any_ip is the same class:
		// the default source-IP sub-oneof member inside a client_matcher
		// ip_threat_category_list arm, echoed alongside ip_threat_categories (variant 002).
		"any_client",
		"any_ip",
		"skip_response_validation",
		// SP4 API Testing (#45): "standard" is the server-default base member of the
		// api_testing credentials credentials_choice oneof, echoed on every credential.
		// It appears on the LB inline api_testing block here and on the standalone
		// xcsh_api_testing (APITesting, below). Suppress so a concrete-arm credential
		// (api_key/basic_auth/bearer_token) is round-trip-import clean. Verified live.
		"standard",
	},
	"APITesting": {
		"standard",
	},
	// Coverage Batch B (#51): the server materializes the base member of each
	// client-matcher oneof on a rate_limiter_policy rule that omits that matcher
	// (any_country observed live on a rule with asn_list but no country; any_asn /
	// any_ip are the same class for the ASN / IP oneofs). The module omits a matcher
	// for "match any", so suppress these on import to keep the standalone
	// xcsh_rate_limiter_policy round-trip clean.
	"RateLimiterPolicy": {
		"any_asn",
		"any_country",
		"any_ip",
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
			for r, members := range parseSuppressionsJSON(data) {
				add(r, members)
			}
		}
	}
}

// parseSuppressionsJSON parses the suppression data file into resource -> members.
// It parses via RawMessage so the string "_comment" field does not break
// unmarshalling of the []string resource entries (a regression that silently
// disabled the whole JSON, leaving only the built-in seed active).
func parseSuppressionsJSON(data []byte) map[string][]string {
	out := map[string][]string{}
	var raw map[string]json.RawMessage
	if json.Unmarshal(data, &raw) != nil {
		return out
	}
	for r, rawMembers := range raw {
		if r == "_comment" {
			continue
		}
		var members []string
		if json.Unmarshal(rawMembers, &members) == nil {
			out[r] = members
		}
	}
	return out
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
