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
var capabilityProbeMu sync.Mutex

type quotaCapabilityFixture struct {
	path string
	spec map[string]interface{}
}

var quotaLimitedCapabilityTests = map[string]quotaCapabilityFixture{
	"TestAccAPICrawlerResource_basic": {
		path: "/api/config/namespaces/system/api_crawlers",
		spec: map[string]interface{}{
			"domains": []interface{}{map[string]interface{}{"domain": "example.com"}},
		},
	},
	"TestAccFleetDataSource_basic": {
		path: "/api/config/namespaces/system/fleets",
		spec: map[string]interface{}{
			"fleet_label":                          "tf-acc-test-capability",
			"enable_default_fleet_config_download": false,
			"operating_system_version":             "default",
			"volterra_software_version":            "default",
		},
	},
	"TestAccFleetResource_basic": {
		path: "/api/config/namespaces/system/fleets",
		spec: map[string]interface{}{
			"fleet_label":                          "tf-acc-test-capability",
			"enable_default_fleet_config_download": false,
			"operating_system_version":             "default",
			"volterra_software_version":            "default",
		},
	},
	"TestAccMaliciousUserMitigationResource_captchaChallenge": maliciousUserMitigationQuotaFixture(),
	"TestAccMaliciousUserMitigationResource_fullLifecycle":    maliciousUserMitigationQuotaFixture(),
	"TestAccMaliciousUserMitigationResource_jsChallenge":      maliciousUserMitigationQuotaFixture(),
	"TestAccMaliciousUserMitigationResource_switchAction":     maliciousUserMitigationQuotaFixture(),
}

func maliciousUserMitigationQuotaFixture() quotaCapabilityFixture {
	return quotaCapabilityFixture{
		path: "/api/config/namespaces/system/malicious_user_mitigations",
		spec: map[string]interface{}{
			"mitigation_type": map[string]interface{}{
				"rules": []interface{}{map[string]interface{}{
					"threat_level":      map[string]interface{}{"high": map[string]interface{}{}},
					"mitigation_action": map[string]interface{}{"block_temporarily": map[string]interface{}{}},
				}},
			},
		},
	}
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

func executeQuotaCapabilityProbe(ctx context.Context, c interface {
	Post(context.Context, string, interface{}, interface{}) error
	Delete(context.Context, string) error
}, fixture quotaCapabilityFixture, name string) (probeErr, cleanupErr error) {
	spec := make(map[string]interface{}, len(fixture.spec))
	for key, value := range fixture.spec {
		spec[key] = value
	}
	if _, ok := spec["fleet_label"]; ok {
		spec["fleet_label"] = name
	}
	request := map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":      name,
			"namespace": "system",
		},
		"spec": spec,
	}
	var response map[string]interface{}
	if err := c.Post(ctx, fixture.path, request, &response); err != nil {
		return err, nil
	}
	return nil, c.Delete(ctx, fixture.path+"/"+name)
}

func precheckLiveCapability(t *testing.T) {
	t.Helper()
	_, namespaceCreate := namespaceCreationCapabilityTests[t.Name()]
	quotaFixture, quotaLimited := quotaLimitedCapabilityTests[t.Name()]
	path, ok := capabilityProbePath(t.Name())
	if namespaceCreate {
		path, ok = "/api/web/namespaces", true
	} else if quotaLimited {
		path, ok = quotaFixture.path, true
	}
	if !ok {
		return
	}
	cacheKey := "GET " + path
	mutationProbe := namespaceCreate || quotaLimited
	if mutationProbe {
		cacheKey = "POST " + path
	}
	capabilityProbeMu.Lock()
	defer capabilityProbeMu.Unlock()
	if cached, found := capabilityProbeCache.Load(cacheKey); found {
		if !cached.(bool) {
			t.Skip("authorized tenant capability preflight proved the required surface unavailable")
		}
		return
	}

	c, err := GetTestClient()
	if err != nil {
		t.Fatalf("capability preflight could not initialize the API client")
	}
	var probeErr error
	if quotaLimited {
		probeName := RandomName("tf-acc-test-capability")
		var cleanupErr error
		probeErr, cleanupErr = executeQuotaCapabilityProbe(context.Background(), c, quotaFixture, probeName)
		if cleanupErr != nil {
			t.Fatalf("capability preflight created its disposable object but could not remove it")
		}
	} else {
		probeErr = executeCapabilityProbe(context.Background(), c, path, mutationProbe)
	}
	decision := classifyCapabilityProbe(probeErr)
	if quotaLimited {
		decision = classifyQuotaCapabilityProbe(probeErr)
	}
	switch decision {
	case capabilityDenied:
		capabilityProbeCache.Store(cacheKey, false)
		t.Skip("authorized tenant capability preflight proved the required surface unavailable")
	case capabilityAllowed:
		capabilityProbeCache.Store(cacheKey, true)
	default:
		t.Log("capability preflight was inconclusive; running the acceptance test")
	}
}
