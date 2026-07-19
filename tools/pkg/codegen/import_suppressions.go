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
		// l7_ddos_protection nested server-default markers (#1155): the server materializes
		// all four for any l7_ddos_protection config (even empty), so a config that enables it
		// to set a custom rps_threshold / clientside action drifts on import without these.
		"mitigation_block",
		"default_rps_threshold",
		"clientside_action_none",
		"ddos_policy_none",
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
		// #1125: more_option empty-marker sub-blocks the API always materializes when a
		// config sets more_option (e.g. for header manipulation): custom_errors {} and
		// no_request_limit_per_connection {}. Same class as endpoint_subsets — the module
		// omits them, so suppress on import or the whole-LB round-trip drifts (+ each
		// marker). discover-defaults.go can't observe them (its probe LB sets no
		// more_option). Verified live (webapp-api-protection LPC-2 more_option matrix).
		"custom_errors",
		"no_request_limit_per_connection",
		// CR-1 (#1134): custom routes[] server-default empty markers. route_state_enabled is the
		// route enable/disable oneof default (routes are enabled by default); auto_host_rewrite is
		// the simple_route host-rewrite oneof base (server default when neither host_rewrite nor
		// disable_host_rewrite is set). A minimal simple_route omits both, so the whole-LB import
		// round-trip drifts (- each) without suppression. Same class as endpoint_subsets (#1103).
		// Both appear only under routes[]; matched by leaf name at any depth. Verified live
		// (webapp-api-protection CR-1).
		"route_state_enabled",
		"auto_host_rewrite",
		// CR-3 (#1138): simple_route.advanced_options oneof-base empty markers the API
		// materializes whenever advanced_options is set — buffer_policy (common_buffering),
		// hash_policy (common_hash_policy), retry (default_retry_policy), mirror
		// (disable_mirroring), rewrite (disable_prefix_rewrite, when no prefix_rewrite), spdy
		// (disable_spdy), web_socket (disable_web_socket_config), cluster-retract
		// (retract_cluster). A minimal advanced_options omits them, so the whole-LB import
		// round-trip drifts (- each). All appear only under routes[].simple_route.advanced_options.
		// Verified live (webapp-api-protection CR-3, two probes with/without prefix_rewrite+WAF).
		"common_buffering",
		"common_hash_policy",
		"default_retry_policy",
		"disable_mirroring",
		"disable_prefix_rewrite",
		"disable_spdy",
		"disable_web_socket_config",
		"retract_cluster",
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
		// LPC-5b (#1130): advanced_options oneof base members the API materializes whenever
		// advanced_options is set — the circuit_breaker / outlier_detection / subsets /
		// panic_threshold / request-limit "default/disable/no" choices plus auto_http_config.
		// A minimal advanced_options config (e.g. just connection_timeout/http_idle_timeout)
		// omits them, so import populates them and the next plan drifts (- each marker). Same
		// class as endpoint_subsets (#1103) / more_option (#1125). Verified live
		// (f5-sales-demo webapp-api-protection LPC-5b origin_pool matrix).
		"auto_http_config",
		"default_circuit_breaker",
		"disable_outlier_detection",
		"disable_subsets",
		"no_panic_threshold",
		"no_request_limit_per_connection",
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
	// SPol effort (service_policy coverage). Suppress ONLY the oneof base members the F5 XC
	// API echoes when the module OMITS their parent — verified live by GETting a created
	// service_policy (f5-sales-demo): a rule that declares no client/asn/ip matcher still
	// comes back with any_client{}/any_asn{}/any_ip{}, and a policy with no server scope
	// comes back with any_server{}. Those are true "server adds it on omit" defaults, so
	// suppressing them keeps a minimal config round-trip-import clean.
	//
	// Do NOT suppress DECLARED oneof members like waf_action.none / bot_action.none /
	// mum_action.default / segment_policy.{src_any,dst_any,intra_segment} /
	// headers[].{check_present,check_not_present} / request_constraints.max_*_none: the API
	// returns those ONLY when the parent block is set (waf_action=null / bot_action=null /
	// no segment_policy when omitted — confirmed live), so they are mandatory declared
	// values, not defaults. Suppressing them drops the declared value on import ("was
	// present, now absent" drift) — exactly how the live rule_list matrix caught an earlier
	// over-broad seed that guessed from schema structure instead of server behavior. Re-add
	// a specific base here only if a later sub-project's live matrix proves the server
	// echoes it alongside a concrete arm (the any_client pattern).
	"ServicePolicy": {
		"any_server",
		"any_client",
		"any_asn",
		"any_ip",
	},
	"ServicePolicyRule": {
		"any_client",
		"any_asn",
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
