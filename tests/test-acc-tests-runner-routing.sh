#!/usr/bin/env bash
set -euo pipefail

repo_root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
workflow="$repo_root/.github/workflows/acc-tests.yml"

job_runner() {
  local job=$1
  awk -v job="$job" '
    $0 == "  " job ":" { in_job = 1; next }
    in_job && /^  [[:alnum:]_-]+:/ { exit }
    in_job && /^    runs-on:/ {
      sub(/^    runs-on:[[:space:]]*/, "")
      print
      exit
    }
  ' "$workflow"
}

assert_runner() {
  local job=$1 expected=$2 actual
  actual=$(job_runner "$job")
  if [[ "$actual" != "$expected" ]]; then
    printf 'FAIL: %s runner: expected %s, got %s\n' "$job" "$expected" "${actual:-<missing>}" >&2
    return 1
  fi
}

# Pull-request-safe and artifact-only work must remain runnable after retirement
# of the repository-scoped ARC fleet. Live tenant mutation stays isolated.
assert_runner mock-tests ubuntu-latest
assert_runner compare-results ubuntu-latest
assert_runner summary ubuntu-latest
assert_runner real-api-tests managed-socketless
if [[ -n "$(job_runner cleanup)" ]]; then
  printf 'FAIL: separate cleanup job remains\n' >&2
  exit 1
fi
python3 - "$workflow" <<'PY'
import pathlib, sys
text = pathlib.Path(sys.argv[1]).read_text()
start = text.index("      - name: Verify journal-owned cleanup")
step = text[start:text.index("      - name: Upload real API test reports", start)]
assert "if: always()" in step
assert "XCSH_ACC_JOURNAL" in step
assert "second_pass_remaining" in step
assert "-sweep" not in step
PY

echo "PASS: acceptance-test runner routing"
