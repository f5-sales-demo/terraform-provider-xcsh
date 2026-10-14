// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package validators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ConfiguredObject distinguishes a selected root block from the synthetic
// unknown descendants Terraform supplies for an unresolved dynamic for_each.
// It shares presence semantics with nested oneof validators, including known
// empty blocks. Decoded model pointers cannot retain this distinction.
func ConfiguredObject(ctx context.Context, config tfsdk.Config, location path.Path, diagnostics *diag.Diagnostics) bool {
	var value types.Object
	diagnostics.Append(config.GetAttribute(ctx, location, &value)...)
	if diagnostics.HasError() {
		return false
	}
	return configuredPresence(value) == configurationPresent
}

// ConfiguredList applies the same check to repeated root blocks, retaining
// known nonempty selections while deferring unresolved synthetic elements.
func ConfiguredList(ctx context.Context, config tfsdk.Config, location path.Path, diagnostics *diag.Diagnostics) bool {
	var value types.List
	diagnostics.Append(config.GetAttribute(ctx, location, &value)...)
	if diagnostics.HasError() {
		return false
	}
	return configuredPresence(value) == configurationPresent
}
