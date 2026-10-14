// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccTunnelDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_tunnel.test"
	dataSourceName := "data.xcsh_tunnel.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTunnelDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccTunnelDataSourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_tunnel" "test" {
  name      = %[1]q
  namespace = "system"
  tunnel_type = "IPSEC_PSK"
  local_ip {
    ip_address {
      auto = {}
      virtual_network_type { site_local = {} }
    }
  }
  remote_ip {
    ip {
      ipv4 { addr = "192.0.2.1" }
    }
  }
  params {
    ipsec {
      ipsec_psk {
        clear_secret_info { url = "string:///dGVzdC1wc2stc2VjcmV0" }
      }
    }
  }
}

data "xcsh_tunnel" "test" {
  depends_on = [xcsh_tunnel.test]
  name       = xcsh_tunnel.test.name
  namespace  = xcsh_tunnel.test.namespace
}
`, name))
}
