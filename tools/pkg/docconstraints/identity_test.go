package docconstraints

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestMetadataIdentityRequiresExactRoleAndTypedReceiver(t *testing.T) {
	for _, fixture := range []struct {
		body, response string
		want           string
	}{
		{`resp.TypeName=req.ProviderTypeName+"_thing"`, "r.MetadataResponse", "thing_resource.go"},
		{`resp.TypeName=req.ProviderTypeName+"_thing"`, "d.MetadataResponse", ""},
		{`if condition {resp.TypeName=req.ProviderTypeName+"_thing"}`, "r.MetadataResponse", ""},
		{`resp.TypeName=helper()`, "r.MetadataResponse", ""},
		{`resp.TypeName=req.ProviderTypeName+"_thing";resp.TypeName=req.ProviderTypeName+"_other"`, "r.MetadataResponse", ""},
		{`resp.TypeName=req.ProviderTypeName+"../thing"`, "r.MetadataResponse", ""},
	} {
		source := `package fixture
import r "github.com/hashicorp/terraform-plugin-framework/resource"
import d "github.com/hashicorp/terraform-plugin-framework/datasource"
func(x *Thing)Metadata(req r.MetadataRequest,resp *` + fixture.response + `){` + fixture.body + `}
`
		file, err := parser.ParseFile(token.NewFileSet(), "arbitrary.go", source, 0)
		if err != nil {
			t.Fatal(err)
		}
		got := DeclaredDocumentSources(map[string]*ast.File{"arbitrary.go": file}, map[string]string{"Thing": "github.com/hashicorp/terraform-plugin-framework/resource"})["Thing"]
		if got != fixture.want {
			t.Fatalf("body %s source=%q expected=%q", fixture.body, got, fixture.want)
		}
	}
}

func TestAllRoleDestinations(t *testing.T) {
	for method, role := range registrationMethods {
		iface := frameworkInterfaces[role]
		source := `package fixture
import p "github.com/hashicorp/terraform-plugin-framework/provider"
import r "` + role + `"
import s "` + role + `/schema"
var _ p.Provider=&Provider{}
var _ r.` + iface + `=&Thing{}
func(x *Provider)` + method + `() []func() r.` + iface + ` {return []func() r.` + iface + `{NewThing}}
func NewThing() r.` + iface + ` {return &Thing{}}
func(x *Thing)Metadata(req r.MetadataRequest,resp *r.MetadataResponse){resp.TypeName=req.ProviderTypeName+"_thing"}
func(x *Thing)Schema(resp *r.SchemaResponse){resp.Schema=s.Schema{Attributes:map[string]s.Attribute{"protocol":s.StringAttribute{}}}}
`
		file, err := parser.ParseFile(token.NewFileSet(), "arbitrary.go", source, 0)
		if err != nil {
			t.Fatal(err)
		}
		files := map[string]*ast.File{"arbitrary.go": file}
		contracts := PackageReceiverContracts(files)
		roots := RegisteredSchemaRootsWithContracts(file, DeclaredRegistrations(files), contracts)
		if len(roots) != 1 {
			t.Fatalf("role %s roots=%d", role, len(roots))
		}
		receiver := SchemaRootReceiver(file, roots[0])
		destination := DeclaredDocumentSources(files, contracts)[receiver]
		if destination != "thing_"+documentSuffix[role]+".go" {
			t.Fatalf("role %s destination=%s", role, destination)
		}
		count := 0
		if err := WalkSchemaFields(file, roots[0], func(path []string, value *ast.CompositeLit) error { count++; return nil }); err != nil || count != 1 {
			t.Fatalf("role %s fields=%d err=%v", role, count, err)
		}
	}
}
func TestDuplicateMetadataDestinationsRemainUnresolved(t *testing.T) {
	source := `package fixture
import r "github.com/hashicorp/terraform-plugin-framework/resource"
func(x *First)Metadata(req r.MetadataRequest,resp *r.MetadataResponse){resp.TypeName=req.ProviderTypeName+"_thing"}
func(x *Second)Metadata(req r.MetadataRequest,resp *r.MetadataResponse){resp.TypeName=req.ProviderTypeName+"_thing"}
func(x *Third)Metadata(req r.MetadataRequest,resp *r.MetadataResponse){resp.TypeName=req.ProviderTypeName+"_thing"}
`
	file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}
	role := "github.com/hashicorp/terraform-plugin-framework/resource"
	sources := DeclaredDocumentSources(map[string]*ast.File{"fixture.go": file}, map[string]string{"First": role, "Second": role, "Third": role})
	for _, source := range sources {
		if source != "" {
			t.Fatal("duplicate destination asserted")
		}
	}
}

func TestUniqueSchemaOwnersRejectAmbiguousProvidersAndDuplicateMethods(t *testing.T) {
	role := "github.com/hashicorp/terraform-plugin-framework/resource"
	providerRole := "github.com/hashicorp/terraform-plugin-framework/provider"
	source := `package fixture
func(x *Thing)Schema(){}
func(x *FirstProvider)Schema(){}
func(x *SecondProvider)Schema(){}
`
	file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}
	contracts := map[string]string{"Thing": role, "FirstProvider": providerRole, "SecondProvider": providerRole}
	owners := UniqueSchemaOwners(map[string]*ast.File{"fixture.go": file}, contracts)
	if !owners["Thing"] || owners["FirstProvider"] || owners["SecondProvider"] {
		t.Fatal("ambiguous providers accepted")
	}
	owners = UniqueSchemaOwners(map[string]*ast.File{"first.go": file, "second.go": file}, contracts)
	if owners["Thing"] {
		t.Fatal("duplicate schema union accepted")
	}
}

func TestCompleteSplitFileDocumentIdentityAllRoles(t *testing.T) {
	for method, role := range registrationMethods {
		iface := frameworkInterfaces[role]
		sources := map[string]string{
			"assertions.go": `package fixture
import p "github.com/hashicorp/terraform-plugin-framework/provider"
import r "` + role + `"
var _ p.Provider=&Provider{}
var _ r.` + iface + `=&Thing{}`,
			"registrations.go": `package fixture
import r "` + role + `"
func(x *Provider)` + method + `() []func() r.` + iface + `{return []func() r.` + iface + `{NewThing}}`,
			"constructors.go": `package fixture
import r "` + role + `"
func NewThing() r.` + iface + `{return &Thing{}}`,
			"metadata.go": `package fixture
import r "` + role + `"
func(x *Thing)Metadata(req r.MetadataRequest,resp *r.MetadataResponse){resp.TypeName=req.ProviderTypeName+"_thing"}`,
			"schema.go": `package fixture
import r "` + role + `"
import s "` + role + `/schema"
func(x *Thing)Schema(resp *r.SchemaResponse){resp.Schema=s.Schema{Attributes:map[string]s.Attribute{"protocol":s.StringAttribute{}}}}`,
		}
		files := map[string]*ast.File{}
		for name, source := range sources {
			file, err := parser.ParseFile(token.NewFileSet(), name, source, 0)
			if err != nil {
				t.Fatal(err)
			}
			files[name] = file
		}
		contracts := PackageReceiverContracts(files)
		registrations := DeclaredRegistrations(files)
		roots := RegisteredSchemaRootsWithContracts(files["schema.go"], registrations, contracts)
		if len(roots) != 1 || !UniqueSchemaOwners(files, contracts)["Thing"] {
			t.Fatalf("split root failed %s", role)
		}
		destination := DeclaredDocumentSources(files, contracts)[SchemaRootReceiver(files["schema.go"], roots[0])]
		if destination != "thing_"+documentSuffix[role]+".go" {
			t.Fatalf("split destination failed %s: %s", role, destination)
		}
	}
}

func TestMetadataParameterIdentityAndDuplicateMethods(t *testing.T) {
	role := "github.com/hashicorp/terraform-plugin-framework/resource"
	for _, methods := range []string{
		`func(x *Thing)Metadata(req d.MetadataRequest,resp *r.MetadataResponse){resp.TypeName=req.ProviderTypeName+"_thing"}`,
		`func(x *Thing)Metadata(req r.MetadataRequest,resp r.MetadataResponse){resp.TypeName=req.ProviderTypeName+"_thing"}`,
		`func(x *Thing)Metadata(req r.MetadataRequest,resp *r.MetadataResponse){req:=custom;resp.TypeName=req.ProviderTypeName+"_thing"}`,
		`func(x *Thing)Metadata(req r.MetadataRequest,resp *r.MetadataResponse){resp.TypeName=req.ProviderTypeName+"_thing"};func(x *Thing)Metadata(req r.MetadataRequest,resp *r.MetadataResponse){resp.TypeName=req.ProviderTypeName+"_thing"}`,
	} {
		source := `package fixture
import r "github.com/hashicorp/terraform-plugin-framework/resource"
import d "github.com/hashicorp/terraform-plugin-framework/datasource"
` + methods
		file, err := parser.ParseFile(token.NewFileSet(), "fixture.go", source, 0)
		if err != nil {
			t.Fatal(err)
		}
		if DeclaredDocumentSources(map[string]*ast.File{"fixture.go": file}, map[string]string{"Thing": role})["Thing"] != "" {
			t.Fatal("unowned metadata identity accepted")
		}
	}
}
func TestPackageAssertionConflictsStayUnresolved(t *testing.T) {
	files := map[string]*ast.File{}
	for name, role := range map[string]string{"first.go": "resource", "second.go": "datasource", "third.go": "resource"} {
		iface := "Resource"
		if role == "datasource" {
			iface = "DataSource"
		}
		source := `package fixture
import r "github.com/hashicorp/terraform-plugin-framework/` + role + `"
var _ r.` + iface + `=&Thing{}`
		file, err := parser.ParseFile(token.NewFileSet(), name, source, 0)
		if err != nil {
			t.Fatal(err)
		}
		files[name] = file
	}
	if PackageReceiverContracts(files)["Thing"] != "" {
		t.Fatal("conflicting package role asserted")
	}
}
