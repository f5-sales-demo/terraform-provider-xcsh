//go:build linux

package provider

import (
	"context"
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/codegen"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func TestTokenCreateOriginalCaptureLifecycle(t *testing.T) {
	if os.Getenv("XCSH_ORIGINAL_CREATION_GENERATED_CHILD") != "1" {
		verifyOriginalCreationGeneratedLifecycle(t)
		return
	}
	for _, kind := range []string{"followup-failure", "sink-failure", "missing-uid", "wrong-scope"} {
		t.Run(kind, func(t *testing.T) {
			file, dir := captureFixture(t)
			raw, err := os.ReadFile(file)
			if err != nil {
				t.Fatal(err)
			}
			var scope OriginalCreationScope
			if err = json.Unmarshal(raw, &scope); err != nil {
				t.Fatal(err)
			}
			if kind != "wrong-scope" {
				scope.Resources[0].Name = "fixture"
			}
			raw, err = json.Marshal(scope)
			if err != nil {
				t.Fatal(err)
			}
			if err = os.Chmod(file, 0600); err != nil {
				t.Fatal(err)
			}
			if err = os.WriteFile(file, raw, 0600); err != nil {
				t.Fatal(err)
			}
			if err = os.Chmod(file, 0400); err != nil {
				t.Fatal(err)
			}
			var posts, gets atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if r.Method == http.MethodPost {
					posts.Add(1)
					if kind == "sink-failure" {
						if err := os.Chmod(dir, 0755); err != nil {
							t.Error(err)
						}
					}
					if kind == "missing-uid" {
						w.Write([]byte(`{"metadata":{"name":"fixture","namespace":"system"}}`))
					} else {
						w.Write([]byte(`{"metadata":{"name":"fixture","namespace":"system"},"system_metadata":{"uid":"original-uid"}}`))
					}
					return
				}
				gets.Add(1)
				entries, err := os.ReadDir(dir)
				if err != nil {
					t.Error(err)
				}
				found := false
				for _, entry := range entries {
					if strings.HasPrefix(entry.Name(), "receipt-") {
						found = true
						raw, err := os.ReadFile(filepath.Join(dir, entry.Name()))
						if err != nil {
							t.Error(err)
						}
						if !strings.Contains(string(raw), strings.TrimPrefix(creationIdentityReceipt("original-uid").ValueString(), "sha256:")) {
							t.Error("original receipt missing before GET")
						}
					}
				}
				if !found {
					t.Error("follow-up read ran before durable original capture")
				}
				w.WriteHeader(http.StatusServiceUnavailable)
				w.Write([]byte(`{}`))
			}))
			defer server.Close()
			implementation := &TokenResource{client: client.NewClient(server.URL, "fixture-token", client.WithMaxRetries(0))}
			state := tokenDeleteState(t, implementation)
			response := &resource.CreateResponse{State: tfsdk.State{Schema: state.Schema}}
			ctx := context.Background()
			implementation.Create(ctx, resource.CreateRequest{Plan: tfsdk.Plan(state)}, response)
			if !response.Diagnostics.HasError() {
				t.Fatal("failed lifecycle not reported")
			}
			if kind == "wrong-scope" {
				if posts.Load() != 0 {
					t.Fatal("unsafe scope submitted POST")
				}
				return
			}
			if posts.Load() != 1 {
				t.Fatal("original create replayed or absent")
			}
			var retained TokenResourceModel
			if d := response.State.Get(ctx, &retained); d.HasError() {
				t.Fatal(d)
			}
			if retained.ID.ValueString() != "fixture" {
				t.Fatal("successful create lost cleanup identity")
			}
			if kind == "missing-uid" {
				if !retained.CreationIdentitySHA256.IsNull() {
					t.Fatal("missing UID fabricated")
				}
			} else if !retained.CreationIdentitySHA256.Equal(creationIdentityReceipt("original-uid")) {
				t.Fatal("original UID discarded")
			}
			if kind != "followup-failure" && gets.Load() != 0 {
				t.Fatal("capture failure continued to GET")
			}
			if kind == "followup-failure" && gets.Load() != 1 {
				t.Fatal("expected original followup")
			}
		})
	}
}

// verifyOriginalCreationGeneratedLifecycle compiles the authored generator's
// Create method against the real provider and HTTP client. Source PRs do not
// commit generated resources; the isolated Go overlay keeps this regression
// effective before On Merge produces the generated delivery.
func verifyOriginalCreationGeneratedLifecycle(t *testing.T) {
	t.Helper()
	repository, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	specDir := os.Getenv("XCSH_SPEC_DIR")
	if specDir == "" {
		specDir = filepath.Join(repository, "docs", "specifications", "api")
	}
	catalog, err := openapi.ParseOperationCatalogFromDir(specDir)
	if err != nil {
		t.Fatal(err)
	}
	spec, err := openapi.ParseFile(filepath.Join(specDir, "domains", "users.json"))
	if err != nil {
		t.Fatal(err)
	}
	resolved, err := catalog.ResolveResource(spec, "token")
	if err != nil {
		t.Fatal(err)
	}
	subject, err := schema.ExtractResourceSchema(spec, "token", func(*openapi.Spec, string) (string, string, bool) {
		return resolved.CollectionPath, resolved.ItemPath, resolved.HasNamespace
	})
	if err != nil {
		t.Fatal(err)
	}
	if resolved.Replace == nil {
		schema.ForceReplaceForCreateDeleteOnly(subject.Attributes)
		subject.UpdateDisabled = true
	}
	schema.RefreshResourcePlanModifierUsage(subject)
	directory := t.TempDir()
	if err = codegen.GenerateResourceFile(subject, directory); err != nil {
		t.Fatal(err)
	}
	generated, err := os.ReadFile(filepath.Join(directory, "token_resource.go"))
	if err != nil {
		t.Fatal(err)
	}
	originalPath := filepath.Join(repository, "internal", "provider", "token_resource.go")
	original, err := os.ReadFile(originalPath)
	if err != nil {
		t.Fatal(err)
	}
	methodRange := func(raw []byte) (int, int) {
		fset := token.NewFileSet()
		tree, err := parser.ParseFile(fset, "token_resource.go", raw, 0)
		if err != nil {
			t.Fatal(err)
		}
		for _, decl := range tree.Decls {
			method, ok := decl.(*ast.FuncDecl)
			if ok && method.Recv != nil && method.Name.Name == "Create" {
				return fset.Position(method.Pos()).Offset, fset.Position(method.End()).Offset
			}
		}
		t.Fatal("generated original Create method missing")
		return 0, 0
	}
	generatedStart, generatedEnd := methodRange(generated)
	originalStart, originalEnd := methodRange(original)
	replacement := append(append(append([]byte{}, original[:originalStart]...), generated[generatedStart:generatedEnd]...), original[originalEnd:]...)
	overlaid := filepath.Join(directory, "token-original-create.go")
	if err = os.WriteFile(overlaid, replacement, 0600); err != nil {
		t.Fatal(err)
	}
	raw, err := json.Marshal(map[string]any{"Replace": map[string]string{originalPath: overlaid}})
	if err != nil {
		t.Fatal(err)
	}
	overlay := filepath.Join(directory, "overlay.json")
	if err = os.WriteFile(overlay, raw, 0600); err != nil {
		t.Fatal(err)
	}
	command := exec.Command("go", "test", "-race", "-overlay="+overlay, "./internal/provider", "-run", "^TestTokenCreateOriginalCaptureLifecycle$", "-count=1")
	command.Dir = repository
	command.Env = append(os.Environ(), "XCSH_ORIGINAL_CREATION_GENERATED_CHILD=1")
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("generated original Create lifecycle failed: %v\n%s", err, output)
	}
}
