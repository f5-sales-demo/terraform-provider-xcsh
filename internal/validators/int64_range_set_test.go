// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package validators

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestInt64RangeSetValidator(t *testing.T) {
	v := Int64RangeSetValidator(
		Int64Range{Minimum: 0, Maximum: 0},
		Int64Range{Minimum: 512, Maximum: 16384},
	)
	for value, wantError := range map[int64]bool{
		-1: true, 0: false, 1: true, 511: true,
		512: false, 1500: false, 16384: false, 16385: true,
	} {
		t.Run(types.Int64Value(value).String(), func(t *testing.T) {
			resp := &validator.Int64Response{}
			v.ValidateInt64(context.Background(), validator.Int64Request{
				ConfigValue: types.Int64Value(value),
			}, resp)
			if got := resp.Diagnostics.HasError(); got != wantError {
				t.Fatalf("value %d: HasError=%v, want %v", value, got, wantError)
			}
		})
	}
}

func TestInt64RangeSetValidatorSkipsNullAndUnknown(t *testing.T) {
	v := Int64RangeSetValidator(Int64Range{Minimum: 0, Maximum: 0})
	for _, value := range []types.Int64{types.Int64Null(), types.Int64Unknown()} {
		resp := &validator.Int64Response{}
		v.ValidateInt64(context.Background(), validator.Int64Request{ConfigValue: value}, resp)
		if resp.Diagnostics.HasError() {
			t.Fatalf("%s produced diagnostics: %v", value, resp.Diagnostics)
		}
	}
}
