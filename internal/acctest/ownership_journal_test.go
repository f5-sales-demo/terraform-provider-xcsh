// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"errors"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"path/filepath"
	"testing"
)

type journalBackend struct {
	present   map[string]string
	reads     int
	deletes   []string
	readError error
	pending   int
}

func (b *journalBackend) Read(_ context.Context, r OwnedResource) (client.Metadata, error) {
	b.reads++
	if b.readError != nil {
		return client.Metadata{}, b.readError
	}
	if uid, ok := b.present[r.Name]; ok {
		return client.Metadata{Name: r.Name, UID: uid, Annotations: map[string]string{ownershipAnnotation: "runone"}}, nil
	}
	return client.Metadata{}, errOwnedAbsent
}
func (b *journalBackend) Delete(_ context.Context, r OwnedResource) error {
	b.deletes = append(b.deletes, r.Name)
	if b.pending > 0 {
		b.pending--
		return nil
	}
	delete(b.present, r.Name)
	return nil
}
func TestOwnershipJournalFailsClosed(t *testing.T) {
	path := filepath.Join(t.TempDir(), "ownership.json")
	if _, err := LoadOwnershipJournal(path, "runone"); err == nil {
		t.Fatal("missing journal accepted")
	}
	j, err := CreateOwnershipJournal(path, "runone")
	if err != nil {
		t.Fatal(err)
	}
	if _, err = LoadOwnershipJournal(path, "runtwo"); err == nil {
		t.Fatal("other run accepted")
	}
	r := OwnedResource{Type: "xcsh_origin_pool", Name: "tf-acc-test-runone-pool", Namespace: "system"}
	if err = j.Intent(r); err != nil {
		t.Fatal(err)
	}
	b := &journalBackend{present: map[string]string{r.Name: "owned-uid"}}
	receipt, err := j.Cleanup(context.Background(), b, 3, 0)
	if err != nil || !receipt.Complete || receipt.SecondPassRemaining != 0 || len(b.deletes) != 1 {
		t.Fatalf("interrupted write not recovered: %#v %v", receipt, err)
	}
	receipt, err = j.Cleanup(context.Background(), b, 3, 0)
	if err != nil || !receipt.Complete || len(b.deletes) != 1 {
		t.Fatal("cleanup not idempotent")
	}
}
func TestOwnershipJournalPreservesReplacementAndUnknownReads(t *testing.T) {
	for _, readError := range []error{nil, errors.New("404 in authorization error")} {
		j, err := CreateOwnershipJournal(filepath.Join(t.TempDir(), "journal.json"), "runone")
		if err != nil {
			t.Fatal(err)
		}
		r := OwnedResource{Type: "xcsh_origin_pool", Name: "tf-acc-test-runone-pool", Namespace: "system"}
		if err = j.Intent(r); err != nil {
			t.Fatal(err)
		}
		if err = j.Confirm(r, "original"); err != nil {
			t.Fatal(err)
		}
		b := &journalBackend{present: map[string]string{r.Name: "replacement"}, readError: readError}
		receipt, err := j.Cleanup(context.Background(), b, 3, 0)
		if err == nil || receipt.Complete || len(b.deletes) != 0 {
			t.Fatal("unowned replacement or ambiguous read deleted")
		}
	}
}
func TestOwnershipJournalReverseDependenciesAndAsyncDeletion(t *testing.T) {
	j, err := CreateOwnershipJournal(filepath.Join(t.TempDir(), "journal.json"), "runone")
	if err != nil {
		t.Fatal(err)
	}
	parent := OwnedResource{Type: "xcsh_namespace", Name: "tf-acc-test-runone-parent"}
	child := OwnedResource{Type: "xcsh_origin_pool", Name: "tf-acc-test-runone-child", Namespace: parent.Name, Dependencies: []string{ownedKey(parent)}}
	for _, r := range []OwnedResource{parent, child} {
		if err = j.Intent(r); err != nil {
			t.Fatal(err)
		}
		if err = j.Confirm(r, r.Name); err != nil {
			t.Fatal(err)
		}
	}
	b := &journalBackend{present: map[string]string{parent.Name: parent.Name, child.Name: child.Name}, pending: 1}
	receipt, err := j.Cleanup(context.Background(), b, 4, 0)
	if err != nil || !receipt.Complete || b.deletes[0] != child.Name || b.deletes[len(b.deletes)-1] != parent.Name {
		t.Fatalf("dependency/async cleanup: %#v %v %v", receipt, err, b.deletes)
	}
}
func TestMissingOwnedVerifierFails(t *testing.T) {
	b := apiOwnershipBackend{}
	_, err := b.Read(context.Background(), OwnedResource{Type: "xcsh_" + "not_registered"})
	if err == nil {
		t.Fatal("missing verifier accepted")
	}
}
