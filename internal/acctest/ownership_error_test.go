// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"errors"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"testing"
)

type transientBadReadBackend struct{ reads int }

func (b *transientBadReadBackend) Read(context.Context, OwnedResource) (client.Metadata, error) {
	b.reads++
	if b.reads == 1 {
		return client.Metadata{}, errors.New("authorization unavailable")
	}
	return client.Metadata{}, errOwnedAbsent
}
func (b *transientBadReadBackend) Delete(context.Context, OwnedResource) error { return nil }
func TestOwnershipCleanupRetainsAuthorizationFailure(t *testing.T) {
	j, err := CreateOwnershipJournal(t.TempDir()+"/journal.json", "runone")
	if err != nil {
		t.Fatal(err)
	}
	r := OwnedResource{Type: "xcsh_origin_pool", Name: "tf-acc-test-runone-pool", Namespace: "system"}
	if err = j.Intent(r); err != nil {
		t.Fatal(err)
	}
	receipt, err := j.Cleanup(context.Background(), &transientBadReadBackend{}, 3, 0)
	if err == nil || receipt.Complete || receipt.Failures == 0 {
		t.Fatal("authorization failure masked by subsequent absence")
	}
}
