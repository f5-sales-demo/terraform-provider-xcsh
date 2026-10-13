// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore
// +build ignore

// This script generates data source acceptance tests for all data sources.
// It uses resource configurations based on OpenAPI specs and existing resource tests.
//
// Features:
//   - resourceConfigs: Custom HCL configuration for each resource's required fields
//   - skipResources: Resources that cannot be tested (permissions, broken APIs, etc.)
//   - systemNamespaceResources: Resources that work in "system" namespace
//
// Run with: go run tools/generate-datasource-tests.go
package main

import (
	"bytes"
	"fmt"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/schema"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// ResourceConfig contains the Terraform configuration for a resource
type ResourceConfig struct {
	Name           string
	StructName     string
	ResourceName   string
	TestFuncName   string
	ConfigFuncName string
	NeedsNamespace bool
	IsNamespace    bool
	// Custom resource configuration block (the HCL inside the resource block)
	ResourceConfig string
	// Whether to skip this resource (e.g., FORBIDDEN errors)
	Skip       bool
	SkipReason string
}

func shouldRegenerateDataSourceTest(content []byte) bool {
	return bytes.Contains(content, []byte("\npackage provider_test\n"))
}

// Resource configurations based on OpenAPI specs
// Key is the resource name (without xcsh_ prefix)
var resourceConfigs = map[string]string{
	"tunnel": `
  tunnel_type = "IPSEC_PSK"
  local_ip {
    ip_address {
      auto = {}
      virtual_network_type { site_local = {} }
    }
  }
  remote_ip {
    ip {
      ipv4 { addr = "192.0.2.1" }
    }
  }
  params {
    ipsec {
      ipsec_psk {
        clear_secret_info { url = "string:///dGVzdC1wc2stc2VjcmV0" }
      }
    }
  }`,

	"subnet": `
  isolated_nw = {}
  site_subnet_params {
    dhcp = {}
    subnet_dhcp_server_params {
      dhcp_networks { network_prefix = "192.0.2.0/24" }
    }
    site {
      name = %[2]q
      namespace = "system"
    }
  }`,
	"site_mesh_group": `
  bfd_disabled = {}
  enable_re_fallback = {}
  virtual_site {
    name = %[2]q
    namespace = %[3]q
  }
  full_mesh {
    data_plane_mesh = {}
  }`,

	"bot_defense_app_infrastructure": `
 environment_type = "PRODUCTION"
 traffic_type = "WEB"
 cloud_hosted {
  infra_host_name = "test.example.com"
  region = "US"
  ingress {
   host_name = "test.example.com"
   location = "AWS_US_EAST_1"
  }
  egress {
   ip_address = "192.0.2.1"
   location = "AWS_US_EAST_1"
  }
 }
`,

	"route": `
  routes {
    disable_location_add = false
    match {
      http_method = "ANY"
      path { prefix = "/" }
      incoming_port { no_port_match = {} }
    }
    route_direct_response {
      response_code = 200
      response_body_encoded = "string:///T0s="
    }
    waf_type { inherit_waf = {} }
  }`,

	"secret_management_access": `
  provider_name = "fixture"`,
	"cloud_elastic_ip": `
  item_count = 1`,
	"cdn_loadbalancer": `
  domains = ["%[1]s.example.com"]
  http { port = 80 }
  origin_pool {
    origin_servers {
      port = 80
      public_name { dns_name = "example.com" }
    }
    no_tls = {}
  }`,
	// Simple resources - just name/namespace
	"alert_policy": "",

	"api_crawler":           "",
	"api_definition":        "\n  swagger_specs = [xcsh_swagger_object.fixture.path]",
	"api_discovery":         "",
	"api_testing":           "",
	"app_api_group":         "",
	"app_firewall":          "",
	"cdn_cache_rule":        "",
	"cluster":               "",
	"code_base_integration": "",
	"dns_compliance_checks": `
  domain_denylist = ["example.com"]`,
	"enhanced_firewall_policy": "",
	"ike1":                     "",
	"ike2":                     "",
	"ike_phase1_profile":       "",
	"ike_phase2_profile":       "",
	"irule": `
  description_spec = "Acceptance fixture"
  irule = "when HTTP_REQUEST { return }"`,
	"malicious_user_mitigation": "",
	"namespace":                 "",
	"rate_limiter":              "",
	"rate_limiter_policy":       "",

	"sensitive_data_policy": "",
	"securemesh_site_v2": `
  baremetal {
    not_managed {}
  }`,

	"trusted_ca_list": "",
	"virtual_k8s":     "",

	// Resources with required fields
	"address_allocator": `
  address_pool = ["10.0.0.0/24"]
  mode = "LOCAL"
  address_allocation_scheme {
    allocation_unit = 0
    local_interface_address_type = "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN"
    local_interface_address_offset = 0
  }`,

	"advertise_policy": `
  port = 80`,

	"alert_receiver": `
  email {
    email = "test@example.com"
  }`,

	"app_setting": `
  app_type_settings {
    app_type_ref {
      name = xcsh_app_type.test.name
      namespace = xcsh_namespace.test.name
    }
  }`,

	"app_type": `
  business_logic_markup_setting {
    enable = {}
  }`,

	"authentication": ``,

	"bgp": `
  bgp_parameters {
    asn          = 65000
    from_site = {}
  }
  where {
    site {
      network_type = "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE"
      ref {
        name = "example-site"
        namespace = "system"
      }
    }
  }`,

	"bgp_asn_set": `
  as_numbers = ["65000"]`,

	"bgp_routing_policy": `
  rules {
    action { allow = {} }
  }`,

	"certificate": `
  certificate_url = "string:///%[2]s"
  private_key {
    clear_secret_info {
      url = "string:///%[3]s"
    }
  }
  disable_ocsp_stapling = {}`,
	"certificate_chain": `
  certificate_url = "string:///%[2]s"`,

	"container_registry": `
  registry          = "registry.example.com"
  user_name         = "testuser"
  email = "test@example.com"
  password {
    clear_secret_info {
      url = "string:///dGVzdC1wYXNzd29yZA=="
    }
  }`,

	"crl": `
  server_address = "crl.example.com"
  server_port = 80
  refresh_interval = 24
  timeout = 10
  http_access { path = "/fixture.crl" }`,

	"data_group": `
  string_records {}`,

	"data_type": `
  is_pii = true
  rules {
    key_pattern {
      substring_value = "test"
    }
  }`,

	"discovery": `
  no_cluster_id = {}
  discovery_consul {
    access_info {
      connection_info { api_server = "http://consul.example.com:8500" }
    }
    publish_info { disable_spec = {} }
  }
  where {
    site {
      network_type = "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE"
      disable_internet_vip = {}
      ref {
        name = %[2]q
        namespace = "system"
      }
    }
  }`,

	"dns_lb_health_check": `
  http {
    path = "/"
  }`,

	"dns_lb_pool": `
  type = "A"
  a_pool {
    members {
      public_ip {
        ip = "192.0.2.1"
      }
    }
  }`,

	"dns_load_balancer": `
  record_type = "A"
  dns_a_record {
    values = ["192.0.2.1"]
  }`,

	"dns_zone": `
  dns_zone_type = "primary"
  primary {
    default_rr_set_group = []
  }`,

	"endpoint": `
  ip = "192.0.2.10"
  port = 8080
  protocol = "TCP"`,

	"fast_acl": `
  re_acl {}`,

	"fast_acl_rule": `
  action { simple_action = "DENY" }
  ip_prefix_set {}`,

	"filter_set": `
  context_key = "dashboard"
  filter_fields {
    field_id = "test-field"
    string_field {
      field_values = ["test-value"]
    }
  }`,

	"fleet": `
  fleet_label = "%[1]s"
  operating_system_version = "default"
  volterra_software_version = "default"
  enable_default_fleet_config_download = false`,

	"forward_proxy_policy": `
  any_proxy = {}
  allow_all = {}`,

	"forwarding_class": `
  queue_id_to_use = "DSCP_BEST_EFFORT"
  no_policer = {}
  interface_group = "ANY_AVAILABLE_INTERFACE"
  tos_value = 40`,

	"geo_location_set": ``,

	"global_log_receiver": `
  http_receiver {
    uri = "http://logs.example.com/fixture"
    auth_none = {}
    no_tls = {}
  }
  ns_current = {}
  audit_logs = {}`,

	"healthcheck": `
  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5
  tcp_health_check {}`,

	"http_loadbalancer": `
  domains = ["test.example.com"]
  http {
    dns_volterra_managed = false
  }
  default_route_pools {
    pool {
      name      = xcsh_origin_pool.test.name
      namespace = xcsh_origin_pool.test.namespace
    }
  }`,

	"ip_prefix_set": `
  ipv4_prefixes {
    ipv4_prefix      = "10.0.0.0/8"
    description_spec = "Test prefix"
  }`,

	"log_receiver": `
  site_local {}`,

	"nat_policy": `
  site {
    refs {
      name = %[2]q
      namespace = "system"
    }
  }
  rules {
    name = "fixture"
    enable = {}
    segment {
      refs {
        name = xcsh_segment.fixture.name
        namespace = xcsh_segment.fixture.namespace
      }
    }
    criteria { any = {} }
    action { virtual_cidr = "192.0.2.0/24" }
  }`,

	"network_connector": `
  disable_forward_proxy = {}
  sli_to_slo_snat {
    default_gw_snat {}
  }`,

	"network_firewall": `
  active_forward_proxy_policies {
    forward_proxy_policies {
      name      = xcsh_forward_proxy_policy.test.name
      namespace = xcsh_forward_proxy_policy.test.namespace
    }
  }`,

	"network_interface": `
  ethernet_interface {
    device = "eth0"
    mtu = 0
    priority = 0
    cluster = {}
    untagged = {}
    dhcp_client = {}
    site_local_inside_network = {}
    monitor_disabled = {}
    not_primary = {}
    no_ipv6_address = {}
  }`,

	"network_policy": `
  endpoint { any = {} }
  rules {
    ingress_rules {
      action = "ALLOW"
      any = {}
      all_traffic = {}
      metadata { name = "allow-all" }
    }
  }`,

	"network_policy_rule": `
  action = "DENY"
  protocol = "TCP"`,

	"network_policy_view": `
  endpoint { any = {} }
  ingress_rules {
    action = "ALLOW"
    any = {}
    all_traffic = {}
    metadata { name = "allow-ingress" }
  }
  egress_rules {
    action = "ALLOW"
    any = {}
    all_traffic = {}
    metadata { name = "allow-egress" }
  }`,

	"origin_pool": `
  port = 443
  origin_servers {
    public_name { dns_name = "example.com" }
  }
  no_tls = {}`,

	"policer": `
  committed_information_rate = 10000
  burst_size                 = 5000`,

	"policy_based_routing": `
  forwarding_class_list {
    name = xcsh_forwarding_class.fixture.name
    namespace = xcsh_forwarding_class.fixture.namespace
  }`,

	"protocol_inspection": `
  enable_disable_compliance_checks {
    disable_compliance_checks = {}
  }
  enable_disable_signatures {
    disable_signature = {}
  }`,

	"protocol_policer": `
  protocol_policer {
    protocol { dns = {} }
    policer {
      name = xcsh_policer.test.name
      namespace = "system"
    }
  }`,

	"proxy": `
  http_proxy {
    enable_http {}
  }
  no_forward_proxy_policy = {}
  site_local_network = {}
  do_not_advertise = {}`,

	"service_policy": `
  allow_all_requests = {}
  any_server = {}`,

	"service_policy_rule": `
  action = "ALLOW"
  waf_action { none = {} }`,

	"tcp_loadbalancer": `
  domains = ["%[1]s.example.com"]
  tcp = {}
  listen_port = 80
  no_sni = {}
  no_service_policies = {}
  do_not_advertise = {}
  origin_pools_weights {
    pool {
      name = xcsh_origin_pool.test.name
      namespace = xcsh_origin_pool.test.namespace
    }
    weight = 1
  }`,

	"token": `
  type = 0`,

	"udp_loadbalancer": `
  domains     = ["%[1]s.example.com"]
  listen_port = 53
  dns_volterra_managed = false
  idle_timeout = 30000
  origin_pools_weights {
    pool {
      name      = xcsh_origin_pool.test.name
      namespace = xcsh_origin_pool.test.namespace
    }
    weight = 1
  }
  do_not_advertise = {}`,

	"usb_policy": `
  allowed_devices {
    id_vendor = "0x046d"
    id_product = "0xc52b"
  }`,

	"user_identification": `
  rules { client_ip = {} }`,

	"virtual_host": `
  domains = ["test.example.com"]
  proxy   = "SMA_PROXY"`,

	"virtual_network": `
  global_network = {}`,

	"virtual_site": `
  site_type = "CUSTOMER_EDGE"
  site_selector {
    expressions = ["site-type=ce"]
  }`,

	"waf_exclusion_policy": `
  waf_exclusion_rules {
    metadata { name = "test-rule" }
  }`,

	"workload": `
  simple_service {
    do_not_advertise = {}
    disabled = {}
    container {
      name = "fixture"
      default_flavor = {}
      image {
        name = "nginx:stable"
        public = {}
      }
    }
  }`,

	"workload_flavor": ``,
}

// Resources to skip (FORBIDDEN errors, special permissions required)
var skipResources = map[string]string{
	"secret_management_access": "External secret-management authentication is required",

	"cloud_credentials": "Requires system-level namespace access",
	"cminstance":        "Requires CM instance management permissions",

	"registration":         "Requires site registration permissions",
	"tenant_configuration": "API returns 500 errors",
}

// Resources that need special dependencies in the config
var resourceDependencies = map[string]string{
	"policy_based_routing": `
resource "xcsh_forwarding_class" "fixture" {
 name = "%[1]s-fc"
 namespace = "system"
 queue_id_to_use = "DSCP_BEST_EFFORT"
 no_policer = {}
 interface_group = "ANY_AVAILABLE_INTERFACE"
}`,

	"nat_policy": `
resource "xcsh_segment" "fixture" {
 name = "%[1]s-segment"
 namespace = "system"
}`,

	"api_definition": `
resource "xcsh_swagger_object" "fixture" {
 name = "%[1]s-spec"
 namespace = "shared"
 content = "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Acceptance API\",\"version\":\"1.0.0\"},\"servers\":[{\"url\":\"https://example.com\"}],\"paths\":{\"/health\":{\"get\":{\"responses\":{\"200\":{\"description\":\"OK\"}}}}}}"
}
`,
	"protocol_policer": `
resource "xcsh_policer" "test" {
 name = "%[1]s-policer"
 namespace = "system"
 burst_size = 1000
 committed_information_rate = 1000
}
`,

	"app_setting": `
resource "xcsh_app_type" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = "${%[2]q}-apptype"
  namespace  = xcsh_namespace.test.name
  business_logic_markup_setting {
    enable = {}
  }
}
`,
	"http_loadbalancer": `
resource "xcsh_origin_pool" "test" {
  name       = "%[1]s-pool"
  namespace  = "system"
  origin_servers {
    public_ip {
      ip = "192.0.2.1"
    }
  }
  port               = 80
  endpoint_selection = "LOCAL_PREFERRED"
  loadbalancer_algorithm = "ROUND_ROBIN"
}
`,
	"tcp_loadbalancer": `
resource "xcsh_origin_pool" "test" {
  name       = "%[1]s-pool"
  namespace  = "system"
  origin_servers {
    public_ip {
      ip = "192.0.2.1"
    }
  }
  port               = 80
  endpoint_selection = "LOCAL_PREFERRED"
  loadbalancer_algorithm = "ROUND_ROBIN"
}
`,
	"udp_loadbalancer": `
resource "xcsh_origin_pool" "test" {
  name       = "%[1]s-pool"
  namespace  = "system"
  origin_servers {
    public_ip {
      ip = "192.0.2.1"
    }
  }
  port               = 53
  endpoint_selection = "LOCAL_PREFERRED"
  loadbalancer_algorithm = "ROUND_ROBIN"
}
`,
	"network_firewall": `
resource "xcsh_forward_proxy_policy" "test" {
  name = "%[1]s-forward-proxy"
  namespace = "system"
  any_proxy = {}
  allow_all = {}
}
`,
}

// Resources that don't need a namespace (they're system-level)
var systemLevelResources = map[string]bool{
	"namespace": true,
}

// Resources that MUST use namespace = "system" (not custom namespace)
var systemNamespaceResources = map[string]bool{
	"network_interface":    true,
	"discovery":            true,
	"nat_policy":           true,
	"network_connector":    true,
	"tcp_loadbalancer":     true,
	"tunnel":               true,
	"waf_exclusion_policy": true,
	"global_log_receiver":  true,
	"crl":                  true,
	"api_crawler":          true,
	"network_firewall":     true,
	"forward_proxy_policy": true,
	"token":                true,
	"site_mesh_group":      true,
	"address_allocator":    true,
	"forwarding_class":     true,
	"protocol_policer":     true,
	"route":                true,
	"udp_loadbalancer":     true,
	"fast_acl":             true,
	"fast_acl_rule":        true,
	"dc_cluster_group":     true,
	"network_policy_view":  true,
	"segment":              true,
	"virtual_network":      true,
	"policer":              true,
	"user_identification":  true,
	"bgp_asn_set":          true,
	"geo_location_set":     true,
	"healthcheck":          true,
	"rate_limiter":         true,
	"service_policy":       true,
	"data_group":           true,
	"data_type":            true,
	"filter_set":           true,
	"origin_pool":          true,
	"app_firewall":         true,
	"securemesh_site_v2":   true,
}

func toStructName(name string) string {
	parts := strings.Split(name, "_")
	var result string
	for _, part := range parts {
		if len(part) > 0 {
			result += strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return result
}

func main() {
	specDir := os.Getenv("XCSH_SPEC_DIR")
	if specDir == "" {
		specDir = "docs/specifications/api"
	}
	if err := openapi.RequireSpecs(specDir); err != nil {
		panic(err)
	}
	fixtureNamespaces := map[string]string{}
	domains, err := openapi.FindDomainSpecFiles(specDir)
	if err != nil {
		panic(err)
	}
	for _, path := range domains {
		domain, err := openapi.ExtractResourcesFromDomain(path)
		if err != nil {
			panic(err)
		}
		for _, resource := range domain.Resources {
			if _, err := os.Stat(filepath.Join("internal/provider", resource.Name+"_resource.go")); os.IsNotExist(err) {
				continue
			}
			ns, err := schema.AcceptanceReusableFixtureNamespace(domain.Spec, resource.Name)
			if err != nil {
				panic(err)
			}
			if resource.Name == "subnet" {
				ns = "system"
			}
			if systemNamespaceResources[resource.Name] {
				ns = "system"
			}
			if ns != "" {
				fixtureNamespaces[resource.Name] = ns
			}
		}
	}
	// Template for the test file
	const testTemplate = `// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

{{if .Skip}}
import (
	"testing"
 "github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func {{.TestFuncName}}(t *testing.T) {
	{{if eq .Name "tenant_configuration"}}acctest.RequireNonDisposableScope(t)
{{else if eq .Name "cloud_credentials"}}acctest.RequireExternalCredential(t, "cloud")
{{else if eq .Name "cminstance"}}acctest.RequireExternalCredential(t, "central_manager")
{{else if eq .Name "secret_management_access"}}acctest.RequireExternalCredential(t, "secret_management")
{{else}}acctest.FailHistoricalPrerequisite(t)
{{end}}
}
{{else}}
import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
)

func {{.TestFuncName}}(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
 {{if eq .Name "site_mesh_group"}}acctest.RequireReadOnlyVirtualSite(t){{else if or (eq .Name "subnet") (eq .Name "nat_policy") (eq .Name "discovery")}}acctest.RequireExistingInfrastructure(t,"site"){{end}}

	{{if eq .Name "api_definition"}}rName := acctest.CompactRandomName(20){{else}}rName := acctest.RandomName("tf-acc-test"){{end}}
{{- if and .NeedsNamespace (not .IsNamespace) (not .UseSystemNamespace)}}
	nsName := acctest.RandomName("tf-acc-test-ns")
{{- end}}
	resourceName := "{{.ResourceName}}.test"
	dataSourceName := "data.{{.ResourceName}}.test"

{{- if .RunSequential}}
	testCase := resource.TestCase{
{{- else}}
	resource.ParallelTest(t, resource.TestCase{
{{- end}}
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
{{- if and .NeedsNamespace (not .IsNamespace) (not .UseSystemNamespace)}}
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
{{- end}}
		Steps: []resource.TestStep{
			{
{{- if .IsNamespace}}
				Config: {{.ConfigFuncName}}(rName),
{{- else if .UseSystemNamespace}}
				Config: {{.ConfigFuncName}}(rName{{if eq .Name "site_mesh_group"}}, acctest.ReadOnlyVirtualSiteName(), acctest.ReadOnlyVirtualSiteNamespace(){{else if or (eq .Name "subnet") (eq .Name "nat_policy") (eq .Name "discovery")}}, acctest.ReadOnlySiteName(){{end}}),
{{- else if .NeedsNamespace}}
				Config: {{.ConfigFuncName}}(nsName, rName),
{{- else}}
				Config: {{.ConfigFuncName}}(rName),
{{- end}}
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
{{- if and .NeedsNamespace (not .IsNamespace)}}
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
{{- end}}
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
{{- if .RunSequential}}
	}
	acctest.RunWithMockOrReal(t, testCase, nil)
{{- else}}
	})
{{- end}}
}
{{end}}
{{- if .IsNamespace}}

func {{.ConfigFuncName}}(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(` + "`" + `
resource "{{.ResourceName}}" "test" {
  name = %[1]q
}

data "{{.ResourceName}}" "test" {
  name      = {{.ResourceName}}.test.name
  namespace = "system"
}
` + "`" + `, name))
}
{{- else if .UseSystemNamespace}}

func {{.ConfigFuncName}}(name string{{if eq .Name "site_mesh_group"}}, virtualSiteName, virtualSiteNamespace string{{else if or (eq .Name "subnet") (eq .Name "nat_policy") (eq .Name "discovery")}}, siteName string{{end}}) string {
{{- if or (eq .Name "certificate") (eq .Name "certificate_chain")}}
 certs := acctest.MustGenerateTestCertificates()
{{- end}}
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(` + "`" + `
{{.Dependencies}}
resource "{{.ResourceName}}" "test" {
  name      = %[1]q
  namespace = {{if .FixedNamespace}}{{printf "%q" .FixedNamespace}}{{else}}"system"{{end}}{{.ResourceConfig}}
}

data "{{.ResourceName}}" "test" {
  depends_on = [{{.ResourceName}}.test]
  name       = {{.ResourceName}}.test.name
  namespace  = {{.ResourceName}}.test.namespace
}
` + "`" + `, name{{if eq .Name "site_mesh_group"}}, virtualSiteName, virtualSiteNamespace{{else if or (eq .Name "subnet") (eq .Name "nat_policy") (eq .Name "discovery")}}, siteName{{end}}{{if eq .Name "certificate"}}, certs.ServerCertBase64, certs.ServerKeyBase64{{else if eq .Name "certificate_chain"}}, certs.IntermediateCABase64{{end}}))
}
{{- else if and .NeedsNamespace (not .Skip)}}

func {{.ConfigFuncName}}(nsName, name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(` + "`" + `
resource "xcsh_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [xcsh_namespace.test]
  create_duration = "5s"
}
{{.Dependencies}}
resource "{{.ResourceName}}" "test" {
  depends_on = [time_sleep.wait_for_namespace{{.DependsOnExtra}}]
  name       = %[2]q
  namespace  = xcsh_namespace.test.name{{.ResourceConfig}}
}

data "{{.ResourceName}}" "test" {
  depends_on = [{{.ResourceName}}.test]
  name       = {{.ResourceName}}.test.name
  namespace  = {{.ResourceName}}.test.namespace
}
` + "`" + `, nsName, name))
}
{{- else if not .Skip}}

func {{.ConfigFuncName}}(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(` + "`" + `
resource "{{.ResourceName}}" "test" {
  name = %[1]q
}

data "{{.ResourceName}}" "test" {
  name      = {{.ResourceName}}.test.name
  namespace = "system"
}
` + "`" + `, name))
}
{{- end}}
`

	// Find all data source test files
	dataSourceDir := "internal/provider"
	pattern := filepath.Join(dataSourceDir, "*_data_source_test.go")
	files, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Printf("Error finding test files: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d data source test files to fix\n", len(files))

	fixed := 0
	skipped := 0

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", file, err)
			continue
		}
		if !shouldRegenerateDataSourceTest(content) {
			fmt.Printf("Preserving hand-maintained %s\n", file)
			continue
		}

		baseName := filepath.Base(file)
		resourceName := strings.TrimSuffix(baseName, "_data_source_test.go")

		// Check if we should skip this resource
		if reason, skip := skipResources[resourceName]; skip {
			fmt.Printf("Marking %s as skipped: %s\n", resourceName, reason)
			// Continue to regenerate with skip flag
		}

		structName := toStructName(resourceName)

		config := ResourceConfig{
			Name:           resourceName,
			StructName:     structName,
			ResourceName:   "xcsh_" + resourceName,
			TestFuncName:   "TestAcc" + structName + "DataSource_basic",
			ConfigFuncName: "testAcc" + structName + "DataSourceConfig_basic",
			NeedsNamespace: !systemLevelResources[resourceName],
			IsNamespace:    resourceName == "namespace",
		}

		// Check for skip
		if reason, skip := skipResources[resourceName]; skip {
			config.Skip = true
			config.SkipReason = reason
		}

		// Get resource-specific config
		if rcfg, exists := resourceConfigs[resourceName]; exists && rcfg != "" {
			config.ResourceConfig = rcfg
		}

		// Add dependencies
		deps := ""
		dependsOnExtra := ""
		if dep, exists := resourceDependencies[resourceName]; exists {
			deps = dep
			// Extract resource type from dependency
			if strings.Contains(dep, "xcsh_origin_pool") {
				dependsOnExtra = ", xcsh_origin_pool.test"
			} else if strings.Contains(dep, "xcsh_app_type") {
				dependsOnExtra = ", xcsh_app_type.test"
			} else if strings.Contains(dep, "xcsh_forward_proxy_policy") {
				dependsOnExtra = ", xcsh_forward_proxy_policy.test"
			}
		}

		// Create template data without embedding
		type templateData struct {
			Name               string
			StructName         string
			ResourceName       string
			TestFuncName       string
			ConfigFuncName     string
			NeedsNamespace     bool
			IsNamespace        bool
			UseSystemNamespace bool
			FixedNamespace     string
			RunSequential      bool
			ResourceConfig     string
			Skip               bool
			SkipReason         string
			Dependencies       string
			DependsOnExtra     string
		}

		if ns := fixtureNamespaces[resourceName]; ns != "" {
			config.ResourceConfig = strings.ReplaceAll(config.ResourceConfig, "xcsh_namespace.test.name", fmt.Sprintf("%q", ns))
			deps = strings.ReplaceAll(deps, "xcsh_namespace.test.name", fmt.Sprintf("%q", ns))
		}
		data := templateData{
			Name:               config.Name,
			StructName:         config.StructName,
			ResourceName:       config.ResourceName,
			TestFuncName:       config.TestFuncName,
			ConfigFuncName:     config.ConfigFuncName,
			NeedsNamespace:     config.NeedsNamespace,
			IsNamespace:        config.IsNamespace,
			UseSystemNamespace: systemNamespaceResources[resourceName] || fixtureNamespaces[resourceName] != "",
			FixedNamespace:     fixtureNamespaces[resourceName],
			RunSequential:      resourceName == "securemesh_site_v2",
			ResourceConfig:     config.ResourceConfig,
			Skip:               config.Skip,
			SkipReason:         config.SkipReason,
			Dependencies:       deps,
			DependsOnExtra:     dependsOnExtra,
		}

		tmpl, err := template.New("test").Parse(testTemplate)
		if err != nil {
			fmt.Printf("Error parsing template: %v\n", err)
			continue
		}

		// Execute template to buffer first
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, data)
		if err != nil {
			fmt.Printf("Error executing template for %s: %v\n", file, err)
			continue
		}

		// Format the generated code with gofmt
		formatted, err := format.Source(buf.Bytes())
		if err != nil {
			// If formatting fails, write unformatted code with warning
			fmt.Printf("⚠️  gofmt failed for %s: %v (writing unformatted)\n", file, err)
			formatted = buf.Bytes()
		}

		// Write formatted content to file
		err = os.WriteFile(file, formatted, 0644)
		if err != nil {
			fmt.Printf("Error writing %s: %v\n", file, err)
			continue
		}

		if config.Skip {
			skipped++
		} else {
			fixed++
		}
		fmt.Printf("Fixed %s\n", file)
	}

	fmt.Printf("\nSummary: Fixed %d test files, skipped %d\n", fixed, skipped)
}
