// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

func TestMeshSelectorRequiresVerifiedExpansion(t *testing.T) {
	for _, tc := range []struct {
		name, expansion string
		status          int
		selected        bool
		failed          bool
	}{
		{"populated", `{"items":[{"name":"site-fixture","namespace":"system","kind":"site"}]}`, 200, true, false},
		{"empty", `{"items":[]}`, 200, false, false},
		{"malformed", `{}`, 200, false, true},
		{"missing identity", `{"items":[{}]}`, 200, false, true},
		{"embedded error", `{"items":[],"error":{"code":"UNKNOWN","message":"error"}}`, 200, false, true},
		{"authorization", `{"code":"FORBIDDEN"}`, 403, false, true},
		{"server error", `{"code":"UNKNOWN"}`, 500, false, true},
		{"stale site", `{"items":[{"name":"stale","namespace":"system","kind":"site"}]}`, 200, false, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "GET" {
					t.Error("appliance mutation attempted")
				}
				w.Header().Set("Content-Type", "application/json")
				switch r.URL.Path {
				case "/api/config/namespaces/system/virtual_sites":
					fmt.Fprint(w, `{"items":[{"name":"selector"}]}`)
				case "/api/config/namespaces/system/virtual_sites/selector":
					fmt.Fprint(w, `{"metadata":{"name":"selector","namespace":"system"},"system_metadata":{"uid":"immutable"},"spec":{"site_type":"CUSTOMER_EDGE","site_selector":{"expressions":["fixture=true"]}}}`)
				case "/api/config/namespaces/system/virtual_sites/selector/selectees":
					w.WriteHeader(tc.status)
					fmt.Fprint(w, tc.expansion)
				case "/api/config/namespaces/system/sites/site-fixture":
					fmt.Fprint(w, `{"metadata":{"name":"site-fixture","namespace":"system"},"system_metadata":{"uid":"immutable-site"}}`)
				case "/api/config/namespaces/system/sites/stale":
					w.WriteHeader(404)
					fmt.Fprint(w, `{"code":"NOT_FOUND"}`)
				default:
					t.Error("unexpected route")
				}
			}))
			defer server.Close()
			c := client.NewClient(server.URL, "synthetic")
			c.MaxRetries = 0
			ref, err := discoverMeshVirtualSite(context.Background(), c)
			if tc.selected {
				if err != nil || ref.Name != "selector" {
					t.Fatal("verified expansion not selected")
				}
				return
			}
			if err == nil {
				t.Fatal("unverified selector accepted")
			}
			absent := strings.Contains(err.Error(), "existing infrastructure absent")
			if absent == tc.failed {
				t.Fatal("failed expansion reinterpreted as absence")
			}
		})
	}
}
