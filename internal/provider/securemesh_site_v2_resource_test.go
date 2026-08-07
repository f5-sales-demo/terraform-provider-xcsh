// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

// =============================================================================
// SECUREMESH SITE V2 RESOURCE ACCEPTANCE TESTS
// =============================================================================

func TestAccSecuremeshSiteV2Resource_basic(t *testing.T) {
	t.Skip("Skipping: requires physical/cloud site infrastructure and registration token")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	rName := acctest.RandomName("tf-acc-test-smsite-v2")
	resourceName := "xcsh_securemesh_site_v2.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("xcsh_securemesh_site_v2"),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create securemesh_site_v2 with minimal configuration
			{
				Config: testAccSecuremeshSiteV2ResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Step 2: Import state verification
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccSecuremeshSiteV2ImportStateIdFunc(resourceName),
			},
		},
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

func testAccSecuremeshSiteV2ResourceConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "xcsh_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [xcsh_namespace.test]
  create_duration = "5s"
}

resource "xcsh_securemesh_site_v2" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = xcsh_namespace.test.name
}
`, nsName, name))
}
