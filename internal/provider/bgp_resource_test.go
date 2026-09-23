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
	t.Skip("Skipping: bgp resource requires site infrastructure (CE/RE site) which is not available in acceptance tests")

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
				Config: testAccBGPConfig_basic(nsName, rName),
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

func testAccBGPConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "xcsh_bgp" "test" {
  name       = %[2]q
  namespace  = %[1]q

  peers {
    metadata {
      name = "test-peer"
    }
    external {
      asn     = 64512
      address = "192.168.1.1"
      no_authentication = {}
    }
    bfd_disabled = {}
    passive_mode_disabled = {}
    disable {}
  }

  where {
    site {
      kind      = "site"
      name      = "example-site"
      namespace = "system"
      tenant    = "example-corp"
    }
  }
}
`, nsName, name))
}
