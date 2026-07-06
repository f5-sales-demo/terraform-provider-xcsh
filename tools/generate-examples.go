// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore

// Command generate-examples regenerates the Terraform examples.
//
// Examples are generated schema-driven (directly from the TerraformAttribute tree) by
// tools/generate-all-schemas.go — the single source of truth — so they can never drift
// from the provider schema. This command is retained as the historical entrypoint that
// several CI jobs invoke; it delegates to generate-all-schemas so every caller produces
// the same fresh, schema-valid examples (and provider code).
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	specDir := os.Getenv("XCSH_SPEC_DIR")
	if specDir == "" {
		specDir = "docs/specifications/api"
	}
	fmt.Println("generate-examples: delegating to tools/generate-all-schemas.go (single source of truth for schema-driven examples)")
	cmd := exec.Command("go", "run", "tools/generate-all-schemas.go", "--spec-dir="+specDir) //nolint:gosec // fixed args
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "generate-examples: delegation to generate-all-schemas failed: %v\n", err)
		os.Exit(1)
	}
}
