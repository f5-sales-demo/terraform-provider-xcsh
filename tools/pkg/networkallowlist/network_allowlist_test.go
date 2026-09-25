// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package networkallowlist

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"
	"testing"
)

func TestExtractPinnedV802Manifest(t *testing.T) {
	openAPI, pin := fixture(t, validManifest())
	artifact, err := Extract(openAPI, pin, "v8.0.2")
	if err != nil {
		t.Fatalf("Extract() error = %v", err)
	}
	if artifact.APIReleaseTag != "v8.0.2" ||
		artifact.SourceSHA256 == "" ||
		artifact.GeneratedAt != "2025-09-10T00:00:00Z" {
		t.Fatalf("unexpected provenance: %#v", artifact)
	}
	if got := artifact.Services["global_log_receiver"].SourceEntries; got[2] != "3.64.239.6" {
		t.Fatalf("mixed global log receiver source entries = %v", got)
	}
	if got := artifact.Services["data_intelligence"].Regions["us"].SourceEntries[0]; got != "35.247.100.206" {
		t.Fatalf("US Data Intelligence source entry = %q", got)
	}
	if got := artifact.CustomerEdge.SecureMeshV2.Domains[0]; got != ".volterra.io" {
		t.Fatalf("leading-dot domain = %q", got)
	}
	encoded, err := json.Marshal(artifact)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), "legacy") || strings.Contains(string(encoded), "20.33.0.0/16") {
		t.Fatalf("generated artifact retained legacy Customer Edge data: %s", encoded)
	}
}

func TestExtractRejectsInvalidExtension(t *testing.T) {
	tests := []struct {
		name          string
		mutate        func(map[string]any)
		match         string
		keepBadDigest bool
	}{
		{"missing extension", func(root map[string]any) { delete(root["info"].(map[string]any), "x-f5xc-network-allowlist") }, "missing", false},
		{"digest mismatch", func(root map[string]any) {
			root["info"].(map[string]any)["x-f5xc-network-allowlist"].(map[string]any)["sha256"] = strings.Repeat("0", 64)
		}, "digest", true},
		{"missing service", func(root map[string]any) {
			delete(root["info"].(map[string]any)["x-f5xc-network-allowlist"].(map[string]any)["manifest"].(map[string]any)["services"].(map[string]any), "cdn")
		}, "cdn", false},
		{"unknown service", func(root map[string]any) {
			root["info"].(map[string]any)["x-f5xc-network-allowlist"].(map[string]any)["manifest"].(map[string]any)["services"].(map[string]any)["surprise"] = map[string]any{"ipv4_ips": []any{"192.0.2.1"}}
		}, "surprise", false},
		{"malformed address", func(root map[string]any) {
			root["info"].(map[string]any)["x-f5xc-network-allowlist"].(map[string]any)["manifest"].(map[string]any)["services"].(map[string]any)["dnslb_health_checks"] = map[string]any{"ipv4_ips": []any{"not-an-ip"}}
		}, "not-an-ip", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			openAPI, pin := fixture(t, validManifest())
			var root map[string]any
			if err := json.Unmarshal(openAPI, &root); err != nil {
				t.Fatal(err)
			}
			test.mutate(root)
			if !test.keepBadDigest {
				refreshExtensionDigest(t, root)
			}
			openAPI, err := json.Marshal(root)
			if err != nil {
				t.Fatal(err)
			}
			pin = releasePin(t, openAPI)
			if _, err := Extract(openAPI, pin, "v8.0.2"); err == nil || !strings.Contains(err.Error(), test.match) {
				t.Fatalf("Extract() error = %v, want containing %q", err, test.match)
			}
		})
	}
}

func refreshExtensionDigest(t *testing.T, root map[string]any) {
	t.Helper()
	info := root["info"].(map[string]any)
	raw, ok := info["x-f5xc-network-allowlist"]
	if !ok {
		return
	}
	extension := raw.(map[string]any)
	canonical, err := json.Marshal(extension["manifest"])
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(canonical)
	extension["sha256"] = hex.EncodeToString(sum[:])
}

func fixture(t *testing.T, manifest map[string]any) ([]byte, []byte) {
	t.Helper()
	canonical, err := json.Marshal(manifest)
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(canonical)
	root := map[string]any{"info": map[string]any{"x-f5xc-network-allowlist": map[string]any{
		"source_url": "https://docs.cloud.f5.com/ips-domains.json",
		"sha256":     hex.EncodeToString(sum[:]),
		"manifest":   manifest,
	}}}
	openAPI, err := json.Marshal(root)
	if err != nil {
		t.Fatal(err)
	}
	return openAPI, releasePin(t, openAPI)
}

func releasePin(t *testing.T, openAPI []byte) []byte {
	t.Helper()
	sum := sha256.Sum256(openAPI)
	pin, err := json.Marshal(map[string]any{
		"release_tag": "v8.0.2",
		"assets":      map[string]string{"openapi.json": "sha256:" + hex.EncodeToString(sum[:])},
	})
	if err != nil {
		t.Fatal(err)
	}
	return pin
}

func validManifest() map[string]any {
	return map[string]any{
		"generated_at":     "2025-09-10T00:00:00Z",
		"manifest_type":    "firewall_proxy_allowlist",
		"manifest_version": "1.0.0",
		"schema_version":   "1.0.0",
		"reference":        "https://docs.cloud.f5.com/network-cloud-ref",
		"services": map[string]any{
			"regional_edges":               map[string]any{"regions": map[string]any{"americas": map[string]any{"ipv4_cidrs": []any{"192.0.2.0/24"}}, "europe": map[string]any{"ipv4_cidrs": []any{"198.51.100.0/24"}}, "asia": map[string]any{"ipv4_cidrs": []any{"203.0.113.0/24"}}}},
			"cdn":                          map[string]any{"ipv4_cidrs": []any{"192.0.2.0/24"}},
			"secondary_dns_zone_transfer":  map[string]any{"ipv4_ips": []any{"192.0.2.53"}},
			"global_log_receiver":          map[string]any{"ipv4_cidrs": []any{"198.51.100.0/29", "203.0.113.8", "3.64.239.6"}},
			"dnslb_health_checks":          map[string]any{"ipv4_ips": []any{"192.0.2.80"}},
			"global_controller_sso_egress": map[string]any{"ipv4_ips": []any{"192.0.2.81"}},
			"bot_defense":                  map[string]any{"domains": []any{"ibd.example.test"}},
			"data_intelligence":            map[string]any{"regions": map[string]any{"us": map[string]any{"ipv4_ips": []any{"35.247.100.206"}}, "eu": map[string]any{"ipv4_cidrs": []any{"198.51.100.4/32"}}}},
		},
		"customer_edge": map[string]any{
			"defaults": map[string]any{"dns_servers": []any{"8.8.8.8"}, "ntp_servers": []any{"216.239.35.4"}},
			"site_types": map[string]any{
				"secure_mesh_v2": map[string]any{"egress_domain_rules": map[string]any{"domains": []any{".volterra.io"}}, "egress_ip_rules": map[string]any{"registration_updates": map[string]any{"ipv4_ips": []any{"159.60.141.140"}}}},
				"legacy":         map[string]any{"egress_ip_rules": map[string]any{"registration_updates": map[string]any{"ipv4_cidrs": []any{"20.33.0.0/16"}}}},
			},
		},
	}
}
