#!/usr/bin/env bash
set -euo pipefail

state=${MOCK_GH_STATE:?MOCK_GH_STATE is required}
attempt=0
if [ -f "$state" ]; then
  read -r attempt <"$state"
fi
attempt=$((attempt + 1))
printf '%s\n' "$attempt" >"$state"

case "${MOCK_GH_MODE:?MOCK_GH_MODE is required}" in
transport_then_valid)
  if [ "$attempt" -eq 1 ]; then
    printf 'partial-response'
    printf '%s\n' 'read: connection reset by peer' >&2
    exit 1
  fi
  printf '%s' "${MOCK_GH_PAYLOAD:?MOCK_GH_PAYLOAD is required}"
  ;;
malformed_then_valid)
  if [ "$attempt" -eq 1 ]; then
    printf '{"partial":'
  else
    printf '%s' "${MOCK_GH_PAYLOAD:?MOCK_GH_PAYLOAD is required}"
  fi
  ;;
checksum_then_valid)
  if [ "$attempt" -eq 1 ]; then
    printf 'corrupt'
  else
    printf '%s' "${MOCK_GH_PAYLOAD:?MOCK_GH_PAYLOAD is required}"
  fi
  ;;
persistent_checksum_mismatch)
  printf 'corrupt'
  ;;
*)
  printf 'unexpected mock mode: %s\n' "$MOCK_GH_MODE" >&2
  exit 2
  ;;
esac
