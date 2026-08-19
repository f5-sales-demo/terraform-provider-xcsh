package codegen

import (
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestMapRoundTrip(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{
			Name:        "metadata_labels",
			GoName:      "MetadataLabels",
			TfsdkTag:    "metadata_labels",
			Type:        "map",
			ElementType: "string",
			JsonName:    "metadata_labels",
			IsSpecField: true,
		},
	}

	_, err := RenderSpecMarshalCode(attrs, "\t", "Test")
	if err != nil {
		t.Fatalf("unexpected marshal error: %v", err)
	}

	unmarshal, err := RenderSpecUnmarshalCode(attrs, "\t", "Test")
	if err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	if !strings.Contains(unmarshal, "UnmarshalStringMap") {
		t.Errorf("Expected UnmarshalStringMap helper invocation in unmarshal code, got:\n%s", unmarshal)
	}
}

func TestTopLevelListRoundTripPreservesConfiguredEmptyValueWhenAPIElidesIt(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{
			Name:        "required_values",
			GoName:      "RequiredValues",
			TfsdkTag:    "required_values",
			Type:        "list",
			ElementType: "string",
			JsonName:    "required_values",
			IsSpecField: true,
			Required:    true,
		},
	}

	unmarshal, err := RenderSpecUnmarshalCode(attrs, "\t", "Test")
	if err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	for _, want := range []string{
		`if v, ok := apiResource.Spec["required_values"].([]interface{}); ok {`,
		`required_valuesList := make([]string, 0, len(v))`,
		`} else if data.RequiredValues.IsNull() || data.RequiredValues.IsUnknown() {`,
	} {
		if !strings.Contains(unmarshal, want) {
			t.Errorf("generated list unmarshal does not preserve configured empty values; missing %q:\n%s", want, unmarshal)
		}
	}
	if strings.Contains(unmarshal, "ok && len(v) > 0") {
		t.Errorf("generated list unmarshal incorrectly treats an API-provided empty list as absent:\n%s", unmarshal)
	}
}
