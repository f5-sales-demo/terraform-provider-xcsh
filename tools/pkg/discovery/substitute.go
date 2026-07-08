// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package discovery holds testable helpers for tools/discover-defaults.go
// (which is a //go:build ignore program and cannot be unit-tested directly).
package discovery

import "strings"

// SubstitutePlaceholders deep-copies v, replacing "@prereq:<kind>" string values
// with names[kind] and "@prereq-ns:<kind>" with ns. It wires the created names of
// prerequisite dependencies into the spec of the resource under discovery (e.g. an
// http_loadbalancer's default_route_pools[].pool.name must point at a real
// origin_pool created first). Strings with no matching placeholder are unchanged.
func SubstitutePlaceholders(v interface{}, names map[string]string, ns string) interface{} {
	switch t := v.(type) {
	case string:
		if kind, ok := strings.CutPrefix(t, "@prereq:"); ok {
			if n, found := names[kind]; found {
				return n
			}
			return t
		}
		if _, ok := strings.CutPrefix(t, "@prereq-ns:"); ok {
			return ns
		}
		return t
	case map[string]interface{}:
		out := make(map[string]interface{}, len(t))
		for k, val := range t {
			out[k] = SubstitutePlaceholders(val, names, ns)
		}
		return out
	case []interface{}:
		out := make([]interface{}, len(t))
		for i, val := range t {
			out[i] = SubstitutePlaceholders(val, names, ns)
		}
		return out
	default:
		return v
	}
}
