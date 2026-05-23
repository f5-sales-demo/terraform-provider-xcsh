#!/usr/bin/env bash
# Retry wrapper for Go commands that may fail due to network issues.
# Usage: scripts/go-retry.sh <max_attempts> <command> [args...]
# Example: scripts/go-retry.sh 3 go mod tidy
#          scripts/go-retry.sh 3 go build -v .
#          scripts/go-retry.sh 2 go test -v ./internal/...
set -euo pipefail

MAX_ATTEMPTS="${1:?Usage: go-retry.sh <max_attempts> <command> [args...]}"
shift
CMD=("$@")

attempt=1
while [ "$attempt" -le "$MAX_ATTEMPTS" ]; do
  echo "::group::${CMD[*]} - Attempt $attempt of $MAX_ATTEMPTS"
  if "${CMD[@]}"; then
    echo "::endgroup::"
    exit 0
  fi
  echo "::endgroup::"
  attempt=$((attempt + 1))
  if [ "$attempt" -le "$MAX_ATTEMPTS" ]; then
    sleep_time=$((2 ** attempt * 5))
    echo "::warning::${CMD[0]} failed, retrying in ${sleep_time}s..."
    sleep "$sleep_time"
  fi
done

echo "::error::${CMD[*]} failed after $MAX_ATTEMPTS attempts"
exit 1
