// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package conflicts

import (
	"fmt"
	"strings"
)

// Attr describes an attribute with conflict relationships.
type Attr struct {
	TfsdkTag      string
	GoName        string
	ConflictsWith []string
}

// GenerateChecks returns Go code for ValidateConfig body that checks mutual exclusivity.
// Deduplicates: if A conflicts with B and B conflicts with A, only one check is emitted.
// The goNameLookup map resolves tfsdk tags to Go field names for conflict targets.
// If a conflict target is not found in the lookup, it is skipped (it may be a block or
// an attribute that doesn't exist in the model).
func GenerateChecks(attrs []Attr, goNameLookup map[string]string) string {
	var sb strings.Builder
	seen := make(map[string]bool)

	for _, attr := range attrs {
		for _, conflict := range attr.ConflictsWith {
			conflictGoName, ok := goNameLookup[conflict]
			if !ok {
				// Conflict target not in lookup — skip (may be a block or missing attribute)
				continue
			}

			pairKey := attr.TfsdkTag + ":" + conflict
			reverseKey := conflict + ":" + attr.TfsdkTag
			if seen[reverseKey] {
				continue
			}
			seen[pairKey] = true

			sb.WriteString(fmt.Sprintf("\tif !data.%s.IsNull() && !data.%s.IsNull() {\n", attr.GoName, conflictGoName))
			sb.WriteString("\t\tresp.Diagnostics.AddAttributeError(\n")
			sb.WriteString(fmt.Sprintf("\t\t\tpath.Root(%q),\n", attr.TfsdkTag))
			sb.WriteString("\t\t\t\"Conflicting Configuration\",\n")
			sb.WriteString(fmt.Sprintf("\t\t\t\"%s and %s are mutually exclusive.\",\n", attr.TfsdkTag, conflict))
			sb.WriteString("\t\t)\n\t}\n")
		}
	}
	return sb.String()
}

// TfsdkToGoName converts snake_case tfsdk tag to TitleCase Go field name.
func TfsdkToGoName(s string) string {
	parts := strings.Split(s, "_")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, "")
}
