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

// platformLabelPrefixes are the two label namespaces F5 XC reserves for labels it
// sets itself. `internal.ves.io/` does not start with `ves.io/`, so one prefix test
// does not cover both.
var platformLabelPrefixes = []string{"ves.io/", "internal.ves.io/"}

// discoveredSiteLabels are the labels F5 XC writes onto a node-backed site from the
// hardware and OS it booted on, once the node registers. They carry no reserved
// prefix, which is why a `ves.io/` test alone let them through into state, where a
// configuration with no `labels` block then proposed deleting every one of them on
// every plan (#1391). A user cannot meaningfully own `hw-serial-number` on a site,
// and an empty `labels` block plainly does not mean "delete the serial number".
//
// The set is fixed and platform-defined: measured in this tenant, every decorated
// site carries exactly these six keys on all four infrastructure providers.
var discoveredSiteLabels = map[string]struct{}{
	"domain":           {},
	"host-os-version":  {},
	"hw-model":         {},
	"hw-serial-number": {},
	"hw-vendor":        {},
	"hw-version":       {},
}

// filterSystemLabels returns labels with the platform's own entries removed, so that
// Terraform state holds only what a configuration could have authored.
//
// siteDiscovery additionally removes the six hardware/OS discovery labels. It is set
// only for the resources F5 XC actually decorates (tools/discovered-site-labels.json),
// because `domain` is an ordinary enough name that a user could legitimately own a
// label called that on an unrelated object.
//
// Filtering is unconditional — deliberately not conditioned on the prior state's
// keys. Prior state looks like evidence of what a configuration owns, but it is not:
// an earlier provider version wrote these very keys into state, so honouring state
// would keep every existing user on the broken behaviour after upgrading, which is
// the population that has the problem. Filtering by key and never by value means an
// out-of-band edit to a label a configuration owns still surfaces as drift.
//
// The cost is that a `ves.io/`-prefixed or discovery-named label a configuration
// genuinely sets cannot be managed on a decorated resource: the read-back removes it,
// so the plan proposes adding it forever. That is pre-existing for the prefixes and is
// tracked separately — fixing it needs recorded ownership that legacy state cannot
// forge.
//
// nolint:unused // Used by generated resource/data source Read methods
func filterSystemLabels(labels map[string]string, siteDiscovery bool) map[string]string {
	filtered := make(map[string]string, len(labels))
	for k, v := range labels {
		if !isPlatformLabel(k, siteDiscovery) {
			filtered[k] = v
		}
	}
	return filtered
}

// isPlatformLabel reports whether F5 XC, rather than a user, is the author of the
// label key.
func isPlatformLabel(key string, siteDiscovery bool) bool {
	if siteDiscovery {
		if _, ok := discoveredSiteLabels[key]; ok {
			return true
		}
	}
	for _, prefix := range platformLabelPrefixes {
		if strings.HasPrefix(key, prefix) {
			return true
		}
	}
	return false
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
