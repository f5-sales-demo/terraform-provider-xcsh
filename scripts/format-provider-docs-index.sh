#!/usr/bin/env bash
set -euo pipefail

if [ "${PROVIDER_FORK_ISOLATION:-false}" = true ]; then
  npx --yes @biomejs/biome@2.5.6 format --write docs/terraform-llms-index.json
else
  test "$(biome --version)" = 'Version: 2.5.6' || test "$(biome --version)" = '2.5.6'
  biome format --write docs/terraform-llms-index.json
fi
