// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

// importDefaultSuppressions lists, per resource (title-case model prefix), the
// empty-marker oneof members that the F5 XC API ALWAYS returns as the server
// default for their oneof group. On `terraform import` there is no prior config
// to preserve, so the flatten would otherwise populate every such default member
// and the next plan would show spurious drift (it wants to remove blocks the user
// never configured).
//
// Suppressing the DEFAULT member of a oneof on import is semantically safe:
// omitting it means the server re-applies the same default, so behavior is
// unchanged. Non-default members (e.g. app_firewall vs disable_waf) are NOT
// listed and still import normally, as are user-intent markers such as
// advertise_on_public_default_vip.
//
// Data source: discovered from the live tenant (create a minimal object, diff the
// response against the request). Seeded here for http_loadbalancer from the
// staging tenant; see tracking issue #1006 for auto-populating this from the
// discover-defaults pipeline across all resources.
var importDefaultSuppressions = map[string]map[string]bool{
	// Healthcheck: http_health_check.headers is a server-default empty marker
	// (verified via live API — returned as {} for a minimal health check).
	"Healthcheck": {
		"headers": true,
	},
	"HTTPLoadBalancer": {
		"default_sensitive_data_policy":    true,
		"disable_api_definition":           true,
		"disable_api_discovery":            true,
		"disable_api_testing":              true,
		"disable_malicious_user_detection": true,
		"disable_malware_protection":       true,
		"disable_rate_limit":               true,
		"disable_threat_mesh":              true,
		"disable_trust_client_ip_headers":  true,
		"disable_waf":                      true,
		"l7_ddos_protection":               true,
		"no_challenge":                     true,
		"round_robin":                      true,
		"service_policies_from_namespace":  true,
		"user_id_client_ip":                true,
	},
}

// isImportDefaultSuppressed reports whether the given empty-marker member of the
// given resource is a server-default oneof member that must not be populated from
// the API response on the import path.
func isImportDefaultSuppressed(resourceTitleCase, jsonName string) bool {
	members, ok := importDefaultSuppressions[resourceTitleCase]
	if !ok {
		return false
	}
	return members[jsonName]
}
