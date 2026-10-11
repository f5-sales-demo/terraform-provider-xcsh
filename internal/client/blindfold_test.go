package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNativeBlindfoldUncertainWrite(t *testing.T) {
	desired := map[string]any{"metadata": map[string]any{"name": "synthetic", "annotations": map[string]any{"f5-sales-demo.com/blindfold": "synthetic"}}, "spec": map[string]any{"certificate_url": "string:///public", "private_key": map[string]any{"blindfold_secret_info": map[string]any{"location": "string:///encrypted"}}}}
	writes := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			writes++
			w.WriteHeader(503)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"replace_form": desired, "resource_version": "version"})
	}))
	defer server.Close()
	c := NewClient(server.URL, "synthetic")
	var result map[string]any
	if err := c.Post(context.Background(), "/api/config/namespaces/synthetic/certificates", desired, &result); err != nil {
		t.Fatal(err)
	}
	if writes != 1 {
		t.Fatal("mutation retried")
	}
	if result["resource_version"] != "version" {
		t.Fatal("token lost")
	}
}

func TestNativeBlindfoldDefinitiveRejection(t *testing.T) {
	for _, status := range []int{400, 401, 403, 404, 413, 422, 429} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			desired := map[string]any{"metadata": map[string]any{"name": "synthetic", "annotations": map[string]any{"f5-sales-demo.com/blindfold": "synthetic"}}, "spec": map[string]any{"certificate_url": "string:///public"}}
			writes, reads := 0, 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodPost {
					writes++
					w.WriteHeader(status)
					_, _ = w.Write([]byte(`{"message":"PRIVATE-DIAGNOSTIC-MARKER"}`))
					return
				}
				reads++
				w.WriteHeader(http.StatusNotFound)
			}))
			defer server.Close()
			c := NewClient(server.URL, "synthetic")
			err := c.Post(context.Background(), "/api/config/namespaces/synthetic/certificates", desired, nil)
			if err == nil || !HasHTTPStatus(err, status) {
				t.Fatalf("rejection status lost: %v", err)
			}
			if strings.Contains(err.Error(), "PRIVATE-DIAGNOSTIC-MARKER") || strings.Contains(err.Error(), "unresolved") {
				t.Fatalf("unsafe or uncertain rejection: %v", err)
			}
			if writes != 1 || reads != 0 {
				t.Fatalf("definitive rejection made %d writes and %d reads", writes, reads)
			}
		})
	}
}
