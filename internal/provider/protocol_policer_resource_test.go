// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccProtocolPolicerResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Skip: protocol_policer requires policer reference but API validation is
	// rejecting the request format (BAD_REQUEST). Need further investigation of
	// the exact API payload structure for the policer reference.

	rName := acctest.RandomName("tf-acc-test-pp")
	nsName := "system"
	resourceName := "xcsh_protocol_policer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("xcsh_protocol_policer"),
		Steps: []resource.TestStep{
			{
				Config: testAccProtocolPolicerConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccProtocolPolicerImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccProtocolPolicerImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccProtocolPolicerConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_policer" "test" {
  name                      = "%[2]s-policer"
  namespace = "system"
  burst_size                = 1000
  committed_information_rate = 1000
}

resource "xcsh_protocol_policer" "test" {
  depends_on = [xcsh_policer.test]
  name       = %[2]q
  namespace = "system"

  protocol_policer {
    protocol {
      dns = {}
    }

    policer {
      name      = xcsh_policer.test.name
      namespace = "system"
    }
  }
}
`, nsName, name))
}
