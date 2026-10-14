// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCapabilityProbeSkipsOnlyForbidden(t *testing.T) {
	for _, status := range []int{400, 403, 404, 429, 500} {
		want := capabilityInconclusive
		if status == 403 {
			want = capabilityDenied
		}
		if got := classifyCapabilityProbe(&xcsherrors.XCSHError{StatusCode: status}); got != want {
			t.Fatalf("status %d=%v", status, got)
		}
	}
}
func TestCapabilityOperationUsesInvalidSinglePost(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if r.Method != "POST" || r.URL.Path != "/api/web/namespaces" {
			t.Error("wrong operation")
		}
		w.WriteHeader(403)
	}))
	defer server.Close()
	c := client.NewClient(server.URL, "test")
	if decision := classifyCapabilityProbe(executeCapabilityProbe(context.Background(), c, "/api/web/namespaces")); decision != capabilityDenied || attempts != 1 {
		t.Fatalf("decision=%v attempts=%d", decision, attempts)
	}
}
