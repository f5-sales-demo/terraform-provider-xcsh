package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetUnimplementedIsTerminal(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		w.WriteHeader(http.StatusNotImplemented)
		_, _ = fmt.Fprint(w, `{"code":12,"message":"unimplemented"}`)
	}))
	defer server.Close()
	c := NewClient(server.URL, "fixture-token", WithMaxRetries(3), WithRetryWait(time.Millisecond, 2*time.Millisecond))
	var value map[string]interface{}
	err := c.Get(context.Background(), "/unavailable", &value)
	if !HasHTTPStatus(err, http.StatusNotImplemented) {
		t.Fatalf("expected typed HTTP 501, got %v", err)
	}
	if requests != 1 {
		t.Fatalf("terminal 501 was retried: %d requests", requests)
	}
}
