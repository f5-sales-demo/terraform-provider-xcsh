// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
	"testing"
)

func TestExhaustionRequiresTypedExactDiagnostic(t *testing.T) {
	for _, tc := range []struct {
		status  int
		message string
		want    bool
	}{{429, "Object kind synthetic has exhausted limits", true}, {429, "rate limit exceeded", false}, {500, "Object kind synthetic has exhausted limits", false}, {400, "Object kind synthetic has exhausted limits", false}} {
		err := &xcsherrors.XCSHError{StatusCode: tc.status, Message: tc.message}
		if confirmedObjectExhaustion(err) != tc.want {
			t.Fatal("generic API failure treated as exhaustion")
		}
	}
}
