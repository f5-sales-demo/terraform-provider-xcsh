//go:build ignore

// Export schema validators and defaults which Terraform's JSON protocol omits.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	parsed := map[string]*ast.File{}
	positions := map[string]*token.FileSet{}
	for _, name := range files {
		if strings.HasSuffix(name, "_test.go") {
			continue
		}
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, name, nil, 0)
		if err != nil {
			log.Fatal(err)
		}
		parsed[name] = file
		positions[name] = fset
	}
	registrations := docconstraints.DeclaredRegistrations(parsed)
	for _, name := range files {
		if strings.HasSuffix(name, "_test.go") {
			continue
		}
		fset := positions[name]
		file := parsed[name]
		fields := map[string]map[string]string{}

		for _, root := range docconstraints.RegisteredSchemaRoots(file, registrations) {
			err := docconstraints.WalkSchemaFields(file, root, func(child []string, value *ast.CompositeLit) error {
				metadata := map[string]string{}
				for _, element := range value.Elts {
					item, ok := element.(*ast.KeyValueExpr)
					if !ok {
						continue
					}
					key, ok := item.Key.(*ast.Ident)
					if !ok {
						continue
					}
					if key.Name != "Validators" && key.Name != "Default" {
						continue
					}
					var buffer bytes.Buffer
					if err := printer.Fprint(&buffer, fset, item.Value); err != nil {
						return err
					}
					metadata[key.Name] = buffer.String()
					if key.Name == "Validators" {
						metadata["EnumExtractionComplete"] = strconv.FormatBool(docconstraints.EnumCoverage(file, item.Value))
						enums := docconstraints.LiteralEnums(file, item.Value)
						if len(enums) > 0 {
							encoded, err := json.Marshal(enums)
							if err != nil {
								return err
							}
							metadata["EnumValidators"] = string(encoded)
						}
					}
				}
				schemaKey := strings.Join(child, ".")
				if previous, exists := fields[schemaKey]; exists && !reflect.DeepEqual(previous, metadata) {
					return fmt.Errorf("conflicting schema constraint evidence for %s in %s", schemaKey, name)
				}
				fields[schemaKey] = metadata
				return nil
			})
			if err != nil {
				log.Fatal(err)
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
