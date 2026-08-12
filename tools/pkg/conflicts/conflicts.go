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
	IsPointer     bool
	ConflictsWith []string
}

// Field describes how a top-level Terraform field is represented in the
// generated resource model. Single nested blocks are pointers; attributes and
// list nested blocks are framework values with IsNull/IsUnknown methods.
type Field struct {
	GoName    string
	IsPointer bool
}

// GenerateChecks returns Go code for ValidateConfig body that checks mutual exclusivity.
// Deduplicates: if A conflicts with B and B conflicts with A, only one check is emitted.
// The fieldLookup map resolves tfsdk tags to generated model fields.
// If a conflict target is not found in the lookup, it is skipped because it does not
// exist in the generated model.
func GenerateChecks(attrs []Attr, fieldLookup map[string]Field) string {
	var sb strings.Builder
	seen := make(map[string]bool)

	for _, attr := range attrs {
		for _, conflict := range attr.ConflictsWith {
			conflictField, ok := fieldLookup[conflict]
			if !ok {
				// Conflict target not in lookup — skip the missing field.
				continue
			}

			pairKey := attr.TfsdkTag + ":" + conflict
			reverseKey := conflict + ":" + attr.TfsdkTag
			if seen[reverseKey] {
				continue
			}
			seen[pairKey] = true

			sb.WriteString(fmt.Sprintf("\tif %s && %s {\n", configuredExpression(Field{GoName: attr.GoName, IsPointer: attr.IsPointer}), configuredExpression(conflictField)))
			sb.WriteString("\t\tresp.Diagnostics.AddAttributeError(\n")
			sb.WriteString(fmt.Sprintf("\t\t\tpath.Root(%q),\n", attr.TfsdkTag))
			sb.WriteString("\t\t\t\"Conflicting Configuration\",\n")
			sb.WriteString(fmt.Sprintf("\t\t\t\"%s and %s are mutually exclusive.\",\n", attr.TfsdkTag, conflict))
			sb.WriteString("\t\t)\n\t}\n")
		}
	}
	return sb.String()
}

func configuredExpression(field Field) string {
	if field.IsPointer {
		return fmt.Sprintf("data.%s != nil", field.GoName)
	}
	return fmt.Sprintf("!data.%s.IsNull() && !data.%s.IsUnknown()", field.GoName, field.GoName)
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
