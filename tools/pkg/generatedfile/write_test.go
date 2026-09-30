// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package generatedfile

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestUnchangedOutputPreservesFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "output.go")
	content := []byte("package fixture\n")
	if err := os.WriteFile(path, content, 0600); err != nil {
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
	if err := WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}
	after, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if !os.SameFile(before, after) || !after.ModTime().Equal(before.ModTime()) || after.Mode().Perm() != before.Mode().Perm() {
		t.Fatalf("unchanged output metadata differs: before=%v after=%v", before, after)
	}
}

func TestNewAndChangedOutput(t *testing.T) {
	path := filepath.Join(t.TempDir(), "output.go")
	for _, value := range []string{"before\n", "after\n", ""} {
		if err := WriteFile(path, []byte(value), 0644); err != nil {
			t.Fatal(err)
		}
		got, err := os.ReadFile(path)
		if err != nil || string(got) != value {
			t.Fatalf("output=%q err=%v", got, err)
		}
	}
}

func TestInvalidOutputCannotSucceed(t *testing.T) {
	dir := t.TempDir()
	for _, path := range []string{dir, filepath.Join(dir, "missing", "output.go")} {
		if err := WriteFile(path, []byte("content"), 0644); err == nil {
			t.Fatalf("invalid output %q succeeded", path)
		}
	}
}
