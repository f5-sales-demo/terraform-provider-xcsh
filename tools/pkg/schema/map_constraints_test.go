package schema

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"strings"
	"testing"
)

func TestStructuredMapConstraintConversion(t *testing.T) {
	s := openapi.Schema{Type: "object", AdditionalProperties: map[string]interface{}{"type": "string"}, XF5XCConstraints: map[string]interface{}{
		"constraintType": "map", "deterministic": true, "values": map[string]interface{}{"type": "string", "maxLength": float64(65536)}, "cardinality": map[string]interface{}{"maxProperties": float64(16)},
	}}
	attribute := ConvertToTerraformAttribute("custom_errors", s, false, "", &openapi.Spec{})
	if attribute.Type != "map" || !strings.Contains(attribute.MapConstraintsJSON, "65536") {
		t.Fatalf("lost structured map: %+v", attribute)
	}
	if attribute.MaxLength != 0 {
		t.Fatal("value bound was applied to container")
	}
}
