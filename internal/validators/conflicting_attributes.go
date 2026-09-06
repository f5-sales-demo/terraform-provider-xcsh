// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package validators

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type conflictingAttributesValidator struct{ left, right string }

// ConflictingObjectAttributes rejects two known, non-null sibling values,
// including empty blocks and scalar zero values. Unknown values defer validation.
func ConflictingObjectAttributes(left, right string) validator.Object {
	return conflictingAttributesValidator{left: left, right: right}
}

// ConflictingListObjectAttributes checks each object independently, retaining
// its list index in diagnostics. Values in different elements cannot conflict.
func ConflictingListObjectAttributes(left, right string) validator.List {
	return conflictingAttributesValidator{left: left, right: right}
}

func (v conflictingAttributesValidator) Description(context.Context) string {
	return fmt.Sprintf("%q and %q are mutually exclusive.", v.left, v.right)
}

func (v conflictingAttributesValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v conflictingAttributesValidator) validate(ctx context.Context, value types.Object, location path.Path, diagnostics *diag.Diagnostics) {
	if value.IsNull() || value.IsUnknown() {
		return
	}
	attributes := value.Attributes()
	left, leftExists := attributes[v.left]
	right, rightExists := attributes[v.right]
	if !leftExists || !rightExists || left.IsNull() || right.IsNull() || left.IsUnknown() || right.IsUnknown() {
		return
	}
	diagnostics.AddAttributeError(location.AtName(v.right), "Conflicting Configuration", v.Description(ctx))
}

func (v conflictingAttributesValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	v.validate(ctx, req.ConfigValue, req.Path, &resp.Diagnostics)
}

func (v conflictingAttributesValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	for index, element := range req.ConfigValue.Elements() {
		if object, ok := element.(types.Object); ok {
			v.validate(ctx, object, req.Path.AtListIndex(index), &resp.Diagnostics)
		}
	}
}
