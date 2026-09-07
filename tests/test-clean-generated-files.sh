#!/usr/bin/env bash
set -euo pipefail

root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT
provider="$work/provider"
client="$work/client"
mkdir -p "$provider" "$client"

printf '%s\n' '// Code generated from a fixture. DO NOT EDIT.' >"$provider/generated_resource.go"
printf '%s\n' '// Code generated from a fixture. DO NOT EDIT.' >"$provider/generated_data_source.go"
printf '%s\n' '// Code generated from a fixture. DO NOT EDIT.' >"$provider/generated_action.go"
printf '%s\n' '// Code generated from a fixture. DO NOT EDIT.' >"$provider/provider.go"
printf '%s\n' '// Code generated from a fixture. DO NOT EDIT.' >"$client/generated_types.go"

printf '%s\n' '// Manually maintained runtime data source.' >"$provider/runtime_data_source.go"
printf '%s\n' '// Manually maintained action.' >"$provider/runtime_action.go"
printf '%s\n' '// Manually maintained client types.' >"$client/runtime_types.go"
printf '%s\n' '// Unrelated source.' >"$provider/helper.go"

"$root/scripts/clean-generated-files.sh" "$provider" "$client"

for removed in \
  "$provider/generated_resource.go" \
  "$provider/generated_data_source.go" \
  "$provider/generated_action.go" \
  "$provider/provider.go" \
  "$client/generated_types.go"; do
  [ ! -e "$removed" ] || {
    printf 'generated file was retained: %s\n' "$removed" >&2
    exit 1
  }
done

for retained in \
  "$provider/runtime_data_source.go" \
  "$provider/runtime_action.go" \
  "$client/runtime_types.go" \
  "$provider/helper.go"; do
  [ -e "$retained" ] || {
    printf 'manually maintained file was deleted: %s\n' "$retained" >&2
    exit 1
  }
done

printf '%s\n' 'generated-file cleanup tests passed'
