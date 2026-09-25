#!/usr/bin/env bash
set -euo pipefail
printf '%s\n' "${1:?sleep duration is required}" >>"${MOCK_SLEEP_LOG:?MOCK_SLEEP_LOG is required}"
