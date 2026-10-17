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

// Execute generated Delete methods with real framework state and diagnostics.
// Error text is untrusted: only typed HTTP status can establish absence.
func TestGeneratedDeleteFailuresPreserveState(t *testing.T) {
	for _, title := range []string{"Synthetic", "Namespace", "AlertPolicy"} {
		t.Run(title, func(t *testing.T) {
			generated := t.TempDir()
			r := &openapi.ResourceTemplate{Name: "synthetic", TitleCase: title, HasNamespaceInPath: true}
			if err := GenerateResourceFile(r, generated); err != nil {
				t.Fatal(err)
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
			var method string
			for _, declaration := range parsed.Decls {
				function, ok := declaration.(*ast.FuncDecl)
				if ok && function.Name.Name == "Delete" {
					method = string(text[fset.Position(function.Pos()).Offset:fset.Position(function.End()).Offset])
				}
			}
			if method == "" {
				t.Fatal("generated Delete method unavailable")
			}
			method = strings.NewReplacer(title+"Resource", "SyntheticResource", "Delete"+title, "DeleteSynthetic", "Get"+title, "GetSynthetic").Replace(method)
			imports := ""
			if strings.Contains(method, "strings.") {
				imports = "\n\"strings\""
			}
			root := repoRootFromTest(t)
			dir, err := os.MkdirTemp(filepath.Join(root, "internal"), "zz_deletefailure_")
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { _ = os.RemoveAll(dir) })
			source := `package probe
import (
 "context"
 "errors"
 "fmt"
 "testing"
 "time"
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
type SyntheticResourceModel struct {
 ID types.String ` + "`tfsdk:\"id\"`" + `
 Name types.String ` + "`tfsdk:\"name\"`" + `
 Namespace types.String ` + "`tfsdk:\"namespace\"`" + `
 Timeouts timeouts.Value ` + "`tfsdk:\"timeouts\"`" + `
}
type failingClient struct {err,verifyErr error; calls int; cancel context.CancelFunc}
func (c *failingClient) DeleteSynthetic(_ context.Context, namespace,name string) error {
 if namespace!="demo" || name!="owned" {panic("wrong exact delete identity")}
 c.calls++
 if client.HasHTTPStatus(c.err,400) {c.cancel()}
 return c.err
}
func (c *failingClient) CascadeDeleteSynthetic(ctx context.Context,name string) error {
 return c.DeleteSynthetic(ctx,"demo",name)
}
func (c *failingClient) GetSynthetic(_ context.Context,namespace,name string) (interface{},error) {
 if namespace!="demo" || name!="owned" {panic("wrong exact absence read identity")}
 return nil,c.verifyErr
}
type SyntheticResource struct {client *failingClient}
` + method + `
func TestDeleteStatus(t *testing.T) {
 status:=func(code int,message string) error {return &xcsherrors.XCSHError{StatusCode:code,Message:message}}
 for _,tc:=range []struct{name string;err,verifyErr error;success bool;alertOnly bool}{
  {"deleted",nil,nil,true,false},
  {"typed_absence",status(404,"missing"),nil,true,false},
  {"wrapped_absence",fmt.Errorf("wrapped: %w",status(404,"missing")),nil,true,false},
  {"unimplemented",status(501,"404 NOT_FOUND BAD_REQUEST 400"),status(501,"404 NOT_FOUND"),false,false},
  {"server_error",status(500,"404 NOT_FOUND"),status(500,"404 NOT_FOUND"),false,false},
  {"unauthorized",status(401,"404 NOT_FOUND"),status(401,"404 NOT_FOUND"),false,false},
  {"forbidden",status(403,"404 NOT_FOUND"),status(403,"404 NOT_FOUND"),false,false},
  {"unavailable",status(503,"404 NOT_FOUND"),status(503,"404 NOT_FOUND"),false,false},
  {"transport",errors.New("404 NOT_FOUND BAD_REQUEST 400"),errors.New("404 NOT_FOUND"),false,false},
  {"cancelled_referential_retry",status(400,"404 NOT_FOUND 501"),nil,false,false},
  {"wrapped_cancelled_retry",fmt.Errorf("wrapped: %w",status(400,"404 NOT_FOUND 501")),nil,false,false},
  {"confirmed_absence_after_error",status(500,"failed"),status(404,"missing"),true,true},
  {"wrapped_confirmed_absence",status(500,"failed"),fmt.Errorf("wrapped: %w",status(404,"missing")),true,true},
  {"unverified_absence",status(500,"failed"),errors.New("404 NOT_FOUND"),false,true},
 } {
  if tc.alertOnly && ` + fmtBool(title == "AlertPolicy") + `==false {continue}
  t.Run(tc.name,func(t *testing.T) {
   ctx,cancel:=context.WithCancel(context.Background());defer cancel()
   timeoutTypes:=map[string]attr.Type{"create":types.StringType,"read":types.StringType,"update":types.StringType,"delete":types.StringType}
   model:=SyntheticResourceModel{ID:types.StringValue("owned"),Name:types.StringValue("owned"),Namespace:types.StringValue("demo"),Timeouts:timeouts.Value{Object:types.ObjectNull(timeoutTypes)}}
   state:=tfsdk.State{Schema:schema.Schema{Attributes:map[string]schema.Attribute{"id":schema.StringAttribute{Computed:true},"name":schema.StringAttribute{Required:true},"namespace":schema.StringAttribute{Required:true}},Blocks:map[string]schema.Block{"timeouts":inttimeouts.Block(inttimeouts.DefaultConfig())}}}
   if diagnostics:=state.Set(ctx,&model);diagnostics.HasError(){t.Fatal(diagnostics)}
   c:=&failingClient{err:tc.err,verifyErr:tc.verifyErr,cancel:cancel}
   response:=&resource.DeleteResponse{State:state}
   (&SyntheticResource{client:c}).Delete(ctx,resource.DeleteRequest{State:state},response)
   if response.Diagnostics.HasError()==tc.success {t.Errorf("delete success=%v, want %v: %v",!response.Diagnostics.HasError(),tc.success,response.Diagnostics)}
   if !response.State.Raw.Equal(state.Raw) {t.Error("failed deletion changed owned state")}
   if c.calls!=1 {t.Errorf("delete called %d times",c.calls)}
  })
 }
}
`
			if err := os.WriteFile(filepath.Join(dir, "delete_test.go"), []byte(source), 0600); err != nil {
				t.Fatal(err)
			}
			command := exec.Command("go", "test", "-count=1", "./internal/"+filepath.Base(dir))
			command.Dir = root
			output, err := command.CombinedOutput()
			if err != nil {
				t.Fatalf("generated Delete regression: %v\n%s", err, output)
			}
		})
	}
}

func fmtBool(value bool) string {
	if value {
		return "true"
	}
	return "false"
}

// Strings is needed by namespace/name imports and JWT site validation, but not
// by a namespace's name-only import or typed delete status checks.
func TestGeneratedResourceStringsImportTracksUsage(t *testing.T) {
	for _, r := range []*openapi.ResourceTemplate{
		{Name: "namespace", TitleCase: "Namespace"},
		{Name: "synthetic", TitleCase: "Synthetic", HasNamespaceInPath: true},
		{Name: "token", TitleCase: "Token"},
	} {
		t.Run(r.Name, func(t *testing.T) {
			dir := t.TempDir()
			if err := GenerateResourceFile(r, dir); err != nil {
				t.Fatal(err)
			}
			parsed, err := parser.ParseFile(token.NewFileSet(), filepath.Join(dir, r.Name+"_resource.go"), nil, 0)
			if err != nil {
				t.Fatal(err)
			}
			imported, used := false, false
			for _, i := range parsed.Imports {
				if i.Path.Value == `"strings"` {
					imported = true
				}
			}
			ast.Inspect(parsed, func(n ast.Node) bool {
				if selector, ok := n.(*ast.SelectorExpr); ok {
					if identifier, ok := selector.X.(*ast.Ident); ok && identifier.Name == "strings" {
						used = true
					}
				}
				return true
			})
			if imported != used {
				t.Errorf("strings imported=%v, used=%v", imported, used)
			}
		})
	}
}
