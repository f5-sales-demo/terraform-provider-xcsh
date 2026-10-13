// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccExternalConnectorDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.RequireConnectorEndpoint(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_external_connector.test"
	dataSourceName := "data.xcsh_external_connector.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccExternalConnectorDataSourceConfig_basic(rName, acctest.ReadOnlySiteName()),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccExternalConnectorDataSourceConfig_basic(name string, siteName string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_external_connector" "test" {
  name      = %[1]q
  namespace = "system"
  ce_site_reference {
    name = %[2]q
    namespace = "system"
  }
  gre {
    gre_parameters {
      peer_ip_address { addr = "192.0.2.2" }
      site_local_network = {}
      tunnel_mtu = 1370
      tunnel_eps {
        node = %[3]q
        interface = %[4]q
        local_tunnel_ip = "192.0.2.1/30"
        remote_tunnel_ip = "192.0.2.2/30"
      }
    }
  }
}

data "xcsh_external_connector" "test" {
  depends_on = [xcsh_external_connector.test]
  name       = xcsh_external_connector.test.name
  namespace  = xcsh_external_connector.test.namespace
}
`, name, siteName, acctest.ReadOnlyConnectorNode(), acctest.ReadOnlyConnectorInterface()))
}
