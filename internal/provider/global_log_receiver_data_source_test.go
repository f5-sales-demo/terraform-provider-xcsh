// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccGlobalLogReceiverDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_global_log_receiver.test"
	dataSourceName := "data.xcsh_global_log_receiver.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGlobalLogReceiverDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccGlobalLogReceiverDataSourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_global_log_receiver" "test" {
  name      = %[1]q
  namespace = "system"
  http_receiver {
    uri = "http://logs.example.com/fixture"
    auth_none = {}
    no_tls = {}
  }
  ns_current = {}
  audit_logs = {}
}

data "xcsh_global_log_receiver" "test" {
  depends_on = [xcsh_global_log_receiver.test]
  name       = xcsh_global_log_receiver.test.name
  namespace  = xcsh_global_log_receiver.test.namespace
}
`, name))
}
