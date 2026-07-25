// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package discovery

import (
	"errors"
	"testing"
)

// A cleanup DELETE that fails must be retried and, when every attempt fails, must surface
// the error — discover-defaults previously discarded it, silently stranding a tf-discover-*
// probe object in the shared "system" namespace next to the live demo sites.
func TestDeleteWithRetry(t *testing.T) {
	boom := errors.New("500 internal error")

	tests := []struct {
		name         string
		attempts     int
		failuresLeft int // how many calls fail before succeeding; >attempts = never succeeds
		wantCalls    int
		wantErr      bool
	}{
		{name: "succeeds first try", attempts: 3, failuresLeft: 0, wantCalls: 1},
		{name: "succeeds after transient failure", attempts: 3, failuresLeft: 2, wantCalls: 3},
		{name: "exhausts attempts and reports the error", attempts: 3, failuresLeft: 9, wantCalls: 3, wantErr: true},
		{name: "attempts below one still tries once", attempts: 0, failuresLeft: 9, wantCalls: 1, wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			calls, left := 0, tc.failuresLeft
			err := DeleteWithRetry(func() error {
				calls++
				if left > 0 {
					left--
					return boom
				}
				return nil
			}, tc.attempts, 0)

			if calls != tc.wantCalls {
				t.Errorf("called del %d times, want %d", calls, tc.wantCalls)
			}
			if tc.wantErr {
				if !errors.Is(err, boom) {
					t.Errorf("err = %v, want %v so the caller can report the stranded object", err, boom)
				}
			} else if err != nil {
				t.Errorf("err = %v, want nil", err)
			}
		})
	}
}
