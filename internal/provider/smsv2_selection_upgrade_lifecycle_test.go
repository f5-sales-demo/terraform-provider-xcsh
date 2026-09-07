// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestMockSMSv2SelectionAndDrainLifecycle(t *testing.T) {
	acctest.SkipIfNoMockMode(t)
	mock := acctest.SetupMockTest(t)
	defer mock.Cleanup()
	name := acctest.RandomName("tf-acc-test-smsv2-choices")
	address := "xcsh_securemesh_site_v2.test"
	endpoint := "/api/config/namespaces/system/securemesh_site_v2s/" + name
	config := func(geography string, percentage int) string {
		return acctest.ConfigCompose(mock.MockProviderConfig(), fmt.Sprintf(`
resource "xcsh_securemesh_site_v2" "test" {
  name = %q
  namespace = "system"
  aws {
    not_managed {}
  }
  disable_ha              = {}
  block_all_services      = {}
  logs_streaming_disabled = {}
  re_select { specific_geography = %q }
  upgrade_settings {
    kubernetes_upgrade_drain {
      enable_upgrade_drain {
        drain_max_unavailable_node_percentage = %d
        drain_node_timeout = 300
      }
    }
  }
}
`, name, geography, percentage))
	}
	check := func(method, geography string, percentage int) resource.TestCheckFunc {
		return resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr(address, "re_select.specific_geography", geography),
			resource.TestCheckResourceAttr(address, "upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage", fmt.Sprint(percentage)),
			func(_ *terraform.State) error {
				matches := 0
				for _, request := range mock.Server.GetRequestLog() {
					if request.Method != method || (request.Path != endpoint && request.Path != "/api/config/namespaces/system/securemesh_site_v2s") {
						continue
					}
					var body struct {
						Spec struct {
							Selection map[string]any `json:"re_select"`
							Upgrade   struct {
								Drain struct {
									Enabled map[string]any `json:"enable_upgrade_drain"`
								} `json:"kubernetes_upgrade_drain"`
							} `json:"upgrade_settings"`
						} `json:"spec"`
					}
					if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
						return err
					}
					if len(body.Spec.Selection) != 1 || body.Spec.Selection["specific_geography"] != geography {
						return fmt.Errorf("%s lost geographic selection or sent another selection choice", method)
					}
					drain := body.Spec.Upgrade.Drain.Enabled
					if drain["drain_max_unavailable_node_percentage"] != float64(percentage) || drain["drain_node_timeout"] != float64(300) {
						return fmt.Errorf("%s lost percentage or timeout", method)
					}
					if _, exists := drain["drain_max_unavailable_node_count"]; exists {
						return fmt.Errorf("%s sent both drain choices", method)
					}
					matches++
				}
				if matches != 1 {
					return fmt.Errorf("expected one %s request, got %d", method, matches)
				}
				return nil
			},
		)
	}
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: mock.ProtoV6ProviderFactories(),
		CheckDestroy: func(_ *terraform.State) error {
			if _, exists := mock.Server.GetResource(endpoint); exists {
				return fmt.Errorf("SMSv2 object remains after destroy")
			}
			return nil
		},
		Steps: []resource.TestStep{
			{Config: config("US", 50), Check: check("POST", "US", 50)},
			{
				Config:      strings.Replace(config("US", 50), `re_select { specific_geography = "US" }`, "re_select {\n specific_geography = \"US\"\n geo_proximity = {}\n}", 1),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("Conflicting Configuration"),
			},
			{
				Config:      strings.Replace(config("US", 50), "drain_max_unavailable_node_percentage = 50", "drain_max_unavailable_node_percentage = 50\n drain_max_unavailable_node_count = 1", 1),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("Conflicting Configuration"),
			},
			{PreConfig: mock.Server.ClearRequestLog, Config: config("EU", 25), Check: check("PUT", "EU", 25)},
			{
				Config: strings.Replace(config("EU", 25), `re_select { specific_geography = "EU" }`, `re_select {
  geo_proximity = {}
  specific_re {
    primary_re = "example-primary-re"
    backup_re = "example-backup-re"
  }
}`, 1),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("Conflicting Configuration"),
			},
			{ResourceName: address, ImportState: true, ImportStateId: "system/" + name, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"timeouts"}},
			{Config: config("EU", 25), PlanOnly: true},
		},
	})
}
