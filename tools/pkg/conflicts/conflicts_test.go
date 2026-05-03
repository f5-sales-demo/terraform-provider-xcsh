// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package conflicts

import (
	"strings"
	"testing"
)

func TestTfsdkToGoName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"single word", "name", "Name"},
		{"snake_case", "field_name", "FieldName"},
		{"multiple underscores", "my_long_field_name", "MyLongFieldName"},
		{"already capitalized parts", "my_URL_field", "MyURLField"},
		{"empty string", "", ""},
		{"single underscore", "_", ""},
		{"trailing underscore", "field_", "Field"},
		{"leading underscore", "_field", "Field"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TfsdkToGoName(tt.input)
			if result != tt.expected {
				t.Errorf("TfsdkToGoName(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGenerateChecks_EmptyAttrs(t *testing.T) {
	result := GenerateChecks(nil, nil)
	if result != "" {
		t.Errorf("GenerateChecks(nil, nil) = %q, want empty string", result)
	}

	result = GenerateChecks([]Attr{}, map[string]string{})
	if result != "" {
		t.Errorf("GenerateChecks([], {}) = %q, want empty string", result)
	}
}

func TestGenerateChecks_SingleConflictPair(t *testing.T) {
	attrs := []Attr{
		{TfsdkTag: "field_a", GoName: "FieldA", ConflictsWith: []string{"field_b"}},
	}
	lookup := map[string]string{"field_a": "FieldA", "field_b": "FieldB"}
	result := GenerateChecks(attrs, lookup)

	// Check for expected code components
	if !strings.Contains(result, "data.FieldA.IsNull()") {
		t.Error("Expected data.FieldA.IsNull() in output")
	}
	if !strings.Contains(result, "data.FieldB.IsNull()") {
		t.Error("Expected data.FieldB.IsNull() in output")
	}
	if !strings.Contains(result, `path.Root("field_a")`) {
		t.Error(`Expected path.Root("field_a") in output`)
	}
	if !strings.Contains(result, "Conflicting Configuration") {
		t.Error("Expected 'Conflicting Configuration' error title in output")
	}
	if !strings.Contains(result, "field_a and field_b are mutually exclusive") {
		t.Error("Expected mutual exclusivity message in output")
	}
}

func TestGenerateChecks_Deduplication(t *testing.T) {
	// Both A->B and B->A should result in only one check
	attrs := []Attr{
		{TfsdkTag: "field_a", GoName: "FieldA", ConflictsWith: []string{"field_b"}},
		{TfsdkTag: "field_b", GoName: "FieldB", ConflictsWith: []string{"field_a"}},
	}
	lookup := map[string]string{"field_a": "FieldA", "field_b": "FieldB"}
	result := GenerateChecks(attrs, lookup)

	// Count occurrences of the if statement pattern
	count := strings.Count(result, "if !data.")
	if count != 1 {
		t.Errorf("Expected exactly 1 conflict check due to deduplication, got %d", count)
	}
}

func TestGenerateChecks_MultipleConflictPairs(t *testing.T) {
	attrs := []Attr{
		{TfsdkTag: "field_a", GoName: "FieldA", ConflictsWith: []string{"field_b", "field_c"}},
		{TfsdkTag: "field_d", GoName: "FieldD", ConflictsWith: []string{"field_e"}},
	}
	lookup := map[string]string{
		"field_a": "FieldA", "field_b": "FieldB", "field_c": "FieldC",
		"field_d": "FieldD", "field_e": "FieldE",
	}
	result := GenerateChecks(attrs, lookup)

	// Should have 3 distinct checks: A-B, A-C, D-E
	count := strings.Count(result, "if !data.")
	if count != 3 {
		t.Errorf("Expected 3 conflict checks, got %d", count)
	}

	// Verify each pair is checked
	if !strings.Contains(result, "field_a and field_b are mutually exclusive") {
		t.Error("Expected field_a/field_b conflict check")
	}
	if !strings.Contains(result, "field_a and field_c are mutually exclusive") {
		t.Error("Expected field_a/field_c conflict check")
	}
	if !strings.Contains(result, "field_d and field_e are mutually exclusive") {
		t.Error("Expected field_d/field_e conflict check")
	}
}

func TestGenerateChecks_GoNameUsedInNullCheck(t *testing.T) {
	attrs := []Attr{
		{TfsdkTag: "my_field", GoName: "MyField", ConflictsWith: []string{"other_field"}},
	}
	lookup := map[string]string{"my_field": "MyField", "other_field": "OtherField"}
	result := GenerateChecks(attrs, lookup)

	// The GoName should be used in the data.X.IsNull() check
	if !strings.Contains(result, "data.MyField.IsNull()") {
		t.Error("Expected GoName (MyField) to be used in null check")
	}
	// The TfsdkTag should be used in path.Root
	if !strings.Contains(result, `path.Root("my_field")`) {
		t.Error("Expected TfsdkTag (my_field) to be used in path.Root")
	}
}

func TestGenerateChecks_NoConflicts(t *testing.T) {
	attrs := []Attr{
		{TfsdkTag: "field_a", GoName: "FieldA", ConflictsWith: nil},
		{TfsdkTag: "field_b", GoName: "FieldB", ConflictsWith: []string{}},
	}
	lookup := map[string]string{"field_a": "FieldA", "field_b": "FieldB"}
	result := GenerateChecks(attrs, lookup)

	if result != "" {
		t.Errorf("Expected empty string for attrs with no conflicts, got %q", result)
	}
}

func TestGenerateChecks_ConflictTargetNotInLookup(t *testing.T) {
	// If a conflict target is not in the lookup map (e.g., it's a block), it should be skipped
	attrs := []Attr{
		{TfsdkTag: "field_a", GoName: "FieldA", ConflictsWith: []string{"block_field"}},
	}
	lookup := map[string]string{"field_a": "FieldA"} // block_field not in lookup
	result := GenerateChecks(attrs, lookup)

	if result != "" {
		t.Errorf("Expected empty string when conflict target not in lookup, got %q", result)
	}
}

func TestGenerateChecks_CodeIndentation(t *testing.T) {
	attrs := []Attr{
		{TfsdkTag: "field_a", GoName: "FieldA", ConflictsWith: []string{"field_b"}},
	}
	lookup := map[string]string{"field_a": "FieldA", "field_b": "FieldB"}
	result := GenerateChecks(attrs, lookup)

	// Check for proper indentation with tabs
	if !strings.Contains(result, "\tif !data.") {
		t.Error("Expected tab-indented if statement")
	}
	if !strings.Contains(result, "\t\tresp.Diagnostics.AddAttributeError(") {
		t.Error("Expected double-tab-indented AddAttributeError")
	}
	if !strings.Contains(result, "\t\t\tpath.Root(") {
		t.Error("Expected triple-tab-indented path.Root")
	}
}
