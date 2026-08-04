// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"regexp"
	"strconv"
)

var (
	generatedBetweenRe = regexp.MustCompile(`int64validator\.Between\(\s*(-?\d+)\s*,\s*(-?\d+)\s*\)`)
	generatedAtLeastRe = regexp.MustCompile(`int64validator\.AtLeast\(\s*(-?\d+)\s*\)`)
	generatedAtMostRe  = regexp.MustCompile(`int64validator\.AtMost\(\s*(-?\d+)\s*\)`)
)

// ParseGeneratedInt64Bounds extracts the numeric range from a generated
// schema.Int64Attribute body. The generator emits only integer literals here,
// so a regexp match is sufficient and keeps example generation independent of
// the OpenAPI bundle.
func ParseGeneratedInt64Bounds(body string) (minimum, maximum int, hasMinimum, hasMaximum bool) {
	if match := generatedBetweenRe.FindStringSubmatch(body); match != nil {
		return mustParseGeneratedBound(match[1]), mustParseGeneratedBound(match[2]), true, true
	}

	if match := generatedAtLeastRe.FindStringSubmatch(body); match != nil {
		minimum, hasMinimum = mustParseGeneratedBound(match[1]), true
	}
	if match := generatedAtMostRe.FindStringSubmatch(body); match != nil {
		maximum, hasMaximum = mustParseGeneratedBound(match[1]), true
	}
	return minimum, maximum, hasMinimum, hasMaximum
}

func mustParseGeneratedBound(value string) int {
	bound, err := strconv.Atoi(value)
	if err != nil {
		panic("generated int64 validator bound is outside the platform int range: " + value)
	}
	return bound
}
