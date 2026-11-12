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
		resource := openapi.ResourceTemplate{Name: name, HasConcurrencyToken: name != "network", ConcurrencyTokenGoName: "ConcurrencyToken", ConcurrencyTokenJSONName: "concurrency_token", TitleCase: map[string]string{"token": "Token", "securemesh_site_v2": "SecuremeshSiteV2", "network": "Network"}[name]}
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

func TestOriginalCreationCaptureGeneratedCheckpoint(t *testing.T) {
	for _, name := range []string{"token", "securemesh_site_v2", "network"} {
		titles := map[string]string{"token": "Token", "securemesh_site_v2": "SecuremeshSiteV2", "network": "Network"}
		resource := openapi.ResourceTemplate{Name: name, TitleCase: titles[name], HasConcurrencyToken: name != "network", ConcurrencyTokenGoName: "ConcurrencyToken", ConcurrencyTokenJSONName: "concurrency_token"}
		dir := t.TempDir()
		if err := GenerateResourceFile(&resource, dir); err != nil {
			t.Fatal(err)
		}
		raw, err := os.ReadFile(filepath.Join(dir, name+"_resource.go"))
		if err != nil {
			t.Fatal(err)
		}
		s := string(raw)
		preflight := strings.Index(s, "OpenOriginalCreationCapture(")
		if name == "network" {
			if preflight >= 0 {
				t.Fatal("unrelated resource changed")
			}
			continue
		}
		post := strings.Index(s, "r.client.Create"+titles[name]+"(")
		initial := strings.Index(s, "data.CreationIdentitySHA256 = creationIdentityReceipt(")
		state := strings.Index(s[initial:], "resp.State.Set(ctx, &data)") + initial
		receipt := strings.Index(s, "originalCapture.RecordCreation(")
		followup := strings.Index(s, "r.client.Get"+titles[name]+"(")
		if preflight < 0 || post < preflight || initial < post || state < initial || receipt < state || followup < receipt {
			t.Fatal("capture must preflight before POST and persist state and receipt before GET")
		}
		if !strings.Contains(s, "defer originalCapture.Close()") || !strings.Contains(s, "Original Creation Capture Failed") {
			t.Fatal("capture lifetime or explicit failure missing")
		}
	}
}
