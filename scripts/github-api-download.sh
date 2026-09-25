#!/usr/bin/env bash
# Download one GitHub API response with bounded retry, validation, and atomic promotion.
set -euo pipefail

usage() {
  echo "Usage: $0 [--accept MEDIA_TYPE] [--json-object] [--sha256 sha256:HEX] ENDPOINT DESTINATION" >&2
  exit 64
}

accept='application/vnd.github+json'
expected_sha=''
validate_json_object=false
while [ "$#" -gt 0 ]; do
  case "$1" in
  --accept)
    [ "$#" -ge 2 ] || usage
    accept=$2
    shift 2
    ;;
  --json-object)
    validate_json_object=true
    shift
    ;;
  --sha256)
    [ "$#" -ge 2 ] || usage
    expected_sha=$2
    shift 2
    ;;
  --)
    shift
    break
    ;;
  -*) usage ;;
  *) break ;;
  esac
done
[ "$#" -eq 2 ] || usage
endpoint=$1
destination=$2

[ -n "$endpoint" ] || usage
[ -n "$accept" ] || usage
if [ -n "$expected_sha" ] && [[ ! "$expected_sha" =~ ^sha256:[0-9a-f]{64}$ ]]; then
  echo "::error::Expected SHA-256 must use sha256:<64 lowercase hex characters>" >&2
  exit 64
fi
command -v gh >/dev/null 2>&1 || {
  echo "::error::gh CLI is required for GitHub API downloads" >&2
  exit 69
}
if [ "$validate_json_object" = true ]; then
  command -v jq >/dev/null 2>&1 || {
    echo "::error::jq is required for JSON response validation" >&2
    exit 69
  }
fi
if [ -n "$expected_sha" ]; then
  command -v shasum >/dev/null 2>&1 || {
    echo "::error::shasum is required for release-asset verification" >&2
    exit 69
  }
fi

destination_dir=$(dirname -- "$destination")
[ -d "$destination_dir" ] || {
  echo "::error::Destination directory does not exist: ${destination_dir}" >&2
  exit 73
}
candidate=$(mktemp "${destination}.download.XXXXXX")
error_log=$(mktemp "${destination}.error.XXXXXX")
trap 'rm -f -- "$candidate" "$error_log"' EXIT

readonly max_attempts=4
attempt=1
backoff_seconds=1
last_failure='unknown failure'
while [ "$attempt" -le "$max_attempts" ]; do
  : >"$candidate"
  : >"$error_log"
  if gh api "$endpoint" -H "Accept: ${accept}" >"$candidate" 2>"$error_log"; then
    if [ "$validate_json_object" = true ] && ! jq -e 'type == "object"' "$candidate" >/dev/null 2>&1; then
      last_failure='response was not a complete JSON object'
    elif [ -n "$expected_sha" ]; then
      actual_sha="sha256:$(shasum -a 256 "$candidate" | awk '{print $1}')"
      if [ "$actual_sha" = "$expected_sha" ]; then
        mv -f -- "$candidate" "$destination"
        exit 0
      fi
      last_failure='SHA-256 mismatch'
    else
      mv -f -- "$candidate" "$destination"
      exit 0
    fi
  else
    status=$?
    last_failure="GitHub API command exited ${status}"
  fi

  if [ "$attempt" -lt "$max_attempts" ]; then
    echo "::warning::GitHub API download attempt ${attempt}/${max_attempts} failed (${last_failure}); retrying in ${backoff_seconds}s" >&2
    sleep "$backoff_seconds"
    backoff_seconds=$((backoff_seconds * 2))
  fi
  attempt=$((attempt + 1))
done

echo "::error::GitHub API download failed after ${max_attempts} attempts (${last_failure}): ${endpoint}" >&2
exit 1
