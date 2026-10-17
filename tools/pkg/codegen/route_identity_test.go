package codegen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestHTTPLoadBalancerRouteIdentityWiring(t *testing.T) {
	r := &openapi.ResourceTemplate{Name: "http_loadbalancer", TitleCase: "HTTPLoadBalancer", HasConcurrencyToken: true, ConcurrencyTokenGoName: "ResourceVersion", HasNamespaceInPath: true}
	dir := t.TempDir()
	if err := GenerateResourceFile(r, dir); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(filepath.Join(dir, "http_loadbalancer_resource.go"))
	if err != nil {
		t.Fatal(err)
	}
	text := string(raw)
	for _, method := range []string{"Create", "Read", "Update"} {
		start := strings.Index(text, "func (r *HTTPLoadBalancerResource) "+method+"(")
		if start < 0 {
			t.Fatalf("missing %s method", method)
		}
		tail := text[start:]
		end := strings.Index(tail[1:], "\nfunc ")
		if end >= 0 {
			tail = tail[:end+1]
		}
		if !strings.Contains(tail, "encodeRouteIdentities(") || !strings.Contains(tail, "routeIdentitiesPrivateSlot") {
			t.Fatalf("%s does not retain route evidence", method)
		}
		if method == "Update" {
			preserve := strings.Index(tail, "preserveRouteIdentities(")
			write := strings.Index(tail, "r.client.UpdateHTTPLoadBalancer(")
			if preserve < 0 || write < 0 || preserve >= write {
				t.Fatal("PUT precedes identity preservation")
			}
			if strings.Index(tail, "r.client.GetHTTPLoadBalancer(") < write {
				t.Fatal("update adopts fresh read before PUT")
			}
		}
	}
}
