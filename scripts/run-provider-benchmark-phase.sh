#!/usr/bin/env bash
set -euo pipefail

[ "$#" -eq 3 ] || {
  echo "usage: $0 <phase> <go-concurrency> <evidence-directory>" >&2
  exit 2
}

phase=$1
concurrency=$2
evidence_dir=$3
[[ "$concurrency" =~ ^(1|2|4|8)$ ]] || {
  echo "go concurrency must be one of 1, 2, 4, or 8" >&2
  exit 2
}

mkdir -p "$evidence_dir"
raw_log="$evidence_dir/workload.log"
set +e
{
  case "$phase" in
    build)
      scripts/go-retry.sh 3 go build -p "$concurrency" ./...
      ;;
    vet)
      scripts/go-retry.sh 3 go vet -p "$concurrency" ./...
      scripts/vet-build-ignored-tools.sh
      ;;
    race)
      XCSH_SPEC_DIR=docs/specifications/api \
        scripts/go-retry.sh 2 go test -p "$concurrency" -timeout=45m -race \
          ./internal/... ./tools/...
      ;;
    provider-generation)
      go run tools/generate-all-schemas.go --spec-dir=docs/specifications/api
      scripts/go-retry.sh 3 go mod tidy
      ;;
    documentation-generation)
      scripts/generate-provider-docs.sh
      ;;
    release-preflight)
      go run tools/generate-all-schemas.go --spec-dir=docs/specifications/api
      scripts/go-retry.sh 3 go mod tidy
      scripts/generate-provider-docs.sh
      git diff --exit-code
      ;;
    *)
      echo "unknown benchmark phase: $phase" >&2
      exit 2
      ;;
  esac
} > >(tee "$raw_log") 2>&1
status=$?
set -e

go list -mod=readonly -f '{{.ImportPath}} {{with .Module}}{{.Path}}@{{.Version}}{{end}}' \
  ./... | LC_ALL=C sort -u > "$evidence_dir/package-inventory.txt"
sed -E \
  -e 's#(^|[[:space:]])/[^[:space:]]+/(_work|go-build|go/pkg/mod)/#\1<path>/#g' \
  -e 's/[[:space:]][0-9]+(\.[0-9]+)?s$/ <duration>/' \
  "$raw_log" | LC_ALL=C sort > "$evidence_dir/normalized-output.txt"
git diff --binary --no-ext-diff | sed -E 's/index [0-9a-f]+\.\.[0-9a-f]+/index <digest>..<digest>/' \
  > "$evidence_dir/worktree-output.patch"
{
  sha256sum "$evidence_dir/package-inventory.txt"
  sha256sum "$evidence_dir/normalized-output.txt"
  sha256sum "$evidence_dir/worktree-output.patch"
} | sha256sum | awk '{print "sha256:" $1}' > "$evidence_dir/output-digest.txt"
exit "$status"
