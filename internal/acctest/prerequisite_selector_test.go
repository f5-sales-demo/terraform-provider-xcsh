// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDisposableMeshSelectorClonesVerifiedSelection(t *testing.T) {
	t.Setenv("XCSH_ACC_RUN_ID", "runone")
	t.Setenv("XCSH_ACC_PRIVATE", t.TempDir())
	posts := 0
	var created string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "POST" {
			var body client.VirtualSite
			if json.NewDecoder(r.Body).Decode(&body) != nil {
				t.Fatal("malformed create")
			}
			if body.Metadata.Namespace != "system" || body.Spec["site_type"] != "CUSTOMER_EDGE" {
				t.Error("wrong selector scope")
			}
			created = body.Metadata.Name
			posts++
			json.NewEncoder(w).Encode(body)
			return
		}
		if r.URL.Path == "/api/config/namespaces/system/virtual_sites" {
			w.Write([]byte(`{"items":[{"name":"existing"}]}`))
			return
		}
		if r.URL.Path == "/api/config/namespaces/system/sites/site" {
			w.Write([]byte(`{"metadata":{"name":"site","namespace":"system"},"system_metadata":{"uid":"site-uid"}}`))
			return
		}
		if r.URL.Path == "/api/config/namespaces/system/virtual_sites/existing/selectees" || r.URL.Path == "/api/config/namespaces/system/virtual_sites/"+created+"/selectees" {
			w.Write([]byte(`{"items":[{"name":"site","namespace":"system","kind":"site"}]}`))
			return
		}
		if r.URL.Path == "/api/config/namespaces/system/virtual_sites/existing" {
			w.Write([]byte(`{"metadata":{"name":"existing","namespace":"system"},"system_metadata":{"uid":"virtual-uid"},"spec":{"site_type":"CUSTOMER_EDGE","site_selector":{"expressions":["fixture=true"]}}}`))
			return
		}
		t.Error("unexpected read")
	}))
	defer server.Close()
	if err := provisionMeshSelector(context.Background(), client.NewClient(server.URL, "synthetic")); err != nil || posts != 1 {
		t.Fatal("verified disposable selector not provisioned")
	}
}
