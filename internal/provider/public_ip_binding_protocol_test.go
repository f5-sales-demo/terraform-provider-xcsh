package provider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestPublicIPBindingTerraformRetargetAndZeroChange(t *testing.T) {
	original := []interface{}{map[string]interface{}{"name": "all", "namespace": "shared", "tenant": "ves-io"}}
	object := map[string]interface{}{
		"metadata":         map[string]interface{}{"name": "ip-example", "namespace": "shared"},
		"resource_version": "1",
		"spec":             map[string]interface{}{"ip": "192.0.2.55", "virtual_sites": original},
	}
	writes := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && (r.URL.Path == "/api/config/namespaces/demo/virtual_sites/canada" || r.URL.Path == "/api/config/namespaces/demo/virtual_sites/canada-two") {
			_ = json.NewEncoder(w).Encode(map[string]interface{}{"metadata": map[string]interface{}{"name": "canada", "namespace": "demo"}, "spec": map[string]interface{}{"site_type": "REGIONAL_EDGE"}})
			return
		}
		if r.URL.Path != "/api/config/namespaces/shared/public_ips/ip-example" {
			t.Errorf("unexpected API path: %s", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Method == http.MethodPut {
			writes++
			_ = json.NewDecoder(r.Body).Decode(&object)
		} else if r.Method != http.MethodGet {
			t.Errorf("unexpected mutation: %s", r.Method)
		}
		_ = json.NewEncoder(w).Encode(object)
	}))
	defer server.Close()
	config := func(target string) string {
		return fmt.Sprintf(`
provider "xcsh" {
 api_url = %q
 api_token = "fixture-token"
}
resource "xcsh_public_ip_binding" "test" {
 name = "ip-example"
 namespace = "shared"
 expected_ip = "192.0.2.55"
 virtual_site = %q
 virtual_site_namespace = "demo"
}
`, server.URL, target)
	}
	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"xcsh": providerserver.NewProtocol6WithError(New("test")()),
		},
		Steps: []resource.TestStep{
			{Config: config("canada"), Check: resource.TestCheckResourceAttr("xcsh_public_ip_binding.test", "virtual_site", "canada")},
			{Config: config("canada"), PlanOnly: true},
			{Config: config("canada-two"), Check: resource.TestCheckResourceAttr("xcsh_public_ip_binding.test", "virtual_site", "canada-two")},
			{Config: config("canada-two"), PlanOnly: true},
		},
	})
	if writes != 3 || !samePublicIPBindings(object["spec"].(map[string]interface{})["virtual_sites"].([]interface{}), original) {
		t.Fatalf("expected adoption, retarget and restore only; got %d writes", writes)
	}
}
