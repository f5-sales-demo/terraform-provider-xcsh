// Code generated from api-specs-enriched v5.0.1 smsv2-contract.json. DO NOT EDIT.

package provider

const (
	smsv2ContractID        = "f5xc-ce-automation/v2"
	smsv2ContractVersion   = "5.0.0"
	smsv2APIReleaseTag     = "v5.0.1"
	smsv2SourceCommit      = "03297e8b403ac7df1cb0b897387edd74f836114a"
	smsv2TelemetrySchemaID = "f5xc-smsv2-aws-tgw-telemetry/v1"
)

var smsv2ContractCapabilities = map[string]string{"aws_ce_create": "unavailable", "runtime_status": "unavailable", "tgw_connect": "unavailable"}
var smsv2ContractF5XCAuthorities = []string{"smsv2_configuration", "runtime_health", "bgp_peers", "bgp_routes", "simplified_routes"}
var smsv2ContractAWSAuthorities = []string{"eni", "transit_gateway", "transit_gateway_connect", "gre_endpoints", "bgp_inside_cidrs"}
