package codegen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestCreationIdentityGeneratedOnlyForCEObjects(t *testing.T) {
	for _, name := range []string{"token", "securemesh_site_v2", "network"} {
		resource := openapi.ResourceTemplate{Name: name, HasConcurrencyToken: name != "network", TitleCase: map[string]string{"token": "Token", "securemesh_site_v2": "SecuremeshSiteV2", "network": "Network"}[name]}
		dir := t.TempDir()
		err := GenerateResourceFile(&resource, dir)
		generated, readErr := os.ReadFile(filepath.Join(dir, name+"_resource.go"))
		if readErr != nil {
			t.Fatal(readErr)
		}
		if err != nil {
			t.Fatal(err)
		}
		wants := name != "network"
		if strings.Contains(string(generated), "creation_identity_sha256") != wants {
			t.Fatalf("receipt scope for %s", name)
		}
		if wants {
			capture := strings.Index(string(generated), "data.CreationIdentitySHA256 = creationIdentityReceipt(")
			persist := strings.Index(string(generated)[capture:], "resp.State.Set(ctx, &data)")
			read := strings.Index(string(generated)[capture:], "verifyCreationIdentity(")
			if persist < 0 || read < 0 || persist >= read {
				t.Fatal("original POST receipt must be persisted before follow-up verification")
			}
		}
	}
}
