// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package suppress derives the terraform-import default-suppression map from the
// defaults discovered by tools/discover-defaults.go. A discovered field is a
// server default (safe to suppress on import) when it was absent from the create
// request but present in the response as either an empty marker block or a bool
// at its false zero-default. See issue #1006.
package suppress

import (
	"sort"
	"strings"

	defaultspkg "github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/defaults"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/naming"
)

// These aliases keep suppression derivation on the same persisted contract as
// discovery. In particular, resources are a typed array rather than an object
// whose keys duplicate resource identity.
type FieldDefault = defaultspkg.DefaultValue
type ResourceResult = defaultspkg.ResourceEntry
type Database = defaultspkg.APIDefaultsFile

// leaf returns the final dot-separated segment of a field path (the member name).
func leaf(path string) string {
	parts := strings.Split(path, ".")
	return parts[len(parts)-1]
}

// isServerDefaultMember reports whether a discovered field is a server-default
// oneof member that is safe to suppress on import.
func isServerDefaultMember(fd FieldDefault) bool {
	if fd.IsMarkerBlock {
		return true
	}
	if fd.Type == "bool" {
		if b, ok := fd.DefaultValue.(bool); ok && !b {
			return true
		}
	}
	return false
}

// Derive builds a resource(title-case) -> sorted member names map from discovered
// defaults. Only successfully-discovered resources contribute.
func Derive(db Database) map[string][]string {
	acc := map[string]map[string]bool{}
	for _, res := range db.Resources {
		if res.Status != "discovered" {
			continue
		}
		rc := naming.ToResourceTypeName(res.ResourceName)
		for _, fd := range res.Defaults {
			if !isServerDefaultMember(fd) {
				continue
			}
			if acc[rc] == nil {
				acc[rc] = map[string]bool{}
			}
			acc[rc][leaf(fd.Path)] = true
		}
	}
	out := map[string][]string{}
	for rc, members := range acc {
		list := make([]string, 0, len(members))
		for m := range members {
			list = append(list, m)
		}
		sort.Strings(list)
		out[rc] = list
	}
	return out
}

// Merge unions derived members into an existing map (from the JSON data file),
// preserving hand-seeded entries. Returns a fresh sorted map.
func Merge(existing, derived map[string][]string) map[string][]string {
	acc := map[string]map[string]bool{}
	add := func(rc string, members []string) {
		if acc[rc] == nil {
			acc[rc] = map[string]bool{}
		}
		for _, m := range members {
			acc[rc][m] = true
		}
	}
	for rc, m := range existing {
		add(rc, m)
	}
	for rc, m := range derived {
		add(rc, m)
	}
	out := map[string][]string{}
	for rc, members := range acc {
		list := make([]string, 0, len(members))
		for m := range members {
			list = append(list, m)
		}
		sort.Strings(list)
		out[rc] = list
	}
	return out
}
