// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOwnershipProjectionUsesSystemUID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"metadata":{"name":"fixture","namespace":"system","annotations":{"f5-sales-demo.com/acceptance-run":"runone"}},"system_metadata":{"uid":"immutable-uid"},"spec":{}}`))
	}))
	defer server.Close()
	c := client.NewClient(server.URL, "test")
	metadata, err := readOwnershipProjection(context.Background(), c, "/api/config/namespaces/system/origin_pools/fixture", "system", "fixture")
	if err != nil || metadata.UID != "immutable-uid" || metadata.Annotations[ownershipAnnotation] != "runone" {
		t.Fatalf("ownership projection %#v %v", metadata, err)
	}
}
