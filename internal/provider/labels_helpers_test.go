// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"reflect"
	"testing"
)

// The seven labels a real Customer Edge site carries: the six F5 XC discovers from
// the node's hardware and OS, plus the one platform-prefixed key it injects.
func decoratedSiteLabels() map[string]string {
	return map[string]string{
		"domain":           "",
		"host-os-version":  "rhel-9-2024-6",
		"hw-model":         "virtual-machine",
		"hw-serial-number": "0000-0011-9249-1643-8520-4494-21",
		"hw-vendor":        "microsoft-corporation",
		"hw-version":       "7-0",
		"ves.io/provider":  "ves-io-AZURE",
	}
}

// #1391: the discovery labels carry no reserved prefix, so a `ves.io/` test alone let
// them through into state, where a configuration with no `labels` block then proposed
// deleting every one of them on every plan — and `terraform plan` stopped being usable
// as a "did I change anything" signal on any stack owning a site.
func TestFilterSystemLabels_DropsDiscoveredSiteLabels_Issue1391(t *testing.T) {
	in := decoratedSiteLabels()
	in["env"] = "demo" // one genuine user label, which must survive

	got := filterSystemLabels(in, true)
	want := map[string]string{"env": "demo"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("filterSystemLabels(siteDiscovery=true) = %v, want %v", got, want)
	}
}

// The filter must be unconditional, not conditioned on the prior state's keys. Prior
// state looks like evidence of what a configuration owns, but an earlier provider
// version wrote these very keys into state — so honouring it would leave every
// existing user on the broken behaviour after upgrading, which is exactly the
// population that has the problem. Reproduced on a live site before this test existed:
// with state written by 3.81.1, an ownership-guarded filter still planned
// "- labels = {...} -> null".
func TestFilterSystemLabels_LegacyStateDoesNotResurrectTheBug_Issue1391(t *testing.T) {
	// Whatever a prior provider left in state, the result depends only on the response.
	got := filterSystemLabels(decoratedSiteLabels(), true)
	if len(got) != 0 {
		t.Errorf("a decorated site must filter down to nothing, got %v", got)
	}
}

// The six names are ordinary enough that a user could own one on an object F5 XC does
// not decorate, so the discovery filter is opt-in per resource. Off, only the prefixes
// are filtered.
func TestFilterSystemLabels_DiscoveryFilterIsScopedToDecoratedResources(t *testing.T) {
	in := map[string]string{
		"domain":          "example.com",
		"hw-model":        "a name a user chose",
		"ves.io/provider": "ves-io-AZURE",
	}

	got := filterSystemLabels(in, false)
	want := map[string]string{"domain": "example.com", "hw-model": "a name a user chose"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("filterSystemLabels(siteDiscovery=false) = %v, want %v", got, want)
	}
}

// The platform reserves two label namespaces, not one. `internal.ves.io/` does not
// start with `ves.io/`, so a prefix test for the latter alone lets it through — on
// every resource, decorated or not.
func TestFilterSystemLabels_DropsBothPlatformPrefixes(t *testing.T) {
	in := map[string]string{
		"ves.io/siteType":         "ves-io-ce",
		"internal.ves.io/batch":   "batch-3",
		"internal.ves.io/griffin": "griffin-2",
		"location":                "lab",
	}
	want := map[string]string{"location": "lab"}

	for _, siteDiscovery := range []bool{false, true} {
		got := filterSystemLabels(in, siteDiscovery)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("filterSystemLabels(siteDiscovery=%v) = %v, want %v", siteDiscovery, got, want)
		}
	}
}

// Filtering is by key and never by value, so a label a configuration owns still
// reconciles: whatever the response says it is, is what state gets.
func TestFilterSystemLabels_PassesOwnedLabelsThroughByValue(t *testing.T) {
	in := decoratedSiteLabels()
	in["env"] = "matrix-2" // changed out from under us

	got := filterSystemLabels(in, true)
	if got["env"] != "matrix-2" {
		t.Errorf("owned label must reflect the response value, got %q", got["env"])
	}
}

func TestFilterSystemLabels_EmptyAndNil(t *testing.T) {
	for _, siteDiscovery := range []bool{false, true} {
		if got := filterSystemLabels(nil, siteDiscovery); len(got) != 0 {
			t.Errorf("filterSystemLabels(nil, %v) = %v, want empty", siteDiscovery, got)
		}
		if got := filterSystemLabels(map[string]string{}, siteDiscovery); len(got) != 0 {
			t.Errorf("filterSystemLabels(empty, %v) = %v, want empty", siteDiscovery, got)
		}
	}
}

// #1396: F5 XC replaces metadata.labels on write rather than merging, so a PUT built
// only from the configuration erases the discovery labels. Measured live: a site
// carrying all six, applied with labels = { env, tier }, came back holding only
// { env, tier, ves.io/provider }. Read stashes what it filtered so Update can send it
// back.
func TestPreservedPlatformLabels_KeepsOnlyWhatAWriteWouldDestroy(t *testing.T) {
	in := decoratedSiteLabels()
	in["env"] = "demo"

	got := preservedPlatformLabels(in)
	want := map[string]string{
		"domain":           "",
		"host-os-version":  "rhel-9-2024-6",
		"hw-model":         "virtual-machine",
		"hw-serial-number": "0000-0011-9249-1643-8520-4494-21",
		"hw-vendor":        "microsoft-corporation",
		"hw-version":       "7-0",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("preservedPlatformLabels() = %v, want %v", got, want)
	}
	// ves.io/provider is deliberately absent: the platform re-injects it after every
	// write that drops it, and sending a reserved-namespace label back invites a
	// rejection for no gain.
	if _, ok := got["ves.io/provider"]; ok {
		t.Error("prefixed platform labels must not be preserved for write-back")
	}
	// The user's own label is not the write path's business — the configuration carries it.
	if _, ok := got["env"]; ok {
		t.Error("a user label must not be preserved via the platform stash")
	}
}

func TestMergePreservedLabels_ConfigurationWinsOnACollision(t *testing.T) {
	outgoing := map[string]string{"env": "demo", "hw-model": "the operator overrode this"}
	preserved := map[string]string{"hw-model": "virtual-machine", "hw-vendor": "microsoft-corporation"}

	got := mergePreservedLabels(outgoing, preserved)
	want := map[string]string{
		"env":       "demo",
		"hw-model":  "the operator overrode this",
		"hw-vendor": "microsoft-corporation",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("mergePreservedLabels() = %v, want %v", got, want)
	}
}

// The case that actually erased the fleet's labels: a configuration with no labels at
// all. The outgoing map is nil, and the merge still has to produce the preserved six.
func TestMergePreservedLabels_RestoresWhenTheConfigurationSetsNothing(t *testing.T) {
	preserved := map[string]string{"hw-model": "virtual-machine", "hw-vendor": "microsoft-corporation"}

	got := mergePreservedLabels(nil, preserved)
	if !reflect.DeepEqual(got, preserved) {
		t.Errorf("mergePreservedLabels(nil, preserved) = %v, want %v", got, preserved)
	}
}

// Nothing stashed must not conjure an empty map where the caller had nil, or a site
// that genuinely has no labels would start being written with one.
func TestMergePreservedLabels_NothingStashedIsAPassThrough(t *testing.T) {
	if got := mergePreservedLabels(nil, nil); got != nil {
		t.Errorf("mergePreservedLabels(nil, nil) = %v, want nil", got)
	}
	outgoing := map[string]string{"env": "demo"}
	if got := mergePreservedLabels(outgoing, map[string]string{}); !reflect.DeepEqual(got, outgoing) {
		t.Errorf("mergePreservedLabels(outgoing, empty) = %v, want %v", got, outgoing)
	}
}

// ValidateConfig cannot inspect a labels map that is unknown at validate time — labels =
// var.x, or a value derived from another resource — so Create and Update run this against
// the resolved map as a backstop. Sorted, so the error names the keys deterministically.
func TestDiscoveredSiteLabelKeys_ReportsOffendersSorted(t *testing.T) {
	in := map[string]string{
		"hw-vendor": "microsoft-corporation",
		"env":       "demo",
		"domain":    "example.com",
		"hw-model":  "virtual-machine",
	}
	got := discoveredSiteLabelKeys(in)
	want := []string{"domain", "hw-model", "hw-vendor"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("discoveredSiteLabelKeys() = %v, want %v", got, want)
	}
}

func TestDiscoveredSiteLabelKeys_SilentOnACleanMap(t *testing.T) {
	if got := discoveredSiteLabelKeys(map[string]string{"env": "demo"}); got != nil {
		t.Errorf("discoveredSiteLabelKeys(clean) = %v, want nil", got)
	}
	if got := discoveredSiteLabelKeys(nil); got != nil {
		t.Errorf("discoveredSiteLabelKeys(nil) = %v, want nil", got)
	}
}
