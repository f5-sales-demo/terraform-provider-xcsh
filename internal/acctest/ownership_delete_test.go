// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOwnershipTransportRejectsUnownedDelete(t *testing.T) {
	mutations := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			mutations++
		}
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()
	j, err := CreateOwnershipJournal(t.TempDir()+"/journal.json", "runone")
	if err != nil {
		t.Fatal(err)
	}
	c := client.NewClient(server.URL, "synthetic")
	c.HTTPClient.Transport = &ownershipTransport{base: http.DefaultTransport, c: c, j: j}
	err = c.Delete(context.Background(), "/api/config/namespaces/system/origin_pools/tf-acc-test-otherrun-pool")
	if err == nil || mutations != 0 {
		t.Fatal("another run resource deletion reached API")
	}
}
func TestOwnershipTransportRejectsUnownedNamespaceCascade(t *testing.T) {
	mutations := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			mutations++
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	}))
	defer server.Close()
	j, err := CreateOwnershipJournal(t.TempDir()+"/journal.json", "runone")
	if err != nil {
		t.Fatal(err)
	}
	c := client.NewClient(server.URL, "synthetic")
	c.MaxRetries = 0
	c.HTTPClient.Transport = &ownershipTransport{base: http.DefaultTransport, c: c, j: j}
	if err = c.CascadeDeleteNamespace(context.Background(), "tf-acc-test-otherrun-ns"); err == nil || mutations != 0 {
		t.Fatal("unowned namespace cascade reached API")
	}
}
