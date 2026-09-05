// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package schema

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/conflicts"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/description"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/namespace"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/naming"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

// ExtractConfig holds resource metadata maps that extractResourceSchema needs
// from the monolith's global state.
type ExtractConfig struct {
	TierMap         map[string]string
	DependencyMap   map[string]*openapi.ResourceDependencies
	ReferencedByMap map[string][]string
	CategoryMap     map[string]string
}

// ExtractResponseOperationSchema converts one exact catalog response operation
// into a shared Terraform IR. Path, query, and body bindings remain explicit so
// generation never guesses a wire location from an attribute name.
func ExtractResponseOperationSchema(spec *openapi.Spec, operation openapi.ResolvedResponseOperation) (*openapi.ResponseOperationTemplate, error) {
	if spec == nil {
		return nil, fmt.Errorf("response operation %q requires an OpenAPI spec", operation.Name)
	}
	if operation.Name == "" || operation.Role == "" || operation.Path == "" || operation.OperationID == "" || operation.ResponseSchema == "" {
		return nil, fmt.Errorf("response operation is missing a name, role, path, operation ID, or response schema")
	}
	pathValue, ok := spec.Paths[operation.Path]
	if !ok {
		return nil, fmt.Errorf("response operation %s path %q is absent", operation.Name, operation.Path)
	}
	pathItem, ok := pathValue.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("response operation %s path %q is not an object", operation.Name, operation.Path)
	}
	methodValue, ok := pathItem[strings.ToLower(operation.Method)]
	if !ok {
		return nil, fmt.Errorf("response operation %s method %s is absent", operation.Name, operation.Method)
	}
	method, ok := methodValue.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("response operation %s method %s is not an object", operation.Name, operation.Method)
	}
	if operationID, _ := method["operationId"].(string); operationID != operation.OperationID {
		return nil, fmt.Errorf("response operation %s OpenAPI operationId %q does not match %q", operation.Name, operationID, operation.OperationID)
	}
	if operation.Method == "POST" && operation.RequestSchema == "" {
		return nil, fmt.Errorf("response operation %s POST request schema is required", operation.Name)
	}

	required := make(map[string]bool)
	if rawRequired, ok := method["x-f5xc-required-fields"].([]interface{}); ok {
		for _, raw := range rawRequired {
			name, ok := raw.(string)
			if !ok || name == "" {
				return nil, fmt.Errorf("response operation %s has an invalid required field", operation.Name)
			}
			required[name] = true
		}
	}

	inputsByTag := make(map[string]*openapi.ResponseOperationInput)
	addInput := func(name, location string, property openapi.Schema, isRequired bool) error {
		attr := ConvertToTerraformAttribute(name, property, isRequired, "", spec)
		if attr.ConversionError != "" {
			return fmt.Errorf("response operation %s input %q: %s", operation.Name, name, attr.ConversionError)
		}
		if attr.IsBlock || !supportedResponseOperationInput(attr) {
			return fmt.Errorf("response operation %s input %q has unsupported %s shape", operation.Name, name, location)
		}
		// Operation inputs are practitioner configuration, never response-owned
		// computed values. The one exception below is namespace with a provider
		// default, which must be Optional+Computed so the default can enter state.
		attr.Computed = false
		if !isRequired {
			attr.Required = false
			attr.Optional = true
		}
		if operation.Name == "site_image" && name == "provider" {
			attr.GoName = "ProviderRef"
			attr.TfsdkTag = "provider_ref"
		}
		if name == "namespace" && (operation.Name == "site_registrations_by_site" || operation.Name == "site_registrations_by_state") {
			attr.Required = false
			attr.Optional = true
			attr.Computed = true
			attr.StringDefault = "system"
		}
		if operation.Role == "action" && name == "force" {
			attr.Required = false
			attr.Optional = true
			attr.Computed = false
			attr.Default = false
			if !strings.Contains(strings.ToLower(attr.Description), "default") {
				attr.Description = strings.TrimSpace(attr.Description) + " Defaults to `false`."
			}
		}
		existing := inputsByTag[attr.TfsdkTag]
		if existing == nil {
			input := &openapi.ResponseOperationInput{Attribute: attr}
			inputsByTag[attr.TfsdkTag] = input
			existing = input
		} else if existing.Attribute.Type != attr.Type || existing.Attribute.JsonName != attr.JsonName {
			return fmt.Errorf("response operation %s inputs %q and %q collide as Terraform attribute %q", operation.Name, existing.Attribute.Name, name, attr.TfsdkTag)
		}
		existing.Attribute.Required = existing.Attribute.Required || attr.Required
		if existing.Attribute.Required {
			existing.Attribute.Optional = false
			existing.Attribute.Computed = false
		}
		existing.Bindings = append(existing.Bindings, openapi.OperationBinding{Location: location, Name: name})
		return nil
	}

	parameters := make([]interface{}, 0)
	if pathParameters, ok := pathItem["parameters"].([]interface{}); ok {
		parameters = append(parameters, pathParameters...)
	}
	if operationParameters, ok := method["parameters"].([]interface{}); ok {
		parameters = append(parameters, operationParameters...)
	}
	for _, raw := range parameters {
		parameter, ok := raw.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("response operation %s has a non-object parameter", operation.Name)
		}
		location, _ := parameter["in"].(string)
		if location != "path" && location != "query" {
			continue
		}
		name, _ := parameter["name"].(string)
		if name == "" {
			return nil, fmt.Errorf("response operation %s has an unnamed %s parameter", operation.Name, location)
		}
		property, err := decodeOperationParameterSchema(parameter)
		if err != nil {
			return nil, fmt.Errorf("response operation %s parameter %q: %w", operation.Name, name, err)
		}
		parameterRequired, _ := parameter["required"].(bool)
		if err := addInput(name, location, property, required[name] || parameterRequired); err != nil {
			return nil, err
		}
	}

	if operation.RequestSchema != "" {
		request, ok := spec.Components.Schemas[operation.RequestSchema]
		if !ok {
			return nil, fmt.Errorf("response operation %s request schema %q is absent", operation.Name, operation.RequestSchema)
		}
		if request.Type != "object" {
			return nil, fmt.Errorf("response operation %s request schema %q must be an object", operation.Name, operation.RequestSchema)
		}
		names := sortedSchemaPropertyNames(request.Properties)
		for _, name := range names {
			if err := addInput(name, "body", request.Properties[name], required[name]); err != nil {
				return nil, err
			}
		}
	}
	for name := range required {
		found := false
		for _, input := range inputsByTag {
			for _, binding := range input.Bindings {
				if binding.Name == name {
					found = true
				}
			}
		}
		if !found {
			return nil, fmt.Errorf("response operation %s required field %q has no path, query, or body binding", operation.Name, name)
		}
	}

	response, ok := spec.Components.Schemas[operation.ResponseSchema]
	if !ok {
		return nil, fmt.Errorf("response operation %s response schema %q is absent", operation.Name, operation.ResponseSchema)
	}
	responseAttrs, err := responseOperationAttributes(spec, operation.Name, response)
	if err != nil {
		return nil, err
	}
	for _, attr := range responseAttrs {
		if input := inputsByTag[attr.TfsdkTag]; input != nil {
			return nil, fmt.Errorf("response operation %s input %q collides with response attribute %q", operation.Name, input.Attribute.Name, attr.Name)
		}
	}

	inputs := make([]openapi.ResponseOperationInput, 0, len(inputsByTag))
	for _, input := range inputsByTag {
		inputs = append(inputs, *input)
	}
	sort.Slice(inputs, func(i, j int) bool {
		if inputs[i].Attribute.Required != inputs[j].Attribute.Required {
			return inputs[i].Attribute.Required
		}
		return inputs[i].Attribute.TfsdkTag < inputs[j].Attribute.TfsdkTag
	})

	descriptionSource := ""
	if metadata, ok := method["x-f5xc-operation-metadata"].(map[string]interface{}); ok {
		descriptionSource, _ = metadata["purpose"].(string)
	}
	if descriptionSource == "" {
		descriptionSource, _ = method["description"].(string)
	}
	if descriptionSource == "" {
		descriptionSource, _ = method["summary"].(string)
	}
	descriptionText := description.Clean(descriptionSource, operation.Name)
	if descriptionText == "" {
		descriptionText = fmt.Sprintf("Invoke the %s operation.", naming.ToHumanReadableName(operation.Name))
	} else if !strings.ContainsAny(descriptionText[len(descriptionText)-1:], ".!?") {
		descriptionText += "."
	}
	return &openapi.ResponseOperationTemplate{
		Name: operation.Name, Role: operation.Role, TitleCase: naming.ToResourceTypeName(operation.Name), Method: operation.Method,
		APIPath: operation.Path, OperationID: operation.OperationID, RequestSchema: operation.RequestSchema,
		ResponseSchema: operation.ResponseSchema, Description: descriptionText, Inputs: inputs, ResponseAttributes: responseAttrs,
		ResponseIsScalar: response.Type != "object",
	}, nil
}

func supportedResponseOperationInput(attribute openapi.TerraformAttribute) bool {
	switch attribute.Type {
	case "string", "int64", "bool":
		return true
	case "list", "map":
		return attribute.ElementType == "string" || attribute.ElementType == "int64" || attribute.ElementType == "bool"
	default:
		return false
	}
}

func decodeOperationParameterSchema(parameter map[string]interface{}) (openapi.Schema, error) {
	schemaMap, ok := parameter["schema"].(map[string]interface{})
	if !ok {
		return openapi.Schema{}, fmt.Errorf("schema is absent or not an object")
	}
	encoded, err := json.Marshal(schemaMap)
	if err != nil {
		return openapi.Schema{}, err
	}
	var property openapi.Schema
	if err := json.Unmarshal(encoded, &property); err != nil {
		return openapi.Schema{}, err
	}
	if property.Description == "" {
		property.Description, _ = parameter["description"].(string)
		// F5 path parameters commonly encode a display label, an x-required
		// marker, and the useful sentence on separate lines. Clean() removes
		// the marker; preserve a sentence boundary so "Name" and "Site name"
		// do not collapse into "NameSite name".
		property.Description = strings.ReplaceAll(property.Description, "\r\n\r\nx-required\r\n", ". ")
		property.Description = strings.ReplaceAll(property.Description, "\n\nx-required\n", ". ")
		property.Description = strings.ReplaceAll(property.Description, "\r\nx-required\r\n", ". ")
		property.Description = strings.ReplaceAll(property.Description, "\nx-required\n", ". ")
	}
	return property, nil
}

func responseOperationAttributes(spec *openapi.Spec, operationName string, response openapi.Schema) ([]openapi.TerraformAttribute, error) {
	properties := response.Properties
	if response.Type != "object" {
		if response.Type != "string" && response.Type != "integer" && response.Type != "boolean" {
			return nil, fmt.Errorf("response operation %s has unsupported response schema type %q", operationName, response.Type)
		}
		properties = map[string]openapi.Schema{"result": response}
	}
	attributes := make([]openapi.TerraformAttribute, 0, len(properties))
	for _, name := range sortedSchemaPropertyNames(properties) {
		attr := ConvertToTerraformAttribute(name, properties[name], false, "", spec)
		if attr.ConversionError != "" {
			return nil, fmt.Errorf("response operation %s response field %q: %s", operationName, name, attr.ConversionError)
		}
		markResponseAttributeComputed(&attr)
		attributes = append(attributes, attr)
	}
	return attributes, nil
}

func markResponseAttributeComputed(attr *openapi.TerraformAttribute) {
	attr.Required = false
	attr.Optional = false
	attr.Computed = true
	attr.PlanModifier = ""
	if attr.Sensitive && !strings.Contains(strings.ToLower(attr.Description), "terraform state") {
		attr.Description = strings.TrimSpace(attr.Description) + " This sensitive value is stored in Terraform state; protect state access accordingly."
	}
	for index := range attr.NestedAttributes {
		markResponseAttributeComputed(&attr.NestedAttributes[index])
	}
}

func sortedSchemaPropertyNames(properties map[string]openapi.Schema) []string {
	names := make([]string, 0, len(properties))
	for name := range properties {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// ExtractResourceSchema extracts a Terraform resource schema from an OpenAPI spec.
// extractAPIPath is the operation-catalog resolver supplied by the generator.
// ForceReplaceForCreateDeleteOnly marks every user-settable, non-computed attribute as
// RequiresReplace. Create/delete-only F5 XC resources (those declared in import-id-fields.json —
// e.g. the CSD shape/csd domain objects) support only create/list/delete: there is no PUT/update
// endpoint and no by-name GET. An in-place update therefore 404s ("API Group could not be
// determined for Method: PUT"), so any field change must force delete+create instead of a phantom
// update. Computed-only attributes (e.g. id) are left untouched.
func ForceReplaceForCreateDeleteOnly(attrs []openapi.TerraformAttribute) {
	for i := range attrs {
		a := &attrs[i]
		if (a.Required || a.Optional) && !a.Computed && a.PlanModifier != "RequiresReplace" {
			a.PlanModifier = "RequiresReplace"
		}
	}
}

func ExtractResourceSchema(spec *openapi.Spec, resourceName string, extractAPIPath func(spec *openapi.Spec, resourceName string) (string, string, bool)) (*openapi.ResourceTemplate, error) {
	createSpec, createSpecKey, found, resolveErr := ResolveEnvelopeSchema(spec, resourceName, "CreateSpecType")
	if resolveErr != nil {
		return nil, resolveErr
	}
	if !found {
		return nil, fmt.Errorf("no CreateSpecType found")
	}

	// Extract OneOf groups from x-ves-oneof-field annotations
	oneOfGroups := ExtractOneOfGroups(spec, createSpecKey)

	// Create reverse mapping: field -> group name + all fields in group
	// Also track which field should get the constraint (first alphabetically)
	fieldToOneOf := make(map[string][]string)
	fieldToGroupName := make(map[string]string) // Track the group name for AI-friendly defaults
	fieldIsFirst := make(map[string]bool)       // Only first field in each group gets the constraint
	for groupName, fields := range oneOfGroups {
		// Sort fields to determine which is first
		sortedFields := make([]string, len(fields))
		copy(sortedFields, fields)
		sort.Strings(sortedFields)
		firstField := sortedFields[0]

		for _, field := range fields {
			fieldToOneOf[field] = fields
			fieldToGroupName[field] = groupName
			if field == firstField {
				fieldIsFirst[field] = true
			}
		}
	}

	// Convert properties to Terraform attributes
	attributes := []openapi.TerraformAttribute{}
	requiredSet := make(map[string]bool)
	for _, r := range createSpec.Required {
		requiredSet[r] = true
	}

	for propName, propSchema := range createSpec.Properties {
		oneOfFields := fieldToOneOf[propName]
		groupName := fieldToGroupName[propName]
		attr := ConvertToTerraformAttribute(propName, propSchema, requiredSet[propName], "", spec)
		// Add OneOf constraint hint to description only for the first field in each group
		// Include group name for AI-friendly default recommendations
		if len(oneOfFields) > 1 && fieldIsFirst[propName] {
			attr.Description = description.AddOneOfConstraintWithGroup(attr.Description, groupName, oneOfFields)
		}
		attributes = append(attributes, attr)
	}

	// Sort attributes per HashiCorp documentation standards:
	// Arguments: 1) ID components first, 2) Required alphabetically, 3) Optional alphabetically
	// Attributes: 1) id first, 2) remaining alphabetically
	sort.Slice(attributes, func(i, j int) bool {
		// Computed attributes go after arguments
		if attributes[i].Computed != attributes[j].Computed {
			return !attributes[i].Computed
		}
		// Required before optional
		if attributes[i].Required != attributes[j].Required {
			return attributes[i].Required
		}
		// Alphabetical within each group
		return attributes[i].Name < attributes[j].Name
	})

	// Extract correct API path from OpenAPI spec first to determine if namespace is required
	_, _, hasNamespace := extractAPIPath(spec, resourceName)

	// Add standard metadata attributes in HashiCorp-compliant order:
	// 1. ID components (name, namespace) - these form the resource ID
	// 2. Other required args alphabetically
	// 3. Optional args alphabetically (annotations, labels)
	// 4. Computed attributes (id first)

	// DNS zone resource uses domain names (with dots), not standard names
	useDomainValidator := resourceName == "dns_zone"
	nameDescription := fmt.Sprintf("Name of the %s. Must be unique within the namespace.", naming.ToHumanReadableName(resourceName))
	if useDomainValidator {
		nameDescription = fmt.Sprintf("Domain name for the %s (e.g., example.com). Must be a valid DNS domain name.", naming.ToHumanReadableName(resourceName))
	}

	idComponentAttrs := []openapi.TerraformAttribute{
		{Name: "name", GoName: "Name", TfsdkTag: "name", Type: "string",
			Description: nameDescription,
			Required:    true, PlanModifier: "RequiresReplace", UseDomainValidator: useDomainValidator},
	}

	// Namespace emission is driven by the spec-declared namespace constraint
	// (x-f5xc-namespace-profile). Precedence:
	//   1. Profile restricts the resource to a single namespace AND that constraint is
	//      enforced (verification-gated in the spec: only verified classifications set
	//      enforced=true) -> Optional+Computed, defaulted to that value + OneOf, so it
	//      may be omitted and can't be set wrong.
	//   2. Path has {namespace} but the resource is multi-namespace or its single-namespace
	//      constraint is unverified (enforced=false) -> Required (don't over-restrict).
	//   3. No namespace in the API path -> Optional+Computed (omit/empty).
	if prof, ok := namespace.GetProfile(resourceName); ok && len(prof.Allowed) == 1 && prof.Enforced {
		fixedNS := string(prof.Allowed[0])
		idComponentAttrs = append(idComponentAttrs, openapi.TerraformAttribute{
			Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", Type: "string",
			Description:   fmt.Sprintf("Namespace for the %s. The F5 XC API restricts this resource to the %s namespace; it defaults to that value and may be omitted.", naming.ToHumanReadableName(resourceName), fixedNS),
			Optional:      true,
			Computed:      true,
			StringDefault: fixedNS,
			EnumValues:    []string{fixedNS},
			PlanModifier:  "RequiresReplace"})
	} else if hasNamespace {
		idComponentAttrs = append(idComponentAttrs, openapi.TerraformAttribute{
			Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", Type: "string",
			Description: fmt.Sprintf("Namespace where the %s is created.", naming.ToHumanReadableName(resourceName)),
			Required:    true, PlanModifier: "RequiresReplace"})
	} else {
		idComponentAttrs = append(idComponentAttrs, openapi.TerraformAttribute{
			Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", Type: "string",
			Description: fmt.Sprintf("Namespace for the %s. For this resource type, namespace should be empty or omitted.", naming.ToHumanReadableName(resourceName)),
			Optional:    true, Computed: true, PlanModifier: "UseStateForUnknown"})
	}

	// Optional standard attrs will be sorted with other optionals
	// These match the F5XC schemaObjectCreateMetaType fields from OpenAPI specs
	optionalStdAttrs := []openapi.TerraformAttribute{
		{Name: "annotations", GoName: "Annotations", TfsdkTag: "annotations", Type: "map", ElementType: "string",
			Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.", Optional: true},
		{Name: "description", GoName: "Description", TfsdkTag: "description", Type: "string",
			Description: "Human readable description for the object.", Optional: true},
		{Name: "disable", GoName: "Disable", TfsdkTag: "disable", Type: "bool",
			Description: "A value of true administratively disables the object.", Optional: true},
		{Name: "labels", GoName: "Labels", TfsdkTag: "labels", Type: "map", ElementType: "string",
			Description: "Labels is a user defined key value map that can be attached to resources for organization and filtering.", Optional: true},
	}

	// Computed attrs - id first per HashiCorp standards
	computedAttrs := []openapi.TerraformAttribute{
		{Name: "id", GoName: "ID", TfsdkTag: "id", Type: "string",
			Description: "Unique identifier for the resource.", Computed: true, PlanModifier: "UseStateForUnknown"},
	}

	// Surface the object's server-generated system_metadata.uid as a Computed,
	// read-only `uid` attribute — but only when the resource is opted in
	// (tools/expose-uid.json) AND its response schema actually carries
	// system_metadata.uid. This keeps the mechanism schema-driven yet surgically
	// scoped: `id` is the object name, whereas `uid` is the value a consumer such
	// as a CE registration flow needs. It is not a spec field, so it is excluded
	// from spec marshal/unmarshal and populated directly from the API response.
	exposeUID := openapi.LoadExposeUID(naming.ToResourceTypeName(resourceName)) &&
		ResponseHasSystemMetadataUID(spec, resourceName)
	if exposeUID {
		attr := SystemMetadataUIDAttribute(resourceName)
		attr.PlanModifier = "UseStateForUnknown"
		computedAttrs = append(computedAttrs, attr)
	}

	// Combine: ID components first, then other required, then optional (incl. standard), then computed
	var sortedAttrs []openapi.TerraformAttribute
	sortedAttrs = append(sortedAttrs, idComponentAttrs...)

	// Add remaining required attributes (alphabetically)
	for _, attr := range attributes {
		if attr.Required && !attr.Computed {
			sortedAttrs = append(sortedAttrs, attr)
		}
	}

	// Add optional attributes (standard + schema-derived, alphabetically)
	// First, filter out standard attrs that already exist in schema-derived attrs to avoid duplicates
	schemaOptional := FilterOptional(attributes)
	schemaAttrNames := make(map[string]bool)
	for _, attr := range schemaOptional {
		schemaAttrNames[attr.Name] = true
	}
	var filteredStdAttrs []openapi.TerraformAttribute
	for _, stdAttr := range optionalStdAttrs {
		if !schemaAttrNames[stdAttr.Name] {
			filteredStdAttrs = append(filteredStdAttrs, stdAttr)
		}
	}
	allOptional := append(filteredStdAttrs, schemaOptional...)
	sort.Slice(allOptional, func(i, j int) bool {
		return allOptional[i].Name < allOptional[j].Name
	})
	sortedAttrs = append(sortedAttrs, allOptional...)

	// Add computed attributes (id first, then others alphabetically)
	sortedAttrs = append(sortedAttrs, computedAttrs...)
	for _, attr := range attributes {
		if attr.Computed && attr.Name != "id" {
			sortedAttrs = append(sortedAttrs, attr)
		}
	}

	attributes = sortedAttrs

	// Create/delete-only resources (declared in import-id-fields.json, because their create-only
	// spec fields are not readable via the 501 by-name GET) also lack a PUT/update endpoint on the
	// F5 XC API. Force every settable field to RequiresReplace so terraform reconciles via
	// delete+create instead of a phantom in-place update that 404s.
	if len(openapi.LoadImportIDFields(naming.ToResourceTypeName(resourceName))) > 0 {
		ForceReplaceForCreateDeleteOnly(attributes)
	}

	concurrencyToken, err := ExtractConcurrencyTokenContract(spec, resourceName)
	if err != nil {
		return nil, err
	}

	// Get best description with enrichment extension priority:
	// 1. x-f5xc-description-medium (preferred - detailed but concise)
	// 2. x-f5xc-description-short (fallback - ultra-short)
	// 3. description (original)
	bestDescription := createSpec.XF5XCDescriptionMed
	if bestDescription == "" {
		bestDescription = createSpec.XF5XCDescriptionShort
	}
	if bestDescription == "" {
		bestDescription = createSpec.Description
	}
	resourceDescription := description.TransformResourceDescription(resourceName, bestDescription)

	// Generate example usage HCL
	exampleUsage := GenerateExampleUsage(resourceName, attributes)

	// Generate API docs URL
	apiDocsURL := fmt.Sprintf("https://docs.cloud.f5.com/docs/api/%s", strings.ReplaceAll(resourceName, "_", "-"))

	// Extract correct API path from OpenAPI spec
	apiPath, apiPathItem, hasNamespace := extractAPIPath(spec, resourceName)

	// Scan attributes to determine which plan modifier imports are needed
	usesBool, usesInt64, usesString, usesList, usesMap := ScanPlanModifierUsage(attributes)

	// Check if the resource has any nested models that would generate AttrTypes
	// AttrTypes (which use attr.Type) are generated for any block with nested attributes
	hasBlocks := HasNestedModelsWithAttrTypes(attributes)

	// Check for max length validators (including nested attributes)
	hasMaxLengthValidators := HasMaxLengthValidatorsAny(attributes)

	// Collect conflict attributes and generate ValidateConfig checks
	conflictAttrs, goNameLookup := CollectConflictAttrs(attributes)
	conflictCode := conflicts.GenerateChecks(conflictAttrs, goNameLookup)

	// Compile declared apply-time prerequisites into this resource. The source of
	// truth is x-f5xc-requires in the enriched spec: DeriveRequirementPreflights
	// turns each structured cross-resource requirement into a preflight, resolving
	// the required resource's LIST path from the spec's own paths. Hand-maintained
	// preflight-requirements.json entries override the derived ones by trigger
	// field, so the file is an override/supplement rather than the sole source.
	// Each preflight's trigger field is then resolved to its Go model field so
	// Create/Update can nil-check it; those whose trigger is not a top-level
	// attribute are dropped.
	titleCase := naming.ToResourceTypeName(resourceName)
	// extractAPIPath already returns the collection path with the namespace
	// placeholder substituted to %s (e.g. /api/shape/csd/namespaces/%s/protected_domains),
	// which is exactly the LIST-path form a preflight needs. Require a namespaced
	// path with a single %s so we never emit a malformed list_path.
	resolveListPath := func(resource string) (string, bool) {
		p, _, hasNamespace := extractAPIPath(spec, resource)
		if p == "" || !hasNamespace || strings.Count(p, "%s") != 1 {
			return "", false
		}
		return p, true
	}
	derivedPreflights := openapi.DeriveRequirementPreflights(createSpec, resolveListPath)
	mergedPreflights := openapi.MergePreflights(derivedPreflights, openapi.LoadPreflights(titleCase))
	preflights := ResolvePreflightGoFields(mergedPreflights, attributes)

	return &openapi.ResourceTemplate{
		Name:                resourceName,
		TitleCase:           titleCase,
		Preflights:          preflights,
		ImportIDExtraFields: openapi.LoadImportIDFields(titleCase),
		ExposeUID:           exposeUID,
		// #1391: only the resources F5 XC decorates with hardware/OS discovery labels
		// filter those six unprefixed keys out of the read-back.
		FiltersDiscoveredSiteLabels: openapi.LoadDiscoveredSiteLabels(titleCase),
		HasConcurrencyToken:         concurrencyToken != nil,
		ConcurrencyTokenJSONName:    concurrencyTokenJSONName(concurrencyToken),
		ConcurrencyTokenGoName:      concurrencyTokenGoName(concurrencyToken),
		APIPath:                     apiPath,
		APIPathPlural:               resourceName + "s",
		APIPathItem:                 apiPathItem,
		HasNamespaceInPath:          hasNamespace,
		Description:                 resourceDescription,
		Attributes:                  attributes,
		OneOfGroups:                 oneOfGroups, // Now properly preserving extracted OneOf groups
		ExampleUsage:                exampleUsage,
		APIDocsURL:                  apiDocsURL,
		UsesBoolPlanModifier:        usesBool,
		UsesInt64PlanModifier:       usesInt64,
		UsesStringPlanModifier:      usesString,
		UsesListPlanModifier:        usesList,
		UsesMapPlanModifier:         usesMap,
		HasBlocks:                   hasBlocks,
		HasMaxLengthValidators:      hasMaxLengthValidators,
		HasEnumValidators:           HasEnumValidatorsAny(attributes),
		HasPatternValidators:        HasPatternValidatorsAny(attributes),
		HasListSizeValidators:       HasListSizeValidatorsAny(attributes),
		HasInt64RangeValidators:     HasInt64RangeValidatorsAny(attributes),
		HasStringDefaults:           HasStringDefaultsAny(attributes),
		HasConflicts:                conflictCode != "",
		ConflictCheckCode:           conflictCode,
	}, nil
}

// ExtractActionResourceSchema builds an action-style resource template from a
// request-body schema carrying a schema-level x-f5xc-action. Unlike a CRUD
// resource it has no CreateSpecType/GetSpecType: its attributes derive directly
// from the flat action request body. Only scalar string props (plus the `state`
// enum) become attributes; decorative object props (annotations, labels) and
// non-enum $ref props (tunnel_type) are skipped because they are not
// representable as plain string attributes.
//
// A skipped prop the action REQUIRES is a different matter, and dropping one is
// how #1355 shipped: the filter keyed on the Go type, so the required `passport`
// object vanished and every approve POST came back 500 "Validation approval:
// Passport is required" with no user workaround. Two things now prevent that.
// Props declared in tools/action-derived-fields.json are carried as
// ActionDerivedFields — read off the object being acted on at Create rather than
// exposed as attributes, because the API accepts only the value it already
// holds. And any remaining REQUIRED prop that is neither an attribute nor a
// declared derived field fails extraction outright, so the class cannot recur
// silently for a different action.
//
// `namespace` is a path parameter (not a body property) and is injected so the
// generated model carries a Namespace field for the action/read path Sprintf.
// The singular action POST path and the pluralized sibling object GET path are
// captured for Create/Read, `state` constant-defaults to APPROVED, and every
// user-settable field forces replace (there is no in-place update).
// action carries exact POST and sibling GET paths resolved from apiOperations.
func ExtractActionResourceSchema(spec *openapi.Spec, action openapi.ResourcePath) (*openapi.ResourceTemplate, error) {
	resourceName := action.ResourceName
	if resourceName == "" || action.SchemaName == "" || action.ActionPath == "" || action.ReadObjectPath == "" {
		return nil, fmt.Errorf("incomplete catalog action for %q", resourceName)
	}

	reqSchema, ok := spec.Components.Schemas[action.SchemaName]
	if !ok {
		return nil, fmt.Errorf("action request schema %s not found", action.SchemaName)
	}

	// Derive attributes from the request-body properties (deterministic order).
	propNames := make([]string, 0, len(reqSchema.Properties))
	for name := range reqSchema.Properties {
		propNames = append(propNames, name)
	}
	sort.Strings(propNames)

	titleCase := naming.ToResourceTypeName(resourceName)

	// Fields the action requires but that are facts about the object being acted
	// on: read at Create off the sibling object, never exposed as attributes.
	derived, err := openapi.LoadActionDerivedFields(titleCase)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", resourceName, err)
	}
	isDerived := make(map[string]bool, len(derived))
	for _, d := range derived {
		if _, ok := reqSchema.Properties[d.Field]; !ok {
			return nil, fmt.Errorf("%s: action-derived-fields.json declares %q, which the %s request body does not define", resourceName, d.Field, action.SchemaName)
		}
		isDerived[d.Field] = true
	}

	// The action model is rendered as a flat set of string attributes. Only
	// scalar string props (plus the `state` enum $ref, special-cased below) map
	// cleanly to a StringAttribute. Object props (annotations, labels) and
	// non-enum $ref props (passport, tunnel_type) are structured values that must
	// NOT be emitted as strings, so they are skipped as attributes — a required
	// one is picked up by the derived-field declaration or the guard below.
	var attributes []openapi.TerraformAttribute
	haveNamespace := false
	for _, name := range propNames {
		prop := reqSchema.Properties[name]
		isState := name == "state"
		if prop.Type != "string" && !isState {
			continue
		}
		if isDerived[name] {
			// Server-derived: the API accepts only its own value, so it must not
			// become user input even when it would render as a string.
			continue
		}
		if name == "namespace" {
			haveNamespace = true
		}
		attr := openapi.TerraformAttribute{
			Name:        name,
			GoName:      naming.ToResourceTypeName(name),
			TfsdkTag:    naming.ToSnakeCaseTerraform(name),
			Type:        "string",
			JsonName:    prop.WireName(name), // x-f5xc-wire-name override, else the property name (#1323)
			IsSpecField: true,
		}
		switch name {
		case "name", "namespace":
			// Path parameters ({namespace}/{name}); the request schema's `required`
			// list is null, so force these Required.
			attr.Required = true
		case "state":
			// state constant-defaults to APPROVED; Optional+Computed so it may be
			// omitted yet always materializes to the action's target state.
			attr.Optional = true
			attr.Computed = true
			attr.StringDefault = "APPROVED"
		default:
			attr.Optional = true
		}
		attributes = append(attributes, attr)
	}

	// `namespace` is a PATH parameter ({namespace}), not a request-body property,
	// so it is absent from reqSchema.Properties. Inject it (mirroring how CRUD
	// resources declare namespace) so the generated model carries a Namespace
	// field for the action/read path Sprintf. Placed first to keep it adjacent to
	// the other ID component.
	if !haveNamespace {
		attributes = append([]openapi.TerraformAttribute{{
			Name:        "namespace",
			GoName:      "Namespace",
			TfsdkTag:    "namespace",
			Type:        "string",
			JsonName:    "namespace",
			Required:    true,
			IsSpecField: false,
		}}, attributes...)
	}

	// Guard (#1355): a property the action REQUIRES must reach the request body,
	// as an attribute or as a declared server-derived field. Silently skipping
	// one produces a resource that can never create, and the failure surfaces
	// only as an opaque server error against a live tenant — so fail generation
	// here instead, naming the property and both ways to resolve it.
	covered := make(map[string]bool, len(attributes)+len(derived))
	for _, a := range attributes {
		covered[a.Name] = true
	}
	for name := range isDerived {
		covered[name] = true
	}
	for _, name := range propNames {
		if covered[name] || !actionPropIsRequired(reqSchema, name) {
			continue
		}
		return nil, fmt.Errorf(
			"%s: request property %q is required by the %s action but reaches neither a Terraform attribute nor the request body; declare it in tools/action-derived-fields.json if the API derives it from the object, or make it representable as an attribute",
			resourceName, name, action.ActionValue)
	}

	// An action has no PUT/update endpoint, so force every settable field to
	// replace. ForceReplaceForCreateDeleteOnly skips Computed attributes, but the
	// Optional+Computed `state` must force replace too, so backfill afterwards.
	ForceReplaceForCreateDeleteOnly(attributes)
	for i := range attributes {
		attributes[i].PlanModifier = "RequiresReplace"
	}

	return &openapi.ResourceTemplate{
		Name:                resourceName,
		TitleCase:           titleCase,
		Description:         description.TransformResourceDescription(resourceName, reqSchema.Description),
		Attributes:          attributes,
		HasNamespaceInPath:  true,
		HasStringDefaults:   true,
		IsAction:            true,
		ActionPath:          action.ActionPath,
		ActionState:         "APPROVED",
		ReadObjectPath:      action.ReadObjectPath,
		ActionDerivedFields: derived,
	}, nil
}

// actionPropIsRequired reports whether the action request body requires the
// property. F5 expresses that three different ways across the spec pipeline —
// the OpenAPI `required` list, the enrichment's x-f5xc-required-for.create, and
// the upstream x-ves-required marker — and any of them is enough to make an
// omitted field a server-side rejection. x-f5xc-minimum-configuration
// .required_fields is deliberately NOT consulted: on the approve schema it lists
// every property including the decorative annotations and labels, so it does not
// distinguish what the API actually enforces.
func actionPropIsRequired(reqSchema openapi.Schema, name string) bool {
	if reqSchema.IsRequired(name) {
		return true
	}
	prop, ok := reqSchema.Properties[name]
	if !ok {
		return false
	}
	return prop.XF5XCRequiredFor.Create || strings.EqualFold(prop.XVesRequired, "true")
}

// ResolvePreflightGoFields binds each preflight's declared trigger field (WhenField,
// a tfsdk tag such as "client_side_defense") to its generated Go model field name
// (WhenGoField, e.g. "ClientSideDefense") using the resource's top-level attributes.
// A preflight whose trigger is not a top-level attribute is dropped so the generator
// never emits a reference to a field that does not exist.
func ResolvePreflightGoFields(preflights []openapi.RequirementPreflight, attributes []openapi.TerraformAttribute) []openapi.RequirementPreflight {
	if len(preflights) == 0 {
		return nil
	}
	goByTfsdk := make(map[string]string, len(attributes))
	for _, a := range attributes {
		goByTfsdk[a.TfsdkTag] = a.GoName
	}
	resolved := make([]openapi.RequirementPreflight, 0, len(preflights))
	for _, p := range preflights {
		goName, ok := goByTfsdk[p.WhenField]
		if !ok || goName == "" {
			continue
		}
		p.WhenGoField = goName
		resolved = append(resolved, p)
	}
	return resolved
}

// ExtractReadOnlyResourceSchema extracts a data-source-only schema from a GetSpecType.
// All spec properties become Computed attributes. No plan modifiers or conflict checks.
func ExtractReadOnlyResourceSchema(spec *openapi.Spec, resourceName string, extractAPIPath func(spec *openapi.Spec, resourceName string) (string, string, bool)) (*openapi.ResourceTemplate, error) {
	getSpec, _, found, resolveErr := ResolveEnvelopeSchema(spec, resourceName, "GetSpecType")
	if resolveErr != nil {
		return nil, resolveErr
	}
	if !found {
		return nil, fmt.Errorf("no GetSpecType found for %s", resourceName)
	}

	// Convert spec properties to Computed-only attributes
	var attributes []openapi.TerraformAttribute
	for propName, propSchema := range getSpec.Properties {
		if IsMetadataField(propName) {
			continue
		}
		attr := ConvertToTerraformAttribute(propName, propSchema, false, "", spec)
		attr.Required = false
		attr.Optional = false
		attr.Computed = true
		attr.PlanModifier = ""
		attributes = append(attributes, attr)
	}
	sort.Slice(attributes, func(i, j int) bool {
		return attributes[i].Name < attributes[j].Name
	})

	apiPath, apiPathItem, hasNamespace := extractAPIPath(spec, resourceName)

	// Standard metadata attributes for data sources
	metaAttrs := []openapi.TerraformAttribute{
		{Name: "name", GoName: "Name", TfsdkTag: "name", Type: "string",
			Description: fmt.Sprintf("Name of the %s to look up.", naming.ToHumanReadableName(resourceName)),
			Required:    true},
	}
	if hasNamespace {
		metaAttrs = append(metaAttrs, openapi.TerraformAttribute{
			Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", Type: "string",
			Description: fmt.Sprintf("Namespace of the %s.", naming.ToHumanReadableName(resourceName)),
			Required:    true})
	} else {
		metaAttrs = append(metaAttrs, openapi.TerraformAttribute{
			Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", Type: "string",
			Description: "Namespace.", Optional: true, Computed: true})
	}

	computedMeta := []openapi.TerraformAttribute{
		{Name: "id", GoName: "ID", TfsdkTag: "id", Type: "string",
			Description: "Unique identifier.", Computed: true},
		{Name: "description", GoName: "Description", TfsdkTag: "description", Type: "string",
			Description: "Description of the resource.", Computed: true},
		{Name: "labels", GoName: "Labels", TfsdkTag: "labels", Type: "map", ElementType: "string",
			Description: "User-defined labels.", Computed: true},
		{Name: "annotations", GoName: "Annotations", TfsdkTag: "annotations", Type: "map", ElementType: "string",
			Description: "Annotations.", Computed: true},
	}

	exposeUID := openapi.LoadExposeUID(naming.ToResourceTypeName(resourceName)) &&
		ResponseHasSystemMetadataUID(spec, resourceName)
	if exposeUID {
		computedMeta = append(computedMeta, SystemMetadataUIDAttribute(resourceName))
	}

	var allAttrs []openapi.TerraformAttribute
	allAttrs = append(allAttrs, metaAttrs...)
	allAttrs = append(allAttrs, computedMeta...)
	allAttrs = append(allAttrs, attributes...)

	bestDescription := getSpec.Description
	resourceDescription := description.TransformResourceDescription(resourceName, bestDescription)

	return &openapi.ResourceTemplate{
		Name:               resourceName,
		TitleCase:          naming.ToResourceTypeName(resourceName),
		APIPath:            apiPath,
		APIPathPlural:      resourceName + "s",
		APIPathItem:        apiPathItem,
		HasNamespaceInPath: hasNamespace,
		Description:        resourceDescription,
		Attributes:         allAttrs,
		IsReadOnly:         true,
		ExposeUID:          exposeUID,
	}, nil
}

// SystemMetadataUIDAttribute constructs the Terraform attribute for the system_metadata.uid field.
// It applies sensitivity and state-file warnings appropriately for resources like Token.
func SystemMetadataUIDAttribute(resourceName string) openapi.TerraformAttribute {
	uidDesc := "Server-generated unique identifier (`system_metadata.uid`). Read-only; assigned by F5 Distributed Cloud on creation."
	uidSens := false
	if naming.ToResourceTypeName(resourceName) == "Token" {
		uidDesc = "Effective sensitive CE registration credential. NORMAL tokens use `system_metadata.uid`; JWT tokens use `spec.content`. This value is stored in plain text in the Terraform state file; ensure your state file is properly secured."
		uidSens = true
	}
	return openapi.TerraformAttribute{
		Name:        "uid",
		GoName:      "Uid",
		TfsdkTag:    "uid",
		Type:        "string",
		Description: uidDesc,
		Computed:    true,
		Sensitive:   uidSens,
		IsSpecField: false,
	}
}

// ResponseHasSystemMetadataUID reports whether the resource's API response
// envelope carries a system_metadata object with a uid property. F5 XC response
// schemas wrap system_metadata as an allOf-$ref to a shared system-metadata type
// (e.g. schemaSystemObjectGetMetaType) whose properties include uid. This is the
// schema guard behind ExposeUID: the generator only surfaces uid when the object
// can actually return it. Checks the Get and Create response envelopes and the
// object schema, resolving refs via ResolveRef (spec components then SchemaCache).
func ResponseHasSystemMetadataUID(spec *openapi.Spec, resourceName string) bool {
	candidates := []string{
		resourceName + "GetResponse",
		"schema" + resourceName + "GetResponse",
		resourceName + "CreateResponse",
		"schema" + resourceName + "CreateResponse",
		resourceName + "Object",
	}
	for _, name := range candidates {
		envelope, ok := spec.Components.Schemas[name]
		if !ok {
			continue
		}
		sm, ok := envelope.Properties["system_metadata"]
		if !ok {
			continue
		}
		// system_metadata is typically an allOf-wrapped $ref; unwrap it.
		if sm.Ref == "" && len(sm.AllOf) > 0 {
			for _, item := range sm.AllOf {
				if item.Ref != "" {
					sm.Ref = item.Ref
					break
				}
			}
		}
		target := sm
		if sm.Ref != "" {
			target = ResolveRef(sm.Ref, spec)
		}
		if _, ok := target.Properties["uid"]; ok {
			return true
		}
	}
	return false
}

// ExtractOneOfGroups extracts x-ves-oneof-field annotations from the raw schema JSON.
func ExtractOneOfGroups(spec *openapi.Spec, schemaKey string) map[string][]string {
	oneOfGroups := make(map[string][]string)

	// Get raw schema from cache
	rawSchema, ok := RawSpecCache[schemaKey]
	if !ok {
		return oneOfGroups
	}

	// Look for x-ves-oneof-field-* in the raw schema
	for key, value := range rawSchema {
		if strings.HasPrefix(key, "x-ves-oneof-field-") {
			groupName := strings.TrimPrefix(key, "x-ves-oneof-field-")
			// Value can be either a JSON array string or actual array
			switch v := value.(type) {
			case string:
				// Parse JSON array format: "[\"field1\",\"field2\"]"
				v = strings.Trim(v, "[]")
				fields := strings.Split(v, ",")
				for i, f := range fields {
					fields[i] = strings.Trim(strings.TrimSpace(f), "\"")
				}
				oneOfGroups[groupName] = fields
			case []interface{}:
				fields := make([]string, len(v))
				for i, f := range v {
					if s, ok := f.(string); ok {
						fields[i] = s
					}
				}
				oneOfGroups[groupName] = fields
			}
		}
	}

	return oneOfGroups
}
