package codegen

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestGeneratedRouteTimeoutDriftReadback(t *testing.T) {
	field := openapi.TerraformAttribute{Name: "timeout", GoName: "Timeout", TfsdkTag: "timeout", JsonName: "timeout", Type: "int64", Optional: true}
	var code strings.Builder
	if err := renderUnmarshalChild(&code, "HTTPLoadBalancer", "RoutesSimpleRouteAdvancedOptions", field, "wire", "prior", "prior != nil", "single", "\t"); err != nil {
		t.Fatal(err)
	}
	root := repoRootFromTest(t)
	dir, err := os.MkdirTemp(filepath.Join(root, "internal"), "zz_routetimeout_")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	source := `package probe
import (
 "testing"
 "github.com/hashicorp/terraform-plugin-framework/types"
)
type Model struct {Timeout types.Int64}
func read(prior *Model,wire map[string]interface{},isImport bool) *Model {return &Model{
` + code.String() + `} }
func TestTimeout(t *testing.T) {
 for _, imported:=range []bool{false,true} {
  got:=read(&Model{Timeout:types.Int64Value(10)},map[string]interface{}{"timeout":float64(35)},imported)
  if got.Timeout.ValueInt64()!=35 {t.Fatal("configured timeout drift hidden")}
  got=read(&Model{Timeout:types.Int64Value(35)},map[string]interface{}{"timeout":float64(10)},imported)
  if got.Timeout.ValueInt64()!=10 {t.Fatal("restored timeout not observed")}
 }
 got:=read(&Model{Timeout:types.Int64Unknown()},map[string]interface{}{"timeout":float64(35)},false)
 if got.Timeout.IsUnknown() || got.Timeout.ValueInt64()!=35 {t.Fatal("unknown timeout unresolved")}
 got=read(&Model{Timeout:types.Int64Null()},map[string]interface{}{"timeout":float64(0)},false)
 if !got.Timeout.IsNull() {t.Fatal("omitted default became explicit")}
 for _, configured:=range []int64{0,10} {
  got=read(&Model{Timeout:types.Int64Value(configured)},map[string]interface{}{"timeout":float64(0)},false)
  if got.Timeout.IsNull() || got.Timeout.ValueInt64()!=0 {t.Fatal("infinite timeout drift or explicit zero hidden")}
 }
 got=read(&Model{Timeout:types.Int64Null()},map[string]interface{}{"timeout":float64(35)},false)
 if got.Timeout.ValueInt64()!=35 {t.Fatal("non-default drift from omitted timeout hidden")}
 got=read(&Model{Timeout:types.Int64Value(10)},map[string]interface{}{},false)
 if !got.Timeout.IsNull() {t.Fatal("removed timeout drift hidden")}
}
`
	if err := os.WriteFile(filepath.Join(dir, "probe_test.go"), []byte(source), 0600); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "test", dir)
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated timeout decoder: %v\n%s", err, out)
	}
}
