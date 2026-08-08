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

  echo "Running test: $name"

  # Create mock gh
  cat <<'EOF' >"$TEMP_DIR/gh"
#!/bin/bash
EOF
  echo "$mock_gh_body" >>"$TEMP_DIR/gh"
  chmod +x "$TEMP_DIR/gh"

  # Run the script and capture output/exit code
  set +e
  output=$("$SCRIPT_UNDER_TEST" "$SPEC_FILE" 2>&1)
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

# Test 3: API Failure
echo "v1.2.3" >"$SPEC_FILE"
run_test "API Failure" \
  "echo 'API Error'; exit 1" \
  2 \
  "Could not query upstream releases from GitHub API"

# Test 4: Malformed Tag (Null/Empty)
echo "v1.2.3" >"$SPEC_FILE"
run_test "Malformed Tag" \
  "echo 'null'; exit 0" \
  2 \
  "Could not query upstream releases from GitHub API or received invalid/empty tag"

# Test 5: Missing gh CLI
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
