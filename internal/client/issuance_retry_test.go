package client

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
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

func TestPostEOFIsAmbiguousAndNeverRetries(t *testing.T) {
	transport := &issuanceTransport{}
	const secret = "test-super-secret"
	c := NewClient("https://example.invalid", secret, WithMaxRetries(3), WithRetryWait(time.Millisecond, time.Millisecond))
	c.HTTPClient.Transport = transport
	err := c.Post(context.Background(), "/mutation", map[string]string{"diagnostic": "invalid"}, nil)
	if err == nil {
		t.Fatal("expected network failure")
	}
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Fatalf("expected EOF in error chain, got %T", err)
	}
	if transport.calls != 1 {
		t.Fatalf("mutation replayed after ambiguous EOF: %d calls", transport.calls)
	}
	if strings.Contains(err.Error(), secret) {
		t.Fatal("network diagnostic exposed the authorization credential")
	}
}

func TestPostDoesNotFollowRedirect(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if r.URL.Path == "/mutation" {
			http.Redirect(w, r, "/replay", http.StatusTemporaryRedirect)
		}
	}))
	defer server.Close()
	c := NewClient(server.URL, "test-token")
	if err := c.Post(context.Background(), "/mutation", map[string]string{"diagnostic": "invalid"}, nil); err == nil {
		t.Fatal("redirect must be returned as a failure")
	}
	if requests != 1 {
		t.Fatalf("redirect replayed mutation: %d", requests)
	}
}

func TestPostUsesFixedLengthHTTP1AfterSourceTransportUsedHTTP2(t *testing.T) {
	type observation struct {
		protoMajor    int
		contentLength int64
		close         bool
		body          []byte
	}
	observations := make(chan observation, 3)
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/prime" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("read request body: %v", err)
		}
		observations <- observation{r.ProtoMajor, r.ContentLength, r.Close, body}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"code":"invalid"}`))
	}))
	server.EnableHTTP2 = true
	server.StartTLS()
	defer server.Close()

	transport := server.Client().Transport.(*http.Transport).Clone()
	transport.ForceAttemptHTTP2 = true
	primeClient := &http.Client{Transport: transport}
	response, err := primeClient.Get(server.URL + "/prime")
	if err != nil {
		t.Fatalf("prime HTTP/2 transport: %v", err)
	}
	_ = response.Body.Close()
	if response.ProtoMajor != 2 {
		t.Fatalf("source transport did not negotiate HTTP/2: %s", response.Proto)
	}

	payload := map[string]string{"diagnostic": "invalid"}
	expectedBody, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	c := NewClient(server.URL, "test-token", WithHTTPClient(&http.Client{Transport: transport}))
	post := func() {
		err := c.Post(context.Background(), "/mutation", payload, nil)
		var apiErr *xcsherrors.XCSHError
		if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusBadRequest {
			t.Errorf("expected HTTP 400 response, got %T", err)
		}
	}

	post()
	var group sync.WaitGroup
	for range 2 {
		group.Add(1)
		go func() {
			defer group.Done()
			post()
		}()
	}
	group.Wait()
	close(observations)
	count := 0
	for got := range observations {
		count++
		if got.protoMajor != 1 {
			t.Errorf("mutation protocol = HTTP/%d, want HTTP/1.1", got.protoMajor)
		}
		if got.contentLength != int64(len(expectedBody)) {
			t.Errorf("content length = %d, want %d", got.contentLength, len(expectedBody))
		}
		if !got.close {
			t.Error("mutation connection was not marked close")
		}
		if string(got.body) != string(expectedBody) {
			t.Error("mutation body changed in transport")
		}
	}
	if count != 3 {
		t.Fatalf("mutation request count = %d, want 3", count)
	}
}
