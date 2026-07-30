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
