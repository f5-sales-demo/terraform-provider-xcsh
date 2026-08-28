// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/naming"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

// GenerateResponseOperation writes the provider implementation for one
// catalog-owned non-CRUD operation.
func GenerateResponseOperation(operation *openapi.ResponseOperationTemplate, providerDir string) error {
	if operation == nil {
		return fmt.Errorf("response operation template is required")
	}
	var suffix, source string
	switch operation.Role {
	case "query", "collection":
		suffix, source = "_data_source.go", responseDataSourceTemplate
	case "issuance":
		suffix, source = "_resource.go", responseIssuanceTemplate
	case "action":
		suffix, source = "_action.go", responseActionTemplate
	default:
		return fmt.Errorf("unsupported response-operation role %q", operation.Role)
	}
	path := filepath.Join(providerDir, operation.Name+suffix)
	if err := ensureResponseOperationTarget(path); err != nil {
		return err
	}
	responseUnmarshal, err := renderResponseOperationUnmarshal(operation)
	if err != nil {
		return err
	}
	funcs := template.FuncMap{
		"escapeGo":           EscapeGoString,
		"modelFields":        renderResponseOperationModelFields,
		"nestedModels":       renderResponseOperationNestedModels,
		"schemaAttributes":   renderResponseOperationSchema,
		"requestSetup":       renderResponseOperationRequestSetup,
		"invoke":             renderResponseOperationInvoke,
		"responseUnmarshal":  func(*openapi.ResponseOperationTemplate) string { return responseUnmarshal },
		"issuanceIdentifier": renderIssuanceIdentifier,
	}
	tmpl, err := template.New("response operation").Funcs(funcs).Parse(source)
	if err != nil {
		return fmt.Errorf("parse response-operation template: %w", err)
	}
	var output bytes.Buffer
	if err := tmpl.Execute(&output, operation); err != nil {
		return fmt.Errorf("execute response-operation template: %w", err)
	}
	formatted, err := imports.Process(path, output.Bytes(), nil)
	if err != nil {
		return fmt.Errorf("format generated response operation %s: %w", path, err)
	}
	if err := os.WriteFile(path, formatted, 0o644); err != nil {
		return fmt.Errorf("write generated response operation %s: %w", path, err)
	}
	return nil
}

func ensureResponseOperationTarget(path string) error {
	content, err := os.ReadFile(path)
	if err == nil {
		if !bytes.Contains(content, []byte("Code generated")) || !bytes.Contains(content, []byte("DO NOT EDIT")) {
			return fmt.Errorf("refusing to overwrite handwritten file %s", path)
		}
		return nil
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("read response-operation target %s: %w", path, err)
	}
	return nil
}

func responseOperationAttributes(operation *openapi.ResponseOperationTemplate, includeResponse bool) []openapi.TerraformAttribute {
	attributes := make([]openapi.TerraformAttribute, 0, len(operation.Inputs)+len(operation.ResponseAttributes))
	for _, input := range operation.Inputs {
		attributes = append(attributes, input.Attribute)
	}
	if includeResponse {
		attributes = append(attributes, operation.ResponseAttributes...)
	}
	return attributes
}

func renderResponseOperationModelFields(operation *openapi.ResponseOperationTemplate, includeResponse bool) string {
	attributes := responseOperationAttributes(operation, includeResponse)
	var result strings.Builder
	for _, attribute := range attributes {
		if attribute.IsBlock {
			if attribute.NestedBlockType == "list" {
				fmt.Fprintf(&result, "\t%s types.List `tfsdk:%q`\n", attribute.GoName, attribute.TfsdkTag)
			} else {
				typeName := operation.TitleCase + naming.ToResourceTypeName(attribute.TfsdkTag) + "Model"
				fmt.Fprintf(&result, "\t%s *%s `tfsdk:%q`\n", attribute.GoName, typeName, attribute.TfsdkTag)
			}
			continue
		}
		goType := map[string]string{"string": "String", "int64": "Int64", "bool": "Bool", "list": "List", "map": "Map"}[attribute.Type]
		if goType == "" {
			goType = "Dynamic"
		}
		fmt.Fprintf(&result, "\t%s types.%s `tfsdk:%q`\n", attribute.GoName, goType, attribute.TfsdkTag)
	}
	return result.String()
}

func renderResponseOperationNestedModels(operation *openapi.ResponseOperationTemplate, includeResponse bool) string {
	return RenderNestedModelTypes(operation.TitleCase, responseOperationAttributes(operation, includeResponse))
}

func renderResponseOperationSchema(operation *openapi.ResponseOperationTemplate, includeResponse bool) string {
	attributes := responseOperationAttributes(operation, includeResponse)
	if operation.Role == "issuance" {
		attributes = append(attributes, openapi.TerraformAttribute{
			Name: "id", GoName: "ID", TfsdkTag: "id", JsonName: "id", Type: "string", Computed: true,
			Description: "Stable Terraform identifier for this create-once issuance.",
		})
	}
	return renderResponseOperationSchemaMap(attributes, "\t\t", operation.Role == "issuance")
}

func renderResponseOperationSchemaMap(attributes []openapi.TerraformAttribute, indent string, requiresReplace bool) string {
	var result strings.Builder
	result.WriteString(indent + "Attributes: map[string]schema.Attribute{\n")
	for _, attribute := range attributes {
		description := EscapeGoString(attribute.Description)
		if attribute.IsBlock {
			typeName := "SingleNestedAttribute"
			if attribute.NestedBlockType == "list" {
				typeName = "ListNestedAttribute"
			}
			fmt.Fprintf(&result, "%s\t%q: schema.%s{\n", indent, attribute.TfsdkTag, typeName)
			fmt.Fprintf(&result, "%s\t\tMarkdownDescription: %q,\n", indent, description)
			if attribute.NestedBlockType == "list" {
				fmt.Fprintf(&result, "%s\t\tNestedObject: schema.NestedAttributeObject{\n", indent)
				result.WriteString(renderResponseOperationSchemaMap(attribute.NestedAttributes, indent+"\t\t\t", false))
				fmt.Fprintf(&result, "%s\t\t},\n", indent)
			} else {
				result.WriteString(renderResponseOperationSchemaMap(attribute.NestedAttributes, indent+"\t\t", false))
			}
			renderResponseOperationFlags(&result, attribute, indent+"\t\t", false)
			fmt.Fprintf(&result, "%s\t},\n", indent)
			continue
		}
		typeName := map[string]string{"string": "String", "int64": "Int64", "bool": "Bool", "list": "List", "map": "Map"}[attribute.Type]
		if typeName == "" {
			typeName = "Dynamic"
		}
		fmt.Fprintf(&result, "%s\t%q: schema.%sAttribute{\n", indent, attribute.TfsdkTag, typeName)
		fmt.Fprintf(&result, "%s\t\tMarkdownDescription: %q,\n", indent, description)
		renderResponseOperationFlags(&result, attribute, indent+"\t\t", requiresReplace)
		if attribute.Type == "list" || attribute.Type == "map" {
			elementType := map[string]string{"string": "types.StringType", "int64": "types.Int64Type", "bool": "types.BoolType"}[attribute.ElementType]
			if elementType == "" {
				elementType = "types.StringType"
			}
			fmt.Fprintf(&result, "%s\t\tElementType: %s,\n", indent, elementType)
		}
		fmt.Fprintf(&result, "%s\t},\n", indent)
	}
	result.WriteString(indent + "},\n")
	return result.String()
}

func renderResponseOperationFlags(result *strings.Builder, attribute openapi.TerraformAttribute, indent string, requiresReplace bool) {
	if attribute.Required {
		result.WriteString(indent + "Required: true,\n")
	} else {
		if attribute.Optional {
			result.WriteString(indent + "Optional: true,\n")
		}
		if attribute.Computed {
			result.WriteString(indent + "Computed: true,\n")
		}
	}
	if attribute.Sensitive {
		result.WriteString(indent + "Sensitive: true,\n")
	}
	renderResponseOperationValidators(result, attribute, indent)
	if requiresReplace && (attribute.Required || attribute.Optional) {
		typeName := map[string]string{"string": "String", "int64": "Int64", "bool": "Bool", "list": "List", "map": "Map"}[attribute.Type]
		packageName := map[string]string{"string": "stringplanmodifier", "int64": "int64planmodifier", "bool": "boolplanmodifier", "list": "listplanmodifier", "map": "mapplanmodifier"}[attribute.Type]
		if typeName != "" {
			fmt.Fprintf(result, "%sPlanModifiers: []planmodifier.%s{%s.RequiresReplace()},\n", indent, typeName, packageName)
		}
	}
}

func renderResponseOperationValidators(result *strings.Builder, attribute openapi.TerraformAttribute, indent string) {
	if attribute.Type == "string" {
		var validators []string
		switch {
		case attribute.MinLength > 0 && attribute.MaxLength > 0:
			validators = append(validators, fmt.Sprintf("stringvalidator.LengthBetween(%d, %d)", attribute.MinLength, attribute.MaxLength))
		case attribute.MinLength > 0:
			validators = append(validators, fmt.Sprintf("stringvalidator.LengthAtLeast(%d)", attribute.MinLength))
		case attribute.MaxLength > 0:
			validators = append(validators, fmt.Sprintf("stringvalidator.LengthAtMost(%d)", attribute.MaxLength))
		}
		if attribute.Pattern != "" {
			validators = append(validators, fmt.Sprintf("stringvalidator.RegexMatches(regexp.MustCompile(%s), \"\")", RegexLiteral(attribute.Pattern)))
		}
		if len(attribute.EnumValues) > 0 {
			quoted := make([]string, len(attribute.EnumValues))
			for index, value := range attribute.EnumValues {
				quoted[index] = strconv.Quote(value)
			}
			validators = append(validators, "stringvalidator.OneOf("+strings.Join(quoted, ", ")+")")
		}
		if len(validators) > 0 {
			result.WriteString(indent + "Validators: []validator.String{\n")
			for _, expression := range validators {
				fmt.Fprintf(result, "%s\t%s,\n", indent, expression)
			}
			result.WriteString(indent + "},\n")
		}
	}
	if attribute.Type == "list" && (attribute.MinItems > 0 || attribute.MaxItems > 0) {
		result.WriteString(indent + "Validators: []validator.List{\n")
		switch {
		case attribute.MinItems > 0 && attribute.MaxItems > 0:
			fmt.Fprintf(result, "%s\tlistvalidator.SizeBetween(%d, %d),\n", indent, attribute.MinItems, attribute.MaxItems)
		case attribute.MinItems > 0:
			fmt.Fprintf(result, "%s\tlistvalidator.SizeAtLeast(%d),\n", indent, attribute.MinItems)
		default:
			fmt.Fprintf(result, "%s\tlistvalidator.SizeAtMost(%d),\n", indent, attribute.MaxItems)
		}
		result.WriteString(indent + "},\n")
	}
	if attribute.Type != "int64" {
		return
	}
	if len(attribute.Int64RangeSpans) > 0 {
		result.WriteString(indent + "Validators: []validator.Int64{\n")
		result.WriteString(indent + "\tvalidators.Int64RangeSetValidator(\n")
		for _, span := range attribute.Int64RangeSpans {
			fmt.Fprintf(result, "%s\t\tvalidators.Int64Range{Minimum: %d, Maximum: %d},\n", indent, span.Minimum, span.Maximum)
		}
		result.WriteString(indent + "\t),\n" + indent + "},\n")
		return
	}
	if !attribute.HasMinimum && !attribute.HasMaximum {
		return
	}
	result.WriteString(indent + "Validators: []validator.Int64{\n")
	switch {
	case attribute.HasMinimum && attribute.HasMaximum:
		fmt.Fprintf(result, "%s\tint64validator.Between(%d, %d),\n", indent, attribute.Minimum, attribute.Maximum)
	case attribute.HasMinimum:
		fmt.Fprintf(result, "%s\tint64validator.AtLeast(%d),\n", indent, attribute.Minimum)
	default:
		fmt.Fprintf(result, "%s\tint64validator.AtMost(%d),\n", indent, attribute.Maximum)
	}
	result.WriteString(indent + "},\n")
}

func renderResponseOperationRequestSetup(operation *openapi.ResponseOperationTemplate) string {
	var result strings.Builder
	for _, input := range operation.Inputs {
		attribute := input.Attribute
		if attribute.StringDefault != "" {
			fmt.Fprintf(&result, "\tif data.%s.IsNull() { data.%s = types.StringValue(%q) }\n", attribute.GoName, attribute.GoName, attribute.StringDefault)
		}
		if value, ok := attribute.Default.(bool); ok {
			fmt.Fprintf(&result, "\tif data.%s.IsNull() { data.%s = types.BoolValue(%t) }\n", attribute.GoName, attribute.GoName, value)
		}
	}
	fmt.Fprintf(&result, "\tapiPath := %q\n", operation.APIPath)
	for _, input := range operation.Inputs {
		for _, binding := range input.Bindings {
			if binding.Location != "path" {
				continue
			}
			fmt.Fprintf(&result, "\tapiPath = strings.ReplaceAll(apiPath, %q, url.PathEscape(data.%s.ValueString()))\n", "{"+binding.Name+"}", input.Attribute.GoName)
		}
	}
	result.WriteString("\tqueryValues := url.Values{}\n")
	for _, input := range operation.Inputs {
		for _, binding := range input.Bindings {
			if binding.Location != "query" {
				continue
			}
			result.WriteString(renderOperationQueryBinding(input.Attribute, binding.Name))
		}
	}
	result.WriteString("\tif encoded := queryValues.Encode(); encoded != \"\" { apiPath += \"?\" + encoded }\n")
	if operation.Method == "POST" {
		result.WriteString("\tbody := map[string]interface{}{}\n")
		for _, input := range operation.Inputs {
			for _, binding := range input.Bindings {
				if binding.Location != "body" {
					continue
				}
				result.WriteString(renderOperationBodyBinding(input.Attribute, binding.Name))
			}
		}
	}
	return result.String()
}

func renderOperationQueryBinding(attribute openapi.TerraformAttribute, name string) string {
	guard := fmt.Sprintf("!data.%s.IsNull() && !data.%s.IsUnknown()", attribute.GoName, attribute.GoName)
	var value string
	switch attribute.Type {
	case "string":
		value = "data." + attribute.GoName + ".ValueString()"
	case "int64":
		value = "strconv.FormatInt(data." + attribute.GoName + ".ValueInt64(), 10)"
	case "bool":
		value = "strconv.FormatBool(data." + attribute.GoName + ".ValueBool())"
	case "list":
		return fmt.Sprintf("\tif %s { var values []string; data.%s.ElementsAs(ctx, &values, false); for _, value := range values { queryValues.Add(%q, value) } }\n", guard, attribute.GoName, name)
	default:
		return ""
	}
	return fmt.Sprintf("\tif %s { queryValues.Set(%q, %s) }\n", guard, name, value)
}

func renderOperationBodyBinding(attribute openapi.TerraformAttribute, name string) string {
	guard := fmt.Sprintf("!data.%s.IsNull() && !data.%s.IsUnknown()", attribute.GoName, attribute.GoName)
	var value string
	switch attribute.Type {
	case "string":
		value = "data." + attribute.GoName + ".ValueString()"
	case "int64":
		value = "data." + attribute.GoName + ".ValueInt64()"
	case "bool":
		value = "data." + attribute.GoName + ".ValueBool()"
	case "list":
		return fmt.Sprintf("\tif %s { var values []interface{}; data.%s.ElementsAs(ctx, &values, false); body[%q] = values }\n", guard, attribute.GoName, name)
	case "map":
		return fmt.Sprintf("\tif %s { var values map[string]interface{}; data.%s.ElementsAs(ctx, &values, false); body[%q] = values }\n", guard, attribute.GoName, name)
	default:
		return ""
	}
	return fmt.Sprintf("\tif %s { body[%q] = %s }\n", guard, name, value)
}

func renderResponseOperationInvoke(operation *openapi.ResponseOperationTemplate, receiver string, useResult bool) string {
	var result strings.Builder
	resultTarget := "nil"
	if useResult {
		if operation.ResponseIsScalar {
			result.WriteString("\tvar scalarResult interface{}\n")
			resultTarget = "&scalarResult"
		} else {
			result.WriteString("\tapiResult := map[string]interface{}{}\n")
			resultTarget = "&apiResult"
		}
	}
	method := "Get"
	arguments := "ctx, apiPath, " + resultTarget
	if operation.Method == "POST" {
		method = "Post"
		arguments = "ctx, apiPath, body, " + resultTarget
	}
	if operation.Role == "collection" {
		method += "Lenient"
	}
	fmt.Fprintf(&result, "\tif err := %s.client.%s(%s); err != nil {\n", receiver, method, arguments)
	result.WriteString("\t\tresp.Diagnostics.AddError(\"Client Error\", fmt.Sprintf(\"Unable to invoke response operation: %s\", err))\n\t\treturn\n\t}\n")
	if operation.ResponseIsScalar && useResult {
		result.WriteString("\tapiResult := map[string]interface{}{\"result\": scalarResult}\n")
	}
	return result.String()
}

func renderResponseOperationUnmarshal(operation *openapi.ResponseOperationTemplate) (string, error) {
	if operation.Role == "action" {
		return "", nil
	}
	var code strings.Builder
	for _, attribute := range operation.ResponseAttributes {
		var err error
		switch {
		case attribute.IsBlock && attribute.NestedBlockType == "list":
			err = renderUnmarshalTopLevelList(&code, operation.TitleCase, attribute, "\t")
		case attribute.IsBlock:
			err = renderUnmarshalTopLevelSingle(&code, operation.TitleCase, attribute, "\t")
		default:
			err = renderUnmarshalTopLevelScalar(&code, attribute, "\t")
		}
		if err != nil {
			return "", fmt.Errorf("render response operation %s field %q: %w", operation.Name, attribute.Name, err)
		}
	}
	return "\tapiResource := struct{ Spec map[string]interface{} }{Spec: apiResult}\n\tisImport := true\n\t_ = isImport\n" + code.String(), nil
}

func renderIssuanceIdentifier(operation *openapi.ResponseOperationTemplate) string {
	values := []string{strconv.Quote(operation.Name)}
	for _, input := range operation.Inputs {
		attribute := input.Attribute
		switch attribute.Type {
		case "string":
			values = append(values, "data."+attribute.GoName+".ValueString()")
		case "int64":
			values = append(values, "strconv.FormatInt(data."+attribute.GoName+".ValueInt64(), 10)")
		case "bool":
			values = append(values, "strconv.FormatBool(data."+attribute.GoName+".ValueBool())")
		}
	}
	return fmt.Sprintf("\tsum := sha256.Sum256([]byte(strings.Join([]string{%s}, \"\\x00\")))\n\tdata.ID = types.StringValue(hex.EncodeToString(sum[:]))\n", strings.Join(values, ", "))
}

const responseDataSourceTemplate = `// Code generated by generate-all-schemas.go. DO NOT EDIT.
// Source: F5 XC enriched API response-operation contract

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"regexp"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/validators"
)

var _ datasource.DataSource = &{{.TitleCase}}DataSource{}
var _ datasource.DataSourceWithConfigure = &{{.TitleCase}}DataSource{}

func New{{.TitleCase}}DataSource() datasource.DataSource { return &{{.TitleCase}}DataSource{} }
type {{.TitleCase}}DataSource struct { client *client.Client }
type {{.TitleCase}}DataSourceModel struct {
{{modelFields . true -}}
}
{{nestedModels . true -}}
func (d *{{.TitleCase}}DataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) { resp.TypeName = req.ProviderTypeName + "_{{.Name}}" }
func (d *{{.TitleCase}}DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "{{escapeGo .Description}}",
{{schemaAttributes . true -}}
	}
}
func (d *{{.TitleCase}}DataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil { return }
	configured, ok := req.ProviderData.(*client.Client); if !ok { resp.Diagnostics.AddError("Unexpected Data Source Configure Type", "Expected *client.Client"); return }; d.client = configured
}
func (d *{{.TitleCase}}DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data {{.TitleCase}}DataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...); if resp.Diagnostics.HasError() { return }
{{requestSetup . -}}
{{invoke . "d" true -}}
{{responseUnmarshal . -}}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
`

const responseIssuanceTemplate = `// Code generated by generate-all-schemas.go. DO NOT EDIT.
// Source: F5 XC enriched API issuance contract

package provider

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"regexp"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/validators"
)

var _ resource.Resource = &{{.TitleCase}}Resource{}
var _ resource.ResourceWithConfigure = &{{.TitleCase}}Resource{}
func New{{.TitleCase}}Resource() resource.Resource { return &{{.TitleCase}}Resource{} }
type {{.TitleCase}}Resource struct { client *client.Client }
type {{.TitleCase}}ResourceModel struct {
{{modelFields . true -}}
	ID types.String ` + "`tfsdk:\"id\"`" + `
}
{{nestedModels . true -}}
func (r *{{.TitleCase}}Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) { resp.TypeName = req.ProviderTypeName + "_{{.Name}}" }
func (r *{{.TitleCase}}Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "{{escapeGo .Description}}",
{{schemaAttributes . true -}}
	}
}
func (r *{{.TitleCase}}Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) { if req.ProviderData == nil { return }; configured, ok := req.ProviderData.(*client.Client); if !ok { resp.Diagnostics.AddError("Unexpected Resource Configure Type", "Expected *client.Client"); return }; r.client = configured }
func (r *{{.TitleCase}}Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data {{.TitleCase}}ResourceModel; resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...); if resp.Diagnostics.HasError() { return }
{{requestSetup . -}}
{{invoke . "r" true -}}
{{responseUnmarshal . -}}
{{issuanceIdentifier . -}}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *{{.TitleCase}}Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) { var data {{.TitleCase}}ResourceModel; resp.Diagnostics.Append(req.State.Get(ctx, &data)...); if !resp.Diagnostics.HasError() { resp.Diagnostics.Append(resp.State.Set(ctx, &data)...) } }
func (r *{{.TitleCase}}Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) { resp.Diagnostics.AddError("Update Not Supported", "All issuance inputs require replacement.") }
func (r *{{.TitleCase}}Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {}
`

const responseActionTemplate = `// Code generated by generate-all-schemas.go. DO NOT EDIT.
// Source: F5 XC enriched API action contract

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"regexp"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/validators"
)

var _ action.Action = &{{.TitleCase}}Action{}
var _ action.ActionWithConfigure = &{{.TitleCase}}Action{}
func New{{.TitleCase}}Action() action.Action { return &{{.TitleCase}}Action{} }
type {{.TitleCase}}Action struct { client *client.Client }
type {{.TitleCase}}ActionModel struct {
{{modelFields . false -}}
}
{{nestedModels . false -}}
func (a *{{.TitleCase}}Action) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) { resp.TypeName = req.ProviderTypeName + "_{{.Name}}" }
func (a *{{.TitleCase}}Action) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) { resp.Schema = schema.Schema{MarkdownDescription: "{{escapeGo .Description}}",
{{schemaAttributes . false -}}
	} }
func (a *{{.TitleCase}}Action) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) { if req.ProviderData == nil { return }; configured, ok := req.ProviderData.(*client.Client); if !ok { resp.Diagnostics.AddError("Unexpected Action Configure Type", "Expected *client.Client"); return }; a.client = configured }
func (a *{{.TitleCase}}Action) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data {{.TitleCase}}ActionModel; resp.Diagnostics.Append(req.Config.Get(ctx, &data)...); if resp.Diagnostics.HasError() { return }
{{requestSetup . -}}
{{invoke . "a" false -}}
}
`
