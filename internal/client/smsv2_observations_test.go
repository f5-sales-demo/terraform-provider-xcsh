// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSMSv2ObservationEndpoints(t *testing.T) {
	requests := []struct{ method, path string }{
		{http.MethodGet, "/api/config/namespaces/system/securemesh_site_v2s/lab-site"},
		{http.MethodGet, "/api/operate/namespaces/system/sites/lab-site/vpm/debug/global/health"},
		{http.MethodGet, "/api/operate/namespaces/system/sites/lab-site/ver/bgp_peers"},
		{http.MethodGet, "/api/operate/namespaces/system/sites/lab-site/ver/bgp_routes"},
		{http.MethodPost, "/api/operate/namespaces/system/sites/lab-site/ver/simplified_routes"},
		{http.MethodGet, "/api/config/namespaces/system/sites/lab-site"},
		{http.MethodGet, "/api/maurice/upgradable_sw_versions"},
		{http.MethodGet, "/api/maurice/namespaces/system/sites/lab-site/pre_upgrade_check"},
		{http.MethodGet, "/api/maurice/namespaces/system/sites/lab-site/upgrade_status"},
	}
	index := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if index >= len(requests) {
			t.Errorf("unexpected request %s %s", request.Method, request.URL.Path)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		want := requests[index]
		index++
		if request.Method != want.method || request.URL.Path != want.path {
			t.Errorf("request = %s %s, want %s %s", request.Method, request.URL.Path, want.method, want.path)
		}
		if request.Method == http.MethodPost {
			var body map[string]interface{}
			if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			if _, ok := body["slo"]; !ok {
				t.Errorf("simplified route request has no slo selector: %#v", body)
			}
		}
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()
	c := NewClient(server.URL, "token", WithMaxRetries(0))
	ctx := context.Background()
	if _, err := c.GetSMSv2Configuration(ctx, "system", "lab-site"); err != nil {
		t.Fatal(err)
	}
	if _, err := c.GetSMSv2Health(ctx, "lab-site"); err != nil {
		t.Fatal(err)
	}
	if _, err := c.GetSMSv2BGPPeers(ctx, "system", "lab-site"); err != nil {
		t.Fatal(err)
	}
	if _, err := c.GetSMSv2BGPRoutes(ctx, "system", "lab-site"); err != nil {
		t.Fatal(err)
	}
	if _, err := c.GetSMSv2SimplifiedRoutes(ctx, "system", "lab-site", "slo"); err != nil {
		t.Fatal(err)
	}
	if _, err := c.GetSMSv2SiteUpgradeStatus(ctx, "system", "lab-site"); err != nil {
		t.Fatal(err)
	}
	if _, err := c.GetSMSv2UpgradableSoftwareVersions(ctx, "9.2026.10", "crt-20251002-0027"); err != nil {
		t.Fatal(err)
	}
	if _, err := c.GetSMSv2PreUpgradeCheck(ctx, "system", "lab-site", "crt-20260201-0179"); err != nil {
		t.Fatal(err)
	}
	if _, err := c.GetSMSv2UpgradeProgress(ctx, "system", "lab-site"); err != nil {
		t.Fatal(err)
	}
	if index != len(requests) {
		t.Fatalf("made %d requests, want %d", index, len(requests))
	}
}

func TestSMSv2ObservationHTTPFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		http.Error(writer, "unavailable", http.StatusBadRequest)
	}))
	defer server.Close()
	c := NewClient(server.URL, "token", WithMaxRetries(0))
	if _, err := c.GetSMSv2BGPPeers(context.Background(), "system", "lab-site"); err == nil {
		t.Fatal("expected HTTP failure")
	}
	if _, err := c.GetSMSv2SimplifiedRoutes(context.Background(), "system", "lab-site", "external"); err == nil {
		t.Fatal("expected unsupported role failure")
	}
}
