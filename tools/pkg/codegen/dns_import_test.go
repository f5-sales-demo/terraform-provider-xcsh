package codegen

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestGeneratedDNSImportOwnershipConvergence(t *testing.T) {
	field := openapi.TerraformAttribute{Name: "dns_volterra_managed", GoName: "DnsVolterraManaged", TfsdkTag: "dns_volterra_managed", JsonName: "dns_volterra_managed", Type: "bool", Optional: true, PlanModifier: "RequiresReplace"}
	var decoder strings.Builder
	if err := renderUnmarshalChild(&decoder, "HTTPLoadBalancer", "Http", field, "wire", "prior", "prior != nil", "single", "\t"); err != nil {
		t.Fatal(err)
	}
	root := repoRootFromTest(t)
	dir, err := os.MkdirTemp(filepath.Join(root, "internal"), "zz_dnsimport_")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	source := `package probe
import (
 "context"
 "testing"
 "github.com/hashicorp/terraform-plugin-framework/types"
 "github.com/hashicorp/terraform-plugin-framework/resource/schema"
 "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
 "github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
 "github.com/hashicorp/terraform-plugin-framework/tfsdk"
 "github.com/hashicorp/terraform-plugin-go/tftypes"
)
type Model struct { DnsVolterraManaged types.Bool }
func read(prior *Model,wire map[string]interface{},isImport bool) *Model {return &Model{
` + decoder.String() + `} }
func TestOwnership(t *testing.T) {
 ctx:=context.Background()
 nested:=schema.SingleNestedAttribute{
` + RenderNestedAttributes([]openapi.TerraformAttribute{field}, " ") + `}
 modifier:=nested.Attributes["dns_volterra_managed"].(schema.BoolAttribute).PlanModifiers[0]
 s:=schema.Schema{Attributes:nested.Attributes}
 raw:=tftypes.NewValue(s.Type().TerraformType(ctx),map[string]tftypes.Value{"dns_volterra_managed":tftypes.NewValue(tftypes.Bool,false)})
 for _,tc:=range []struct {name string; before,after types.Bool; replace bool}{
  {"import false to explicit false",types.BoolNull(),types.BoolValue(false),false},
  {"explicit false to omitted",types.BoolValue(false),types.BoolNull(),false},
  {"omitted to delegated",types.BoolNull(),types.BoolValue(true),true},
  {"delegated to omitted",types.BoolValue(true),types.BoolNull(),true},
  {"tenant to delegated",types.BoolValue(false),types.BoolValue(true),true},
  {"delegated to tenant",types.BoolValue(true),types.BoolValue(false),true},
  {"unchanged",types.BoolValue(false),types.BoolValue(false),false},
  {"both omitted",types.BoolNull(),types.BoolNull(),false},
  {"unknown ownership",types.BoolValue(false),types.BoolUnknown(),true},
 } {
  t.Run(tc.name,func(t *testing.T){
   req:=planmodifier.BoolRequest{State:tfsdk.State{Schema:s,Raw:raw},Plan:tfsdk.Plan{Schema:s,Raw:raw},StateValue:tc.before,PlanValue:tc.after}
   resp:=&planmodifier.BoolResponse{}
   modifier.PlanModifyBool(ctx,req,resp)
   if resp.RequiresReplace!=tc.replace {t.Fatalf("replace=%v want %v",resp.RequiresReplace,tc.replace)}
  })
 }
 got:=read(&Model{DnsVolterraManaged:types.BoolValue(false)},map[string]interface{}{"dns_volterra_managed":true},false)
 if !got.DnsVolterraManaged.ValueBool() {t.Fatal("DNS ownership drift hidden")}
 got=read(&Model{DnsVolterraManaged:types.BoolNull()},map[string]interface{}{"dns_volterra_managed":false},true)
 if !got.DnsVolterraManaged.IsNull() {t.Fatal("omitted import default no longer suppressed")}
 got=read(&Model{DnsVolterraManaged:types.BoolValue(true)},map[string]interface{}{"dns_volterra_managed":false},false)
 if got.DnsVolterraManaged.IsNull() || got.DnsVolterraManaged.ValueBool() {t.Fatal("tenant ownership drift hidden")}
}
`
	if err := os.WriteFile(filepath.Join(dir, "probe_test.go"), []byte(source), 0600); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "test", dir)
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated DNS convergence: %v\n%s", err, out)
	}
}
