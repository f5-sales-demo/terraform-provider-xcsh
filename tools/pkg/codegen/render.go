// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package codegen provides code rendering functions for generating Go source
// code from Terraform attribute definitions. These functions produce Go code
// strings that are embedded into generated resource files via text/template.
package codegen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/naming"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/schema"
)

// RenderRequirementPreflights emits apply-time prerequisite guards for a resource's
// Create/Update. For each declared preflight it nil-checks the triggering block, LISTs
// the requirement's collection in the resource's own namespace, and fails fast with the
// remediation message when the collection is empty — turning an opaque server error
// (e.g. CSD's 500 "Failed to get CSD JS Configuration") into an actionable diagnostic.
// This compiles the dependency declared by x-f5xc-requires into the shipped binary, so
// every remote workstation enforces it identically. recvVar is the resource receiver
// (typically "r"); the emitted code also references the ambient ctx, data, and resp
// that both the Create and Update method bodies provide. Empty input emits nothing, so
// resources without preflights are byte-identical to before.
func RenderRequirementPreflights(preflights []openapi.RequirementPreflight, recvVar string) string {
	if len(preflights) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, p := range preflights {
		listErr := "Failed to verify the " + p.WhenField + " prerequisite (namespace %s): %s"
		sb.WriteString("\n")
		sb.WriteString("\t// Requirement pre-flight (generated; source: x-f5xc-requires):\n")
		sb.WriteString("\t// " + p.Requires + "\n")
		fmt.Fprintf(&sb, "\tif data.%s != nil {\n", p.WhenGoField)
		fmt.Fprintf(&sb, "\t\tpreflightPath := fmt.Sprintf(%s, data.Namespace.ValueString())\n", strconv.Quote(p.ListPath))
		sb.WriteString("\t\tvar preflightResp struct {\n")
		sb.WriteString("\t\t\tItems []map[string]interface{} `json:\"items\"`\n")
		sb.WriteString("\t\t}\n")
		fmt.Fprintf(&sb, "\t\tif err := %s.client.Get(ctx, preflightPath, &preflightResp); err != nil {\n", recvVar)
		fmt.Fprintf(&sb, "\t\t\tresp.Diagnostics.AddError(%s, fmt.Sprintf(%s, data.Namespace.ValueString(), err))\n", strconv.Quote(p.ErrorTitle), strconv.Quote(listErr))
		sb.WriteString("\t\t\treturn\n")
		sb.WriteString("\t\t}\n")
		sb.WriteString("\t\tif len(preflightResp.Items) == 0 {\n")
		fmt.Fprintf(&sb, "\t\t\tresp.Diagnostics.AddError(%s, fmt.Sprintf(%s, data.Namespace.ValueString()))\n", strconv.Quote(p.ErrorTitle), strconv.Quote(p.ErrorDetail))
		sb.WriteString("\t\t\treturn\n")
		sb.WriteString("\t\t}\n")
		sb.WriteString("\t}\n")
	}
	return sb.String()
}

// NestedModelInfo holds information needed to generate a nested model type
type NestedModelInfo struct {
	TypeName    string
	Description string
	Prefix      string // Full prefix path for generating nested type references
	Attributes  []openapi.TerraformAttribute
	IsEmpty     bool
}

// nestedListUsesTypesList reports whether a nested block is a list block, which is always
// modeled as types.List rather than a native Go slice. A native slice cannot represent the
// unknown values a config may carry during planning — whether from a Computed descendant or
// from an element field sourced by an unresolved reference (e.g. the inline API crawler
// domains[].simple_login.password). This matches the top-level list representation
// (RenderBlockFields), and the same predicate drives the model-type decision
// (RenderNestedModelTypes) and the marshal/unmarshal emitters, so they never diverge.
// See #1083.
func nestedListUsesTypesList(attr openapi.TerraformAttribute) bool {
	return attr.NestedBlockType == "list"
}

// EscapeGoString escapes a string for use in a Go string literal
func EscapeGoString(s string) string {
	// Replace backslashes first, then other special characters
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	return s
}

// RegexLiteral returns a Go literal representation of a regex pattern.
// Uses a raw string literal (backticks) unless the pattern contains backticks,
// in which case it falls back to a quoted string with proper escaping.
func RegexLiteral(pattern string) string {
	if !strings.Contains(pattern, "`") {
		return "`" + pattern + "`"
	}
	// Fall back to quoted string - need to escape backslashes and quotes
	escaped := strings.ReplaceAll(pattern, `\`, `\\`)
	escaped = strings.ReplaceAll(escaped, `"`, `\"`)
	return `"` + escaped + `"`
}

// GetGoClientType returns the Go type for use in client structs
func GetGoClientType(attr openapi.TerraformAttribute) string {
	if attr.IsBlock {
		// Nested blocks become pointers to nested structs or slices
		if attr.NestedBlockType == "list" {
			return "[]map[string]interface{}"
		}
		return "map[string]interface{}"
	}

	switch attr.Type {
	case "string":
		return "string"
	case "int64":
		return "int64"
	case "bool":
		return "bool"
	case "list":
		if attr.ElementType == "string" {
			return "[]string"
		} else if attr.ElementType == "int64" {
			return "[]int64"
		}
		return "[]interface{}"
	case "map":
		return "map[string]string"
	default:
		return "interface{}"
	}
}

// RenderSpecStructFields generates Go struct fields for spec attributes
func RenderSpecStructFields(attrs []openapi.TerraformAttribute, indent string) string {
	specFields := schema.FilterSpecFields(attrs)
	if len(specFields) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, attr := range specFields {
		goType := GetGoClientType(attr)
		jsonTag := attr.JsonName
		if jsonTag == "" {
			jsonTag = attr.TfsdkTag
		}
		// For nested blocks, don't use omitempty - the API needs empty objects to be sent
		// For primitive fields, use omitempty to avoid sending zero values
		if attr.IsBlock {
			sb.WriteString(fmt.Sprintf("%s%s %s `json:\"%s\"`\n", indent, attr.GoName, goType, jsonTag))
		} else {
			sb.WriteString(fmt.Sprintf("%s%s %s `json:\"%s,omitempty\"`\n", indent, attr.GoName, goType, jsonTag))
		}
	}
	return sb.String()
}

// RenderSpecMarshalCodeForCreate generates Go code for Create (uses "createReq" variable)
func RenderSpecMarshalCodeForCreate(attrs []openapi.TerraformAttribute, indent string, resourceTitleCase string) string {
	return RenderSpecMarshalCodeWithVar(attrs, indent, "createReq", resourceTitleCase)
}

// RenderSpecMarshalCode generates Go code to marshal spec fields from Terraform state to API struct (uses "apiResource" variable)
func RenderSpecMarshalCode(attrs []openapi.TerraformAttribute, indent string, resourceTitleCase string) string {
	return RenderSpecMarshalCodeWithVar(attrs, indent, "apiResource", resourceTitleCase)
}

// RenderSpecMarshalCodeWithVar generates Go code to marshal spec fields with configurable variable name
func RenderSpecMarshalCodeWithVar(attrs []openapi.TerraformAttribute, indent string, varName string, resourceTitleCase string) string {
	specFields := schema.FilterSpecFields(attrs)
	if len(specFields) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, attr := range specFields {
		if attr.IsBlock {
			// Top-level list blocks are always modeled as types.List (RenderBlockFields);
			// single blocks as pointers.
			renderMarshalBlock(&sb, resourceTitleCase, "", attr, "data."+attr.GoName, varName+".Spec", indent, attr.NestedBlockType == "list")
		} else {
			renderMarshalScalar(&sb, attr, "data."+attr.GoName, varName+".Spec", indent)
		}
	}
	return sb.String()
}

// renderMarshalScalar emits code marshaling a primitive or primitive-list attribute from its
// Terraform value (src) into dstMap[jsonName]. dstMap is a Go expression for a
// map[string]interface{} (e.g. "apiResource.Spec" or a local item map).
func renderMarshalScalar(sb *strings.Builder, attr openapi.TerraformAttribute, src, dstMap, indent string) {
	jsonName := attr.JsonName
	if jsonName == "" {
		jsonName = attr.TfsdkTag
	}
	switch attr.Type {
	case "string":
		sb.WriteString(fmt.Sprintf("%sif !%s.IsNull() && !%s.IsUnknown() {\n", indent, src, src))
		sb.WriteString(fmt.Sprintf("%s\t%s[\"%s\"] = %s.ValueString()\n", indent, dstMap, jsonName, src))
		sb.WriteString(fmt.Sprintf("%s}\n", indent))
	case "int64":
		sb.WriteString(fmt.Sprintf("%sif !%s.IsNull() && !%s.IsUnknown() {\n", indent, src, src))
		sb.WriteString(fmt.Sprintf("%s\t%s[\"%s\"] = %s.ValueInt64()\n", indent, dstMap, jsonName, src))
		sb.WriteString(fmt.Sprintf("%s}\n", indent))
	case "bool":
		sb.WriteString(fmt.Sprintf("%sif !%s.IsNull() && !%s.IsUnknown() {\n", indent, src, src))
		sb.WriteString(fmt.Sprintf("%s\t%s[\"%s\"] = %s.ValueBool()\n", indent, dstMap, jsonName, src))
		sb.WriteString(fmt.Sprintf("%s}\n", indent))
	case "list":
		elemGo := ""
		switch attr.ElementType {
		case "string":
			elemGo = "string"
		case "int64":
			elemGo = "int64"
		}
		if elemGo == "" {
			return
		}
		itemsVar := attr.GoName + "Items"
		sb.WriteString(fmt.Sprintf("%sif !%s.IsNull() && !%s.IsUnknown() {\n", indent, src, src))
		sb.WriteString(fmt.Sprintf("%s\tvar %s []%s\n", indent, itemsVar, elemGo))
		sb.WriteString(fmt.Sprintf("%s\tdiags := %s.ElementsAs(ctx, &%s, false)\n", indent, src, itemsVar))
		sb.WriteString(fmt.Sprintf("%s\tif !diags.HasError() {\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t\t%s[\"%s\"] = %s\n", indent, dstMap, jsonName, itemsVar))
		sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		sb.WriteString(fmt.Sprintf("%s}\n", indent))
	}
}

// renderMarshalBlock emits code marshaling a nested block (attr) from its Terraform model value
// (src) into dstMap[jsonName]. prefixPath is the accumulated nested-model type-name prefix
// (matching CollectNestedModelTypes). fieldIsTypesList reports whether the Go model represents
// this list block as types.List (top-level lists, or nested lists with a Computed descendant)
// rather than a native slice. It recurses to arbitrary depth.
func renderMarshalBlock(sb *strings.Builder, resourceTitleCase, prefixPath string, attr openapi.TerraformAttribute, src, dstMap, indent string, fieldIsTypesList bool) {
	jsonName := attr.JsonName
	if jsonName == "" {
		jsonName = attr.TfsdkTag
	}
	childPath := prefixPath + naming.ToResourceTypeName(attr.TfsdkTag)
	base := attr.GoName

	if attr.NestedBlockType == "list" {
		loopVar := base + "Item"
		mapVar := base + "ItemMap"
		listVar := base + "List"

		emitLoop := func(bodyIndent, rangeExpr string) {
			sb.WriteString(fmt.Sprintf("%svar %s []map[string]interface{}\n", bodyIndent, listVar))
			if len(attr.NestedAttributes) == 0 {
				sb.WriteString(fmt.Sprintf("%sfor range %s {\n", bodyIndent, rangeExpr))
			} else {
				sb.WriteString(fmt.Sprintf("%sfor _, %s := range %s {\n", bodyIndent, loopVar, rangeExpr))
			}
			sb.WriteString(fmt.Sprintf("%s\t%s := make(map[string]interface{})\n", bodyIndent, mapVar))
			for _, child := range attr.NestedAttributes {
				childSrc := loopVar + "." + child.GoName
				if child.IsBlock {
					renderMarshalBlock(sb, resourceTitleCase, childPath, child, childSrc, mapVar, bodyIndent+"\t", nestedListUsesTypesList(child))
				} else {
					renderMarshalScalar(sb, child, childSrc, mapVar, bodyIndent+"\t")
				}
			}
			sb.WriteString(fmt.Sprintf("%s\t%s = append(%s, %s)\n", bodyIndent, listVar, listVar, mapVar))
			sb.WriteString(fmt.Sprintf("%s}\n", bodyIndent))
			sb.WriteString(fmt.Sprintf("%s%s[\"%s\"] = %s\n", bodyIndent, dstMap, jsonName, listVar))
		}

		if fieldIsTypesList {
			elemsVar := base + "Elems"
			elemType := resourceTitleCase + childPath + "Model"
			if len(attr.NestedAttributes) == 0 {
				elemType = resourceTitleCase + "EmptyModel"
			}
			sb.WriteString(fmt.Sprintf("%sif !%s.IsNull() && !%s.IsUnknown() {\n", indent, src, src))
			sb.WriteString(fmt.Sprintf("%s\tvar %s []%s\n", indent, elemsVar, elemType))
			sb.WriteString(fmt.Sprintf("%s\tdiags := %s.ElementsAs(ctx, &%s, false)\n", indent, src, elemsVar))
			sb.WriteString(fmt.Sprintf("%s\tresp.Diagnostics.Append(diags...)\n", indent))
			sb.WriteString(fmt.Sprintf("%s\tif !resp.Diagnostics.HasError() && len(%s) > 0 {\n", indent, elemsVar))
			emitLoop(indent+"\t\t", elemsVar)
			sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
		} else {
			sb.WriteString(fmt.Sprintf("%sif len(%s) > 0 {\n", indent, src))
			emitLoop(indent+"\t", src)
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
		}
		return
	}

	// Single nested block.
	if len(attr.NestedAttributes) == 0 {
		sb.WriteString(fmt.Sprintf("%sif %s != nil {\n", indent, src))
		sb.WriteString(fmt.Sprintf("%s\t%s[\"%s\"] = map[string]interface{}{}\n", indent, dstMap, jsonName))
		sb.WriteString(fmt.Sprintf("%s}\n", indent))
		return
	}
	// Name the map var by the full nested path (childPath), not just the leaf GoName:
	// a single nested block whose child is another single block with the SAME GoName
	// (F5 XC view-ref wrappers, e.g. http_loadbalancer { http_loadbalancer { name } })
	// would otherwise collide the map var, shadowing the outer and emitting a
	// self-referential assignment while the outer map sent to the API stayed empty.
	subVar := childPath + "Map"
	sb.WriteString(fmt.Sprintf("%sif %s != nil {\n", indent, src))
	sb.WriteString(fmt.Sprintf("%s\t%s := make(map[string]interface{})\n", indent, subVar))
	for _, child := range attr.NestedAttributes {
		childSrc := src + "." + child.GoName
		if child.IsBlock {
			renderMarshalBlock(sb, resourceTitleCase, childPath, child, childSrc, subVar, indent+"\t", nestedListUsesTypesList(child))
		} else {
			renderMarshalScalar(sb, child, childSrc, subVar, indent+"\t")
		}
	}
	sb.WriteString(fmt.Sprintf("%s\t%s[\"%s\"] = %s\n", indent, dstMap, jsonName, subVar))
	sb.WriteString(fmt.Sprintf("%s}\n", indent))
}

// RenderComputedFieldsCode generates Go code to set Computed+Optional fields from API response
// This ensures that fields with UseStateForUnknown plan modifier have known values after Create/Update
// The varName parameter specifies the API response variable name (e.g., "created" or "updated")
func RenderComputedFieldsCode(attrs []openapi.TerraformAttribute, indent string, varName string) string {
	specFields := schema.FilterSpecFields(attrs)
	if len(specFields) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(indent + "// Set computed fields from API response\n")

	for _, attr := range specFields {
		// Only generate code for Computed+Optional scalar fields (UseStateForUnknown pattern)
		if !attr.Computed || !attr.Optional || attr.IsBlock {
			continue
		}

		fieldName := attr.GoName
		jsonName := attr.JsonName
		if jsonName == "" {
			jsonName = attr.TfsdkTag
		}

		switch attr.Type {
		case "bool":
			// Set value if API returns it; otherwise handle based on plan value:
			// - If plan was unknown, set to null (resolves unknown state after apply)
			// - If plan had a value, preserve it (user specified this value)
			sb.WriteString(fmt.Sprintf("%sif v, ok := %s.Spec[\"%s\"].(bool); ok {\n", indent, varName, jsonName))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.BoolValue(v)\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s} else if data.%s.IsUnknown() {\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t// API didn't return value and plan was unknown - set to null\n", indent))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.BoolNull()\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
			sb.WriteString(fmt.Sprintf("%s// If plan had a value, preserve it\n", indent))
		case "int64":
			// Set value if API returns it; otherwise handle based on plan value
			sb.WriteString(fmt.Sprintf("%sif v, ok := %s.Spec[\"%s\"].(float64); ok {\n", indent, varName, jsonName))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.Int64Value(int64(v))\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s} else if data.%s.IsUnknown() {\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t// API didn't return value and plan was unknown - set to null\n", indent))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.Int64Null()\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
			sb.WriteString(fmt.Sprintf("%s// If plan had a value, preserve it\n", indent))
		case "string":
			// Set value if API returns it; otherwise handle based on plan value
			sb.WriteString(fmt.Sprintf("%sif v, ok := %s.Spec[\"%s\"].(string); ok && v != \"\" {\n", indent, varName, jsonName))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.StringValue(v)\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s} else if data.%s.IsUnknown() {\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t// API didn't return value and plan was unknown - set to null\n", indent))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.StringNull()\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
			sb.WriteString(fmt.Sprintf("%s// If plan had a value, preserve it\n", indent))
		}
	}

	return sb.String()
}

// RenderCreateComputedFieldsCode generates code for Create method (uses "created" variable)
func RenderCreateComputedFieldsCode(attrs []openapi.TerraformAttribute, indent string) string {
	return RenderComputedFieldsCode(attrs, indent, "created")
}

// RenderUpdateComputedFieldsCode generates code for Update method (uses "updated" variable)
// Deprecated: Use RenderFetchedComputedFieldsCode instead since Update now uses GET after PUT
func RenderUpdateComputedFieldsCode(attrs []openapi.TerraformAttribute, indent string) string {
	return RenderComputedFieldsCode(attrs, indent, "updated")
}

// RenderFetchedComputedFieldsCode generates code for Update method after GET (uses "fetched" variable)
func RenderFetchedComputedFieldsCode(attrs []openapi.TerraformAttribute, indent string) string {
	return RenderComputedFieldsCode(attrs, indent, "fetched")
}

// RenderSpecUnmarshalCode generates Go code to unmarshal spec fields from API response to Terraform state
func RenderSpecUnmarshalCode(attrs []openapi.TerraformAttribute, indent string, resourceTitleCase string) string {
	specFields := schema.FilterSpecFields(attrs)
	if len(specFields) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, attr := range specFields {
		if attr.IsBlock {
			if attr.NestedBlockType == "list" {
				renderUnmarshalTopLevelList(&sb, resourceTitleCase, attr, indent)
			} else {
				renderUnmarshalTopLevelSingle(&sb, resourceTitleCase, attr, indent)
			}
			continue
		}
		renderUnmarshalTopLevelScalar(&sb, attr, indent)
	}
	return sb.String()
}

// nestedElemModel returns the Go model type name for elements of a nested block.
func nestedElemModel(resourceTitleCase, childPath string, attr openapi.TerraformAttribute) string {
	if len(attr.NestedAttributes) == 0 {
		return resourceTitleCase + "EmptyModel"
	}
	return resourceTitleCase + childPath + "Model"
}

// nestedObjectTypeExpr returns the types.ObjectType expression describing a nested block's
// element, referencing the generated <Model>AttrTypes var (or an inline empty map).
func nestedObjectTypeExpr(resourceTitleCase, childPath string, attr openapi.TerraformAttribute) string {
	if len(attr.NestedAttributes) == 0 {
		return "types.ObjectType{AttrTypes: map[string]attr.Type{}}"
	}
	return fmt.Sprintf("types.ObjectType{AttrTypes: %s%sModelAttrTypes}", resourceTitleCase, childPath)
}

// renderUnmarshalTopLevelScalar assigns a primitive/list spec field directly to data.<Field>.
func renderUnmarshalTopLevelScalar(sb *strings.Builder, attr openapi.TerraformAttribute, indent string) {
	fieldName := attr.GoName
	jsonName := attr.JsonName
	if jsonName == "" {
		jsonName = attr.TfsdkTag
	}
	switch attr.Type {
	case "string":
		sb.WriteString(fmt.Sprintf("%sif v, ok := apiResource.Spec[\"%s\"].(string); ok && v != \"\" {\n", indent, jsonName))
		sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.StringValue(v)\n", indent, fieldName))
		sb.WriteString(fmt.Sprintf("%s} else {\n", indent))
		sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.StringNull()\n", indent, fieldName))
		sb.WriteString(fmt.Sprintf("%s}\n", indent))
	case "int64":
		sb.WriteString(fmt.Sprintf("%sif v, ok := apiResource.Spec[\"%s\"].(float64); ok {\n", indent, jsonName))
		sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.Int64Value(int64(v))\n", indent, fieldName))
		sb.WriteString(fmt.Sprintf("%s} else {\n", indent))
		sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.Int64Null()\n", indent, fieldName))
		sb.WriteString(fmt.Sprintf("%s}\n", indent))
	case "bool":
		if attr.Optional {
			sb.WriteString(fmt.Sprintf("%s// Top-level Optional bool: preserve prior state to avoid API default drift\n", indent))
			sb.WriteString(fmt.Sprintf("%sif !isImport && !data.%s.IsNull() && !data.%s.IsUnknown() {\n", indent, fieldName, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t// Normal Read: preserve existing state value (do nothing)\n", indent))
			sb.WriteString(fmt.Sprintf("%s} else {\n", indent))
			sb.WriteString(fmt.Sprintf("%s\tif v, ok := apiResource.Spec[\"%s\"].(bool); ok {\n", indent, jsonName))
			sb.WriteString(fmt.Sprintf("%s\t\tdata.%s = types.BoolValue(v)\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t} else {\n", indent))
			sb.WriteString(fmt.Sprintf("%s\t\tdata.%s = types.BoolNull()\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
		} else {
			sb.WriteString(fmt.Sprintf("%sif v, ok := apiResource.Spec[\"%s\"].(bool); ok {\n", indent, jsonName))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.BoolValue(v)\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s} else {\n", indent))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.BoolNull()\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
		}
	case "list":
		var elemType, goElem, cast, conv string
		switch attr.ElementType {
		case "string":
			elemType, goElem, cast, conv = "types.StringType", "string", "string", "s"
		case "int64":
			elemType, goElem, cast, conv = "types.Int64Type", "int64", "float64", "int64(s)"
		default:
			return
		}
		listVar := attr.TfsdkTag + "List"
		sb.WriteString(fmt.Sprintf("%sif v, ok := apiResource.Spec[\"%s\"].([]interface{}); ok && len(v) > 0 {\n", indent, jsonName))
		sb.WriteString(fmt.Sprintf("%s\tvar %s []%s\n", indent, listVar, goElem))
		sb.WriteString(fmt.Sprintf("%s\tfor _, item := range v {\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t\tif s, ok := item.(%s); ok {\n", indent, cast))
		sb.WriteString(fmt.Sprintf("%s\t\t\t%s = append(%s, %s)\n", indent, listVar, listVar, conv))
		sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		sb.WriteString(fmt.Sprintf("%s\tlistVal, diags := types.ListValueFrom(ctx, %s, %s)\n", indent, elemType, listVar))
		sb.WriteString(fmt.Sprintf("%s\tresp.Diagnostics.Append(diags...)\n", indent))
		sb.WriteString(fmt.Sprintf("%s\tif !resp.Diagnostics.HasError() {\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t\tdata.%s = listVal\n", indent, fieldName))
		sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		sb.WriteString(fmt.Sprintf("%s} else {\n", indent))
		sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.ListNull(%s)\n", indent, fieldName, elemType))
		sb.WriteString(fmt.Sprintf("%s}\n", indent))
	}
}

// renderUnmarshalTopLevelList rebuilds a top-level list block (always types.List) from the API
// response, converting to types.List via ListValueFrom.
func renderUnmarshalTopLevelList(sb *strings.Builder, rc string, attr openapi.TerraformAttribute, indent string) {
	fieldName := attr.GoName
	jsonName := attr.JsonName
	if jsonName == "" {
		jsonName = attr.TfsdkTag
	}
	childPath := naming.ToResourceTypeName(attr.TfsdkTag)
	elemModel := nestedElemModel(rc, childPath, attr)
	objType := nestedObjectTypeExpr(rc, childPath, attr)
	listVar := fieldName + "List"
	existingVar := "existing" + fieldName + "Items"

	// Preserve an unconfigured (null) top-level list on normal Read/Create so a
	// server-managed list the user did not configure does not drift the plan
	// ("Provider produced inconsistent result after apply"). Import still reads the API.
	sb.WriteString(fmt.Sprintf("%sif !isImport && (data.%s.IsNull() || len(data.%s.Elements()) == 0) {\n", indent, fieldName, fieldName))
	sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.ListNull(%s)\n", indent, fieldName, objType))
	sb.WriteString(fmt.Sprintf("%s} else if listData, ok := apiResource.Spec[\"%s\"].([]interface{}); ok && len(listData) > 0 {\n", indent, jsonName))
	sb.WriteString(fmt.Sprintf("%s\tvar %s []%s\n", indent, listVar, elemModel))
	if len(attr.NestedAttributes) == 0 {
		sb.WriteString(fmt.Sprintf("%s\tfor range listData {\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t\t%s = append(%s, %s{})\n", indent, listVar, listVar, elemModel))
		sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
	} else {
		sb.WriteString(fmt.Sprintf("%s\tvar %s []%s\n", indent, existingVar, elemModel))
		sb.WriteString(fmt.Sprintf("%s\tif !data.%s.IsNull() && !data.%s.IsUnknown() {\n", indent, fieldName, fieldName))
		sb.WriteString(fmt.Sprintf("%s\t\tdata.%s.ElementsAs(ctx, &%s, false)\n", indent, fieldName, existingVar))
		sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		sb.WriteString(fmt.Sprintf("%s\tfor listIdx, item := range listData {\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t\t_ = listIdx\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t\tif itemMap, ok := item.(map[string]interface{}); ok {\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t\t\t%s = append(%s, %s{\n", indent, listVar, listVar, elemModel))
		for _, child := range attr.NestedAttributes {
			renderUnmarshalChild(sb, rc, childPath, child, "itemMap", existingVar+"[listIdx]", fmt.Sprintf("len(%s) > listIdx", existingVar), "list", indent+"\t\t\t\t")
		}
		sb.WriteString(fmt.Sprintf("%s\t\t\t})\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
	}
	sb.WriteString(fmt.Sprintf("%s\tlistVal, diags := types.ListValueFrom(ctx, %s, %s)\n", indent, objType, listVar))
	sb.WriteString(fmt.Sprintf("%s\tresp.Diagnostics.Append(diags...)\n", indent))
	sb.WriteString(fmt.Sprintf("%s\tif !resp.Diagnostics.HasError() {\n", indent))
	sb.WriteString(fmt.Sprintf("%s\t\tdata.%s = listVal\n", indent, fieldName))
	sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
	sb.WriteString(fmt.Sprintf("%s} else {\n", indent))
	sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.ListNull(%s)\n", indent, fieldName, objType))
	sb.WriteString(fmt.Sprintf("%s}\n", indent))
}

// renderUnmarshalTopLevelSingle rebuilds a top-level single block (pointer) from the API response.
func renderUnmarshalTopLevelSingle(sb *strings.Builder, rc string, attr openapi.TerraformAttribute, indent string) {
	fieldName := attr.GoName
	jsonName := attr.JsonName
	if jsonName == "" {
		jsonName = attr.TfsdkTag
	}
	childPath := naming.ToResourceTypeName(attr.TfsdkTag)

	if len(attr.NestedAttributes) == 0 {
		// Skip populating server-default oneof empty markers on import: with no
		// prior config to preserve, importing them causes spurious post-import
		// drift. Omitting a default member = the server re-applies the same
		// default, so behavior is unchanged. Non-default/user-intent markers still
		// import normally.
		if !isImportDefaultSuppressed(rc, jsonName) {
			sb.WriteString(fmt.Sprintf("%sif _, ok := apiResource.Spec[\"%s\"].(map[string]interface{}); ok && isImport && data.%s == nil {\n", indent, jsonName, fieldName))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = &%sEmptyModel{}\n", indent, fieldName, rc))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
		}
		return
	}

	model := rc + childPath + "Model"
	// For server-default blocks (suppressed), skip building on import when the API
	// returned an empty object — otherwise import materializes an all-defaults block
	// the user never configured, causing post-import drift. A non-empty response
	// (user configured something) still imports normally.
	buildGuard := fmt.Sprintf("(isImport || data.%s != nil)", fieldName)
	if isImportDefaultSuppressed(rc, jsonName) {
		buildGuard = fmt.Sprintf("((isImport && len(blockData) > 0) || data.%s != nil)", fieldName)
	}
	sb.WriteString(fmt.Sprintf("%sif blockData, ok := apiResource.Spec[\"%s\"].(map[string]interface{}); ok && %s {\n", indent, jsonName, buildGuard))
	sb.WriteString(fmt.Sprintf("%s\tdata.%s = &%s{\n", indent, fieldName, model))
	for _, child := range attr.NestedAttributes {
		renderUnmarshalChild(sb, rc, childPath, child, "blockData", "data."+fieldName, "data."+fieldName+" != nil", "single", indent+"\t\t")
	}
	sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
	sb.WriteString(fmt.Sprintf("%s}\n", indent))
}

// renderUnmarshalChild writes a "Field: <closure>," entry for a child attribute inside a model
// struct literal. srcMap is the Go expr for the map[string]interface{} to read from. container is
// "single" or "list" (the block kind this child lives in). stateBase is the Go expr for the
// existing-state model value at this level (for drift-preserving reads), or "" when unavailable;
// stateGuard is a boolean Go expr true when stateBase is safe to read. It recurses to any depth.
func renderUnmarshalChild(sb *strings.Builder, rc, prefixPath string, child openapi.TerraformAttribute, srcMap, stateBase, stateGuard, container, indent string) {
	if !child.IsBlock {
		renderUnmarshalScalarChild(sb, rc, child, srcMap, stateBase, stateGuard, container, indent)
		return
	}
	childPath := prefixPath + naming.ToResourceTypeName(child.TfsdkTag)
	if child.NestedBlockType == "list" {
		renderUnmarshalListChild(sb, rc, childPath, child, srcMap, stateBase, stateGuard, container, indent)
		return
	}
	renderUnmarshalSingleChild(sb, rc, childPath, child, srcMap, stateBase, stateGuard, container, indent)
}

// renderUnmarshalScalarChild emits the closure entry for a primitive/list child.
// rc is the resource title-case prefix, used for import-default suppression lookups.
func renderUnmarshalScalarChild(sb *strings.Builder, rc string, attr openapi.TerraformAttribute, srcMap, stateBase, stateGuard, container, indent string) {
	fieldName := attr.GoName
	jsonName := attr.JsonName
	if jsonName == "" {
		jsonName = attr.TfsdkTag
	}
	// Single-block Optional int64/bool preserve prior state on normal read to avoid API default drift.
	preserve := container == "single" && attr.Optional && stateBase != "" && (attr.Type == "int64" || attr.Type == "bool")

	switch attr.Type {
	case "string":
		sb.WriteString(fmt.Sprintf("%s%s: func() types.String {\n", indent, fieldName))
		sb.WriteString(fmt.Sprintf("%s\tif v, ok := %s[\"%s\"].(string); ok && v != \"\" {\n", indent, srcMap, jsonName))
		sb.WriteString(fmt.Sprintf("%s\t\treturn types.StringValue(v)\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		sb.WriteString(fmt.Sprintf("%s\treturn types.StringNull()\n", indent))
		sb.WriteString(fmt.Sprintf("%s}(),\n", indent))
	case "int64":
		sb.WriteString(fmt.Sprintf("%s%s: func() types.Int64 {\n", indent, fieldName))
		if preserve {
			// Preserve an explicitly-set prior value to avoid API-default drift, but
			// only when it is KNOWN. If the planned value is unknown (Computed+Optional
			// field left unset by the user), fall through to the API response / null —
			// returning an unknown here trips "invalid result object after apply".
			sb.WriteString(fmt.Sprintf("%s\tif !isImport && %s && !%s.%s.IsUnknown() {\n", indent, stateGuard, stateBase, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t\treturn %s.%s\n", indent, stateBase, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		}
		// Most int64 fields treat a returned 0 as "unset" (v != 0 guard); meaningful-zero
		// leaves (e.g. signature_id, where 0 = "all") must read 0 back faithfully (#1129).
		if isMeaningfulZeroInt64(rc, jsonName) {
			sb.WriteString(fmt.Sprintf("%s\tif v, ok := %s[\"%s\"].(float64); ok {\n", indent, srcMap, jsonName))
		} else {
			sb.WriteString(fmt.Sprintf("%s\tif v, ok := %s[\"%s\"].(float64); ok && v != 0 {\n", indent, srcMap, jsonName))
		}
		sb.WriteString(fmt.Sprintf("%s\t\treturn types.Int64Value(int64(v))\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		sb.WriteString(fmt.Sprintf("%s\treturn types.Int64Null()\n", indent))
		sb.WriteString(fmt.Sprintf("%s}(),\n", indent))
	case "bool":
		sb.WriteString(fmt.Sprintf("%s%s: func() types.Bool {\n", indent, fieldName))
		if preserve {
			// Preserve an explicitly-set prior value to avoid API-default drift, but
			// only when it is KNOWN. If the planned value is unknown (Computed+Optional
			// field left unset by the user), fall through to the API response / null —
			// returning an unknown here trips "invalid result object after apply".
			sb.WriteString(fmt.Sprintf("%s\tif !isImport && %s && !%s.%s.IsUnknown() {\n", indent, stateGuard, stateBase, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t\treturn %s.%s\n", indent, stateBase, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		}
		if isImportDefaultSuppressed(rc, jsonName) {
			// On import a suppressed server-default bool at its false default must not
			// be populated (config omits it) — otherwise the next plan shows a spurious
			// "false -> null". A true value is a real user choice and still imports.
			sb.WriteString(fmt.Sprintf("%s\tif isImport {\n", indent))
			sb.WriteString(fmt.Sprintf("%s\t\tif v, ok := %s[\"%s\"].(bool); ok && !v {\n", indent, srcMap, jsonName))
			sb.WriteString(fmt.Sprintf("%s\t\t\treturn types.BoolNull()\n", indent))
			sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
			sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		}
		sb.WriteString(fmt.Sprintf("%s\tif v, ok := %s[\"%s\"].(bool); ok {\n", indent, srcMap, jsonName))
		sb.WriteString(fmt.Sprintf("%s\t\treturn types.BoolValue(v)\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		sb.WriteString(fmt.Sprintf("%s\treturn types.BoolNull()\n", indent))
		sb.WriteString(fmt.Sprintf("%s}(),\n", indent))
	case "list":
		var elemType, goElem, cast, conv string
		switch attr.ElementType {
		case "string":
			elemType, goElem, cast, conv = "types.StringType", "string", "string", "s"
		case "int64":
			elemType, goElem, cast, conv = "types.Int64Type", "int64", "float64", "int64(s)"
		default:
			return
		}
		sb.WriteString(fmt.Sprintf("%s%s: func() types.List {\n", indent, fieldName))
		sb.WriteString(fmt.Sprintf("%s\tif v, ok := %s[\"%s\"].([]interface{}); ok && len(v) > 0 {\n", indent, srcMap, jsonName))
		sb.WriteString(fmt.Sprintf("%s\t\tvar items []%s\n", indent, goElem))
		sb.WriteString(fmt.Sprintf("%s\t\tfor _, item := range v {\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t\t\tif s, ok := item.(%s); ok {\n", indent, cast))
		sb.WriteString(fmt.Sprintf("%s\t\t\t\titems = append(items, %s)\n", indent, conv))
		sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t\tlistVal, _ := types.ListValueFrom(ctx, %s, items)\n", indent, elemType))
		sb.WriteString(fmt.Sprintf("%s\t\treturn listVal\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		sb.WriteString(fmt.Sprintf("%s\treturn types.ListNull(%s)\n", indent, elemType))
		sb.WriteString(fmt.Sprintf("%s}(),\n", indent))
	}
}

// isObjectReferenceBlock reports whether a nested block is an F5 XC object reference
// (kind/name/namespace/tenant/uid pattern). Such blocks have a server-derived "tenant"
// child. Their server fields (tenant/uid/kind) are Computed-only, so the read-back must
// reconstruct them from the API response rather than preserving the planned value — a
// preserved unknown tenant yields "invalid result object after apply" and a preserved
// null yields "inconsistent result after apply" on newly-added blocks. See #1079.
func isObjectReferenceBlock(attr openapi.TerraformAttribute) bool {
	for _, child := range attr.NestedAttributes {
		tag := child.TfsdkTag
		if tag == "" {
			tag = child.JsonName
		}
		if tag == "tenant" {
			return true
		}
	}
	return false
}

// hasObjectReferenceDescendant reports whether a nested block is, or transitively
// contains at any depth, an F5 XC object reference (a block with a server-derived
// "tenant" child). A block that merely nests a reference deeper down (e.g.
// custom_api_auth_discovery -> api_discovery_ref) must also reconstruct from the API
// response: preserving its planned value carries the nested ref's unknown Computed-only
// tenant and yields "invalid result object after apply". #1080 only handled blocks that
// ARE references; this generalizes it to any nesting depth. See #1091.
func hasObjectReferenceDescendant(attr openapi.TerraformAttribute) bool {
	if isObjectReferenceBlock(attr) {
		return true
	}
	for _, child := range attr.NestedAttributes {
		if hasObjectReferenceDescendant(child) {
			return true
		}
	}
	return false
}

// renderUnmarshalSingleChild emits the closure entry for a single nested block child.
func renderUnmarshalSingleChild(sb *strings.Builder, rc, childPath string, attr openapi.TerraformAttribute, srcMap, stateBase, stateGuard, container, indent string) {
	fieldName := attr.GoName
	jsonName := attr.JsonName
	if jsonName == "" {
		jsonName = attr.TfsdkTag
	}

	if len(attr.NestedAttributes) == 0 {
		// Empty marker block.
		sb.WriteString(fmt.Sprintf("%s%s: func() *%sEmptyModel {\n", indent, fieldName, rc))
		if stateBase != "" {
			// Preserve the PLANNED marker exactly (presence AND absence) on the apply
			// path; import still reads the API below. stateBase is the state accessor
			// at this level — for a single block a pointer (data.Foo.Field), for a list
			// element a positional value (existingItems[idx].Field) — and stateGuard
			// already includes the list len() bound, so returning stateBase.Field is
			// nil-safe. Returning it (rather than the old list-only "&Empty{} when
			// present" heuristic) preserves marker ABSENCE too, so a server-echoed
			// default oneof member the plan omitted does not drift ("was absent, now
			// present"). See #41 (single-container any_client) and #45 (list-element
			// credentials_choice base marker "standard").
			sb.WriteString(fmt.Sprintf("%s\tif !isImport && %s {\n", indent, stateGuard))
			sb.WriteString(fmt.Sprintf("%s\t\treturn %s.%s\n", indent, stateBase, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		}
		// Server-default oneof members must NOT be populated from the API
		// response on import: there is no prior config to preserve, so importing
		// them causes spurious "was absent, now present" drift on the next plan.
		// Guard the response-populate with !isImport for those members; omitting a
		// default member means the server re-applies the same default (safe).
		// Root-only leaves (#1145, e.g. disable_waf) are a server default at the resource
		// root but a DECLARED oneof arm here (nested / inside a list element), so they must
		// read back and round-trip — skip suppression for them at this nested site.
		suppress := isImportDefaultSuppressed(rc, jsonName) && !isSuppressionRootOnly(rc, jsonName)
		popIndent := indent
		if suppress {
			sb.WriteString(fmt.Sprintf("%s\tif !isImport {\n", indent))
			popIndent = indent + "\t"
		}
		sb.WriteString(fmt.Sprintf("%s\tif _, ok := %s[\"%s\"].(map[string]interface{}); ok {\n", popIndent, srcMap, jsonName))
		sb.WriteString(fmt.Sprintf("%s\t\treturn &%sEmptyModel{}\n", popIndent, rc))
		sb.WriteString(fmt.Sprintf("%s\t}\n", popIndent))
		if suppress {
			sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		}
		sb.WriteString(fmt.Sprintf("%s\treturn nil\n", indent))
		sb.WriteString(fmt.Sprintf("%s}(),\n", indent))
		return
	}

	model := rc + childPath + "Model"
	dataVar := fieldName + "Data"
	isRef := isObjectReferenceBlock(attr)
	// Preserve the whole planned block only when it is configured and contains NO object
	// reference anywhere: an object-ref's tenant/uid/kind are Computed-only and unknown in
	// state on create, so any block ON THE PATH to a reference must read those leaves from
	// the API. See #1079 (direct refs) and #1091 (nested refs, e.g. custom_api_auth_discovery).
	preserveWhole := container == "single" && stateBase != "" && !hasObjectReferenceDescendant(attr)

	// When a block cannot be preserved whole because it merely CONTAINS a reference on one
	// arm (a "spine" block), reconstruct from the API but thread the prior-state accessor
	// into its children so off-spine Optional markers/scalars still preserve the planned
	// value (avoiding server-echoed defaults materializing as "was absent/null, now
	// present/false"), while the reference arm reconstructs its Computed tenant from the
	// API. A block that IS itself a reference threads no state — all its leaves are
	// server-derived. See #41 (SP3 API Protection: client_matcher any_client/invert_matcher).
	childStateBase, childStateGuard := "", ""
	if stateBase != "" && !isRef {
		childStateBase = stateBase + "." + fieldName
		childStateGuard = fmt.Sprintf("%s && %s != nil", stateGuard, childStateBase)
	}

	sb.WriteString(fmt.Sprintf("%s%s: func() *%s {\n", indent, fieldName, model))
	if preserveWhole {
		sb.WriteString(fmt.Sprintf("%s\tif !isImport && %s && %s.%s != nil {\n", indent, stateGuard, stateBase, fieldName))
		sb.WriteString(fmt.Sprintf("%s\t\treturn %s.%s\n", indent, stateBase, fieldName))
		sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
	}
	sb.WriteString(fmt.Sprintf("%s\tif %s, ok := %s[\"%s\"].(map[string]interface{}); ok {\n", indent, dataVar, srcMap, jsonName))
	sb.WriteString(fmt.Sprintf("%s\t\treturn &%s{\n", indent, model))
	for _, child := range attr.NestedAttributes {
		renderUnmarshalChild(sb, rc, childPath, child, dataVar, childStateBase, childStateGuard, "single", indent+"\t\t\t")
	}
	sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
	sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
	sb.WriteString(fmt.Sprintf("%s\treturn nil\n", indent))
	sb.WriteString(fmt.Sprintf("%s}(),\n", indent))
}

// renderUnmarshalListChild emits the closure entry for a list nested block child. It always
// returns types.List, matching the Go model (a native slice cannot hold unknown values).
func renderUnmarshalListChild(sb *strings.Builder, rc, childPath string, attr openapi.TerraformAttribute, srcMap, stateBase, stateGuard, container, indent string) {
	fieldName := attr.GoName
	jsonName := attr.JsonName
	if jsonName == "" {
		jsonName = attr.TfsdkTag
	}
	isTypesList := nestedListUsesTypesList(attr)
	elemModel := nestedElemModel(rc, childPath, attr)
	objType := nestedObjectTypeExpr(rc, childPath, attr)
	loopVar := fieldName + "Item"
	mapVar := fieldName + "ItemMap"
	resultVar := fieldName + "Result"

	retType := "[]" + elemModel
	if isTypesList {
		retType = "types.List"
	}

	sb.WriteString(fmt.Sprintf("%s%s: func() %s {\n", indent, fieldName, retType))
	// A suppressed server-computed list (e.g. app_firewall detection_settings.
	// violations_view — the server materializes the full violation catalog whenever
	// detection_settings is configured) must NOT be populated from the API on import,
	// or a config that omits it drifts on round-trip. Matched by leaf name at any
	// depth, mirroring the top-level list suppression. Verified live on the
	// f5-sales-demo WAF exhaustive-coverage matrix.
	if isImportDefaultSuppressed(rc, jsonName) {
		if isTypesList {
			sb.WriteString(fmt.Sprintf("%s\tif isImport {\n%s\t\treturn types.ListNull(%s)\n%s\t}\n", indent, indent, objType, indent))
		} else {
			sb.WriteString(fmt.Sprintf("%s\tif isImport {\n%s\t\treturn nil\n%s\t}\n", indent, indent, indent))
		}
	}
	// Preserve an unconfigured (null/empty) list block on normal Read/Create so a
	// server-managed list the user did not configure does not drift the plan
	// ("Provider produced inconsistent result after apply"). Import still reads the API.
	// Applies whenever prior state is threaded at this level — including a list nested
	// inside a LIST element (container=="list"), e.g. api_testing domains[].credentials[].
	if stateBase != "" {
		if isTypesList {
			sb.WriteString(fmt.Sprintf("%s\tif !isImport && %s && (%s.%s.IsNull() || len(%s.%s.Elements()) == 0) {\n", indent, stateGuard, stateBase, fieldName, stateBase, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t\treturn types.ListNull(%s)\n", indent, objType))
			sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		} else {
			sb.WriteString(fmt.Sprintf("%s\tif !isImport && %s && len(%s.%s) == 0 {\n", indent, stateGuard, stateBase, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t\treturn nil\n", indent))
			sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		}
	}
	// Thread prior-state elements positionally into element children so element Optional
	// markers/scalars preserve the planned value on Read/Create (mirrors the top-level
	// list renderer, render.go renderUnmarshalTopLevelList) instead of materializing
	// server-echoed defaults. Applies whenever prior state is threaded at this level,
	// including a list nested inside a LIST element (container=="list") — e.g. the
	// api_testing domains[].credentials[] where the server echoes the credentials_choice
	// base marker "standard". Nested lists are always types.List, so ElementsAs applies.
	// Import reads the API. See #41 (SP3 client_matcher) and #45 (SP4 credentials).
	threadElem := stateBase != "" && len(attr.NestedAttributes) > 0
	existingVar := fieldName + "Existing"
	idxVar := fieldName + "Idx"
	if threadElem {
		sb.WriteString(fmt.Sprintf("%s\tvar %s []%s\n", indent, existingVar, elemModel))
		sb.WriteString(fmt.Sprintf("%s\tif !isImport && %s && !%s.%s.IsNull() && !%s.%s.IsUnknown() {\n", indent, stateGuard, stateBase, fieldName, stateBase, fieldName))
		sb.WriteString(fmt.Sprintf("%s\t\t%s.%s.ElementsAs(ctx, &%s, false)\n", indent, stateBase, fieldName, existingVar))
		sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
	}
	sb.WriteString(fmt.Sprintf("%s\tif rawList, ok := %s[\"%s\"].([]interface{}); ok && len(rawList) > 0 {\n", indent, srcMap, jsonName))
	sb.WriteString(fmt.Sprintf("%s\t\tvar %s []%s\n", indent, resultVar, elemModel))
	if len(attr.NestedAttributes) == 0 {
		sb.WriteString(fmt.Sprintf("%s\t\tfor range rawList {\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t\t\t%s = append(%s, %s{})\n", indent, resultVar, resultVar, elemModel))
		sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
	} else {
		childStateBase, childStateGuard := "", ""
		if threadElem {
			sb.WriteString(fmt.Sprintf("%s\t\tfor %s, %s := range rawList {\n", indent, idxVar, loopVar))
			sb.WriteString(fmt.Sprintf("%s\t\t\t_ = %s\n", indent, idxVar))
			childStateBase = fmt.Sprintf("%s[%s]", existingVar, idxVar)
			childStateGuard = fmt.Sprintf("len(%s) > %s", existingVar, idxVar)
		} else {
			sb.WriteString(fmt.Sprintf("%s\t\tfor _, %s := range rawList {\n", indent, loopVar))
		}
		sb.WriteString(fmt.Sprintf("%s\t\t\tif %s, ok := %s.(map[string]interface{}); ok {\n", indent, mapVar, loopVar))
		sb.WriteString(fmt.Sprintf("%s\t\t\t\t%s = append(%s, %s{\n", indent, resultVar, resultVar, elemModel))
		for _, child := range attr.NestedAttributes {
			renderUnmarshalChild(sb, rc, childPath, child, mapVar, childStateBase, childStateGuard, "list", indent+"\t\t\t\t\t")
		}
		sb.WriteString(fmt.Sprintf("%s\t\t\t\t})\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
	}
	if isTypesList {
		sb.WriteString(fmt.Sprintf("%s\t\tlistVal, _ := types.ListValueFrom(ctx, %s, %s)\n", indent, objType, resultVar))
		sb.WriteString(fmt.Sprintf("%s\t\treturn listVal\n", indent))
		sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		sb.WriteString(fmt.Sprintf("%s\treturn types.ListNull(%s)\n", indent, objType))
	} else {
		sb.WriteString(fmt.Sprintf("%s\t\treturn %s\n", indent, resultVar))
		sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
		sb.WriteString(fmt.Sprintf("%s\treturn nil\n", indent))
	}
	sb.WriteString(fmt.Sprintf("%s}(),\n", indent))
}

// RenderNestedAttributes generates the Attributes map for nested blocks
func RenderNestedAttributes(attrs []openapi.TerraformAttribute, indent string) string {
	if len(attrs) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(indent + "Attributes: map[string]schema.Attribute{\n")

	for _, attr := range attrs {
		if attr.IsBlock {
			continue // Blocks are handled separately
		}

		attrType := "String"
		switch attr.Type {
		case "int64":
			attrType = "Int64"
		case "bool":
			attrType = "Bool"
		case "map":
			attrType = "Map"
		case "list":
			attrType = "List"
		}

		// Escape backslashes and quotes in descriptions for Go string literals
		desc := EscapeGoString(attr.Description)

		sb.WriteString(fmt.Sprintf("%s\t\"%s\": schema.%sAttribute{\n", indent, attr.TfsdkTag, attrType))
		sb.WriteString(fmt.Sprintf("%s\t\tMarkdownDescription: \"%s\",\n", indent, desc))

		if attr.DeprecationMessage != "" {
			sb.WriteString(fmt.Sprintf("%s\t\tDeprecationMessage: \"%s\",\n", indent, EscapeGoString(attr.DeprecationMessage)))
		}

		if attr.Required {
			sb.WriteString(fmt.Sprintf("%s\t\tRequired: true,\n", indent))
		} else {
			// Handle Optional and Computed flags (can be both true for fields like 'tenant')
			if attr.Optional {
				sb.WriteString(fmt.Sprintf("%s\t\tOptional: true,\n", indent))
			}
			if attr.Computed {
				sb.WriteString(fmt.Sprintf("%s\t\tComputed: true,\n", indent))
			}
		}

		if attr.PlanModifier != "" {
			typeName := "String"
			pkgName := "stringplanmodifier"
			switch attr.Type {
			case "bool":
				typeName = "Bool"
				pkgName = "boolplanmodifier"
			case "int64":
				typeName = "Int64"
				pkgName = "int64planmodifier"
			case "list":
				typeName = "List"
				pkgName = "listplanmodifier"
			case "map":
				typeName = "Map"
				pkgName = "mapplanmodifier"
			}
			sb.WriteString(fmt.Sprintf("%s\t\tPlanModifiers: []planmodifier.%s{\n", indent, typeName))
			switch attr.PlanModifier {
			case "RequiresReplace":
				sb.WriteString(fmt.Sprintf("%s\t\t\t%s.RequiresReplace(),\n", indent, pkgName))
			default:
				sb.WriteString(fmt.Sprintf("%s\t\t\t%s.UseStateForUnknown(),\n", indent, pkgName))
			}
			sb.WriteString(fmt.Sprintf("%s\t\t},\n", indent))
		}

		if attr.Type == "map" || attr.Type == "list" {
			// Map ElementType to Terraform attr type
			elementTfType := "types.StringType"
			switch attr.ElementType {
			case "int64":
				elementTfType = "types.Int64Type"
			case "bool":
				elementTfType = "types.BoolType"
			}
			sb.WriteString(fmt.Sprintf("%s\t\tElementType: %s,\n", indent, elementTfType))
		}

		// Add string validators (LengthBetween/LengthAtMost/LengthAtLeast, RegexMatches, OneOf)
		if attr.Type == "string" {
			var stringValidators []string
			if attr.MinLength > 0 && attr.MaxLength > 0 {
				stringValidators = append(stringValidators, fmt.Sprintf("stringvalidator.LengthBetween(%d, %d)", attr.MinLength, attr.MaxLength))
			} else if attr.MaxLength > 0 {
				stringValidators = append(stringValidators, fmt.Sprintf("stringvalidator.LengthAtMost(%d)", attr.MaxLength))
			} else if attr.MinLength > 0 {
				stringValidators = append(stringValidators, fmt.Sprintf("stringvalidator.LengthAtLeast(%d)", attr.MinLength))
			}
			if attr.Pattern != "" {
				stringValidators = append(stringValidators, fmt.Sprintf("stringvalidator.RegexMatches(regexp.MustCompile(%s), \"\")", RegexLiteral(attr.Pattern)))
			}
			if len(attr.EnumValues) > 0 {
				quoted := make([]string, len(attr.EnumValues))
				for i, v := range attr.EnumValues {
					quoted[i] = fmt.Sprintf("%q", v)
				}
				stringValidators = append(stringValidators, fmt.Sprintf("stringvalidator.OneOf(%s)", strings.Join(quoted, ", ")))
			}
			if attr.ETLDPlusOne {
				stringValidators = append(stringValidators, "validators.ETLDPlusOneValidator()")
			}
			if len(stringValidators) > 0 {
				sb.WriteString(fmt.Sprintf("%s\t\tValidators: []validator.String{\n", indent))
				for _, sv := range stringValidators {
					sb.WriteString(fmt.Sprintf("%s\t\t\t%s,\n", indent, sv))
				}
				sb.WriteString(fmt.Sprintf("%s\t\t},\n", indent))
			}
		}

		// List/set size validators
		if (attr.Type == "list" || attr.Type == "set") && (attr.MinItems > 0 || attr.MaxItems > 0) {
			sb.WriteString(fmt.Sprintf("%s\t\tValidators: []validator.List{\n", indent))
			if attr.MinItems > 0 && attr.MaxItems > 0 {
				sb.WriteString(fmt.Sprintf("%s\t\t\tlistvalidator.SizeBetween(%d, %d),\n", indent, attr.MinItems, attr.MaxItems))
			} else if attr.MaxItems > 0 {
				sb.WriteString(fmt.Sprintf("%s\t\t\tlistvalidator.SizeAtMost(%d),\n", indent, attr.MaxItems))
			} else {
				sb.WriteString(fmt.Sprintf("%s\t\t\tlistvalidator.SizeAtLeast(%d),\n", indent, attr.MinItems))
			}
			sb.WriteString(fmt.Sprintf("%s\t\t},\n", indent))
		}

		sb.WriteString(fmt.Sprintf("%s\t},\n", indent))
	}

	sb.WriteString(indent + "},\n")
	return sb.String()
}

// CollectNestedModelTypes recursively collects all nested model type definitions
func CollectNestedModelTypes(resourceTitleCase string, attrs []openapi.TerraformAttribute, prefix string, collected *[]NestedModelInfo) {
	for _, attr := range attrs {
		if !attr.IsBlock {
			continue
		}

		// Build the type name: ResourceName + Prefix + AttributeName + Model
		currentPrefix := prefix + naming.ToResourceTypeName(attr.TfsdkTag)
		typeName := resourceTitleCase + currentPrefix + "Model"

		// Check if this block has any nested attributes (non-empty)
		hasContent := false
		for _, nested := range attr.NestedAttributes {
			if !nested.IsBlock {
				hasContent = true
				break
			}
		}

		// Also check for nested blocks
		for _, nested := range attr.NestedAttributes {
			if nested.IsBlock {
				hasContent = true
				break
			}
		}

		*collected = append(*collected, NestedModelInfo{
			TypeName:    typeName,
			Description: attr.TfsdkTag,
			Prefix:      currentPrefix, // Store the full prefix for generating nested type references
			Attributes:  attr.NestedAttributes,
			IsEmpty:     !hasContent && len(attr.NestedAttributes) == 0,
		})

		// Recursively collect nested types
		if len(attr.NestedAttributes) > 0 {
			CollectNestedModelTypes(resourceTitleCase, attr.NestedAttributes, currentPrefix, collected)
		}
	}
}

// RenderNestedModelTypes generates all nested model struct definitions for a resource
func RenderNestedModelTypes(resourceTitleCase string, attrs []openapi.TerraformAttribute) string {
	// First check if there are any blocks
	hasBlocks := false
	for _, attr := range attrs {
		if attr.IsBlock {
			hasBlocks = true
			break
		}
	}
	if !hasBlocks {
		return ""
	}

	var sb strings.Builder

	// Add empty model type for blocks with no attributes
	sb.WriteString(fmt.Sprintf("// %sEmptyModel represents empty nested blocks\n", resourceTitleCase))
	sb.WriteString(fmt.Sprintf("type %sEmptyModel struct {\n}\n\n", resourceTitleCase))

	// Collect all nested model types
	var models []NestedModelInfo
	CollectNestedModelTypes(resourceTitleCase, attrs, "", &models)

	// Generate each model type
	for _, model := range models {
		if model.IsEmpty {
			continue // Empty models use the shared EmptyModel
		}

		sb.WriteString(fmt.Sprintf("// %s represents %s block\n", model.TypeName, model.Description))
		sb.WriteString(fmt.Sprintf("type %s struct {\n", model.TypeName))

		// Generate fields for non-block attributes
		for _, attr := range model.Attributes {
			if attr.IsBlock {
				continue
			}
			goType := "String"
			switch attr.Type {
			case "int64":
				goType = "Int64"
			case "bool":
				goType = "Bool"
			case "map":
				goType = "Map"
			case "list":
				goType = "List"
			}
			sb.WriteString(fmt.Sprintf("\t%s types.%s `tfsdk:\"%s\"`\n", attr.GoName, goType, attr.TfsdkTag))
		}

		// Generate pointer fields for nested block attributes
		for _, attr := range model.Attributes {
			if !attr.IsBlock {
				continue
			}
			// Use the model's full prefix to build the nested type name
			nestedTypeName := resourceTitleCase + model.Prefix + naming.ToResourceTypeName(attr.TfsdkTag) + "Model"

			// Check if this nested block is empty
			isNestedEmpty := len(attr.NestedAttributes) == 0
			if !isNestedEmpty {
				hasNonBlockAttrs := false
				for _, nested := range attr.NestedAttributes {
					if !nested.IsBlock {
						hasNonBlockAttrs = true
						break
					}
				}
				isNestedEmpty = !hasNonBlockAttrs && len(attr.NestedAttributes) == 0
			}

			if isNestedEmpty {
				nestedTypeName = resourceTitleCase + "EmptyModel"
			}

			// Nested list blocks are always modeled as types.List (matching top-level lists):
			// a native slice cannot hold the unknown values a config may carry during planning,
			// whether from a Computed descendant or an element sourced by an unresolved
			// reference. Single blocks stay pointers. See #1083.
			if attr.NestedBlockType == "list" {
				sb.WriteString(fmt.Sprintf("\t%s types.List `tfsdk:\"%s\"`\n", attr.GoName, attr.TfsdkTag))
			} else {
				sb.WriteString(fmt.Sprintf("\t%s *%s `tfsdk:\"%s\"`\n", attr.GoName, nestedTypeName, attr.TfsdkTag))
			}
		}

		sb.WriteString("}\n\n")

		// Generate AttrTypes map for this model (needed for types.List conversion)
		sb.WriteString(fmt.Sprintf("// %sAttrTypes defines the attribute types for %s\n", model.TypeName, model.TypeName))
		sb.WriteString(fmt.Sprintf("var %sAttrTypes = map[string]attr.Type{\n", model.TypeName))

		// Generate attr.Type for non-block attributes
		for _, attr := range model.Attributes {
			if attr.IsBlock {
				continue
			}
			var attrType string
			switch attr.Type {
			case "string":
				attrType = "types.StringType"
			case "int64":
				attrType = "types.Int64Type"
			case "bool":
				attrType = "types.BoolType"
			case "map":
				attrType = "types.MapType{ElemType: types.StringType}"
			case "list":
				elemType := "types.StringType"
				switch attr.ElementType {
				case "int64":
					elemType = "types.Int64Type"
				case "bool":
					elemType = "types.BoolType"
				}
				attrType = fmt.Sprintf("types.ListType{ElemType: %s}", elemType)
			default:
				attrType = "types.StringType"
			}
			sb.WriteString(fmt.Sprintf("\t\"%s\": %s,\n", attr.TfsdkTag, attrType))
		}

		// Generate attr.Type for block attributes
		for _, attr := range model.Attributes {
			if !attr.IsBlock {
				continue
			}

			isNestedEmpty := len(attr.NestedAttributes) == 0

			if isNestedEmpty {
				// Empty nested block - use inline empty map
				if attr.NestedBlockType == "list" {
					sb.WriteString(fmt.Sprintf("\t\"%s\": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{}}},\n", attr.TfsdkTag))
				} else {
					sb.WriteString(fmt.Sprintf("\t\"%s\": types.ObjectType{AttrTypes: map[string]attr.Type{}},\n", attr.TfsdkTag))
				}
			} else {
				// Non-empty nested block - reference its AttrTypes variable
				nestedAttrTypesName := resourceTitleCase + model.Prefix + naming.ToResourceTypeName(attr.TfsdkTag) + "ModelAttrTypes"
				if attr.NestedBlockType == "list" {
					sb.WriteString(fmt.Sprintf("\t\"%s\": types.ListType{ElemType: types.ObjectType{AttrTypes: %s}},\n", attr.TfsdkTag, nestedAttrTypesName))
				} else {
					sb.WriteString(fmt.Sprintf("\t\"%s\": types.ObjectType{AttrTypes: %s},\n", attr.TfsdkTag, nestedAttrTypesName))
				}
			}
		}

		sb.WriteString("}\n\n")
	}

	return sb.String()
}

// RenderBlockFields generates the block fields for the main ResourceModel struct
func RenderBlockFields(resourceTitleCase string, attrs []openapi.TerraformAttribute) string {
	var sb strings.Builder

	for _, attr := range attrs {
		if !attr.IsBlock {
			continue
		}

		// Determine the type name
		typeName := resourceTitleCase + naming.ToResourceTypeName(attr.TfsdkTag) + "Model"

		// Check if this block is empty
		isBlockEmpty := len(attr.NestedAttributes) == 0
		if !isBlockEmpty {
			hasNonBlockAttrs := false
			for _, nested := range attr.NestedAttributes {
				if !nested.IsBlock {
					hasNonBlockAttrs = true
					break
				}
			}
			// Also check for any nested blocks
			hasNestedBlocks := false
			for _, nested := range attr.NestedAttributes {
				if nested.IsBlock {
					hasNestedBlocks = true
					break
				}
			}
			isBlockEmpty = !hasNonBlockAttrs && !hasNestedBlocks
		}

		if isBlockEmpty {
			typeName = resourceTitleCase + "EmptyModel"
		}

		// For list nested blocks, use types.List to properly handle unknown values during planning.
		// For single nested blocks, use pointer type.
		if attr.NestedBlockType == "list" {
			sb.WriteString(fmt.Sprintf("\t%s types.List `tfsdk:\"%s\"`\n", attr.GoName, attr.TfsdkTag))
		} else {
			sb.WriteString(fmt.Sprintf("\t%s *%s `tfsdk:\"%s\"`\n", attr.GoName, typeName, attr.TfsdkTag))
		}
	}

	return sb.String()
}

// RenderNestedBlocks generates the Blocks map for nested blocks within a block
func RenderNestedBlocks(attrs []openapi.TerraformAttribute, indent string) string {
	var hasBlocks bool
	for _, attr := range attrs {
		if attr.IsBlock {
			hasBlocks = true
			break
		}
	}

	if !hasBlocks {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(indent + "Blocks: map[string]schema.Block{\n")

	for _, attr := range attrs {
		if !attr.IsBlock {
			continue
		}

		blockType := "SingleNestedBlock"
		if attr.NestedBlockType == "list" {
			blockType = "ListNestedBlock"
		}

		// Escape backslashes and quotes in descriptions
		desc := EscapeGoString(attr.Description)

		sb.WriteString(fmt.Sprintf("%s\t\"%s\": schema.%s{\n", indent, attr.TfsdkTag, blockType))
		sb.WriteString(fmt.Sprintf("%s\t\tMarkdownDescription: \"%s\",\n", indent, desc))

		if attr.NestedBlockType == "list" {
			sb.WriteString(fmt.Sprintf("%s\t\tNestedObject: schema.NestedBlockObject{\n", indent))
			if len(attr.NestedAttributes) > 0 {
				sb.WriteString(RenderNestedAttributes(attr.NestedAttributes, indent+"\t\t\t"))
				sb.WriteString(RenderNestedBlocks(attr.NestedAttributes, indent+"\t\t\t"))
			} else {
				sb.WriteString(fmt.Sprintf("%s\t\t\tAttributes: map[string]schema.Attribute{},\n", indent))
			}
			sb.WriteString(fmt.Sprintf("%s\t\t},\n", indent))
		} else {
			// SingleNestedBlock
			if len(attr.NestedAttributes) > 0 {
				sb.WriteString(RenderNestedAttributes(attr.NestedAttributes, indent+"\t\t"))
				sb.WriteString(RenderNestedBlocks(attr.NestedAttributes, indent+"\t\t"))
			}
		}

		sb.WriteString(fmt.Sprintf("%s\t},\n", indent))
	}

	sb.WriteString(indent + "},\n")
	return sb.String()
}
