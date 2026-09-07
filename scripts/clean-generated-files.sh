#!/usr/bin/env bash
set -euo pipefail

provider_dir=${1:-internal/provider}
client_dir=${2:-internal/client}

remove_generated_matches() {
  local directory=$1 pattern=$2 file
  [ -d "$directory" ] || return 0
  while IFS= read -r -d '' file; do
    if head -c 500 "$file" | grep -Fq 'DO NOT EDIT'; then
      rm -f -- "$file"
    fi
  done < <(find "$directory" -maxdepth 1 -type f -name "$pattern" -print0)
}

remove_generated_matches "$provider_dir" '*_resource.go'
remove_generated_matches "$provider_dir" '*_data_source.go'
remove_generated_matches "$provider_dir" '*_action.go'
remove_generated_matches "$provider_dir" 'provider.go'
remove_generated_matches "$client_dir" '*_types.go'
