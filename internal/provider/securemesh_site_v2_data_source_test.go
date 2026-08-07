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

	rName := acctest.RandomName("tf-acc-test-v2")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "xcsh_securemesh_site_v2.test"
	dataSourceName := "data.xcsh_securemesh_site_v2.test"

	testCase := resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccSecuremeshSiteV2DataSourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	}

	acctest.RunWithMockOrReal(t, testCase, func(mockCfg *acctest.MockTestConfig) {
		mockCfg.SetupNamespaceMock(nsName)
	})
}

func testAccSecuremeshSiteV2DataSourceConfig_basic(nsName, name string) string {
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

resource "xcsh_securemesh_site_v2" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = xcsh_namespace.test.name
}

data "xcsh_securemesh_site_v2" "test" {
  depends_on = [xcsh_securemesh_site_v2.test]
  name       = xcsh_securemesh_site_v2.test.name
  namespace  = xcsh_securemesh_site_v2.test.namespace
}
`, nsName, name))
}
