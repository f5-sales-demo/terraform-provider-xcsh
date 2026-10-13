// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package provider_test

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
	"os"
	"strings"
	"testing"
)

func TestTenantObjectFixturesUseSupportedNamespacesWithoutAppliances(t *testing.T) {
	for _, tc := range []struct{ name, config, typ, namespace string }{
		{"API definition", testAccAPIDefinitionConfig_basic("tf-acc-test-fixture"), "xcsh_api_definition", "shared"},
		{"workload flavor", testAccWorkloadFlavorResourceConfig_basic("unused", "tf-acc-test-fixture"), "xcsh_workload_flavor", "shared"},
		{"site mesh group", testAccSiteMeshGroupResourceConfig_fullMesh("unused", "tf-acc-test-fixture", "fixture", "example-virtual-site", "system"), "xcsh_site_mesh_group", "system"},
		{"cluster group", testAccDcClusterGroupConfig_basic("unused", "tf-acc-test-fixture"), "xcsh_dc_cluster_group", "system"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			body, err := acctest.ParseFixture(tc.config)
			if err != nil {
				t.Fatal(err)
			}
			found := false
			for _, block := range body.Blocks {
				if block.Type != "resource" {
					continue
				}
				if block.Labels[0] == "xcsh_namespace" {
					t.Error("tenant object fixture unnecessarily creates a namespace")
				}
				if block.Labels[0] != tc.typ {
					continue
				}
				found = true
				value, diags := block.Body.Attributes["namespace"].Expr.Value(nil)
				if diags.HasErrors() || value.AsString() != tc.namespace {
					t.Error("tenant object uses unsupported namespace")
				}
			}
			if !found {
				t.Fatal("tenant object missing")
			}
		})
	}
}

// The datasource's creating resource must use the same mesh prerequisite as
// resource lifecycle tests rather than silently exercising an invalid fixture.
func TestMeshDataSourceFixtureHasRequiredSelection(t *testing.T) {
	body, err := acctest.ParseFixture(testAccSiteMeshGroupDataSourceConfig_basic("tf-acc-test-fixture", "example-virtual-site", "system"))
	if err != nil {
		t.Fatal(err)
	}
	for _, block := range body.Blocks {
		if block.Type != "resource" || block.Labels[0] != "xcsh_site_mesh_group" {
			continue
		}
		for _, name := range []string{"bfd_disabled", "disable_re_fallback"} {
			if block.Body.Attributes[name] == nil {
				t.Errorf("mesh datasource missing required %s", name)
			}
		}
		found := map[string]bool{}
		for _, nested := range block.Body.Blocks {
			found[nested.Type] = true
		}
		if !found["virtual_site"] || !found["full_mesh"] {
			t.Error("mesh datasource missing selector or mesh mode")
		}
		return
	}
	t.Fatal("mesh datasource fixture has no creating resource")
}

// Segment names are composed with the owning site name by the API; keep both
// disposable identities compact enough for the resulting DNS label.
func TestSegmentVRFFixtureUsesCompactSiteIdentity(t *testing.T) {
	data, err := os.ReadFile("securemesh_site_v2_resource_test.go")
	if err != nil {
		t.Fatal(err)
	}
	source := string(data)
	start := strings.Index(source, "func TestAccSecuremeshSiteV2Resource_segmentVRF(")
	end := strings.Index(source[start:], "func testAccSecuremeshSiteV2ImportStateIdFunc")
	body := source[start : start+end]
	if strings.Count(body, "acctest.CompactRandomName(20)") != 2 {
		t.Fatal("segment VRF site identity can exceed derived API label limit")
	}
}
