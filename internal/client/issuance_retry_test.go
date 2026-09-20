package client

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestIssuanceGetNeverRetries(t *testing.T) {
	for _, status := range []int{429, 500, 502, 503, 504} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			calls := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				calls++
				w.WriteHeader(status)
			}))
			defer server.Close()
			c := NewClient(server.URL, "test-token", WithMaxRetries(3), WithRetryWait(time.Millisecond, time.Millisecond))
			if err := c.GetOnce(context.Background(), "/issue", nil); err == nil {
				t.Fatal("expected server error")
			}
			if calls != 1 {
				t.Fatalf("credential operation replayed: %d calls", calls)
			}
		})
	}
}

type issuanceTransport struct{ calls int }

func (transport *issuanceTransport) RoundTrip(*http.Request) (*http.Response, error) {
	transport.calls++
	return nil, io.ErrUnexpectedEOF
}

func TestIssuanceGetNetworkFailureNeverRetries(t *testing.T) {
	transport := &issuanceTransport{}
	c := NewClient("https://example.invalid", "test-token", WithMaxRetries(3), WithRetryWait(time.Millisecond, time.Millisecond))
	c.HTTPClient.Transport = transport
	if err := c.GetOnce(context.Background(), "/issue", nil); err == nil {
		t.Fatal("expected network failure")
	}
	if transport.calls != 1 {
		t.Fatalf("credential operation replayed: %d calls", transport.calls)
	}
}

func TestIssuanceGetDoesNotFollowRedirect(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if r.URL.Path == "/issue" {
			http.Redirect(w, r, "/replay", http.StatusTemporaryRedirect)
		}
	}))
	defer server.Close()
	c := NewClient(server.URL, "test-token")
	if err := c.GetOnce(context.Background(), "/issue", nil); err == nil {
		t.Fatal("redirect must be returned as a failure")
	}
	if requests != 1 {
		t.Fatalf("redirect replayed issuance: %d", requests)
	}
	if c.HTTPClient.CheckRedirect != nil || c.HTTPClient.Transport != nil {
		t.Fatal("shared client was modified")
	}
}
