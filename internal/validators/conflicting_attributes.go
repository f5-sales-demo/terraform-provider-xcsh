// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package validators

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type conflictingAttributesValidator struct{ left, right string }

type configurationPresence uint8

const (
	configurationAbsent configurationPresence = iota
	configurationUnresolved
	configurationPresent
)

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
	if !leftExists || !rightExists || configuredPresence(left) != configurationPresent || configuredPresence(right) != configurationPresent {
		return
	}
	diagnostics.AddAttributeError(location.AtName(v.right), "Conflicting Configuration", v.Description(ctx))
}

// configuredPresence distinguishes a configured block from the synthetic
// object Terraform constructs while a dynamic block's for_each is unresolved.
// The latter has only unknown descendants and must defer validation. A known
// empty object remains configured because empty protobuf oneof members use it
// as their payload.
func configuredPresence(value attr.Value) configurationPresence {
	if value == nil || value.IsNull() {
		return configurationAbsent
	}
	if value.IsUnknown() {
		return configurationUnresolved
	}
	switch typed := value.(type) {
	case types.Object:
		if len(typed.Attributes()) == 0 {
			return configurationPresent
		}
		unresolved := false
		for _, child := range typed.Attributes() {
			switch configuredPresence(child) {
			case configurationPresent:
				return configurationPresent
			case configurationUnresolved:
				unresolved = true
			}
		}
		if unresolved {
			return configurationUnresolved
		}
		// A known object whose optional children are all null represents an
		// explicitly configured empty block.
		return configurationPresent
	case types.List:
		if len(typed.Elements()) == 0 {
			return configurationAbsent
		}
	}
	return configurationPresent
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
