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
			t.Errorf("expected strict string map, got Type=%q ElementType=%q GoName=%q", attr.Type, attr.ElementType, attr.GoName)
		}
	})

	// 2. Non-String Map Element (Should Panic)
	t.Run("Non-String Map Element", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Error("expected panic for non-string map element type")
			}
		}()
		schema := openapi.Schema{
			Type: "object",
			AdditionalProperties: map[string]interface{}{
				"type": "integer",
			},
		}
		ConvertToTerraformAttribute("test_map", schema, false, "", spec)
	})

	// 3. Boolean additionalProperties Form (True / False - Should Not Be Map)
	t.Run("Boolean AdditionalProperties True", func(t *testing.T) {
		schema := openapi.Schema{
			Type:                 "object",
			AdditionalProperties: true,
		}
		attr := ConvertToTerraformAttribute("test_map", schema, false, "", spec)
		if attr.Type == "map" {
			t.Error("expected additionalProperties: true NOT to be classified as a map")
		}
	})

	t.Run("Boolean AdditionalProperties False", func(t *testing.T) {
		schema := openapi.Schema{
			Type:                 "object",
			AdditionalProperties: false,
		}
		attr := ConvertToTerraformAttribute("test_map", schema, false, "", spec)
		if attr.Type == "map" {
			t.Error("expected additionalProperties: false NOT to be classified as a map")
		}
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

	// 5. Ref pointing to non-string type (Should Panic)
	t.Run("Int Ref Panic", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Error("expected panic for ref pointing to non-string type")
			}
		}()
		schema := openapi.Schema{
			Type: "object",
			AdditionalProperties: map[string]interface{}{
				"$ref": "#/components/schemas/IntRef",
			},
		}
		ConvertToTerraformAttribute("test_map", schema, false, "", spec)
	})

	// 6. Map Rules without Element Schema (Should Not Be Map)
	t.Run("Map Rules Without Element Schema", func(t *testing.T) {
		schema := openapi.Schema{
			Type: "object",
			XValidationRules: map[string]string{
				"ves.io.schema.rules.map.keys.string.min_len": "1",
			},
		}
		attr := ConvertToTerraformAttribute("test_map", schema, false, "", spec)
		if attr.Type == "map" {
			t.Error("expected map rules without additionalProperties NOT to be classified as a map")
		}
	})

	// 7. Genuine empty-marker object {} (Should Not Be Map)
	t.Run("Empty-Marker Object", func(t *testing.T) {
		schema := openapi.Schema{
			Type:                 "object",
			AdditionalProperties: map[string]interface{}{},
		}
		attr := ConvertToTerraformAttribute("test_map", schema, false, "", spec)
		if attr.Type == "map" {
			t.Error("expected empty-marker object (additionalProperties: {}) NOT to be classified as a map")
		}
	})
}
