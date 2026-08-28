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
		`} else if isImport || data.RequiredValues.IsUnknown() {`,
	} {
		if !strings.Contains(unmarshal, want) {
			t.Errorf("generated list unmarshal does not preserve configured empty values; missing %q:\n%s", want, unmarshal)
		}
	}
	if strings.Contains(unmarshal, "ok && len(v) > 0") {
		t.Errorf("generated list unmarshal incorrectly treats an API-provided empty list as absent:\n%s", unmarshal)
	}
}

func TestOptionalTopLevelListPreservesPriorNullOrConfiguredValueWhenAPIEchoesEmpty(t *testing.T) {
	attrs := []openapi.TerraformAttribute{{
		Name: "compliances", GoName: "Compliances", TfsdkTag: "compliances",
		Type: "list", ElementType: "string", JsonName: "compliances",
		IsSpecField: true, Optional: true,
	}}

	unmarshal, err := RenderSpecUnmarshalCode(attrs, "\t", "DataType")
	if err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	for _, want := range []string{
		`ok && (len(v) > 0 || isImport || data.Compliances.IsUnknown())`,
		`} else if isImport || data.Compliances.IsUnknown() {`,
	} {
		if !strings.Contains(unmarshal, want) {
			t.Errorf("optional list read must preserve prior state when the API echoes empty; missing %q:\n%s", want, unmarshal)
		}
	}
}

func TestOptionalComputedTopLevelListResolvesUnknownPlanWhenAPIOmitsIt(t *testing.T) {
	attrs := []openapi.TerraformAttribute{{
		Name: "swagger_specs", GoName: "SwaggerSpecs", TfsdkTag: "swagger_specs",
		Type: "list", ElementType: "string", JsonName: "swagger_specs",
		IsSpecField: true, Optional: true, Computed: true,
	}}

	unmarshal, err := RenderSpecUnmarshalCode(attrs, "\t", "APIDefinition")
	if err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	for _, want := range []string{
		`len(v) > 0 || isImport || data.SwaggerSpecs.IsUnknown()`,
		`} else if isImport || data.SwaggerSpecs.IsUnknown() {`,
	} {
		if !strings.Contains(unmarshal, want) {
			t.Errorf("optional+computed list must resolve an unknown plan; missing %q:\n%s", want, unmarshal)
		}
	}
}

func TestTopLevelBlockWithComputedDescendantResolvesOmittedAPIValue(t *testing.T) {
	attrs := []openapi.TerraformAttribute{{
		Name: "origins", GoName: "Origins", TfsdkTag: "origins", JsonName: "origins",
		IsSpecField: true, IsBlock: true, NestedBlockType: "list",
		NestedAttributes: []openapi.TerraformAttribute{
			{GoName: "Name", TfsdkTag: "name", JsonName: "name", Type: "string", Optional: true},
			{GoName: "Tenant", TfsdkTag: "tenant", JsonName: "tenant", Type: "string", Computed: true},
		},
	}}

	unmarshal, err := RenderSpecUnmarshalCode(attrs, "\t", "APIDefinition")
	if err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if !strings.Contains(unmarshal, "} else {\n\t\tdata.Origins = types.ListNull") {
		t.Fatalf("computed-descendant block omission must resolve to known null:\n%s", unmarshal)
	}
}

func TestNestedMapRoundTripUsesImportAwareNullPreservation(t *testing.T) {
	attr := openapi.TerraformAttribute{
		Name:        "Labels",
		GoName:      "Labels",
		TfsdkTag:    "labels",
		JsonName:    "labels",
		Type:        "map",
		ElementType: "string",
		Optional:    true,
	}

	var rendered strings.Builder
	if err := renderUnmarshalChild(&rendered, "Example", "Item", attr, "itemData", "data.Item", "data.Item != nil", "single", "\t"); err != nil {
		t.Fatalf("render nested map: %v", err)
	}
	got := rendered.String()
	if !strings.Contains(got, `UnmarshalStringMapForRead(ctx, itemData["labels"]`) ||
		!strings.Contains(got, `"Labels", isImport, &resp.Diagnostics)`) {
		t.Fatalf("generated nested map does not preserve null/empty semantics by read mode:\n%s", got)
	}
}
