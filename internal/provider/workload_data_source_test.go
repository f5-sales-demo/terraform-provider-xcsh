// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccWorkloadDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_workload.test"
	dataSourceName := "data.xcsh_workload.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccWorkloadDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccWorkloadDataSourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_workload" "test" {
  name      = %[1]q
  namespace = "shared"
  simple_service { do_not_advertise = {} }
}

data "xcsh_workload" "test" {
  depends_on = [xcsh_workload.test]
  name       = xcsh_workload.test.name
  namespace  = xcsh_workload.test.namespace
}
`, name))
}
