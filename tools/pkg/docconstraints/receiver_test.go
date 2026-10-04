package docconstraints

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestReceiverRoleAndResponseMustMatch(t *testing.T) {
	for _, fixture := range []struct {
		assertion, response string
		count               int
	}{
		{`var _ r.Resource=&Thing{}`, `r.SchemaResponse`, 1},
		{``, `r.SchemaResponse`, 0},
		{`var _ r.Resource=&Other{}`, `r.SchemaResponse`, 0},
		{`var _ r.Resource=&Thing{}`, `d.SchemaResponse`, 0},
		{`var _ fake.Resource=&Thing{}`, `r.SchemaResponse`, 0},
	} {
		source := `package fixture
 import r "github.com/hashicorp/terraform-plugin-framework/resource"
 import d "github.com/hashicorp/terraform-plugin-framework/datasource"
 import fake "example.com/resource"
 import s "github.com/hashicorp/terraform-plugin-framework/resource/schema"
 ` + fixture.assertion + `
 func(x *Thing)Schema(resp *` + fixture.response + `){resp.Schema=s.Schema{}}`
		file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
		if err != nil {
			t.Fatal(err)
		}
		if roots := SchemaRoots(file); len(roots) != fixture.count {
			t.Fatalf("assertion %s response %s roots %d", fixture.assertion, fixture.response, len(roots))
		}
	}
}

func TestReceiverContractsRejectConflictsAndUnownedExpressions(t *testing.T) {
	for _, assertion := range []string{
		"var _ r.Resource = &Thing{}; var _ d.DataSource = &Thing{}",
		"var _ d.DataSource = &Thing{}; var _ r.Resource = &Thing{}",
		"var _ r.Resource = &Thing{}; var _ d.DataSource = &Thing{}; var _ r.Resource = &Thing{}",
		"var _ r.Resource = *Thing{}",
	} {
		source := `package fixture
import r "github.com/hashicorp/terraform-plugin-framework/resource"
import d "github.com/hashicorp/terraform-plugin-framework/datasource"
` + assertion
		file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
		if err != nil {
			t.Fatal(err)
		}
		if role := ReceiverContracts(file)["Thing"]; role != "" {
			t.Fatalf("assertion %s resolved as %s", assertion, role)
		}
	}
}

func TestSchemaLiteralMustMatchReceiverRole(t *testing.T) {
	for _, fixture := range []struct {
		schema string
		count  int
	}{
		{"github.com/hashicorp/terraform-plugin-framework/resource/schema", 1},
		{"github.com/hashicorp/terraform-plugin-framework/datasource/schema", 0},
		{"github.com/hashicorp/terraform-plugin-framework/impostor/schema", 0},
		{"example.com/resource/schema", 0},
	} {
		source := `package fixture
import r "github.com/hashicorp/terraform-plugin-framework/resource"
import s "` + fixture.schema + `"
var _ r.Resource=&Thing{}
func(x *Thing)Schema(resp *r.SchemaResponse){resp.Schema=s.Schema{}}
`
		file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
		if err != nil {
			t.Fatal(err)
		}
		if roots := SchemaRoots(file); len(roots) != fixture.count {
			t.Fatalf("schema %s roots %d", fixture.schema, len(roots))
		}
	}
}
