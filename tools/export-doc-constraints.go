//go:build ignore

// Export schema validators and defaults which Terraform's JSON protocol omits.
package main

import (
	"bytes"
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/docconstraints"
)

func main() {
	result := map[string]map[string]map[string]string{}
	files, err := filepath.Glob("internal/provider/*.go")
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range files {
		if strings.HasSuffix(name, "_test.go") {
			continue
		}
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, name, nil, 0)
		if err != nil {
			log.Fatal(err)
		}
		fields := map[string]map[string]string{}
		var walk func(ast.Node, []string)
		walk = func(node ast.Node, path []string) {
			ast.Inspect(node, func(n ast.Node) bool {
				if _, closure := n.(*ast.FuncLit); closure {
					return false
				}
				kv, ok := n.(*ast.KeyValueExpr)
				if !ok {
					return true
				}
				key, ok := kv.Key.(*ast.BasicLit)
				if !ok || key.Kind != token.STRING {
					return true
				}
				value, ok := kv.Value.(*ast.CompositeLit)
				if !ok {
					return true
				}
				typ, ok := value.Type.(*ast.SelectorExpr)
				if !ok || !(docconstraints.FrameworkSelector(file, typ, "Attribute") || docconstraints.FrameworkSelector(file, typ, "Block")) {
					return true
				}
				field, err := strconv.Unquote(key.Value)
				if err != nil {
					log.Fatal(err)
				}
				child := append(append([]string{}, path...), field)
				metadata := map[string]string{}
				for _, elt := range value.Elts {
					item, ok := elt.(*ast.KeyValueExpr)
					if !ok {
						continue
					}
					id, ok := item.Key.(*ast.Ident)
					if !ok {
						continue
					}
					if id.Name == "Validators" || id.Name == "Default" {
						var buf bytes.Buffer
						if err := printer.Fprint(&buf, fset, item.Value); err != nil {
							log.Fatal(err)
						}
						metadata[id.Name] = buf.String()
						if id.Name == "Validators" {
							metadata["EnumExtractionComplete"] = strconv.FormatBool(docconstraints.EnumCoverage(file, item.Value))
							enums := docconstraints.LiteralEnums(file, item.Value)
							if len(enums) > 0 {
								encoded, err := json.Marshal(enums)
								if err != nil {
									log.Fatal(err)
								}
								metadata["EnumValidators"] = string(encoded)
							}
						}
					}
				}
				schemaKey := strings.Join(child, ".")
				if previous, exists := fields[schemaKey]; exists && !reflect.DeepEqual(previous, metadata) {
					log.Fatalf("conflicting schema constraint evidence for %s in %s", schemaKey, name)
				}
				fields[schemaKey] = metadata
				walk(value, child)
				return false
			})
		}
		for _, root := range docconstraints.SchemaRoots(file) {
			walk(root, nil)
		}
		if len(fields) > 0 {
			result[filepath.Base(name)] = fields
		}
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(result); err != nil {
		log.Fatal(err)
	}
}
