// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeFiles(t *testing.T, dir string, names ...string) {
	t.Helper()
	for _, name := range names {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("x"), 0o600); err != nil {
			t.Fatalf("seeding %s: %v", name, err)
		}
	}
}

func TestPruneOrphanFilesRemovesOnlyUnkeptNames(t *testing.T) {
	dir := t.TempDir()
	writeFiles(t, dir, "apm.txt", "log_receiver.txt", "alert_policy.txt", "notes.md")

	removed, err := PruneOrphanFiles(dir, ".txt", map[string]bool{
		"log_receiver": true,
		"alert_policy": true,
	})
	if err != nil {
		t.Fatalf("PruneOrphanFiles: %v", err)
	}

	if len(removed) != 1 || filepath.Base(removed[0]) != "apm.txt" {
		t.Fatalf("removed = %v, want exactly [apm.txt]", removed)
	}
	for _, keptName := range []string{"log_receiver.txt", "alert_policy.txt", "notes.md"} {
		if _, err := os.Stat(filepath.Join(dir, keptName)); err != nil {
			t.Errorf("%s should have survived: %v", keptName, err)
		}
	}
	if _, err := os.Stat(filepath.Join(dir, "apm.txt")); !os.IsNotExist(err) {
		t.Errorf("apm.txt should be gone, stat err = %v", err)
	}
}

// An empty keep set means the caller found no resources at all — a failed fetch,
// a bad glob, a generator that bailed early. Pruning on that signal would delete
// the whole directory, so it must be an error instead.
func TestPruneOrphanFilesRefusesEmptyKeepSet(t *testing.T) {
	dir := t.TempDir()
	writeFiles(t, dir, "apm.txt", "log_receiver.txt")

	removed, err := PruneOrphanFiles(dir, ".txt", nil)
	if err == nil {
		t.Fatal("expected an error for an empty keep set, got nil")
	}
	if len(removed) != 0 {
		t.Errorf("removed = %v, want nothing", removed)
	}
	for _, name := range []string{"apm.txt", "log_receiver.txt"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Errorf("%s must survive a refused prune: %v", name, err)
		}
	}
}

func TestPruneOrphanFilesRejectsBadExtension(t *testing.T) {
	for _, ext := range []string{"", "txt"} {
		if _, err := PruneOrphanFiles(t.TempDir(), ext, map[string]bool{"a": true}); err == nil {
			t.Errorf("extension %q should have been rejected", ext)
		}
	}
}

func TestPruneOrphanFilesIsIdempotent(t *testing.T) {
	dir := t.TempDir()
	writeFiles(t, dir, "apm.txt", "log_receiver.txt")
	keep := map[string]bool{"log_receiver": true}

	first, err := PruneOrphanFiles(dir, ".txt", keep)
	if err != nil {
		t.Fatalf("first prune: %v", err)
	}
	if len(first) != 1 {
		t.Fatalf("first prune removed %v, want one file", first)
	}

	second, err := PruneOrphanFiles(dir, ".txt", keep)
	if err != nil {
		t.Fatalf("second prune: %v", err)
	}
	if len(second) != 0 {
		t.Errorf("second prune removed %v, want nothing — pruning must be idempotent", second)
	}
}

func TestPruneOrphanFilesReportsRemovalFailure(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("running as root; a read-only directory would not block os.Remove")
	}

	parent := t.TempDir()
	dir := filepath.Join(parent, "resources")
	if err := os.Mkdir(dir, 0o700); err != nil {
		t.Fatalf("creating %s: %v", dir, err)
	}
	writeFiles(t, dir, "apm.txt")

	// Removing a directory entry needs write permission on the DIRECTORY, so
	// clearing it makes os.Remove fail while the file itself is untouched.
	if err := os.Chmod(dir, 0o500); err != nil {
		t.Fatalf("chmod %s: %v", dir, err)
	}
	t.Cleanup(func() { _ = os.Chmod(dir, 0o700) })

	_, err := PruneOrphanFiles(dir, ".txt", map[string]bool{"log_receiver": true})
	if err == nil {
		t.Fatal("expected a removal failure to be reported, got nil")
	}
	if !strings.Contains(err.Error(), "apm.txt") {
		t.Errorf("error %q should name the file it could not remove", err)
	}
}
