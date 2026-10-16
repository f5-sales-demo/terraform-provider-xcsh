// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"
)

func TestCommandOwnershipInventoryIsExplicit(t *testing.T) {
	data, err := os.ReadFile("ownership_handler_exclusions.json")
	if err != nil {
		t.Fatal(err)
	}
	var inventory struct {
		Version   int `json:"version"`
		Resources []struct {
			Type    string `json:"type"`
			Handler string `json:"handler"`
			Reason  string `json:"reason"`
		} `json:"resources"`
	}
	if err = json.Unmarshal(data, &inventory); err != nil {
		t.Fatal(err)
	}
	expected := map[string]string{
		"public_ip_binding":           "non_disposable_action_scope",
		"registration_approval":       "non_disposable_action_scope",
		"smsv2_kvm_runtime_interface": "non_disposable_action_scope",
		"swagger_object":              "exact_version_content_digest",
	}
	if inventory.Version != 1 || len(inventory.Resources) != len(expected) {
		t.Fatal("incomplete command inventory")
	}
	for _, resource := range inventory.Resources {
		if resource.Reason == "" {
			t.Fatal("missing operation evidence")
		}
		handler, ok := expected[resource.Type]
		if !ok || resource.Handler != handler {
			t.Fatalf("unreviewed cleanup contract: %s %s", resource.Type, resource.Handler)
		}
		delete(expected, resource.Type)
	}
	if len(expected) != 0 {
		t.Fatal("missing explicit ownership contracts")
	}
}

func TestOwnershipInventoryGeneratorContracts(t *testing.T) {
	output, err := exec.Command("python3", "../../tools/test_generate_acceptance_inventory.py").CombinedOutput()
	if err != nil {
		t.Fatalf("inventory generator: %v\n%s", err, output)
	}
}
