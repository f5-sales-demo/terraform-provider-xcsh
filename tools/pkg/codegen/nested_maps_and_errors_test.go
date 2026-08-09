// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestRenderMarshalScalar_UnsupportedTypeError(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{
			Name:        "unsupported_field",
			GoName:      "UnsupportedField",
			TfsdkTag:    "unsupported_field",
			Type:        "complex",
			IsSpecField: true,
		},
	}

	_, err := RenderSpecMarshalCode(attrs, "\t", "Test")
	if err == nil {
		t.Error("expected error for unsupported field type, got nil")
	} else if !strings.Contains(err.Error(), "unsupported type \"complex\"") {
		t.Errorf("expected error to contain unsupported type message, got: %v", err)
	}
}

func TestRenderMarshalScalar_UnsupportedListElementError(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{
			Name:        "unsupported_list",
			GoName:      "UnsupportedList",
			TfsdkTag:    "unsupported_list",
			Type:        "list",
			ElementType: "unsupported_elem",
			IsSpecField: true,
		},
	}

	_, err := RenderSpecMarshalCode(attrs, "\t", "Test")
	if err == nil {
		t.Error("expected error for unsupported list element type, got nil")
	} else if !strings.Contains(err.Error(), "unsupported list element type \"unsupported_elem\"") {
		t.Errorf("expected error to contain unsupported list element message, got: %v", err)
	}
}

func TestRenderUnmarshalTopLevelScalar_UnsupportedTypeError(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{
			Name:        "unsupported_field",
			GoName:      "UnsupportedField",
			TfsdkTag:    "unsupported_field",
			Type:        "complex",
			IsSpecField: true,
		},
	}

	_, err := RenderSpecUnmarshalCode(attrs, "\t", "Test")
	if err == nil {
		t.Error("expected error for unsupported field type during unmarshal, got nil")
	} else if !strings.Contains(err.Error(), "unsupported type \"complex\"") {
		t.Errorf("expected error to contain unsupported type message, got: %v", err)
	}
}

func TestRenderMapSupport_Success(t *testing.T) {
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

	marshal, err := RenderSpecMarshalCode(attrs, "\t", "Test")
	if err != nil {
		t.Fatalf("unexpected marshal error: %v", err)
	}
	if !strings.Contains(marshal, "MetadataLabelsMap") || !strings.Contains(marshal, "ElementsAs(ctx, &MetadataLabelsMap, false)") {
		t.Errorf("expected marshal code to contain ElementsAs for Map, got:\n%s", marshal)
	}

	unmarshal, err := RenderSpecUnmarshalCode(attrs, "\t", "Test")
	if err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if !strings.Contains(unmarshal, "UnmarshalStringMap") {
		t.Errorf("expected unmarshal code to invoke UnmarshalStringMap, got:\n%s", unmarshal)
	}
}

func TestNestedPathAwareErrors(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{
			Name:            "outer",
			GoName:          "Outer",
			TfsdkTag:        "outer",
			IsBlock:         true,
			NestedBlockType: "single",
			IsSpecField:     true,
			NestedAttributes: []openapi.TerraformAttribute{
				{
					Name:            "inner",
					GoName:          "Inner",
					TfsdkTag:        "inner",
					IsBlock:         true,
					NestedBlockType: "single",
					NestedAttributes: []openapi.TerraformAttribute{
						{
							Name:     "bad_field",
							GoName:   "BadField",
							TfsdkTag: "bad_field",
							Type:     "complex",
						},
					},
				},
			},
		},
	}

	_, err := RenderSpecMarshalCodeWithVar(attrs, "\t", "data", "Test")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	expectedMsg := `marshaling Test: field "outer": field "inner": field "bad_field": unsupported type "complex" at field bad_field`
	if !strings.Contains(err.Error(), expectedMsg) {
		t.Errorf("expected error to contain %q, got: %v", expectedMsg, err)
	}

	_, err = RenderSpecUnmarshalCode(attrs, "\t", "Test")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	expectedUnmarshal := `unmarshaling Test: field "outer": field "inner": field "bad_field": unsupported type "complex" at field bad_field`
	if !strings.Contains(err.Error(), expectedUnmarshal) {
		t.Errorf("expected error to contain %q, got: %v", expectedUnmarshal, err)
	}
}

func TestNestedMapSupport_Success(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{
			Name:            "outer",
			GoName:          "Outer",
			TfsdkTag:        "outer",
			IsBlock:         true,
			NestedBlockType: "single",
			IsSpecField:     true,
			NestedAttributes: []openapi.TerraformAttribute{
				{
					Name:        "nested_map",
					GoName:      "NestedMap",
					TfsdkTag:    "nested_map",
					Type:        "map",
					ElementType: "string",
					JsonName:    "nested_map",
				},
			},
		},
	}

	marshal, err := RenderSpecMarshalCodeWithVar(attrs, "\t", "data", "Test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(marshal, "ElementsAs(ctx, &NestedMapMap, false)") {
		t.Errorf("expected nested marshal to contain ElementsAs for map, got:\n%s", marshal)
	}
	if !strings.Contains(marshal, "resp.Diagnostics.Append(diags...)") {
		t.Errorf("expected map diagnostics to be surfaced, got:\n%s", marshal)
	}

	unmarshal, err := RenderSpecUnmarshalCode(attrs, "\t", "Test")
	if err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if !strings.Contains(unmarshal, "UnmarshalStringMap") {
		t.Errorf("expected nested unmarshal to invoke UnmarshalStringMap, got:\n%s", unmarshal)
	}
}
