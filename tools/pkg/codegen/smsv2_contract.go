// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// SMSv2DataSourceTemplate describes one provider-owned composite data source
// enabled by the immutable API release contract.
type SMSv2DataSourceTemplate struct {
	Name string
	Kind string
}

type smsv2ReleaseContract struct {
	Version    string `json:"version"`
	ContractID string `json:"contract_id"`
	Providers  struct {
		AWS struct {
			Capabilities map[string]string            `json:"capabilities"`
			Runtime      map[string]map[string]string `json:"runtime"`
			Authorities  map[string][]string          `json:"authorities"`
		} `json:"aws"`
	} `json:"providers"`
}

// SMSv2DataSourceTemplates selects the clean-break provider surfaces only from
// a v5-or-newer contract with complete F5 and AWS authority declarations.
func SMSv2DataSourceTemplates(contractJSON []byte) ([]SMSv2DataSourceTemplate, error) {
	var contract smsv2ReleaseContract
	if err := json.Unmarshal(contractJSON, &contract); err != nil {
		return nil, fmt.Errorf("decode SMSv2 contract: %w", err)
	}
	majorText, _, found := strings.Cut(contract.Version, ".")
	major, err := strconv.Atoi(majorText)
	if err != nil || !found || major < 5 || contract.ContractID != "f5xc-ce-automation/v2" {
		return nil, fmt.Errorf("SMSv2 data sources require the v5 clean-break v2 contract")
	}
	wantCapabilities := map[string]string{"aws_ce_create": "available", "runtime_status": "available", "tgw_connect": "available"}
	if !equalStringMap(contract.Providers.AWS.Capabilities, wantCapabilities) {
		return nil, fmt.Errorf("SMSv2 v2 capabilities are incomplete")
	}
	wantRuntime := map[string]struct {
		method string
		path   string
	}{
		"configuration":     {"GET", "/api/config/namespaces/{namespace}/securemesh_site_v2s/{site}"},
		"health":            {"GET", "/api/operate/namespaces/system/sites/{site}/vpm/debug/global/health"},
		"bgp_peers":         {"GET", "/api/operate/namespaces/{namespace}/sites/{site}/ver/bgp_peers"},
		"bgp_routes":        {"GET", "/api/operate/namespaces/{namespace}/sites/{site}/ver/bgp_routes"},
		"simplified_routes": {"POST", "/api/operate/namespaces/{namespace}/sites/{site}/ver/simplified_routes"},
	}
	if len(contract.Providers.AWS.Runtime) != len(wantRuntime) {
		return nil, fmt.Errorf("SMSv2 v2 runtime endpoint set is incomplete")
	}
	for name, want := range wantRuntime {
		got := contract.Providers.AWS.Runtime[name]
		if got["method"] != want.method || got["path"] != want.path || got["operation_id"] == "" || got["response_schema"] == "" {
			return nil, fmt.Errorf("SMSv2 v2 runtime endpoint %q is incomplete", name)
		}
	}
	wantF5XC := []string{"smsv2_configuration", "runtime_health", "bgp_peers", "bgp_routes", "simplified_routes"}
	wantAWS := []string{"eni", "transit_gateway", "transit_gateway_connect", "gre_endpoints", "bgp_inside_cidrs"}
	if !equalStrings(contract.Providers.AWS.Authorities["f5xc"], wantF5XC) || !equalStrings(contract.Providers.AWS.Authorities["aws"], wantAWS) {
		return nil, fmt.Errorf("SMSv2 v2 authority mapping is incomplete")
	}
	return []SMSv2DataSourceTemplate{
		{Name: "smsv2_contract", Kind: "contract"},
		{Name: "smsv2_aws_runtime", Kind: "runtime"},
		{Name: "site_bgp_status", Kind: "convergence"},
	}, nil
}

func equalStringMap(left, right map[string]string) bool {
	if len(left) != len(right) {
		return false
	}
	for key, value := range right {
		if left[key] != value {
			return false
		}
	}
	return true
}

func equalStrings(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range right {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}
