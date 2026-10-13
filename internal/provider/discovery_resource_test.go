// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccDiscoveryResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Skip: discovery requires site/virtual_site with k8s or consul infrastructure
	// This cannot be tested without real k8s cluster or Consul server attached to a site

	acctest.RequireExistingInfrastructure(t, "site")

	rName := acctest.RandomName("tf-acc-test-disc")
	nsName := "system"
	resourceName := "xcsh_discovery.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("xcsh_discovery"),
		Steps: []resource.TestStep{
			{
				Config: testAccDiscoveryConfig_basic(nsName, rName, acctest.ReadOnlySiteName()),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccDiscoveryImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccDiscoveryImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccDiscoveryConfig_basic(nsName, name, siteName string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "xcsh_discovery" "test" {
  name       = %[2]q
  namespace  = "system"
  discovery_consul {
    access_info {
      connection_info { api_server = "http://consul.example.com:8500" }
    }
    publish_info { disable_spec = {} }
  }
  no_cluster_id = {}
  where {
    site {
      network_type = "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE"
      disable_internet_vip = {}
      ref {
        name = %[3]q
        namespace = "system"
      }
    }
  }
}
`, nsName, name, siteName))
}
