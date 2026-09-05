#!/usr/bin/env bash
set -euo pipefail

root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
validator="$root/scripts/validate-smsv2-release.py"
work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT

tag=v6.0.0
commit=0123456789abcdef0123456789abcdef01234567
observed=$(date -u +%Y-%m-%dT%H:%M:%SZ)
jq -n --arg observed "$observed" '{
  contract_id:"f5xc-ce-automation/v3",
  recorded_at:$observed,
  receipts:[{operations:["create","read","replace","delete"],result:"accepted",sanitized:true,redaction:"fixture"}]
}' >"$work/smsv2-evidence-receipt.json"
jq -n '{
  version:"6.0.0",
  contract_id:"f5xc-ce-automation/v3",
  resource:"securemesh_site_v2",
  api:{namespace:"system",operations:["create","read","replace","delete"]},
  providers:{aws:{
    availability:"evidence_backed",
    node_list_path:"aws.not_managed.node_list[]",
    interface_list_path:"aws.not_managed.node_list[].interface_list[]",
    capabilities:{aws_ce_create:"available",runtime_status:"available",tgw_connect:"available"},
    interface_identity:{field:"ethernet_interface.mac",guest_device:"observational_only",known_macs:"non_empty_unique_per_node"},
    roles:[{name:"slo",network_option:"site_local_network"},{name:"sli",network_option:"site_local_inside_network"}],
    telemetry_intake:{
      schema_id:"f5xc-smsv2-aws-tgw-telemetry/v2",availability:"available",complete:true,
      required_facts:["runtime","gre","bgp","mtu","route","bgp_inside_cidr_block"],
      observed_facts:["runtime","gre","bgp","mtu","route","bgp_inside_cidr_block"],
      unavailable_facts:[]
    },
    runtime:{
      configuration:{method:"GET",path:"/api/config/namespaces/{namespace}/securemesh_site_v2s/{site}",operation_id:"ves.io.schema.views.securemesh_site_v2.API.Get",response_schema:"securemesh_site_v2GetResponse"},
      health:{method:"GET",path:"/api/operate/namespaces/system/sites/{site}/vpm/debug/global/health",operation_id:"ves.io.schema.operate.debug.CustomPublicAPI.HealthPublic",response_schema:"debugHealthResponse"},
      bgp_peers:{method:"GET",path:"/api/operate/namespaces/{namespace}/sites/{site}/ver/bgp_peers",operation_id:"ves.io.schema.operate.bgp.CustomPublicAPI.ShowBGPPeers",response_schema:"bgpBGPPeersResponse"},
      bgp_routes:{method:"GET",path:"/api/operate/namespaces/{namespace}/sites/{site}/ver/bgp_routes",operation_id:"ves.io.schema.operate.bgp.CustomPublicAPI.ShowBGPRoutes",response_schema:"bgpBGPRoutesResponse"},
      simplified_routes:{method:"POST",path:"/api/operate/namespaces/{namespace}/sites/{site}/ver/simplified_routes",operation_id:"ves.io.schema.operate.route.CustomPublicAPI.ShowSimplifiedRoutes",request_schema:"routeSimplifiedRouteRequest",response_schema:"routeSimplifiedRouteResponse"}
    },
    authorities:{
      f5xc:["smsv2_configuration","runtime_health","bgp_peers","bgp_routes","simplified_routes"],
      aws:["eni","transit_gateway","transit_gateway_connect","gre_endpoints","bgp_inside_cidrs"]
    },
    prohibited_legacy_apis:["aws_vpc_site","aws_tgw_site"],
    unavailable_capabilities:[]
  }}
}' >"$work/smsv2-contract.json"
jq '
  .providers.aws.interface_identity = {
    fields:["node","ethernet_interface.mac"],
    guest_device:"rejected",
    known_value_policy:"reject_null_incomplete_malformed_ambiguous_or_inconsistent",
    mac:{configuration_path:"spec.aws.not_managed.node_list[].interface_list[].ethernet_interface.mac",input_field:"mac",normalization:"ieee802_lowercase_colon",nullable:false},
    node:{configuration_path:"spec.aws.not_managed.node_list[].hostname",input_field:"node",normalization:"trim",nullable:false},
    uniqueness_scope:"node",
    unknown_value_policy:"defer"
  }
  | .providers.aws.authorities.aws += ["autonomous_system_numbers"]
  | .providers.aws.runtime.configuration += {
      authority:"f5xc",semantics:"configuration",correlation:["node","normalized_mac"],
      response_mappings:{nodes:"spec.aws.not_managed.node_list[]"},
      normalization:{mac:"ieee802_lowercase_colon",node:"trim",role:"slo_or_sli"},
      nullability:{all_identity_fields:"non_null",public_ip:"nullable"}
    }
  | .providers.aws.runtime.health += {
      authority:"f5xc",semantics:"observational_read_only",
      correlation:["node"],response_mappings:{node:"hostname"}
    }
  | .providers.aws.runtime.bgp_peers += {
      authority:"f5xc",semantics:"observational_read_only",
      correlation:["node","peer_address"],
      response_mappings:{state_changed_at:"ver[].peer[].up_down_timestamp"}
    }
  | .providers.aws.runtime.bgp_routes += {
      authority:"f5xc",semantics:"observational_read_only",
      correlation:["node"],response_mappings:{nodes:"ver[]"}
    }
  | .providers.aws.runtime.simplified_routes += {
      authority:"f5xc",semantics:"observational_read_only",
      correlation:["node","role"],request_mappings:{node_scope:"all_nodes",roles:["slo","sli"]},
      response_mappings:{nodes:"ver_routes[]"}
    }
' "$work/smsv2-contract.json" >"$work/contract-v3.json"
mv "$work/contract-v3.json" "$work/smsv2-contract.json"
jq -n '{
  version:"6.0.0",eligible_count:1,covered_count:1,excluded_count:1,
  resources:[{api_identity:"ves.io.schema.probe.API",get:{schema:"probeGetResponse"},replace:{schema:"probeReplaceRequest"},token:"resource_version"}],
  exclusions:[{api_identity:"ves.io.schema.command.API",reason:"command endpoint"}]
}' >"$work/concurrency_contracts.json"
jq -n '{
  version:"6.0.0",resource:"securemesh_site_v2",path_count:1,
  paths:[{path:"spec.segment_vrf[].segment_network",type:"object"}],choice_groups:{},
  deprecated_exclusions:["spec.log_receiver","spec.private_adn","spec.rseries"],
  current_platform_removals:[
    "spec.segment_vrf[].segment_config.nameserver_v6",
    "spec.segment_vrf[].segment_config.secondary_nameserver_v6"
  ]
}' >"$work/smsv2_parity_manifest.json"
jq -n '{
  components:{schemas:{viewssecuremesh_site_v2Node:{properties:{
    public_ip:{type:"string",nullable:true}
  }}}}
}' >"$work/openapi.json"

refresh_manifest() {
  local directory=$1
  local contract_sha evidence_sha contract_version
  contract_sha="sha256:$(sha256sum "$directory/smsv2-contract.json" | awk '{print $1}')"
  evidence_sha="sha256:$(sha256sum "$directory/smsv2-evidence-receipt.json" | awk '{print $1}')"
  contract_version=${2:-$(jq -r .version "$directory/smsv2-contract.json")}
  jq -n --arg tag "$tag" --arg commit "$commit" --arg contract "$contract_sha" --arg evidence "$evidence_sha" --arg contract_version "$contract_version" '{
    schema_version:1,contract_id:"f5xc-ce-automation/v3",contract_version:$contract_version,
    release:{tag:$tag,commit:$commit},
    assets:{"smsv2-contract.json":$contract,"smsv2-evidence-receipt.json":$evidence}
  }' >"$directory/smsv2-contract-manifest.json"
}

reject_contract_mutation() {
  local name=$1 filter=$2 directory="$work/$1"
  mkdir "$directory"
  cp "$work"/*.json "$directory/"
  jq "$filter" "$directory/smsv2-contract.json" >"$directory/updated.json"
  mv "$directory/updated.json" "$directory/smsv2-contract.json"
  if [ "$name" = contract-version-mismatch ]; then
    refresh_manifest "$directory" 6.0.0
  else
    refresh_manifest "$directory"
  fi
  if python3 "$validator" "$directory" "$tag" "$commit" >/dev/null 2>&1; then
    printf 'invalid contract was accepted: %s\n' "$name" >&2
    exit 1
  fi
}

refresh_manifest "$work"
python3 "$validator" "$work" "$tag" "$commit"

schema_only="$work/schema-only"
mkdir "$schema_only"
cp "$work"/*.json "$schema_only/"
jq '.providers.aws.availability = "schema_only"
  | .providers.aws.capabilities = {aws_ce_create:"unavailable",runtime_status:"unavailable",tgw_connect:"unavailable"}
  | .providers.aws.unavailable_capabilities = ["aws_ce_create","runtime_status","tgw_connect"]
  | .providers.aws.telemetry_intake.availability = "unavailable"
  | .providers.aws.telemetry_intake.complete = false' "$schema_only/smsv2-contract.json" >"$schema_only/contract.json"
mv "$schema_only/contract.json" "$schema_only/smsv2-contract.json"
jq '.receipts = [{
  operations:["replace"],result:"rejected",
  blocking_conditions:["mac_only_interface_rejected_by_live_api","public_ip_empty_string_null_round_trip"],
  sanitized:true,redaction:"fixture"
}]' "$schema_only/smsv2-evidence-receipt.json" >"$schema_only/evidence.json"
mv "$schema_only/evidence.json" "$schema_only/smsv2-evidence-receipt.json"
refresh_manifest "$schema_only"
python3 "$validator" "$schema_only" "$tag" "$commit"

bad_blocker="$work/bad-blocker"
mkdir "$bad_blocker"
cp "$schema_only"/*.json "$bad_blocker/"
jq '.receipts[0].blocking_conditions = ["unverified"]' "$bad_blocker/smsv2-evidence-receipt.json" >"$bad_blocker/evidence.json"
mv "$bad_blocker/evidence.json" "$bad_blocker/smsv2-evidence-receipt.json"
refresh_manifest "$bad_blocker"
if python3 "$validator" "$bad_blocker" "$tag" "$commit" >/dev/null 2>&1; then
  echo "schema-only contract with unverified blocker was accepted" >&2
  exit 1
fi

non_nullable="$work/non-nullable-public-ip"
mkdir "$non_nullable"
cp "$schema_only"/*.json "$non_nullable/"
jq 'del(.components.schemas.viewssecuremesh_site_v2Node.properties.public_ip.nullable)' \
  "$non_nullable/openapi.json" >"$non_nullable/updated.json"
mv "$non_nullable/updated.json" "$non_nullable/openapi.json"
if python3 "$validator" "$non_nullable" "$tag" "$commit" >/dev/null 2>&1; then
  echo "non-nullable SMSv2 node public_ip was accepted" >&2
  exit 1
fi

reject_contract_mutation retired-v1 '.contract_id = "f5xc-ce-automation/v1"'
reject_contract_mutation contract-version-mismatch '.version = "5.0.1"'
reject_contract_mutation unavailable-only '.providers.aws.capabilities.runtime_status = "unavailable"'
reject_contract_mutation incomplete-telemetry '.providers.aws.telemetry_intake.complete = false'
reject_contract_mutation legacy-interface-path '.providers.aws.runtime.configuration.path = "/api/config/namespaces/{namespace}/sites/{site}/interface"'
reject_contract_mutation general-routes-path '.providers.aws.runtime.simplified_routes.path = "/api/operate/namespaces/{namespace}/sites/{site}/ver/routes"'
reject_contract_mutation guest-device-identity '.providers.aws.interface_identity.field = "ethernet_interface.device"'
reject_contract_mutation authority-mismatch '.providers.aws.authorities.aws += ["runtime_health"]'

naive="$work/naive"
mkdir "$naive"
cp "$work"/*.json "$naive/"
jq '.recorded_at = "2026-08-01T12:00:00"' "$naive/smsv2-evidence-receipt.json" >"$naive/evidence.json"
mv "$naive/evidence.json" "$naive/smsv2-evidence-receipt.json"
refresh_manifest "$naive"
if python3 "$validator" "$naive" "$tag" "$commit" >/dev/null 2>&1; then
  echo "timezone-naive evidence was accepted" >&2
  exit 1
fi
site_upgrade="$work/site-upgrade"
mkdir "$site_upgrade"
cp "$work"/*.json "$site_upgrade/"
tag=v6.1.0
jq '
  .version = "6.1.0"
  | .providers.aws.capabilities.site_upgrade = "available"
  | .providers.aws.authorities.f5xc += ["site_upgrade_observation"]
  | .providers.aws.site_upgrade = {
      site_status:{method:"GET",path:"/api/config/namespaces/{namespace}/sites/{site}",operation_id:"ves.io.schema.site.API.Get",response_schema:"siteGetResponse",response_mappings:{software_installed_version:"status[].volterra_software_status.last_installed_version",software_available_version:"status[].volterra_software_status.available_version",software_deployment_phase:"status[].volterra_software_status.deployment_state.phase",software_deployment_result:"status[].volterra_software_status.deployment_state.result",os_installed_version:"status[].operating_system_status.deployment_state.version",os_available_version:"status[].operating_system_status.available_version",os_deployment_phase:"status[].operating_system_status.deployment_state.phase",os_deployment_result:"status[].operating_system_status.deployment_state.result",site_state:"spec.site_state"}},
      target_discovery:{method:"GET",path:"/api/maurice/upgradable_sw_versions",operation_id:"ves.io.schema.upgrade_status.UpgradeStatusCustomApi.GetUpgradableSWVersions",response_schema:"upgrade_statusGetUpgradableSWVersionsResponse",query_mappings:{installed_os_version:"current_os_version",installed_software_version:"current_sw_version"},response_mappings:{upgradable_software_versions:"sw_versions[]"}},
      precheck:{method:"GET",path:"/api/maurice/namespaces/{namespace}/sites/{site}/pre_upgrade_check",operation_id:"ves.io.schema.upgrade_status.UpgradeStatusCustomApi.PreUpgradeCheck",response_schema:"upgrade_statusPreUpgradeCheckResponse",query_mappings:{software_version:"sw_version"},response_mappings:{checks:"checklist[]",name:"checklist[].item",status:"checklist[].status"},passing_statuses:["CHECKLIST_PASSED","CHECKLIST_WARNING"],failure_statuses:["CHECKLIST_FAILED","CHECKLIST_UNKNOWN"]},
      upgrade_status:{method:"GET",path:"/api/maurice/namespaces/{namespace}/sites/{site}/upgrade_status",operation_id:"ves.io.schema.upgrade_status.UpgradeStatusCustomApi.GetUpgradeStatus",response_schema:"upgrade_statusGetUpgradeStatusResponse",response_mappings:{version:"upgrade_status.sw_upgrade_progress.version",status:"upgrade_status.sw_upgrade_progress.status",site_level_status:"upgrade_status.sw_upgrade_progress.site_level_upgrade.status",node_level_status:"upgrade_status.sw_upgrade_progress.node_level_upgrade.status",validation_status:"upgrade_status.sw_upgrade_progress.validation.status",os_setup_status:"upgrade_status.sw_upgrade_progress.os_setup.status"}},
      software_upgrade:{method:"POST",path:"/api/config/namespaces/{namespace}/sites/{site}/upgrade_sw",operation_id:"ves.io.schema.site.UpgradeAPI.UpgradeSW",request_schema:"siteUpgradeSWRequest",request_mappings:{site:"name",software_version:"version",force:"force"},force:false,semantics:"asynchronous"},
      os_upgrade:{method:"POST",path:"/api/config/namespaces/{namespace}/sites/{site}/upgrade_os",operation_id:"ves.io.schema.site.UpgradeAPI.UpgradeOS",request_schema:"siteUpgradeOSRequest",request_mappings:{site:"name",os_version:"version",force:"force"},force:false,semantics:"asynchronous"},
      eligibility:{site_state:"ONLINE",os_target:"equals_advertised_os_available_version",software_target:"listed_in_upgradable_software_versions",software_prechecks:"all_pass_or_warning"},
      polling:{transient_failure_values:["UPGRADE_FAILED","FAILED"],failure_semantics:"transient_until_bounded_timeout",completion:"supplied_targets_installed_and_site_online",timeout_authority:"caller"},
      redaction:{exported:"sanitized_status_fields_only",prohibited:["raw_api_messages","node_identifiers","urls"]},
      verified_path:{installed_software:"crt-20251002-0027",installed_os:"9.2026.10",target_software:"crt-20260201-0179",target_os:"9.2026.17",software_prechecks:"passed"}
    }
' "$site_upgrade/smsv2-contract.json" >"$site_upgrade/updated.json"
mv "$site_upgrade/updated.json" "$site_upgrade/smsv2-contract.json"
jq '.receipts += [{
  operations:["site_status","target_discovery","precheck","software_upgrade","os_upgrade"],
  result:"accepted",sanitized:true,redaction:"fixture",
  upgrade_path:{installed_software:"crt-20251002-0027",installed_os:"9.2026.10",target_software:"crt-20260201-0179",target_os:"9.2026.17",software_prechecks:"passed"},
  validated_facts:["installed_and_available_versions","deployment_phase_and_result","site_state","software_target_advertised","software_prechecks_passed","transient_failure_observed"]
}]' "$site_upgrade/smsv2-evidence-receipt.json" >"$site_upgrade/evidence.json"
mv "$site_upgrade/evidence.json" "$site_upgrade/smsv2-evidence-receipt.json"
jq '.version = "6.1.0"' "$site_upgrade/concurrency_contracts.json" >"$site_upgrade/concurrency.json"
mv "$site_upgrade/concurrency.json" "$site_upgrade/concurrency_contracts.json"
jq '.version = "6.1.0"' "$site_upgrade/smsv2_parity_manifest.json" >"$site_upgrade/parity.json"
mv "$site_upgrade/parity.json" "$site_upgrade/smsv2_parity_manifest.json"
refresh_manifest "$site_upgrade"
python3 "$validator" "$site_upgrade" "$tag" "$commit"

reject_upgrade_mutation() {
  local name=$1 filter=$2 directory="$work/upgrade-$1"
  mkdir "$directory"
  cp "$site_upgrade"/*.json "$directory/"
  jq "$filter" "$directory/smsv2-contract.json" >"$directory/updated.json"
  mv "$directory/updated.json" "$directory/smsv2-contract.json"
  refresh_manifest "$directory"
  if python3 "$validator" "$directory" "$tag" "$commit" >/dev/null 2>&1; then
    printf 'invalid site upgrade contract was accepted: %s\n' "$name" >&2
    exit 1
  fi
}

reject_upgrade_mutation status-mapping '.providers.aws.site_upgrade.site_status.response_mappings.site_state = "status[].site_state"'
reject_upgrade_mutation mutable-force '.providers.aws.site_upgrade.software_upgrade.force = true'
reject_upgrade_mutation terminal-failure '.providers.aws.site_upgrade.polling.failure_semantics = "terminal"'
reject_upgrade_mutation raw-export '.providers.aws.site_upgrade.redaction.exported = "raw_response"'

tag=v6.0.0
printf x >>"$work/smsv2-contract.json"
if python3 "$validator" "$work" "$tag" "$commit" >/dev/null 2>&1; then
  echo "tampered contract was accepted" >&2
  exit 1
fi

printf '%s\n' 'SMSv2 v3 release validator tests passed'
