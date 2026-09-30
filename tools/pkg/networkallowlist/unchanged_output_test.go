// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package networkallowlist

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestWriteGoPreservesIdenticalOutputAndReplacesChangedOutput(t *testing.T) {
	openAPI, pin := fixture(t, validManifest())
	artifact, err := Extract(openAPI, pin, "v8.0.2")
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "nested", "allowlist.go")
	if err := WriteGo(path, artifact); err != nil {
		t.Fatal(err)
	}
	fixed := time.Unix(1000000000, 0)
	if err := os.Chtimes(path, fixed, fixed); err != nil {
		t.Fatal(err)
	}
	before, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := WriteGo(path, artifact); err != nil {
		t.Fatal(err)
	}
	after, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if !os.SameFile(before, after) || !before.ModTime().Equal(after.ModTime()) {
		t.Fatal("unchanged allowlist was replaced")
	}
	artifact.GeneratedAt = "2026-01-01T00:00:00Z"
	if err := WriteGo(path, artifact); err != nil {
		t.Fatal(err)
	}
	changed, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if os.SameFile(before, changed) {
		t.Fatal("changed allowlist lost atomic replacement")
	}
	if changed.Mode().Perm() != 0644 {
		t.Fatalf("changed output mode=%v", changed.Mode())
	}
}
