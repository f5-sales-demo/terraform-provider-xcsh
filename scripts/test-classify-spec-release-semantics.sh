#!/usr/bin/env bash

set -euo pipefail

repo_root=$(git rev-parse --show-toplevel)
classifier="$repo_root/scripts/classify-spec-release-semantics.sh"
test_root=$(mktemp -d)
trap 'rm -rf "$test_root"' EXIT

assert_output() {
  local current=$1
  local incoming=$2
  local expected_breaking=$3
  local expected_title=$4
  local output="$test_root/output"

  : >"$output"
  GITHUB_OUTPUT="$output" "$classifier" "$current" "$incoming"
  grep -Fx "breaking=$expected_breaking" "$output" >/dev/null
  grep -Fx "pr_title=$expected_title" "$output" >/dev/null
}

assert_output v2.1.225 v3.0.0 true \
  'feat!: update F5 Distributed Cloud OpenAPI specifications'
grep -Fx \
  'breaking_footer=BREAKING CHANGE: updates the provider to API v3.0.0 and removes superseded contract fields.' \
  "$test_root/output" >/dev/null

assert_output v2.1.225 v2.2.0 false \
  'feat: update F5 Distributed Cloud OpenAPI specifications'
grep -Fx 'breaking_footer=' "$test_root/output" >/dev/null

for stale in v2.1.225 v2.1.224 v1.99.99 malformed; do
  if GITHUB_OUTPUT="$test_root/output" "$classifier" v2.1.225 "$stale" >/dev/null 2>&1; then
    printf 'classifier accepted stale or malformed version %s\n' "$stale" >&2
    exit 1
  fi
done

printf 'spec release semantic classification tests passed\n'
