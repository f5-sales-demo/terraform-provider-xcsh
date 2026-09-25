#!/usr/bin/env bash
set -euo pipefail

ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)
SCRIPT="$ROOT/scripts/github-api-download.sh"
TEMP_DIR=$(mktemp -d)
trap 'rm -rf "$TEMP_DIR"' EXIT
BIN_DIR="$TEMP_DIR/bin"
mkdir -p "$BIN_DIR"
ln -s "$ROOT/scripts/testdata/gh-api-download-mock.sh" "$BIN_DIR/gh"
ln -s "$ROOT/scripts/testdata/sleep-recorder.sh" "$BIN_DIR/sleep"
export PATH="$BIN_DIR:$PATH"
export MOCK_GH_STATE="$TEMP_DIR/attempts"
export MOCK_SLEEP_LOG="$TEMP_DIR/sleeps"

fail() {
  printf 'FAIL: %s\n' "$*" >&2
  exit 1
}

assert_file_equals() {
  local expected=$1 path=$2
  [ "$(cat "$path")" = "$expected" ] || fail "$path did not contain expected bytes"
}

reset_mocks() {
  rm -f "$MOCK_GH_STATE" "$MOCK_SLEEP_LOG"
}

payload='verified release asset'
expected_sha="sha256:$(printf '%s' "$payload" | shasum -a 256 | awk '{print $1}')"
destination="$TEMP_DIR/asset"

reset_mocks
export MOCK_GH_MODE=transport_then_valid MOCK_GH_PAYLOAD=$payload
"$SCRIPT" --accept application/octet-stream --sha256 "$expected_sha" \
  repos/example/releases/assets/1 "$destination"
assert_file_equals "$payload" "$destination"
[ "$(cat "$MOCK_GH_STATE")" = 2 ] || fail 'transport failure was not retried once'
assert_file_equals 1 "$MOCK_SLEEP_LOG"

reset_mocks
export MOCK_GH_MODE=malformed_then_valid MOCK_GH_PAYLOAD='{"complete":true}'
"$SCRIPT" --json-object repos/example/releases/tags/v1.2.3 "$destination"
assert_file_equals '{"complete":true}' "$destination"
[ "$(cat "$MOCK_GH_STATE")" = 2 ] || fail 'malformed JSON was not retried once'
assert_file_equals 1 "$MOCK_SLEEP_LOG"

reset_mocks
export MOCK_GH_MODE=checksum_then_valid MOCK_GH_PAYLOAD=$payload
"$SCRIPT" --sha256 "$expected_sha" repos/example/releases/assets/2 "$destination"
assert_file_equals "$payload" "$destination"
[ "$(cat "$MOCK_GH_STATE")" = 2 ] || fail 'checksum mismatch was not retried once'
assert_file_equals 1 "$MOCK_SLEEP_LOG"

reset_mocks
printf 'preserve-me' > "$destination"
export MOCK_GH_MODE=persistent_checksum_mismatch MOCK_GH_PAYLOAD=$payload
set +e
output=$("$SCRIPT" --sha256 "$expected_sha" repos/example/releases/assets/3 "$destination" 2>&1)
status=$?
set -e
[ "$status" -ne 0 ] || fail 'persistent checksum mismatch unexpectedly succeeded'
assert_file_equals preserve-me "$destination"
[ "$(cat "$MOCK_GH_STATE")" = 4 ] || fail 'retry limit was not four attempts'
[ "$(cat "$MOCK_SLEEP_LOG")" = $'1\n2\n4' ] || fail 'backoff was not 1, 2, 4 seconds'
case "$output" in
  *'failed after 4 attempts'*'SHA-256 mismatch'*) ;;
  *) fail 'exhaustion diagnostic was not specific' ;;
esac

printf '%s\n' 'All GitHub API download retry tests passed'
