#!/usr/bin/env bash
set -euo pipefail
id=${1:?delivery ID required}
base=${2:?base ref required}
: "${TARGET_REPOSITORY:?target repository required}"
[[ "$id" =~ ^[0-9a-f]{64}$ ]] || exit 1
expected=$(printf '%s\n' 'M tools/provider-publication-receipts.json' 'M tools/spec-deliveries.json')
actual=$(git diff --name-status "$base" HEAD | sort -k2,2 | tr '\t' ' ')
[ "$actual" = "$expected" ] || exit 1
staging=$(mktemp -d)
trap 'rm -rf "$staging"' EXIT
git show "$base:tools/spec-deliveries.json" >"$staging/old-common.json"
git show "$base:tools/provider-publication-receipts.json" >"$staging/old-publications.json"
jq -en --arg id "$id" --slurpfile old "$staging/old-common.json" \
  --slurpfile oldp "$staging/old-publications.json" \
  --slurpfile new tools/spec-deliveries.json \
  --slurpfile newp tools/provider-publication-receipts.json '
  (($new[0].deliveries | keys) - ($old[0].deliveries | keys)) == [$id] and
  (($newp[0].receipts | keys) - ($oldp[0].receipts | keys)) == [$id] and
  $newp[0].receipts[$id].delivery == $new[0].deliveries[$id]
' >/dev/null
scripts/validate-provider-delivery-state.sh --base-ref "$base"
jq -e --arg id "$id" '.receipts[$id].publication' tools/provider-publication-receipts.json >"$staging/publication.json"
tag=$(jq -er '.tag' "$staging/publication.json")
commit=$(jq -er '.commit' "$staging/publication.json")
[[ "$tag" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]] || exit 1
[[ "$commit" =~ ^[0-9a-f]{40}$ ]] || exit 1
gh api "repos/${TARGET_REPOSITORY}/releases/tags/${tag}" >"$staging/release.json"
jq -en --slurpfile release "$staging/release.json" --slurpfile publication "$staging/publication.json" '
  $release[0].immutable == true and $release[0].draft == false and $release[0].prerelease == false and
  $release[0].tag_name == $publication[0].tag and
  ($release[0].assets | length) == ($publication[0].assets | length) and
  ($release[0].assets | map({key:.name,value:.digest}) | from_entries) == $publication[0].assets
' >/dev/null
jq -r '.body' "$staging/release.json" | sed -n 's/^<!-- provider-publication-receipt:\(.*\) -->$/\1/p' >"$staging/body-receipt.json"
jq -es --slurpfile expected "$staging/publication.json" 'length == 1 and .[0] == $expected[0]' "$staging/body-receipt.json" >/dev/null
git fetch --no-tags origin "refs/tags/${tag}"
[ "$(git rev-list -n 1 FETCH_HEAD)" = "$commit" ]
git show "$commit:tools/spec-release.json" >"$staging/released-pin.json"
[ "$(sha256sum "$staging/released-pin.json" | awk '{print $1}')" = "$(jq -er '.spec_release_sha256' "$staging/publication.json")" ]
