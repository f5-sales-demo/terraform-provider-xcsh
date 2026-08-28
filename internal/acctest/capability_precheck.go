// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"testing"

	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
)

var capabilityProbeCache sync.Map

var quotaLimitedCapabilityTests = map[string]string{
	"TestAccAPICrawlerResource_basic": "/api/config/namespaces/system/api_crawlers",
	"TestAccFleetDataSource_basic":    "/api/config/namespaces/system/fleets",
	"TestAccFleetResource_basic":      "/api/config/namespaces/system/fleets",
}

type capabilityRoute struct {
	testPrefix   string
	resourceType string
	webNamespace bool
}

// These surfaces returned a stable 403 in consecutive exact-SHA live inventories.
// Each affected test performs a read-only probe before creating anything. A skip is
// permitted only when that probe itself returns 403; all other outcomes run the test.
var capabilityRoutes = []capabilityRoute{
	{testPrefix: "TestAccHTTPLoadBalancerResource_httpsAutoCertAdvertiseCustomVirtualSite", resourceType: "virtual_sites"},
	{testPrefix: "TestAccTLSChain_blindfoldCertWithOriginPool", resourceType: "certificates"},
	{testPrefix: "TestAccBotDefenseAppInfrastructure", resourceType: "bot_defense_app_infrastructures"},
	{testPrefix: "TestAccMaliciousUserMitigation", resourceType: "malicious_user_mitigations"},
	{testPrefix: "TestAccEnhancedFirewallPolicy", resourceType: "enhanced_firewall_policys"},
	{testPrefix: "TestAccSensitiveDataPolicy", resourceType: "sensitive_data_policys"},
	{testPrefix: "TestAccCodeBaseIntegration", resourceType: "code_base_integrations"},
	{testPrefix: "TestAccCertificateChain", resourceType: "certificate_chains"},
	{testPrefix: "TestAccContainerRegistry", resourceType: "container_registrys"},
	{testPrefix: "TestAccProtocolInspection", resourceType: "protocol_inspections"},
	{testPrefix: "TestAccGlobalLogReceiver", resourceType: "global_log_receivers"},
	{testPrefix: "TestAccRateLimiterPolicy", resourceType: "rate_limiter_policys"},
	{testPrefix: "TestAccTrustedCaList", resourceType: "trusted_ca_lists"},
	{testPrefix: "TestAccIPPrefixSet", resourceType: "ip_prefix_sets"},
	{testPrefix: "TestAccIpPrefixSet", resourceType: "ip_prefix_sets"},
	{testPrefix: "TestAccAlertReceiver", resourceType: "alert_receivers"},
	{testPrefix: "TestAccAlertPolicy", resourceType: "alert_policys"},
	{testPrefix: "TestAccApiDefinition", resourceType: "api_definitions"},
	{testPrefix: "TestAccApiDiscovery", resourceType: "api_discoverys"},
	{testPrefix: "TestAccApiTesting", resourceType: "api_testings"},
	{testPrefix: "TestAccApiCrawler", resourceType: "api_crawlers"},
	{testPrefix: "TestAccAPICrawler", resourceType: "api_crawlers"},
	{testPrefix: "TestAccAppApiGroup", resourceType: "app_api_groups"},
	{testPrefix: "TestAccAppSetting", resourceType: "app_settings"},
	{testPrefix: "TestAccCDNCacheRule", resourceType: "cdn_cache_rules"},
	{testPrefix: "TestAccCdnCacheRule", resourceType: "cdn_cache_rules"},
	{testPrefix: "TestAccCertificate", resourceType: "certificates"},
	{testPrefix: "TestAccCRL", resourceType: "crls"},
	{testPrefix: "TestAccCrl", resourceType: "crls"},
	{testPrefix: "TestAccCluster", resourceType: "clusters"},
	{testPrefix: "TestAccEndpoint", resourceType: "endpoints"},
	{testPrefix: "TestAccDNSComplianceChecks", resourceType: "dns_compliance_checkss"},
	{testPrefix: "TestAccDnsComplianceChecks", resourceType: "dns_compliance_checkss"},
	{testPrefix: "TestAccIrule", resourceType: "irules"},
	{testPrefix: "TestAccLogReceiver", resourceType: "log_receivers"},
	{testPrefix: "TestAccNamespace", webNamespace: true},
	{testPrefix: "TestAccNetworkPolicy", resourceType: "network_policys"},
	{testPrefix: "TestAccProxy", resourceType: "proxys"},
	{testPrefix: "TestAccVirtualHost", resourceType: "virtual_hosts"},
	{testPrefix: "TestAccVirtualSite", resourceType: "virtual_sites"},
	{testPrefix: "TestAccUDPLoadBalancer", resourceType: "udp_loadbalancers"},
}

func capabilityProbePath(testName string) (string, bool) {
	for _, route := range capabilityRoutes {
		if !strings.HasPrefix(testName, route.testPrefix) {
			continue
		}
		if route.webNamespace {
			return "/api/web/namespaces", true
		}
		return fmt.Sprintf("/api/config/namespaces/system/%s", route.resourceType), true
	}
	return "", false
}

type capabilityDecision int

const (
	capabilityInconclusive capabilityDecision = iota
	capabilityAllowed
	capabilityDenied
)

func classifyCapabilityProbe(err error) capabilityDecision {
	if err == nil {
		return capabilityAllowed
	}
	var apiErr *xcsherrors.XCSHError
	if !errors.As(err, &apiErr) {
		return capabilityInconclusive
	}
	switch apiErr.StatusCode {
	case http.StatusForbidden:
		return capabilityDenied
	case http.StatusBadRequest, http.StatusNotFound:
		return capabilityAllowed
	default:
		return capabilityInconclusive
	}
}

func classifyQuotaCapabilityProbe(err error) capabilityDecision {
	var apiErr *xcsherrors.XCSHError
	if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusTooManyRequests {
		return capabilityDenied
	}
	return classifyCapabilityProbe(err)
}

func executeCapabilityProbe(ctx context.Context, c interface {
	Get(context.Context, string, interface{}) error
	Post(context.Context, string, interface{}, interface{}) error
}, path string, namespaceCreate bool) error {
	var response map[string]interface{}
	if namespaceCreate {
		// The deliberately invalid name cannot create an object. Reaching request
		// validation proves mutation authorization; an authorization failure returns
		// 403 first. POST is non-idempotent and the client therefore never retries it.
		request := map[string]interface{}{
			"metadata": map[string]interface{}{"name": "-"},
		}
		return c.Post(ctx, path, request, &response)
	}
	return c.Get(ctx, path, &response)
}

func precheckLiveCapability(t *testing.T) {
	t.Helper()
	_, namespaceCreate := namespaceCreationCapabilityTests[t.Name()]
	quotaPath, quotaLimited := quotaLimitedCapabilityTests[t.Name()]
	path, ok := capabilityProbePath(t.Name())
	if namespaceCreate {
		path, ok = "/api/web/namespaces", true
	} else if quotaLimited {
		path, ok = quotaPath, true
	}
	if !ok {
		return
	}
	cacheKey := "GET " + path
	mutationProbe := namespaceCreate || quotaLimited
	if mutationProbe {
		cacheKey = "POST " + path
	}
	if cached, found := capabilityProbeCache.Load(cacheKey); found {
		if !cached.(bool) {
			t.Skip("authorized tenant capability preflight returned HTTP 403 for a required surface")
		}
		return
	}

	c, err := GetTestClient()
	if err != nil {
		t.Fatalf("capability preflight could not initialize the API client")
	}
	probeErr := executeCapabilityProbe(context.Background(), c, path, mutationProbe)
	decision := classifyCapabilityProbe(probeErr)
	if quotaLimited {
		decision = classifyQuotaCapabilityProbe(probeErr)
	}
	switch decision {
	case capabilityDenied:
		capabilityProbeCache.Store(cacheKey, false)
		t.Skip("authorized tenant capability preflight returned HTTP 403 for a required surface")
	case capabilityAllowed:
		capabilityProbeCache.Store(cacheKey, true)
	default:
		t.Log("capability preflight was inconclusive; running the acceptance test")
	}
}
