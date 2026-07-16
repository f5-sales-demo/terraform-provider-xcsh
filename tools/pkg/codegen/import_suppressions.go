// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"sync"
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
		// #1103: endpoint_subsets {} is a plain optional empty-marker block (NOT a oneof
		// member) that the API echoes on every default_route_pools[] / routes[].pools[]
		// element. The module omits it (empty carries no meaning), so it must be
		// suppressed on import or the minimal config drifts every plan (- endpoint_subsets
		// {}), cascading into computed tenant re-planning on pool/app_firewall/api refs.
		// Missed until now because the auto-derive differ was blind to list-element
		// defaults (fixed in tools/pkg/suppress/diff.go).
		"endpoint_subsets",
	},
	"APITesting": {
		"standard",
	},
	// #1103: labels {} is a plain optional empty-marker block the API echoes on every
	// origin_servers[] element (a label selector the schema models as empty-only). The
	// module omits it, so suppress on import to keep origin_pool round-trip clean.
	// Does NOT affect the top-level metadata.labels types.Map (a different read path
	// that never consults isImportDefaultSuppressed).
	"OriginPool": {
		"labels",
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
	// SPol effort (service_policy coverage): the API echoes server-default empty markers
	// throughout xcsh_service_policy (and the standalone xcsh_service_policy_rule, which
	// shares the rule shape). Seeded up front so every service_policy matrix variant is
	// round-trip-import clean. Matched by leaf name at any depth, so one list per resource
	// covers all nesting/list depths. Classes: policy server oneof base (any_server);
	// per-rule client/asn/ip oneof bases (any_client/any_asn/any_ip); action oneof bases
	// (waf_action & bot_action "none", mum_action "default"); segment_policy bases
	// (src_any/dst_any/intra_segment); per-list-element present/not-present markers
	// (check_present/check_not_present); and the 13 request_constraints max_*_none bases.
	// NOTE: port_matcher is a NON-empty server-default block ("Server applies default when
	// omitted") — the empty-marker suppression path does not cover it; handle in lock-step
	// (mirroring l7_ddos_protection) only if the live import matrix surfaces its drift.
	"ServicePolicy": {
		"any_server",
		"any_client",
		"any_asn",
		"any_ip",
		"none",
		"default",
		"src_any",
		"dst_any",
		"intra_segment",
		"check_present",
		"check_not_present",
		"max_cookie_count_none",
		"max_cookie_key_size_none",
		"max_cookie_value_size_none",
		"max_header_count_none",
		"max_header_key_size_none",
		"max_header_value_size_none",
		"max_parameter_count_none",
		"max_parameter_name_size_none",
		"max_parameter_value_size_none",
		"max_query_size_none",
		"max_request_line_size_none",
		"max_request_size_none",
		"max_url_size_none",
	},
	"ServicePolicyRule": {
		"any_client",
		"any_asn",
		"any_ip",
		"none",
		"default",
		"src_any",
		"dst_any",
		"intra_segment",
		"check_present",
		"check_not_present",
		"max_cookie_count_none",
		"max_cookie_key_size_none",
		"max_cookie_value_size_none",
		"max_header_count_none",
		"max_header_key_size_none",
		"max_header_value_size_none",
		"max_parameter_count_none",
		"max_parameter_name_size_none",
		"max_parameter_value_size_none",
		"max_query_size_none",
		"max_request_line_size_none",
		"max_request_size_none",
		"max_url_size_none",
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
