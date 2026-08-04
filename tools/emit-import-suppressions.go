// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore
// +build ignore

// emit-import-suppressions updates the canonical Terraform import suppression
// measurements from a discovered defaults database.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/suppress"
)

func main() {
	inDB := flag.String("from-db", "tools/api-defaults.json", "discovered defaults database")
	outFile := flag.String("out", "tools/import-default-suppressions.json", "suppression data file to write")
	flag.Parse()
	if err := suppress.EmitImportSuppressions(*inDB, *outFile); err != nil {
		fmt.Fprintf(os.Stderr, "emit-import-suppressions: %v\n", err)
		os.Exit(1)
	}
}
