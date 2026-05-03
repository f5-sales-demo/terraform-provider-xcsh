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
