// Package planmodifiers implements Terraform planning policies.
package planmodifiers

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// RequireImmutableOneOfSelection compares member presence without decoding child
// settings. Omission is distinct from selection, and unknown presence cannot
// establish that the selection remains unchanged.
func RequireImmutableOneOfSelection(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse, group string, members []string) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}
	old, next := []string{}, []string{}
	affected := []path.Path{}
	for _, member := range members {
		before, err := selectionMember(req.State.Raw, member)
		if err != nil {
			resp.Diagnostics.AddError("Invalid Immutable Selection Contract", fmt.Sprintf("Cannot inspect %s state member %s: %s", group, member, err))
			return
		}
		after, err := selectionMember(req.Plan.Raw, member)
		if err != nil {
			resp.Diagnostics.AddError("Invalid Immutable Selection Contract", fmt.Sprintf("Cannot inspect %s plan member %s: %s", group, member, err))
			return
		}
		if !before.IsKnown() {
			old = append(old, member+" (unknown presence)")
		} else if !before.IsNull() {
			old = append(old, member)
		}
		if !after.IsKnown() {
			next = append(next, member+" (unknown presence)")
		} else if !after.IsNull() {
			next = append(next, member)
		}
		if !before.IsKnown() || !after.IsKnown() || before.IsNull() != after.IsNull() {
			affected = append(affected, path.Root(member))
		}
	}
	if len(affected) == 0 {
		return
	}
	label := func(selection []string) string {
		if len(selection) == 0 {
			return "omitted"
		}
		return strings.Join(selection, ", ")
	}
	resp.RequiresReplace = append(resp.RequiresReplace, affected...)
	resp.Diagnostics.AddWarning(
		"Load Balancer Type Requires Replacement",
		fmt.Sprintf("The %s selection changes from %s to %s, or its presence is unknown. F5 Distributed Cloud does not support changing this selection in place. Terraform must recreate the resource, which may interrupt service. Supported settings within the same selected type remain updatable.", group, label(old), label(next)),
	)
}

func selectionMember(raw tftypes.Value, member string) (tftypes.Value, error) {
	value, err := raw.ApplyTerraform5AttributePathStep(tftypes.AttributeName(member))
	if err != nil {
		return tftypes.Value{}, err
	}
	result, ok := value.(tftypes.Value)
	if !ok {
		return tftypes.Value{}, fmt.Errorf("member %s is not a Terraform value", member)
	}
	return result, nil
}
