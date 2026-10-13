// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOwnershipJournalRejectsUnknownOrTrailingData(t *testing.T) {
	for _, data := range []string{
		`{"version":1,"run_id":"runone","resources":[],"credential":"SECRET_SENTINEL"}`,
		`{"version":1,"run_id":"runone","resources":[]} {}`,
	} {
		path := filepath.Join(t.TempDir(), "journal.json")
		if err := os.WriteFile(path, []byte(data), 0600); err != nil {
			t.Fatal(err)
		}
		if _, err := LoadOwnershipJournal(path, "runone"); err == nil {
			t.Fatal("malformed journal accepted")
		}
	}
}
func TestDependencyReferenceUsesExactPathComponents(t *testing.T) {
	if !containsOwnedReference("/api/object_store/namespaces/shared/stored_objects/swagger/owned/v1", "owned") {
		t.Fatal("versioned dependency missed")
	}
	if containsOwnedReference("/path/owned-other/v1", "owned") {
		t.Fatal("prefix mistaken for ownership")
	}
}
