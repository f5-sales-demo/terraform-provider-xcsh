#!/bin/bash

set -e

# Setup temp directory for the test workspace
TEMP_DIR=$(mktemp -d)
trap 'rm -rf "$TEMP_DIR"' EXIT

# Mock path setup
export PATH="$TEMP_DIR:$PATH"
export SPEC_FILE="$TEMP_DIR/spec-version.txt"

SCRIPT_UNDER_TEST="$(pwd)/scripts/check-spec-version-freshness.sh"

run_test() {
  local name="$1"
  local mock_gh_body="$2"
  local expected_exit="$3"
  local expected_stdout="$4"
  local pending_file="${5:-$TEMP_DIR/missing-pending.json}"

  echo "Running test: $name"

  # Create mock gh
  cat <<'EOF' >"$TEMP_DIR/gh"
#!/bin/bash
EOF
  echo "$mock_gh_body" >>"$TEMP_DIR/gh"
  chmod +x "$TEMP_DIR/gh"

  # Run the script and capture output/exit code
  set +e
  output=$("$SCRIPT_UNDER_TEST" "$SPEC_FILE" "$pending_file" 2>&1)
  exit_code=$?
  set -e

  if [ "$exit_code" -ne "$expected_exit" ]; then
    echo "❌ FAILED: $name"
    echo "Expected exit code $expected_exit, got $exit_code"
    echo "Output: $output"
    exit 1
  fi

  if [ -n "$expected_stdout" ]; then
    if ! echo "$output" | grep -q "$expected_stdout"; then
      echo "❌ FAILED: $name"
      echo "Expected output to contain '$expected_stdout', but it did not."
      echo "Output: $output"
      exit 1
    fi
  fi

  echo "✅ PASSED: $name"
}

# Test 1: Matching Versions
echo "v1.2.3" >"$SPEC_FILE"
run_test "Matching Versions" \
  "echo 'v1.2.3'; exit 0" \
  0 \
  "Spec version is up to date"

# Test 2: Stale Pin
echo "v1.2.3" >"$SPEC_FILE"
run_test "Stale Pin" \
  "echo 'v1.2.4'; exit 0" \
  1 \
  "::error::tools/spec-version.txt (v1.2.3) lags latest upstream api-specs-enriched release (v1.2.4)"

# Test 3: An exact pending delivery permits only an explicitly authorized
# workflow-recovery PR to repair the release path before advancing the pin.
commit=$(printf 'a%.0s' {1..40})
canonical=$(jq -cnS \
  --arg commit "$commit" \
  --arg event_type enriched-specs-updated \
  --arg source f5-sales-demo/api-specs-enriched \
  --arg tag v1.2.3 \
  --arg target f5-sales-demo/terraform-provider-xcsh \
  --arg version 1.2.3 \
  '{commit:$commit,event_type:$event_type,source:$source,tag:$tag,target:$target,version:$version}')
delivery_id=$(printf '%s' "$canonical" | shasum -a 256 | awk '{print $1}')
pending_file="$TEMP_DIR/pending.json"
jq -nS \
  --arg id "$delivery_id" \
  --arg tag v1.2.3 \
  --arg commit "$commit" \
  --arg version 1.2.3 \
  '{delivery_id:$id,release_tag:$tag,target_commit:$commit,version:$version}' \
  >"$pending_file"
ALLOW_PENDING_DELIVERY=true run_test "Exact Pending Recovery" \
  "echo 'v1.2.4'; exit 0" \
  0 \
  "Allowing workflow-only recovery for exact pending delivery" \
  "$pending_file"

# Test 4: Recovery infrastructure may repair a completed delivery while the pin is stale.
ALLOW_RECOVERY_INFRASTRUCTURE=true run_test "Completed Delivery Infrastructure Recovery" \
  "echo 'v1.2.4'; exit 0" \
  0 \
  "Allowing delivery-infrastructure recovery with no pending delivery"

# Test 5: Recovery infrastructure cannot bypass an extant pending delivery.
ALLOW_RECOVERY_INFRASTRUCTURE=true run_test "Pending Delivery Requires Exact Recovery" \
  "echo 'v1.2.4'; exit 0" \
  1 \
  "Pending delivery requires exact pending-delivery recovery validation" \
  "$pending_file"

# Test 6: A forged pending identity remains a hard failure.
jq '.delivery_id = ("0" * 64)' "$pending_file" >"$TEMP_DIR/forged-pending.json"
ALLOW_PENDING_DELIVERY=true run_test "Forged Pending Recovery" \
  "echo 'v1.2.4'; exit 0" \
  1 \
  "Pending delivery ID is not canonical" \
  "$TEMP_DIR/forged-pending.json"

# Test 7: API Failure
echo "v1.2.3" >"$SPEC_FILE"
run_test "API Failure" \
  "echo 'API Error'; exit 1" \
  2 \
  "Could not query upstream releases from GitHub API"

# Test 8: Malformed Tag (Null/Empty)
echo "v1.2.3" >"$SPEC_FILE"
run_test "Malformed Tag" \
  "echo 'null'; exit 0" \
  2 \
  "Could not query upstream releases from GitHub API or received invalid/empty tag"

# Test 9: Missing gh CLI
rm -f "$TEMP_DIR/gh"
echo "v1.2.3" >"$SPEC_FILE"
cp "$SCRIPT_UNDER_TEST" "$TEMP_DIR/check_no_gh.sh"
sed -i 's/command -v gh/false/g' "$TEMP_DIR/check_no_gh.sh"
set +e
output=$("$TEMP_DIR/check_no_gh.sh" "$SPEC_FILE" 2>&1)
exit_code=$?
set -e

if [ "$exit_code" -ne 2 ]; then
  echo "❌ FAILED: Missing gh CLI"
  echo "Expected exit code 2, got $exit_code"
  echo "Output: $output"
  exit 1
fi
if ! echo "$output" | grep -q "gh CLI not available"; then
  echo "❌ FAILED: Missing gh CLI output check"
  echo "Output: $output"
  exit 1
fi
echo "✅ PASSED: Missing gh CLI"

echo "All tests passed successfully!"
