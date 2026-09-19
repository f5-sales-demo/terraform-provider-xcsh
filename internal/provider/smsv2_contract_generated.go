// Code generated from api-specs-enriched v7.0.6 smsv2-contract.json. DO NOT EDIT.

package provider

const (
	smsv2ContractID        = "f5xc-smsv2-api/v1"
	smsv2ContractVersion   = "7.0.0"
	smsv2APIReleaseTag     = "v7.0.6"
	smsv2SourceCommit      = "9bae0474d119572574156b9ba1e538e6b0431cf0"
	smsv2TelemetrySchemaID = "f5xc-smsv2-aws-tgw-telemetry/v2"
)

var smsv2ContractCapabilities = map[string]string{"aws_ce_create": "available", "aws_node_configuration": "available", "runtime_status": "available", "site_upgrade": "available", "tgw_connect": "available"}
var smsv2AWSNodeConfigurationJSON = "{\"availability\":\"evidence_backed\",\"enforcement\":\"required\",\"field_paths\":[\"resource_version\",\"spec.aws.not_managed.node_list[]\",\"spec.aws.not_managed.node_list[].hostname\",\"spec.aws.not_managed.node_list[].interface_list[]\",\"spec.aws.not_managed.node_list[].interface_list[].ethernet_interface.device\",\"spec.aws.not_managed.node_list[].interface_list[].ethernet_interface.mac\"],\"invariants\":{\"device_source\":\"observed_registration_only\",\"ha\":\"disabled\",\"interface_count\":2,\"interface_role_cardinality\":\"exactly_one_each\",\"interface_roles\":[\"slo\",\"sli\"],\"mac_normalization\":\"ieee802_lowercase_colon\",\"node_count\":1},\"mapping\":{\"cardinality\":\"one_to_one\",\"device_policy\":\"observed_only\",\"device_value_path\":\"interfaces[].device\",\"join_key\":\"normalized_mac\",\"mac_value_path\":\"interfaces[].mac\",\"registration_source\":\"site_registration_hardware_inventory\",\"terraform_mac_source\":\"aws_network_interface.mac_address\"},\"operation\":{\"method\":\"PUT\",\"operation_id\":\"ves.io.schema.views.securemesh_site_v2.API.Replace\",\"path\":\"/api/config/namespaces/{metadata.namespace}/securemesh_site_v2s/{metadata.name}\",\"request_schema\":\"securemesh_site_v2ReplaceRequest\"},\"provenance\":{\"evidence_receipt_sha256\":\"a5e423f8223bce56b83b74d0c546072984240761d392f15b664920e9fbfeec29\",\"issue\":\"f5-sales-demo/api-specs-enriched#1776\",\"probe_date\":\"2026-09-17\",\"source_commit\":\"3be0310daf3658bc90383048cc18ab6447d18e35\",\"source_spec_sha256\":\"c6145b7096d7b08416732c8ebbd80e83be1902f73096ba08309e360d9a746eec\"},\"strategy\":\"discovery_rebuild\",\"unsupported_reasons\":{\"ambiguous_mapping\":\"aws_node_configuration_mapping_ambiguous\",\"direct_rebuild_mode_transition\":\"aws_node_configuration_discovery_rebuild_requires_distinct_site\",\"guessed_device\":\"aws_node_configuration_device_must_be_observed\",\"incomplete_request_semantics\":\"aws_node_configuration_request_semantics_incomplete\",\"malformed_mac\":\"aws_node_configuration_mac_malformed\",\"missing_mapping\":\"aws_node_configuration_mapping_missing\",\"multi_node_or_ha_input\":\"aws_node_configuration_requires_single_non_ha_node\"}}"
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
