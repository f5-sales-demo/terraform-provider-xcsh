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
    scripts/prepare-spec-delivery-receipt.sh | \
    scripts/validate-provider-delivery-state.sh | \
    scripts/is-pending-delivery-active.sh | \
    scripts/is-release-recovery-only.sh | \
    tools/provider-publication-receipts.json | \
    tools/spec-deliveries.json | \
    tools/spec-pending-delivery.json | \
    tools/spec-regeneration-receipt.json) ;;
  *) exit 1 ;;
  esac
done

[ "$saw_path" = true ]
