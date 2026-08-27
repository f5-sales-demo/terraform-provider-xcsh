// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore

// Command extract-legacy-smsv2 converts the generated v0.11.49 SDK resource
// schema into a compact, reviewable path manifest. The upstream Go source is
// an input rather than a vendored dependency; the recorded SHA-256 makes any
// source substitution fail closed.
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"sort"
	"strconv"
	"strings"
)

const legacySourceURL = "https://github.com/volterraedge/terraform-provider-volterra/blob/v0.11.49/volterra/resource_auto_volterra_securemesh_site_v2.go"

type legacyManifest struct {
	Version      string        `json:"version"`
	Resource     string        `json:"resource"`
	SourceURL    string        `json:"source_url"`
	SourceSHA256 string        `json:"source_sha256"`
	PathCount    int           `json:"path_count"`
	Paths        []legacyField `json:"paths"`
}

type legacyField struct {
	Path          string      `json:"path"`
	WireKey       string      `json:"wire_key"`
	Type          string      `json:"type"`
	Cardinality   string      `json:"cardinality"`
	Required      bool        `json:"required"`
	Optional      bool        `json:"optional"`
	Computed      bool        `json:"computed"`
	ForceNew      bool        `json:"force_new"`
	Deprecated    bool        `json:"deprecated"`
	Default       interface{} `json:"default,omitempty"`
	ConflictsWith []string    `json:"conflicts_with"`
}

func main() {
	if len(os.Args) != 3 {
		fatalf("usage: go run tools/extract-legacy-smsv2.go UPSTREAM_GO OUTPUT_JSON")
	}
	source, err := os.ReadFile(os.Args[1])
	if err != nil {
		fatalf("read source: %v", err)
	}
	parsed, err := parser.ParseFile(token.NewFileSet(), os.Args[1], source, 0)
	if err != nil {
		fatalf("parse source: %v", err)
	}
	root := findRootSchema(parsed)
	if root == nil {
		fatalf("resourceVolterraSecuremeshSiteV2 root Schema map not found")
	}
	paths, err := parseSchemaMap(root, "")
	if err != nil {
		fatalf("parse legacy schema: %v", err)
	}
	sort.Slice(paths, func(i, j int) bool { return paths[i].Path < paths[j].Path })
	digest := sha256.Sum256(source)
	document := legacyManifest{
		Version:      "0.11.49",
		Resource:     "volterra_securemesh_site_v2",
		SourceURL:    legacySourceURL,
		SourceSHA256: "sha256:" + hex.EncodeToString(digest[:]),
		PathCount:    len(paths),
		Paths:        paths,
	}
	encoded, err := json.MarshalIndent(document, "", "  ")
	if err != nil {
		fatalf("encode manifest: %v", err)
	}
	encoded = append(encoded, '\n')
	if err := os.WriteFile(os.Args[2], encoded, 0o644); err != nil {
		fatalf("write manifest: %v", err)
	}
}

func findRootSchema(file *ast.File) *ast.CompositeLit {
	for _, declaration := range file.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok || function.Name.Name != "resourceVolterraSecuremeshSiteV2" {
			continue
		}
		for _, statement := range function.Body.List {
			ret, ok := statement.(*ast.ReturnStmt)
			if !ok || len(ret.Results) != 1 {
				continue
			}
			resource := dereferenceComposite(ret.Results[0])
			if resource == nil {
				continue
			}
			for _, element := range resource.Elts {
				field, ok := element.(*ast.KeyValueExpr)
				if ok && identifier(field.Key) == "Schema" {
					return dereferenceComposite(field.Value)
				}
			}
		}
	}
	return nil
}

func parseSchemaMap(schemaMap *ast.CompositeLit, prefix string) ([]legacyField, error) {
	fields := make([]legacyField, 0, len(schemaMap.Elts))
	for _, element := range schemaMap.Elts {
		entry, ok := element.(*ast.KeyValueExpr)
		if !ok {
			return nil, fmt.Errorf("unexpected schema map element %T", element)
		}
		name, err := stringLiteral(entry.Key)
		if err != nil {
			return nil, err
		}
		schema := dereferenceComposite(entry.Value)
		if schema == nil {
			return nil, fmt.Errorf("%s: schema value is not a composite literal", name)
		}
		properties := keyedFields(schema)
		typeName := selectorName(properties["Type"])
		maxItems, _ := integerLiteral(properties["MaxItems"])
		cardinality := "single"
		nestedSuffix := ""
		switch typeName {
		case "TypeList", "TypeSet":
			if maxItems == 1 {
				cardinality = "single_block"
			} else {
				cardinality = strings.ToLower(strings.TrimPrefix(typeName, "Type"))
				nestedSuffix = "[]"
			}
		case "TypeMap":
			cardinality = "map"
		}
		path := name
		if prefix != "" {
			path = prefix + "." + name
		}
		field := legacyField{
			Path:          path + nestedSuffix,
			WireKey:       name,
			Type:          terraformType(typeName),
			Cardinality:   cardinality,
			Required:      booleanLiteral(properties["Required"]),
			Optional:      booleanLiteral(properties["Optional"]),
			Computed:      booleanLiteral(properties["Computed"]),
			ForceNew:      booleanLiteral(properties["ForceNew"]),
			Deprecated:    properties["Deprecated"] != nil,
			ConflictsWith: []string{},
		}
		if value, ok := scalarLiteral(properties["Default"]); ok {
			field.Default = value
		}
		fields = append(fields, field)

		elementSchema := dereferenceComposite(properties["Elem"])
		if elementSchema == nil {
			continue
		}
		elementProperties := keyedFields(elementSchema)
		nestedMap := dereferenceComposite(elementProperties["Schema"])
		if nestedMap == nil {
			continue
		}
		nested, err := parseSchemaMap(nestedMap, field.Path)
		if err != nil {
			return nil, err
		}
		fields = append(fields, nested...)
	}
	return fields, nil
}

func keyedFields(literal *ast.CompositeLit) map[string]ast.Expr {
	result := make(map[string]ast.Expr, len(literal.Elts))
	for _, element := range literal.Elts {
		if field, ok := element.(*ast.KeyValueExpr); ok {
			result[identifier(field.Key)] = field.Value
		}
	}
	return result
}

func dereferenceComposite(expression ast.Expr) *ast.CompositeLit {
	if unary, ok := expression.(*ast.UnaryExpr); ok && unary.Op == token.AND {
		expression = unary.X
	}
	literal, _ := expression.(*ast.CompositeLit)
	return literal
}

func identifier(expression ast.Expr) string {
	value, _ := expression.(*ast.Ident)
	if value == nil {
		return ""
	}
	return value.Name
}

func selectorName(expression ast.Expr) string {
	selector, _ := expression.(*ast.SelectorExpr)
	if selector == nil {
		return ""
	}
	return selector.Sel.Name
}

func stringLiteral(expression ast.Expr) (string, error) {
	literal, ok := expression.(*ast.BasicLit)
	if !ok || literal.Kind != token.STRING {
		return "", fmt.Errorf("expected string literal, got %T", expression)
	}
	return strconv.Unquote(literal.Value)
}

func integerLiteral(expression ast.Expr) (int, bool) {
	literal, ok := expression.(*ast.BasicLit)
	if !ok || literal.Kind != token.INT {
		return 0, false
	}
	value, err := strconv.Atoi(literal.Value)
	return value, err == nil
}

func booleanLiteral(expression ast.Expr) bool {
	value, _ := expression.(*ast.Ident)
	return value != nil && value.Name == "true"
}

func scalarLiteral(expression ast.Expr) (interface{}, bool) {
	if expression == nil {
		return nil, false
	}
	if value, ok := expression.(*ast.Ident); ok {
		switch value.Name {
		case "true":
			return true, true
		case "false":
			return false, true
		}
	}
	literal, ok := expression.(*ast.BasicLit)
	if !ok {
		return nil, false
	}
	switch literal.Kind {
	case token.STRING:
		value, err := strconv.Unquote(literal.Value)
		return value, err == nil
	case token.INT:
		value, err := strconv.Atoi(literal.Value)
		return value, err == nil
	}
	return nil, false
}

func terraformType(name string) string {
	switch name {
	case "TypeBool":
		return "boolean"
	case "TypeFloat":
		return "number"
	case "TypeInt":
		return "integer"
	case "TypeList":
		return "list"
	case "TypeMap":
		return "map"
	case "TypeSet":
		return "set"
	case "TypeString":
		return "string"
	default:
		return "unknown"
	}
}

func fatalf(format string, arguments ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", arguments...)
	os.Exit(1)
}
