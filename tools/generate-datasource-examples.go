// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore

// Command generate-datasource-examples is retained for pipeline compatibility only.
// Data-source examples are generated schema-driven by tools/generate-all-schemas.go
// (make generate), so this command performs no work.
package main

import "fmt"

func main() {
	fmt.Println("generate-datasource-examples: no-op. Data-source examples are generated " +
		"by tools/generate-all-schemas.go (make generate).")
}
