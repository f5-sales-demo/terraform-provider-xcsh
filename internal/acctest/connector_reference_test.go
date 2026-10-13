// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

func TestConnectorEndpointCorrelatesSiteUID(t *testing.T) {
	for _, tc := range []struct {
		name, owner string
		want        bool
	}{
		{"matching", "site-uid", true}, {"other site", "other-uid", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch r.URL.Path {
				case "/api/config/namespaces/system/sites/fixture":
					_, _ = w.Write([]byte(`{"metadata":{"name":"fixture","namespace":"system"},"system_metadata":{"uid":"site-uid"},"spec":{"site_type":"CUSTOMER_EDGE","main_nodes":[{"name":"fixture-node"}]}}`))
				case "/api/config/namespaces/system/network_interfaces":
					_, _ = w.Write([]byte(`{"items":[{"owner_view":{"uid":"` + tc.owner + `"},"get_spec":{"type":"NETWORK_INTERFACE_ETHERNET","device_name":"eth1"}}],"errors":[]}`))
				default:
					t.Errorf("unexpected read")
					w.WriteHeader(http.StatusNotFound)
				}
			}))
			defer server.Close()
			ref, err := discoverConnectorEndpoint(context.Background(), client.NewClient(server.URL, "fixture"), readOnlySiteReference{Name: "fixture", Namespace: "system"})
			if (err == nil) != tc.want {
				t.Fatal("unverified site interface accepted or verified interface rejected")
			}
			if tc.want && (ref.Node != "fixture-node" || ref.Interface != "eth1") {
				t.Fatal("endpoint correlation lost")
			}
		})
	}
}
