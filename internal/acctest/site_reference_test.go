// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExistingSiteReferenceRequiresTypedExactReadback(t *testing.T) {
	for _, tc := range []struct {
		name, listing, read string
		valid               bool
	}{
		{"verified", `{"items":[{"name":"site-b"},{"name":"site-a"}]}`, `{"metadata":{"name":"site-a","namespace":"system"},"system_metadata":{"uid":"immutable"}}`, true},
		{"missing list", `{}`, `{}`, false},
		{"missing name", `{"items":[{}]}`, `{}`, false},
		{"wrong name", `{"items":[{"name":"site-a"}]}`, `{"metadata":{"name":"site-b"},"system_metadata":{"uid":"immutable"}}`, false},
		{"missing uid", `{"items":[{"name":"site-a"}]}`, `{"metadata":{"name":"site-a"}}`, false},
		{"other run", `{"items":[{"name":"tf-acc-test-otherrun-site"}]}`, `{}`, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			reads := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "GET" {
					t.Error("existing appliance mutation attempted")
				}
				w.Header().Set("Content-Type", "application/json")
				if r.URL.Path == "/api/config/namespaces/system/sites" {
					_, _ = w.Write([]byte(tc.listing))
					return
				}
				if r.URL.Path != "/api/config/namespaces/system/sites/site-a" {
					t.Error("unstable site selection")
				}
				reads++
				_, _ = w.Write([]byte(tc.read))
			}))
			defer server.Close()
			ref, err := discoverReadOnlySite(context.Background(), client.NewClient(server.URL, "synthetic"))
			if (err == nil) != tc.valid {
				t.Fatalf("valid=%v expected=%v", err == nil, tc.valid)
			}
			if tc.valid && (reads != 1 || ref.Name != "site-a" || ref.Namespace != "system") {
				t.Fatal("reference not verified")
			}
		})
	}
}
