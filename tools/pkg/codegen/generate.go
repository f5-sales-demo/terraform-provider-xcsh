// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package codegen (generate.go) provides file generation functions that
// execute Go text/templates against ResourceTemplate data and write the
// formatted output to disk.
package codegen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"golang.org/x/tools/go/ast/astutil"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/schema"
)

type generatedImport struct {
	path  string
	alias string
}

// generatedSelectorImports contains packages that render helpers may reference
// conditionally. Imports are derived mechanically from selector expressions in
// the rendered source, rather than from package discovery in the local module
// cache. This keeps generation identical on a cold immutable runner and a warm
// hosted runner.
var generatedSelectorImports = map[string]generatedImport{
	"attr":              {path: "github.com/hashicorp/terraform-plugin-framework/attr"},
	"boolplanmodifier":  {path: "github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"},
	"int64planmodifier": {path: "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"},
	"listplanmodifier":  {path: "github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"},
	"mapplanmodifier":   {path: "github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"},
	"stringdefault":     {path: "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"},
	"int64validator":    {path: "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"},
	"listvalidator":     {path: "github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"},
	"stringvalidator":   {path: "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"},
	"regexp":            {path: "regexp"},
	"errors":            {path: "errors"},
	"xcsherrors":        {path: "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors", alias: "xcsherrors"},
}

func formatGeneratedSource(filename string, source []byte) ([]byte, error) {
	fileset := token.NewFileSet()
	file, err := parser.ParseFile(fileset, filename, source, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	imported := make(map[string]struct{}, len(file.Imports))
	for _, spec := range file.Imports {
		name := ""
		if spec.Name != nil {
			name = spec.Name.Name
		} else if path, unquoteErr := strconv.Unquote(spec.Path.Value); unquoteErr == nil {
			name = filepath.Base(path)
		}
		if name != "" {
			imported[name] = struct{}{}
		}
	}

	used := make(map[string]struct{})
	ast.Inspect(file, func(node ast.Node) bool {
		selector, ok := node.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		identifier, ok := selector.X.(*ast.Ident)
		if ok {
			used[identifier.Name] = struct{}{}
		}
		return true
	})

	for name, dependency := range generatedSelectorImports {
		if _, referenced := used[name]; !referenced {
			continue
		}
		if _, present := imported[name]; present {
			continue
		}
		if dependency.alias == "" {
			astutil.AddImport(fileset, file, dependency.path)
		} else {
			astutil.AddNamedImport(fileset, file, dependency.alias, dependency.path)
		}
	}

	var formatted bytes.Buffer
	if err := format.Node(&formatted, fileset, file); err != nil {
		return nil, err
	}
	return formatted.Bytes(), nil
}

// GenerateResourceFile generates the Terraform resource Go file for a single resource.
// outputDir is the directory where the file will be written (e.g. "internal/provider").
func GenerateResourceFile(resource *openapi.ResourceTemplate, outputDir string) error {
	outputPath := filepath.Join(outputDir, resource.Name+"_resource.go")
	// Derive every conditional plan-modifier flag at the renderer boundary. This
	// gives direct callers the same final-IR guarantee as generator orchestration
	// and removes both stale-positive and stale-negative template state.
	schema.RefreshResourcePlanModifierUsage(resource)

	// Create template with custom functions
	funcMap := template.FuncMap{
		"renderNestedAttrs":               RenderNestedAttributes,
		"renderNestedBlocks":              RenderNestedBlocks,
		"renderConditionalRequired":       RenderConditionalRequiredValidators,
		"renderNestedModelTypes":          RenderNestedModelTypes,
		"renderBlockFields":               RenderBlockFields,
		"renderSpecStructFields":          RenderSpecStructFields,
		"renderSpecMarshalCode":           RenderSpecMarshalCode,
		"renderSpecMarshalCodeForCreate":  RenderSpecMarshalCodeForCreate,
		"renderSpecUnmarshalCode":         RenderSpecUnmarshalCode,
		"renderPreflights":                RenderRequirementPreflights,
		"add":                             func(a, b int) int { return a + b },
		"renderCreateComputedFieldsCode":  RenderCreateComputedFieldsCode,
		"renderUpdateComputedFieldsCode":  RenderUpdateComputedFieldsCode,
		"renderFetchedComputedFieldsCode": RenderFetchedComputedFieldsCode,
		"filterSpecFields":                schema.FilterSpecFields,
		"enumValuesLiteral": func(values []string) string {
			quoted := make([]string, len(values))
			for i, v := range values {
				quoted[i] = fmt.Sprintf("%q", v)
			}
			return strings.Join(quoted, ", ")
		},
		"regexLiteral": RegexLiteral,
	}

	tmpl, err := template.New("resource").Funcs(funcMap).Parse(ResourceTemplate)
	if err != nil {
		return fmt.Errorf("template parse error: %w", err)
	}

	// Execute template to buffer first
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, resource); err != nil {
		return fmt.Errorf("template execute error: %w", err)
	}

	// Format the generated code with gofmt
	formatted, err := formatGeneratedSource(outputPath, buf.Bytes())
	if err != nil {
		// If formatting fails, write unformatted code with warning
		fmt.Printf("Warning: gofmt failed for %s: %v (writing unformatted)\n", outputPath, err)
		formatted = buf.Bytes()
	}

	return os.WriteFile(outputPath, formatted, 0644)
}

// GenerateClientTypes generates the client type Go file for a single resource.
// clientDir is the directory where the file will be written (e.g. "internal/client").
func GenerateClientTypes(resource *openapi.ResourceTemplate, clientDir string) error {
	outputPath := filepath.Join(clientDir, resource.Name+"_types.go")

	// Create template with custom functions for spec field generation
	funcMap := template.FuncMap{
		"renderSpecStructFields": func(attrs []openapi.TerraformAttribute) string {
			return RenderSpecStructFields(attrs, "\t")
		},
	}

	tmpl, err := template.New("client").Funcs(funcMap).Parse(ClientTemplate)
	if err != nil {
		return fmt.Errorf("template parse error: %w", err)
	}

	// Execute template to buffer first
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, resource); err != nil {
		return fmt.Errorf("template execute error: %w", err)
	}

	// Format the generated code with gofmt
	formatted, err := formatGeneratedSource(outputPath, buf.Bytes())
	if err != nil {
		// If formatting fails, write unformatted code with warning
		fmt.Printf("Warning: gofmt failed for %s: %v (writing unformatted)\n", outputPath, err)
		formatted = buf.Bytes()
	}

	return os.WriteFile(outputPath, formatted, 0644)
}

// GenerateReadOnlyDataSource generates a data-source-only file for a read-only resource.
func GenerateReadOnlyDataSource(resource *openapi.ResourceTemplate, outputDir string) error {
	outputPath := filepath.Join(outputDir, resource.Name+"_data_source.go")

	tmpl, err := template.New("readonly_ds").Parse(ReadOnlyDataSourceTemplate)
	if err != nil {
		return fmt.Errorf("template parse error: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, resource); err != nil {
		return fmt.Errorf("template execute error: %w", err)
	}

	formatted, err := formatGeneratedSource(outputPath, buf.Bytes())
	if err != nil {
		fmt.Printf("Warning: gofmt failed for %s: %v (writing unformatted)\n", outputPath, err)
		formatted = buf.Bytes()
	}

	return os.WriteFile(outputPath, formatted, 0644)
}

// GenerateReadOnlyClientTypes generates a Get-only client type file for a read-only resource.
func GenerateReadOnlyClientTypes(resource *openapi.ResourceTemplate, clientDir string) error {
	outputPath := filepath.Join(clientDir, resource.Name+"_types.go")

	tmpl, err := template.New("readonly_client").Parse(ReadOnlyClientTemplate)
	if err != nil {
		return fmt.Errorf("template parse error: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, resource); err != nil {
		return fmt.Errorf("template execute error: %w", err)
	}

	formatted, err := formatGeneratedSource(outputPath, buf.Bytes())
	if err != nil {
		fmt.Printf("Warning: gofmt failed for %s: %v (writing unformatted)\n", outputPath, err)
		formatted = buf.Bytes()
	}

	return os.WriteFile(outputPath, formatted, 0644)
}

// GenerateActionResource generates the resource file and the client request-body
// types file for an action-style resource (x-f5xc-action). It emits no data
// source and no CRUD client: Create marshals the request struct via the generic
// Post, and Read uses GetLenient into a generic map.
func GenerateActionResource(resource *openapi.ResourceTemplate, outputDir, clientDir string) error {
	// Resource file.
	resourcePath := filepath.Join(outputDir, resource.Name+"_resource.go")
	rtmpl, err := template.New("action_resource").Parse(ActionResourceTemplate)
	if err != nil {
		return fmt.Errorf("action resource template parse error: %w", err)
	}
	var rbuf bytes.Buffer
	if err := rtmpl.Execute(&rbuf, resource); err != nil {
		return fmt.Errorf("action resource template execute error: %w", err)
	}
	rformatted, err := formatGeneratedSource(resourcePath, rbuf.Bytes())
	if err != nil {
		fmt.Printf("Warning: gofmt failed for %s: %v (writing unformatted)\n", resourcePath, err)
		rformatted = rbuf.Bytes()
	}
	if err := os.WriteFile(resourcePath, rformatted, 0644); err != nil {
		return err
	}

	// Client request-body types file.
	clientPath := filepath.Join(clientDir, resource.Name+"_types.go")
	ctmpl, err := template.New("action_client").Parse(ActionClientTemplate)
	if err != nil {
		return fmt.Errorf("action client template parse error: %w", err)
	}
	var cbuf bytes.Buffer
	if err := ctmpl.Execute(&cbuf, resource); err != nil {
		return fmt.Errorf("action client template execute error: %w", err)
	}
	cformatted, err := formatGeneratedSource(clientPath, cbuf.Bytes())
	if err != nil {
		fmt.Printf("Warning: gofmt failed for %s: %v (writing unformatted)\n", clientPath, err)
		cformatted = cbuf.Bytes()
	}
	return os.WriteFile(clientPath, cformatted, 0644)
}

// GenerateDataSource generates the Terraform data source Go file for a single resource.
// outputDir is the directory where the file will be written (e.g. "internal/provider").
func GenerateDataSource(resource *openapi.ResourceTemplate, outputDir string) error {
	outputPath := filepath.Join(outputDir, resource.Name+"_data_source.go")

	tmpl, err := template.New("datasource").Parse(DataSourceTemplate)
	if err != nil {
		return fmt.Errorf("template parse error: %w", err)
	}

	// Execute template to buffer first
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, resource); err != nil {
		return fmt.Errorf("template execute error: %w", err)
	}

	// Format the generated code with gofmt
	formatted, err := formatGeneratedSource(outputPath, buf.Bytes())
	if err != nil {
		// If formatting fails, write unformatted code with warning
		fmt.Printf("Warning: gofmt failed for %s: %v (writing unformatted)\n", outputPath, err)
		formatted = buf.Bytes()
	}

	return os.WriteFile(outputPath, formatted, 0644)
}
