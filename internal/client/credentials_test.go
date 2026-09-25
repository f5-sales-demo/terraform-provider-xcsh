// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package client

import (
	"context"
	"net/http"
	"strings"
	"testing"
)

type countingRoundTripper struct {
	calls int
}

func (r *countingRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	r.calls++
	return nil, nil
}

func TestUnauthenticatedClientRejectsBeforeHTTPRequest(t *testing.T) {
	transport := &countingRoundTripper{}
	client := NewUnauthenticatedClient("https://example.invalid", WithHTTPClient(&http.Client{Transport: transport}))
	_, err := client.doRequest(context.Background(), http.MethodGet, "/api/config/namespaces", nil)
	if err == nil || !strings.Contains(err.Error(), "F5XC API credentials are required") {
		t.Fatalf("doRequest() error = %v", err)
	}
	if transport.calls != 0 {
		t.Fatalf("HTTP transport called %d times, want 0", transport.calls)
	}
}
