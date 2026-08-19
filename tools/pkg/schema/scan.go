// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package schema

import (
	"slices"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/conflicts"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

// ApplyOneOfConflicts translates x-ves-oneof-field groups into pairwise
// ConflictsWith relationships consumed by generated ValidateConfig checks.
func ApplyOneOfConflicts(attributes []openapi.TerraformAttribute, fieldToOneOf map[string][]string) {
	for i := range attributes {
		for _, peer := range fieldToOneOf[attributes[i].TfsdkTag] {
			if peer != attributes[i].TfsdkTag && !slices.Contains(attributes[i].ConflictsWith, peer) {
				attributes[i].ConflictsWith = append(attributes[i].ConflictsWith, peer)
			}
		}
	}
}

// HasNestedModelsWithAttrTypes checks recursively if any nested blocks would generate AttrTypes.
// This is needed to determine if the attr import is required.
// AttrTypes are generated for any nested model that has ANY nested attributes (block or non-block).
func HasNestedModelsWithAttrTypes(attributes []openapi.TerraformAttribute) bool {
	for _, attr := range attributes {
		if attr.IsBlock {
			// If this block has ANY nested attributes, AttrTypes will be generated for it
			if len(attr.NestedAttributes) > 0 {
				return true
			}
			// Note: Even if NestedAttributes is empty, we don't need to recurse
			// because empty blocks use EmptyModel which doesn't have AttrTypes
		}
	}
	return false
}

// HasMaxLengthValidatorsAny recursively checks if any attribute (top-level or nested)
// has a non-zero MaxLength. This determines whether stringvalidator must be imported.
func HasMaxLengthValidatorsAny(attributes []openapi.TerraformAttribute) bool {
	for _, attr := range attributes {
		if attr.MaxLength > 0 {
			return true
		}
		if len(attr.NestedAttributes) > 0 {
			if HasMaxLengthValidatorsAny(attr.NestedAttributes) {
				return true
			}
		}
	}
	return false
}

// HasEnumValidatorsAny recursively checks if any attribute (top-level or nested)
// has EnumValues. This determines whether stringvalidator must be imported for OneOf.
func HasEnumValidatorsAny(attributes []openapi.TerraformAttribute) bool {
	for _, attr := range attributes {
		if len(attr.EnumValues) > 0 {
			return true
		}
		if HasEnumValidatorsAny(attr.NestedAttributes) {
			return true
		}
	}
	return false
}

// HasStringDefaultsAny recursively checks if any attribute (top-level or nested)
// has a StringDefault. This determines whether stringdefault must be imported.
func HasStringDefaultsAny(attributes []openapi.TerraformAttribute) bool {
	for _, attr := range attributes {
		if attr.StringDefault != "" {
			return true
		}
		if HasStringDefaultsAny(attr.NestedAttributes) {
			return true
		}
	}
	return false
}

// HasPatternValidatorsAny recursively checks if any attribute (top-level or nested)
// has a Pattern regex. This determines whether regexp must be imported.
func HasPatternValidatorsAny(attributes []openapi.TerraformAttribute) bool {
	for _, attr := range attributes {
		if attr.Pattern != "" {
			return true
		}
		if HasPatternValidatorsAny(attr.NestedAttributes) {
			return true
		}
	}
	return false
}

// HasListSizeValidatorsAny recursively checks if any non-block list/set attribute (top-level or nested)
// has MinItems or MaxItems constraints. This determines whether listvalidator must be imported.
// Only non-block attributes are checked because blocks use schema.ListNestedBlock which
// does not support the same validator interface.
func HasListSizeValidatorsAny(attributes []openapi.TerraformAttribute) bool {
	for _, attr := range attributes {
		if !attr.IsBlock && (attr.MinItems > 0 || attr.MaxItems > 0) && (attr.Type == "list" || attr.Type == "set") {
			return true
		}
		if HasListSizeValidatorsAny(attr.NestedAttributes) {
			return true
		}
	}
	return false
}

// CollectConflictAttrs collects top-level attributes and blocks that have
// ConflictsWith relationships. Generated validation checks pointer presence for
// blocks and framework null/unknown state for scalar attributes.
func CollectConflictAttrs(attributes []openapi.TerraformAttribute) ([]conflicts.Attr, map[string]conflicts.Field) {
	fieldLookup := make(map[string]conflicts.Field)
	for _, attr := range attributes {
		fieldLookup[attr.TfsdkTag] = conflicts.Field{
			GoName:    attr.GoName,
			IsPointer: attr.IsBlock && attr.NestedBlockType == "single",
		}
	}

	var result []conflicts.Attr
	for _, attr := range attributes {
		if len(attr.ConflictsWith) > 0 {
			result = append(result, conflicts.Attr{
				TfsdkTag:      attr.TfsdkTag,
				GoName:        attr.GoName,
				IsPointer:     attr.IsBlock && attr.NestedBlockType == "single",
				ConflictsWith: attr.ConflictsWith,
			})
		}
	}
	return result, fieldLookup
}

// HasInt64RangeValidatorsAny returns true if any non-block int64 attribute has Minimum or Maximum set.
// Block attributes are excluded because blocks use nested schema types that don't support int64validator.
func HasInt64RangeValidatorsAny(attributes []openapi.TerraformAttribute) bool {
	for _, attr := range attributes {
		if !attr.IsBlock && (attr.HasMinimum || attr.HasMaximum) {
			return true
		}
		for _, nested := range attr.NestedAttributes {
			if !nested.IsBlock && (nested.HasMinimum || nested.HasMaximum) {
				return true
			}
		}
	}
	return false
}

// ScanPlanModifierUsage recursively scans attributes to determine which plan modifier imports are needed.
func ScanPlanModifierUsage(attributes []openapi.TerraformAttribute) (usesBool, usesInt64, usesString, usesList, usesMap bool) {
	for _, attr := range attributes {
		if attr.PlanModifier != "" && !attr.IsBlock {
			switch attr.Type {
			case "bool":
				usesBool = true
			case "int64":
				usesInt64 = true
			case "string":
				usesString = true
			case "list", "set":
				usesList = true
			case "map":
				usesMap = true
			}
		}
		if len(attr.NestedAttributes) > 0 {
			nb, ni, ns, nl, nm := ScanPlanModifierUsage(attr.NestedAttributes)
			usesBool = usesBool || nb
			usesInt64 = usesInt64 || ni
			usesString = usesString || ns
			usesList = usesList || nl
			usesMap = usesMap || nm
		}
	}
	return
}
