#!/usr/bin/env bash
# Hermetic tests for the staged generated-file constitution check.
set -euo pipefail

REPO_ROOT=$(cd "$(dirname "$0")/.." && pwd)
SCRIPT="${REPO_ROOT}/scripts/check-no-generated-files.sh"

FAIL=0
WORK=$(mktemp -d)
cleanup() { rm -rf "$WORK"; }
trap cleanup EXIT

new_repo() {
  local dir="${WORK}/$1"
  mkdir -p "$dir"
  git -C "$dir" init -q -b main
  git -C "$dir" config user.email generated-check@test
  git -C "$dir" config user.name "Generated Check Test"
  printf 'baseline\n' >"${dir}/README.md"
  git -C "$dir" add -A
  git -C "$dir" commit -qm baseline
  echo "$dir"
}

stage_file() {
  local dir="$1" path="$2"
  mkdir -p "${dir}/$(dirname "$path")"
  printf 'changed\n' >"${dir}/${path}"
  git -C "$dir" add "$path"
}

run_check() {
  local dir="$1" rc=0
  (cd "$dir" && bash "$SCRIPT") >/dev/null 2>&1 || rc=$?
  echo "$rc"
}

assert_passes() {
  local label="$1" dir="$2" rc
  rc=$(run_check "$dir")
  if [ "$rc" -eq 0 ]; then
    echo "[OK] $label -> allowed"
  else
    echo "[FAIL] $label — expected success, got $rc"
    FAIL=1
  fi
}

assert_rejected() {
  local label="$1" dir="$2" rc
  rc=$(run_check "$dir")
  if [ "$rc" -eq 1 ]; then
    echo "[OK] $label -> rejected"
  else
    echo "[FAIL] $label — expected rejection, got $rc"
    FAIL=1
  fi
}

repo=$(new_repo generated-only)
stage_file "$repo" docs/resources/example.md
assert_rejected "generated documentation without its source" "$repo"

repo=$(new_repo transform-source)
stage_file "$repo" tools/transform-docs.go
stage_file "$repo" docs/resources/example.md
assert_passes "documentation paired with its transformer" "$repo"

repo=$(new_repo template-source)
stage_file "$repo" templates/guides/example.md
stage_file "$repo" docs/guides/example.md
assert_passes "guide paired with its template" "$repo"

repo=$(new_repo unrelated-source)
stage_file "$repo" internal/client/example.go
stage_file "$repo" docs/resources/example.md
assert_rejected "unrelated source cannot authorize generated output" "$repo"

repo=$(new_repo provider-source)
stage_file "$repo" tools/generate-all-schemas.go
stage_file "$repo" internal/provider/example_resource.go
assert_passes "generated provider paired with its generator" "$repo"

repo=$(new_repo description-source)
stage_file "$repo" tools/pkg/description/description.go
stage_file "$repo" internal/provider/example_resource.go
assert_passes "generated provider paired with description normalization" "$repo"

repo=$(new_repo manual-exception)
stage_file "$repo" internal/provider/site_registration_data_source.go
assert_passes "manually maintained provider source" "$repo"

repo=$(new_repo no-files)
assert_passes "repository with no staged changes" "$repo"

if [ "$FAIL" -ne 0 ]; then
  echo "generated-file check tests FAILED"
  exit 1
fi
echo "generated-file check tests passed"
