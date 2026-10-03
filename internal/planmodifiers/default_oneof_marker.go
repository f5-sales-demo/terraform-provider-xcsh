package planmodifiers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// DefaultOneOfMarker clears an omitted computed marker when configuration
// explicitly selects a sibling. It never changes a configured marker or assumes
// that an unknown sibling has been selected.
func DefaultOneOfMarker(member string, siblings []string) planmodifier.Object {
	return defaultOneOfMarker{member: member, siblings: siblings}
}

type defaultOneOfMarker struct {
	member   string
	siblings []string
}

func (m defaultOneOfMarker) Description(context.Context) string {
	return "Clear an omitted default oneof marker when a sibling is explicitly selected."
}
func (m defaultOneOfMarker) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}
func (m defaultOneOfMarker) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if !req.ConfigValue.IsNull() || req.Config.Raw.IsNull() {
		return
	}
	for _, sibling := range m.siblings {
		if sibling == m.member {
			continue
		}
		value, err := selectionMember(req.Config.Raw, sibling)
		if err != nil {
			resp.Diagnostics.AddError("Invalid Default Selection Contract", fmt.Sprintf("Cannot inspect sibling %s: %s", sibling, err))
			return
		}
		if value.IsKnown() && !value.IsNull() {
			resp.PlanValue = types.ObjectNull(req.PlanValue.AttributeTypes(ctx))
			return
		}
	}
}
