package client

import (
	"strings"
	"testing"
)

func TestValidateResponseFields(t *testing.T) {
	fields := []ResponseField{{Name: "image", Required: true, Type: "string", MinLength: 1, Format: "uri"}}
	for _, value := range []any{nil, "", "   ", 42, false, []any{}, "relative/path", "https://", "https://invalid host/image", "https://user:secret@example.invalid/image"} {
		err := ValidateResponseFields(map[string]any{"image": value}, fields)
		if err == nil {
			t.Errorf("expected invalid response for %T", value)
		} else if strings.Contains(err.Error(), "secret") || strings.Contains(err.Error(), "example.invalid") {
			t.Fatal("diagnostic leaked response content")
		}
	}
	if ValidateResponseFields(map[string]any{}, fields) == nil {
		t.Fatal("missing required payload was accepted")
	}
	if err := ValidateResponseFields(map[string]any{"image": "https://example.invalid/image?signature=private"}, fields); err != nil {
		t.Fatal(err)
	}
	if err := ValidateResponseFields(map[string]any{}, []ResponseField{{Name: "optional", Type: "string"}}); err != nil {
		t.Fatal(err)
	}
	if err := ValidateResponseFields(map[string]any{"template": "#cloud-config\nToken: placeholder\n"}, []ResponseField{{Name: "template", Required: true, Type: "string", MinLength: 1}}); err != nil {
		t.Fatal(err)
	}
}

func TestValidateResponseMapFields(t *testing.T) {
	fields := []ResponseField{{Name: "images", Type: "object", Required: true, MapFields: []ResponseField{{Name: "checksum", Type: "string", Required: true, Pattern: "^[a-f0-9]{32}$"}, {Name: "error", Type: "string", Required: true}}}}
	for _, item := range []any{nil, "private-value", map[string]any{}, map[string]any{"checksum": "private-value", "error": ""}, map[string]any{"checksum": strings.Repeat("a", 32)}} {
		err := ValidateResponseFields(map[string]any{"images": map[string]any{"private-uid": item}}, fields)
		if err == nil {
			t.Fatal("invalid map entry accepted")
		}
		if strings.Contains(err.Error(), "private-") {
			t.Fatal("response map diagnostic exposed private data")
		}
	}
	if err := ValidateResponseFields(map[string]any{"images": map[string]any{"private-uid": map[string]any{"checksum": strings.Repeat("a", 32), "error": ""}}}, fields); err != nil {
		t.Fatal(err)
	}
}
