#!/usr/bin/env bash

set -euo pipefail

if [ "$#" -ne 2 ]; then
  echo "usage: $0 <vMAJOR.MINOR.PATCH> <output.tar.gz>" >&2
  exit 2
fi

tag=$1
output=$2
[[ "$tag" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]] || {
  echo "release tag must be vMAJOR.MINOR.PATCH" >&2
  exit 1
}

commit=$(git rev-list -n 1 "$tag^{commit}")
[ "$commit" = "$(git rev-parse HEAD)" ] || {
  echo "release tag does not identify the checked-out commit" >&2
  exit 1
}
source_date_epoch=$(git show -s --format=%ct "$commit")
[[ "$source_date_epoch" =~ ^[0-9]+$ ]] || {
  echo "release commit timestamp is malformed" >&2
  exit 1
}

required_paths=(
  docs/resources
  docs/data-sources
  docs/functions
  docs/guides
  docs/index.md
  docs/specifications/api
  tools/minimum-configs.json
)
for path in "${required_paths[@]}"; do
  [ -e "$path" ] || {
    echo "required MCP input is missing: $path" >&2
    exit 1
  }
done

for path in "${required_paths[@]}"; do
  if find "$path" -type l -print -quit | grep -q .; then
    echo "MCP inputs must not contain symbolic links: $path" >&2
    exit 1
  fi
  if find "$path" ! -type f ! -type d -print -quit | grep -q .; then
    echo "MCP inputs contain a non-regular filesystem entry: $path" >&2
    exit 1
  fi
done

output_dir=$(dirname "$output")
mkdir -p "$output_dir"
output=$(cd "$output_dir" && pwd)/$(basename "$output")
file_list=$(mktemp)
temporary_index=$(mktemp)
archive=$(mktemp "${output}.tmp.XXXXXX")
rm "$temporary_index"
trap 'rm -f "$file_list" "$temporary_index" "$archive"' EXIT

find "${required_paths[@]}" -type f -print0 | LC_ALL=C sort -z >"$file_list"
[ -s "$file_list" ] || {
  echo "MCP input file list is empty" >&2
  exit 1
}

GIT_INDEX_FILE=$temporary_index git read-tree --empty
GIT_INDEX_FILE=$temporary_index git -c core.autocrlf=false add -f -- "${required_paths[@]}"
while IFS= read -r -d '' file; do
  cmp "$file" <(GIT_INDEX_FILE=$temporary_index git show ":$file") >/dev/null || {
    echo "Git clean filters changed MCP input bytes: $file" >&2
    exit 1
  }
done <"$file_list"
tree=$(GIT_INDEX_FILE=$temporary_index git write-tree)
git archive \
  --format=tar \
  --prefix=mcp-data/ \
  --mtime="@${source_date_epoch}" \
  "$tree" | gzip -n -9 >"$archive"

gzip -t "$archive"
mv "$archive" "$output"
