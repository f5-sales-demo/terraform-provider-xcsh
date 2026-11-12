package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestProviderDeliveryInventory(t *testing.T) {
	for _, test := range []struct {
		name      string
		extras    []string
		badDigest bool
		wantError bool
	}{
		{name: "legacy"},
		{name: "complete_pair", extras: []string{"gcp-operational-health-v2.json", "gcp-operational-qualification-v2.json"}},
		{name: "missing_health", extras: []string{"gcp-operational-qualification-v2.json"}, wantError: true},
		{name: "missing_qualification", extras: []string{"gcp-operational-health-v2.json"}, wantError: true},
		{name: "foreign_asset", extras: []string{"gcp-operational-health-v2.json", "gcp-operational-qualification-v2.json", "unreviewed.json"}, wantError: true},
		{name: "bad_digest", extras: []string{"gcp-operational-health-v2.json", "gcp-operational-qualification-v2.json"}, badDigest: true, wantError: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			directory := t.TempDir()
			tools := filepath.Join(directory, "tools")
			if err := os.Mkdir(tools, 0700); err != nil {
				t.Fatal(err)
			}
			names := []string{"api-catalog.json", "concurrency_contracts.json", "enrichment-coverage.json", "f5xc-api-specs-v12.99.0.zip", "index.json", "minimal-export-defaults.json", "openapi.json", "smsv2-contract-manifest.json", "smsv2-contract.json", "smsv2-evidence-receipt.json", "smsv2_parity_manifest.json", "upstream-contract-changes.json", "upstream-contract-removals.json"}
			names = append(names, test.extras...)
			assets := map[string]string{}
			for _, name := range names {
				assets[name] = "sha256:" + strings.Repeat("b", 64)
			}
			if test.badDigest {
				assets["gcp-operational-health-v2.json"] = "invalid"
			}
			canonical := map[string]string{"commit": strings.Repeat("a", 40), "event_type": "enriched-specs-updated", "source": "f5-sales-demo/api-specs-enriched", "tag": "v12.99.0", "target": "f5-sales-demo/terraform-provider-xcsh", "version": "12.99.0"}
			raw, err := json.Marshal(canonical)
			if err != nil {
				t.Fatal(err)
			}
			id := fmt.Sprintf("%x", sha256.Sum256(raw))
			state := map[string]any{
				"spec-release.json":                  map[string]any{"version": "12.99.0", "release_tag": "v12.99.0", "target_commit": canonical["commit"], "assets": assets},
				"spec-deliveries.json":               map[string]any{"version": 1, "deliveries": map[string]any{}},
				"provider-publication-receipts.json": map[string]any{"version": 1, "receipts": map[string]any{}},
				"spec-pending-delivery.json":         map[string]string{"delivery_id": id, "release_tag": "v12.99.0", "target_commit": canonical["commit"], "version": "12.99.0"},
			}
			for name, value := range state {
				raw, err := json.Marshal(value)
				if err != nil {
					t.Fatal(err)
				}
				if err = os.WriteFile(filepath.Join(tools, name), raw, 0600); err != nil {
					t.Fatal(err)
				}
			}
			if err = os.WriteFile(filepath.Join(tools, "spec-version.txt"), []byte("v12.99.0\n"), 0600); err != nil {
				t.Fatal(err)
			}
			script, err := filepath.Abs(filepath.Join("..", "scripts", "validate-provider-delivery-state.sh"))
			if err != nil {
				t.Fatal(err)
			}
			command := exec.Command("bash", script)
			command.Dir = directory
			output, err := command.CombinedOutput()
			if (err != nil) != test.wantError {
				t.Fatalf("inventory result: %v\n%s", err, output)
			}
			if test.wantError && !strings.Contains(string(output), "spec release pin is malformed") {
				t.Fatalf("unexpected refusal: %s", output)
			}
		})
	}
}
