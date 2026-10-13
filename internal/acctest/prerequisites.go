// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"os"
	"strings"
	"testing"
	"time"
)

func RequireExternalCredential(t *testing.T, key string) {
	t.Helper()
	switch key {
	case "vcs", "central_manager", "cloud", "secret_management":
	default:
		t.Fatal("unregistered credential prerequisite")
	}
	if os.Getenv("XCSH_ACC_"+strings.ToUpper(key)+"_CREDENTIAL") == "" {
		t.Logf("[PREREQUISITE:external_credential_missing method=GET type=%s]", key)
		t.Skip("external credential prerequisite unavailable")
	}
	t.Fatal("external credential-backed fixture requires explicit configuration")
}
func RequireNonDisposableScope(t *testing.T) {
	t.Helper()
	t.Log("[PREREQUISITE:non_disposable_scope method=GET type=tenant_configuration]")
	t.Skip("singleton lifecycle would adopt independently owned tenant configuration")
}
func RequireExistingInfrastructure(t *testing.T, kind string) {
	t.Helper()
	if strings.HasPrefix(t.Name(), "TestAcc") {
		SkipIfNotAccTest(t)
	}
	if IsMockMode() {
		return
	}
	if kind == "site" {
		requireReadOnlySite(t)
		return
	}
	paths := map[string]string{
		"vk8s":               "/api/config/namespaces/system/virtual_k8ss",
		"site":               "/api/config/namespaces/system/sites",
		"cloud_site":         "/api/config/namespaces/system/aws_tgw_sites",
		"nfv":                "/api/config/namespaces/system/nfv_services",
		"discovery":          "/api/config/namespaces/system/discoverys",
		"external_connector": "/api/config/namespaces/system/external_connectors",
	}
	path, ok := paths[kind]
	if !ok {
		t.Fatal("unregistered infrastructure prerequisite")
	}
	c, err := GetTestClient()
	if err != nil {
		t.Fatal("infrastructure client unavailable")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var listing struct {
		Items []struct {
			Name     string          `json:"name"`
			Metadata client.Metadata `json:"metadata"`
		} `json:"items"`
	}
	if err = c.Get(ctx, path, &listing); err != nil {
		if classifyCapabilityProbe(err) == capabilityDenied {
			t.Logf("[PREREQUISITE:operation_denied method=GET type=%s status=403]", kind)
			t.Skip("existing infrastructure read denied")
		}
		t.Fatal("infrastructure inventory unavailable")
	}
	if listing.Items == nil {
		t.Fatal("infrastructure inventory malformed")
	}
	if len(listing.Items) == 0 {
		t.Logf("[PREREQUISITE:infrastructure_absent method=GET type=%s]", kind)
		t.Skip("existing infrastructure collection is empty")
	}
	t.Fatal("existing infrastructure found; configure verified read-only fixture references")
}

func FailHistoricalPrerequisite(t *testing.T) {
	t.Helper()
	t.Fatal("unclassified historical prerequisite skip")
}
