// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccExternalConnectorResource_basic(t *testing.T) {

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.RequireConnectorEndpoint(t)

	rName := acctest.RandomName("tf-acc-test-ec")
	nsName := "system"
	resourceName := "xcsh_external_connector.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("xcsh_external_connector"),
		Steps: []resource.TestStep{
			{
				Config: testAccExternalConnectorConfig_basic(nsName, rName, acctest.ReadOnlySiteName()),
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
				ImportStateIdFunc:       testAccExternalConnectorImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccExternalConnectorImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccExternalConnectorConfig_basic(nsName, name, siteName string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "xcsh_external_connector" "test" {
  name      = %[2]q
  namespace = %[1]q
  ce_site_reference {
    name = %[3]q
    namespace = "system"
  }
  gre {
    gre_parameters {
      peer_ip_address { addr = "192.0.2.2" }
      site_local_inside_network = {}
      tunnel_mtu = 1370
      tunnel_eps {
        node = %[4]q
        interface = %[5]q
        local_tunnel_ip = "192.0.2.1/30"
        remote_tunnel_ip = "192.0.2.2/30"
      }
    }
  }
}
`, nsName, name, siteName, acctest.ReadOnlyConnectorNode(), acctest.ReadOnlyConnectorInterface()))
}
