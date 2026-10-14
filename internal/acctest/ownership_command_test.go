// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"encoding/json"
	"os"
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
	if inventory.Version != 1 || len(inventory.Resources) != 4 {
		t.Fatal("incomplete command inventory")
	}
	for _, resource := range inventory.Resources {
		if resource.Reason == "" {
			t.Fatal("missing operation evidence")
		}
		if resource.Type == "swagger_object" {
			if resource.Handler != "exact_version_content_digest" {
				t.Fatal("missing Swagger handler")
			}
		} else if resource.Handler != "non_disposable_action_scope" {
			t.Fatal("unsafe command cleanup handler")
		}
	}
}
