// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"testing"
)

func TestOwnershipResetAllowsOnlyConfirmedNewLifecycle(t *testing.T) {
	j, err := CreateOwnershipJournal(t.TempDir()+"/journal.json", "runone")
	if err != nil {
		t.Fatal(err)
	}
	r := OwnedResource{Type: "xcsh_origin_pool", Name: "tf-acc-test-runone-pool", Namespace: "system"}
	if err = j.Intent(r); err != nil {
		t.Fatal(err)
	}
	if err = j.Confirm(r, "first"); err != nil {
		t.Fatal(err)
	}
	if err = j.Confirm(r, "replacement"); err == nil {
		t.Fatal("replacement UID accepted without absence")
	}
	if err = j.ResetAfterVerifiedAbsence(r); err != nil {
		t.Fatal(err)
	}
	if err = j.Confirm(r, "replacement"); err != nil {
		t.Fatal(err)
	}
	b := &journalBackend{present: map[string]string{r.Name: "replacement"}}
	receipt, err := j.Cleanup(context.Background(), b, 3, 0)
	if err != nil || !receipt.Complete {
		t.Fatalf("new lifecycle cleanup failed %#v %v", receipt, err)
	}
}
