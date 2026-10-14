// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccCRLResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-crl")
	nsName := "system"
	resourceName := "xcsh_crl.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("xcsh_crl"),
		Steps: []resource.TestStep{
			{
				Config: testAccCRLConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttr(resourceName, "server_address", "crl.example.com"),
					resource.TestCheckResourceAttr(resourceName, "server_port", "80"),
					resource.TestCheckResourceAttr(resourceName, "refresh_interval", "24"),
					resource.TestCheckResourceAttr(resourceName, "timeout", "30"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccCRLImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccCRLResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-crl")
	nsName := "system"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckResourceDestroyed("xcsh_crl"),
		Steps: []resource.TestStep{
			{Config: testAccCRLConfig_basic(nsName, rName)},
			{Config: testAccCRLConfig_basic(nsName, rName), PlanOnly: true, ExpectNonEmptyPlan: false},
		},
	})
}

func testAccCRLImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		namespace := rs.Primary.Attributes["namespace"]
		name := rs.Primary.Attributes["name"]
		return fmt.Sprintf("%s/%s", namespace, name), nil
	}
}

func testAccCRLConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_crl" "test" {
  name       = %[2]q
  namespace  = "system"

  server_address   = "crl.example.com"
  server_port      = 80
  refresh_interval = 24
  timeout          = 30

  http_access {
    path = "/crl/test.crl"
  }
}
`, nsName, name))
}
