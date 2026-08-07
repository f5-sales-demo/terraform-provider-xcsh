package provider_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestMockTokenDataSource_basic(t *testing.T) {
	dataSourceName := "data.xcsh_token.test"

	acctest.SkipIfNoMockMode(t)
	mockCfg := acctest.SetupMockTest(t)
	defer mockCfg.Cleanup()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: `
				resource "xcsh_token" "test" {
					name      = "test-token-ds"
					namespace = "system"
				}

				data "xcsh_token" "test" {
					name      = xcsh_token.test.name
					namespace = xcsh_token.test.namespace
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", "test-token-ds"),
					resource.TestCheckResourceAttr(dataSourceName, "namespace", "system"),
					// Assert the data source returns a non-empty uid that matches a standard ID.
					resource.TestMatchResourceAttr(dataSourceName, "uid", regexp.MustCompile(`(^[0-9a-fA-F-]+$|^mock-uid-[0-9]+$)`)),
				),
			},
		},
	})
}
