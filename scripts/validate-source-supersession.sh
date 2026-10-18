#!/usr/bin/env bash
set -euo pipefail
base=${1:?base ref required}
expected=$(printf '%s\n' \
  'M tools/spec-pending-delivery.json' \
  'D tools/spec-regeneration-receipt.json' \
  'M tools/spec-release.json' \
  'M tools/spec-version.txt')
actual=$(git diff --name-status "$base" HEAD | sort -k2,2 | tr '\t' ' ')
[ "$actual" = "$expected" ] || exit 1
staging=$(mktemp -d)
trap 'rm -rf "$staging"' EXIT
git show "$base:tools/spec-pending-delivery.json" >"$staging/pending.json"
git show "$base:tools/spec-regeneration-receipt.json" >"$staging/generation.json"
jq -en --slurpfile old "$staging/pending.json" \
  --slurpfile generation "$staging/generation.json" \
  --slurpfile new tools/spec-pending-delivery.json '
  def version: split(".") | map(tonumber);
  $old[0].delivery_id == $generation[0].delivery_id and
  $old[0].release_tag == $generation[0].release_tag and
  $old[0].target_commit == $generation[0].target_commit and
  $old[0].version == $generation[0].version and
  $old[0].delivery_id != $new[0].delivery_id and
  ($new[0].version | version) > ($old[0].version | version)
' >/dev/null
scripts/validate-provider-delivery-state.sh --base-ref "$base"
