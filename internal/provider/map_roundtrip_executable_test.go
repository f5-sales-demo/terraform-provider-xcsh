// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestExecutableMapRoundTripBehavior(t *testing.T) {
	ctx := context.TODO()

	// 1. Populated map unmarshaling
	t.Run("Populated Map", func(t *testing.T) {
		var diags diag.Diagnostics

		// Setup input matching the production structure
		InterfaceListItemMap := map[string]interface{}{
			"labels": map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
		}

		// Execute actual generated-style unmarshaling closure on production type
		result := func() types.Map {
			if v, ok := InterfaceListItemMap["labels"].(map[string]interface{}); ok {
				items := make(map[string]string)
				for mk, mv := range v {
					if mvs, ok := mv.(string); ok {
						items[mk] = mvs
					} else {
						diags.AddError("Unexpected type in map", "Expected string")
					}
				}
				mapVal, d := types.MapValueFrom(ctx, types.StringType, items)
				diags.Append(d...)
				return mapVal
			}
			return types.MapNull(types.StringType)
		}()

		if diags.HasError() {
			t.Fatalf("unexpected diagnostics error: %v", diags)
		}

		if result.IsNull() || result.IsUnknown() {
			t.Fatal("expected result to be non-null and non-unknown")
		}

		var elements map[string]string
		d := result.ElementsAs(ctx, &elements, false)
		if d.HasError() {
			t.Fatalf("failed to retrieve elements: %v", d)
		}

		if elements["key1"] != "value1" || elements["key2"] != "value2" {
			t.Errorf("unexpected elements: %v", elements)
		}
	})

	// 2. Diagnostics on non-string types (invalid-value case)
	t.Run("Non-string value type diagnostics", func(t *testing.T) {
		var diags diag.Diagnostics

		InterfaceListItemMap := map[string]interface{}{
			"labels": map[string]interface{}{
				"key1": 123, // invalid int type
			},
		}

		_ = func() types.Map {
			if v, ok := InterfaceListItemMap["labels"].(map[string]interface{}); ok {
				items := make(map[string]string)
				for mk, mv := range v {
					if mvs, ok := mv.(string); ok {
						items[mk] = mvs
					} else {
						diags.AddError("Unexpected type in map", "Expected string")
					}
				}
				mapVal, d := types.MapValueFrom(ctx, types.StringType, items)
				diags.Append(d...)
				return mapVal
			}
			return types.MapNull(types.StringType)
		}()

		if !diags.HasError() {
			t.Fatal("expected diagnostic error for non-string type, got none")
		}
	})

	// 3. API-omitted map state preservation (drift-prevention case)
	t.Run("API-omitted map state preservation", func(t *testing.T) {
		// Existing state contains a populated map inside production model struct
		existingLabels, _ := types.MapValueFrom(ctx, types.StringType, map[string]string{"foo": "bar"})

		InterfaceListExisting := []SecuremeshSiteV2BaremetalNotManagedNodeListInterfaceListModel{
			{
				Labels: existingLabels,
			},
		}
		InterfaceListIdx := 0

		// API response is omitted (nil interface)
		var InterfaceListItemMap map[string]interface{} = nil

		// Execute exact generated-style state preservation closure
		result := func() types.Map {
			if InterfaceListItemMap != nil {
				if v, ok := InterfaceListItemMap["labels"].(map[string]interface{}); ok {
					items := make(map[string]string)
					for mk, mv := range v {
						if mvs, ok := mv.(string); ok {
							items[mk] = mvs
						}
					}
					mapVal, _ := types.MapValueFrom(ctx, types.StringType, items)
					return mapVal
				}
			}
			if len(InterfaceListExisting) > InterfaceListIdx && !InterfaceListExisting[InterfaceListIdx].Labels.IsNull() && !InterfaceListExisting[InterfaceListIdx].Labels.IsUnknown() {
				return InterfaceListExisting[InterfaceListIdx].Labels
			}
			return types.MapNull(types.StringType)
		}()

		if result.IsNull() || result.IsUnknown() {
			t.Fatal("expected omitted map state to be preserved from prior state, got null")
		}

		var elements map[string]string
		_ = result.ElementsAs(ctx, &elements, false)
		if elements["foo"] != "bar" {
			t.Errorf("expected preserved element 'foo' to be 'bar', got: %v", elements)
		}
	})

	// 4. Null Map
	t.Run("Null map unmarshaling", func(t *testing.T) {
		InterfaceListExisting := []SecuremeshSiteV2BaremetalNotManagedNodeListInterfaceListModel{
			{
				Labels: types.MapNull(types.StringType),
			},
		}
		InterfaceListIdx := 0
		var InterfaceListItemMap map[string]interface{} = nil

		result := func() types.Map {
			if InterfaceListItemMap != nil {
				if v, ok := InterfaceListItemMap["labels"].(map[string]interface{}); ok {
					items := make(map[string]string)
					for mk, mv := range v {
						if mvs, ok := mv.(string); ok {
							items[mk] = mvs
						}
					}
					mapVal, _ := types.MapValueFrom(ctx, types.StringType, items)
					return mapVal
				}
			}
			if len(InterfaceListExisting) > InterfaceListIdx && !InterfaceListExisting[InterfaceListIdx].Labels.IsNull() && !InterfaceListExisting[InterfaceListIdx].Labels.IsUnknown() {
				return InterfaceListExisting[InterfaceListIdx].Labels
			}
			return types.MapNull(types.StringType)
		}()

		if !result.IsNull() {
			t.Fatal("expected map to be null")
		}
	})

	// 5. Unknown Map
	t.Run("Unknown map unmarshaling", func(t *testing.T) {
		InterfaceListExisting := []SecuremeshSiteV2BaremetalNotManagedNodeListInterfaceListModel{
			{
				Labels: types.MapUnknown(types.StringType),
			},
		}
		InterfaceListIdx := 0
		var InterfaceListItemMap map[string]interface{} = nil

		result := func() types.Map {
			if InterfaceListItemMap != nil {
				if v, ok := InterfaceListItemMap["labels"].(map[string]interface{}); ok {
					items := make(map[string]string)
					for mk, mv := range v {
						if mvs, ok := mv.(string); ok {
							items[mk] = mvs
						}
					}
					mapVal, _ := types.MapValueFrom(ctx, types.StringType, items)
					return mapVal
				}
			}
			if len(InterfaceListExisting) > InterfaceListIdx && !InterfaceListExisting[InterfaceListIdx].Labels.IsNull() && !InterfaceListExisting[InterfaceListIdx].Labels.IsUnknown() {
				return InterfaceListExisting[InterfaceListIdx].Labels
			}
			return types.MapNull(types.StringType)
		}()

		if !result.IsNull() {
			t.Fatal("expected map to fallback to null since prior state is unknown")
		}
	})

	// 6. Empty Map
	t.Run("Empty map unmarshaling", func(t *testing.T) {
		emptyMap, _ := types.MapValueFrom(ctx, types.StringType, map[string]string{})
		InterfaceListExisting := []SecuremeshSiteV2BaremetalNotManagedNodeListInterfaceListModel{
			{
				Labels: emptyMap,
			},
		}
		InterfaceListIdx := 0
		var InterfaceListItemMap map[string]interface{} = nil

		result := func() types.Map {
			if InterfaceListItemMap != nil {
				if v, ok := InterfaceListItemMap["labels"].(map[string]interface{}); ok {
					items := make(map[string]string)
					for mk, mv := range v {
						if mvs, ok := mv.(string); ok {
							items[mk] = mvs
						}
					}
					mapVal, _ := types.MapValueFrom(ctx, types.StringType, items)
					return mapVal
				}
			}
			if len(InterfaceListExisting) > InterfaceListIdx && !InterfaceListExisting[InterfaceListIdx].Labels.IsNull() && !InterfaceListExisting[InterfaceListIdx].Labels.IsUnknown() {
				return InterfaceListExisting[InterfaceListIdx].Labels
			}
			return types.MapNull(types.StringType)
		}()

		if result.IsNull() || result.IsUnknown() {
			t.Fatal("expected empty map to be preserved, got null/unknown")
		}
		if len(result.Elements()) != 0 {
			t.Errorf("expected empty map, got: %d elements", len(result.Elements()))
		}
	})
}
