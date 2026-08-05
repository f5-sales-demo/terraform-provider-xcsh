// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore
// +build ignore

// normalize-network-examples rewrites hand-maintained network examples to use
// RFC 5737 documentation addresses.
package main

import (
	"fmt"
	"os"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/docfmt"
)

var defaultNetworkExamplePaths = []string{
	"internal/mocks/fixtures.go",
	"internal/provider/origin_pool_resource_test.go",
	"tools/discover-defaults.go",
	"tools/generate-datasource-tests.go",
	"tools/pkg/suppress/diff_test.go",
}

func main() {
	paths := os.Args[1:]
	if len(paths) == 0 {
		paths = defaultNetworkExamplePaths
	}

	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read network example source: %v\n", err)
			os.Exit(1)
		}
		normalized := docfmt.CanonicalizeNetworkLiterals(string(data))
		if normalized == string(data) {
			continue
		}
		if err := os.WriteFile(path, []byte(normalized), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "write network example source: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Normalized network examples: %s\n", path)
	}
}
