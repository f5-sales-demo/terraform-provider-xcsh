// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package schema

import (
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestConvertToTerraformAttribute_StrictMaps(t *testing.T) {
	spec := &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				"StringRef": {
					Type: "string",
				},
				"IntRef": {
					Type: "integer",
				},
			},
		},
	}

	// Helper to assert the complete fallback object-block shape
	assertFallbackObjectShape := func(t *testing.T, attr openapi.TerraformAttribute) {
		t.Helper()
		if attr.Type != "object" {
			t.Errorf("expected Type='object', got %q", attr.Type)
		}
		if !attr.IsBlock {
			t.Error("expected IsBlock=true")
		}
		if attr.NestedBlockType != "single" {
			t.Errorf("expected NestedBlockType='single', got %q", attr.NestedBlockType)
		}
		if attr.GoType != "map[string]interface{}" {
			t.Errorf("expected GoType='map[string]interface{}', got %q", attr.GoType)
		}
	}

	// 1. Strict String-Valued AdditionalProperties Map
	t.Run("String AdditionalProperties", func(t *testing.T) {
		schema := openapi.Schema{
			Type: "object",
			AdditionalProperties: map[string]interface{}{
				"type": "string",
			},
		}
		attr := ConvertToTerraformAttribute("test_map", schema, false, "", spec)
		if attr.Type != "map" || attr.ElementType != "string" || attr.GoType != "map[string]string" {
			t.Errorf("expected strict string map, got Type=%q ElementType=%q GoType=%q", attr.Type, attr.ElementType, attr.GoType)
		}
		if attr.ConversionError != "" {
			t.Errorf("unexpected conversion error: %s", attr.ConversionError)
		}
	})

	// 2. Non-String Map Element (Should not panic, registers ConversionError)
	t.Run("Non-String Map Element", func(t *testing.T) {
		schema := openapi.Schema{
			Type: "object",
			AdditionalProperties: map[string]interface{}{
				"type": "integer",
			},
		}
		attr := ConvertToTerraformAttribute("test_map", schema, false, "", spec)
		if attr.ConversionError == "" {
			t.Error("expected controlled ConversionError, got none")
		}
		if attr.Type != "map" {
			t.Errorf("expected Type='map', got %q", attr.Type)
		}
	})

	// 3. Boolean additionalProperties Form (True / False - Should Fallback to Object unless overridden/ruled)
	t.Run("Boolean AdditionalProperties True Fallback", func(t *testing.T) {
		schema := openapi.Schema{
			Type:                 "object",
			AdditionalProperties: true,
		}
		attr := ConvertToTerraformAttribute("test_unruled_map_not_whitelisted", schema, false, "", spec)
		assertFallbackObjectShape(t, attr)
	})

	t.Run("Boolean AdditionalProperties True Whitelisted Fallback", func(t *testing.T) {
		schema := openapi.Schema{
			Type:                 "object",
			AdditionalProperties: true,
		}
		attr := ConvertToTerraformAttribute("labels", schema, false, "", spec)
		assertFallbackObjectShape(t, attr)
	})

	t.Run("Boolean AdditionalProperties True with Map Rules", func(t *testing.T) {
		schema := openapi.Schema{
			Type:                 "object",
			AdditionalProperties: true,
			XVesValidationRules: map[string]string{
				"ves.io.schema.rules.map.values.string.min_bytes": "1",
			},
		}
		attr := ConvertToTerraformAttribute("test_map", schema, false, "", spec)
		if attr.Type != "map" || attr.ElementType != "string" || attr.GoType != "map[string]string" {
			t.Errorf("expected string map for schema with map rules, got Type=%q ElementType=%q GoType=%q", attr.Type, attr.ElementType, attr.GoType)
		}
	})

	t.Run("Boolean AdditionalProperties False", func(t *testing.T) {
		schema := openapi.Schema{
			Type:                 "object",
			AdditionalProperties: false,
		}
		attr := ConvertToTerraformAttribute("test_map", schema, false, "", spec)
		assertFallbackObjectShape(t, attr)
	})

	// 4. Ref pointing to string type
	t.Run("String Ref", func(t *testing.T) {
		schema := openapi.Schema{
			Type: "object",
			AdditionalProperties: map[string]interface{}{
				"$ref": "#/components/schemas/StringRef",
			},
		}
		attr := ConvertToTerraformAttribute("test_map", schema, false, "", spec)
		if attr.Type != "map" || attr.ElementType != "string" {
			t.Errorf("expected string ref to map strictly, got Type=%q", attr.Type)
		}
	})

	// 5. Ref pointing to non-string type (Should register ConversionError, no panic)
	t.Run("Int Ref Error", func(t *testing.T) {
		schema := openapi.Schema{
			Type: "object",
			AdditionalProperties: map[string]interface{}{
				"$ref": "#/components/schemas/IntRef",
			},
		}
		attr := ConvertToTerraformAttribute("test_map", schema, false, "", spec)
		if attr.ConversionError == "" {
			t.Error("expected controlled ConversionError for non-string ref, got none")
		}
		if attr.Type != "map" {
			t.Errorf("expected Type='map', got %q", attr.Type)
		}
	})

	// 6. Map Rules without Element Schema (Should Fallback to Object)
	t.Run("Map Rules Without Element Schema", func(t *testing.T) {
		schema := openapi.Schema{
			Type: "object",
			XValidationRules: map[string]string{
				"ves.io.schema.rules.map.keys.string.min_len": "1",
			},
		}
		attr := ConvertToTerraformAttribute("test_unruled_map_not_whitelisted", schema, false, "", spec)
		assertFallbackObjectShape(t, attr)
	})

	// 7. Genuine empty-marker object {} (Should Fallback to Object unless overridden/ruled)
	t.Run("Empty-Marker Object Fallback", func(t *testing.T) {
		schema := openapi.Schema{
			Type:                 "object",
			AdditionalProperties: map[string]interface{}{},
		}
		attr := ConvertToTerraformAttribute("test_unruled_map_not_whitelisted", schema, false, "", spec)
		assertFallbackObjectShape(t, attr)
	})

	t.Run("Empty-Marker Object Whitelisted Fallback", func(t *testing.T) {
		schema := openapi.Schema{
			Type:                 "object",
			AdditionalProperties: map[string]interface{}{},
		}
		attr := ConvertToTerraformAttribute("fixed_ip_map", schema, false, "", spec)
		assertFallbackObjectShape(t, attr)
	})

	t.Run("Empty-Marker Object with Map Rules", func(t *testing.T) {
		schema := openapi.Schema{
			Type:                 "object",
			AdditionalProperties: map[string]interface{}{},
			XValidationRules: map[string]string{
				"ves.io.schema.rules.map.values.string.min_bytes": "1",
			},
		}
		attr := ConvertToTerraformAttribute("test_map", schema, false, "", spec)
		if attr.Type != "map" || attr.ElementType != "string" || attr.GoType != "map[string]string" {
			t.Errorf("expected string map for empty-marker with map rules, got Type=%q", attr.Type)
		}
	})

	// 8. Unresolved Ref (Should register ConversionError, no panic)
	t.Run("Unresolved Ref Error", func(t *testing.T) {
		schema := openapi.Schema{
			Type: "object",
			AdditionalProperties: map[string]interface{}{
				"$ref": "#/components/schemas/UnresolvedRef",
			},
		}
		attr := ConvertToTerraformAttribute("test_map", schema, false, "", spec)
		if attr.ConversionError == "" {
			t.Error("expected controlled ConversionError for unresolved ref, got none")
		}
	})

	// 9. Malformed additionalProperties (e.g., unexpected string - Should Fallback to Object)
	t.Run("Malformed AdditionalProperties String", func(t *testing.T) {
		schema := openapi.Schema{
			Type:                 "object",
			AdditionalProperties: "some-malformed-string-here",
		}
		attr := ConvertToTerraformAttribute("test_map", schema, false, "", spec)
		assertFallbackObjectShape(t, attr)
	})

	// 10. Nil AdditionalProperties (Fallback to Object unless overridden/ruled)
	t.Run("Nil AdditionalProperties Fallback", func(t *testing.T) {
		schema := openapi.Schema{
			Type:                 "object",
			AdditionalProperties: nil,
		}
		attr := ConvertToTerraformAttribute("test_unruled_map_not_whitelisted", schema, false, "", spec)
		assertFallbackObjectShape(t, attr)
	})

	t.Run("Nil AdditionalProperties Whitelisted Fallback", func(t *testing.T) {
		schema := openapi.Schema{
			Type:                 "object",
			AdditionalProperties: nil,
		}
		attr := ConvertToTerraformAttribute("fixed_ip_map", schema, false, "", spec)
		assertFallbackObjectShape(t, attr)
	})

	t.Run("Nil AdditionalProperties with Map Rules", func(t *testing.T) {
		schema := openapi.Schema{
			Type:                 "object",
			AdditionalProperties: nil,
			XVesValidationRules: map[string]string{
				"ves.io.schema.rules.map.values.string.min_bytes": "1",
			},
		}
		attr := ConvertToTerraformAttribute("test_map", schema, false, "", spec)
		if attr.Type != "map" || attr.ElementType != "string" || attr.GoType != "map[string]string" {
			t.Errorf("expected string map for nil AdditionalProperties with map rules, got Type=%q", attr.Type)
		}
	})
}
