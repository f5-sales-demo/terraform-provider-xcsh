package codegen

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestGeneratedRoutePriorityDefaultReadback(t *testing.T) {
	field := openapi.TerraformAttribute{Name: "priority", GoName: "Priority", TfsdkTag: "priority", JsonName: "priority", Type: "string", Optional: true}
	var code strings.Builder
	if err := renderUnmarshalChild(&code, "HTTPLoadBalancer", "RoutesSimpleRouteAdvancedOptions", field, "wire", "prior", "prior != nil", "single", "\t"); err != nil {
		t.Fatal(err)
	}
	root := repoRootFromTest(t)
	dir, err := os.MkdirTemp(filepath.Join(root, "internal"), "zz_routepriority_")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	source := `package probe
import (
 "testing"
 "github.com/hashicorp/terraform-plugin-framework/types"
)
type Model struct { Priority types.String }
func read(prior *Model, wire map[string]interface{}, isImport bool) *Model { return &Model{
` + code.String() + `} }
func TestPriority(t *testing.T) {
 for _, imported := range []bool{false,true} {
  got:=read(&Model{Priority:types.StringNull()},map[string]interface{}{"priority":"DEFAULT"},imported)
  if !got.Priority.IsNull() {t.Fatal("omitted DEFAULT became configured intent")}
  got=read(&Model{Priority:types.StringNull()},map[string]interface{}{"priority":"HIGH"},imported)
  if got.Priority.ValueString()!="HIGH" {t.Fatal("non-default priority drift hidden")}
 }
 for _, value:=range []string{"DEFAULT","HIGH"} {
  got:=read(&Model{Priority:types.StringValue(value)},map[string]interface{}{"priority":"DEFAULT"},false)
  if got.Priority.ValueString()!="DEFAULT" {t.Fatal("explicit priority or drift suppressed")}
 }
 got:=read(&Model{Priority:types.StringUnknown()},map[string]interface{}{"priority":"DEFAULT"},false)
 if got.Priority.IsUnknown() || got.Priority.ValueString()!="DEFAULT" {t.Fatal("unknown planned priority unresolved")}
}
`
	if err := os.WriteFile(filepath.Join(dir, "probe_test.go"), []byte(source), 0o600); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "test", dir)
	cmd.Dir = root
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated priority decoder: %v\n%s", err, output)
	}
	var authoritative strings.Builder
	if err := renderUnmarshalChild(&authoritative, "HTTPLoadBalancer", "RoutesSimpleRouteAdvancedOptions", field, "wire", "prior", "prior != nil", "single", "\t", true); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(authoritative.String(), `v == "DEFAULT"`) {
		t.Fatal("data source suppresses live default priority")
	}
}
