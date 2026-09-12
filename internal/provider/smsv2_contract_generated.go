// Code generated from api-specs-enriched v7.0.1 smsv2-contract.json. DO NOT EDIT.

package provider

const (
	smsv2ContractID        = "f5xc-smsv2-api/v1"
	smsv2ContractVersion   = "7.0.0"
	smsv2APIReleaseTag     = "v7.0.1"
	smsv2SourceCommit      = "2513fe498149c98fb737ff2ab207704b8a86fec6"
	smsv2TelemetrySchemaID = "f5xc-smsv2-aws-tgw-telemetry/v2"
)

var smsv2ContractCapabilities = map[string]string{"aws_ce_create": "available", "runtime_status": "available", "site_upgrade": "available", "tgw_connect": "available"}
var smsv2ContractF5XCAuthorities = []string{"smsv2_configuration", "runtime_health", "bgp_peers", "bgp_routes", "simplified_routes", "site_upgrade_observation"}
var smsv2ContractAWSAuthorities = []string{"eni", "transit_gateway", "transit_gateway_connect", "gre_endpoints", "bgp_inside_cidrs", "autonomous_system_numbers"}
var smsv2AzureRouteServerEBGPMultihop = smsv2CapabilityBoundaryContract{
	Availability: "unavailable",
	Enforcement:  "reject_before_mutation",
	Reason:       "no_schema_valid_ebgp_multihop_request_control",
	Source: smsv2CapabilitySourceContract{
		Repository:  "f5-sales-demo/api-specs-enriched",
		Commit:      "322c202ed49c8cfcd5015a524f3195bbd2a8f2bc",
		AssetPath:   "docs/specifications/api/network.json",
		AssetSHA256: "sha256:a7398d85475409c93750a04ccdec7c0b1a7ae12bc362ebfbc76866275904762a",
		SchemaPaths: []string{"components.schemas.bgpPeer", "components.schemas.bgpPeerExternal", "components.schemas.bgpBgpParameters"},
	},
}
