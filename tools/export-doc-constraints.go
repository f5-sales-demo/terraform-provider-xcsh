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
	"strconv"
	"strings"
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
				if !ok || !(strings.HasSuffix(typ.Sel.Name, "Attribute") || strings.HasSuffix(typ.Sel.Name, "Block")) {
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
					}
				}
				fields[strings.Join(child, ".")] = metadata
				walk(value, child)
				return false
			})
		}
		// Schema literals occur in the Schema method; exclude model attr type maps.
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if ok && fn.Name.Name == "Schema" && fn.Body != nil {
				walk(fn.Body, nil)
			}
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
