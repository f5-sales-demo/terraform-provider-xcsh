#!/usr/bin/env bash

set -euo pipefail

if [ "$#" -lt 4 ] || [ "$#" -gt 5 ]; then
  echo "usage: $0 <tag> <commit> <release.json> <asset-dir> [expected-receipt.json]" >&2
  exit 2
fi

tag=$1
commit=$2
release_json=$3
asset_dir=$4
expected_receipt=${5:-}
[[ "$tag" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]] || {
  echo "provider release tag is malformed" >&2
  exit 1
}
[[ "$commit" =~ ^[0-9a-f]{40}$ ]] || {
  echo "provider release commit is malformed" >&2
  exit 1
}
jq -e --arg tag "$tag" '.tag_name == $tag and (.assets | type == "array")' \
  "$release_json" >/dev/null || {
  echo "release metadata does not identify the requested tag" >&2
  exit 1
}

version=${tag#v}
expected_assets=$(printf '%s\n' \
  "terraform-provider-xcsh_${version}_darwin_amd64.zip" \
  "terraform-provider-xcsh_${version}_darwin_arm64.zip" \
  "terraform-provider-xcsh_${version}_freebsd_386.zip" \
  "terraform-provider-xcsh_${version}_freebsd_amd64.zip" \
  "terraform-provider-xcsh_${version}_linux_386.zip" \
  "terraform-provider-xcsh_${version}_linux_amd64.zip" \
  "terraform-provider-xcsh_${version}_linux_arm.zip" \
  "terraform-provider-xcsh_${version}_linux_arm64.zip" \
  "terraform-provider-xcsh_${version}_manifest.json" \
  "terraform-provider-xcsh_${version}_SHA256SUMS" \
  "terraform-provider-xcsh_${version}_SHA256SUMS.sig" \
  "terraform-provider-xcsh_${version}_windows_386.zip" \
  "terraform-provider-xcsh_${version}_windows_amd64.zip" \
  "mcp-data-${version}.tar.gz" | LC_ALL=C sort)
actual_assets=$(jq -r '.assets[].name' "$release_json" | LC_ALL=C sort)
[ "$actual_assets" = "$expected_assets" ] || {
  echo "provider release asset set is incomplete or unexpected" >&2
  diff -u <(printf '%s\n' "$expected_assets") <(printf '%s\n' "$actual_assets") >&2 || true
  exit 1
}

while IFS= read -r name; do
  [ -f "${asset_dir}/${name}" ] || {
    echo "downloaded release asset is missing: $name" >&2
    exit 1
  }
  api_digest=$(jq -er --arg name "$name" '
    [.assets[] | select(.name == $name)] |
    if length == 1 then .[0].digest else empty end
  ' "$release_json")
  [[ "$api_digest" =~ ^sha256:[0-9a-f]{64}$ ]] || {
    echo "GitHub did not report an immutable SHA-256 digest for $name" >&2
    exit 1
  }
  measured=$(shasum -a 256 "${asset_dir}/${name}" | awk '{print $1}')
  [ "$api_digest" = "sha256:${measured}" ] || {
    echo "downloaded bytes differ from the GitHub digest for $name" >&2
    exit 1
  }
done <<<"$expected_assets"

checksums="terraform-provider-xcsh_${version}_SHA256SUMS"
signature="${checksums}.sig"
checksum_names=$(awk '{print $2}' "${asset_dir}/${checksums}" | LC_ALL=C sort)
expected_checksum_names=$(printf '%s\n' "$expected_assets" |
  grep -vE "(_SHA256SUMS|_SHA256SUMS\.sig|^mcp-data-)" |
  LC_ALL=C sort)
[ "$checksum_names" = "$expected_checksum_names" ] || {
  echo "signed checksum manifest has an incomplete or unexpected file set" >&2
  exit 1
}
(
  cd "$asset_dir"
  gpg --batch --verify "$signature" "$checksums"
  shasum -a 256 --check "$checksums" >/dev/null
)

receipt=$(mktemp)
trap 'rm -f "$receipt"' EXIT
assets_json='{}'
while IFS= read -r name; do
  measured=$(shasum -a 256 "${asset_dir}/${name}" | awk '{print $1}')
  assets_json=$(jq -cS --arg name "$name" --arg sha "sha256:$measured" \
    '. + {($name):$sha}' <<<"$assets_json")
done <<<"$expected_assets"
pin_sha=$(shasum -a 256 tools/spec-release.json | awk '{print $1}')
jq -nS \
  --argjson assets "$assets_json" \
  --arg commit "$commit" \
  --arg pin_sha "$pin_sha" \
  --arg tag "$tag" \
  --arg version "$version" \
  '{assets:$assets,commit:$commit,spec_release_sha256:$pin_sha,tag:$tag,version:$version}' \
  >"$receipt"

if [ -n "$expected_receipt" ]; then
  jq -e . "$expected_receipt" >/dev/null || {
    echo "published provider receipt is malformed" >&2
    exit 1
  }
  [ "$(jq -cS . "$expected_receipt")" = "$(jq -cS . "$receipt")" ] || {
    echo "published provider receipt differs from measured release bytes" >&2
    exit 1
  }
fi

cat "$receipt"
