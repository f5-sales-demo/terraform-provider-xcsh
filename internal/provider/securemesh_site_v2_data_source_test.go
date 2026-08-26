// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccSecuremeshSiteV2DataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_securemesh_site_v2.test"
	dataSourceName := "data.xcsh_securemesh_site_v2.test"
	testCase := resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSecuremeshSiteV2DataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	}
	acctest.RunWithMockOrReal(t, testCase, nil)
}

func testAccSecuremeshSiteV2DataSourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "xcsh_securemesh_site_v2" "test" {
  name      = %[1]q
  namespace = "system"
  baremetal {
    not_managed {}
  }

  disable_ha {}
  no_network_policy {}
  no_forward_proxy {}
  logs_streaming_disabled {}
  block_all_services {}
}

data "xcsh_securemesh_site_v2" "test" {
  depends_on = [xcsh_securemesh_site_v2.test]
  name       = xcsh_securemesh_site_v2.test.name
  namespace  = xcsh_securemesh_site_v2.test.namespace
}
`, name))
}
