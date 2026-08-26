package codegen

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"strings"
	"testing"
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
