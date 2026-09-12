#!/usr/bin/env bash

set -euo pipefail

[ "$#" -eq 1 ] || {
  echo "usage: $0 <fixed-workload-id>" >&2
  exit 2
}

case $1 in
go-build)
  exec scripts/go-retry.sh 3 go build ./...
  ;;
go-vet)
  scripts/go-retry.sh 3 go vet ./...
  exec scripts/vet-build-ignored-tools.sh
  ;;
go-race)
  exec scripts/go-retry.sh 2 go test -race ./internal/... ./tools/...
  ;;
provider-generation)
  exec go run tools/generate-all-schemas.go --spec-dir=docs/specifications/api
  ;;
documentation-generation)
  exec scripts/generate-provider-docs.sh --generate-docs-only
  ;;
terraform-example-validation)
  exec scripts/generate-provider-docs.sh --validate-examples-only
  ;;
release-source-reproduction)
  exec scripts/reproduce-release-source.sh --check-clean
  ;;
*)
  echo "unknown fixed workload: $1" >&2
  exit 2
  ;;
esac
