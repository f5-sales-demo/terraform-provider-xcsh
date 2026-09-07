// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package validators

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// RequiredOneOfObjectAttributes requires at least one known, configured child.
// A separate conflicting-attributes validator enforces that no more than one is
// configured. Null parents remain optional and unknown values defer validation.
func RequiredOneOfObjectAttributes(names ...string) validator.Object {
	return requiredOneOfAttributesValidator{names: append([]string(nil), names...)}
}

// RequiredOneOfListObjectAttributes applies RequiredOneOfObjectAttributes to
// every known object element of a configured nested list block.
func RequiredOneOfListObjectAttributes(names ...string) validator.List {
	return requiredOneOfAttributesValidator{names: append([]string(nil), names...)}
}

type requiredOneOfAttributesValidator struct{ names []string }

func (v requiredOneOfAttributesValidator) Description(context.Context) string {
	return "requires one of these attributes when the containing block is configured: " + strings.Join(v.names, ", ")
}

func (v requiredOneOfAttributesValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v requiredOneOfAttributesValidator) validate(ctx context.Context, object types.Object, location path.Path, diagnostics *diag.Diagnostics) {
	if object.IsNull() || object.IsUnknown() {
		return
	}
	unresolved := false
	attributes := object.Attributes()
	for _, name := range v.names {
		value, ok := attributes[name]
		if !ok {
			continue
		}
		switch configuredPresence(value) {
		case configurationPresent:
			return
		case configurationUnresolved:
			unresolved = true
		}
	}
	if unresolved {
		return
	}
	diagnostics.AddAttributeError(
		location,
		"Missing Required Choice in Configured Block",
		fmt.Sprintf("Exactly one of %s must be configured when this block is present.", strings.Join(v.names, ", ")),
	)
}

func (v requiredOneOfAttributesValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	v.validate(ctx, req.ConfigValue, req.Path, &resp.Diagnostics)
}

func (v requiredOneOfAttributesValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	for index, element := range req.ConfigValue.Elements() {
		if object, ok := element.(types.Object); ok {
			v.validate(ctx, object, req.Path.AtListIndex(index), &resp.Diagnostics)
		}
	}
}
