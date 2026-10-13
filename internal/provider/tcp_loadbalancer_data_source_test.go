// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccTcpLoadbalancerDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_tcp_loadbalancer.test"
	dataSourceName := "data.xcsh_tcp_loadbalancer.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTcpLoadbalancerDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccTcpLoadbalancerDataSourceConfig_basic(name string) string {
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

resource "xcsh_tcp_loadbalancer" "test" {
  name      = %[1]q
  namespace = "shared"
  domains = ["%[1]s.example.com"]
  tcp = {}
  listen_port = 80
  no_sni = {}
  no_service_policies = {}
  do_not_advertise = {}
  origin_pools_weights {
    pool {
      name = xcsh_origin_pool.test.name
      namespace = xcsh_origin_pool.test.namespace
    }
    weight = 1
  }
}

data "xcsh_tcp_loadbalancer" "test" {
  depends_on = [xcsh_tcp_loadbalancer.test]
  name       = xcsh_tcp_loadbalancer.test.name
  namespace  = xcsh_tcp_loadbalancer.test.namespace
}
`, name))
}
