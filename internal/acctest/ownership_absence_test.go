// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
	"net/http"
	"testing"
)

func TestOwnedAbsenceRejectsMalformedAndWrongStatus(t *testing.T) {
	for _, tc := range []struct {
		status int
		body   string
		want   bool
	}{
		{404, `{"code":"NOT_FOUND","message":"absent"}`, true},
		{404, `{"code":5,"message":"absent"}`, true},
		{404, "not found", false}, {403, `{"code":"NOT_FOUND"}`, false},
		{404, `{"message":"not found"}`, false}, {500, `{"code":"NOT_FOUND"}`, false},
	} {
		err := xcsherrors.NewAPIError(tc.status, []byte(tc.body), "fixture", http.MethodGet)
		if got := authoritativeOwnedAbsence(err); got != tc.want {
			t.Fatalf("status %d body %q=%v", tc.status, tc.body, got)
		}
	}
}
