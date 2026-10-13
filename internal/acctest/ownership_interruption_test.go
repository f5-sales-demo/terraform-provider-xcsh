// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"testing"
)

type interruptedCleanupBackend struct {
	cancel         context.CancelFunc
	cancelAt       int
	reads, deletes int
	visible        bool
}

func (b *interruptedCleanupBackend) Read(_ context.Context, r OwnedResource) (client.Metadata, error) {
	b.reads++
	if b.cancel != nil && (b.cancelAt == 0 || b.reads == b.cancelAt) {
		b.cancel()
	}
	if b.visible {
		return client.Metadata{Name: r.Name, UID: "owned", Annotations: map[string]string{ownershipAnnotation: "runone"}}, nil
	}
	return client.Metadata{}, errOwnedAbsent
}
func (b *interruptedCleanupBackend) Delete(context.Context, OwnedResource) error {
	b.deletes++
	return nil
}

func TestInterruptedCleanupFailsEvenWhenBackendReturnsAbsence(t *testing.T) {
	for _, when := range []string{"before cleanup", "during read", "before delete", "during second pass"} {
		t.Run(when, func(t *testing.T) {
			j, err := CreateOwnershipJournal(t.TempDir()+"/journal.json", "runone")
			if err != nil {
				t.Fatal(err)
			}
			r := OwnedResource{Type: "xcsh_origin_pool", Namespace: "system", Name: "tf-acc-test-runone-pool"}
			if err := j.Intent(r); err != nil {
				t.Fatal(err)
			}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			backend := &interruptedCleanupBackend{}
			if when == "before cleanup" {
				cancel()
			} else {
				backend.cancel = cancel
			}
			if when == "during second pass" {
				backend.cancelAt = 2
			}
			backend.visible = when == "before delete"
			receipt, err := j.Cleanup(ctx, backend, 3, 0)
			if err == nil || receipt.Complete || receipt.Diagnostics["cleanup_interrupted"] == 0 {
				t.Fatal("interrupted cleanup accepted")
			}
			if backend.deletes != 0 {
				t.Fatal("cleanup deleted after interruption")
			}
		})
	}
}
