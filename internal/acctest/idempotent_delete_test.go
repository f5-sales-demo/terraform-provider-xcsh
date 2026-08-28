// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
)

func TestRetryIdempotentDeleteRetriesOnlyTransientErrors(t *testing.T) {
	t.Parallel()

	attempts := 0
	transientThenSuccess := func(context.Context, *client.Client, string, string) error {
		attempts++
		if attempts < 3 {
			return &xcsherrors.XCSHError{Code: xcsherrors.ErrCodeServerError, StatusCode: http.StatusInternalServerError}
		}
		return nil
	}
	if err := retryIdempotentDelete(context.Background(), nil, transientThenSuccess, "system", "fixture", 3, 0); err != nil {
		t.Fatalf("transient idempotent delete did not recover: %v", err)
	}
	if attempts != 3 {
		t.Fatalf("attempts = %d, want 3", attempts)
	}

	persistent := &xcsherrors.XCSHError{Code: xcsherrors.ErrCodeBadRequest, StatusCode: http.StatusBadRequest}
	attempts = 0
	err := retryIdempotentDelete(context.Background(), nil, func(context.Context, *client.Client, string, string) error {
		attempts++
		return persistent
	}, "system", "fixture", 3, 0)
	if !errors.Is(err, persistent) || attempts != 1 {
		t.Fatalf("persistent delete error = %v after %d attempts; want original error after one attempt", err, attempts)
	}
}

func TestAlertPolicyDisappearanceDeleteUsesGeneratedRoute(t *testing.T) {
	t.Parallel()
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if r.Method != http.MethodDelete || r.URL.Path != "/api/config/namespaces/system/alert_policys/fixture" {
			t.Errorf("delete request = %s %s", r.Method, r.URL.Path)
		}
		if r.ContentLength > 0 {
			t.Errorf("delete content length = %d, want no request body", r.ContentLength)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	c := client.NewClient(server.URL, "fixture-token")
	deleter := resourceDeleterRegistry["xcsh_alert_policy"]
	if err := deleter(context.Background(), c, "system", "fixture"); err != nil {
		t.Fatalf("alert-policy external delete: %v", err)
	}
	if requests != 1 {
		t.Fatalf("delete requests = %d, want 1", requests)
	}
}

func TestWaitForResourceDisappearance(t *testing.T) {
	t.Parallel()

	t.Run("retries visible and transient reads until not found", func(t *testing.T) {
		calls := 0
		verifier := func(context.Context, *client.Client, string, string) error {
			calls++
			switch calls {
			case 1:
				return nil
			case 2:
				return &xcsherrors.XCSHError{Code: xcsherrors.ErrCodeServerError, StatusCode: http.StatusInternalServerError}
			default:
				return &xcsherrors.XCSHError{Code: xcsherrors.ErrCodeNotFound, StatusCode: http.StatusNotFound}
			}
		}
		if err := waitForResourceDisappearance(context.Background(), nil, verifier, "system", "fixture", 3, 0); err != nil {
			t.Fatalf("wait for disappearance: %v", err)
		}
		if calls != 3 {
			t.Fatalf("verification calls = %d, want 3", calls)
		}
	})

	t.Run("returns persistent errors unchanged", func(t *testing.T) {
		persistent := &xcsherrors.XCSHError{Code: xcsherrors.ErrCodeForbidden, StatusCode: http.StatusForbidden}
		calls := 0
		err := waitForResourceDisappearance(context.Background(), nil, func(context.Context, *client.Client, string, string) error {
			calls++
			return persistent
		}, "system", "fixture", 3, 0)
		if !errors.Is(err, persistent) || calls != 1 {
			t.Fatalf("persistent verification error = %v after %d calls", err, calls)
		}
	})

	t.Run("fails when resource remains visible", func(t *testing.T) {
		calls := 0
		err := waitForResourceDisappearance(context.Background(), nil, func(context.Context, *client.Client, string, string) error {
			calls++
			return nil
		}, "system", "fixture", 3, 0)
		if err == nil || !strings.Contains(err.Error(), "remained visible") || calls != 3 {
			t.Fatalf("visible resource error = %v after %d calls", err, calls)
		}
	})
}
