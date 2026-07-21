// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package schema

import (
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

// ExtractResourceSchema extracts a Terraform resource schema from an OpenAPI spec.
// extractAPIPath is passed as a callback because it remains in the monolith.
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
	// Find CreateSpecType schema
	var createSpec openapi.Schema
	var found bool
	var createSpecKey string

	// Try exact match first (most specs), then schema-prefixed variant
	exactPatterns := []string{
		resourceName + "CreateSpecType",
		"schema" + resourceName + "CreateSpecType",
	}
	for _, pattern := range exactPatterns {
		if schema, ok := spec.Components.Schemas[pattern]; ok {
			createSpec = schema
			createSpecKey = pattern
			found = true
			break
		}
	}

	if !found {
		// Sort schema keys for deterministic fallback (map order is random in Go)
		schemaKeys := make([]string, 0, len(spec.Components.Schemas))
		for key := range spec.Components.Schemas {
			schemaKeys = append(schemaKeys, key)
		}
		sort.Strings(schemaKeys)

		for _, key := range schemaKeys {
			keyLower := strings.ToLower(key)
			if strings.Contains(keyLower, strings.ToLower(resourceName)) &&
				strings.Contains(keyLower, "createspectype") {
				createSpec = spec.Components.Schemas[key]
				createSpecKey = key
				found = true
				break
			}
		}
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

	// DNS zone and DNS domain resources use domain names (with dots), not standard names
	useDomainValidator := resourceName == "dns_zone" || resourceName == "dns_domain"
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
			Description: fmt.Sprintf("Namespace where the %s will be created.", naming.ToHumanReadableName(resourceName)),
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
			Description: "A value of true will administratively disable the object.", Optional: true},
		{Name: "labels", GoName: "Labels", TfsdkTag: "labels", Type: "map", ElementType: "string",
			Description: "Labels is a user defined key value map that can be attached to resources for organization and filtering.", Optional: true},
	}

	// Computed attrs - id first per HashiCorp standards
	computedAttrs := []openapi.TerraformAttribute{
		{Name: "id", GoName: "ID", TfsdkTag: "id", Type: "string",
			Description: "Unique identifier for the resource.", Computed: true, PlanModifier: "UseStateForUnknown"},
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

	// Apply x-f5xc-minimum-configuration to improve Required field accuracy
	minConfigFields := ParseMinConfigRequiredFields(createSpec.XF5XCMinimumConfiguration)
	if len(minConfigFields) > 0 {
		minFieldSet := make(map[string]bool, len(minConfigFields))
		for _, f := range minConfigFields {
			minFieldSet[f] = true
		}
		PromoteMinConfigRequired(attributes, minFieldSet)
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
		Name:                    resourceName,
		TitleCase:               titleCase,
		Preflights:              preflights,
		ImportIDExtraFields:     openapi.LoadImportIDFields(titleCase),
		APIPath:                 apiPath,
		APIPathPlural:           resourceName + "s",
		APIPathItem:             apiPathItem,
		HasNamespaceInPath:      hasNamespace,
		Description:             resourceDescription,
		Attributes:              attributes,
		OneOfGroups:             oneOfGroups, // Now properly preserving extracted OneOf groups
		ExampleUsage:            exampleUsage,
		APIDocsURL:              apiDocsURL,
		UsesBoolPlanModifier:    usesBool,
		UsesInt64PlanModifier:   usesInt64,
		UsesStringPlanModifier:  usesString,
		UsesListPlanModifier:    usesList,
		UsesMapPlanModifier:     usesMap,
		HasBlocks:               hasBlocks,
		HasMaxLengthValidators:  hasMaxLengthValidators,
		HasEnumValidators:       HasEnumValidatorsAny(attributes),
		HasPatternValidators:    HasPatternValidatorsAny(attributes),
		HasListSizeValidators:   HasListSizeValidatorsAny(attributes),
		HasInt64RangeValidators: HasInt64RangeValidatorsAny(attributes),
		HasStringDefaults:       HasStringDefaultsAny(attributes),
		HasConflicts:            conflictCode != "",
		ConflictCheckCode:       conflictCode,
	}, nil
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
	// Find GetSpecType schema with all prefix variants
	getPatterns := []string{
		resourceName + "GetSpecType",
		"schema" + resourceName + "GetSpecType",
		"views" + resourceName + "GetSpecType",
	}
	var getSpec openapi.Schema
	var found bool
	for _, pattern := range getPatterns {
		if schema, ok := spec.Components.Schemas[pattern]; ok {
			getSpec = schema
			found = true
			break
		}
	}
	if !found {
		// Fallback: case-insensitive
		lowerName := strings.ToLower(resourceName)
		schemaKeys := make([]string, 0, len(spec.Components.Schemas))
		for key := range spec.Components.Schemas {
			schemaKeys = append(schemaKeys, key)
		}
		sort.Strings(schemaKeys)
		for _, key := range schemaKeys {
			if strings.Contains(strings.ToLower(key), lowerName) && strings.Contains(strings.ToLower(key), "getspectype") {
				getSpec = spec.Components.Schemas[key]
				found = true
				break
			}
		}
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
	}, nil
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
