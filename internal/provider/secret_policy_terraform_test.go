package provider_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	tfresource "github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// This is real Terraform against a protocol fixture, never fresh live API qualification.
func TestSecretPolicyTerraformStablePlansAndImport(t *testing.T) {
	exists := false
	var p client.SecretPolicy
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if !exists {
				http.NotFound(w, r)
				return
			}
			_ = json.NewEncoder(w).Encode(p)
		case "POST", "PUT":
			_ = json.NewDecoder(r.Body).Decode(&p)
			p.SystemMetadata.UID = "fixture-uid"
			p.Spec.Algorithm = "FIRST_MATCH"
			exists = true
			_, _ = w.Write([]byte(`{}`))
		case "DELETE":
			exists = false
			_, _ = w.Write([]byte(`{}`))
		default:
			http.Error(w, "unexpected", 400)
		}
	}))
	defer server.Close()
	config := fmt.Sprintf(`provider "xcsh" {
 api_url = %q
 api_token = "synthetic-test-token"
}
resource "xcsh_secret_policy" "test" {
 name = "fixture-policy"
 namespace = "demo"
 allow_f5xc = true
 decrypt_cache_timeout = "30s"
 rules = [{name="allow", action="ALLOW", client_selector=["app=demo"]}, {name="deny",action="DENY"}]
}`, server.URL)
	tfresource.UnitTest(t, tfresource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"xcsh": providerserver.NewProtocol6WithError(provider.New("test")())},
		Steps: []tfresource.TestStep{
			{Config: config, Check: tfresource.TestCheckResourceAttr("xcsh_secret_policy.test", "rule_combining_algorithm", "FIRST_MATCH")},
			{Config: config, PlanOnly: true},
			{ResourceName: "xcsh_secret_policy.test", ImportState: true, ImportStateId: "demo/fixture-policy", ImportStateVerify: true},
			{Config: config, PlanOnly: true},
		},
	})
}
