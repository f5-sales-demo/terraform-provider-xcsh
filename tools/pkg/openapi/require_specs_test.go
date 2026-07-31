// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// RequireSpecs exists because a missing spec directory used to be a *note*, not an
// error. transform-docs.go printed "No OpenAPI specs found" and carried on, emitting
// documentation stripped of every spec-derived enrichment — per-domain API reference
// links, dependency notes and danger-level callouts. Measured on a spec-less checkout
// of the commit that added the CI diff assertion: 256 files changed, 1548 deletions,
// exit code 0.
//
// That is the defect this guard closes: a missing INPUT must not look like a
// successful RUN.
func TestRequireSpecsRejectsMissingDirectory(t *testing.T) {
	err := RequireSpecs(filepath.Join(t.TempDir(), "definitely-not-here"))
	if err == nil {
		t.Fatal("RequireSpecs accepted a directory that does not exist; a missing spec bundle must be an error, not a note")
	}
	if !strings.Contains(err.Error(), "download-specs") {
		t.Errorf("error must tell the caller how to recover, got: %v", err)
	}
}

// An empty directory is the shape CI actually produced: the checkout creates nothing,
// and `mkdir -p` in an unrelated step can leave the path present but bare. Presence of
// the path is not presence of the specs.
func TestRequireSpecsRejectsEmptyDirectory(t *testing.T) {
	if err := RequireSpecs(t.TempDir()); err == nil {
		t.Fatal("RequireSpecs accepted an empty directory; index.json and domains/ are both required")
	}
}

// index.json without domains/ is a partial unzip — a truncated download that would
// otherwise parse far enough to look usable.
func TestRequireSpecsRejectsPartialBundle(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "index.json"), []byte(`{"specifications":[]}`), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := RequireSpecs(dir); err == nil {
		t.Fatal("RequireSpecs accepted index.json with no domains/ directory")
	}
}

func TestRequireSpecsAcceptsCompleteBundle(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "index.json"), []byte(`{"specifications":[]}`), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(dir, "domains"), 0o750); err != nil {
		t.Fatal(err)
	}
	if err := RequireSpecs(dir); err != nil {
		t.Fatalf("RequireSpecs rejected a complete v2 bundle: %v", err)
	}
}
