// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"errors"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"testing"
)

func TestQuotaProbeReleaseVerifiesAbsenceWithoutTouchingOtherRun(t *testing.T) {
	for _, tc := range []struct {
		name                         string
		failDelete, failRead, remain bool
	}{
		{"deleted", false, false, false}, {"uncertain delete", true, false, false}, {"uncertain read", false, true, false}, {"asynchronous remains", false, false, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			b := &quotaReleaseBackend{present: true, failDelete: tc.failDelete, failRead: tc.failRead, remain: tc.remain}
			err := releaseQuotaProbe(context.Background(), b, OwnedResource{Type: "xcsh_crl", Name: "owned"}, 2, 0)
			want := !tc.failDelete && !tc.failRead && !tc.remain
			if (err == nil) != want {
				t.Fatal("unverified quota probe release accepted")
			}
			if b.deleted != "owned" {
				t.Fatal("quota release selected another identity")
			}
		})
	}
}

type quotaReleaseBackend struct {
	present, failDelete, failRead, remain bool
	deleted                               string
}

func (b *quotaReleaseBackend) Delete(_ context.Context, r OwnedResource) error {
	b.deleted = r.Name
	if b.failDelete {
		return errors.New("uncertain")
	}
	if !b.remain {
		b.present = false
	}
	return nil
}
func (b *quotaReleaseBackend) Read(context.Context, OwnedResource) (client.Metadata, error) {
	if b.failRead {
		return client.Metadata{}, errors.New("uncertain")
	}
	if !b.present {
		return client.Metadata{}, errOwnedAbsent
	}
	return client.Metadata{Name: "owned", UID: "immutable"}, nil
}
