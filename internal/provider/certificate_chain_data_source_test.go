// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccCertificateChainDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_certificate_chain.test"
	dataSourceName := "data.xcsh_certificate_chain.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateChainDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccCertificateChainDataSourceConfig_basic(name string) string {
	certs := acctest.MustGenerateTestCertificates()
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_certificate_chain" "test" {
  name      = %[1]q
  namespace = "shared"
  certificate_url = "string:///%[2]s"
}

data "xcsh_certificate_chain" "test" {
  depends_on = [xcsh_certificate_chain.test]
  name       = xcsh_certificate_chain.test.name
  namespace  = xcsh_certificate_chain.test.namespace
}
`, name, certs.IntermediateCABase64))
}
