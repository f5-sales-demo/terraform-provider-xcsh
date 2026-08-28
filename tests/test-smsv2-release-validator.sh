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
  receipts:[{sanitized:true,redaction:"fixture"}]
}' >"$work/smsv2-evidence-receipt.json"
jq -n '{
  contract_id:"f5xc-ce-automation/v2",
  api:{namespace:"system",operations:["create","read","replace","delete"]},
  providers:{aws:{availability:"evidence_backed",capabilities:{
    aws_ce_create:"available",runtime_status:"unavailable",tgw_connect:"unavailable"
  }}}
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
contract_sha="sha256:$(sha256sum "$work/smsv2-contract.json" | awk '{print $1}')"
evidence_sha="sha256:$(sha256sum "$work/smsv2-evidence-receipt.json" | awk '{print $1}')"
jq -n --arg tag "$tag" --arg commit "$commit" --arg contract "$contract_sha" --arg evidence "$evidence_sha" '{
  schema_version:1,
  contract_id:"f5xc-ce-automation/v2",
  contract_version:"1.0.0",
  release:{tag:$tag,commit:$commit},
  assets:{"smsv2-contract.json":$contract,"smsv2-evidence-receipt.json":$evidence}
}' >"$work/smsv2-contract-manifest.json"

python3 "$validator" "$work" "$tag" "$commit"
retired="$work/retired-v1"
mkdir "$retired"
cp "$work/smsv2-contract.json" "$work/smsv2-evidence-receipt.json" "$work/smsv2-contract-manifest.json" \
  "$work/concurrency_contracts.json" "$work/smsv2_parity_manifest.json" "$retired/"
for retired_asset in smsv2-contract.json smsv2-evidence-receipt.json smsv2-contract-manifest.json; do
  jq '.contract_id = "f5xc-ce-automation/v1"' "$retired/$retired_asset" >"$retired/updated.json"
  mv "$retired/updated.json" "$retired/$retired_asset"
done
retired_contract_sha="sha256:$(sha256sum "$retired/smsv2-contract.json" | awk '{print $1}')"
retired_evidence_sha="sha256:$(sha256sum "$retired/smsv2-evidence-receipt.json" | awk '{print $1}')"
jq --arg contract "$retired_contract_sha" --arg evidence "$retired_evidence_sha" '
  .assets["smsv2-contract.json"] = $contract |
  .assets["smsv2-evidence-receipt.json"] = $evidence
' "$retired/smsv2-contract-manifest.json" >"$retired/updated.json"
mv "$retired/updated.json" "$retired/smsv2-contract-manifest.json"
if python3 "$validator" "$retired" "$tag" "$commit" >/dev/null 2>&1; then
  echo "retired v1 contract identity was accepted" >&2
  exit 1
fi
naive="$work/naive"
mkdir "$naive"
cp "$work/smsv2-contract.json" "$work/smsv2-evidence-receipt.json" "$work/smsv2-contract-manifest.json" \
  "$work/concurrency_contracts.json" "$work/smsv2_parity_manifest.json" "$naive/"
jq '.observed_at = "2026-08-01T12:00:00"' "$naive/smsv2-evidence-receipt.json" >"$naive/evidence.json"
mv "$naive/evidence.json" "$naive/smsv2-evidence-receipt.json"
naive_evidence_sha="sha256:$(sha256sum "$naive/smsv2-evidence-receipt.json" | awk '{print $1}')"
jq --arg digest "$naive_evidence_sha" '.assets["smsv2-evidence-receipt.json"] = $digest' "$naive/smsv2-contract-manifest.json" >"$naive/manifest.json"
mv "$naive/manifest.json" "$naive/smsv2-contract-manifest.json"
if python3 "$validator" "$naive" "$tag" "$commit" >/dev/null 2>&1; then
  echo "timezone-naive evidence was accepted" >&2
  exit 1
fi
printf x >>"$work/smsv2-contract.json"
if python3 "$validator" "$work" "$tag" "$commit"; then
  echo "tampered contract was accepted" >&2
  exit 1
fi
echo "SMSv2 release validator tests passed"
