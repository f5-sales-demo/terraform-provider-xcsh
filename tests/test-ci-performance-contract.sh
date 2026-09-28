#!/usr/bin/env bash
set -euo pipefail

root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
benchmark="$root/.github/workflows/workload-benchmark.yml"
build="$root/.github/workflows/_build-test.yml"
ci="$root/.github/workflows/ci.yml"
provider="$root/.github/workflows/_generate-provider.yml"
docs="$root/.github/workflows/_generate-docs.yml"
merge="$root/.github/workflows/on-merge.yml"

fail() {
  printf 'FAIL: %s\n' "$1" >&2
  exit 1
}
require() { grep -Fq -- "$2" "$1" || fail "$1 is missing: $2"; }

require "$benchmark" 'source_sha:'
require "$benchmark" 'expected_image_digest:'
require "$benchmark" 'dkr\.ecr\.us-east-1\.amazonaws\.com'
require "$benchmark" 'cache_state:'
require "$benchmark" 'pair_id:'
require "$benchmark" 'concurrency:'
require "$benchmark" 'phase:'
require "$benchmark" 'runner-profile.py'
require "$benchmark" 'terraform-provider-xcsh-compute'
require "$benchmark" 'runs-on: ${{ needs.validate.outputs.runner_label }}'
require "$benchmark" '$candidate.memory.peak_limit_ratio < 0.8'
require "$benchmark" '$candidate.duration_seconds <= ($hosted.duration_seconds * 0.8)'
require "$root/scripts/run-provider-benchmark.sh" 'script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)'
require "$root/scripts/run-provider-benchmark.sh" '"$script_dir/run-provider-benchmark-phase.sh"'
require "$root/scripts/run-provider-benchmark.sh" '"${observed_image##*@}" != "${expected_image##*@}"'
require "$root/scripts/run-provider-benchmark.sh" 'runner image mismatch: expected'
require "$root/scripts/run-provider-benchmark.sh" 'dkr\.ecr\.us-east-1\.amazonaws\.com'
require "$root/scripts/run-provider-benchmark.sh" 'image-resident runner-profile differs'
require "$root/scripts/run-provider-benchmark.sh" 'export GOGC=20'
require "$root/scripts/run-provider-benchmark.sh" 'export GOMEMLIMIT=4GiB'
require "$root/scripts/run-provider-benchmark.sh" 'export GOMAXPROCS="$concurrency"'
require "$root/scripts/run-provider-benchmark-phase.sh" 'sha256sum <"$evidence_dir/package-inventory.txt"'
require "$root/scripts/run-provider-benchmark-phase.sh" 'sha256sum <"$evidence_dir/normalized-output.txt"'
require "$root/scripts/run-provider-benchmark-phase.sh" 'sha256sum <"$evidence_dir/worktree-output.patch"'
require "$root/scripts/run-provider-benchmark-phase.sh" "s/-p [0-9]+/-p <concurrency>/g"
require "$root/scripts/run-provider-benchmark-phase.sh" 'go run -p "$concurrency" tools/generate-all-schemas.go'

# Stable PR gate with independent compute shards and explicit reusable inputs.
require "$build" 'runner-label:'
require "$build" 'go-concurrency:'
require "$build" 'default: terraform-provider-xcsh-compute'
require "$build" 'default: 4'
require "$build" 'build:'
require "$build" 'vet:'
require "$build" 'race:'
require "$build" 'lint:'
require "$build" 'aggregate:'
require "$build" 'name: Build and Test'
require "$build" "'16GiB'"
require "$build" 'runner-profile'
require "$build" 'retention-days: 30'
require "$ci" 'group: ci-${{ github.event.pull_request.head.ref || github.ref_name }}'

# Read-only generation jobs hand source-bound artifacts forward; the publisher
# applies them and never performs generation or Go compilation itself.
for workflow in "$provider" "$docs"; do
  require "$workflow" 'runner-label:'
  require "$workflow" 'go-concurrency:'
  require "$workflow" 'generation-artifact-name:'
  require "$workflow" 'source-sha:'
  require "$workflow" 'digest:'
  require "$workflow" 'changed:'
  require "$workflow" 'generation-manifest.json'
  require "$workflow" 'retention-days: 30'
done
require "$provider" 'provider-generation.patch'
require "$provider" 'artifact_kind:$artifact_kind'
require "$docs" 'provider-artifact-name:'
require "$docs" 'combined-generation.patch'
require "$docs" 'Verify and apply provider generation artifact'
require "$merge" 'needs: [detect-changes, build-test, regenerate-provider]'
require "$merge" 'outputs.generation-artifact-name'
require "$merge" 'Verify and apply combined generation artifact'
require "$merge" 'needs: [detect-changes, generation-state, create-regeneration-pr]'

publisher=$(sed -n '/^  create-regeneration-pr:/,/^  # STEP 7:/p' "$merge")
if grep -Eq 'go run tools/generate-all-schemas|go build|go vet|generate-provider-docs' <<<"$publisher"; then
  fail 'create-regeneration-pr must apply the verified artifact without regeneration or compilation'
fi
if rg -n ' \+ {6,}' "$build" "$provider" "$docs" "$ci" "$benchmark" "$merge" >/dev/null; then
  fail 'workflow shell blocks contain a collapsed continuation marker'
fi

bash -n "$root/scripts/run-provider-benchmark.sh"
bash -n "$root/scripts/run-provider-benchmark-phase.sh"
printf 'CI performance workflow contract tests passed\n'
