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
