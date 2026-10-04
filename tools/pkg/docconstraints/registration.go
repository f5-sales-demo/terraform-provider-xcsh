package docconstraints

import (
	"go/ast"
	"go/token"
)

var registrationMethods = map[string]string{
	"Resources":          "github.com/hashicorp/terraform-plugin-framework/resource",
	"DataSources":        "github.com/hashicorp/terraform-plugin-framework/datasource",
	"Actions":            "github.com/hashicorp/terraform-plugin-framework/action",
	"EphemeralResources": "github.com/hashicorp/terraform-plugin-framework/ephemeral",
}

// DeclaredRegistrations binds direct constructor lists to their declared receiver
// contracts. It is static source evidence, not proof of an executed provider.
func DeclaredRegistrations(files map[string]*ast.File) map[string]string {
	type factory struct {
		file string
		fn   *ast.FuncDecl
	}
	factories := map[string]factory{}
	duplicates := map[string]bool{}
	packageName := ""
	for _, file := range files {
		if packageName == "" {
			packageName = file.Name.Name
		}
		if file.Name.Name != packageName {
			return map[string]string{}
		}
	}
	for name, file := range files {
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Recv != nil {
				continue
			}
			if _, exists := factories[fn.Name.Name]; exists {
				duplicates[fn.Name.Name] = true
			}
			factories[fn.Name.Name] = factory{name, fn}
		}
	}
	result := map[string]string{}
	packageContracts := map[string]string{}
	for _, file := range files {
		for receiver, role := range ReceiverContracts(file) {
			previous, exists := packageContracts[receiver]
			if exists && previous != role {
				packageContracts[receiver] = ""
			} else {
				packageContracts[receiver] = role
			}
		}
	}
	conflicts := map[string]bool{}
	for _, file := range files {
		for _, decl := range file.Decls {
			method, ok := decl.(*ast.FuncDecl)
			if !ok || method.Body == nil || packageContracts[receiverName(method)] != "github.com/hashicorp/terraform-plugin-framework/provider" {
				continue
			}
			role, ok := registrationMethods[method.Name.Name]
			if !ok || len(method.Body.List) != 1 {
				continue
			}
			returned, ok := method.Body.List[0].(*ast.ReturnStmt)
			if !ok || len(returned.Results) != 1 {
				continue
			}
			list, ok := returned.Results[0].(*ast.CompositeLit)
			if !ok {
				continue
			}
			array, ok := list.Type.(*ast.ArrayType)
			if !ok || array.Len != nil {
				continue
			}
			function, ok := array.Elt.(*ast.FuncType)
			if !ok || function.Params == nil || len(function.Params.List) != 0 || function.Results == nil || len(function.Results.List) != 1 {
				continue
			}
			if !frameworkInterface(file, function.Results.List[0].Type, role) {
				continue
			}
			for _, element := range list.Elts {
				identifier, ok := element.(*ast.Ident)
				if !ok || duplicates[identifier.Name] {
					continue
				}
				candidate, exists := factories[identifier.Name]
				if identifier.Obj != nil && identifier.Obj.Decl != candidate.fn {
					continue
				}
				if !exists || candidate.fn.Body == nil || len(candidate.fn.Body.List) != 1 || candidate.fn.Type.Params == nil || len(candidate.fn.Type.Params.List) != 0 || candidate.fn.Type.Results == nil || len(candidate.fn.Type.Results.List) != 1 {
					continue
				}
				factoryFile := files[candidate.file]
				if !frameworkInterface(factoryFile, candidate.fn.Type.Results.List[0].Type, role) {
					continue
				}
				returned, ok := candidate.fn.Body.List[0].(*ast.ReturnStmt)
				if !ok || len(returned.Results) != 1 {
					continue
				}
				address, ok := returned.Results[0].(*ast.UnaryExpr)
				if !ok || address.Op != token.AND {
					continue
				}
				literal, ok := address.X.(*ast.CompositeLit)
				if !ok {
					continue
				}
				receiver, ok := literal.Type.(*ast.Ident)
				if !ok || packageContracts[receiver.Name] != role {
					continue
				}
				key := receiver.Name
				previous, exists := result[key]
				if conflicts[key] || (exists && previous != role) {
					conflicts[key] = true
					result[key] = ""
				} else {
					result[key] = role
				}
			}
		}
	}
	return result
}

func frameworkInterface(file *ast.File, expression ast.Expr, role string) bool {
	selector, ok := expression.(*ast.SelectorExpr)
	if !ok || selector.Sel.Name != frameworkInterfaces[role] {
		return false
	}
	qualifier, ok := selector.X.(*ast.Ident)
	return ok && importPath(file, qualifier) == role
}

// RegisteredSchemaRoots retains only directly registered receivers, plus the
// provider's own declared schema. Runtime provider selection remains unverified.
func RegisteredSchemaRoots(file *ast.File, registrations map[string]string) []*ast.CompositeLit {
	roots := SchemaRoots(file)
	result := []*ast.CompositeLit{}
	contracts := ReceiverContracts(file)
	for _, root := range roots {
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Body == nil || root.Pos() < fn.Body.Pos() || root.End() > fn.Body.End() {
				continue
			}
			receiver := receiverName(fn)
			role := contracts[receiver]
			if role == "github.com/hashicorp/terraform-plugin-framework/provider" || (role != "" && registrations[receiver] == role) {
				result = append(result, root)
			}
		}
	}
	return result
}
