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
