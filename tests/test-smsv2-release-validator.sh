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
  contract_id:"f5xc-ce-automation/v1",
  observed_at:$observed,
  receipts:[{sanitized:true,redaction:"fixture"}]
}' > "$work/smsv2-evidence-receipt.json"
jq -n '{
  contract_id:"f5xc-ce-automation/v1",
  api:{namespace:"system",operations:["create","read","replace","delete"]},
  providers:{aws:{availability:"evidence_backed",capabilities:{
    aws_ce_create:"available",runtime_status:"unavailable",tgw_connect:"unavailable"
  }}}
}' > "$work/smsv2-contract.json"
contract_sha="sha256:$(sha256sum "$work/smsv2-contract.json" | awk '{print $1}')"
evidence_sha="sha256:$(sha256sum "$work/smsv2-evidence-receipt.json" | awk '{print $1}')"
jq -n --arg tag "$tag" --arg commit "$commit" --arg contract "$contract_sha" --arg evidence "$evidence_sha" '{
  schema_version:1,
  contract_id:"f5xc-ce-automation/v1",
  contract_version:"1.0.0",
  release:{tag:$tag,commit:$commit},
  assets:{"smsv2-contract.json":$contract,"smsv2-evidence-receipt.json":$evidence}
}' > "$work/smsv2-contract-manifest.json"

python3 "$validator" "$work" "$tag" "$commit"
printf x >> "$work/smsv2-contract.json"
if python3 "$validator" "$work" "$tag" "$commit"; then
  echo "tampered contract was accepted" >&2
  exit 1
fi
echo "SMSv2 release validator tests passed"
