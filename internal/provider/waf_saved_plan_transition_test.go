package provider_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sync"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestWAFSavedPlanTransitions(t *testing.T) {
	var mu sync.Mutex
	var current map[string]interface{}
	var updates []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST", "PUT":
			var next map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&next); err != nil {
				t.Error(err)
				w.WriteHeader(400)
				return
			}
			spec := next["spec"].(map[string]interface{})
			_, custom := spec["blocking_page"]
			_, def := spec["use_default_blocking_page"]
			if custom && def {
				t.Error("payload has both blocking-page variants")
				w.WriteHeader(400)
				return
			}
			if r.Method == "PUT" {
				if next["resource_version"] != "fixture-version" {
					t.Error("concurrency token missing")
				}
				selected := "omitted"
				if custom {
					selected = "custom"
				}
				if def {
					selected = "default"
				}
				updates = append(updates, selected)
			}
			if !custom && !def {
				spec["use_default_blocking_page"] = map[string]interface{}{}
			}
			next["resource_version"] = "fixture-version"
			current = next
			_ = json.NewEncoder(w).Encode(current)
		case "GET":
			if current == nil {
				w.WriteHeader(404)
				return
			}
			_ = json.NewEncoder(w).Encode(current)
		case "DELETE":
			current = nil
			fmt.Fprint(w, "{}")
		default:
			t.Errorf("unexpected method %s", r.Method)
			w.WriteHeader(405)
		}
	}))
	defer server.Close()
	config := func(choice string) string {
		variant := ""
		if choice == "custom" {
			variant = "blocking_page {\n blocking_page = \"string:///Zml4dHVyZQ==\"\n response_code = \"Forbidden\"\n }"
		}
		if choice == "default" {
			variant = "use_default_blocking_page = {}"
		}
		if choice == "contradictory" {
			variant = "use_default_blocking_page = {}\nblocking_page {\n blocking_page = \"string:///Zml4dHVyZQ==\"\n response_code = \"Forbidden\"\n }"
		}
		return fmt.Sprintf("provider \"xcsh\" {\n api_url = %q\n api_token = \"fixture-token\"\n }\nresource \"xcsh_app_firewall\" \"this\" {\n name = \"fixture\"\n namespace = %q\n blocking = {}\n default_detection_settings = {}\n %s\n }\n", server.URL, "system", variant)
	}
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"xcsh": providerserver.NewProtocol6WithError(provider.New("test")())},
		Steps: []resource.TestStep{
			{Config: config("omitted"), Check: resource.TestCheckResourceAttr("xcsh_app_firewall.this", "id", "fixture")},
			{Config: config("custom"), ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("xcsh_app_firewall.this", plancheck.ResourceActionUpdate),
				plancheck.ExpectKnownValue("xcsh_app_firewall.this", tfjsonpath.New("use_default_blocking_page"), knownvalue.Null()),
			}}},
			{Config: config("default"), ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("xcsh_app_firewall.this", plancheck.ResourceActionUpdate),
				plancheck.ExpectKnownValue("xcsh_app_firewall.this", tfjsonpath.New("blocking_page"), knownvalue.Null()),
			}}},
			{Config: config("custom")},
			{Config: config("omitted"), ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("xcsh_app_firewall.this", plancheck.ResourceActionUpdate),
				plancheck.ExpectKnownValue("xcsh_app_firewall.this", tfjsonpath.New("blocking_page"), knownvalue.Null()),
			}}},
			{Config: config("contradictory"), PlanOnly: true, ExpectError: regexp.MustCompile("(?i)(conflict|exclusive|oneof|only one)")},
			{Config: config("omitted")},
		},
	})
	mu.Lock()
	defer mu.Unlock()
	if len(updates) != 4 {
		t.Fatalf("updates=%v", updates)
	}
	if updates[0] != "custom" || updates[1] != "default" || updates[2] != "custom" {
		t.Fatalf("updates=%v", updates)
	}
}
