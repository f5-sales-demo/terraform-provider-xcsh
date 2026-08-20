#!/usr/bin/env bash

set -euo pipefail

release_ready=false
base_ref=
while [ "$#" -gt 0 ]; do
  case "$1" in
  --release-ready)
    release_ready=true
    shift
    ;;
  --base-ref)
    [ "$#" -ge 2 ] || {
      echo "--base-ref requires a value" >&2
      exit 2
    }
    base_ref=$2
    shift 2
    ;;
  *)
    echo "unknown argument: $1" >&2
    exit 2
    ;;
  esac
done

target_repository=${TARGET_REPOSITORY:-f5-sales-demo/terraform-provider-xcsh}
ledger=tools/spec-deliveries.json
detailed=tools/provider-publication-receipts.json
pending=tools/spec-pending-delivery.json
regeneration=tools/spec-regeneration-receipt.json
pin=tools/spec-release.json

fail() {
  echo "delivery state invalid: $*" >&2
  exit 1
}

canonical_id() {
  local version=$1 tag=$2 commit=$3 canonical
  canonical=$(jq -cnS \
    --arg commit "$commit" \
    --arg event_type enriched-specs-updated \
    --arg source f5-sales-demo/api-specs-enriched \
    --arg tag "$tag" \
    --arg target "$target_repository" \
    --arg version "$version" \
    '{commit:$commit,event_type:$event_type,source:$source,tag:$tag,target:$target,version:$version}')
  printf '%s' "$canonical" | shasum -a 256 | awk '{print $1}'
}

jq -e '
  type == "object" and (keys | sort) == ["deliveries", "version"] and
  .version == 1 and (.deliveries | type == "object")
' "$ledger" >/dev/null || fail "common ledger is malformed"
jq -e '
  type == "object" and (keys | sort) == ["receipts", "version"] and
  .version == 1 and (.receipts | type == "object")
' "$detailed" >/dev/null || fail "publication ledger is malformed"
jq -en --slurpfile common "$ledger" --slurpfile evidence "$detailed" '
  ($common[0].deliveries | keys | sort) == ($evidence[0].receipts | keys | sort)
' >/dev/null || fail "common and publication ledgers have different keys"
jq -e '
  def unique_values(values): (values | length) == (values | unique | length);
  unique_values([.deliveries[].release_tag]) and
  unique_values([.deliveries[].version]) and
  unique_values([.deliveries[].target_commit])
' "$ledger" >/dev/null ||
  fail "durable spec deliveries reuse a release tag, version, or commit"
jq -e '
  def unique_values(values): (values | length) == (values | unique | length);
  unique_values([.receipts[].publication.tag]) and
  unique_values([.receipts[].publication.version]) and
  unique_values([.receipts[].publication.commit]) and
  unique_values([.receipts[].publication.spec_release_sha256])
' "$detailed" >/dev/null ||
  fail "durable receipts reuse a provider publication or spec-release binding"

while IFS= read -r entry; do
  key=$(jq -r '.key' <<<"$entry")
  jq -e '
    (.value | type == "object") and
    (.value | keys | sort) == ["release_tag", "target_commit", "version"] and
    .value.release_tag == ("v" + .value.version) and
    (.value.target_commit | test("^[0-9a-f]{40}$"))
  ' <<<"$entry" >/dev/null || fail "common ledger entry is malformed"
  expected=$(canonical_id \
    "$(jq -r '.value.version' <<<"$entry")" \
    "$(jq -r '.value.release_tag' <<<"$entry")" \
    "$(jq -r '.value.target_commit' <<<"$entry")")
  [ "$key" = "$expected" ] || fail "common ledger key is not canonical"
done < <(jq -c '.deliveries | to_entries[]' "$ledger")

while IFS= read -r entry; do
  key=$(jq -r '.key' <<<"$entry")
  common=$(jq -c --arg id "$key" '.deliveries[$id]' "$ledger")
  jq -e --argjson common "$common" '
    (.value | type == "object") and
    (.value | keys | sort) == ["delivery", "publication"] and
    .value.delivery == $common and
    (.value.publication | type == "object") and
    (.value.publication | keys | sort) == ["assets", "commit", "spec_release_sha256", "tag", "version"] and
    (.value.publication.commit | test("^[0-9a-f]{40}$")) and
    (.value.publication.spec_release_sha256 | test("^[0-9a-f]{64}$")) and
    .value.publication.tag == ("v" + .value.publication.version) and
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
  ' <<<"$entry" >/dev/null || fail "publication evidence is malformed or cross-bound"
done < <(jq -c '.receipts | to_entries[]' "$detailed")

jq -e '
  type == "object" and
  (keys | sort) == ["assets", "release_tag", "target_commit", "version"] and
  .release_tag == ("v" + .version) and
  (.target_commit | test("^[0-9a-f]{40}$")) and
  (.assets | keys | sort) == [
    "api-catalog.json",
    ("f5xc-api-specs-" + .release_tag + ".zip"),
    "index.json", "minimal-export-defaults.json", "openapi.json",
    "smsv2-contract-manifest.json", "smsv2-contract.json", "smsv2-evidence-receipt.json"
  ] and
  ([.assets[] | test("^sha256:[0-9a-f]{64}$")] | all)
' "$pin" >/dev/null || fail "spec release pin is malformed"
[ "$(tr -d '[:space:]' <tools/spec-version.txt)" = "$(jq -r '.release_tag' "$pin")" ] ||
  fail "spec version marker and release pin disagree"

if [ -e "$pending" ]; then
  jq -e '
    type == "object" and
    (keys | sort) == ["delivery_id", "release_tag", "target_commit", "version"] and
    (.delivery_id | test("^[0-9a-f]{64}$")) and
    .release_tag == ("v" + .version) and
    (.target_commit | test("^[0-9a-f]{40}$"))
  ' "$pending" >/dev/null || fail "pending delivery is malformed"
  expected=$(canonical_id \
    "$(jq -r '.version' "$pending")" \
    "$(jq -r '.release_tag' "$pending")" \
    "$(jq -r '.target_commit' "$pending")")
  [ "$(jq -r '.delivery_id' "$pending")" = "$expected" ] ||
    fail "pending delivery ID is not canonical"
  jq -e --slurpfile pending "$pending" '
    .release_tag == $pending[0].release_tag and
    .target_commit == $pending[0].target_commit and
    .version == $pending[0].version
  ' "$pin" >/dev/null || fail "pending delivery and spec pin disagree"
  jq -e --arg id "$expected" '.deliveries[$id] == null' "$ledger" >/dev/null ||
    fail "pending delivery is already durable"
  jq -e --arg tag "$(jq -r '.release_tag' "$pending")" \
    '[.deliveries[] | select(.release_tag == $tag)] | length == 0' "$ledger" >/dev/null ||
    fail "pending release tag already exists in the durable ledger"

  if [ -e "$regeneration" ]; then
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
    ' "$regeneration" >/dev/null || fail "regeneration receipt is malformed or cross-bound"
    source_commit=$(jq -r '.source_commit' "$regeneration")
    if [ "$release_ready" = true ]; then
      receipt_status=$(git diff --name-status HEAD^ HEAD -- "$regeneration" |
        awk 'NR == 1 { print $1 }')
      case "$receipt_status" in
      A)
        ;;
      M)
        previous_receipt=$(git show "HEAD^:$regeneration") ||
          fail "pre-release regeneration receipt could not be read"
        jq -e --slurpfile current "$regeneration" '
          type == "object" and
          (keys | sort) == [
            "delivery_id", "release_tag", "source_commit", "spec_release_sha256",
            "target_commit", "version"
          ] and
          (.source_commit | test("^[0-9a-f]{40}$")) and
          .source_commit != $current[0].source_commit and
          del(.source_commit) == ($current[0] | del(.source_commit))
        ' <<<"$previous_receipt" >/dev/null ||
          fail "regeneration receipt may change only its source commit during recovery"
        ;;
      *)
        fail "release commit did not introduce or exactly rebind the regeneration receipt"
        ;;
      esac
      [ "$source_commit" = "$(git rev-parse HEAD^)" ] ||
        fail "regeneration receipt source is not the exact release parent"
    else
      # `source_commit` identifies the source event that was regenerated. A
      # source PR may be squash-merged, in which case GitHub puts a distinct
      # squash commit on main and the attested source is intentionally not an
      # ancestor of HEAD. The on-merge workflow binds this identity to the
      # exact associated PR and release commit; this reusable check verifies
      # that the retained identity names a real commit without imposing an
      # incompatible merge topology.
      git cat-file -e "${source_commit}^{commit}" 2>/dev/null ||
        fail "regeneration source commit is unavailable"
    fi
  elif [ "$release_ready" = true ]; then
    fail "pending delivery has no regeneration receipt"
  fi
else
  [ ! -e "$regeneration" ] ||
    fail "regeneration receipt exists without a pending delivery"
  pin_sha=$(shasum -a 256 "$pin" | awk '{print $1}')
  delivery_count=$(jq '.deliveries | length' "$ledger")
  if [ "$delivery_count" -eq 0 ]; then
    # Bootstrap attestation: before any delivery is receipted there is no ledger
    # entry to bind the pin to, so the pin's own digest is the only thing that
    # makes it tamper-evident. This constant therefore moves with tools/
    # spec-release.json on every specification bump, and the two must be updated
    # together — see the note in that file's PR checklist.
    #
    # It stops being load-bearing once api-specs-enriched publishes a
    # publication receipt per release (f5-sales-demo/api-specs-enriched#1321):
    # deliveries become receipted, delivery_count goes above zero, and the
    # ledger-binding branch below takes over permanently.
    [ "$pin_sha" = "76c721aaf04c39f281d651a3a46dcb6853407cbd98861f8f0f82afc3b36d69cd" ] ||
      fail "unreceipted bootstrap pin differs from the measured v2.1.218 baseline"
  else
    jq -en \
      --arg pin_sha "$pin_sha" \
      --slurpfile pin "$pin" \
      --slurpfile evidence "$detailed" '
      any($evidence[0].receipts[];
        .delivery.release_tag == $pin[0].release_tag and
        .delivery.target_commit == $pin[0].target_commit and
        .delivery.version == $pin[0].version and
        .publication.spec_release_sha256 == $pin_sha)
    ' >/dev/null || fail "current spec pin is not bound to durable publication evidence"
  fi
fi

if [ -n "$base_ref" ] && git cat-file -e "${base_ref}:${ledger}" 2>/dev/null; then
  old_common=$(mktemp)
  old_detailed=$(mktemp)
  trap 'rm -f "$old_common" "$old_detailed"' EXIT
  git show "${base_ref}:${ledger}" >"$old_common"
  git show "${base_ref}:${detailed}" >"$old_detailed"
  jq -en --slurpfile old "$old_common" --slurpfile new "$ledger" '
    all($old[0].deliveries | to_entries[]; $new[0].deliveries[.key] == .value)
  ' >/dev/null || fail "common ledger changed or deleted an existing entry"
  jq -en --slurpfile old "$old_detailed" --slurpfile new "$detailed" '
    all($old[0].receipts | to_entries[]; $new[0].receipts[.key] == .value)
  ' >/dev/null || fail "publication ledger changed or deleted existing evidence"
fi
