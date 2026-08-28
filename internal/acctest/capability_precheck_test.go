// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
)

func TestCapabilityProbePathUsesExactReadOnlySurface(t *testing.T) {
	t.Parallel()
	tests := map[string]string{
		"TestAccCertificateResource_basic":     "/api/config/namespaces/system/certificates",
		"TestAccNamespaceResource_basic":       "/api/web/namespaces",
		"TestAccUDPLoadBalancerResource_basic": "/api/config/namespaces/system/udp_loadbalancers",
		"TestAccOriginPoolResource_basic":      "",
	}
	for testName, want := range tests {
		got, ok := capabilityProbePath(testName)
		if ok != (want != "") || got != want {
			t.Errorf("capabilityProbePath(%q) = %q, %t; want %q", testName, got, ok, want)
		}
	}
}

func TestCapabilityProbeSkipsOnlyForbidden(t *testing.T) {
	t.Parallel()
	tests := []struct {
		status int
		want   capabilityDecision
	}{
		{status: http.StatusForbidden, want: capabilityDenied},
		{status: http.StatusBadRequest, want: capabilityAllowed},
		{status: http.StatusNotFound, want: capabilityAllowed},
		{status: http.StatusTooManyRequests, want: capabilityInconclusive},
		{status: http.StatusInternalServerError, want: capabilityInconclusive},
	}
	for _, test := range tests {
		err := &xcsherrors.XCSHError{StatusCode: test.status}
		if got := classifyCapabilityProbe(err); got != test.want {
			t.Errorf("status %d classified as %v, want %v", test.status, got, test.want)
		}
	}
}

func TestNamespaceCapabilityInventoryIsExact(t *testing.T) {
	t.Parallel()
	if len(namespaceCreationCapabilityTests) != 126 {
		t.Fatalf("namespace capability inventory has %d tests, want 126", len(namespaceCreationCapabilityTests))
	}
	for _, name := range []string{"TestAccNamespaceResource_basic", "TestAccAlertPolicyDataSource_basic"} {
		if _, ok := namespaceCreationCapabilityTests[name]; !ok {
			t.Errorf("missing evidenced namespace-dependent test %q", name)
		}
	}
	if _, ok := namespaceCreationCapabilityTests["TestAccAlertPolicyResource_basic"]; ok {
		t.Error("system-namespace alert policy test must not be capability-skipped")
	}
}

func TestQuotaCapabilityProbeTreatsOnlyPersistentCapacitySignalAsDenied(t *testing.T) {
	t.Parallel()
	if got := classifyQuotaCapabilityProbe(&xcsherrors.XCSHError{StatusCode: http.StatusTooManyRequests}); got != capabilityDenied {
		t.Fatalf("quota capability 429 = %v, want denied", got)
	}
	if got := classifyQuotaCapabilityProbe(&xcsherrors.XCSHError{StatusCode: http.StatusInternalServerError}); got != capabilityInconclusive {
		t.Fatalf("quota capability 500 = %v, want inconclusive", got)
	}
}

func TestNamespaceCapabilityProbeIsInvalidNonRetriedPost(t *testing.T) {
	t.Parallel()
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if r.Method != http.MethodPost || r.URL.Path != "/api/web/namespaces" {
			t.Errorf("probe request = %s %s", r.Method, r.URL.Path)
		}
		var request struct {
			Metadata struct {
				Name string `json:"name"`
			} `json:"metadata"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil || request.Metadata.Name != "-" {
			t.Errorf("probe body is not the deliberately invalid namespace request")
		}
		w.WriteHeader(http.StatusForbidden)
	}))
	defer server.Close()

	c := client.NewClient(server.URL, "fixture-token")
	err := executeCapabilityProbe(context.Background(), c, "/api/web/namespaces", true)
	if classifyCapabilityProbe(err) != capabilityDenied || attempts != 1 {
		t.Fatalf("namespace probe decision = %v after %d attempts; want denied after one", classifyCapabilityProbe(err), attempts)
	}
}
