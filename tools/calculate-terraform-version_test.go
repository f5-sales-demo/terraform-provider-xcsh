// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUpdateTemplateFilesPreservesFeatureSpecificFunctionMinimum(t *testing.T) {
	root := t.TempDir()
	templates := filepath.Join(root, "templates")
	if err := os.MkdirAll(templates, 0o755); err != nil {
		t.Fatal(err)
	}
	functionText := "This function requires Terraform 1.8 or later.\n"
	indexText := "| terraform | >= 1.8  |\n"
	if err := os.WriteFile(filepath.Join(templates, "functions.md.tmpl"), []byte(functionText), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(templates, "index.md.tmpl"), []byte(indexText), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := updateTemplateFiles(root, "1.14"); err != nil {
		t.Fatal(err)
	}
	functionResult, err := os.ReadFile(filepath.Join(templates, "functions.md.tmpl"))
	if err != nil {
		t.Fatal(err)
	}
	if string(functionResult) != functionText {
		t.Fatalf("function-specific minimum changed: %q", functionResult)
	}
	indexResult, err := os.ReadFile(filepath.Join(templates, "index.md.tmpl"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(indexResult), ">= 1.14") {
		t.Fatalf("provider-wide minimum was not updated: %q", indexResult)
	}
}
