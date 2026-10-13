// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccNatPolicyDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.RequireExistingInfrastructure(t, "site")

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_nat_policy.test"
	dataSourceName := "data.xcsh_nat_policy.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNatPolicyDataSourceConfig_basic(rName, acctest.ReadOnlySiteName()),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccNatPolicyDataSourceConfig_basic(name string, siteName string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_virtual_network" "fixture" {
 name = "%[1]s-network"
 namespace = "system"
 global_network = {}
}
resource "xcsh_nat_policy" "test" {
  name      = %[1]q
  namespace = "system"
  site {
    refs {
      name = %[2]q
      namespace = "system"
    }
  }
  rules {
    name = "fixture"
    enable = {}
    virtual_network {
      refs {
        name = xcsh_virtual_network.fixture.name
        namespace = xcsh_virtual_network.fixture.namespace
      }
    }
    criteria {
      any = {}
      source_cidr = ["192.0.2.0/24"]
    }
    action { virtual_cidr = "198.51.100.0/24" }
  }
}

data "xcsh_nat_policy" "test" {
  depends_on = [xcsh_nat_policy.test]
  name       = xcsh_nat_policy.test.name
  namespace  = xcsh_nat_policy.test.namespace
}
`, name, siteName))
}
