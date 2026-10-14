// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadOnlySiteSkipsOnlyTypedStaleInventoryEntry(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/config/namespaces/system/sites":
			_, _ = w.Write([]byte(`{"items":[{"name":"site-a"},{"name":"site-b"}]}`))
		case "/api/config/namespaces/system/sites/site-a":
			w.WriteHeader(404)
			_, _ = w.Write([]byte(`{"code":"NOT_FOUND"}`))
		case "/api/config/namespaces/system/sites/site-b":
			_, _ = w.Write([]byte(`{"metadata":{"name":"site-b","namespace":"system"},"system_metadata":{"uid":"immutable"}}`))
		default:
			t.Error("unexpected route")
		}
	}))
	defer server.Close()
	ref, err := discoverReadOnlySite(context.Background(), client.NewClient(server.URL, "synthetic"))
	if err != nil || ref.Name != "site-b" {
		t.Fatal("typed stale entry prevented verified reference")
	}
}
