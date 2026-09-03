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
			Availability            string                    `json:"availability"`
			Capabilities            map[string]string         `json:"capabilities"`
			UnavailableCapabilities []string                  `json:"unavailable_capabilities"`
			Runtime                 map[string]map[string]any `json:"runtime"`
			Authorities             map[string][]string       `json:"authorities"`
			Telemetry               struct {
				SchemaID         string   `json:"schema_id"`
				Availability     string   `json:"availability"`
				Complete         bool     `json:"complete"`
				RequiredFacts    []string `json:"required_facts"`
				ObservedFacts    []string `json:"observed_facts"`
				UnavailableFacts []string `json:"unavailable_facts"`
			} `json:"telemetry_intake"`
		} `json:"aws"`
	} `json:"providers"`
}

// SMSv2DataSourceTemplates selects the clean-break provider surfaces only from
// a v6-or-newer contract with complete F5 and AWS authority declarations.
func SMSv2DataSourceTemplates(contractJSON []byte) ([]SMSv2DataSourceTemplate, error) {
	var contract smsv2ReleaseContract
	if err := json.Unmarshal(contractJSON, &contract); err != nil {
		return nil, fmt.Errorf("decode SMSv2 contract: %w", err)
	}
	majorText, _, found := strings.Cut(contract.Version, ".")
	major, err := strconv.Atoi(majorText)
	if err != nil || !found || major < 6 || contract.ContractID != "f5xc-ce-automation/v3" {
		return nil, fmt.Errorf("SMSv2 data sources require the v6 clean-break v3 contract")
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
		return nil, fmt.Errorf("SMSv2 v3 runtime endpoint set is incomplete")
	}
	for name, want := range wantRuntime {
		got := contract.Providers.AWS.Runtime[name]
		if got["method"] != want.method || got["path"] != want.path || got["operation_id"] == "" || got["response_schema"] == "" {
			return nil, fmt.Errorf("SMSv2 v3 runtime endpoint %q is incomplete", name)
		}
	}
	wantF5XC := []string{"smsv2_configuration", "runtime_health", "bgp_peers", "bgp_routes", "simplified_routes"}
	wantAWS := []string{"eni", "transit_gateway", "transit_gateway_connect", "gre_endpoints", "bgp_inside_cidrs", "autonomous_system_numbers"}
	if !equalStrings(contract.Providers.AWS.Authorities["f5xc"], wantF5XC) || !equalStrings(contract.Providers.AWS.Authorities["aws"], wantAWS) {
		return nil, fmt.Errorf("SMSv2 v3 authority mapping is incomplete")
	}
	wantFacts := []string{"runtime", "gre", "bgp", "mtu", "route", "bgp_inside_cidr_block"}
	telemetry := contract.Providers.AWS.Telemetry
	if telemetry.SchemaID != "f5xc-smsv2-aws-tgw-telemetry/v2" ||
		!equalStringSets(telemetry.RequiredFacts, wantFacts) ||
		!equalStringSets(telemetry.ObservedFacts, wantFacts) ||
		len(telemetry.UnavailableFacts) != 0 {
		return nil, fmt.Errorf("SMSv2 v3 telemetry declaration is incomplete")
	}
	available := map[string]string{"aws_ce_create": "available", "runtime_status": "available", "tgw_connect": "available"}
	unavailable := map[string]string{"aws_ce_create": "unavailable", "runtime_status": "unavailable", "tgw_connect": "unavailable"}
	switch contract.Providers.AWS.Availability {
	case "evidence_backed":
		if !equalStringMap(contract.Providers.AWS.Capabilities, available) ||
			len(contract.Providers.AWS.UnavailableCapabilities) != 0 ||
			telemetry.Availability != "available" || !telemetry.Complete {
			return nil, fmt.Errorf("SMSv2 evidence-backed capabilities are incoherent")
		}
	case "schema_only":
		if !equalStringMap(contract.Providers.AWS.Capabilities, unavailable) ||
			!equalStringSets(contract.Providers.AWS.UnavailableCapabilities, []string{"aws_ce_create", "runtime_status", "tgw_connect"}) ||
			telemetry.Availability != "unavailable" || telemetry.Complete {
			return nil, fmt.Errorf("SMSv2 schema-only capabilities must fail closed")
		}
	default:
		return nil, fmt.Errorf("SMSv2 AWS availability is unsupported")
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

func equalStringSets(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	values := make(map[string]struct{}, len(left))
	for _, value := range left {
		if _, exists := values[value]; exists {
			return false
		}
		values[value] = struct{}{}
	}
	for _, value := range right {
		if _, exists := values[value]; !exists {
			return false
		}
	}
	return true
}
