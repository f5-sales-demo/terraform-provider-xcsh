// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSubnetResource_basic(t *testing.T) {

	acctest.RequireExistingInfrastructure(t, "site")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "xcsh_subnet.test"
	nsName := "system"
	name := acctest.RandomName("tf-acc-test-tf-acc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccSubnetResourceConfig_basic(nsName, name, acctest.ReadOnlySiteName()),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccSubnetResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccSubnetResourceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes["namespace"], rs.Primary.Attributes["name"]), nil
	}
}

func testAccSubnetResourceConfig_basic(nsName, name, siteName string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "xcsh_subnet" "test" {
  name       = %[2]q
  namespace  = "system"

  isolated_nw = {}
  site_subnet_params {
    static_ip = {}
    site {
      name = %[3]q
      namespace = "system"
    }
  }
}
`, nsName, name, siteName))
}
