// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package constraints

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected *Parsed
	}{
		{
			name:     "nil input returns nil",
			input:    nil,
			expected: nil,
		},
		{
			name:     "empty map returns nil (no metadata)",
			input:    map[string]interface{}{},
			expected: nil,
		},
		{
			name: "low confidence (0.5) and not deterministic returns nil",
			input: map[string]interface{}{
				"metadata": map[string]interface{}{
					"deterministic": false,
					"confidence":    0.5,
				},
				"minLength": float64(5),
			},
			expected: nil,
		},
		{
			name: "deterministic=true overrides low confidence",
			input: map[string]interface{}{
				"metadata": map[string]interface{}{
					"deterministic": true,
					"confidence":    0.3,
				},
				"minLength": float64(10),
				"maxLength": float64(100),
			},
			expected: &Parsed{
				MinLength: 10,
				MaxLength: 100,
			},
		},
		{
			name: "high confidence (0.95) returns Parsed with correct values",
			input: map[string]interface{}{
				"metadata": map[string]interface{}{
					"deterministic": false,
					"confidence":    0.95,
				},
				"minLength": float64(1),
				"maxLength": float64(255),
				"pattern":   "^[a-z]+$",
				"minItems":  float64(1),
				"maxItems":  float64(10),
			},
			expected: &Parsed{
				MinLength: 1,
				MaxLength: 255,
				Pattern:   "^[a-z]+$",
				MinItems:  1,
				MaxItems:  10,
			},
		},
		{
			name: "all fields extracted correctly",
			input: map[string]interface{}{
				"metadata": map[string]interface{}{
					"deterministic": true,
				},
				"minLength": float64(5),
				"maxLength": float64(50),
				"pattern":   "^[A-Za-z0-9_-]+$",
				"minItems":  float64(2),
				"maxItems":  float64(20),
			},
			expected: &Parsed{
				MinLength: 5,
				MaxLength: 50,
				Pattern:   "^[A-Za-z0-9_-]+$",
				MinItems:  2,
				MaxItems:  20,
			},
		},
		{
			name: "missing individual fields are zero-valued",
			input: map[string]interface{}{
				"metadata": map[string]interface{}{
					"deterministic": true,
				},
				"pattern": "^test$",
			},
			expected: &Parsed{
				MinLength: 0,
				MaxLength: 0,
				Pattern:   "^test$",
				MinItems:  0,
				MaxItems:  0,
			},
		},
		{
			name: "confidence exactly at threshold (0.9) returns Parsed",
			input: map[string]interface{}{
				"metadata": map[string]interface{}{
					"confidence": 0.9,
				},
				"minLength": float64(3),
			},
			expected: &Parsed{
				MinLength: 3,
			},
		},
		{
			name: "confidence just below threshold (0.89) returns nil",
			input: map[string]interface{}{
				"metadata": map[string]interface{}{
					"confidence": 0.89,
				},
				"minLength": float64(3),
			},
			expected: nil,
		},
		{
			name: "metadata missing deterministic field, high confidence",
			input: map[string]interface{}{
				"metadata": map[string]interface{}{
					"confidence": 0.92,
				},
				"maxItems": float64(5),
			},
			expected: &Parsed{
				MaxItems: 5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Parse(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Parse() = %+v, expected %+v", result, tt.expected)
			}
		})
	}
}

func TestParse_TopLevelDeterministic(t *testing.T) {
	input := map[string]interface{}{
		"constraintType": "string",
		"category":       "naming",
		"maxLength":      float64(1024),
		"minLength":      float64(1),
		"pattern":        "^[a-z]([-a-z0-9]*[a-z0-9])?$",
		"deterministic":  true,
		"metadata": map[string]interface{}{
			"source":     "api-probed",
			"confidence": 0.99,
		},
	}
	result := Parse(input)
	if result == nil {
		t.Fatal("Parse() returned nil for top-level deterministic=true constraint")
	}
	if result.MinLength != 1 {
		t.Errorf("MinLength = %d, want 1", result.MinLength)
	}
	if result.MaxLength != 1024 {
		t.Errorf("MaxLength = %d, want 1024", result.MaxLength)
	}
	if result.Pattern != "^[a-z]([-a-z0-9]*[a-z0-9])?$" {
		t.Errorf("Pattern = %q, want naming pattern", result.Pattern)
	}
}

func TestParse_NumericConstraints(t *testing.T) {
	input := map[string]interface{}{
		"constraintType": "number",
		"category":       "threshold",
		"minimum":        float64(1),
		"maximum":        float64(16),
		"deterministic":  true,
		"metadata": map[string]interface{}{
			"source":     "api-probed",
			"confidence": 0.99,
		},
	}
	result := Parse(input)
	if result == nil {
		t.Fatal("Parse() returned nil for numeric constraint")
	}
	if result.Minimum != 1 {
		t.Errorf("Minimum = %d, want 1", result.Minimum)
	}
	if result.Maximum != 16 {
		t.Errorf("Maximum = %d, want 16", result.Maximum)
	}
}

func TestParse_DiscoverySourcePatternSkipped(t *testing.T) {
	// discovery-inferred URL patterns (e.g., format:uri → ^(https?|ftp)://)
	// must be skipped because they reject F5 XC-specific schemes like string:///
	input := map[string]interface{}{
		"constraintType": "string",
		"category":       "discovery",
		"format":         "uri",
		"pattern":        `^(https?|ftp)://[^\s/$.?#].[^\s]*$`,
		"maxLength":      float64(131072),
		"deterministic":  true,
		"metadata": map[string]interface{}{
			"source":     "discovery",
			"confidence": float64(0.99),
		},
	}
	result := Parse(input)
	if result == nil {
		t.Fatal("Parse() returned nil, want non-nil Parsed with no pattern")
	}
	if result.Pattern != "" {
		t.Errorf("Pattern = %q, want empty string (discovery patterns should be skipped)", result.Pattern)
	}
	if result.MaxLength != 131072 {
		t.Errorf("MaxLength = %d, want 131072", result.MaxLength)
	}
}

func TestParse_NonDiscoveryPatternKept(t *testing.T) {
	// non-discovery patterns (naming, network, etc.) should still be applied
	input := map[string]interface{}{
		"constraintType": "string",
		"category":       "naming",
		"pattern":        `^[a-z]([-a-z0-9]*[a-z0-9])?$`,
		"deterministic":  true,
		"metadata": map[string]interface{}{
			"source":     "api-probed",
			"confidence": float64(0.99),
		},
	}
	result := Parse(input)
	if result == nil {
		t.Fatal("Parse() returned nil for naming pattern")
	}
	if result.Pattern != `^[a-z]([-a-z0-9]*[a-z0-9])?$` {
		t.Errorf("Pattern = %q, want naming regex", result.Pattern)
	}
}

func TestParse_ConfidenceInMetadata_NoDeterministic(t *testing.T) {
	input := map[string]interface{}{
		"constraintType": "string",
		"maxLength":      float64(1200),
		"metadata": map[string]interface{}{
			"source":     "inferred",
			"confidence": float64(0.9),
		},
	}
	result := Parse(input)
	if result == nil {
		t.Fatal("Parse() returned nil for confidence=0.9 constraint")
	}
	if result.MaxLength != 1200 {
		t.Errorf("MaxLength = %d, want 1200", result.MaxLength)
	}
}

func TestParse_MinimumZeroIsPresent(t *testing.T) {
	p := Parse(map[string]interface{}{
		"minimum": float64(0), "maximum": float64(255), "deterministic": true,
	})
	if p == nil {
		t.Fatal("nil")
	}
	if !p.HasMinimum || p.Minimum != 0 {
		t.Errorf("HasMinimum=%v Minimum=%d, want true/0", p.HasMinimum, p.Minimum)
	}
	if !p.HasMaximum || p.Maximum != 255 {
		t.Errorf("HasMaximum=%v Maximum=%d, want true/255", p.HasMaximum, p.Maximum)
	}
}

func TestParse_Format(t *testing.T) {
	p := Parse(map[string]interface{}{"format": "mac-address", "deterministic": true})
	if p == nil || p.Format != "mac-address" {
		t.Fatalf("Format = %q, want mac-address", func() string {
			if p == nil {
				return "<nil>"
			}
			return p.Format
		}())
	}
}
