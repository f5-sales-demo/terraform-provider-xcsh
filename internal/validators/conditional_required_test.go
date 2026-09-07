// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package validators

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestRequiredObjectAttributes(t *testing.T) {
	v := RequiredObjectAttributes("required_value")
	attrTypes := map[string]attr.Type{"required_value": types.StringType}
	for name, test := range map[string]struct {
		value     types.Object
		wantError bool
	}{
		"absent parent":    {value: types.ObjectNull(attrTypes)},
		"unknown parent":   {value: types.ObjectUnknown(attrTypes)},
		"missing child":    {value: types.ObjectValueMust(attrTypes, map[string]attr.Value{"required_value": types.StringNull()}), wantError: true},
		"unknown child":    {value: types.ObjectValueMust(attrTypes, map[string]attr.Value{"required_value": types.StringUnknown()})},
		"configured child": {value: types.ObjectValueMust(attrTypes, map[string]attr.Value{"required_value": types.StringValue("set")})},
	} {
		t.Run(name, func(t *testing.T) {
			resp := &validator.ObjectResponse{}
			v.ValidateObject(context.Background(), validator.ObjectRequest{ConfigValue: test.value, Path: path.Root("parent")}, resp)
			if got := resp.Diagnostics.HasError(); got != test.wantError {
				t.Fatalf("HasError=%v, want %v: %v", got, test.wantError, resp.Diagnostics)
			}
		})
	}
}

func TestRequiredListObjectAttributes(t *testing.T) {
	objectTypes := map[string]attr.Type{"required_value": types.Int64Type}
	objectType := types.ObjectType{AttrTypes: objectTypes}
	v := RequiredListObjectAttributes("required_value")
	configured := types.ObjectValueMust(objectTypes, map[string]attr.Value{"required_value": types.Int64Value(1)})
	missing := types.ObjectValueMust(objectTypes, map[string]attr.Value{"required_value": types.Int64Null()})

	for name, test := range map[string]struct {
		value     types.List
		wantError bool
	}{
		"absent parent":       {value: types.ListNull(objectType)},
		"unknown parent":      {value: types.ListUnknown(objectType)},
		"configured elements": {value: types.ListValueMust(objectType, []attr.Value{configured})},
		"missing child":       {value: types.ListValueMust(objectType, []attr.Value{configured, missing}), wantError: true},
	} {
		t.Run(name, func(t *testing.T) {
			resp := &validator.ListResponse{}
			v.ValidateList(context.Background(), validator.ListRequest{ConfigValue: test.value, Path: path.Root("parents")}, resp)
			if got := resp.Diagnostics.HasError(); got != test.wantError {
				t.Fatalf("HasError=%v, want %v: %v", got, test.wantError, resp.Diagnostics)
			}
		})
	}
}

func TestRequiredOneOfObjectAttributes(t *testing.T) {
	fields := map[string]attr.Type{"a": types.StringType, "b": types.StringType, "c": types.StringType}
	object := func(a, b, c types.String) types.Object {
		return types.ObjectValueMust(fields, map[string]attr.Value{"a": a, "b": b, "c": c})
	}
	for name, test := range map[string]struct {
		value     types.Object
		wantError bool
	}{
		"absent parent":  {value: types.ObjectNull(fields)},
		"unknown parent": {value: types.ObjectUnknown(fields)},
		"none selected":  {value: object(types.StringNull(), types.StringNull(), types.StringNull()), wantError: true},
		"one selected":   {value: object(types.StringNull(), types.StringValue("set"), types.StringNull())},
		"unknown defers": {value: object(types.StringNull(), types.StringUnknown(), types.StringNull())},
	} {
		t.Run(name, func(t *testing.T) {
			resp := &validator.ObjectResponse{}
			RequiredOneOfObjectAttributes("a", "b", "c").ValidateObject(context.Background(), validator.ObjectRequest{ConfigValue: test.value, Path: path.Root("parent")}, resp)
			if got := resp.Diagnostics.HasError(); got != test.wantError {
				t.Fatalf("HasError=%v, want %v: %v", got, test.wantError, resp.Diagnostics)
			}
		})
	}
}

func TestRequiredOneOfListObjectAttributes(t *testing.T) {
	fields := map[string]attr.Type{"a": types.BoolType, "b": types.BoolType}
	elementType := types.ObjectType{AttrTypes: fields}
	selected := types.ObjectValueMust(fields, map[string]attr.Value{"a": types.BoolValue(false), "b": types.BoolNull()})
	missing := types.ObjectValueMust(fields, map[string]attr.Value{"a": types.BoolNull(), "b": types.BoolNull()})
	unknown := types.ObjectValueMust(fields, map[string]attr.Value{"a": types.BoolUnknown(), "b": types.BoolNull()})
	for name, test := range map[string]struct {
		value     types.List
		wantError bool
	}{
		"absent parent":  {value: types.ListNull(elementType)},
		"unknown parent": {value: types.ListUnknown(elementType)},
		"one selected":   {value: types.ListValueMust(elementType, []attr.Value{selected})},
		"unknown defers": {value: types.ListValueMust(elementType, []attr.Value{unknown})},
		"missing choice": {value: types.ListValueMust(elementType, []attr.Value{selected, missing}), wantError: true},
	} {
		t.Run(name, func(t *testing.T) {
			resp := &validator.ListResponse{}
			RequiredOneOfListObjectAttributes("a", "b").ValidateList(context.Background(), validator.ListRequest{ConfigValue: test.value, Path: path.Root("parents")}, resp)
			if got := resp.Diagnostics.HasError(); got != test.wantError {
				t.Fatalf("HasError=%v, want %v: %v", got, test.wantError, resp.Diagnostics)
			}
		})
	}
}
