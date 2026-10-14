// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccNetworkFirewallDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_network_firewall.test"
	dataSourceName := "data.xcsh_network_firewall.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkFirewallDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccNetworkFirewallDataSourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_forward_proxy_policy" "test" {
  name = "%[1]s-forward-proxy"
  namespace = "system"
  any_proxy = {}
  allow_all = {}
}

resource "xcsh_network_firewall" "test" {
  name      = %[1]q
  namespace = "system"
  active_forward_proxy_policies {
    forward_proxy_policies {
      name      = xcsh_forward_proxy_policy.test.name
      namespace = xcsh_forward_proxy_policy.test.namespace
    }
  }
}

data "xcsh_network_firewall" "test" {
  depends_on = [xcsh_network_firewall.test]
  name       = xcsh_network_firewall.test.name
  namespace  = xcsh_network_firewall.test.namespace
}
`, name))
}
