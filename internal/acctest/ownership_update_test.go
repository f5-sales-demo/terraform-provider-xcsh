// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOwnershipTransportRejectsUnownedUpdate(t *testing.T) {
	mutations := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			mutations++
		}
		w.Write([]byte(`{}`))
	}))
	defer server.Close()
	j, err := CreateOwnershipJournal(t.TempDir()+"/journal.json", "runone")
	if err != nil {
		t.Fatal(err)
	}
	c := client.NewClient(server.URL, "test")
	c.HTTPClient.Transport = &ownershipTransport{base: http.DefaultTransport, c: c, j: j}
	err = c.Put(context.Background(), "/api/config/namespaces/system/origin_pools/unrelated", map[string]interface{}{"metadata": map[string]string{"name": "unrelated", "namespace": "system"}, "spec": map[string]interface{}{}}, nil)
	if err == nil || mutations != 0 {
		t.Fatal("unowned update sent")
	}
}
