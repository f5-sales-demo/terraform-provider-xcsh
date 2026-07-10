// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"sort"
	"strings"
)

// RequiresEntry is one x-f5xc-requires descriptor stamped on a spec property by
// api-specs-enriched. It carries both the opaque requires_field string (for
// documentation/other consumers) and, for cross-resource requirements, the
// structured RequiresResource the provider uses to auto-derive an apply-time
// preflight. See api-specs-enriched#967.
type RequiresEntry struct {
	Field            string            `json:"field"`
	RequiresField    string            `json:"requires_field"`
	RequiresResource *RequiresResource `json:"requires_resource"`
	Required         bool              `json:"required"`
	MinItems         int               `json:"min_items"`
	Reason           string            `json:"reason"`
}

// RequiresResource is the structured form of a cross-resource requirement:
// the triggering field needs a resource of the given kind to exist in the given
// scope (currently only "same_namespace" is actionable as a namespace LIST).
type RequiresResource struct {
	Resource string `json:"resource"`
	Scope    string `json:"scope"`
}

// scopeSameNamespace is the only requires_resource scope that maps to a
// namespace-scoped LIST preflight.
const scopeSameNamespace = "same_namespace"

// DeriveRequirementPreflights builds apply-time preflights from a resource's
// CreateSpecType schema's x-f5xc-requires cross-resource requirements. For each
// property whose requirement names a same-namespace resource, resolveListPath
// maps that target resource to its LIST path (with a single %s for the
// namespace); an unresolvable target is skipped so the generator never emits a
// broken list_path. The source of truth is the enriched spec, making
// preflight-requirements.json an override rather than the sole declaration.
func DeriveRequirementPreflights(createSpec Schema, resolveListPath func(resource string) (string, bool)) []RequirementPreflight {
	var out []RequirementPreflight
	// Deterministic order: iterate property names sorted.
	propNames := make([]string, 0, len(createSpec.Properties))
	for name := range createSpec.Properties {
		propNames = append(propNames, name)
	}
	sort.Strings(propNames)

	for _, propName := range propNames {
		for _, entry := range createSpec.Properties[propName].XF5XCRequires {
			rr := entry.RequiresResource
			if rr == nil || rr.Scope != scopeSameNamespace || rr.Resource == "" {
				continue // sibling-field or non-namespace requirement: not a preflight
			}
			listPath, ok := resolveListPath(rr.Resource)
			if !ok || strings.Count(listPath, "%s") != 1 {
				continue // cannot resolve a well-formed LIST path: skip, never emit broken
			}
			field := entry.Field
			if field == "" {
				field = propName
			}
			out = append(out, RequirementPreflight{
				WhenField:   field,
				ListPath:    listPath,
				Requires:    requirementText(field, rr.Resource, entry.Reason),
				ErrorTitle:  displayResource(rr.Resource) + " prerequisite missing",
				ErrorDetail: requirementDetail(field, rr.Resource, entry.Reason),
			})
		}
	}
	return out
}

// MergePreflights overlays override entries (e.g. preflight-requirements.json)
// onto derived entries: for a given WhenField the override wins, so hand-tuned
// declarations keep byte-identical output while derived-only and override-only
// entries both survive. The result is sorted by WhenField for determinism.
func MergePreflights(derived, override []RequirementPreflight) []RequirementPreflight {
	byField := make(map[string]RequirementPreflight, len(derived)+len(override))
	for _, p := range derived {
		byField[p.WhenField] = p
	}
	for _, p := range override {
		byField[p.WhenField] = p
	}
	fields := make([]string, 0, len(byField))
	for f := range byField {
		fields = append(fields, f)
	}
	sort.Strings(fields)
	out := make([]RequirementPreflight, 0, len(fields))
	for _, f := range fields {
		out = append(out, byField[f])
	}
	return out
}

// requirementText renders the human-readable requirement (a code comment).
func requirementText(field, resource, reason string) string {
	if r := sanitizeReason(reason); r != "" {
		return r
	}
	return field + " requires a " + resource + " in the same namespace"
}

// requirementDetail renders the apply-time diagnostic detail. It contains exactly
// one %s (the namespace), which the generated Create/Update fills at runtime; the
// reason is sanitized so it can never introduce a second format verb.
func requirementDetail(field, resource, reason string) string {
	detail := field + " is enabled but no " + resource +
		" exists in namespace %s. Create one in the same namespace before applying."
	if r := sanitizeReason(reason); r != "" {
		detail += " " + r
	}
	return detail
}

// displayResource turns a snake_case resource name into a Title Case label.
func displayResource(resource string) string {
	words := strings.Split(resource, "_")
	for i, w := range words {
		if w != "" {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}

// sanitizeReason strips '%' so an interpolated reason cannot inject a stray
// format verb into ErrorDetail (which is later passed to fmt.Sprintf).
func sanitizeReason(reason string) string {
	return strings.TrimSpace(strings.ReplaceAll(reason, "%", ""))
}
