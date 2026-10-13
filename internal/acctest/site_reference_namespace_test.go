// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInfrastructureReferenceUsesListedNamespace(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/api/config/namespaces/system/virtual_sites" {
			_, _ = w.Write([]byte(`{"items":[{"name":"fixture","namespace":"shared"}]}`))
			return
		}
		if r.URL.Path != "/api/config/namespaces/shared/virtual_sites/fixture" {
			t.Error("listed namespace ignored")
		}
		_, _ = w.Write([]byte(`{"metadata":{"name":"fixture","namespace":"shared"},"system_metadata":{"uid":"immutable"}}`))
	}))
	defer server.Close()
	ref, err := discoverReadOnlyInfrastructure(context.Background(), client.NewClient(server.URL, "synthetic"), "virtual_sites")
	if err != nil || ref.Namespace != "shared" {
		t.Fatal("verified namespace lost")
	}
}
