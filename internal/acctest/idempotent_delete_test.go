// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
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

func TestAlertPolicyDisappearanceDeleteUsesRequiredBody(t *testing.T) {
	t.Parallel()
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if r.Method != http.MethodDelete || r.URL.Path != "/api/config/namespaces/system/alert_policys/fixture" {
			t.Errorf("delete request = %s %s", r.Method, r.URL.Path)
		}
		var body map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("decode delete request: %v", err)
		}
		if len(body) != 0 {
			t.Errorf("delete body = %#v", body)
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
