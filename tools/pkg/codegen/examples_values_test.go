// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestExampleValueHonorsInt64Bounds(t *testing.T) {
	tests := []struct {
		name string
		attr openapi.TerraformAttribute
		want string
	}{
		{
			name: "positive minimum",
			attr: openapi.TerraformAttribute{Type: "int64", Minimum: 6, HasMinimum: true, Maximum: 168, HasMaximum: true},
			want: "6",
		},
		{
			name: "zero maximum",
			attr: openapi.TerraformAttribute{Type: "int64", Maximum: 0, HasMaximum: true},
			want: "0",
		},
		{
			name: "negative range",
			attr: openapi.TerraformAttribute{Type: "int64", Minimum: -10, HasMinimum: true, Maximum: -2, HasMaximum: true},
			want: "-10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exampleValue(tt.attr); got != tt.want {
				t.Fatalf("exampleValue() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestExampleValueHonorsDomainValidators(t *testing.T) {
	tests := []openapi.TerraformAttribute{
		{Type: "string", ETLDPlusOne: true},
		{Type: "string", UseDomainValidator: true},
	}
	for _, attr := range tests {
		if got := exampleValue(attr); got != `"example.com"` {
			t.Fatalf("exampleValue() = %s, want %s", got, `"example.com"`)
		}
	}
}
