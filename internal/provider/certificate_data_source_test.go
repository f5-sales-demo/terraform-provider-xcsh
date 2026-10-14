// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestAccCertificateDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "xcsh_certificate.test"
	dataSourceName := "data.xcsh_certificate.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccCertificateDataSourceConfig_basic(name string) string {
	certs := acctest.MustGenerateTestCertificates()
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`

resource "xcsh_certificate" "test" {
  name      = %[1]q
  namespace = "shared"
  certificate_url = "string:///%[2]s"
  private_key {
    clear_secret_info {
      url = "string:///%[3]s"
    }
  }
  disable_ocsp_stapling = {}
}

data "xcsh_certificate" "test" {
  depends_on = [xcsh_certificate.test]
  name       = xcsh_certificate.test.name
  namespace  = xcsh_certificate.test.namespace
}
`, name, certs.ServerCertBase64, certs.ServerKeyBase64))
}
