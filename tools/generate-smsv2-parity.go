// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore

// Command generate-smsv2-parity creates the checked-in exhaustive provider
// compatibility matrix from the immutable legacy and current API manifests.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/parity"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "usage: generate-smsv2-parity LEGACY_JSON CURRENT_JSON OUTPUT_JSON")
		os.Exit(2)
	}
	legacy, err := parity.LoadLegacy(os.Args[1])
	if err != nil {
		fail(err)
	}
	current, err := parity.LoadCurrent(os.Args[2])
	if err != nil {
		fail(err)
	}
	matrix, err := parity.BuildSMSv2Matrix(legacy, current)
	if err != nil {
		if matrix != nil {
			for _, path := range matrix.Unclassified {
				fmt.Fprintf(os.Stderr, "unclassified: %s\n", path)
			}
		}
		fail(err)
	}
	data, err := json.MarshalIndent(matrix, "", "  ")
	if err != nil {
		fail(err)
	}
	if err := os.WriteFile(os.Args[3], append(data, '\n'), 0o644); err != nil {
		fail(err)
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
