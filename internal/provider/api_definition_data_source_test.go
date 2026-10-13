// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccApiDefinitionDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.CompactRandomName(20)
	resourceName := "xcsh_api_definition.test"
	dataSourceName := "data.xcsh_api_definition.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApiDefinitionDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccApiDefinitionDataSourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_swagger_object" "fixture" {
 name = "%[1]s-spec"
 namespace = "shared"
 content = "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Acceptance API\",\"version\":\"1.0.0\"},\"servers\":[{\"url\":\"https://example.com\"}],\"paths\":{\"/health\":{\"get\":{\"responses\":{\"200\":{\"description\":\"OK\"}}}}}}"
}

resource "xcsh_api_definition" "test" {
  name      = %[1]q
  namespace = "shared"
  swagger_specs = [xcsh_swagger_object.fixture.path]
}

data "xcsh_api_definition" "test" {
  depends_on = [xcsh_api_definition.test]
  name       = xcsh_api_definition.test.name
  namespace  = xcsh_api_definition.test.namespace
}
`, name))
}
