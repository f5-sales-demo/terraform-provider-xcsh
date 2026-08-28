// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestMockDataGroupResource_APIEmptyMapPreservesConfiguredNull(t *testing.T) {
	acctest.SkipIfNoMockMode(t)
	mockCfg := acctest.SetupMockTest(t)
	defer mockCfg.Cleanup()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
		Steps: []resource.TestStep{{
			Config: fmt.Sprintf(`%s
resource "xcsh_data_group" "test" {
  name      = "state-normalization-data-group"
  namespace = "system"
  string_records {}
}
`, mockCfg.MockProviderConfig()),
		}},
	})
}

func TestMockOriginPoolResource_APIEmptyNestedMapPreservesConfiguredNull(t *testing.T) {
	acctest.SkipIfNoMockMode(t)
	mockCfg := acctest.SetupMockTest(t)
	defer mockCfg.Cleanup()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
		Steps: []resource.TestStep{{
			Config: fmt.Sprintf(`%s
resource "xcsh_origin_pool" "test" {
  name      = "state-normalization-origin-pool"
  namespace = "system"
  port      = 443

  origin_servers {
    public_name { dns_name = "example.com" }
  }

  no_tls {}
  same_as_endpoint_port {}
}
`, mockCfg.MockProviderConfig()),
		}},
	})
}
