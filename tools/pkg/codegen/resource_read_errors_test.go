package codegen

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

// Execute the generated error path using real framework state and diagnostics.
// Cached state is retained on unavailable reads, but never accepted as refreshed.
func TestGeneratedReadFailuresPreserveState(t *testing.T) {
	for _, action := range []bool{false, true} {
		t.Run(map[bool]string{false: "resource", true: "action"}[action], func(t *testing.T) {
			testGeneratedReadFailuresPreserveState(t, action)
		})
	}
}

func testGeneratedReadFailuresPreserveState(t *testing.T, action bool) {
	generated := t.TempDir()
	r := &openapi.ResourceTemplate{Name: "synthetic", TitleCase: "Synthetic", HasNamespaceInPath: true, IsAction: action, ReadObjectPath: "/api/register/namespaces/%s/registrations/%s"}
	var generateErr error
	if action {
		generateErr = GenerateActionResource(r, generated, generated)
	} else {
		generateErr = GenerateResourceFile(r, generated)
	}
	if generateErr != nil {
		t.Fatal(generateErr)
	}
	text, err := os.ReadFile(filepath.Join(generated, "synthetic_resource.go"))
	if err != nil {
		t.Fatal(err)
	}
	fset := token.NewFileSet()
	parsed, err := parser.ParseFile(fset, "generated.go", text, 0)
	if err != nil {
		t.Fatal(err)
	}
	var prefix string
	for _, declaration := range parsed.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok || function.Name.Name != "Read" {
			continue
		}
		for _, statement := range function.Body.List {
			branch, ok := statement.(*ast.IfStmt)
			if !ok {
				continue
			}
			condition := string(text[fset.Position(branch.Cond.Pos()).Offset:fset.Position(branch.Cond.End()).Offset])
			if condition == "err != nil" {
				prefix = string(text[fset.Position(function.Pos()).Offset:fset.Position(branch.End()).Offset]) + "\n}\n"
				break
			}
		}
	}
	if prefix == "" {
		t.Fatal("generated Read error branch unavailable")
	}
	prefix = strings.Replace(prefix, "apiResource, err :=", "_, err :=", 1)
	root := repoRootFromTest(t)
	dir, err := os.MkdirTemp(filepath.Join(root, "internal"), "zz_readfailure_")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	imports := ""
	if strings.Contains(prefix, "strings.") {
		imports += "\n\"strings\""
	}
	source := `package probe
import (
 "context"
 "fmt"
 "errors"
 "testing"
 "github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
 xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
 inttimeouts "github.com/f5-sales-demo/terraform-provider-xcsh/internal/timeouts"
 "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
 "github.com/hashicorp/terraform-plugin-framework/attr"
 "github.com/hashicorp/terraform-plugin-framework/resource"
 "github.com/hashicorp/terraform-plugin-framework/resource/schema"
 "github.com/hashicorp/terraform-plugin-framework/tfsdk"
 "github.com/hashicorp/terraform-plugin-framework/types"
 "github.com/hashicorp/terraform-plugin-log/tflog"
` + imports + `
)
var _ = client.HasHTTPStatus
type SyntheticResourceModel struct {
 ID types.String ` + "`tfsdk:\"id\"`" + `
 Name types.String ` + "`tfsdk:\"name\"`" + `
 Namespace types.String ` + "`tfsdk:\"namespace\"`" + `
 Timeouts timeouts.Value ` + "`tfsdk:\"timeouts\"`" + `
}
type failingClient struct {err error; calls int}
func (c *failingClient) GetSynthetic(_ context.Context, namespace, name string) (interface{}, error) {
 if namespace != "demo" || name != "owned" {panic("wrong exact read identity")}
 c.calls++
 return nil,c.err
}
func (c *failingClient) GetLenient(_ context.Context, path string, _ interface{}) error {
 if path != "/api/register/namespaces/demo/registrations/owned" {panic("wrong exact action read identity")}
 c.calls++
 return c.err
}
type SyntheticResource struct {client *failingClient}
` + prefix + `
func TestUnavailableRead(t *testing.T) {
 for _,tc:=range []struct{name string;err error;absent bool}{
  {"unimplemented", &xcsherrors.XCSHError{StatusCode:501,Message:"unimplemented"}, false},
  {"misleading_unimplemented", &xcsherrors.XCSHError{StatusCode:501,Message:"404 NOT_FOUND"}, false},
  {"unauthorized", &xcsherrors.XCSHError{StatusCode:401,Message:"denied"}, false},
  {"forbidden", &xcsherrors.XCSHError{StatusCode:403,Message:"denied"}, false},
  {"server_error", &xcsherrors.XCSHError{StatusCode:500,Message:"404 NOT_FOUND 501"}, false},
  {"unavailable", &xcsherrors.XCSHError{StatusCode:503,Message:"unavailable"}, false},
  {"transport", errors.New("transport error mentions 404 NOT_FOUND"), false},
  {"verified_absence", &xcsherrors.XCSHError{StatusCode:404,Message:"missing"}, true},
  {"wrapped_absence", fmt.Errorf("wrapped: %w", &xcsherrors.XCSHError{StatusCode:404,Message:"missing"}), true},
 } {
  t.Run(tc.name,func(t *testing.T) {
   ctx:=context.Background()
   timeoutTypes:=map[string]attr.Type{"create":types.StringType,"read":types.StringType,"update":types.StringType,"delete":types.StringType}
   model:=SyntheticResourceModel{ID:types.StringValue("owned"),Name:types.StringValue("owned"),Namespace:types.StringValue("demo"),Timeouts:timeouts.Value{Object:types.ObjectNull(timeoutTypes)}}
   state:=tfsdk.State{Schema:schema.Schema{Attributes:map[string]schema.Attribute{"id":schema.StringAttribute{Computed:true},"name":schema.StringAttribute{Required:true},"namespace":schema.StringAttribute{Required:true}},Blocks:map[string]schema.Block{"timeouts":inttimeouts.Block(inttimeouts.DefaultConfig())}}}
   if diagnostics:=state.Set(ctx,&model);diagnostics.HasError(){t.Fatal(diagnostics)}
   c:=&failingClient{err:tc.err}
   r:=&SyntheticResource{client:c}
   response:=&resource.ReadResponse{State:state}
   r.Read(ctx,resource.ReadRequest{State:state},response)
   if response.Diagnostics.HasError()!=!tc.absent {t.Errorf("unavailable read accepted cached state: %v",response.Diagnostics)}
   if response.State.Raw.IsNull()!=tc.absent {t.Errorf("state removal=%v, want %v",response.State.Raw.IsNull(),tc.absent)}
   if !tc.absent && !response.State.Raw.Equal(state.Raw) {t.Error("unverified read changed owned state")}
   if c.calls!=1 {t.Errorf("read called %d times",c.calls)}
  })
 }
}
`
	if err := os.WriteFile(filepath.Join(dir, "read_test.go"), []byte(source), 0600); err != nil {
		t.Fatal(err)
	}
	command := exec.Command("go", "test", "-count=1", "./internal/"+filepath.Base(dir))
	command.Dir = root
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("generated Read regression: %v\n%s", err, output)
	}
}
