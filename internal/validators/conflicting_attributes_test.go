// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package validators

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestConflictingObjectAttributes(t *testing.T) {
	fields := map[string]attr.Type{"a": types.Int64Type, "b": types.ObjectType{AttrTypes: map[string]attr.Type{}}}
	empty := types.ObjectValueMust(map[string]attr.Type{}, map[string]attr.Value{})
	absent := types.ObjectNull(map[string]attr.Type{})
	object := func(a types.Int64, b types.Object) types.Object {
		return types.ObjectValueMust(fields, map[string]attr.Value{"a": a, "b": b})
	}
	for name, test := range map[string]struct {
		value    types.Object
		conflict bool
	}{
		"absent parent":          {types.ObjectNull(fields), false},
		"unknown parent":         {types.ObjectUnknown(fields), false},
		"both absent":            {object(types.Int64Null(), absent), false},
		"scalar only":            {object(types.Int64Value(0), absent), false},
		"block only":             {object(types.Int64Null(), empty), false},
		"unknown scalar":         {object(types.Int64Unknown(), empty), false},
		"unknown block":          {object(types.Int64Value(1), types.ObjectUnknown(map[string]attr.Type{})), false},
		"zero and empty block":   {object(types.Int64Value(0), empty), true},
		"number and empty block": {object(types.Int64Value(50), empty), true},
	} {
		t.Run(name, func(t *testing.T) {
			var response validator.ObjectResponse
			ConflictingObjectAttributes("a", "b").ValidateObject(context.Background(), validator.ObjectRequest{ConfigValue: test.value, Path: path.Root("parent")}, &response)
			if response.Diagnostics.HasError() != test.conflict {
				t.Fatalf("unexpected diagnostics: %v", response.Diagnostics)
			}
			if test.conflict && !response.Diagnostics[0].(diag.DiagnosticWithPath).Path().Equal(path.Root("parent").AtName("b")) {
				t.Fatal("diagnostic lost nested attribute path")
			}
		})
	}
}

func TestConflictingObjectAttributesDefersUnresolvedDynamicBlocks(t *testing.T) {
	childTypes := map[string]attr.Type{"choice": types.ObjectType{AttrTypes: map[string]attr.Type{}}}
	blockType := types.ObjectType{AttrTypes: childTypes}
	parentTypes := map[string]attr.Type{"left": blockType, "right": blockType}
	unresolved := types.ObjectValueMust(childTypes, map[string]attr.Value{
		"choice": types.ObjectUnknown(map[string]attr.Type{}),
	})
	parent := types.ObjectValueMust(parentTypes, map[string]attr.Value{
		"left": unresolved, "right": unresolved,
	})

	var response validator.ObjectResponse
	ConflictingObjectAttributes("left", "right").ValidateObject(
		context.Background(),
		validator.ObjectRequest{ConfigValue: parent, Path: path.Root("parent")},
		&response,
	)
	if response.Diagnostics.HasError() {
		t.Fatalf("unresolved dynamic branches must defer conflict validation: %v", response.Diagnostics)
	}
}

func TestConflictingListObjectAttributes(t *testing.T) {
	fields := map[string]attr.Type{"a": types.BoolType, "b": types.StringType}
	elementType := types.ObjectType{AttrTypes: fields}
	object := func(a types.Bool, b types.String) types.Object {
		return types.ObjectValueMust(fields, map[string]attr.Value{"a": a, "b": b})
	}
	left := object(types.BoolValue(false), types.StringNull())
	right := object(types.BoolNull(), types.StringValue(""))
	both := object(types.BoolValue(false), types.StringValue(""))
	for name, test := range map[string]struct {
		value    types.List
		conflict bool
	}{
		"null list":              {types.ListNull(elementType), false},
		"unknown list":           {types.ListUnknown(elementType), false},
		"separate elements":      {types.ListValueMust(elementType, []attr.Value{left, right}), false},
		"unknown element":        {types.ListValueMust(elementType, []attr.Value{left, types.ObjectUnknown(fields)}), false},
		"null element":           {types.ListValueMust(elementType, []attr.Value{left, types.ObjectNull(fields)}), false},
		"false and empty string": {types.ListValueMust(elementType, []attr.Value{left, both}), true},
	} {
		t.Run(name, func(t *testing.T) {
			var response validator.ListResponse
			ConflictingListObjectAttributes("a", "b").ValidateList(context.Background(), validator.ListRequest{ConfigValue: test.value, Path: path.Root("parents")}, &response)
			if response.Diagnostics.HasError() != test.conflict {
				t.Fatalf("unexpected diagnostics: %v", response.Diagnostics)
			}
			if test.conflict && !response.Diagnostics[0].(diag.DiagnosticWithPath).Path().Equal(path.Root("parents").AtListIndex(1).AtName("b")) {
				t.Fatal("diagnostic lost list element path")
			}
		})
	}
}
