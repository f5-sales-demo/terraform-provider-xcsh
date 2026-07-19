// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

// Meaningful-zero int64 leaves: fields whose value 0 is a real, user-intended value the API
// stores and returns — NOT the "unset / server default" sentinel that most int64 fields use.
//
// The generated int64 read guard is `if v, ok := m[k].(float64); ok && v != 0` — the `v != 0`
// clause treats a returned 0 as absent and leaves the attribute null. That is correct for the
// common case (the API omits or zero-fills a field it considers unset), but wrong for fields
// where 0 carries meaning: the API returns 0, the read drops it, and a config that set 0 drifts
// ("was 0, now null" on apply / round-trip import). See #1129: waf_exclusion
// exclude_signature_contexts[].signature_id, where 0 means "exclude ALL signatures for the
// context" (per the schema) and is a legitimate value in the range {0} ∪ [200000001,299999999].
//
// Listed here per resource (title-case model prefix) by leaf json name, matched at any depth
// (mirrors isImportDefaultSuppressed). For these, renderUnmarshalScalarChild drops the `v != 0`
// clause so a returned 0 is read back faithfully. Keep this list tight: only add a leaf a live
// round-trip proves the API returns 0 for as a meaningful value.
var meaningfulZeroInt64Seed = map[string][]string{
	// #1129: signature_id 0 = "all signatures excluded for the context". Appears on the LB
	// inline waf_exclusion and the standalone xcsh_waf_exclusion_policy; matched by leaf name at
	// any depth, so one entry per resource covers every nesting.
	"HTTPLoadBalancer":   {"signature_id"},
	"WAFExclusionPolicy": {"signature_id"},
}

// isMeaningfulZeroInt64 reports whether the given int64 leaf of the given resource must be read
// back even when the API returns 0 (i.e. the `v != 0` guard must be dropped).
func isMeaningfulZeroInt64(resourceTitleCase, jsonName string) bool {
	for _, leaf := range meaningfulZeroInt64Seed[resourceTitleCase] {
		if leaf == jsonName {
			return true
		}
	}
	return false
}
