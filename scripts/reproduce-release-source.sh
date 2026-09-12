#!/usr/bin/env bash

set -euo pipefail

check_clean=false
case ${1:-} in
"") ;;
--check-clean) check_clean=true ;;
*)
  echo "usage: $0 [--check-clean]" >&2
  exit 2
  ;;
esac
[ "$#" -le 1 ] || {
  echo "usage: $0 [--check-clean]" >&2
  exit 2
}

repository_root=$(git rev-parse --show-toplevel)
cd "$repository_root"

go run tools/generate-all-schemas.go --spec-dir=docs/specifications/api
scripts/go-retry.sh 3 go mod tidy
scripts/generate-provider-docs.sh

if [ "$check_clean" = true ] && [ -n "$(git status --porcelain --untracked-files=all)" ]; then
  echo "release source does not reproduce byte-for-byte from its pinned inputs" >&2
  git status --short >&2
  exit 1
fi
