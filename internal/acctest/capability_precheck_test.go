// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"net/http"
	"testing"

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
