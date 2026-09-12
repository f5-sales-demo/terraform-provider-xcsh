// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SMSv2DataSourceTemplate describes one provider-owned composite data source
// enabled by the immutable API release contract.
type SMSv2DataSourceTemplate struct {
	Name string
	Kind string
}

type smsv2CapabilitySource struct {
	Repository  string   `json:"repository"`
	Commit      string   `json:"commit"`
	AssetPath   string   `json:"asset_path"`
	AssetSHA256 string   `json:"asset_sha256"`
	SchemaPaths []string `json:"schema_paths"`
}

type smsv2CapabilityBoundary struct {
	Availability              string                `json:"availability"`
	Enforcement               string                `json:"enforcement"`
	Reason                    string                `json:"reason"`
	Source                    smsv2CapabilitySource `json:"source"`
	FutureMappingRequirements map[string]string     `json:"future_mapping_requirements"`
}

type smsv2ReleaseContract struct {
	Version    string `json:"version"`
	ContractID string `json:"contract_id"`
	Providers  struct {
		Azure struct {
			Availability            string                    `json:"availability"`
			RouteServerEBGPMultihop smsv2CapabilityBoundary   `json:"route_server_ebgp_multihop"`
			Runtime                 map[string]map[string]any `json:"runtime"`
		} `json:"azure"`
		AWS struct {
			Availability            string                    `json:"availability"`
			Capabilities            map[string]string         `json:"capabilities"`
			UnavailableCapabilities []string                  `json:"unavailable_capabilities"`
			Runtime                 map[string]map[string]any `json:"runtime"`
			SiteUpgrade             map[string]any            `json:"site_upgrade"`
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
// the v7 API contract with complete F5, AWS, and Azure declarations.
func SMSv2DataSourceTemplates(contractJSON []byte) ([]SMSv2DataSourceTemplate, error) {
	var contract smsv2ReleaseContract
	if err := json.Unmarshal(contractJSON, &contract); err != nil {
		return nil, fmt.Errorf("decode SMSv2 contract: %w", err)
	}
	if contract.Version != "7.0.0" || contract.ContractID != "f5xc-smsv2-api/v1" {
		return nil, fmt.Errorf("SMSv2 data sources require the v7.0 clean-break API v1 contract")
	}
	if err := validateAzureRouteServerContract(contract.Providers.Azure); err != nil {
		return nil, err
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
		return nil, fmt.Errorf("SMSv2 API v1 runtime endpoint set is incomplete")
	}
	for name, want := range wantRuntime {
		got := contract.Providers.AWS.Runtime[name]
		if got["method"] != want.method || got["path"] != want.path || got["operation_id"] == "" || got["response_schema"] == "" {
			return nil, fmt.Errorf("SMSv2 API v1 runtime endpoint %q is incomplete", name)
		}
	}
	wantUpgrade := map[string]struct {
		method string
		path   string
	}{
		"site_status":      {"GET", "/api/config/namespaces/{namespace}/sites/{site}"},
		"target_discovery": {"GET", "/api/maurice/upgradable_sw_versions"},
		"precheck":         {"GET", "/api/maurice/namespaces/{namespace}/sites/{site}/pre_upgrade_check"},
		"upgrade_status":   {"GET", "/api/maurice/namespaces/{namespace}/sites/{site}/upgrade_status"},
		"software_upgrade": {"POST", "/api/config/namespaces/{namespace}/sites/{site}/upgrade_sw"},
		"os_upgrade":       {"POST", "/api/config/namespaces/{namespace}/sites/{site}/upgrade_os"},
	}
	for name, want := range wantUpgrade {
		operation, ok := contract.Providers.AWS.SiteUpgrade[name].(map[string]any)
		if !ok || operation["method"] != want.method || operation["path"] != want.path || operation["operation_id"] == "" {
			return nil, fmt.Errorf("SMSv2 site upgrade operation %q is incomplete", name)
		}
		if want.method == "GET" && operation["response_schema"] == "" {
			return nil, fmt.Errorf("SMSv2 site upgrade operation %q is incomplete", name)
		}
		if want.method == "POST" {
			force, hasForce := operation["force"].(bool)
			if operation["request_schema"] == "" || !hasForce || force || operation["semantics"] != "asynchronous" {
				return nil, fmt.Errorf("SMSv2 site upgrade operation %q is incomplete", name)
			}
		}
	}
	polling, pollingOK := contract.Providers.AWS.SiteUpgrade["polling"].(map[string]any)
	transient, transientOK := polling["transient_failure_values"].([]any)
	redaction, redactionOK := contract.Providers.AWS.SiteUpgrade["redaction"].(map[string]any)
	prohibited, prohibitedOK := redaction["prohibited"].([]any)
	if !pollingOK || !transientOK || fmt.Sprint(transient) != "[UPGRADE_FAILED FAILED]" ||
		polling["failure_semantics"] != "transient_until_bounded_timeout" ||
		polling["completion"] != "supplied_targets_installed_and_site_online" ||
		polling["timeout_authority"] != "caller" || !redactionOK || !prohibitedOK ||
		redaction["exported"] != "sanitized_status_fields_only" ||
		fmt.Sprint(prohibited) != "[raw_api_messages node_identifiers urls]" {
		return nil, fmt.Errorf("SMSv2 site upgrade polling or redaction contract is incomplete")
	}
	wantF5XC := []string{"smsv2_configuration", "runtime_health", "bgp_peers", "bgp_routes", "simplified_routes", "site_upgrade_observation"}
	wantAWS := []string{"eni", "transit_gateway", "transit_gateway_connect", "gre_endpoints", "bgp_inside_cidrs", "autonomous_system_numbers"}
	if !equalStrings(contract.Providers.AWS.Authorities["f5xc"], wantF5XC) || !equalStrings(contract.Providers.AWS.Authorities["aws"], wantAWS) {
		return nil, fmt.Errorf("SMSv2 API v1 authority mapping is incomplete")
	}
	wantFacts := []string{"runtime", "gre", "bgp", "mtu", "route", "bgp_inside_cidr_block"}
	telemetry := contract.Providers.AWS.Telemetry
	if telemetry.SchemaID != "f5xc-smsv2-aws-tgw-telemetry/v2" ||
		!equalStringSets(telemetry.RequiredFacts, wantFacts) ||
		!equalStringSets(telemetry.ObservedFacts, wantFacts) ||
		len(telemetry.UnavailableFacts) != 0 {
		return nil, fmt.Errorf("SMSv2 API v1 telemetry declaration is incomplete")
	}
	available := map[string]string{"aws_ce_create": "available", "runtime_status": "available", "site_upgrade": "available", "tgw_connect": "available"}
	unavailable := map[string]string{"aws_ce_create": "unavailable", "runtime_status": "unavailable", "site_upgrade": "unavailable", "tgw_connect": "unavailable"}
	switch contract.Providers.AWS.Availability {
	case "evidence_backed":
		if !equalStringMap(contract.Providers.AWS.Capabilities, available) ||
			len(contract.Providers.AWS.UnavailableCapabilities) != 0 ||
			telemetry.Availability != "available" || !telemetry.Complete {
			return nil, fmt.Errorf("SMSv2 evidence-backed capabilities are incoherent")
		}
	case "schema_only":
		if !equalStringMap(contract.Providers.AWS.Capabilities, unavailable) ||
			!equalStringSets(contract.Providers.AWS.UnavailableCapabilities, []string{"aws_ce_create", "runtime_status", "site_upgrade", "tgw_connect"}) ||
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
		{Name: "site_upgrade_status", Kind: "upgrade"},
	}, nil
}

func validateAzureRouteServerContract(azure struct {
	Availability            string                    `json:"availability"`
	RouteServerEBGPMultihop smsv2CapabilityBoundary   `json:"route_server_ebgp_multihop"`
	Runtime                 map[string]map[string]any `json:"runtime"`
}) error {
	boundary := azure.RouteServerEBGPMultihop
	if azure.Availability != "evidence_backed" || boundary.Availability != "unavailable" ||
		boundary.Enforcement != "reject_before_mutation" || boundary.Reason != "no_schema_valid_ebgp_multihop_request_control" {
		return fmt.Errorf("Azure Route Server eBGP multihop boundary is incomplete")
	}
	if boundary.Source.Repository != "f5-sales-demo/api-specs-enriched" || len(boundary.Source.Commit) != 40 ||
		boundary.Source.AssetPath != "docs/specifications/api/network.json" ||
		!strings.HasPrefix(boundary.Source.AssetSHA256, "sha256:") || len(boundary.Source.AssetSHA256) != 71 ||
		!equalStrings(boundary.Source.SchemaPaths, []string{
			"components.schemas.bgpPeer",
			"components.schemas.bgpPeerExternal",
			"components.schemas.bgpBgpParameters",
		}) {
		return fmt.Errorf("Azure Route Server eBGP multihop provenance is incomplete")
	}
	wantRequirements := map[string]string{
		"request_schema_path":     "explicit",
		"request_field_path":      "explicit",
		"request_value_semantics": "explicit",
		"schema_validation":       "required",
		"runtime_acceptance":      "required",
	}
	if !equalStringMap(boundary.FutureMappingRequirements, wantRequirements) {
		return fmt.Errorf("Azure Route Server eBGP multihop promotion requirements are incomplete")
	}
	bgp := azure.Runtime["bgp_configuration"]
	responseMappings, mappingsOK := bgp["response_mappings"].(map[string]any)
	responseOnly, responseOnlyOK := bgp["response_only_fields"].([]any)
	if bgp["method"] != "GET" || bgp["path"] != "/api/config/namespaces/{namespace}/bgps/{name}" ||
		bgp["operation_id"] != "ves.io.schema.bgp.API.Get" || bgp["response_schema"] != "bgpGetResponse" ||
		bgp["authority"] != "f5xc" || bgp["semantics"] != "observational_read_only" ||
		bgp["request_eligibility"] != "rejected_without_authoritative_request_schema" ||
		!mappingsOK || responseMappings["target_service"] != "spec.peers[].target_service" ||
		responseMappings["family_inet_v6"] != "spec.peers[].external.family_inet_v6" ||
		!responseOnlyOK || fmt.Sprint(responseOnly) != "[spec.peers[].target_service spec.peers[].external.family_inet_v6]" {
		return fmt.Errorf("Azure BGP response-only observation contract is incomplete")
	}
	if _, writable := bgp["request_mappings"]; writable {
		return fmt.Errorf("Azure BGP response-only fields cannot be promoted to request inputs")
	}
	return nil
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
