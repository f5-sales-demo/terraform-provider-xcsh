// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"errors"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCustomerEdgeInventoryDoesNotChooseRegionalEdge(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/config/namespaces/system/sites":
			_, _ = w.Write([]byte(`{"items":[{"name":"a-regional"},{"name":"b-customer"}]}`))
		case "/api/config/namespaces/system/sites/a-regional":
			_, _ = w.Write([]byte(`{"metadata":{"name":"a-regional","namespace":"system"},"system_metadata":{"uid":"regional"},"spec":{"site_type":"REGIONAL_EDGE"}}`))
		case "/api/config/namespaces/system/sites/b-customer":
			_, _ = w.Write([]byte(`{"metadata":{"name":"b-customer","namespace":"system"},"system_metadata":{"uid":"customer"},"spec":{"site_type":"CUSTOMER_EDGE"}}`))
		default:
			t.Fatal("unexpected appliance operation")
		}
	}))
	defer server.Close()
	refs, err := discoverCustomerEdgeSites(context.Background(), client.NewClient(server.URL, "fixture"))
	if err != nil || len(refs) != 1 || refs[0].Name != "b-customer" {
		t.Fatal("regional edge selected for customer-edge operation")
	}
}

func TestCustomerEdgeMissingTypeFailsInsteadOfSkipping(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/api/config/namespaces/system/sites" {
			_, _ = w.Write([]byte(`{"items":[{"name":"fixture"}]}`))
			return
		}
		_, _ = w.Write([]byte(`{"metadata":{"name":"fixture","namespace":"system"},"system_metadata":{"uid":"fixture"}}`))
	}))
	defer server.Close()
	_, err := discoverCustomerEdgeSites(context.Background(), client.NewClient(server.URL, "fixture"))
	if err == nil || errors.Is(err, errEligibleInfrastructureAbsent) {
		t.Fatal("malformed site became capability skip")
	}
}
