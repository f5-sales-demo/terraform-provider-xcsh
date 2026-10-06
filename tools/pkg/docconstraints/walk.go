package docconstraints

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
)

// WalkSchemaFields follows only framework schema-bearing members. Helper call
// arguments, defaults, validators, descriptions and callbacks are never schema input.
// Indirect container expressions remain outside this bounded traversal.
func WalkSchemaFields(file *ast.File, root *ast.CompositeLit, visit func([]string, *ast.CompositeLit) error) error {
	var shape func(*ast.CompositeLit, []string) error
	var fields func(ast.Expr, []string) error
	fields = func(expression ast.Expr, path []string) error {
		mapping, ok := expression.(*ast.CompositeLit)
		if !ok {
			return nil
		}
		mappingType, ok := mapping.Type.(*ast.MapType)
		if !ok {
			return nil
		}
		if !(FrameworkSelector(file, mappingType.Value, "Attribute") || FrameworkSelector(file, mappingType.Value, "Block")) {
			return nil
		}
		for _, element := range mapping.Elts {
			pair, ok := element.(*ast.KeyValueExpr)
			if !ok {
				return fmt.Errorf("unsupported schema map element")
			}
			key, ok := pair.Key.(*ast.BasicLit)
			if !ok || key.Kind != token.STRING {
				return fmt.Errorf("nonliteral schema field key")
			}
			name, err := strconv.Unquote(key.Value)
			if err != nil {
				return err
			}
			var expression ast.Expr = pair.Value
			if address, ok := expression.(*ast.UnaryExpr); ok && address.Op == token.AND {
				expression = address.X
			}
			value, ok := expression.(*ast.CompositeLit)
			if !ok {
				continue
			}
			if !(FrameworkSelector(file, value.Type, "Attribute") || FrameworkSelector(file, value.Type, "Block")) {
				continue
			}
			child := append(append([]string{}, path...), name)
			if err := visit(child, value); err != nil {
				return err
			}
			if err := shape(value, child); err != nil {
				return err
			}
		}
		return nil
	}
	shape = func(literal *ast.CompositeLit, path []string) error {
		for _, element := range literal.Elts {
			member, ok := element.(*ast.KeyValueExpr)
			if !ok {
				continue
			}
			key, ok := member.Key.(*ast.Ident)
			if !ok {
				continue
			}
			switch key.Name {
			case "Attributes", "Blocks":
				if err := fields(member.Value, path); err != nil {
					return err
				}
			case "NestedObject", "NestedType":
				var expression ast.Expr = member.Value
				if address, ok := expression.(*ast.UnaryExpr); ok && address.Op == token.AND {
					expression = address.X
				}
				nested, ok := expression.(*ast.CompositeLit)
				if !ok {
					continue
				}
				if !(FrameworkSelector(file, nested.Type, "NestedBlockObject") || FrameworkSelector(file, nested.Type, "NestedAttributeObject")) {
					continue
				}
				if err := shape(nested, path); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return shape(root, nil)
}
