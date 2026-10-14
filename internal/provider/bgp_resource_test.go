// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccBGPResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Skip: BGP requires a real site reference which is infrastructure-dependent

	acctest.RequireExistingInfrastructure(t, "site")

	rName := acctest.RandomName("tf-acc-test-bgp")
	nsName := "system"
	resourceName := "xcsh_bgp.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("xcsh_bgp"),
		Steps: []resource.TestStep{
			{
				Config: testAccBGPConfig_basic(nsName, rName, acctest.ReadOnlySiteName()),
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
				ImportStateIdFunc:       testAccBGPImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccBGPImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccBGPConfig_basic(nsName, name, siteName string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "xcsh_bgp" "test" {
  name       = %[2]q
  namespace  = "system"

  bgp_parameters {
    asn = 65000
    from_site = {}
  }

  where {
    site {
      network_type = "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE"
      ref {
        name = %[3]q
        namespace = "system"
      }
    }
  }
}
`, nsName, name, siteName))
}
