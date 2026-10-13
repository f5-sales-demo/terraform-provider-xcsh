// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccSubnetDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.RequireExistingInfrastructure(t, "site")

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_subnet.test"
	dataSourceName := "data.xcsh_subnet.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSubnetDataSourceConfig_basic(rName, acctest.ReadOnlySiteName()),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccSubnetDataSourceConfig_basic(name string, siteName string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_subnet" "test" {
  name      = %[1]q
  namespace = "system"
  isolated_nw = {}
  site_subnet_params {
    dhcp = {}
    site {
      name = %[2]q
      namespace = "system"
    }
  }
}

data "xcsh_subnet" "test" {
  depends_on = [xcsh_subnet.test]
  name       = xcsh_subnet.test.name
  namespace  = xcsh_subnet.test.namespace
}
`, name, siteName))
}
