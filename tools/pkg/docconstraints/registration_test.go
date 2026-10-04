package docconstraints

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestDeclaredRegistrationsBindDirectFactoryToAssertedRole(t *testing.T) {
	for _, fixture := range []struct {
		body, factory string
		want          bool
	}{
		{`return []func() r.Resource{NewThing}`, `return &Thing{}`, true},
		{`return []func() r.Resource{NewOther}`, `return &Thing{}`, false},
		{`if enabled {return []func() r.Resource{NewThing}}; return nil`, `return &Thing{}`, false},
		{`return []func() r.Resource{NewThing}`, `if enabled {return &Thing{}}; return &Other{}`, false},
		{`return []func() r.Resource{NewThing}`, `return helper()`, false},
		{`return []func() d.DataSource{NewThing}`, `return &Thing{}`, false},
	} {
		source := `package fixture
import p "github.com/hashicorp/terraform-plugin-framework/provider"
import r "github.com/hashicorp/terraform-plugin-framework/resource"
import d "github.com/hashicorp/terraform-plugin-framework/datasource"
var _ p.Provider = &Provider{}
var _ r.Resource = &Thing{}
func(x *Provider)Resources() []func() r.Resource {` + fixture.body + `}
func NewThing() r.Resource {` + fixture.factory + `}
`
		file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
		if err != nil {
			t.Fatal(err)
		}
		registrations := DeclaredRegistrations(map[string]*ast.File{"fixture.go": file})
		if got := registrations["Thing"]; (got == "github.com/hashicorp/terraform-plugin-framework/resource") != fixture.want {
			t.Fatalf("body=%s factory=%s registration=%q", fixture.body, fixture.factory, got)
		}
	}
}
func TestDeclaredRegistrationsRejectUnregisteredAndDuplicateFactories(t *testing.T) {
	sources := map[string]string{
		"provider.go": `package fixture
import p "github.com/hashicorp/terraform-plugin-framework/provider"
import r "github.com/hashicorp/terraform-plugin-framework/resource"
var _ p.Provider = &Provider{}
func(x *Provider)Resources() []func() r.Resource {return []func() r.Resource{NewThing}}
`,
		"thing.go": `package fixture
import r "github.com/hashicorp/terraform-plugin-framework/resource"
var _ r.Resource=&Thing{}
func NewThing() r.Resource {return &Thing{}}
`,
	}
	files := map[string]*ast.File{}
	for name, source := range sources {
		file, err := parser.ParseFile(token.NewFileSet(), name, source, 0)
		if err != nil {
			t.Fatal(err)
		}
		files[name] = file
	}
	if DeclaredRegistrations(files)["Thing"] == "" {
		t.Fatal("cross-file factory not bound")
	}
	files["duplicate.go"] = files["thing.go"]
	if DeclaredRegistrations(files)["Thing"] != "" {
		t.Fatal("duplicate factory accepted")
	}
}

func TestRegistrationListShadowedFactoryRemainsUnresolved(t *testing.T) {
	source := `package fixture
import p "github.com/hashicorp/terraform-plugin-framework/provider"
import r "github.com/hashicorp/terraform-plugin-framework/resource"
var _ p.Provider=&Provider{}
var _ r.Resource=&Thing{}
func(x *Provider)Resources(NewThing func() r.Resource) []func() r.Resource {return []func() r.Resource{NewThing}}
func NewThing() r.Resource {return &Thing{}}
`
	file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}
	if DeclaredRegistrations(map[string]*ast.File{"fixture.go": file})["Thing"] != "" {
		t.Fatal("parameter shadow accepted")
	}
}
func TestRegisteredSchemaRootsRejectUnlistedReceiver(t *testing.T) {
	source := `package fixture
import r "github.com/hashicorp/terraform-plugin-framework/resource"
import s "github.com/hashicorp/terraform-plugin-framework/resource/schema"
var _ r.Resource=&Thing{}
func(x *Thing)Schema(resp *r.SchemaResponse){resp.Schema=s.Schema{}}
`
	file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(RegisteredSchemaRoots(file, nil)) != 0 {
		t.Fatal("unregistered root exported")
	}
	if len(RegisteredSchemaRoots(file, map[string]string{"Thing": "github.com/hashicorp/terraform-plugin-framework/resource"})) != 1 {
		t.Fatal("registered root lost")
	}
}

func TestSplitFileFactoryRetainsSchemaDestination(t *testing.T) {
	sources := map[string]string{
		"provider.go": `package fixture
import p "github.com/hashicorp/terraform-plugin-framework/provider"
import r "github.com/hashicorp/terraform-plugin-framework/resource"
var _ p.Provider=&Provider{}
`,
		"registrations.go": `package fixture
import r "github.com/hashicorp/terraform-plugin-framework/resource"
func(x *Provider)Resources() []func() r.Resource {return []func() r.Resource{NewThing}}
`,
		"constructors.go": `package fixture
import r "github.com/hashicorp/terraform-plugin-framework/resource"
func NewThing() r.Resource {return &Thing{}}
`,
		"thing.go": `package fixture
import r "github.com/hashicorp/terraform-plugin-framework/resource"
import s "github.com/hashicorp/terraform-plugin-framework/resource/schema"
var _ r.Resource=&Thing{}
func(x *Thing)Schema(resp *r.SchemaResponse){resp.Schema=s.Schema{Attributes:map[string]s.Attribute{"protocol":s.StringAttribute{}}}}
`,
	}
	files := map[string]*ast.File{}
	for name, source := range sources {
		file, err := parser.ParseFile(token.NewFileSet(), name, source, 0)
		if err != nil {
			t.Fatal(err)
		}
		files[name] = file
	}
	roots := RegisteredSchemaRoots(files["thing.go"], DeclaredRegistrations(files))
	if len(roots) != 1 {
		t.Fatal("split-file root lost")
	}
	count := 0
	err := WalkSchemaFields(files["thing.go"], roots[0], func(path []string, value *ast.CompositeLit) error {
		if path[0] == "protocol" {
			count++
		}
		return nil
	})
	if err != nil || count != 1 {
		t.Fatalf("split-file field lost count=%d err=%v", count, err)
	}
}

func TestAllDirectRegistrationRoles(t *testing.T) {
	for method, role := range registrationMethods {
		iface := frameworkInterfaces[role]
		source := `package fixture
import p "github.com/hashicorp/terraform-plugin-framework/provider"
import r "` + role + `"
var _ p.Provider=&Provider{}
var _ r.` + iface + `=&Thing{}
func(x *Provider)` + method + `() []func() r.` + iface + ` {return []func() r.` + iface + `{NewThing}}
func NewThing() r.` + iface + ` {return &Thing{}}
`
		file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
		if err != nil {
			t.Fatal(err)
		}
		if DeclaredRegistrations(map[string]*ast.File{"fixture.go": file})["Thing"] != role {
			t.Fatalf("role lost %s", role)
		}
	}
}
func TestProviderSchemaRetentionIsDeclaredEvidence(t *testing.T) {
	source := `package fixture
import p "github.com/hashicorp/terraform-plugin-framework/provider"
import s "github.com/hashicorp/terraform-plugin-framework/provider/schema"
var _ p.Provider=&Provider{}
func(x *Provider)Schema(resp *p.SchemaResponse){resp.Schema=s.Schema{}}
`
	file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(RegisteredSchemaRoots(file, nil)) != 1 {
		t.Fatal("provider schema lost")
	}
}

func TestSplitAssertionsRetainResourceAndProviderSchemaRoots(t *testing.T) {
	sources := map[string]string{
		"assertions.go": `package fixture
import p "github.com/hashicorp/terraform-plugin-framework/provider"
import r "github.com/hashicorp/terraform-plugin-framework/resource"
var _ p.Provider=&Provider{}
var _ r.Resource=&Thing{}
`,
		"methods.go": `package fixture
import p "github.com/hashicorp/terraform-plugin-framework/provider"
import r "github.com/hashicorp/terraform-plugin-framework/resource"
import s "github.com/hashicorp/terraform-plugin-framework/resource/schema"
import ps "github.com/hashicorp/terraform-plugin-framework/provider/schema"
func(x *Provider)Resources() []func() r.Resource {return []func() r.Resource{NewThing}}
func NewThing() r.Resource {return &Thing{}}
func(x *Thing)Schema(resp *r.SchemaResponse){resp.Schema=s.Schema{}}
func(x *Provider)Schema(resp *p.SchemaResponse){resp.Schema=ps.Schema{}}
`,
	}
	files := map[string]*ast.File{}
	for name, source := range sources {
		file, err := parser.ParseFile(token.NewFileSet(), name, source, 0)
		if err != nil {
			t.Fatal(err)
		}
		files[name] = file
	}
	roots := RegisteredSchemaRootsWithContracts(files["methods.go"], DeclaredRegistrations(files), PackageReceiverContracts(files))
	if len(roots) != 2 {
		t.Fatalf("split assertion roots=%d", len(roots))
	}
}
