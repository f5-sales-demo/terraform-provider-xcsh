#!/usr/bin/env bash
set -euo pipefail

root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
benchmark="$root/.github/workflows/workload-benchmark.yml"

fail() { printf 'FAIL: %s\n' "$1" >&2; exit 1; }
require() { grep -Fq -- "$2" "$1" || fail "$1 is missing: $2"; }

require "$benchmark" 'source_sha:'
require "$benchmark" 'expected_image_digest:'
require "$benchmark" 'cache_state:'
require "$benchmark" 'pair_id:'
require "$benchmark" 'concurrency:'
require "$benchmark" 'phase:'
require "$benchmark" 'runner-profile.py'
require "$benchmark" 'terraform-provider-xcsh-32vcpu-candidate'
require "$benchmark" 'runs-on: ${{ needs.validate.outputs.runner_label }}'
require "$root/scripts/run-provider-benchmark.sh" 'script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)'
require "$root/scripts/run-provider-benchmark.sh" '"$script_dir/run-provider-benchmark-phase.sh"'
require "$root/scripts/run-provider-benchmark.sh" '"${observed_image##*@}" == "${expected_image##*@}"'
bash -n "$root/scripts/run-provider-benchmark.sh"
bash -n "$root/scripts/run-provider-benchmark-phase.sh"
printf 'CI performance workflow contract tests passed\n'
