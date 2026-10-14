// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package provider_test

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
	"testing"
)

func TestFixedSystemFixturesGenerateSystemOperationScope(t *testing.T) {
	for name, config := range map[string]string{
		"bgp":                testAccBGPConfig_basic("tf-acc-test-fixture", "fixture", "fixture-site"),
		"external_connector": testAccExternalConnectorConfig_basic("tf-acc-test-fixture", "fixture", "fixture-site"),
		"virtual_site":       testAccVirtualSiteResourceConfig_basic("tf-acc-test-fixture", "fixture"),
	} {
		t.Run(name, func(t *testing.T) {
			requirements, err := acctest.FixtureRequirements(config)
			if err != nil {
				t.Fatal(err)
			}
			found := false
			for _, r := range requirements {
				if r.Type == "xcsh_"+name {
					found = true
					if r.Namespace != "system" {
						t.Fatal("capability probe differs from actual fixed system test scope")
					}
				}
			}
			if !found {
				t.Fatal("fixture operation missing")
			}
		})
	}
}
