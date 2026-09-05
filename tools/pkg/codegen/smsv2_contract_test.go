// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"fmt"
	"strings"
	"testing"
)

const syntheticSMSv2V6Contract = `{
  "version":"6.1.0",
  "contract_id":"f5xc-ce-automation/v3",
  "providers":{"aws":{
	"availability":"evidence_backed",
    "capabilities":{"aws_ce_create":"available","runtime_status":"available","site_upgrade":"available","tgw_connect":"available"},
	"unavailable_capabilities":[],
	"telemetry_intake":{
	  "schema_id":"f5xc-smsv2-aws-tgw-telemetry/v2","availability":"available","complete":true,
	  "required_facts":["runtime","gre","bgp","mtu","route","bgp_inside_cidr_block"],
	  "observed_facts":["runtime","gre","bgp","mtu","route","bgp_inside_cidr_block"],
	  "unavailable_facts":[]
	},
    "runtime":{
      "configuration":{"method":"GET","path":"/api/config/namespaces/{namespace}/securemesh_site_v2s/{site}","operation_id":"config.get","response_schema":"securemesh_site_v2GetResponse"},
      "health":{"method":"GET","path":"/api/operate/namespaces/system/sites/{site}/vpm/debug/global/health","operation_id":"health.get","response_schema":"debugHealthResponse"},
      "bgp_peers":{"method":"GET","path":"/api/operate/namespaces/{namespace}/sites/{site}/ver/bgp_peers","operation_id":"bgp.peers","response_schema":"bgpBGPPeersResponse"},
      "bgp_routes":{"method":"GET","path":"/api/operate/namespaces/{namespace}/sites/{site}/ver/bgp_routes","operation_id":"bgp.routes","response_schema":"bgpBGPRoutesResponse"},
      "simplified_routes":{"method":"POST","path":"/api/operate/namespaces/{namespace}/sites/{site}/ver/simplified_routes","operation_id":"routes.simplified","request_schema":"routeSimplifiedRouteRequest","response_schema":"routeSimplifiedRouteResponse"}
    },
    "site_upgrade":{
      "site_status":{"method":"GET","path":"/api/config/namespaces/{namespace}/sites/{site}","operation_id":"site.get","response_schema":"siteGetResponse"},
      "target_discovery":{"method":"GET","path":"/api/maurice/upgradable_sw_versions","operation_id":"upgrade.targets","response_schema":"upgradeTargetsResponse"},
      "precheck":{"method":"GET","path":"/api/maurice/namespaces/{namespace}/sites/{site}/pre_upgrade_check","operation_id":"upgrade.precheck","response_schema":"upgradePrecheckResponse"},
      "upgrade_status":{"method":"GET","path":"/api/maurice/namespaces/{namespace}/sites/{site}/upgrade_status","operation_id":"upgrade.status","response_schema":"upgradeStatusResponse"},
      "software_upgrade":{"method":"POST","path":"/api/config/namespaces/{namespace}/sites/{site}/upgrade_sw","operation_id":"upgrade.sw","request_schema":"upgradeSWRequest","force":false,"semantics":"asynchronous"},
      "os_upgrade":{"method":"POST","path":"/api/config/namespaces/{namespace}/sites/{site}/upgrade_os","operation_id":"upgrade.os","request_schema":"upgradeOSRequest","force":false,"semantics":"asynchronous"},
      "polling":{"transient_failure_values":["UPGRADE_FAILED","FAILED"],"failure_semantics":"transient_until_bounded_timeout","completion":"supplied_targets_installed_and_site_online","timeout_authority":"caller"},
      "redaction":{"exported":"sanitized_status_fields_only","prohibited":["raw_api_messages","node_identifiers","urls"]}
    },
    "authorities":{
      "f5xc":["smsv2_configuration","runtime_health","bgp_peers","bgp_routes","simplified_routes","site_upgrade_observation"],
      "aws":["eni","transit_gateway","transit_gateway_connect","gre_endpoints","bgp_inside_cidrs","autonomous_system_numbers"]
    }
  }}
}`

func TestSMSv2DataSourceTemplatesSelectsCleanBreakSurfaces(t *testing.T) {
	t.Parallel()
	got, err := SMSv2DataSourceTemplates([]byte(syntheticSMSv2V6Contract))
	if err != nil {
		t.Fatal(err)
	}
	want := []SMSv2DataSourceTemplate{{Name: "smsv2_contract", Kind: "contract"}, {Name: "smsv2_aws_runtime", Kind: "runtime"}, {Name: "site_bgp_status", Kind: "convergence"}, {Name: "site_upgrade_status", Kind: "upgrade"}}
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("templates = %#v, want %#v", got, want)
	}
}

func TestSMSv2DataSourceTemplatesRejectsLegacyAndIncompleteContracts(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		from string
		to   string
	}{
		{name: "v2 identity", from: "f5xc-ce-automation/v3", to: "f5xc-ce-automation/v2"},
		{name: "pre-v6.1 release", from: `"version":"6.1.0"`, to: `"version":"6.0.0"`},
		{name: "unavailable runtime", from: `"runtime_status":"available"`, to: `"runtime_status":"unavailable"`},
		{name: "legacy interface endpoint", from: "/securemesh_site_v2s/{site}", to: "/sites/{site}/interface"},
		{name: "legacy routes endpoint", from: "/ver/simplified_routes", to: "/ver/routes"},
		{name: "missing site upgrade capability", from: `"site_upgrade":"available",`, to: ""},
		{name: "mutable upgrade force", from: `"force":false`, to: `"force":true`},
		{name: "authority mismatch", from: `"gre_endpoints","bgp_inside_cidrs"`, to: `"runtime_health","bgp_inside_cidrs"`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fixture := strings.Replace(syntheticSMSv2V6Contract, test.from, test.to, 1)
			if _, err := SMSv2DataSourceTemplates([]byte(fixture)); err == nil {
				t.Fatal("expected contract selection error")
			}
		})
	}
}

func TestSMSv2DataSourceTemplatesRetainsSchemasWhenCapabilitiesFailClosed(t *testing.T) {
	t.Parallel()
	fixture := strings.NewReplacer(
		`"availability":"evidence_backed"`, `"availability":"schema_only"`,
		`"aws_ce_create":"available","runtime_status":"available","site_upgrade":"available","tgw_connect":"available"`, `"aws_ce_create":"unavailable","runtime_status":"unavailable","site_upgrade":"unavailable","tgw_connect":"unavailable"`,
		`"unavailable_capabilities":[]`, `"unavailable_capabilities":["aws_ce_create","runtime_status","site_upgrade","tgw_connect"]`,
		`"availability":"available","complete":true`, `"availability":"unavailable","complete":false`,
	).Replace(syntheticSMSv2V6Contract)
	got, err := SMSv2DataSourceTemplates([]byte(fixture))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 4 {
		t.Fatalf("templates = %#v, want all four fail-closed schemas", got)
	}
}
