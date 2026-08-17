#!/usr/bin/env bash

set -euo pipefail

: "${PROVIDER_TAG:?PROVIDER_TAG is required}"
: "${RELEASED_COMMIT:?RELEASED_COMMIT is required}"
: "${TARGET_REPOSITORY:?TARGET_REPOSITORY is required}"
: "${GITHUB_OUTPUT:?GITHUB_OUTPUT is required}"

pending=tools/spec-pending-delivery.json
ledger=tools/spec-deliveries.json
detailed_ledger=tools/provider-publication-receipts.json
regeneration_receipt=tools/spec-regeneration-receipt.json
pin=tools/spec-release.json

fail() {
  echo "::error::$*" >&2
  exit 1
}

if [ ! -e "$pending" ]; then
  echo "changed=false" >>"$GITHUB_OUTPUT"
  echo "No pending spec delivery to receipt"
  exit 0
fi

[[ "$PROVIDER_TAG" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]] ||
  fail "Provider release tag is malformed"
[[ "$RELEASED_COMMIT" =~ ^[0-9a-f]{40}$ ]] ||
  fail "Released commit is malformed"
[ "$(git rev-parse HEAD)" = "$RELEASED_COMMIT" ] ||
  fail "Receipt preparation is not running on the released commit"
[ "$(git rev-list -n 1 "$PROVIDER_TAG")" = "$RELEASED_COMMIT" ] ||
  fail "Provider tag does not identify the immutable released commit"

jq -e '
  type == "object" and
  (keys | sort) == ["delivery_id", "release_tag", "target_commit", "version"] and
  (.delivery_id | test("^[0-9a-f]{64}$")) and
  (.release_tag | test("^v[0-9]+\\.[0-9]+\\.[0-9]+$")) and
  (.target_commit | test("^[0-9a-f]{40}$")) and
  .release_tag == ("v" + .version)
' "$pending" >/dev/null || fail "Pending delivery is malformed"
canonical_id() {
  local version=$1 tag=$2 commit=$3
  local canonical
  canonical=$(jq -cnS \
    --arg commit "$commit" \
    --arg event_type enriched-specs-updated \
    --arg source f5-sales-demo/api-specs-enriched \
    --arg tag "$tag" \
    --arg target "$TARGET_REPOSITORY" \
    --arg version "$version" \
    '{commit:$commit,event_type:$event_type,source:$source,tag:$tag,target:$target,version:$version}')
  printf '%s' "$canonical" | shasum -a 256 | awk '{print $1}'
}

validate_ledgers() {
  jq -e '
    type == "object" and
    (keys | sort) == ["deliveries", "version"] and
    .version == 1 and
    (.deliveries | type == "object")
  ' "$ledger" >/dev/null || fail "Delivery ledger is malformed"
  jq -e '
    type == "object" and
    (keys | sort) == ["receipts", "version"] and
    .version == 1 and
    (.receipts | type == "object")
  ' "$detailed_ledger" >/dev/null || fail "Provider publication receipt ledger is malformed"
  jq -en --slurpfile common "$ledger" --slurpfile detailed "$detailed_ledger" '
    ($common[0].deliveries | keys | sort) == ($detailed[0].receipts | keys | sort)
  ' >/dev/null || fail "Common and detailed receipt ledgers have different delivery keys"

  while IFS= read -r entry; do
    key=$(jq -r '.key' <<<"$entry")
    jq -e '
      (.value | type == "object") and
      (.value | keys | sort) == ["release_tag", "target_commit", "version"] and
      (.value.release_tag == ("v" + .value.version)) and
      (.value.target_commit | test("^[0-9a-f]{40}$"))
    ' <<<"$entry" >/dev/null || fail "Delivery ledger contains a malformed entry"
    expected=$(canonical_id \
      "$(jq -r '.value.version' <<<"$entry")" \
      "$(jq -r '.value.release_tag' <<<"$entry")" \
      "$(jq -r '.value.target_commit' <<<"$entry")")
    [ "$key" = "$expected" ] || fail "Delivery ledger contains a noncanonical key"
  done < <(jq -c '.deliveries | to_entries[]' "$ledger")

  while IFS= read -r entry; do
    key=$(jq -r '.key' <<<"$entry")
    common_entry=$(jq -c --arg id "$key" '.deliveries[$id]' "$ledger")
    jq -e --argjson common "$common_entry" '
      (.value | type == "object") and
      (.value | keys | sort) == ["delivery", "publication"] and
      .value.delivery == $common and
      (.value.publication | type == "object") and
      (.value.publication | keys | sort) == ["assets", "commit", "spec_release_sha256", "tag", "version"] and
      (.value.publication.commit | test("^[0-9a-f]{40}$")) and
      (.value.publication.spec_release_sha256 | test("^[0-9a-f]{64}$")) and
      (.value.publication.tag == ("v" + .value.publication.version)) and
      (.value.publication.assets | keys | sort) == [
        ("mcp-data-" + .value.publication.version + ".tar.gz"),
        ("terraform-provider-xcsh_" + .value.publication.version + "_SHA256SUMS"),
        ("terraform-provider-xcsh_" + .value.publication.version + "_SHA256SUMS.sig"),
        ("terraform-provider-xcsh_" + .value.publication.version + "_darwin_amd64.zip"),
        ("terraform-provider-xcsh_" + .value.publication.version + "_darwin_arm64.zip"),
        ("terraform-provider-xcsh_" + .value.publication.version + "_freebsd_386.zip"),
        ("terraform-provider-xcsh_" + .value.publication.version + "_freebsd_amd64.zip"),
        ("terraform-provider-xcsh_" + .value.publication.version + "_linux_386.zip"),
        ("terraform-provider-xcsh_" + .value.publication.version + "_linux_amd64.zip"),
        ("terraform-provider-xcsh_" + .value.publication.version + "_linux_arm.zip"),
        ("terraform-provider-xcsh_" + .value.publication.version + "_linux_arm64.zip"),
        ("terraform-provider-xcsh_" + .value.publication.version + "_manifest.json"),
        ("terraform-provider-xcsh_" + .value.publication.version + "_windows_386.zip"),
        ("terraform-provider-xcsh_" + .value.publication.version + "_windows_amd64.zip")
      ] and
      ([.value.publication.assets[] | test("^sha256:[0-9a-f]{64}$")] | all)
    ' <<<"$entry" >/dev/null || fail "Provider publication ledger contains malformed evidence"
  done < <(jq -c '.receipts | to_entries[]' "$detailed_ledger")
}

validate_ledgers

delivery_id=$(jq -r '.delivery_id' "$pending")
expected_id=$(canonical_id \
  "$(jq -r '.version' "$pending")" \
  "$(jq -r '.release_tag' "$pending")" \
  "$(jq -r '.target_commit' "$pending")")
[ "$delivery_id" = "$expected_id" ] || fail "Pending delivery ID is not canonical"

receipt_status=$(git diff --name-status "${RELEASED_COMMIT}^" "$RELEASED_COMMIT" -- "$regeneration_receipt" |
  awk 'NR == 1 { print $1 }')
case "$receipt_status" in
A)
  ;;
M)
  previous_receipt=$(git show "${RELEASED_COMMIT}^:$regeneration_receipt") ||
    fail "Pre-release regeneration receipt could not be read"
  jq -e --slurpfile current "$regeneration_receipt" '
    type == "object" and
    (keys | sort) == [
      "delivery_id", "release_tag", "source_commit", "spec_release_sha256",
      "target_commit", "version"
    ] and
    (.source_commit | test("^[0-9a-f]{40}$")) and
    .source_commit != $current[0].source_commit and
    del(.source_commit) == ($current[0] | del(.source_commit))
  ' <<<"$previous_receipt" >/dev/null ||
    fail "Regeneration receipt may change only its source commit during recovery"
  ;;
*)
  fail "Released commit did not introduce or exactly rebind the regeneration receipt"
  ;;
esac
[ "$(tr -d '[:space:]' <tools/spec-version.txt)" = "$(jq -r '.release_tag' "$pending")" ] ||
  fail "Published spec marker does not match pending delivery"
jq -e --slurpfile pending "$pending" '
  type == "object" and
  (keys | sort) == ["assets", "release_tag", "target_commit", "version"] and
  .release_tag == $pending[0].release_tag and
  .target_commit == $pending[0].target_commit and
  .version == $pending[0].version and
  (.assets | type == "object") and
  (.assets | keys | sort) == [
    "api-catalog.json",
    ("f5xc-api-specs-" + $pending[0].release_tag + ".zip"),
    "index.json",
    "minimal-export-defaults.json",
    "openapi.json"
  ] and
  ([.assets[] | test("^sha256:[0-9a-f]{64}$")] | all)
' "$pin" >/dev/null || fail "Published release pin does not match pending delivery"
pin_sha=$(shasum -a 256 "$pin" | awk '{print $1}')
jq -e --arg pin_sha "$pin_sha" --slurpfile pending "$pending" '
  type == "object" and
  (keys | sort) == [
    "delivery_id", "release_tag", "source_commit", "spec_release_sha256",
    "target_commit", "version"
  ] and
  .delivery_id == $pending[0].delivery_id and
  .release_tag == $pending[0].release_tag and
  .target_commit == $pending[0].target_commit and
  .version == $pending[0].version and
  .spec_release_sha256 == $pin_sha and
  (.source_commit | test("^[0-9a-f]{40}$"))
' "$regeneration_receipt" >/dev/null ||
  fail "Regeneration receipt does not bind the pending delivery"
source_commit=$(jq -r '.source_commit' "$regeneration_receipt")
[ "$source_commit" = "$(git rev-parse "${RELEASED_COMMIT}^")" ] ||
  fail "Regeneration receipt source is not the exact release parent"

release_json=$(mktemp)
provider_receipts=$(mktemp)
provider_receipt=$(mktemp)
provider_body=$(mktemp)
expected_pending=$(mktemp)
expected_regeneration=$(mktemp)
trap 'rm -f "$release_json" "$provider_receipts" "$provider_receipt" "$provider_body" "$expected_pending" "$expected_regeneration"' EXIT
gh api "repos/${TARGET_REPOSITORY}/releases/tags/${PROVIDER_TAG}" >"$release_json"
jq -e --arg tag "$PROVIDER_TAG" '
  .tag_name == $tag and .draft == false and .prerelease == false and .immutable == true
' "$release_json" >/dev/null || fail "Provider release is absent or incomplete"
jq -r '.body // ""' "$release_json" >"$provider_body"
marker_count=$(grep -Ec '^<!-- provider-publication-receipt:.* -->$' "$provider_body" || true)
[ "$marker_count" -eq 1 ] ||
  fail "Provider release must contain exactly one publication receipt"
sed -n 's/^<!-- provider-publication-receipt:\(.*\) -->$/\1/p' \
  "$provider_body" >"$provider_receipts"
jq -S . "$provider_receipts" >"$provider_receipt" ||
  fail "Provider publication receipt is malformed"
provider_version=${PROVIDER_TAG#v}
jq -e \
  --arg commit "$RELEASED_COMMIT" \
  --arg pin_sha "$pin_sha" \
  --arg tag "$PROVIDER_TAG" \
  --arg version "$provider_version" '
  type == "object" and
  (keys | sort) == ["assets", "commit", "spec_release_sha256", "tag", "version"] and
  .commit == $commit and .tag == $tag and .version == $version and
  .spec_release_sha256 == $pin_sha and
  (.assets | type == "object") and
  (.assets | length == 14) and
  ([.assets[] | test("^sha256:[0-9a-f]{64}$")] | all)
' "$provider_receipt" >/dev/null || fail "Provider publication receipt has the wrong identity"
actual_assets=$(jq -r '.assets[].name' "$release_json" | LC_ALL=C sort)
receipted_assets=$(jq -r '.assets | keys[]' "$provider_receipt" | LC_ALL=C sort)
[ "$actual_assets" = "$receipted_assets" ] ||
  fail "Provider receipt and release asset sets differ"
while IFS= read -r name; do
  api_digest=$(jq -er --arg name "$name" '
    [.assets[] | select(.name == $name)] |
    if length == 1 then .[0].digest else empty end
  ' "$release_json")
  receipt_digest=$(jq -r --arg name "$name" '.assets[$name]' "$provider_receipt")
  [ "$api_digest" = "$receipt_digest" ] ||
    fail "Provider receipt differs from GitHub's measured digest for $name"
done <<<"$actual_assets"

jq -cS . "$pending" >"$expected_pending"
jq -cS . "$regeneration_receipt" >"$expected_regeneration"
git fetch origin main
git checkout --detach origin/main
if [ ! -e "$pending" ]; then
  [ ! -e "$regeneration_receipt" ] ||
    fail "Latest main has a regeneration receipt without pending delivery"
  validate_ledgers
  jq -e --arg id "$delivery_id" --slurpfile expected "$expected_pending" '
    .deliveries[$id] == ($expected[0] | del(.delivery_id))
  ' "$ledger" >/dev/null || fail "Latest main removed or changed the durable delivery receipt"
  jq -e \
    --arg id "$delivery_id" \
    --slurpfile expected "$expected_pending" \
    --slurpfile receipt "$provider_receipt" '
    .receipts[$id] == {
      delivery:($expected[0] | del(.delivery_id)),
      publication:$receipt[0]
    }
  ' "$detailed_ledger" >/dev/null ||
    fail "Latest main removed or changed the provider publication evidence"
  echo "changed=false" >>"$GITHUB_OUTPUT"
  echo "Delivery is already durably receipted on main"
  exit 0
fi
[ "$(jq -cS . "$pending")" = "$(cat "$expected_pending")" ] ||
  fail "Latest main no longer carries the released pending identity"
[ "$(jq -cS . "$regeneration_receipt")" = "$(cat "$expected_regeneration")" ] ||
  fail "Latest main no longer carries the released regeneration receipt"
[ "$(shasum -a 256 "$pin" | awk '{print $1}')" = "$pin_sha" ] ||
  fail "Latest main no longer carries the released spec pin"
validate_ledgers

if jq -e --arg id "$delivery_id" '.deliveries[$id] != null' "$ledger" >/dev/null; then
  jq -e --arg id "$delivery_id" --slurpfile pending "$pending" '
    .deliveries[$id] == ($pending[0] | del(.delivery_id))
  ' "$ledger" >/dev/null || fail "Existing delivery receipt conflicts with pending identity"
else
  tmp_ledger=$(mktemp)
  jq -S --arg id "$delivery_id" --slurpfile pending "$pending" '
    .deliveries[$id] = ($pending[0] | del(.delivery_id))
  ' "$ledger" >"$tmp_ledger"
  mv "$tmp_ledger" "$ledger"
fi

if jq -e --arg id "$delivery_id" '.receipts[$id] != null' "$detailed_ledger" >/dev/null; then
  jq -e --arg id "$delivery_id" --slurpfile pending "$pending" --slurpfile receipt "$provider_receipt" '
    .receipts[$id] == {delivery:($pending[0] | del(.delivery_id)),publication:$receipt[0]}
  ' "$detailed_ledger" >/dev/null ||
    fail "Existing provider publication evidence conflicts with measured release"
else
  tmp_detailed=$(mktemp)
  jq -S --arg id "$delivery_id" --slurpfile pending "$pending" --slurpfile receipt "$provider_receipt" '
    .receipts[$id] = {delivery:($pending[0] | del(.delivery_id)),publication:$receipt[0]}
  ' "$detailed_ledger" >"$tmp_detailed"
  mv "$tmp_detailed" "$detailed_ledger"
fi

rm "$pending" "$regeneration_receipt"
TARGET_REPOSITORY=$TARGET_REPOSITORY scripts/validate-provider-delivery-state.sh
{
  echo "changed=true"
  echo "delivery_id=$delivery_id"
} >>"$GITHUB_OUTPUT"
