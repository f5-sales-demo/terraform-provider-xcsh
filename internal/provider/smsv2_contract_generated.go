// Code generated from api-specs-enriched v6.0.0 smsv2-contract.json. DO NOT EDIT.

package provider

const (
	smsv2ContractID        = "f5xc-ce-automation/v3"
	smsv2ContractVersion   = "6.0.0"
	smsv2APIReleaseTag     = "v6.0.0"
	smsv2SourceCommit      = "8a48ca67ad9fc23174d086c0d63a2783e531044b"
	smsv2TelemetrySchemaID = "f5xc-smsv2-aws-tgw-telemetry/v2"
)

var smsv2ContractCapabilities = map[string]string{"aws_ce_create": "unavailable", "runtime_status": "unavailable", "tgw_connect": "unavailable"}
var smsv2ContractF5XCAuthorities = []string{"smsv2_configuration", "runtime_health", "bgp_peers", "bgp_routes", "simplified_routes"}
var smsv2ContractAWSAuthorities = []string{"eni", "transit_gateway", "transit_gateway_connect", "gre_endpoints", "bgp_inside_cidrs", "autonomous_system_numbers"}
