// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"fmt"
	"strings"
	"testing"
)

const syntheticSMSv2V5Contract = `{
  "version":"5.0.0",
  "contract_id":"f5xc-ce-automation/v2",
  "providers":{"aws":{
    "capabilities":{"aws_ce_create":"available","runtime_status":"available","tgw_connect":"available"},
    "runtime":{
      "configuration":{"method":"GET","path":"/api/config/namespaces/{namespace}/securemesh_site_v2s/{site}","operation_id":"config.get","response_schema":"securemesh_site_v2GetResponse"},
      "health":{"method":"GET","path":"/api/operate/namespaces/system/sites/{site}/vpm/debug/global/health","operation_id":"health.get","response_schema":"debugHealthResponse"},
      "bgp_peers":{"method":"GET","path":"/api/operate/namespaces/{namespace}/sites/{site}/ver/bgp_peers","operation_id":"bgp.peers","response_schema":"bgpBGPPeersResponse"},
      "bgp_routes":{"method":"GET","path":"/api/operate/namespaces/{namespace}/sites/{site}/ver/bgp_routes","operation_id":"bgp.routes","response_schema":"bgpBGPRoutesResponse"},
      "simplified_routes":{"method":"POST","path":"/api/operate/namespaces/{namespace}/sites/{site}/ver/simplified_routes","operation_id":"routes.simplified","request_schema":"routeSimplifiedRouteRequest","response_schema":"routeSimplifiedRouteResponse"}
    },
    "authorities":{
      "f5xc":["smsv2_configuration","runtime_health","bgp_peers","bgp_routes","simplified_routes"],
      "aws":["eni","transit_gateway","transit_gateway_connect","gre_endpoints","bgp_inside_cidrs"]
    }
  }}
}`

func TestSMSv2DataSourceTemplatesSelectsCleanBreakSurfaces(t *testing.T) {
	t.Parallel()
	got, err := SMSv2DataSourceTemplates([]byte(syntheticSMSv2V5Contract))
	if err != nil {
		t.Fatal(err)
	}
	want := []SMSv2DataSourceTemplate{{Name: "smsv2_contract", Kind: "contract"}, {Name: "smsv2_aws_runtime", Kind: "runtime"}, {Name: "site_bgp_status", Kind: "convergence"}}
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
		{name: "v1 identity", from: "f5xc-ce-automation/v2", to: "f5xc-ce-automation/v1"},
		{name: "pre-v5 release", from: `"version":"5.0.0"`, to: `"version":"4.9.9"`},
		{name: "unavailable runtime", from: `"runtime_status":"available"`, to: `"runtime_status":"unavailable"`},
		{name: "legacy interface endpoint", from: "/securemesh_site_v2s/{site}", to: "/sites/{site}/interface"},
		{name: "legacy routes endpoint", from: "/ver/simplified_routes", to: "/ver/routes"},
		{name: "authority mismatch", from: `"gre_endpoints","bgp_inside_cidrs"`, to: `"runtime_health","bgp_inside_cidrs"`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fixture := strings.Replace(syntheticSMSv2V5Contract, test.from, test.to, 1)
			if _, err := SMSv2DataSourceTemplates([]byte(fixture)); err == nil {
				t.Fatal("expected contract selection error")
			}
		})
	}
}
