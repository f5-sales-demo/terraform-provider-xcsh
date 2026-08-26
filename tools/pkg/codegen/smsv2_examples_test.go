// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteSecuremeshSiteV2Examples(t *testing.T) {
	root := t.TempDir()
	if err := WriteSecuremeshSiteV2Examples(root); err != nil {
		t.Fatal(err)
	}
	variants := append(append([]string{}, SecuremeshSiteV2ProviderChoices...), "segment-vrf")
	for _, variant := range variants {
		path := filepath.Join(root, "examples", "resources", "xcsh_securemesh_site_v2", "variants", variant, "resource.tf")
		content, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("%s: %v", variant, err)
		}
		text := string(content)
		if !strings.Contains(text, `namespace = "system"`) || !strings.Contains(text, "required_providers") {
			t.Fatalf("%s example is structurally incomplete", variant)
		}
		for _, guidance := range []string{"disable_ha {}", "no_network_policy {}", "no_forward_proxy {}", "logs_streaming_disabled {}", "block_all_services {}"} {
			if !strings.Contains(text, guidance) {
				t.Fatalf("%s example is missing minimum-configuration guidance %q", variant, guidance)
			}
		}
		if variant != "segment-vrf" && !strings.Contains(text, variant+" {") {
			t.Fatalf("%s example does not select its provider block", variant)
		}
		if variant == "segment-vrf" && (!strings.Contains(text, "segment_network {") || !strings.Contains(text, `name      = "example-segment"`)) {
			t.Fatal("Segment VRF example does not use a named Segment reference")
		}
	}
}
