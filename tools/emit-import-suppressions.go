// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore
// +build ignore

// emit-import-suppressions derives the terraform-import default-suppression data
// file (tools/import-default-suppressions.json, consumed by the code generator)
// from the defaults discovered by discover-defaults.go (tools/api-defaults.json).
// It unions newly-derived members over the existing file so hand-seeded entries
// are preserved. Run: go run tools/emit-import-suppressions.go
//
// See issue #1006.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/suppress"
)

const defaultComment = "Per-resource (title-case model prefix) server-default oneof members suppressed on the terraform import path (consumed by the code generator). Auto-derived by tools/emit-import-suppressions.go from tools/api-defaults.json; hand-seeded fallback in tools/pkg/codegen/import_suppressions.go. See issue #1006."

func main() {
	inDB := flag.String("from-db", "tools/api-defaults.json", "discovered defaults database")
	outFile := flag.String("out", "tools/import-default-suppressions.json", "suppression data file to write")
	flag.Parse()

	// Load existing suppression file (preserve hand-seeded entries + comment).
	existing := map[string][]string{}
	comment := defaultComment
	if data, err := os.ReadFile(*outFile); err == nil {
		var raw map[string]json.RawMessage
		if json.Unmarshal(data, &raw) == nil {
			for k, v := range raw {
				if k == "_comment" {
					_ = json.Unmarshal(v, &comment)
					continue
				}
				var members []string
				if json.Unmarshal(v, &members) == nil {
					existing[k] = members
				}
			}
		}
	}

	// Load discovered defaults.
	data, err := os.ReadFile(*inDB)
	if err != nil {
		fmt.Fprintf(os.Stderr, "emit-import-suppressions: %v\n", err)
		os.Exit(1)
	}
	var db suppress.Database
	if err := json.Unmarshal(data, &db); err != nil {
		fmt.Fprintf(os.Stderr, "emit-import-suppressions: parse %s: %v\n", *inDB, err)
		os.Exit(1)
	}

	derived := suppress.Derive(db)
	merged := suppress.Merge(existing, derived)

	// Write JSON with the comment first (maps marshal with sorted keys; "_comment"
	// sorts before capitalized resource names).
	out := map[string]interface{}{"_comment": comment}
	for rc, members := range merged {
		out[rc] = members
	}
	b, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "emit-import-suppressions: marshal: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(*outFile, append(b, '\n'), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "emit-import-suppressions: write %s: %v\n", *outFile, err)
		os.Exit(1)
	}
	fmt.Printf("emit-import-suppressions: %d resources in %s (%d derived from %s)\n", len(merged), *outFile, len(derived), *inDB)
}
