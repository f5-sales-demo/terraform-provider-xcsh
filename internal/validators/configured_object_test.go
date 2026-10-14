// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package validators

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestConfiguredObjectDefersSyntheticRootBlocks(t *testing.T) {
	blockType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"choice": tftypes.String, "optional": tftypes.String}}
	rootType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"branch": blockType}}
	block := func(choice, optional interface{}) tftypes.Value {
		return tftypes.NewValue(blockType, map[string]tftypes.Value{"choice": tftypes.NewValue(tftypes.String, choice), "optional": tftypes.NewValue(tftypes.String, optional)})
	}
	for name, test := range map[string]struct {
		value   tftypes.Value
		present bool
	}{
		"null":                                  {tftypes.NewValue(blockType, nil), false},
		"unknown":                               {tftypes.NewValue(blockType, tftypes.UnknownValue), false},
		"synthetic unknown descendants":         {block(tftypes.UnknownValue, tftypes.UnknownValue), false},
		"unknown and omitted descendants":       {block(tftypes.UnknownValue, nil), false},
		"explicit empty block":                  {block(nil, nil), true},
		"known empty string":                    {block("", nil), true},
		"known selection with unknown optional": {block("chosen", tftypes.UnknownValue), true},
	} {
		t.Run(name, func(t *testing.T) {
			config := tfsdk.Config{Raw: tftypes.NewValue(rootType, map[string]tftypes.Value{"branch": test.value}), Schema: schema.Schema{Blocks: map[string]schema.Block{"branch": schema.SingleNestedBlock{Attributes: map[string]schema.Attribute{"choice": schema.StringAttribute{Optional: true}, "optional": schema.StringAttribute{Optional: true}}}}}}
			var diagnostics diag.Diagnostics
			got := ConfiguredObject(context.Background(), config, path.Root("branch"), &diagnostics)
			if diagnostics.HasError() || got != test.present {
				t.Fatalf("present=%v want=%v diagnostics=%v", got, test.present, diagnostics)
			}
		})
	}
}

func TestConfiguredListDefersSyntheticRootElements(t *testing.T) {
	blockType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"choice": tftypes.String}}
	listType := tftypes.List{ElementType: blockType}
	rootType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"branches": listType}}
	block := func(choice interface{}) tftypes.Value {
		return tftypes.NewValue(blockType, map[string]tftypes.Value{"choice": tftypes.NewValue(tftypes.String, choice)})
	}
	list := func(elements ...tftypes.Value) tftypes.Value { return tftypes.NewValue(listType, elements) }
	for name, test := range map[string]struct {
		value   tftypes.Value
		present bool
	}{
		"null":                 {tftypes.NewValue(listType, nil), false},
		"unknown":              {tftypes.NewValue(listType, tftypes.UnknownValue), false},
		"empty":                {list(), false},
		"unknown element":      {list(tftypes.NewValue(blockType, tftypes.UnknownValue)), false},
		"unknown descendants":  {list(block(tftypes.UnknownValue)), false},
		"known empty block":    {list(block(nil)), true},
		"known selection":      {list(block("chosen")), true},
		"known and unresolved": {list(block("chosen"), block(tftypes.UnknownValue)), true},
	} {
		t.Run(name, func(t *testing.T) {
			config := tfsdk.Config{Raw: tftypes.NewValue(rootType, map[string]tftypes.Value{"branches": test.value}), Schema: schema.Schema{Blocks: map[string]schema.Block{"branches": schema.ListNestedBlock{NestedObject: schema.NestedBlockObject{Attributes: map[string]schema.Attribute{"choice": schema.StringAttribute{Optional: true}}}}}}}
			var diagnostics diag.Diagnostics
			got := ConfiguredList(context.Background(), config, path.Root("branches"), &diagnostics)
			if diagnostics.HasError() || got != test.present {
				t.Fatalf("present=%v want=%v diagnostics=%v", got, test.present, diagnostics)
			}
		})
	}
}
