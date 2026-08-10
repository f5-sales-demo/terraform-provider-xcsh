// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// UnmarshalStringMap converts an API map value to a types.Map, and implements state preservation if the API value is omitted (nil).
// It adds a diagnostic error if any of the map values are not strings.
func UnmarshalStringMap(ctx context.Context, apiVal interface{}, priorState types.Map, fieldName string, diags *diag.Diagnostics) types.Map {
	if apiVal != nil {
		if v, ok := apiVal.(map[string]interface{}); ok {
			items := make(map[string]string)
			for mk, mv := range v {
				if mvs, ok := mv.(string); ok {
					items[mk] = mvs
				} else {
					diags.AddError(
						"Unexpected type in map",
						fmt.Sprintf("Expected string for key %s in field %s, got %T", mk, fieldName, mv),
					)
				}
			}
			mapVal, d := types.MapValueFrom(ctx, types.StringType, items)
			diags.Append(d...)
			return mapVal
		} else {
			diags.AddError(
				"Invalid map container type",
				fmt.Sprintf("Expected map[string]interface{} for field %s, got %T", fieldName, apiVal),
			)
		}
	}
	if !priorState.IsNull() && !priorState.IsUnknown() {
		return priorState
	}
	return types.MapNull(types.StringType)
}
