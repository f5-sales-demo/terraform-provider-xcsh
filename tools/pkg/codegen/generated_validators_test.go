// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import "testing"

func TestParseGeneratedInt64Bounds(t *testing.T) {
	tests := []struct {
		name                   string
		body                   string
		minimum, maximum       int
		hasMinimum, hasMaximum bool
	}{
		{
			name:       "between",
			body:       "Validators: []validator.Int64{\n int64validator.Between(6, 168),\n}",
			minimum:    6,
			maximum:    168,
			hasMinimum: true,
			hasMaximum: true,
		},
		{
			name:       "at least zero",
			body:       "int64validator.AtLeast(0)",
			minimum:    0,
			hasMinimum: true,
		},
		{
			name:       "at most negative",
			body:       "int64validator.AtMost( -2 )",
			maximum:    -2,
			hasMaximum: true,
		},
		{
			name: "no range validator",
			body: "int64validator.NoneOf(1)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minimum, maximum, hasMinimum, hasMaximum := ParseGeneratedInt64Bounds(tt.body)
			if minimum != tt.minimum || maximum != tt.maximum || hasMinimum != tt.hasMinimum || hasMaximum != tt.hasMaximum {
				t.Fatalf("got (%d, %d, %t, %t), want (%d, %d, %t, %t)",
					minimum, maximum, hasMinimum, hasMaximum,
					tt.minimum, tt.maximum, tt.hasMinimum, tt.hasMaximum)
			}
		})
	}
}
