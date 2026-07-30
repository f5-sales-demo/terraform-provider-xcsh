// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// provider_helpers.go - Manually maintained helper functions for the provider.
// This file is NOT auto-generated and contains utility functions used by
// the provider implementation.

package provider

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
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

// discoveredSiteLabels are the top-level metadata labels F5 XC writes onto a site
// from the node's own hardware and OS discovery once it registers. They carry no
// reserved prefix, which is why a `ves.io/` test alone let them through (#1391).
//
// The set is fixed and platform-defined, not per-site: every Customer Edge in the
// tenant carries exactly these six keys, on all four infrastructure providers
// (AWS, Azure, VMware, KVM). A user cannot meaningfully own `hw-serial-number` in
// configuration, and an empty `labels` block plainly does not mean "delete the
// serial number" — but that is what Terraform proposed on every plan, forever,
// until this filter covered them.
var discoveredSiteLabels = map[string]struct{}{
	"domain":           {},
	"host-os-version":  {},
	"hw-model":         {},
	"hw-serial-number": {},
	"hw-vendor":        {},
	"hw-version":       {},
}

// filterSystemLabels returns labels with the platform's own entries removed, so
// that Terraform state holds only what the configuration manages.
//
// priorKeys is the key set of the prior state's labels — that is, what the last
// apply put there on the configuration's behalf. A platform-owned key listed in
// priorKeys is kept: the configuration genuinely declares it, so it must keep
// reconciling. Without that guard the fix would be blanket suppression, and a user
// who set `ves.io/app` (which real sites in this tenant do) could never see it
// change. Filtering is by key only, never by value, so an out-of-band edit to a
// label the configuration owns still surfaces as drift.
//
// nolint:unused // Used by generated resource/data source Read methods
func filterSystemLabels(labels map[string]string, priorKeys map[string]struct{}) map[string]string {
	filtered := make(map[string]string, len(labels))
	for k, v := range labels {
		if _, owned := priorKeys[k]; owned || !isPlatformLabel(k) {
			filtered[k] = v
		}
	}
	return filtered
}

// labelKeySet returns the key set of a prior-state label map, for the ownership
// guard in filterSystemLabels. A null or unknown map yields nil: nothing is owned
// yet, which is the right answer both on import and in a data source, neither of
// which has a prior state to consult.
//
// nolint:unused // Used by generated resource/data source Read methods
func labelKeySet(prior types.Map) map[string]struct{} {
	if prior.IsNull() || prior.IsUnknown() {
		return nil
	}
	elements := prior.Elements()
	out := make(map[string]struct{}, len(elements))
	for k := range elements {
		out[k] = struct{}{}
	}
	return out
}

// isPlatformLabel reports whether F5 XC, rather than a user, is the author of the
// label key.
func isPlatformLabel(key string) bool {
	if _, ok := discoveredSiteLabels[key]; ok {
		return true
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
