// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"os"
	"strings"
	"testing"
)

func TestOwnershipJournalDoesNotStoreCredentialUID(t *testing.T) {
	path := t.TempDir() + "/journal.json"
	j, err := CreateOwnershipJournal(path, "runone")
	if err != nil {
		t.Fatal(err)
	}
	r := OwnedResource{Type: "xcsh_token", Name: "tf-acc-test-runone-token", Namespace: "system"}
	if err = j.Intent(r); err != nil {
		t.Fatal(err)
	}
	if err = j.Confirm(r, "SECRET_SENTINEL_TOKEN_UID"); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "SECRET_SENTINEL") {
		t.Fatal("credential UID stored")
	}
}
