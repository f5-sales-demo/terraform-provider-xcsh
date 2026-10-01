#!/usr/bin/env bash
set -euo pipefail

[ "$#" -eq 8 ] || {
  echo "usage: $0 <source-sha> <expected-image> <runner-kind> <cache-state> <pair-id> <concurrency> <phase> <evidence-dir>" >&2
  exit 2
}

source_sha=$1
expected_image=$2
runner_kind=$3
cache_state=$4
pair_id=$5
concurrency=$6
phase=$7
evidence_dir=$8
script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
[[ "$source_sha" =~ ^[0-9a-f]{40}$ ]]
approved_image_pattern='^(ghcr\.io/f5-sales-demo|f5salesdemoarcca\.azurecr\.io|[0-9]{12}\.dkr\.ecr\.us-east-1\.amazonaws\.com)/self-hosted-runner@sha256:[0-9a-f]{64}$'
[[ "$expected_image" =~ $approved_image_pattern ]]
[[ "$runner_kind" =~ ^(hosted|eks)$ ]]
[[ "$cache_state" =~ ^(cold|warm)$ ]]
[[ "$pair_id" =~ ^[1-5]$ ]]
actual_source_sha=$(git rev-parse HEAD)
if [ "$actual_source_sha" != "$source_sha" ]; then
  echo "source SHA mismatch: expected $source_sha, observed $actual_source_sha" >&2
  exit 1
fi

mkdir -p "$evidence_dir"
export CHECKPOINT_DISABLE=1
# The production build gate deliberately constrains the generated-provider
# compiler footprint on GitHub-hosted runners. Apply the same contract to both
# benchmark sides so a candidate comparison cannot evade the bounded-memory
# correctness gate.
export GOGC=20
export GOMEMLIMIT=4GiB
export GOMAXPROCS="$concurrency"
export GOFLAGS="-p=$concurrency"
observed_image=${RUNNER_IMAGE_DIGEST:-github-hosted}
profiler=.runner-harness/scripts/runner-profile.py
if [ "$runner_kind" = eks ]; then
  if [[ "$observed_image" != "$expected_image" &&
    "${observed_image##*@}" != "${expected_image##*@}" ]]; then
    echo "runner image mismatch: expected $expected_image, observed $observed_image" >&2
    exit 1
  fi
  if [ "$(command -v runner-profile)" != /usr/local/bin/runner-profile ]; then
    echo "runner-profile is not installed at the governed image path" >&2
    exit 1
  fi
  if ! cmp -s /usr/local/bin/runner-profile "$profiler"; then
    echo "image-resident runner-profile differs from the frozen harness copy" >&2
    exit 1
  fi
  profiler=/usr/local/bin/runner-profile
fi

if [ "$cache_state" = cold ]; then
  export GOCACHE="$RUNNER_TEMP/go-build-cold-${pair_id}-${phase}-${concurrency}"
  export GOMODCACHE="$RUNNER_TEMP/go-mod-cold-${pair_id}-${phase}-${concurrency}"
else
  GOCACHE=$(go env GOCACHE)
  GOMODCACHE=$(go env GOMODCACHE)
  export GOCACHE GOMODCACHE
  go mod download
  go test -run '^$' -p "$concurrency" ./internal/... ./tools/... >/dev/null
fi

set +e
python3 "$script_dir/profile_provider_runtime.py" \
  "$evidence_dir" "$source_sha" "$observed_image" "$cache_state" "$concurrency" -- \
  python3 "$profiler" \
  --name "$phase" \
  --output "$evidence_dir/workload-profile.json" \
  --cache-state "$cache_state" \
  --variant "${runner_kind}-p${concurrency}" \
  --pair-id "$pair_id" \
  --repository "$GITHUB_REPOSITORY" \
  --commit "$source_sha" \
  --image-digest "$observed_image" \
  -- python3 "$script_dir/run_with_cache_writeback.py" "$evidence_dir" -- \
  "$script_dir/run-provider-benchmark-phase.sh" "$phase" "$concurrency" "$evidence_dir"
status=$?
set -e

output_digest=$(cat "$evidence_dir/output-digest.txt" 2>/dev/null || printf 'sha256:%064d' 0)
tmp_profile="$evidence_dir/.workload-profile.json.tmp"
jq --arg digest "$output_digest" '.output_digest = $digest' \
  "$evidence_dir/workload-profile.json" >"$tmp_profile"
mv "$tmp_profile" "$evidence_dir/workload-profile.json"

package_digest=$(sha256sum "$evidence_dir/package-inventory.txt" | awk '{print "sha256:" $1}')
tool_versions=$(jq -cn \
  --arg go "$(go env GOVERSION)" \
  --arg terraform "$(terraform version 2>/dev/null | head -1 || true)" \
  --arg shellcheck "$(shellcheck --version 2>/dev/null | awk '/^version:/ {print $2}' || true)" \
  --arg golangci "$(golangci-lint --version 2>/dev/null | awk '{print $4}' || true)" \
  --arg tfplugindocs "$(go version -m "$(command -v tfplugindocs 2>/dev/null)" 2>/dev/null | awk '$1 == "mod" {print $3}' || true)" \
  --arg zizmor "$(zizmor --version 2>/dev/null | awk '{print $NF}' || true)" \
  --arg semgrep "$(semgrep --version 2>/dev/null | tail -1 || true)" \
  --arg pyyaml "$(provider-python -c 'import yaml; print(yaml.__version__)' 2>/dev/null || true)" \
  '{go:$go,terraform:$terraform,shellcheck:$shellcheck,golangci_lint:$golangci,tfplugindocs:$tfplugindocs,zizmor:$zizmor,semgrep:$semgrep,pyyaml:$pyyaml}')
jq -nS \
  --arg source_sha "$source_sha" \
  --arg expected_image_digest "$expected_image" \
  --arg observed_image_digest "$observed_image" \
  --arg runner_kind "$runner_kind" \
  --arg cache_state "$cache_state" \
  --arg pair_id "$pair_id" \
  --arg phase "$phase" \
  --argjson concurrency "$concurrency" \
  --arg package_inventory_digest "$package_digest" \
  --arg output_digest "$output_digest" \
  --argjson tool_versions "$tool_versions" \
  --argjson exit_code "$status" \
  '{schema_version:1,source_sha:$source_sha,expected_image_digest:$expected_image_digest,observed_image_digest:$observed_image_digest,runner_kind:$runner_kind,cache_state:$cache_state,pair_id:$pair_id,phase:$phase,go_concurrency:$concurrency,tool_versions:$tool_versions,package_inventory_digest:$package_inventory_digest,output_digest:$output_digest,exit_code:$exit_code}' \
  >"$evidence_dir/output-manifest.json"
exit "$status"
