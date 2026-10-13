// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQuotaPrerequisitesRequireTypedLimitsAndCollectionCounts(t *testing.T) {
	for _, tc := range []struct {
		name, quota, listing    string
		unconfigured, exhausted bool
		valid                   bool
	}{
		{"missing forwarding", `{"objects":{}}`, `{"items":[]}`, true, false, true},
		{"exhausted receiver", `{"objects":{"global_log_receiver":{"limit":{"maximum":2}}}}`, `{"items":[{},{}]}`, false, true, true},
		{"capacity available", `{"objects":{"global_log_receiver":{"limit":{"maximum":2}}}}`, `{"items":[{}]}`, false, false, true},
		{"unknown quota", `{}`, `{"items":[]}`, false, false, false},
		{"unknown listing", `{"objects":{"global_log_receiver":{"limit":{"maximum":2}}}}`, `{}`, false, false, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "GET" {
					t.Error("quota read mutated tenant")
				}
				w.Header().Set("Content-Type", "application/json")
				if r.URL.Path == "/api/web/namespaces/system/quota/limits" {
					_, _ = w.Write([]byte(tc.quota))
				} else {
					_, _ = w.Write([]byte(tc.listing))
				}
			}))
			defer server.Close()
			typ := "xcsh_global_log_receiver"
			if tc.unconfigured {
				typ = "xcsh_forwarding_class"
			}
			receipt, err := readQuotaPrerequisite(context.Background(), client.NewClient(server.URL, "synthetic"), typ, "system")
			if (err == nil) != tc.valid {
				t.Fatalf("valid=%v want=%v", err == nil, tc.valid)
			}
			if tc.valid && (receipt.Unconfigured != tc.unconfigured || receipt.Exhausted != tc.exhausted) {
				t.Fatal("quota status inferred incorrectly")
			}
		})
	}
}
