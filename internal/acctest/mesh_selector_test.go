// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMeshReferenceRequiresCustomerEdgeSelector(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/api/config/namespaces/system/virtual_sites" {
			_, _ = w.Write([]byte(`{"items":[{"name":"regional"},{"name":"customer"}]}`))
			return
		}
		if r.URL.Path == "/api/config/namespaces/system/virtual_sites/customer/selectees" {
			_, _ = w.Write([]byte(`{"items":[{"name":"fixture-site","namespace":"system","kind":"site"}]}`))
			return
		}
		if r.URL.Path == "/api/config/namespaces/system/sites/fixture-site" {
			_, _ = w.Write([]byte(`{"metadata":{"name":"fixture-site","namespace":"system"},"system_metadata":{"uid":"immutable-site"}}`))
			return
		}
		typ := "REGIONAL_EDGE"
		name := "regional"
		if r.URL.Path == "/api/config/namespaces/system/virtual_sites/customer" {
			typ = "CUSTOMER_EDGE"
			name = "customer"
		}
		_, _ = w.Write([]byte(`{"metadata":{"name":"` + name + `","namespace":"system"},"system_metadata":{"uid":"immutable"},"spec":{"site_type":"` + typ + `","site_selector":{"expressions":["fixture=true"]}}}`))
	}))
	defer server.Close()
	ref, err := discoverMeshVirtualSite(context.Background(), client.NewClient(server.URL, "synthetic"))
	if err != nil || ref.Name != "customer" {
		t.Fatal("mesh selected non-customer infrastructure")
	}
}
