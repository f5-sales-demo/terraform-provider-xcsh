// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// provider_helpers.go - Manually maintained helper functions for the provider.
// This file is NOT auto-generated and contains utility functions used by
// the provider implementation.

package provider

import (
	"strings"
)

// normalizeAPIURL cleans up the API URL to ensure consistent formatting.
// It removes trailing slashes and the /api suffix if present, since API paths
// already include the /api prefix (e.g., /api/web/namespaces).
func normalizeAPIURL(url string) (string, bool) {
	original := url

	// Remove trailing slashes
	url = strings.TrimRight(url, "/")

	// Remove /api suffix (case-insensitive check, preserve original case in removal)
	if strings.HasSuffix(strings.ToLower(url), "/api") {
		url = url[:len(url)-4]
	}

	// Remove any trailing slashes that might have been before /api
	url = strings.TrimRight(url, "/")

	return url, url != original
}

// filterSystemLabels removes F5 XC system-managed labels (ves.io/*) from the label map.
// These labels are injected by the platform and should not be managed by Terraform.
// nolint:unused // Used by generated resource/data source Read methods
func filterSystemLabels(labels map[string]string) map[string]string {
	filtered := make(map[string]string)
	for k, v := range labels {
		if !strings.HasPrefix(k, "ves.io/") {
			filtered[k] = v
		}
	}
	return filtered
}

// systemManagedRrSetGroupName is the reserved name F5 XC gives the rr_set_group it
// auto-creates in a DNS zone to hold records owned by load balancers (created when
// a primary block sets allow_http_lb_managed_records = true). The group is
// platform-owned: any attempt to modify or delete it via the config API returns
// 403 FORBIDDEN.
const systemManagedRrSetGroupName = "x-ves-io-managed"

// isSystemManagedRrSetGroup reports whether an rr_set_group element read from the
// API (as decoded into a map) is the reserved F5 XC system-managed group. The
// reserved name is the authoritative signal.
func isSystemManagedRrSetGroup(item map[string]interface{}) bool {
	md, ok := item["metadata"].(map[string]interface{})
	if !ok {
		return false
	}
	name, _ := md["name"].(string)
	return name == systemManagedRrSetGroupName
}

// filterSystemManagedRrSetGroups returns rawList with the F5 XC system-managed
// rr_set_group ("x-ves-io-managed") removed. Terraform must not surface that group
// as user-managed state: a config that does not declare it would otherwise plan a
// delete the API forbids (403). Filtering the raw API list up front (before the
// flatten loop) keeps prior-state positional threading aligned, since the user
// never declares the system group. Returns a new slice; the input is not mutated.
// nolint:unused // Used by generated DNS zone Read/Create/Update methods
func filterSystemManagedRrSetGroups(rawList []interface{}) []interface{} {
	filtered := make([]interface{}, 0, len(rawList))
	for _, item := range rawList {
		if m, ok := item.(map[string]interface{}); ok && isSystemManagedRrSetGroup(m) {
			continue
		}
		filtered = append(filtered, item)
	}
	return filtered
}
