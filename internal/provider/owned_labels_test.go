// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"reflect"
	"testing"
)

// F5 XC uses the ves.io/ and internal.ves.io/ namespaces for labels it sets itself, and
// filterSystemLabels strips every key in them from the read-back. Users set labels there
// too — real sites in this tenant carry `ves.io/app`, `ves.io/app_type`,
// `ves.io/mcn-demo` — and when they do the label becomes unmanageable: the configuration
// writes it, the read-back removes it from state, the next plan proposes adding it again,
// forever. Measured on cem1-l1 with released provider 3.81.1: two consecutive plans each
// reported `1 to change` for `ves.io/app = demo`, while the server held the label.
//
// #1391 tried to fix this by treating any platform-prefixed key already in prior state as
// owned. That failed: every provider up to 3.81.1 wrote the platform's OWN labels into
// state, so the guard kept those too and went on proposing their deletion. Prior state is
// not evidence of ownership.
//
// Ownership therefore has to be recorded where the configuration is actually visible —
// Create and Update — and read back in Read, which only ever sees prior state. These
// tests cover the encoding and the filter; #1398 tracks the live verification.

func TestOwnedLabelKeysRoundTrip(t *testing.T) {
	keys := []string{"ves.io/app", "env", "internal.ves.io/thing"}
	encoded, err := encodeOwnedLabelKeys(keys)
	if err != nil {
		t.Fatalf("encode: %v", err)
	}
	got := decodeOwnedLabelKeys(encoded)
	// Order must not matter to callers, but the encoding must be stable so an
	// unchanged configuration does not rewrite private state on every apply.
	want := map[string]struct{}{"ves.io/app": {}, "env": {}, "internal.ves.io/thing": {}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("round trip lost keys\ngot:  %v\nwant: %v", got, want)
	}
}

func TestEncodeOwnedLabelKeysIsStable(t *testing.T) {
	a, err := encodeOwnedLabelKeys([]string{"b", "a", "c"})
	if err != nil {
		t.Fatal(err)
	}
	b, err := encodeOwnedLabelKeys([]string{"c", "a", "b"})
	if err != nil {
		t.Fatal(err)
	}
	if string(a) != string(b) {
		t.Errorf("encoding depends on input order: %q vs %q", a, b)
	}
}

// The upgrade path. Every resource created before this change has no private entry, so
// there is nothing to decode. That must mean "owns no platform-namespaced label", which
// reproduces the #1391 behaviour exactly — the platform's labels stay filtered and the
// plan stays clean.
func TestDecodeOwnedLabelKeysTreatsAbsentAndMalformedAsOwningNothing(t *testing.T) {
	for name, raw := range map[string][]byte{
		"nil":       nil,
		"empty":     {},
		"not json":  []byte("ves.io/app"),
		"wrong sha": []byte(`{"unexpected":true}`),
	} {
		t.Run(name, func(t *testing.T) {
			if got := decodeOwnedLabelKeys(raw); len(got) != 0 {
				t.Errorf("expected to own nothing, got %v", got)
			}
		})
	}
}

func TestFilterSystemLabelsKeepsOwnedPlatformKeys(t *testing.T) {
	server := map[string]string{
		"env":              "demo",         // ordinary, always kept
		"ves.io/app":       "demo",         // declared by the configuration
		"ves.io/provider":  "ves-io-AZURE", // the platform's own, never declared
		"hw-serial-number": "ABC123",       // discovery label
	}
	owned := map[string]struct{}{"ves.io/app": {}}

	got := filterSystemLabelsOwning(server, true, owned)

	want := map[string]string{"env": "demo", "ves.io/app": "demo"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("owned platform key not preserved\ngot:  %v\nwant: %v", got, want)
	}
}

// The criterion that #1391 exists to protect: the platform's own label must never be
// proposed for deletion, whether or not an ownership entry exists.
func TestFilterSystemLabelsNeverKeepsUnownedPlatformKeys(t *testing.T) {
	server := map[string]string{"ves.io/provider": "ves-io-AZURE", "env": "demo"}

	for name, owned := range map[string]map[string]struct{}{
		"no entry at all":      nil,
		"entry owning nothing": {},
		"entry owning another": {"ves.io/app": {}},
	} {
		t.Run(name, func(t *testing.T) {
			got := filterSystemLabelsOwning(server, true, owned)
			if _, present := got["ves.io/provider"]; present {
				t.Error("ves.io/provider leaked into state and would be proposed for deletion")
			}
			if got["env"] != "demo" {
				t.Errorf("ordinary label lost: %v", got)
			}
		})
	}
}

// A configuration can deliberately take over a discovery label; #1391 already allows that
// on the write path via mergePreservedLabels, so the read path must agree.
func TestFilterSystemLabelsKeepsOwnedDiscoveryLabel(t *testing.T) {
	server := map[string]string{"hw-vendor": "mine", "hw-model": "theirs"}
	got := filterSystemLabelsOwning(server, true, map[string]struct{}{"hw-vendor": {}})
	if got["hw-vendor"] != "mine" {
		t.Errorf("owned discovery label not preserved: %v", got)
	}
	if _, present := got["hw-model"]; present {
		t.Errorf("unowned discovery label leaked: %v", got)
	}
}

// filterSystemLabels must keep behaving exactly as before, so the existing generated
// callers and the #1391 regression tests are unaffected.
func TestFilterSystemLabelsUnchangedWithoutOwnership(t *testing.T) {
	server := map[string]string{"env": "demo", "ves.io/app": "demo", "hw-model": "x"}
	if !reflect.DeepEqual(filterSystemLabels(server, true), filterSystemLabelsOwning(server, true, nil)) {
		t.Error("filterSystemLabels and filterSystemLabelsOwning(nil) disagree")
	}
}
