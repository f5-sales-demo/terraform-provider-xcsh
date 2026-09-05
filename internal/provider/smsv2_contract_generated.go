// Code generated from api-specs-enriched v6.1.1 smsv2-contract.json. DO NOT EDIT.

package provider

const (
	smsv2ContractID        = "f5xc-ce-automation/v3"
	smsv2ContractVersion   = "6.1.0"
	smsv2APIReleaseTag     = "v6.1.1"
	smsv2SourceCommit      = "2b27355ac9bf4683d3a321f7d6388676f756c2f5"
	smsv2TelemetrySchemaID = "f5xc-smsv2-aws-tgw-telemetry/v2"
)

var smsv2ContractCapabilities = map[string]string{"aws_ce_create": "available", "runtime_status": "available", "site_upgrade": "available", "tgw_connect": "available"}
var smsv2ContractF5XCAuthorities = []string{"smsv2_configuration", "runtime_health", "bgp_peers", "bgp_routes", "simplified_routes", "site_upgrade_observation"}
var smsv2ContractAWSAuthorities = []string{"eni", "transit_gateway", "transit_gateway_connect", "gre_endpoints", "bgp_inside_cidrs", "autonomous_system_numbers"}
