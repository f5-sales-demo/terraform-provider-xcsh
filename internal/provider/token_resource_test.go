package provider_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestMockTokenResource_basic(t *testing.T) {
	resourceName := "xcsh_token.test"

	acctest.SkipIfNoMockMode(t)
	mockCfg := acctest.SetupMockTest(t)
	defer mockCfg.Cleanup()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: `
				resource "xcsh_token" "test" {
					name      = "test-token"
					namespace = "system"
				}

				output "test_token" {
					value = xcsh_token.test.uid
				}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// The uid should be completely masked in output (sensitive)
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test-token"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					// Test that the uid output is generated (but we can't test tf's internal masking here easily,
					// so we just make sure we get *a* value and the resource applies cleanly)
					resource.TestMatchResourceAttr(resourceName, "uid", regexp.MustCompile(`(^[0-9a-fA-F-]+$|^mock-uid-[0-9]+$)`)),
				),
			},
		},
	})
}
