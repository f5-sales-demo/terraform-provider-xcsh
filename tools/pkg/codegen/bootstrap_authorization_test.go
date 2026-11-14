package codegen

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTokenBootstrapAuthorizationPrecedesCredentialStatePublication(t *testing.T) {
	for _, name := range []string{"token", "securemesh_site_v2", "network"} {
		title := map[string]string{"token": "Token", "securemesh_site_v2": "SecuremeshSiteV2", "network": "Network"}[name]
		definition := openapi.ResourceTemplate{Name: name, TitleCase: title, HasConcurrencyToken: name != "network", ConcurrencyTokenGoName: "ConcurrencyToken", ConcurrencyTokenJSONName: "concurrency_token"}
		dir := t.TempDir()
		if err := GenerateResourceFile(&definition, dir); err != nil {
			t.Fatal(err)
		}
		raw, err := os.ReadFile(filepath.Join(dir, name+"_resource.go"))
		if err != nil {
			t.Fatal(err)
		}
		text := string(raw)
		declaration := strings.Contains(text, `const tokenBootstrapPrivateKey = "xcsh_original_bootstrap_token"`)
		if declaration != (name == "token") {
			t.Fatal("private bootstrap key declaration must belong to the generated token consumer")
		}
		start := strings.Index(text, "func (r *"+title+"Resource) Create(")
		end := strings.Index(text[start:], "func (r *"+title+"Resource) Read(") + start
		create := text[start:end]
		gate := strings.Index(create, "originalCapture.AwaitBootstrapAuthorization(")
		if name != "token" {
			if gate >= 0 {
				t.Fatal("unrelated resource changed")
			}
			continue
		}
		created := strings.Index(create, "originalCapture.RecordCreation(")
		populated := strings.Index(create, "populateTokenCredentialState(")
		published := strings.LastIndex(create, "resp.State.Set(ctx, &data)")
		if gate < 0 || created < 0 || populated < created || gate < populated || published < gate {
			t.Fatal("original token content must be authorized before its final state publication")
		}
		if !strings.Contains(create, "Original Bootstrap Authorization Failed") {
			t.Fatal("gate failure diagnostic missing")
		}
	}
}

func TestHeldBootstrapPrivateStatePrecedesFailureAndGuardsRefresh(t *testing.T) {
	resource := openapi.ResourceTemplate{Name: "token", TitleCase: "Token", HasConcurrencyToken: true, ConcurrencyTokenGoName: "ConcurrencyToken", ConcurrencyTokenJSONName: "concurrency_token"}
	dir := t.TempDir()
	if err := GenerateResourceFile(&resource, dir); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(filepath.Join(dir, "token_resource.go"))
	if err != nil {
		t.Fatal(err)
	}
	text := string(raw)
	start := strings.Index(text, "func (r *TokenResource) Create(")
	end := strings.Index(text, "func (r *TokenResource) Read(")
	create := text[start:end]
	held := strings.Index(create, "originalCapture.OriginalBootstrapTokenState(")
	checkpoint := strings.Index(create, "originalCapture.RecordCreation(")
	gate := strings.Index(create, "originalCapture.AwaitBootstrapAuthorization(")
	if held < 0 || checkpoint < held || gate < checkpoint {
		t.Fatal("created identity must retain held bootstrap before follow-up failures")
	}
	readEnd := strings.Index(text[end:], "func (r *TokenResource) Update(") + end
	read := text[end:readEnd]
	guarded := strings.Index(read, "suppressHeldTokenBootstrap(")
	state := strings.LastIndex(read, "resp.State.Set(ctx, &data)")
	if guarded < 0 || state < guarded || !strings.Contains(read, "req.Private.GetKey(ctx, tokenBootstrapPrivateKey)") {
		t.Fatal("held refresh can publish bootstrap credentials")
	}
}
