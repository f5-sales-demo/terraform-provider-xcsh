// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"context"
	"errors"
	"net/http"
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
