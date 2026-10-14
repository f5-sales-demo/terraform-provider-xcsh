// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccForwardingClassDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_forwarding_class.test"
	dataSourceName := "data.xcsh_forwarding_class.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccForwardingClassDataSourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_forwarding_class" "test" {
  name      = %[1]q
  namespace = "system"
  queue_id_to_use = "DSCP_BEST_EFFORT"
  no_policer = {}
  interface_group = "ANY_AVAILABLE_INTERFACE"
  tos_value = 40
}

data "xcsh_forwarding_class" "test" {
  depends_on = [xcsh_forwarding_class.test]
  name       = xcsh_forwarding_class.test.name
  namespace  = xcsh_forwarding_class.test.namespace
}
`, name))
}
