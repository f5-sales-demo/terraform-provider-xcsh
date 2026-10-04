package docconstraints

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func extract(t *testing.T, imports, expression string) []EnumEvidence {
	t.Helper()
	source := "package fixture\n" + imports + "\nvar validators = " + expression
	f, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}
	value := f.Decls[len(f.Decls)-1].(*ast.GenDecl).Specs[0].(*ast.ValueSpec).Values[0]
	return LiteralEnums(f, value)
}

const officialImport = `import sv "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"`

func TestLiteralEnumValuesAreExactAndDeterministic(t *testing.T) {
	evidence := extract(t, officialImport, "[]any{sv.OneOf(\"IPSEC\", `GRE`, \"IPSEC\", \"\")} ")
	if len(evidence) != 1 || !evidence[0].Complete || !evidence[0].CaseSensitive || !reflect.DeepEqual(evidence[0].Values, []string{"", "GRE", "IPSEC"}) {
		t.Fatalf("wrong evidence: %#v", evidence)
	}
	reordered := extract(t, officialImport, "[]any{sv.OneOf(\"\", \"GRE\", \"IPSEC\")} ")
	a, _ := json.Marshal(evidence)
	b, _ := json.Marshal(reordered)
	if string(a) != string(b) {
		t.Fatalf("nondeterministic evidence: %s / %s", a, b)
	}
}
func TestDynamicArgumentsNeverAssertPartialAllowedValues(t *testing.T) {
	for _, expression := range []string{`[]any{sv.OneOf("GRE", dynamic)}`, `[]any{sv.OneOf(values...)}`, `[]any{sv.OneOf()}`, `[]any{sv.OneOf("GRE"+"IPSEC")}`} {
		evidence := extract(t, officialImport, expression)
		if len(evidence) != 1 || evidence[0].Complete || len(evidence[0].Values) != 0 {
			t.Fatalf("partial certainty for %s: %#v", expression, evidence)
		}
	}
}
func TestQualifiedImportsAndValidatorIdentity(t *testing.T) {
	if len(extract(t, `import sv "example.com/stringvalidator"`, `[]any{sv.OneOf("GRE")}`)) != 0 {
		t.Fatal("impostor import accepted")
	}
	if len(extract(t, officialImport, `[]any{other.OneOf("GRE"),sv.LengthAtLeast(3)}`)) != 0 {
		t.Fatal("unowned validator accepted")
	}
	evidence := extract(t, officialImport, `[]any{sv.OneOfCaseInsensitive("gre", "IPSEC")}`)
	if len(evidence) != 1 || evidence[0].CaseSensitive || !evidence[0].Complete {
		t.Fatalf("case semantics lost: %#v", evidence)
	}
}

func TestNestedCallsAndShadowedImportDoNotEstablishEnumEvidence(t *testing.T) {
	if len(extract(t, officialImport, `[]any{wrapper(sv.OneOf("GRE"))}`)) != 0 {
		t.Fatal("wrapped call treated as direct validator")
	}
	if len(extract(t, officialImport, `func() any {sv := custom; return []any{sv.OneOf("GRE")}}()`)) != 0 {
		t.Fatal("shadowed import treated as validator")
	}
	source := `package fixture
 import sv "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
 func f(){sv:=custom; _=[]any{sv.OneOf("GRE")}}`
	file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}
	var expression ast.Expr
	ast.Inspect(file, func(n ast.Node) bool {
		if literal, ok := n.(*ast.CompositeLit); ok {
			expression = literal
		}
		return true
	})
	if len(LiteralEnums(file, expression)) != 0 {
		t.Fatal("local receiver shadowing accepted")
	}
}

func TestSchemaRootsRequireTypedResponseAndFrameworkAttachment(t *testing.T) {
	source := `package fixture
 import r "github.com/hashicorp/terraform-plugin-framework/resource"
 import schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
 import fake "example.com/schema"
 func (x *Thing) Schema(resp *r.SchemaResponse){
 local:=map[string]fake.StringAttribute{"wrong":{}}; _=local
 resp.Schema=schema.Schema{Attributes:map[string]schema.Attribute{"right":schema.StringAttribute{}}}
 }
 func Schema(){other.Schema=schema.Schema{}}
 `
	file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}
	roots := SchemaRoots(file)
	if len(roots) != 1 {
		t.Fatalf("roots=%d", len(roots))
	}
	if FrameworkSelector(file, &ast.SelectorExpr{X: ast.NewIdent("fake"), Sel: ast.NewIdent("StringAttribute")}, "Attribute") {
		t.Fatal("unowned schema type accepted")
	}
}
func TestUnknownValidatorCollectionCoverageIsExplicit(t *testing.T) {
	for _, source := range []string{`func() any {return nil}()`, `[]any{wrapper(sv.OneOf("GRE"))}`} {
		file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", "package fixture\n"+officialImport+"\nvar values="+source, 0)
		if err != nil {
			t.Fatal(err)
		}
		expression := file.Decls[len(file.Decls)-1].(*ast.GenDecl).Specs[0].(*ast.ValueSpec).Values[0]
		if EnumCoverage(file, expression) {
			t.Fatal("unresolved validator coverage accepted")
		}
	}
}

func TestUncalledClosureCannotAttachSchema(t *testing.T) {
	source := `package fixture
 import r "github.com/hashicorp/terraform-plugin-framework/resource"
 import s "github.com/hashicorp/terraform-plugin-framework/resource/schema"
 func (x *Thing) Schema(resp *r.SchemaResponse){unused:=func(){resp.Schema=s.Schema{}};_=unused}`
	file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(SchemaRoots(file)) != 0 {
		t.Fatal("uncalled closure schema accepted")
	}
}

func TestConditionalAndSupersededSchemaAssignmentsRemainUnresolved(t *testing.T) {
	for _, body := range []string{
		`if condition {resp.Schema=s.Schema{}}`,
		`if condition {return};resp.Schema=s.Schema{}`,
		`return;resp.Schema=s.Schema{}`,
		`goto done;resp.Schema=s.Schema{};done: return`,
		`resp.Schema=s.Schema{};resp.Schema=s.Schema{}`,
		`resp.Schema=s.Schema{};resp.Schema=helper()`,
		`resp.Schema=s.Schema{};if condition {resp.Schema=s.Schema{}}`,
		`resp.Schema=s.Schema{};var n int;n,resp.Schema=replacement()`,
		`resp.Schema=s.Schema{};if condition {*resp=r.SchemaResponse{Schema:s.Schema{}}}`,
		`resp.Schema=s.Schema{};alias:=resp;alias.Schema=s.Schema{}`,
		`resp.Schema=s.Schema{};var alias = resp;alias.Schema=s.Schema{}`,
		`resp.Schema=s.Schema{};replace(resp)`,
		`resp.Schema=s.Schema{};replace(&resp.Schema)`,
		`resp.Schema=s.Schema{};resp.Schema.Attributes=map[string]s.Attribute{}`,
		`resp.Schema=s.Schema{};if condition {resp.Schema.Attributes=map[string]s.Attribute{}}`,
		`resp.Schema=s.Schema{};func(){resp.Schema=s.Schema{}}()`,
	} {
		source := `package fixture
 import r "github.com/hashicorp/terraform-plugin-framework/resource"
 import s "github.com/hashicorp/terraform-plugin-framework/resource/schema"
 func (x *Thing) Schema(resp *r.SchemaResponse){` + body + `}`
		file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
		if err != nil {
			t.Fatal(err)
		}
		if len(SchemaRoots(file)) != 0 {
			t.Fatalf("ambiguous assignments accepted: %s", body)
		}
	}
}
