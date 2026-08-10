// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

// =============================================================================
// SECUREMESH SITE V2 RESOURCE ACCEPTANCE TESTS
// =============================================================================

func TestAccSecuremeshSiteV2Resource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-smsite-v2")
	resourceName := "xcsh_securemesh_site_v2.test"

	testCase := resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("xcsh_securemesh_site_v2"),
		Steps: []resource.TestStep{
			// Step 1: Create securemesh_site_v2 with top-level labels map
			{
				Config: testAccSecuremeshSiteV2ResourceConfig_withLabels(
					rName,
					"Initial description",
					map[string]string{"key1": "val1", "key2": "val2"},
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
					resource.TestCheckResourceAttr(resourceName, "labels.key1", "val1"),
					resource.TestCheckResourceAttr(resourceName, "labels.key2", "val2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Step 2: Update securemesh_site_v2 description and top-level labels
			{
				Config: testAccSecuremeshSiteV2ResourceConfig_withLabels(
					rName,
					"Updated description",
					map[string]string{"key1": "val1_updated", "key2": "val2_updated"},
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
					resource.TestCheckResourceAttr(resourceName, "labels.key1", "val1_updated"),
					resource.TestCheckResourceAttr(resourceName, "labels.key2", "val2_updated"),
				),
			},
			// Step 3: Import state verification
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "software_settings"},
				ImportStateIdFunc:       testAccSecuremeshSiteV2ImportStateIdFunc(resourceName),
			},
		},
	}

	acctest.RunWithMockOrReal(t, testCase, func(mockCfg *acctest.MockTestConfig) {
		// Mock setup if applicable
	})
}

func testAccSecuremeshSiteV2ImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		namespace := rs.Primary.Attributes["namespace"]
		name := rs.Primary.Attributes["name"]
		return fmt.Sprintf("%s/%s", namespace, name), nil
	}
}

func testAccSecuremeshSiteV2ResourceConfig_withLabels(name, description string, topLabels map[string]string) string {
	var topLabelsStr strings.Builder
	if len(topLabels) > 0 {
		topLabelsStr.WriteString("  labels = {\n")
		for k, v := range topLabels {
			topLabelsStr.WriteString(fmt.Sprintf("    %q = %q\n", k, v))
		}
		topLabelsStr.WriteString("  }\n")
	}

	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "xcsh_securemesh_site_v2" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q
%[3]s  baremetal {
    not_managed {}
  }
}
`, name, description, topLabelsStr.String()))
}
