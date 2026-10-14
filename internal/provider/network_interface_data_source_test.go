// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccNetworkInterfaceDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_network_interface.test"
	dataSourceName := "data.xcsh_network_interface.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkInterfaceDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccNetworkInterfaceDataSourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_network_interface" "test" {
  name      = %[1]q
  namespace = "system"
  ethernet_interface {
    device = "eth0"
    mtu = 0
    priority = 0
    cluster = {}
    untagged = {}
    dhcp_client = {}
    site_local_inside_network = {}
    monitor_disabled = {}
    not_primary = {}
    no_ipv6_address = {}
  }
}

data "xcsh_network_interface" "test" {
  depends_on = [xcsh_network_interface.test]
  name       = xcsh_network_interface.test.name
  namespace  = xcsh_network_interface.test.namespace
}
`, name))
}
