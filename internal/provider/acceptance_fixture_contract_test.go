// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func TestLiveAcceptanceFixtureSchemaContracts(t *testing.T) {
	t.Parallel()

	t.Run("dns compliance required denylist", func(t *testing.T) {
		body := fixtureResourceBody(t, testAccDNSComplianceChecksConfig_basic("fixture-ns", "fixture-dns"), "xcsh_dns_compliance_checks")
		fixtureRequiresAttributes(t, body, "name", "namespace", "domain_denylist")
	})

	t.Run("dns compliance data source resource fixture", func(t *testing.T) {
		body := fixtureResourceBody(t, testAccDnsComplianceChecksDataSourceConfig_basic("fixture-ns", "fixture-dns"), "xcsh_dns_compliance_checks")
		fixtureRequiresAttributes(t, body, "name", "namespace", "domain_denylist")
	})

	t.Run("irule required content", func(t *testing.T) {
		body := fixtureResourceBody(t, testAccIruleConfig_basic("fixture-ns", "fixture-irule"), "xcsh_irule")
		fixtureRequiresAttributes(t, body, "name", "namespace", "description_spec", "irule")
	})

	t.Run("irule data source resource fixture", func(t *testing.T) {
		body := fixtureResourceBody(t, testAccIruleDataSourceConfig_basic("fixture-ns", "fixture-irule"), "xcsh_irule")
		fixtureRequiresAttributes(t, body, "name", "namespace", "description_spec", "irule")
	})

	t.Run("certificate secret nested blocks", func(t *testing.T) {
		body := fixtureResourceBody(t, testAccCertificateDataSourceConfig_basic(
			"fixture-ns",
			"fixture-cert",
			&acctest.TestCertificates{ServerCertBase64: "Y2VydA==", ServerKeyBase64: "a2V5"},
		), "xcsh_certificate")
		fixtureRequiresAttributes(t, body, "name", "namespace", "certificate_url")
		privateKey := fixtureRequiresBlock(t, body, "private_key")
		clearSecret := fixtureRequiresBlock(t, privateKey, "clear_secret_info")
		fixtureRequiresAttributes(t, clearSecret, "url")
		fixtureRequiresBlock(t, body, "disable_ocsp_stapling")
		fixtureForbidsAttributes(t, body, "private_key", "disable_ocsp_stapling")
	})

	t.Run("fast acl rule nested blocks", func(t *testing.T) {
		body := fixtureResourceBody(t, testAccFastAclDataSourceConfig_basic("fixture-fast-acl"), "xcsh_fast_acl")
		siteACL := fixtureRequiresBlock(t, body, "site_acl")
		rule := fixtureRequiresBlock(t, siteACL, "fast_acl_rules")
		action := fixtureRequiresBlock(t, rule, "action")
		fixtureRequiresAttributes(t, action, "simple_action")
		metadata := fixtureRequiresBlock(t, rule, "metadata")
		fixtureRequiresAttributes(t, metadata, "name")
		prefix := fixtureRequiresBlock(t, rule, "prefix")
		fixtureRequiresAttributes(t, prefix, "prefix")
		fixtureForbidsAttributes(t, body, "action", "metadata", "prefix")
	})

	dataGroupFixtures := map[string]string{
		"resource basic":       testAccDataGroupConfig_basicSystem("fixture-data-group"),
		"resource all":         testAccDataGroupConfig_allAttributesSystem("fixture-data-group"),
		"resource labels":      testAccDataGroupConfig_withLabelsSystem("fixture-data-group", map[string]string{"environment": "test"}),
		"resource description": testAccDataGroupConfig_withDescriptionSystem("fixture-data-group", "fixture"),
		"resource annotations": testAccDataGroupConfig_withAnnotationsSystem("fixture-data-group", map[string]string{"owner": "terraform"}),
		"resource records":     testAccDataGroupConfig_withStringRecordsSystem("fixture-data-group"),
		"data source resource": testAccDataGroupDataSourceConfig_basic("fixture-data-group"),
	}
	for name, config := range dataGroupFixtures {
		t.Run("data group "+name, func(t *testing.T) {
			body := fixtureResourceBody(t, config, "xcsh_data_group")
			records := fixtureRequiresBlock(t, body, "string_records")
			fixtureRequiresAttributes(t, records, "records")
		})
	}

	t.Run("app firewall omits server-materialized violations view", func(t *testing.T) {
		body := fixtureResourceBody(t, testAccAppFirewallConfig_detectionSettings("fixture-waf"), "xcsh_app_firewall")
		detection := fixtureRequiresBlock(t, body, "detection_settings")
		fixtureForbidsBlock(t, detection, "violations_view")
	})

	t.Run("fleet supplies live-required defaults", func(t *testing.T) {
		body := fixtureResourceBody(t, testAccFleetConfig_basic("unused", "fixture-fleet"), "xcsh_fleet")
		fixtureRequiresAttributes(t, body, "name", "namespace", "fleet_label", "enable_default_fleet_config_download", "operating_system_version", "volterra_software_version")
	})

	t.Run("udp load balancer supplies live-required defaults", func(t *testing.T) {
		body := fixtureResourceBody(t, testAccUDPLoadBalancerConfig_basicSystem("fixture-udp"), "xcsh_udp_loadbalancer")
		fixtureRequiresAttributes(t, body, "name", "namespace", "dns_volterra_managed", "idle_timeout", "domains", "listen_port")
		fixtureRequiresBlock(t, body, "udp")
	})
}

func fixtureResourceBody(t *testing.T, config, resourceType string) *hclsyntax.Body {
	t.Helper()
	file, diagnostics := hclsyntax.ParseConfig([]byte(config), "acceptance-fixture.tf", hcl.Pos{Line: 1, Column: 1})
	if diagnostics.HasErrors() {
		t.Fatalf("fixture is not valid HCL: %s", diagnostics.Error())
	}
	for _, block := range file.Body.(*hclsyntax.Body).Blocks {
		if block.Type == "resource" && len(block.Labels) == 2 && block.Labels[0] == resourceType {
			return block.Body
		}
	}
	t.Fatalf("fixture has no resource %q", resourceType)
	return nil
}

func fixtureRequiresAttributes(t *testing.T, body *hclsyntax.Body, names ...string) {
	t.Helper()
	for _, name := range names {
		if _, ok := body.Attributes[name]; !ok {
			t.Errorf("fixture is missing required attribute %q", name)
		}
	}
}

func fixtureForbidsAttributes(t *testing.T, body *hclsyntax.Body, names ...string) {
	t.Helper()
	for _, name := range names {
		if _, ok := body.Attributes[name]; ok {
			t.Errorf("fixture uses stale top-level attribute %q", name)
		}
	}
}

func fixtureRequiresBlock(t *testing.T, body *hclsyntax.Body, name string) *hclsyntax.Body {
	t.Helper()
	for _, block := range body.Blocks {
		if block.Type == name {
			return block.Body
		}
	}
	t.Fatalf("fixture is missing required block %q", name)
	return nil
}

func fixtureForbidsBlock(t *testing.T, body *hclsyntax.Body, name string) {
	t.Helper()
	for _, block := range body.Blocks {
		if block.Type == name {
			t.Errorf("fixture uses server-materialized block %q", name)
		}
	}
}
