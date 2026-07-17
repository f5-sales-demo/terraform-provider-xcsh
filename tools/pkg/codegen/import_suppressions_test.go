// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"reflect"
	"testing"
)

// The suppression data file carries a string "_comment" field alongside the
// []string resource entries. Parsing into map[string][]string fails on that
// string, which silently disabled the whole JSON (only the built-in seed stayed
// active). Regression test: _comment is skipped and resource entries load.
func TestParseSuppressionsJSON_SkipsComment(t *testing.T) {
	data := []byte(`{"_comment":"docs here","AppFirewall":["allow_all_response_codes","default_bot_setting"],"APIDefinition":["strict_schema_origin"]}`)
	got := parseSuppressionsJSON(data)
	want := map[string][]string{
		"AppFirewall":   {"allow_all_response_codes", "default_bot_setting"},
		"APIDefinition": {"strict_schema_origin"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseSuppressionsJSON() = %#v, want %#v", got, want)
	}
}

// End-to-end: an entry present only in the JSON data file (not the Go seed) must
// be honored, proving the file is actually loaded.
func TestImportSuppressions_JSONEntryHonored(t *testing.T) {
	if !isImportDefaultSuppressed("AppFirewall", "allow_all_response_codes") {
		t.Error("AppFirewall.allow_all_response_codes (JSON-only) should be suppressed; JSON not loaded?")
	}
}

// enable_api_discovery's inner oneof defaults are materialized by the server
// whenever discovery is enabled, but discover-defaults.go cannot observe them
// (its probe LB never enables discovery — it's the non-default side of the
// disable/enable oneof). Without suppression, a bare `enable_api_discovery {}`
// hard-errors on first apply and drifts on round-trip import. They are matched by
// leaf name at any depth, so listing them under HTTPLoadBalancer suffices.
func TestImportSuppressions_EnableAPIDiscoveryInnerDefaults(t *testing.T) {
	for _, m := range []string{"default_api_auth_discovery", "disable_learn_from_redirect_traffic"} {
		if !isImportDefaultSuppressed("HTTPLoadBalancer", m) {
			t.Errorf("HTTPLoadBalancer.%s must be a suppressed server-default (enable_api_discovery inner oneof)", m)
		}
	}
}

// disable_client_side_defense is the server default of the client_side_defense
// oneof on the HTTP load balancer. On any LB that does not enable CSD, the server
// materializes disable_client_side_defense {}, so it must be suppressed on import
// (same pattern as disable_waf / disable_api_discovery) to keep round-trip import
// clean.
func TestImportSuppressions_DisableClientSideDefense(t *testing.T) {
	if !isImportDefaultSuppressed("HTTPLoadBalancer", "disable_client_side_defense") {
		t.Error("HTTPLoadBalancer.disable_client_side_defense must be a suppressed server-default (client_side_defense oneof)")
	}
}

// #41 (SP3 API Protection): any_client is the server-default base member of the
// client_matcher oneof. Whenever a client_matcher is configured with a concrete arm
// (ip_prefix_list / ip_threat_category_list), the API normalizes the response to ALSO
// include any_client {} alongside it. A config that selects a concrete arm (and never
// declares any_client — the module omits client_matcher entirely for "match any") then
// drifts on round-trip import. skip_response_validation is the same class: the default
// member of the response-validation sub-oneof inside validation_mode, materialized by
// the API on every open_api_validation_rules entry. Both must be suppressed on import.
// Matched by leaf name at any depth. Verified live (f5-sales-demo webapp-api-protection
// SP3 matrix: variants 004 ip_prefix and 006 custom_list).
// #45 (SP4 API Testing): "standard" is the server-default base member of the
// api_testing credentials credentials_choice oneof — the API echoes standard {} on
// every credential. It appears on both the standalone xcsh_api_testing (APITesting)
// and the LB inline api_testing block (HTTPLoadBalancer), so both must suppress it on
// import (matched by leaf name at any depth), or a config that selects a concrete
// credential arm drifts on round-trip import. Verified live (f5-sales-demo).
func TestImportSuppressions_APITestingStandardMarker(t *testing.T) {
	for _, rc := range []string{"APITesting", "HTTPLoadBalancer"} {
		if !isImportDefaultSuppressed(rc, "standard") {
			t.Errorf("%s.standard must be a suppressed server-default (credentials_choice base marker)", rc)
		}
	}
}

func TestImportSuppressions_APIProtectionServerDefaults(t *testing.T) {
	// any_ip is the default member of the source-IP sub-oneof inside a client_matcher
	// ip_threat_category_list arm: the API echoes any_ip {} alongside the configured
	// ip_threat_categories, and the module never declares it (same class as any_client).
	for _, m := range []string{"any_client", "any_ip", "skip_response_validation"} {
		if !isImportDefaultSuppressed("HTTPLoadBalancer", m) {
			t.Errorf("HTTPLoadBalancer.%s must be a suppressed server-default (SP3 API Protection oneof base marker)", m)
		}
	}
}

// violations_view is the server-materialized catalog of WAF violation checks:
// whenever detection_settings is configured, the API populates the full
// violations_view list (name/title/description_spec/enabled/enabled_by_default)
// regardless of whether the config sets it. Without suppression, a
// detection_settings {} config drifts on round-trip import (the imported state
// carries dozens of server-populated violations_view blocks the config omits).
// Verified live against the f5-sales-demo tenant (webapp-api-protection WAF
// exhaustive-coverage matrix). Matched by leaf name at any depth.
func TestImportSuppressions_AppFirewallViolationsView(t *testing.T) {
	if !isImportDefaultSuppressed("AppFirewall", "violations_view") {
		t.Error("AppFirewall.violations_view must be a suppressed server-computed field (detection_settings round-trip import drift)")
	}
}

// #1103 (recurrence): empty-marker nested blocks the F5 XC API always echoes on
// every list element — origin_pool origin_servers[].labels {} and
// http_loadbalancer default_route_pools[].endpoint_subsets {} (also routes[].pools[],
// etc.). These are NOT oneof members but plain optional message blocks the server
// materializes as present-but-empty. On import there is no prior state to preserve, so
// without suppression the flatten populates the marker and the next plan wants to remove
// it (present-in-read vs absent-in-config) = perpetual drift, cascading into computed
// tenant re-planning on reference blocks. One entry per resource covers all depths
// (matched by leaf name). They slipped through because the auto-derive differ was blind
// to defaults nested inside list elements (fixed in tools/pkg/suppress/diff.go).
func TestImportSuppressions_EmptyMarkerListElementBlocks_Issue1103(t *testing.T) {
	if !isImportDefaultSuppressed("OriginPool", "labels") {
		t.Error("OriginPool.labels must be a suppressed server-default (origin_servers[].labels {} import round-trip drift, #1103)")
	}
	if !isImportDefaultSuppressed("HTTPLoadBalancer", "endpoint_subsets") {
		t.Error("HTTPLoadBalancer.endpoint_subsets must be a suppressed server-default (default_route_pools[].endpoint_subsets {} import round-trip drift, #1103)")
	}
}

// SPol-1 (service_policy coverage): suppress ONLY the oneof base members the F5 XC API
// echoes when their parent is OMITTED — verified live by GETting a created service_policy
// on f5-sales-demo: a rule with no client/asn/ip matcher still returns any_client{}/
// any_asn{}/any_ip{}, and a policy with no server scope returns any_server{}. These four
// are true "server adds it on omit" defaults, so suppressing them keeps a minimal config
// round-trip-import clean.
func TestImportSuppressions_ServicePolicyServerDefaults(t *testing.T) {
	for _, m := range []string{"any_server", "any_client", "any_asn", "any_ip"} {
		if !isImportDefaultSuppressed("ServicePolicy", m) {
			t.Errorf("ServicePolicy.%s must be a suppressed server-default", m)
		}
	}
	for _, m := range []string{"any_client", "any_asn", "any_ip"} {
		if !isImportDefaultSuppressed("ServicePolicyRule", m) {
			t.Errorf("ServicePolicyRule.%s must be a suppressed server-default", m)
		}
	}
	// any_server is policy-level only — the standalone rule has no server scope.
	if isImportDefaultSuppressed("ServicePolicyRule", "any_server") {
		t.Error("ServicePolicyRule.any_server must NOT be suppressed (policy-level only)")
	}

	// Over-suppression guard: DECLARED oneof members must NOT be suppressed. The API returns
	// these only when their parent block is set (waf_action.none / bot_action.none /
	// mum_action.default / segment_policy.{src_any,dst_any,intra_segment} /
	// headers[].{check_present,check_not_present} / request_constraints.max_*_none), so
	// suppressing them would drop the mandatory declared value on import — the exact drift
	// the live rule_list matrix caught. Also the user-intent arms/matchers.
	mustNotSuppress := []string{
		"none", "default", "src_any", "dst_any", "intra_segment",
		"check_present", "check_not_present", "max_cookie_count_none", "max_url_size_none",
		"max_request_size_none", "allow_list", "deny_list", "rule_list", "server_name",
		"server_selector", "server_name_matcher", "client_selector", "ip_prefix_list",
		"asn_list", "tls_fingerprint_matcher", "segment_policy", "waf_action",
	}
	for _, rc := range []string{"ServicePolicy", "ServicePolicyRule"} {
		for _, m := range mustNotSuppress {
			if isImportDefaultSuppressed(rc, m) {
				t.Errorf("%s.%s must NOT be suppressed (declared member / user intent — suppressing drops it on import)", rc, m)
			}
		}
	}
}

// Coverage Batch B (#51): a rate_limiter_policy rule that omits its country client
// matcher gets any_country {} materialized by the server (verified live on
// f5-sales-demo webapp-api-protection: a rule with asn_list but no country came
// back with any_country). any_asn/any_ip are the same server-default base members
// of the ASN/IP client-matcher oneofs. The module never declares them (it omits a
// matcher for "match any"), so they must be suppressed on import to keep the
// standalone xcsh_rate_limiter_policy round-trip clean.
func TestImportSuppressions_RateLimiterPolicyClientMatcherDefaults(t *testing.T) {
	for _, m := range []string{"any_asn", "any_country", "any_ip"} {
		if !isImportDefaultSuppressed("RateLimiterPolicy", m) {
			t.Errorf("RateLimiterPolicy.%s must be a suppressed server-default (rule client_matcher oneof)", m)
		}
	}
}

// #1125 (more_option): custom_errors {} and no_request_limit_per_connection {} are plain
// optional empty-marker sub-blocks the API always materializes when a config sets
// more_option. Same class as endpoint_subsets — suppress on import so the whole-LB round
// trip is clean when the module omits them.
func TestImportSuppressions_MoreOptionEmptyMarkers(t *testing.T) {
	for _, m := range []string{"custom_errors", "no_request_limit_per_connection"} {
		if !isImportDefaultSuppressed("HTTPLoadBalancer", m) {
			t.Errorf("HTTPLoadBalancer.%s must be a suppressed server-default (more_option empty marker)", m)
		}
	}
}
