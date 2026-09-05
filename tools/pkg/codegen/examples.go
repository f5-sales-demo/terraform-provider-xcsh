// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/namespace"
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

// ExampleNamespace returns the namespace value to use in a generated example.
// A single-value namespace constraint in the committed provider schema is
// authoritative; resources without one retain their namespace-classification
// fallback.
func ExampleNamespace(rt *openapi.ResourceTemplate, resourceName string) string {
	if rt != nil {
		for _, attribute := range rt.Attributes {
			if attribute.TfsdkTag == "namespace" && len(attribute.EnumValues) == 1 {
				return attribute.EnumValues[0]
			}
		}
	}
	_, namespaceValue := namespace.ForResource(resourceName)
	return namespaceValue
}

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
	sb.WriteString(fmt.Sprintf("  %s = %q\n", "namespace", namespaceVal))
	if resourceName == "token" {
		sb.WriteString("  type      = 1\n")
		sb.WriteString("  site_name = \"example-securemesh-site\"\n")
	}

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
	sb.WriteString(fmt.Sprintf("  %s = %q\n", "namespace", namespaceVal))
	sb.WriteString("}\n")
	// The output is not decoration: without something referencing it, the data
	// source is an unused declaration and tflint rejects the example
	// (terraform_unused_declarations). Dropping this block is what left 143
	// committed examples disagreeing with the generator — and the committed ones
	// were the correct side, since they lint clean and the generated ones do not
	// (#1397).
	sb.WriteString(fmt.Sprintf("\noutput \"%s_id\" {\n  value = data.xcsh_%s.example.id\n}\n", resourceName, resourceName))
	return sb.String()
}

// RenderResponseOperationExampleHCL renders a minimal valid example for a
// catalog-owned response operation. Unlike CRUD resources and lookup data
// sources, these surfaces do not necessarily expose name or namespace. Their
// required top-level attributes are therefore the complete source of truth.
func RenderResponseOperationExampleHCL(rt *openapi.ResourceTemplate, resourceName, surface string) string {
	human := humanizeResourceName(resourceName)
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# %s %s Example\n\n", human, humanizeResourceName(surface)))
	sb.WriteString("terraform {\n  required_version = \">= 1.14\"\n\n  required_providers {\n    xcsh = {\n      source  = \"f5-sales-demo/xcsh\"\n      version = \">= 0.1.0\"\n    }\n  }\n}\n\n")

	indent := "  "
	switch surface {
	case "resource":
		sb.WriteString(fmt.Sprintf("resource \"xcsh_%s\" \"example\" {\n", resourceName))
	case "data_source":
		sb.WriteString(fmt.Sprintf("data \"xcsh_%s\" \"example\" {\n", resourceName))
	case "action":
		sb.WriteString(fmt.Sprintf("# The API accepts this upgrade request immediately; convergence is asynchronous.\n"))
		sb.WriteString("# This action does not reconcile a site's pinned software_settings.\n")
		sb.WriteString(fmt.Sprintf("action \"xcsh_%s\" \"example\" {\n  config {\n", resourceName))
		indent = "    "
	default:
		return ""
	}

	for _, attr := range rt.Attributes {
		if attr.IsBlock || !attr.Required {
			continue
		}
		sb.WriteString(fmt.Sprintf("%s%s = %s\n", indent, attr.TfsdkTag, exampleValue(attr)))
	}
	if surface == "action" {
		sb.WriteString("  }\n")
	}
	sb.WriteString("}\n")
	if surface == "data_source" {
		sb.WriteString(fmt.Sprintf("\noutput %q {\n  value = data.xcsh_%s.example\n", resourceName+"_result", resourceName))
		for _, attr := range rt.Attributes {
			if attr.Sensitive {
				sb.WriteString("  sensitive = true\n")
				break
			}
		}
		sb.WriteString("}\n")
	}
	return sb.String()
}

// exampleValue synthesizes a schema-valid HCL value for a required attribute.
func exampleValue(attr openapi.TerraformAttribute) string {
	switch attr.Type {
	case "bool":
		return "true"
	case "int64":
		v := 1
		if attr.HasMinimum {
			v = attr.Minimum
		}
		if attr.HasMaximum && v > attr.Maximum {
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
		if attr.ETLDPlusOne || attr.UseDomainValidator {
			return `"example.com"`
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

// WriteResponseOperationExample writes the canonical example for a response
// operation surface. Terraform-plugin-docs requires action examples at
// examples/actions/xcsh_<name>/action.tf.
func WriteResponseOperationExample(rt *openapi.ResourceTemplate, resourceName, examplesRoot, surface string) error {
	var subdir, filename string
	switch surface {
	case "resource":
		subdir, filename = "resources", "resource.tf"
	case "data_source":
		subdir, filename = "data-sources", "data-source.tf"
	case "action":
		subdir, filename = "actions", "action.tf"
	default:
		return fmt.Errorf("unsupported response-operation example surface %q", surface)
	}
	dir := filepath.Join(examplesRoot, subdir, "xcsh_"+resourceName)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, filename), []byte(RenderResponseOperationExampleHCL(rt, resourceName, surface)), 0o644)
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
