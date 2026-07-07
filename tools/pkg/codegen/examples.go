// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/naming"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

// Example generation is schema-driven: it emits a minimal, terraform-valid configuration
// derived directly from the resource's TerraformAttribute tree (the same tree that produces
// the provider schema). It emits metadata identity (name, namespace) plus every REQUIRED
// top-level attribute with a schema-valid value, and deliberately omits optional nested blocks
// (which are never framework-Required and would otherwise risk missing-required-inner-argument
// or unsupported-block drift). Because it reads the live schema tree, examples cannot drift out
// of sync with the generated provider.

// exampleIdentityFields are emitted explicitly (name, namespace), so the required-attribute
// loop skips them. Every OTHER required attribute is emitted — including fields that are usually
// optional metadata (e.g. description, disable) but are marked Required for some resources.
// Optional metadata (labels, annotations) is naturally excluded because it is not Required.
var exampleIdentityFields = map[string]bool{"name": true, "namespace": true}

// RenderResourceExampleHCL renders a minimal valid HCL example for a resource.
func RenderResourceExampleHCL(rt *openapi.ResourceTemplate, resourceName, namespaceVal string) string {
	var sb strings.Builder

	human := humanizeResourceName(resourceName)
	sb.WriteString(fmt.Sprintf("# %s Resource Example\n", human))
	if rt.Description != "" {
		sb.WriteString(fmt.Sprintf("# %s\n", firstSentence(rt.Description)))
	}
	sb.WriteString("\n")
	sb.WriteString("terraform {\n  required_version = \">= 1.0\"\n\n  required_providers {\n    xcsh = {\n      source  = \"f5-sales-demo/xcsh\"\n      version = \">= 0.1.0\"\n    }\n  }\n}\n\n")

	sb.WriteString(fmt.Sprintf("# Basic %s configuration\n", human))
	sb.WriteString(fmt.Sprintf("resource \"xcsh_%s\" \"example\" {\n", resourceName))
	sb.WriteString(fmt.Sprintf("  name      = \"example-%s\"\n", strings.ReplaceAll(resourceName, "_", "-")))
	sb.WriteString(fmt.Sprintf("  namespace = \"%s\"\n", namespaceVal))

	// Required top-level, non-block spec attributes with schema-valid values.
	var required []openapi.TerraformAttribute
	for _, attr := range rt.Attributes {
		if attr.IsBlock || !attr.Required || exampleIdentityFields[attr.TfsdkTag] {
			continue
		}
		required = append(required, attr)
	}
	if len(required) > 0 {
		sb.WriteString("\n")
		for _, attr := range required {
			sb.WriteString(fmt.Sprintf("  %s = %s\n", attr.TfsdkTag, exampleValue(attr)))
		}
	}

	sb.WriteString("}\n")
	return sb.String()
}

// RenderDataSourceExampleHCL renders a minimal valid HCL example for a data source lookup.
func RenderDataSourceExampleHCL(resourceName, namespaceVal string) string {
	human := humanizeResourceName(resourceName)
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# %s Data Source Example\n\n", human))
	sb.WriteString("terraform {\n  required_version = \">= 1.0\"\n\n  required_providers {\n    xcsh = {\n      source  = \"f5-sales-demo/xcsh\"\n      version = \">= 0.1.0\"\n    }\n  }\n}\n\n")
	sb.WriteString(fmt.Sprintf("# Look up an existing %s by name\n", human))
	sb.WriteString(fmt.Sprintf("data \"xcsh_%s\" \"example\" {\n", resourceName))
	sb.WriteString(fmt.Sprintf("  name      = \"example-%s\"\n", strings.ReplaceAll(resourceName, "_", "-")))
	sb.WriteString(fmt.Sprintf("  namespace = \"%s\"\n", namespaceVal))
	sb.WriteString("}\n")
	return sb.String()
}

// exampleValue synthesizes a schema-valid HCL value for a required attribute.
func exampleValue(attr openapi.TerraformAttribute) string {
	switch attr.Type {
	case "bool":
		return "true"
	case "int64":
		v := 1
		if attr.Minimum > 0 {
			v = attr.Minimum
		}
		if attr.Maximum > 0 && v > attr.Maximum {
			v = attr.Maximum
		}
		return fmt.Sprintf("%d", v)
	case "map":
		return "{\n    example = \"value\"\n  }"
	case "list":
		return fmt.Sprintf("[%s]", scalarValue(attr.ElementType, attr))
	default: // string
		return scalarValue("string", attr)
	}
}

// scalarValue produces a valid scalar literal for the given element/attribute type.
func scalarValue(typ string, attr openapi.TerraformAttribute) string {
	switch typ {
	case "int64":
		return "1"
	case "bool":
		return "true"
	default: // string
		if len(attr.EnumValues) > 0 {
			return fmt.Sprintf("%q", attr.EnumValues[0])
		}
		return "\"example-value\""
	}
}

// WriteResourceExample writes examples/resources/xcsh_<name>/resource.tf.
func WriteResourceExample(rt *openapi.ResourceTemplate, resourceName, examplesRoot, namespaceVal string) error {
	dir := filepath.Join(examplesRoot, "resources", "xcsh_"+resourceName)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "resource.tf"), []byte(RenderResourceExampleHCL(rt, resourceName, namespaceVal)), 0o644)
}

// WriteDataSourceExample writes examples/data-sources/xcsh_<name>/data-source.tf.
func WriteDataSourceExample(resourceName, examplesRoot, namespaceVal string) error {
	dir := filepath.Join(examplesRoot, "data-sources", "xcsh_"+resourceName)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "data-source.tf"), []byte(RenderDataSourceExampleHCL(resourceName, namespaceVal)), 0o644)
}

func humanizeResourceName(name string) string {
	return naming.ToResourceTypeName(name)
}

func firstSentence(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	if i := strings.Index(s, ". "); i != -1 {
		return s[:i+1]
	}
	return s
}
