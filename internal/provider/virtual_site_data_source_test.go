// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccVirtualSiteDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_virtual_site.test"
	dataSourceName := "data.xcsh_virtual_site.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualSiteDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccVirtualSiteDataSourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_virtual_site" "test" {
  name      = %[1]q
  namespace = "shared"
  site_type = "CUSTOMER_EDGE"
  site_selector {
    expressions = ["site-type=ce"]
  }
}

data "xcsh_virtual_site" "test" {
  depends_on = [xcsh_virtual_site.test]
  name       = xcsh_virtual_site.test.name
  namespace  = xcsh_virtual_site.test.namespace
}
`, name))
}
