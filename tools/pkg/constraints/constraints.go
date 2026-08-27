// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package constraints

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var int64RangeSpanPattern = regexp.MustCompile(`^(-?[0-9]+)(?:-(-?[0-9]+))?$`)

// Int64RangeSpan is one inclusive interval parsed from an upstream numeric
// range-set rule.
type Int64RangeSpan struct {
	Minimum int64
	Maximum int64
}

// Parsed represents extracted x-f5xc-constraints data.
type Parsed struct {
	MinLength int
	MaxLength int
	Pattern   string
	Format    string
	MinItems  int
	MaxItems  int
	Minimum   int
	Maximum   int

	// HasMinimum/HasMaximum record whether the minimum/maximum key was present,
	// independent of value, so a legitimate minimum:0 (or maximum:0) is detectable.
	HasMinimum bool
	HasMaximum bool
}

// Parse extracts constraint data from x-f5xc-constraints map.
// Returns nil for nil input, low-confidence, or non-deterministic constraints.
// Requires deterministic at top level or metadata confidence >= 0.9.
// Only uses constraints where deterministic==true OR metadata.confidence >= 0.9.
func Parse(raw map[string]interface{}) *Parsed {
	if raw == nil {
		return nil
	}

	deterministic := false
	confidence := 0.0

	// Check top-level deterministic (actual spec format since v2.1.80+)
	if d, ok := raw["deterministic"].(bool); ok {
		deterministic = d
	}

	// Metadata confidence is part of the current enriched-spec contract.
	if meta, ok := raw["metadata"].(map[string]interface{}); ok {
		if c, ok := meta["confidence"].(float64); ok {
			confidence = c
		}
	}

	if !deterministic && confidence < 0.9 {
		return nil
	}

	p := &Parsed{}

	// String constraints
	if v, ok := raw["minLength"].(float64); ok {
		p.MinLength = int(v)
	}
	if v, ok := raw["maxLength"].(float64); ok {
		p.MaxLength = int(v)
	}
	if v, ok := raw["pattern"].(string); ok {
		// Skip discovery-inferred patterns (e.g., format:uri → ^(https?|ftp)://)
		// which may reject vendor-specific schemes like string:///
		source := ""
		if meta, ok := raw["metadata"].(map[string]interface{}); ok {
			if s, ok := meta["source"].(string); ok {
				source = s
			}
		}
		if source != "discovery" {
			if _, err := regexp.Compile(v); err == nil {
				p.Pattern = v
			}
		}
	}

	// Format label (e.g. ipv4, ipv6, ip, cidr, mac-address) drives string
	// format validators in codegen.
	if v, ok := raw["format"].(string); ok {
		p.Format = v
	}

	// List/array constraints
	if v, ok := raw["minItems"].(float64); ok {
		p.MinItems = int(v)
	}
	if v, ok := raw["maxItems"].(float64); ok {
		p.MaxItems = int(v)
	}

	// Numeric constraints
	if v, ok := raw["minimum"].(float64); ok {
		p.Minimum = int(v)
		p.HasMinimum = true
	}
	if v, ok := raw["maximum"].(float64); ok {
		p.Maximum = int(v)
		p.HasMaximum = true
	}

	return p
}

// ParseInt64RangeSpans parses any ves.io numeric `.ranges` validation rule into
// a canonical ordered set of inclusive intervals. Overlapping and adjacent
// intervals are merged; malformed or descending intervals fail closed.
func ParseInt64RangeSpans(rules map[string]string) ([]Int64RangeSpan, error) {
	keys := make([]string, 0)
	for key := range rules {
		if strings.HasSuffix(key, ".ranges") {
			keys = append(keys, key)
		}
	}
	if len(keys) == 0 {
		return nil, nil
	}
	sort.Strings(keys)

	spans := make([]Int64RangeSpan, 0)
	for _, key := range keys {
		value := strings.TrimSpace(rules[key])
		if value == "" {
			return nil, fmt.Errorf("%s is empty", key)
		}
		for _, rawSpan := range strings.Split(value, ",") {
			rawSpan = strings.TrimSpace(rawSpan)
			matches := int64RangeSpanPattern.FindStringSubmatch(rawSpan)
			if matches == nil {
				return nil, fmt.Errorf("%s contains malformed range %q", key, rawSpan)
			}
			minimum, err := strconv.ParseInt(matches[1], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("%s minimum %q: %w", key, matches[1], err)
			}
			maximum := minimum
			if matches[2] != "" {
				maximum, err = strconv.ParseInt(matches[2], 10, 64)
				if err != nil {
					return nil, fmt.Errorf("%s maximum %q: %w", key, matches[2], err)
				}
			}
			if maximum < minimum {
				return nil, fmt.Errorf("%s contains descending range %q", key, rawSpan)
			}
			spans = append(spans, Int64RangeSpan{Minimum: minimum, Maximum: maximum})
		}
	}

	sort.Slice(spans, func(i, j int) bool {
		if spans[i].Minimum == spans[j].Minimum {
			return spans[i].Maximum < spans[j].Maximum
		}
		return spans[i].Minimum < spans[j].Minimum
	})
	merged := make([]Int64RangeSpan, 0, len(spans))
	for _, span := range spans {
		if len(merged) == 0 {
			merged = append(merged, span)
			continue
		}
		last := &merged[len(merged)-1]
		adjacent := last.Maximum != int64(^uint64(0)>>1) && span.Minimum == last.Maximum+1
		if span.Minimum <= last.Maximum || adjacent {
			if span.Maximum > last.Maximum {
				last.Maximum = span.Maximum
			}
			continue
		}
		merged = append(merged, span)
	}
	return merged, nil
}
