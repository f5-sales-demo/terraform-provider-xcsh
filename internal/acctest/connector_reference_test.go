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
		{"matching", "owner-uid", true}, {"other site", "other-uid", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch r.URL.Path {
				case "/api/config/namespaces/system/sites/fixture":
					_, _ = w.Write([]byte(`{"metadata":{"name":"fixture","namespace":"system"},"system_metadata":{"uid":"site-uid","owner_view":{"kind":"securemesh_site_v2","name":"backing","namespace":"system","uid":"owner-uid"}},"spec":{"site_type":"CUSTOMER_EDGE","main_nodes":[{"name":"fixture-node"}]}}`))
				case "/api/config/namespaces/system/securemesh_site_v2s/backing":
					_, _ = w.Write([]byte(`{"metadata":{"name":"backing","namespace":"system"},"system_metadata":{"uid":"owner-uid"}}`))
				case "/api/config/namespaces/system/virtual_networks/ves-io-site-local-inside":
					_, _ = w.Write([]byte(`{"spec":{"site_local_inside_network":{}}}`))
				case "/api/config/namespaces/system/network_interfaces":
					_, _ = w.Write([]byte(`{"items":[{"name":"fixture-interface","owner_view":{"uid":"` + tc.owner + `"},"get_spec":{"type":"NETWORK_INTERFACE_ETHERNET","device_name":"eth1","virtual_network":[{"kind":"virtual_network","name":"ves-io-site-local-inside","namespace":"system"}]}}],"errors":[]}`))
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
			if tc.want && (ref.Node != "fixture-node" || ref.Interface != "fixture-interface") {
				t.Fatal("endpoint correlation lost")
			}
		})
	}
}
