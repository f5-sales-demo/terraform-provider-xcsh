// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"reflect"
	"testing"
)

// keys is a tiny helper to build the prior-state key set the filter consults.
func keys(names ...string) map[string]struct{} {
	out := make(map[string]struct{}, len(names))
	for _, n := range names {
		out[n] = struct{}{}
	}
	return out
}

// #1391: F5 XC populates a fixed set of top-level metadata labels on a site from
// the node's own hardware and OS discovery, once the node registers. They are not
// prefixed `ves.io/`, so the original filter let them through into state, where an
// empty `labels` configuration then proposed deleting every one of them on every
// plan — and `terraform plan` stopped being usable as a "did I change anything"
// signal on any stack owning a site.
func TestFilterSystemLabels_DropsDiscoveredSiteLabels_Issue1391(t *testing.T) {
	// Exactly the six keys observed on every CE site in the tenant, across all
	// four providers (AWS, Azure, VMware, KVM), plus one genuine user label.
	in := map[string]string{
		"domain":           "us-east-2.compute.internal",
		"host-os-version":  "rhel-9-2024-6",
		"hw-model":         "virtual-machine",
		"hw-serial-number": "0000-0011-9249-1643-8520-4494-21",
		"hw-vendor":        "microsoft-corporation",
		"hw-version":       "7-0",
		"env":              "demo",
	}
	want := map[string]string{"env": "demo"}

	got := filterSystemLabels(in, nil)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("filterSystemLabels() = %v, want %v", got, want)
	}
}

// The platform reserves two label namespaces, not one. `internal.ves.io/` does not
// start with `ves.io/`, so a prefix test for the latter alone lets it through.
func TestFilterSystemLabels_DropsBothPlatformPrefixes(t *testing.T) {
	in := map[string]string{
		"ves.io/provider":         "ves-io-AZURE",
		"ves.io/siteType":         "ves-io-ce",
		"internal.ves.io/batch":   "batch-3",
		"internal.ves.io/griffin": "griffin-2",
		"internal.ves.io/network": "v2",
		"location":                "lab",
	}
	want := map[string]string{"location": "lab"}

	got := filterSystemLabels(in, nil)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("filterSystemLabels() = %v, want %v", got, want)
	}
}

// A platform-owned key that the configuration genuinely declares must still be
// reconciled, or the fix is blanket suppression: the user could set it and never
// see it change. Prior state is the signal — after apply it holds what the
// configuration asked for, so a key present there is one Terraform manages.
func TestFilterSystemLabels_KeepsPlatformKeysTheConfigurationOwns(t *testing.T) {
	in := map[string]string{
		"domain":           "lab.example.com", // colliding name, but ours
		"hw-serial-number": "0000-0011-9249",  // discovered, not ours
		"ves.io/app":       "ms-build",        // reserved prefix, but ours
		"ves.io/provider":  "ves-io-AZURE",    // platform's
	}
	prior := keys("domain", "ves.io/app")
	want := map[string]string{"domain": "lab.example.com", "ves.io/app": "ms-build"}

	got := filterSystemLabels(in, prior)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("filterSystemLabels() = %v, want %v", got, want)
	}
}

// An out-of-band deletion of a label the configuration owns must still surface as
// drift: the key is absent from the response, so it must be absent from the result
// even though prior state still lists it.
func TestFilterSystemLabels_ReportsOutOfBandDeletionOfAnOwnedKey(t *testing.T) {
	in := map[string]string{"hw-model": "virtual-machine"}
	prior := keys("env", "hw-model")
	want := map[string]string{"hw-model": "virtual-machine"}

	got := filterSystemLabels(in, prior)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("filterSystemLabels() = %v, want %v", got, want)
	}
}

func TestFilterSystemLabels_EmptyAndNil(t *testing.T) {
	if got := filterSystemLabels(nil, nil); len(got) != 0 {
		t.Errorf("filterSystemLabels(nil, nil) = %v, want empty", got)
	}
	if got := filterSystemLabels(map[string]string{}, nil); len(got) != 0 {
		t.Errorf("filterSystemLabels(empty, nil) = %v, want empty", got)
	}
	// A site that carries nothing but platform labels filters down to nothing,
	// which is what lets an empty `labels` configuration stay null.
	onlyPlatform := map[string]string{"ves.io/provider": "ves-io-AZURE", "hw-vendor": "amazon-ec2"}
	if got := filterSystemLabels(onlyPlatform, nil); len(got) != 0 {
		t.Errorf("filterSystemLabels(onlyPlatform, nil) = %v, want empty", got)
	}
}
