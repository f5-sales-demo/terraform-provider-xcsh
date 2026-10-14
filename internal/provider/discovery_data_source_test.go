// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccDiscoveryDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.RequireExistingInfrastructure(t, "site")

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_discovery.test"
	dataSourceName := "data.xcsh_discovery.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDiscoveryDataSourceConfig_basic(rName, acctest.ReadOnlySiteName()),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccDiscoveryDataSourceConfig_basic(name string, siteName string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_discovery" "test" {
  name      = %[1]q
  namespace = "system"
  no_cluster_id = {}
  discovery_consul {
    access_info {
      connection_info { api_server = "http://consul.example.com:8500" }
    }
    publish_info { disable_spec = {} }
  }
  where {
    site {
      network_type = "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE"
      disable_internet_vip = {}
      ref {
        name = %[2]q
        namespace = "system"
      }
    }
  }
}

data "xcsh_discovery" "test" {
  depends_on = [xcsh_discovery.test]
  name       = xcsh_discovery.test.name
  namespace  = xcsh_discovery.test.namespace
}
`, name, siteName))
}
