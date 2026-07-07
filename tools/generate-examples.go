// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore

// Command generate-examples generates Terraform examples for every resource and data source.
//
// It is spec-free: examples are derived from the committed generated provider schema
// (internal/provider/*_resource.go and *_data_source.go), NOT from the OpenAPI specs. This
// lets it run in CI jobs that do not download specs (e.g. documentation generation), and
// guarantees examples match the provider that is actually shipped.
//
// Rendering is schema-driven via tools/pkg/codegen: a minimal, terraform-valid config with the
// resource identity plus every REQUIRED top-level attribute (enum-aware values, no optional
// blocks). Orphan example directories (no matching provider file) are pruned.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/codegen"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/namespace"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

const providerDir = "internal/provider"

// manuallyMaintained lists resources/data sources whose examples are hand-authored (not
// generated) — they have non-standard schemas or bespoke lookups. See the exceptions in
// scripts/check-no-generated-files.sh. Their example dirs use no xcsh_ prefix, so generating
// xcsh_-prefixed examples for them would be both wrong and duplicative.
var manuallyMaintained = map[string]bool{
	"addon_service":                   true,
	"addon_service_activation_status": true,
}

var (
	attrHeaderRe = regexp.MustCompile(`"([a-z0-9_]+)":\s*schema\.(String|Int64|Bool|Float64|List|Map|Set)Attribute\{`)
	oneOfRe      = regexp.MustCompile(`stringvalidator\.OneOf\(([^)]*)\)`)
	elemTypeRe   = regexp.MustCompile(`ElementType:\s*types\.(String|Int64|Bool)Type`)
	quotedRe     = regexp.MustCompile(`"([^"]*)"`)
	descRe       = regexp.MustCompile(`MarkdownDescription:\s*"((?:[^"\\]|\\.)*)"`)
	requiredRe   = regexp.MustCompile(`Required:\s*true`)
)

func main() {
	resFiles, _ := filepath.Glob(filepath.Join(providerDir, "*_resource.go"))
	dsFiles, _ := filepath.Glob(filepath.Join(providerDir, "*_data_source.go"))

	keep := map[string]bool{}
	var generated, failed int

	for _, f := range resFiles {
		name := strings.TrimSuffix(filepath.Base(f), "_resource.go")
		if manuallyMaintained[name] {
			continue
		}
		keep[name] = true
		rt, err := parseResourceSchema(f, name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ %s: %v\n", name, err)
			failed++
			continue
		}
		ns := exampleNamespace(rt, name)
		if err := codegen.WriteResourceExample(rt, name, "examples", ns); err != nil {
			fmt.Fprintf(os.Stderr, "❌ %s: %v\n", name, err)
			failed++
			continue
		}
		generated++
	}

	for _, f := range dsFiles {
		name := strings.TrimSuffix(filepath.Base(f), "_data_source.go")
		if manuallyMaintained[name] {
			continue
		}
		keep[name] = true
		_, ns := namespace.ForResource(name)
		if err := codegen.WriteDataSourceExample(name, "examples", ns); err != nil {
			fmt.Fprintf(os.Stderr, "❌ %s (data source): %v\n", name, err)
			failed++
		}
	}

	pruneOrphanExampleDirs(keep)
	formatExamples()

	fmt.Printf("\n=== Example generation ===\nGenerated: %d resource examples, %d data-source lookups\n", generated, len(dsFiles))
	if failed > 0 {
		fmt.Fprintf(os.Stderr, "❌ %d example(s) failed\n", failed)
		os.Exit(1)
	}
	fmt.Println("✅ All examples generated (schema-driven, from committed provider)")
}

// exampleNamespace returns the namespace value to use in the generated example. It is
// spec-driven: when the provider schema restricts namespace to a single value (a
// stringvalidator.OneOf captured as EnumValues — e.g. system-only DNS resources), that
// value is used so the example satisfies the constraint. Otherwise it falls back to the
// resource's namespace classification (default "staging").
func exampleNamespace(rt *openapi.ResourceTemplate, name string) string {
	for _, a := range rt.Attributes {
		if a.TfsdkTag == "namespace" && len(a.EnumValues) == 1 {
			return a.EnumValues[0]
		}
	}
	_, ns := namespace.ForResource(name)
	return ns
}

// parseResourceSchema builds a ResourceTemplate carrying the resource's TOP-LEVEL attributes
// (identity + required spec fields) from the generated provider file. Nested block attributes
// live under Blocks{} and are intentionally excluded — minimal valid examples emit no blocks.
func parseResourceSchema(path, name string) (*openapi.ResourceTemplate, error) {
	data, err := os.ReadFile(path) //nolint:gosec // generated file under internal/provider
	if err != nil {
		return nil, err
	}
	content := string(data)

	region, err := topLevelAttributesRegion(content)
	if err != nil {
		return nil, err
	}

	rt := &openapi.ResourceTemplate{Description: firstMarkdownDescription(content)}
	for _, a := range parseAttributes(region) {
		rt.Attributes = append(rt.Attributes, a)
	}
	return rt, nil
}

// topLevelAttributesRegion returns the body of the schema's top-level
// `Attributes: map[string]schema.Attribute{ ... }` using string-aware brace matching.
func topLevelAttributesRegion(content string) (string, error) {
	anchor := strings.Index(content, "resp.Schema = schema.Schema{")
	if anchor == -1 {
		anchor = strings.Index(content, "schema.Schema{")
	}
	if anchor == -1 {
		return "", fmt.Errorf("schema block not found")
	}
	marker := "Attributes: map[string]schema.Attribute{"
	rel := strings.Index(content[anchor:], marker)
	if rel == -1 {
		return "", fmt.Errorf("top-level Attributes map not found")
	}
	open := anchor + rel + len(marker) - 1 // index of the '{'
	body, end := braceBody(content, open)
	if end == -1 {
		return "", fmt.Errorf("unbalanced braces in Attributes map")
	}
	return body, nil
}

// braceBody returns the text between the '{' at openIdx and its matching '}', skipping
// braces inside Go string/rune literals so descriptions containing { } do not confuse matching.
func braceBody(s string, openIdx int) (string, int) {
	depth := 0
	inStr, inRaw, inRune := false, false, false
	for i := openIdx; i < len(s); i++ {
		c := s[i]
		switch {
		case inRaw:
			if c == '`' {
				inRaw = false
			}
			continue
		case inStr:
			if c == '\\' {
				i++
			} else if c == '"' {
				inStr = false
			}
			continue
		case inRune:
			if c == '\\' {
				i++
			} else if c == '\'' {
				inRune = false
			}
			continue
		}
		switch c {
		case '`':
			inRaw = true
		case '"':
			inStr = true
		case '\'':
			inRune = true
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return s[openIdx+1 : i], i
			}
		}
	}
	return "", -1
}

// parseAttributes extracts each top-level attribute (brace-matched body) with the fields the
// example renderer needs: TfsdkTag, Type, Required, EnumValues, ElementType.
func parseAttributes(region string) []openapi.TerraformAttribute {
	var attrs []openapi.TerraformAttribute
	for _, m := range attrHeaderRe.FindAllStringSubmatchIndex(region, -1) {
		tag := region[m[2]:m[3]]
		kind := region[m[4]:m[5]]
		body, end := braceBody(region, m[1]-1) // m[1]-1 is the '{'
		if end == -1 {
			continue
		}
		attr := openapi.TerraformAttribute{
			Name:     tag,
			TfsdkTag: tag,
			Type:     goSchemaKindToType(kind),
			Required: requiredRe.MatchString(body),
		}
		if attr.Type == "list" || attr.Type == "map" {
			attr.ElementType = "string"
			if em := elemTypeRe.FindStringSubmatch(body); em != nil {
				attr.ElementType = strings.ToLower(em[1])
			}
		}
		if attr.Type == "string" {
			if om := oneOfRe.FindStringSubmatch(body); om != nil {
				for _, q := range quotedRe.FindAllStringSubmatch(om[1], -1) {
					attr.EnumValues = append(attr.EnumValues, q[1])
				}
			}
		}
		attrs = append(attrs, attr)
	}
	return attrs
}

func goSchemaKindToType(kind string) string {
	switch kind {
	case "Int64":
		return "int64"
	case "Bool":
		return "bool"
	case "Float64":
		return "int64"
	case "List", "Set":
		return "list"
	case "Map":
		return "map"
	default:
		return "string"
	}
}

func firstMarkdownDescription(content string) string {
	if m := descRe.FindStringSubmatch(content); m != nil {
		return strings.ReplaceAll(m[1], `\"`, `"`)
	}
	return ""
}

// pruneOrphanExampleDirs removes xcsh_-prefixed example dirs with no matching provider file.
func pruneOrphanExampleDirs(keep map[string]bool) {
	for _, sub := range []string{"resources", "data-sources"} {
		matches, _ := filepath.Glob(filepath.Join("examples", sub, "xcsh_*"))
		for _, dir := range matches {
			name := strings.TrimPrefix(filepath.Base(dir), "xcsh_")
			if keep[name] {
				continue
			}
			if err := os.RemoveAll(dir); err == nil {
				fmt.Printf("🧹 Removed orphan example: %s\n", dir)
			}
		}
	}
}

func formatExamples() {
	if _, err := exec.LookPath("terraform"); err != nil {
		return
	}
	_ = exec.Command("terraform", "fmt", "-recursive", "examples").Run() //nolint:gosec // fixed args
}
