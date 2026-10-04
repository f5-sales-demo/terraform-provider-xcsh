package docconstraints

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
	"testing"
)

func TestWalkFollowsOnlySchemaBearingMembers(t *testing.T) {
	source := `package fixture
 import r "github.com/hashicorp/terraform-plugin-framework/resource"
 import s "github.com/hashicorp/terraform-plugin-framework/resource/schema"
 func (x *Thing) Schema(resp *r.SchemaResponse){resp.Schema=s.Schema{
 Attributes:map[string]s.Attribute{
 "choice":s.StringAttribute{Validators:[]any{helper(map[string]s.Attribute{"fake":s.StringAttribute{}})}},
 "nested":&s.ListNestedAttribute{NestedObject:s.NestedAttributeObject{Attributes:map[string]s.Attribute{"actual":s.StringAttribute{}}}},
 },
 Blocks:map[string]s.Block{"block":s.SingleNestedBlock{Attributes:map[string]s.Attribute{"child":s.StringAttribute{}}}},
 Description:helper(map[string]s.Attribute{"unattached":s.StringAttribute{}}),
 }}`
	file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}
	roots := SchemaRoots(file)
	if len(roots) != 1 {
		t.Fatal("missing root")
	}
	paths := []string{}
	err = WalkSchemaFields(file, roots[0], func(path []string, _ *ast.CompositeLit) error {
		paths = append(paths, strings.Join(path, "."))
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(paths, []string{"choice", "nested", "nested.actual", "block", "block.child"}) {
		t.Fatalf("unattached fields included: %#v", paths)
	}
}
func TestWalkDoesNotInspectIndirectSchemaArguments(t *testing.T) {
	source := `package fixture
 import r "github.com/hashicorp/terraform-plugin-framework/resource"
 import s "github.com/hashicorp/terraform-plugin-framework/resource/schema"
 func (x *Thing) Schema(resp *r.SchemaResponse){resp.Schema=s.Schema{Attributes:helper(map[string]s.Attribute{"argument":s.StringAttribute{}})}}`
	file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	err = WalkSchemaFields(file, SchemaRoots(file)[0], func(_ []string, _ *ast.CompositeLit) error { count++; return nil })
	if err != nil || count != 0 {
		t.Fatalf("helper arguments became schema evidence count=%d err=%v", count, err)
	}
}
