// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package codegen

import (
	"fmt"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"sort"
	"strings"
)

func RenderRootBlockConflicts(attributes []openapi.TerraformAttribute) string {
	byName := map[string]openapi.TerraformAttribute{}
	groups := map[string][]string{}
	for _, attribute := range attributes {
		byName[attribute.TfsdkTag] = attribute
		if attribute.OneOfGroup != "" {
			groups[attribute.OneOfGroup] = append(groups[attribute.OneOfGroup], attribute.TfsdkTag)
		}
	}
	pairs := map[[2]string]bool{}
	add := func(a, b string) {
		if a == b {
			return
		}
		if a > b {
			a, b = b, a
		}
		pairs[[2]string{a, b}] = true
	}
	for _, attribute := range attributes {
		for _, target := range attribute.ConflictsWith {
			if _, ok := byName[target]; ok {
				add(attribute.TfsdkTag, target)
			}
		}
	}
	for _, members := range groups {
		for i, a := range members {
			for _, b := range members[i+1:] {
				add(a, b)
			}
		}
	}
	keys := [][2]string{}
	for pair := range pairs {
		keys = append(keys, pair)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i][0]+"/"+keys[i][1] < keys[j][0]+"/"+keys[j][1] })
	configured := func(attribute openapi.TerraformAttribute) string {
		if attribute.IsBlock {
			if attribute.NestedBlockType == "list" {
				return fmt.Sprintf("validators.ConfiguredList(ctx, req.Config, path.Root(%q), &resp.Diagnostics)", attribute.TfsdkTag)
			}
			return fmt.Sprintf("validators.ConfiguredObject(ctx, req.Config, path.Root(%q), &resp.Diagnostics)", attribute.TfsdkTag)
		}
		return fmt.Sprintf("!data.%s.IsNull() && !data.%s.IsUnknown()", attribute.GoName, attribute.GoName)
	}
	var output strings.Builder
	for _, pair := range keys {
		a, b := byName[pair[0]], byName[pair[1]]
		if !a.IsBlock && !b.IsBlock {
			continue
		}
		fmt.Fprintf(&output, "\tif %s && %s {\n\t\tresp.Diagnostics.AddAttributeError(path.Root(%q), \"Conflicting Configuration\", %q)\n\t}\n", configured(a), configured(b), pair[0], pair[0]+" and "+pair[1]+" are mutually exclusive.")
	}
	return output.String()
}
