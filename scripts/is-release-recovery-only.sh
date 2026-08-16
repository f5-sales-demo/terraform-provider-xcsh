#!/bin/bash

set -euo pipefail

saw_path=false
while IFS= read -r path; do
  [ -n "$path" ] || continue
  saw_path=true
  case "$path" in
  .github/workflows/* | \
    tools/validate_workflows_test.go | \
    internal/acctest/release_integrity_test.go | \
    scripts/generate-provider-docs.sh | \
    scripts/check-spec-version-freshness.sh | \
    scripts/test-check-spec-version-freshness.sh | \
    scripts/is-release-recovery-only.sh) ;;
  *) exit 1 ;;
  esac
done

[ "$saw_path" = true ]
