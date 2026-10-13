// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccUdpLoadbalancerDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_udp_loadbalancer.test"
	dataSourceName := "data.xcsh_udp_loadbalancer.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUdpLoadbalancerDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccUdpLoadbalancerDataSourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_origin_pool" "test" {
  name       = "%[1]s-pool"
  namespace  = "system"
  origin_servers {
    public_ip {
      ip = "192.0.2.1"
    }
  }
  port               = 80
  endpoint_selection = "LOCAL_PREFERRED"
  loadbalancer_algorithm = "ROUND_ROBIN"
}

resource "xcsh_udp_loadbalancer" "test" {
  name      = %[1]q
  namespace = "system"
  domains     = ["%[1]s.example.com"]
  listen_port = 53
  origin_pools_weights {
    pool {
      name      = xcsh_origin_pool.test.name
      namespace = xcsh_origin_pool.test.namespace
    }
    weight = 1
  }
  advertise_on_public_default_vip = {}
}

data "xcsh_udp_loadbalancer" "test" {
  depends_on = [xcsh_udp_loadbalancer.test]
  name       = xcsh_udp_loadbalancer.test.name
  namespace  = xcsh_udp_loadbalancer.test.namespace
}
`, name))
}
