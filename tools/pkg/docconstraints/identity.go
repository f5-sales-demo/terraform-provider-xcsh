package docconstraints

import (
	"go/ast"
	"go/token"
	"regexp"
	"strconv"
)

var documentSuffix = map[string]string{
	"github.com/hashicorp/terraform-plugin-framework/resource":   "resource",
	"github.com/hashicorp/terraform-plugin-framework/datasource": "data_source",
	"github.com/hashicorp/terraform-plugin-framework/action":     "action",
	"github.com/hashicorp/terraform-plugin-framework/ephemeral":  "ephemeral_resource",
}
var documentName = regexp.MustCompile(`^_[a-z][a-z0-9_]*$`)

// DeclaredDocumentSources derives canonical collection keys from direct typed
// Metadata methods. Physical Go filenames are not document identities.
func DeclaredDocumentSources(files map[string]*ast.File, contracts map[string]string) map[string]string {
	sources := map[string]string{}
	duplicate := map[string]bool{}
	for _, file := range files {
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Name.Name != "Metadata" || fn.Body == nil || fn.Type.Params == nil {
				continue
			}
			receiver := receiverName(fn)
			role := contracts[receiver]
			suffix, owned := documentSuffix[role]
			if !owned {
				continue
			}
			if _, exists := sources[receiver]; exists {
				duplicate[receiver] = true
				sources[receiver] = ""
				continue
			}
			sources[receiver] = ""
			if len(fn.Body.List) != 1 {
				continue
			}
			requests, responses := map[*ast.Object]bool{}, map[*ast.Object]bool{}
			for _, param := range fn.Type.Params.List {
				expression := param.Type
				pointer, indirect := expression.(*ast.StarExpr)
				if indirect {
					expression = pointer.X
				}
				selector, ok := expression.(*ast.SelectorExpr)
				if !ok {
					continue
				}
				qualifier, ok := selector.X.(*ast.Ident)
				if !ok || importPath(file, qualifier) != role {
					continue
				}
				for _, name := range param.Names {
					if selector.Sel.Name == "MetadataRequest" && !indirect {
						requests[name.Obj] = true
					}
					if selector.Sel.Name == "MetadataResponse" && indirect {
						responses[name.Obj] = true
					}
				}
			}
			assignment, ok := fn.Body.List[0].(*ast.AssignStmt)
			if !ok || assignment.Tok != token.ASSIGN || len(assignment.Lhs) != 1 || len(assignment.Rhs) != 1 {
				continue
			}
			lhs, ok := assignment.Lhs[0].(*ast.SelectorExpr)
			if !ok || lhs.Sel.Name != "TypeName" {
				continue
			}
			response, ok := lhs.X.(*ast.Ident)
			if !ok || !responses[response.Obj] {
				continue
			}
			addition, ok := assignment.Rhs[0].(*ast.BinaryExpr)
			if !ok || addition.Op != token.ADD {
				continue
			}
			member, ok := addition.X.(*ast.SelectorExpr)
			if !ok || member.Sel.Name != "ProviderTypeName" {
				continue
			}
			request, ok := member.X.(*ast.Ident)
			if !ok || !requests[request.Obj] {
				continue
			}
			literal, ok := addition.Y.(*ast.BasicLit)
			if !ok || literal.Kind != token.STRING {
				continue
			}
			name, err := strconv.Unquote(literal.Value)
			if err != nil || !documentName.MatchString(name) {
				continue
			}
			if !duplicate[receiver] {
				sources[receiver] = name[1:] + "_" + suffix + ".go"
			}
		}
	}
	destinations := map[string]string{}
	for receiver, source := range sources {
		if source == "" {
			continue
		}
		if previous, exists := destinations[source]; exists {
			sources[previous] = ""
			sources[receiver] = ""
		} else {
			destinations[source] = receiver
		}
	}
	return sources
}

// SchemaRootReceiver returns the enclosing declared method receiver.
func SchemaRootReceiver(file *ast.File, root *ast.CompositeLit) string {
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok && fn.Body != nil && root.Pos() >= fn.Body.Pos() && root.End() <= fn.Body.End() {
			return receiverName(fn)
		}
	}
	return ""
}

// UniqueSchemaOwners rejects synthetic unions of duplicate schema declarations
// and ambiguous provider receivers before collecting any field evidence.
func UniqueSchemaOwners(files map[string]*ast.File, contracts map[string]string) map[string]bool {
	counts := map[string]int{}
	providers := 0
	for _, role := range contracts {
		if role == "github.com/hashicorp/terraform-plugin-framework/provider" {
			providers++
		}
	}
	for _, file := range files {
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if ok && fn.Name.Name == "Schema" && fn.Recv != nil {
				counts[receiverName(fn)]++
			}
		}
	}
	owners := map[string]bool{}
	for receiver, count := range counts {
		owners[receiver] = count == 1 && contracts[receiver] != "" &&
			(contracts[receiver] != "github.com/hashicorp/terraform-plugin-framework/provider" || providers == 1)
	}
	return owners
}
