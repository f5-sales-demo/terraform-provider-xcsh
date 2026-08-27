// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package validators

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// RequiredObjectAttributes requires named direct children whenever the object
// itself is configured. Null and unknown parent objects are skipped.
func RequiredObjectAttributes(names ...string) validator.Object {
	return requiredObjectAttributesValidator{names: append([]string(nil), names...)}
}

type requiredObjectAttributesValidator struct{ names []string }

func (v requiredObjectAttributesValidator) Description(context.Context) string {
	return "requires these attributes when the containing block is configured: " + strings.Join(v.names, ", ")
}

func (v requiredObjectAttributesValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v requiredObjectAttributesValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	validateRequiredObjectAttributes(req.ConfigValue, req.Path, v.names, resp)
}

// RequiredListObjectAttributes applies RequiredObjectAttributes to every known
// object element of a configured nested list block.
func RequiredListObjectAttributes(names ...string) validator.List {
	return requiredListObjectAttributesValidator{names: append([]string(nil), names...)}
}

type requiredListObjectAttributesValidator struct{ names []string }

func (v requiredListObjectAttributesValidator) Description(context.Context) string {
	return "requires these attributes in each configured block: " + strings.Join(v.names, ", ")
}

func (v requiredListObjectAttributesValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v requiredListObjectAttributesValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	for index, element := range req.ConfigValue.Elements() {
		object, ok := element.(types.Object)
		if !ok || object.IsNull() || object.IsUnknown() {
			continue
		}
		objectResp := &validator.ObjectResponse{}
		validateRequiredObjectAttributes(object, req.Path.AtListIndex(index), v.names, objectResp)
		resp.Diagnostics.Append(objectResp.Diagnostics...)
	}
}

func validateRequiredObjectAttributes(object types.Object, objectPath path.Path, names []string, resp *validator.ObjectResponse) {
	attributes := object.Attributes()
	for _, name := range names {
		value, ok := attributes[name]
		if ok && (!value.IsNull() || value.IsUnknown()) {
			continue
		}
		attributePath := objectPath.AtName(name)
		resp.Diagnostics.AddAttributeError(
			attributePath,
			"Missing Required Attribute in Configured Block",
			fmt.Sprintf("Attribute %q must be configured when this block is present.", name),
		)
	}
}
