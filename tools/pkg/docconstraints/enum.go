// Package docconstraints extracts typed validator evidence from provider Go syntax.
package docconstraints

import (
	"go/ast"
	"go/token"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const stringValidatorPath = "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

// EnumEvidence describes a validator's literal argument set without inferring dynamic values.
type EnumEvidence struct {
	Version       int      `json:"version"`
	Validator     string   `json:"validator"`
	Values        []string `json:"values"`
	CaseSensitive bool     `json:"case_sensitive"`
	Complete      bool     `json:"complete"`
	Source        string   `json:"source"`
}

// LiteralEnums recognizes only import-qualified framework string enum validators.
// One unresolved argument makes the complete set unknown; partial sets are never asserted.
func LiteralEnums(file *ast.File, expression ast.Expr) []EnumEvidence {
	aliases := map[string]bool{}
	for _, spec := range file.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		if err != nil || path != stringValidatorPath {
			continue
		}
		alias := "stringvalidator"
		if spec.Name != nil {
			alias = spec.Name.Name
		}
		if alias != "." && alias != "_" {
			aliases[alias] = true
		}
	}
	evidence := []EnumEvidence{}
	validators, ok := expression.(*ast.CompositeLit)
	if !ok {
		return evidence
	}
	for _, element := range validators.Elts {
		call, ok := element.(*ast.CallExpr)
		if !ok {
			continue
		}
		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			continue
		}
		qualifier, ok := selector.X.(*ast.Ident)
		if !ok || qualifier.Obj != nil || !aliases[qualifier.Name] {
			continue
		}
		validator := selector.Sel.Name
		if validator != "OneOf" && validator != "OneOfCaseInsensitive" {
			continue
		}
		values := []string{}
		complete := len(call.Args) > 0 && !call.Ellipsis.IsValid()
		for _, arg := range call.Args {
			literal, ok := arg.(*ast.BasicLit)
			if !ok || literal.Kind != token.STRING {
				complete = false
				continue
			}
			value, err := strconv.Unquote(literal.Value)
			if err != nil {
				complete = false
				continue
			}
			values = append(values, value)
		}
		if !complete {
			values = []string{}
		} else {
			sort.Strings(values)
			unique := values[:0]
			for _, value := range values {
				if len(unique) == 0 || unique[len(unique)-1] != value {
					unique = append(unique, value)
				}
			}
			values = unique
		}
		evidence = append(evidence, EnumEvidence{Version: 1, Validator: validator, Values: values, CaseSensitive: validator == "OneOf", Complete: complete, Source: "ast-validator:" + stringValidatorPath + "." + validator})
	}
	sort.SliceStable(evidence, func(i, j int) bool {
		if evidence[i].Validator != evidence[j].Validator {
			return evidence[i].Validator < evidence[j].Validator
		}
		if evidence[i].Complete != evidence[j].Complete {
			return !evidence[i].Complete
		}
		for k := 0; k < len(evidence[i].Values) && k < len(evidence[j].Values); k++ {
			if evidence[i].Values[k] != evidence[j].Values[k] {
				return evidence[i].Values[k] < evidence[j].Values[k]
			}
		}
		return len(evidence[i].Values) < len(evidence[j].Values)
	})
	return evidence
}

// FrameworkSelector verifies import ownership without accepting local receivers.
func FrameworkSelector(file *ast.File, expression ast.Expr, suffix string) bool {
	selector, ok := expression.(*ast.SelectorExpr)
	if !ok || !strings.HasSuffix(selector.Sel.Name, suffix) {
		return false
	}
	qualifier, ok := selector.X.(*ast.Ident)
	if !ok || qualifier.Obj != nil {
		return false
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
		if alias == qualifier.Name && strings.HasPrefix(path, "github.com/hashicorp/terraform-plugin-framework/") && strings.HasSuffix(path, "/schema") {
			return true
		}
	}
	return false
}

// SchemaRoots finds framework Schema literals assigned to a typed SchemaResponse.
func SchemaRoots(file *ast.File) []*ast.CompositeLit {
	roots := []*ast.CompositeLit{}
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Name.Name != "Schema" || fn.Body == nil || fn.Type.Params == nil || fn.Recv == nil {
			continue
		}
		responseNames := map[string]*ast.Object{}
		for _, param := range fn.Type.Params.List {
			pointer, ok := param.Type.(*ast.StarExpr)
			if !ok {
				continue
			}
			selector, ok := pointer.X.(*ast.SelectorExpr)
			if !ok || selector.Sel.Name != "SchemaResponse" {
				continue
			}
			qualifier, ok := selector.X.(*ast.Ident)
			if !ok || qualifier.Obj != nil {
				continue
			}
			owned := false
			for _, spec := range file.Imports {
				path, err := strconv.Unquote(spec.Path.Value)
				if err != nil {
					continue
				}
				alias := filepath.Base(path)
				if spec.Name != nil {
					alias = spec.Name.Name
				}
				if alias == qualifier.Name && strings.HasPrefix(path, "github.com/hashicorp/terraform-plugin-framework/") {
					owned = true
				}
			}
			if owned {
				for _, name := range param.Names {
					responseNames[name.Name] = name.Obj
				}
			}
		}
		methodRoots := []*ast.CompositeLit{}
		assignments := 0
		conditionalAssignment := false
		responseEscapes := false
		// A bounded source export cannot prove arbitrary pointer mutation safe.
		// Permit only the response occurrence used by a direct Schema assignment;
		// all other references (aliases, member mutations, pointers and calls)
		// make this method unresolved.
		permitted := map[*ast.Ident]bool{}
		ast.Inspect(fn.Body, func(node ast.Node) bool {
			if _, closure := node.(*ast.FuncLit); closure {
				return false
			}
			assignment, ok := node.(*ast.AssignStmt)
			if !ok {
				return true
			}
			for _, lhs := range assignment.Lhs {
				member, ok := lhs.(*ast.SelectorExpr)
				if !ok || member.Sel.Name != "Schema" {
					continue
				}
				id, ok := member.X.(*ast.Ident)
				if !ok {
					continue
				}
				if object, exists := responseNames[id.Name]; exists && object == id.Obj {
					permitted[id] = true
				}
			}
			return true
		})
		ast.Inspect(fn.Body, func(node ast.Node) bool {
			id, ok := node.(*ast.Ident)
			if !ok {
				return true
			}
			if object, exists := responseNames[id.Name]; exists && object == id.Obj && !permitted[id] {
				responseEscapes = true
			}
			return true
		})

		ast.Inspect(fn.Body, func(node ast.Node) bool {
			switch node.(type) {
			case *ast.FuncLit:
				return false
			case *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.SwitchStmt, *ast.TypeSwitchStmt, *ast.SelectStmt:
				ast.Inspect(node, func(inner ast.Node) bool {
					if _, closure := inner.(*ast.FuncLit); closure {
						return false
					}
					statement, ok := inner.(*ast.AssignStmt)
					if !ok {
						return true
					}
					for _, lhs := range statement.Lhs {
						field, ok := lhs.(*ast.SelectorExpr)
						if !ok || field.Sel.Name != "Schema" {
							continue
						}
						receiver, ok := field.X.(*ast.Ident)
						if !ok {
							continue
						}
						if object, exists := responseNames[receiver.Name]; exists && object == receiver.Obj {
							conditionalAssignment = true
						}
					}
					return true
				})
				return false
			}
			assignment, ok := node.(*ast.AssignStmt)
			if !ok {
				return true
			}
			for i, lhs := range assignment.Lhs {
				selector, ok := lhs.(*ast.SelectorExpr)
				if !ok || selector.Sel.Name != "Schema" {
					continue
				}
				receiver, ok := selector.X.(*ast.Ident)
				if !ok {
					continue
				}
				object, exists := responseNames[receiver.Name]
				if !exists || object != receiver.Obj {
					continue
				}
				assignments++
				if i >= len(assignment.Rhs) {
					continue
				}
				literal, ok := assignment.Rhs[i].(*ast.CompositeLit)
				if ok && FrameworkSelector(file, literal.Type, "Schema") {
					methodRoots = append(methodRoots, literal)
				}
			}
			return true
		})
		bypassable := false
		if len(methodRoots) == 1 {
			attachment := methodRoots[0].Pos()
			ast.Inspect(fn.Body, func(node ast.Node) bool {
				if _, closure := node.(*ast.FuncLit); closure {
					return false
				}
				if jump, ok := node.(*ast.ReturnStmt); ok && jump.Pos() < attachment {
					bypassable = true
				}
				if jump, ok := node.(*ast.BranchStmt); ok && jump.Pos() < attachment {
					bypassable = true
				}
				return true
			})
		}
		// Superseded or bypassable assignments remain unresolved.
		if !bypassable && !responseEscapes && !conditionalAssignment && assignments == 1 && len(methodRoots) == 1 {
			roots = append(roots, methodRoots[0])
		}
	}
	return roots
}

// EnumCoverage distinguishes a fully inspected direct validator collection from skipped syntax.
func EnumCoverage(file *ast.File, expression ast.Expr) bool {
	literal, ok := expression.(*ast.CompositeLit)
	if !ok {
		return false
	}
	for _, element := range literal.Elts {
		call, ok := element.(*ast.CallExpr)
		if !ok {
			return false
		}
		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return false
		}
		qualifier, ok := selector.X.(*ast.Ident)
		if !ok || qualifier.Obj != nil {
			return false
		}
		official := false
		for _, spec := range file.Imports {
			path, err := strconv.Unquote(spec.Path.Value)
			if err != nil {
				continue
			}
			alias := filepath.Base(path)
			if spec.Name != nil {
				alias = spec.Name.Name
			}
			if alias == qualifier.Name && path == stringValidatorPath {
				official = true
			}
		}
		if !official {
			return false
		}
		for _, arg := range call.Args {
			nested := false
			ast.Inspect(arg, func(node ast.Node) bool {
				if _, ok := node.(*ast.CallExpr); ok {
					nested = true
				}
				return true
			})
			if nested {
				return false
			}
		}
	}
	return true
}
