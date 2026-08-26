// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/mocks"
)

type advanceMockResourceVersion struct {
	server *mocks.Server
	path   string
	before *string
	after  *string
}

type advanceLiveResourceVersion struct {
	name   string
	before *string
	after  *string
}

func (check advanceLiveResourceVersion) CheckPlan(ctx context.Context, _ plancheck.CheckPlanRequest, response *plancheck.CheckPlanResponse) {
	apiClient, err := acctest.GetTestClient()
	if err != nil {
		response.Error = fmt.Errorf("create live API client: %w", err)
		return
	}
	current, err := apiClient.GetSecuremeshSiteV2(ctx, "system", check.name)
	if err != nil {
		response.Error = fmt.Errorf("read SecureMesh v2 object before out-of-band advance: %w", err)
		return
	}
	*check.before = current.ResourceVersion
	if *check.before == "" {
		response.Error = fmt.Errorf("live SecureMesh v2 object has no current token")
		return
	}
	originalDescription := current.Metadata.Description
	current.Metadata.Description = "temporary concurrency probe"
	if _, err := apiClient.UpdateSecuremeshSiteV2(ctx, current); err != nil {
		response.Error = fmt.Errorf("advance SecureMesh v2 token out of band: %w", err)
		return
	}
	changed, err := apiClient.GetSecuremeshSiteV2(ctx, "system", check.name)
	if err != nil {
		response.Error = fmt.Errorf("read SecureMesh v2 object after out-of-band advance: %w", err)
		return
	}
	if changed.ResourceVersion == "" || changed.ResourceVersion == *check.before {
		response.Error = fmt.Errorf("live SecureMesh v2 token did not advance")
		return
	}
	changed.Metadata.Description = originalDescription
	if _, err := apiClient.UpdateSecuremeshSiteV2(ctx, changed); err != nil {
		response.Error = fmt.Errorf("restore SecureMesh v2 object after token advance: %w", err)
		return
	}
	restored, err := apiClient.GetSecuremeshSiteV2(ctx, "system", check.name)
	if err != nil {
		response.Error = fmt.Errorf("verify SecureMesh v2 restoration after token advance: %w", err)
		return
	}
	*check.after = restored.ResourceVersion
	if *check.after == "" || *check.after == changed.ResourceVersion || restored.Metadata.Description != originalDescription {
		response.Error = fmt.Errorf("live SecureMesh v2 object was not canonically restored with a newer token")
	}
}

func (check advanceMockResourceVersion) CheckPlan(_ context.Context, _ plancheck.CheckPlanRequest, response *plancheck.CheckPlanResponse) {
	stored, found := check.server.GetResource(check.path)
	if !found {
		response.Error = fmt.Errorf("mock SecureMesh v2 object is missing before stale-plan probe")
		return
	}
	object, ok := stored.(map[string]interface{})
	if !ok {
		response.Error = fmt.Errorf("mock SecureMesh v2 object has unexpected type %T", stored)
		return
	}
	*check.before, _ = object["resource_version"].(string)
	if *check.before == "" {
		response.Error = fmt.Errorf("mock SecureMesh v2 object has no current token")
		return
	}
	advanced, ok := check.server.AdvanceResourceVersion(check.path)
	if !ok || advanced == "" || advanced == *check.before {
		response.Error = fmt.Errorf("mock SecureMesh v2 token did not advance")
		return
	}
	*check.after = advanced
	check.server.ClearRequestLog()
}

func TestMockSecuremeshSiteV2StalePlanConflictsOnceAndPreservesState(t *testing.T) {
	acctest.SkipIfNoMockMode(t)
	mockConfig := acctest.SetupMockTest(t)
	defer mockConfig.Cleanup()

	name := acctest.RandomName("tf-acc-test-smsv2-concurrency")
	resourceName := "xcsh_securemesh_site_v2.test"
	path := "/api/config/namespaces/system/securemesh_site_v2s/" + name
	var plannedToken, advancedToken string

	config := func(description string) string {
		return acctest.ConfigCompose(mockConfig.MockProviderConfig(), fmt.Sprintf(`
resource "xcsh_securemesh_site_v2" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q

  baremetal {
    not_managed {}
  }

  disable_ha {}
  no_network_policy {}
  no_forward_proxy {}
  logs_streaming_disabled {}
  block_all_services {}
}
`, name, description))
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: mockConfig.ProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: config("initial"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "initial"),
				),
			},
			{
				Config: config("stale-overwrite"),
				ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
					advanceMockResourceVersion{
						server: mockConfig.Server,
						path:   path,
						before: &plannedToken,
						after:  &advancedToken,
					},
				}},
				ExpectError: regexp.MustCompile(`(?s)Stale Configuration.*securemesh_site_v2.*object changed`),
			},
			{
				PreConfig: func() {
					assertStaleSMSv2Attempt(t, mockConfig.Server, path, plannedToken, advancedToken)
					mockConfig.Server.ClearRequestLog()
				},
				Config: config("refreshed-update"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "refreshed-update"),
					func(_ *terraform.State) error {
						if puts := countMockPuts(mockConfig.Server.GetRequestLog(), path); puts != 1 {
							return fmt.Errorf("refreshed update made %d PUTs, want exactly one", puts)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccSecuremeshSiteV2StalePlanConflict(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	name := acctest.RandomName("tf-acc-test-smsv2-concurrency")
	resourceName := "xcsh_securemesh_site_v2.test"
	var plannedToken, advancedToken string
	config := func(description string) string {
		return testAccSecuremeshSiteV2ResourceConfig_withLabels(name, description, nil)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("xcsh_securemesh_site_v2"),
		Steps: []resource.TestStep{
			{
				Config: config("initial"),
				Check:  resource.TestCheckResourceAttr(resourceName, "description", "initial"),
			},
			{
				Config: config("stale-overwrite"),
				ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
					advanceLiveResourceVersion{name: name, before: &plannedToken, after: &advancedToken},
				}},
				ExpectError: regexp.MustCompile(`(?s)Stale Configuration.*securemesh_site_v2.*object changed`),
			},
			{
				PreConfig: func() {
					assertLiveSMSv2Unchanged(t, name, "initial", advancedToken)
				},
				Config: config("refreshed-update"),
				Check:  resource.TestCheckResourceAttr(resourceName, "description", "refreshed-update"),
			},
		},
	})
}

func assertLiveSMSv2Unchanged(t *testing.T, name, description, resourceVersion string) {
	t.Helper()
	apiClient, err := acctest.GetTestClient()
	if err != nil {
		t.Fatal(err)
	}
	current, err := apiClient.GetSecuremeshSiteV2(context.Background(), "system", name)
	if err != nil {
		t.Fatal(err)
	}
	if current.Metadata.Description != description {
		t.Fatalf("stale update changed remote description to %q", current.Metadata.Description)
	}
	if resourceVersion == "" || current.ResourceVersion != resourceVersion {
		t.Fatal("stale update changed or lost the out-of-band concurrency token")
	}
}

func assertStaleSMSv2Attempt(t *testing.T, server *mocks.Server, path, plannedToken, advancedToken string) {
	t.Helper()
	requests := server.GetRequestLog()
	if puts := countMockPuts(requests, path); puts != 1 {
		t.Fatalf("stale update made %d PUTs, want exactly one", puts)
	}
	for _, request := range requests {
		if request.Method != "PUT" || request.Path != path {
			continue
		}
		var body map[string]interface{}
		if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
			t.Fatal(err)
		}
		if body["resource_version"] != plannedToken {
			t.Fatalf("stale update sent token %#v, want the plan token", body["resource_version"])
		}
	}

	stored, found := server.GetResource(path)
	if !found {
		t.Fatal("stale update removed the remote object")
	}
	object := stored.(map[string]interface{})
	if object["resource_version"] != advancedToken {
		t.Fatal("stale update changed the remote token")
	}
	metadata := object["metadata"].(map[string]interface{})
	if metadata["description"] != "initial" {
		t.Fatal("stale update mutated the remote specification")
	}
}

func countMockPuts(requests []mocks.RequestRecord, path string) int {
	count := 0
	for _, request := range requests {
		if request.Method == "PUT" && request.Path == path {
			count++
		}
	}
	return count
}
