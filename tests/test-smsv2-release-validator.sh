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
}' >"$work/smsv2-evidence-receipt.json"
jq -n '{
  contract_id:"f5xc-ce-automation/v1",
  api:{namespace:"system",operations:["create","read","replace","delete"]},
  providers:{aws:{availability:"evidence_backed",capabilities:{
    aws_ce_create:"available",runtime_status:"unavailable",tgw_connect:"unavailable"
  }}}
}' >"$work/smsv2-contract.json"
contract_sha="sha256:$(sha256sum "$work/smsv2-contract.json" | awk '{print $1}')"
evidence_sha="sha256:$(sha256sum "$work/smsv2-evidence-receipt.json" | awk '{print $1}')"
jq -n --arg tag "$tag" --arg commit "$commit" --arg contract "$contract_sha" --arg evidence "$evidence_sha" '{
  schema_version:1,
  contract_id:"f5xc-ce-automation/v1",
  contract_version:"1.0.0",
  release:{tag:$tag,commit:$commit},
  assets:{"smsv2-contract.json":$contract,"smsv2-evidence-receipt.json":$evidence}
}' >"$work/smsv2-contract-manifest.json"

python3 "$validator" "$work" "$tag" "$commit"
naive="$work/naive"
mkdir "$naive"
cp "$work/smsv2-contract.json" "$work/smsv2-evidence-receipt.json" "$work/smsv2-contract-manifest.json" "$naive/"
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
