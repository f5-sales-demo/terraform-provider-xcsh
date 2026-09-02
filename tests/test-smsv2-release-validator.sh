#!/usr/bin/env bash
set -euo pipefail

root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
validator="$root/scripts/validate-smsv2-release.py"
work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT

tag=v9.9.9
commit=0123456789abcdef0123456789abcdef01234567
observed=$(date -u +%Y-%m-%dT%H:%M:%SZ)
jq -n --arg observed "$observed" '{
  contract_id:"f5xc-ce-automation/v2",
  observed_at:$observed,
  receipts:[{operations:["create","read","replace","delete"],result:"accepted",sanitized:true,redaction:"fixture"}]
}' >"$work/smsv2-evidence-receipt.json"
jq -n '{
  version:"5.0.0",
  contract_id:"f5xc-ce-automation/v2",
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
      schema_id:"f5xc-smsv2-aws-tgw-telemetry/v1",availability:"available",complete:true,
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
jq -n '{
  version:"9.9.9",eligible_count:1,covered_count:1,excluded_count:1,
  resources:[{api_identity:"ves.io.schema.probe.API",get:{schema:"probeGetResponse"},replace:{schema:"probeReplaceRequest"},token:"resource_version"}],
  exclusions:[{api_identity:"ves.io.schema.command.API",reason:"command endpoint"}]
}' >"$work/concurrency_contracts.json"
jq -n '{
  version:"9.9.9",resource:"securemesh_site_v2",path_count:1,
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
  local contract_sha evidence_sha
  contract_sha="sha256:$(sha256sum "$directory/smsv2-contract.json" | awk '{print $1}')"
  evidence_sha="sha256:$(sha256sum "$directory/smsv2-evidence-receipt.json" | awk '{print $1}')"
  jq -n --arg tag "$tag" --arg commit "$commit" --arg contract "$contract_sha" --arg evidence "$evidence_sha" '{
    schema_version:1,contract_id:"f5xc-ce-automation/v2",contract_version:"5.0.0",
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
  refresh_manifest "$directory"
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
jq '.observed_at = "2026-08-01T12:00:00"' "$naive/smsv2-evidence-receipt.json" >"$naive/evidence.json"
mv "$naive/evidence.json" "$naive/smsv2-evidence-receipt.json"
refresh_manifest "$naive"
if python3 "$validator" "$naive" "$tag" "$commit" >/dev/null 2>&1; then
  echo "timezone-naive evidence was accepted" >&2
  exit 1
fi
printf x >>"$work/smsv2-contract.json"
if python3 "$validator" "$work" "$tag" "$commit" >/dev/null 2>&1; then
  echo "tampered contract was accepted" >&2
  exit 1
fi
printf '%s\n' 'SMSv2 v2 release validator tests passed'
