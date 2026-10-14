// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccProtocolPolicerDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_protocol_policer.test"
	dataSourceName := "data.xcsh_protocol_policer.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProtocolPolicerDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccProtocolPolicerDataSourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_policer" "test" {
 name = "%[1]s-policer"
 namespace = "system"
 burst_size = 1000
 committed_information_rate = 1000
}

resource "xcsh_protocol_policer" "test" {
  name      = %[1]q
  namespace = "system"
  protocol_policer {
    protocol { dns = {} }
    policer {
      name = xcsh_policer.test.name
      namespace = "system"
    }
  }
}

data "xcsh_protocol_policer" "test" {
  depends_on = [xcsh_protocol_policer.test]
  name       = xcsh_protocol_policer.test.name
  namespace  = xcsh_protocol_policer.test.namespace
}
`, name))
}
