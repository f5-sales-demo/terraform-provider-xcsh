// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
)

func TestOwnershipFailedChildPreservesParent(t *testing.T) {
	j, err := CreateOwnershipJournal(t.TempDir()+"/journal.json", "runone")
	if err != nil {
		t.Fatal(err)
	}
	parent := OwnedResource{Type: "xcsh_namespace", Name: "tf-acc-test-runone-ns"}
	child := OwnedResource{Type: "xcsh_origin_pool", Name: "tf-acc-test-runone-pool", Namespace: parent.Name, Dependencies: []string{ownedKey(parent)}}
	for _, r := range []OwnedResource{parent, child} {
		if err = j.Intent(r); err != nil {
			t.Fatal(err)
		}
		if err = j.Confirm(r, "uid"); err != nil {
			t.Fatal(err)
		}
	}
	b := &journalBackend{present: map[string]string{parent.Name: "uid", child.Name: "replacement"}}
	receipt, err := j.Cleanup(context.Background(), b, 3, 0)
	if err == nil || receipt.Complete || len(b.deletes) != 0 {
		t.Fatal("parent deleted while child ownership uncertain")
	}
}
func TestOwnershipReceiptContainsNoIdentity(t *testing.T) {
	b, _ := json.Marshal(CleanupReceipt{Version: 1, Complete: true, Owned: 1, ConfirmedAbsent: 1, JournalSHA256: strings.Repeat("a", 64)})
	if strings.Contains(string(b), "namespace") || strings.Contains(string(b), "name") || strings.Contains(string(b), "uid") {
		t.Fatal("receipt identity leak")
	}
}
