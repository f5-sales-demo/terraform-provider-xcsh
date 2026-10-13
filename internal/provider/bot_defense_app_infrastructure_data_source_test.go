// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccBotDefenseAppInfrastructureDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "xcsh_bot_defense_app_infrastructure.test"
	dataSourceName := "data.xcsh_bot_defense_app_infrastructure.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccBotDefenseAppInfrastructureDataSourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccBotDefenseAppInfrastructureDataSourceConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "xcsh_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [xcsh_namespace.test]
  create_duration = "5s"
}

resource "xcsh_bot_defense_app_infrastructure" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = xcsh_namespace.test.name
 environment_type = "PRODUCTION"
 traffic_type = "WEB"
 cloud_hosted {
  infra_host_name = "test.example.com"
  region = "US"
  ingress {
   host_name = "test.example.com"
   location = "AWS_US_EAST_1"
  }
  egress {
   ip_address = "192.0.2.1"
   location = "AWS_US_EAST_1"
  }
 }

}

data "xcsh_bot_defense_app_infrastructure" "test" {
  depends_on = [xcsh_bot_defense_app_infrastructure.test]
  name       = xcsh_bot_defense_app_infrastructure.test.name
  namespace  = xcsh_bot_defense_app_infrastructure.test.namespace
}
`, nsName, name))
}
