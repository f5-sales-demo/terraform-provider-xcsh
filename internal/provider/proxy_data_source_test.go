// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccProxyDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_proxy.test"
	dataSourceName := "data.xcsh_proxy.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProxyDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccProxyDataSourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_proxy" "test" {
  name      = %[1]q
  namespace = "shared"
  http_proxy {
    enable_http {}
  }
  no_forward_proxy_policy = {}
  site_local_network = {}
  do_not_advertise = {}
}

data "xcsh_proxy" "test" {
  depends_on = [xcsh_proxy.test]
  name       = xcsh_proxy.test.name
  namespace  = xcsh_proxy.test.namespace
}
`, name))
}
