// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package codegen (generate.go) provides file generation functions that
// execute Go text/templates against ResourceTemplate data and write the
// formatted output to disk.
package codegen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/schema"
)

// GenerateResourceFile generates the Terraform resource Go file for a single resource.
// outputDir is the directory where the file will be written (e.g. "internal/provider").
func GenerateResourceFile(resource *openapi.ResourceTemplate, outputDir string) error {
	outputPath := filepath.Join(outputDir, resource.Name+"_resource.go")

	// Create template with custom functions
	funcMap := template.FuncMap{
		"renderNestedAttrs":               RenderNestedAttributes,
		"renderNestedBlocks":              RenderNestedBlocks,
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
	formatted, err := imports.Process(outputPath, buf.Bytes(), nil)
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
	formatted, err := imports.Process(outputPath, buf.Bytes(), nil)
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

	formatted, err := imports.Process(outputPath, buf.Bytes(), nil)
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

	formatted, err := imports.Process(outputPath, buf.Bytes(), nil)
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
	rformatted, err := imports.Process(resourcePath, rbuf.Bytes(), nil)
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
	cformatted, err := imports.Process(clientPath, cbuf.Bytes(), nil)
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
	formatted, err := imports.Process(outputPath, buf.Bytes(), nil)
	if err != nil {
		// If formatting fails, write unformatted code with warning
		fmt.Printf("Warning: gofmt failed for %s: %v (writing unformatted)\n", outputPath, err)
		formatted = buf.Bytes()
	}

	return os.WriteFile(outputPath, formatted, 0644)
}
