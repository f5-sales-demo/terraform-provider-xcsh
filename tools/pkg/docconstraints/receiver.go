package docconstraints

import (
	"go/ast"
	"go/token"
	"path/filepath"
	"strconv"
)

var frameworkInterfaces = map[string]string{
	"github.com/hashicorp/terraform-plugin-framework/resource":   "Resource",
	"github.com/hashicorp/terraform-plugin-framework/datasource": "DataSource",
	"github.com/hashicorp/terraform-plugin-framework/ephemeral":  "EphemeralResource",
	"github.com/hashicorp/terraform-plugin-framework/action":     "Action",
	"github.com/hashicorp/terraform-plugin-framework/provider":   "Provider",
}

func importPath(file *ast.File, qualifier *ast.Ident) string {
	if qualifier.Obj != nil {
		return ""
	}
	for _, spec := range file.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			continue
		}
		alias := filepath.Base(path)
		if spec.Name != nil {
			alias = spec.Name.Name
		}
		if alias == qualifier.Name && alias != "." && alias != "_" {
			return path
		}
	}
	return ""
}

// ReceiverContracts reads explicit compile-time framework interface assertions.
// This binds the receiver's declared role, not runtime constructor registration.
func ReceiverContracts(file *ast.File) map[string]string {
	contracts := map[string]string{}
	for _, decl := range file.Decls {
		declaration, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range declaration.Specs {
			value, ok := spec.(*ast.ValueSpec)
			if !ok || len(value.Names) != 1 || value.Names[0].Name != "_" || len(value.Values) != 1 {
				continue
			}
			selector, ok := value.Type.(*ast.SelectorExpr)
			if !ok {
				continue
			}
			qualifier, ok := selector.X.(*ast.Ident)
			if !ok {
				continue
			}
			path := importPath(file, qualifier)
			iface, owned := frameworkInterfaces[path]
			if !owned || selector.Sel.Name != iface {
				continue
			}
			address, ok := value.Values[0].(*ast.UnaryExpr)
			if !ok || address.Op != token.AND {
				continue
			}
			literal, ok := address.X.(*ast.CompositeLit)
			if !ok {
				continue
			}
			receiver, ok := literal.Type.(*ast.Ident)
			if !ok {
				continue
			}
			if previous, exists := contracts[receiver.Name]; exists && previous != path {
				contracts[receiver.Name] = ""
			} else {
				contracts[receiver.Name] = path
			}
		}
	}
	return contracts
}
func receiverName(fn *ast.FuncDecl) string {
	if fn.Recv == nil || len(fn.Recv.List) != 1 {
		return ""
	}
	expression := fn.Recv.List[0].Type
	if pointer, ok := expression.(*ast.StarExpr); ok {
		expression = pointer.X
	}
	receiver, ok := expression.(*ast.Ident)
	if !ok {
		return ""
	}
	return receiver.Name
}

// frameworkSchemaRole resolves only the schema package belonging to a known role.
func frameworkSchemaRole(file *ast.File, expression ast.Expr) string {
	selector, ok := expression.(*ast.SelectorExpr)
	if !ok || selector.Sel.Name != "Schema" {
		return ""
	}
	qualifier, ok := selector.X.(*ast.Ident)
	if !ok {
		return ""
	}
	path := importPath(file, qualifier)
	for role := range frameworkInterfaces {
		if path == role+"/schema" {
			return role
		}
	}
	return ""
}

// PackageReceiverContracts combines assertions without depending on file placement.
func PackageReceiverContracts(files map[string]*ast.File) map[string]string {
	contracts := map[string]string{}
	for _, file := range files {
		for receiver, role := range ReceiverContracts(file) {
			previous, exists := contracts[receiver]
			if exists && previous != role {
				contracts[receiver] = ""
			} else {
				contracts[receiver] = role
			}
		}
	}
	return contracts
}
