// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"encoding/json"
	"strings"
	"testing"
)

// Use the declared bundled artifact itself. Canonical regeneration and native
// qualification exercise the reviewed v12.0.9 bytes; pending source pins never
// substitute for this binary's actual artifact identity.
func TestNetworkAllowlistReviewedArtifactRejectsChanges(t *testing.T) {
	raw := bundledNetworkAllowlistJSON
	if err := reviewedNetworkAllowlist([]byte(raw)); err != nil {
		t.Fatal(err)
	}
	literal := func(value any) string {
		encoded, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		return string(encoded)
	}
	botDomains := bundledNetworkAllowlist.Services["bot_defense"].Domains
	if len(botDomains) == 0 {
		t.Fatal("owned bot domains are missing")
	}
	domains := literal(botDomains)
	cases := []struct{ name, before, after string }{
		{"unknown_release", literal(bundledNetworkAllowlist.APIReleaseTag), literal("v0.0.0")},
		{"source_digest", literal(bundledNetworkAllowlist.SourceSHA256), literal(strings.Repeat("0", 64))},
		{"source_url", literal(bundledNetworkAllowlist.SourceURL), literal("https://example.invalid/manifest.json")},
		{"generated_at", literal(bundledNetworkAllowlist.GeneratedAt), literal("2000-01-01T00:00:00Z")},
		{"added_domain", domains, strings.TrimSuffix(domains, "]") + `,"unreviewed.example.invalid"]`},
		{"existing_domain", literal(botDomains[0]), literal("changed.example.invalid")},
		{"ip_group", literal(bundledNetworkAllowlist.Services["global_log_receiver"].SourceEntries[2]), literal("192.0.2.1")},
		{"customer_edge", literal(bundledNetworkAllowlist.CustomerEdge.SecureMeshV2.Domains[0]), literal(".example.invalid")},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if strings.Count(raw, tc.before) != 1 {
				t.Fatal("mutation must target exactly one owned artifact value")
			}
			changed := strings.Replace(raw, tc.before, tc.after, 1)
			if !json.Valid([]byte(changed)) {
				t.Fatal("mutation must retain valid artifact JSON")
			}
			if reviewedNetworkAllowlist([]byte(changed)) == nil {
				t.Fatal("accepted an unreviewed artifact mutation")
			}
		})
	}
}
