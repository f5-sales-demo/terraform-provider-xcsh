#!/bin/bash
# Check if tools/spec-version.txt matches the latest release tag from f5-sales-demo/api-specs-enriched
# Issue #1255: Spec version freshness guard

set -e

SPEC_VERSION_FILE="${1:-tools/spec-version.txt}"

if [ ! -f "$SPEC_VERSION_FILE" ]; then
  echo "::error::Spec version file not found at $SPEC_VERSION_FILE"
  exit 1
fi

CURRENT_VERSION=$(cat "$SPEC_VERSION_FILE" | tr -d '[:space:]')
echo "Current pinned spec version: $CURRENT_VERSION"

if command -v gh >/dev/null 2>&1; then
  set +e
  # Retrieve the latest release tag. Exit 2 if gh api fails.
  LATEST_RELEASE=$(gh api repos/f5-sales-demo/api-specs-enriched/releases/latest --jq '.tag_name' 2>/dev/null)
  GH_EXIT_CODE=$?
  set -e

  if [ $GH_EXIT_CODE -ne 0 ] || [ -z "$LATEST_RELEASE" ] || [ "$LATEST_RELEASE" = "null" ]; then
    echo "::error::Could not query upstream releases from GitHub API or received invalid/empty tag. Please check network connectivity or API limits."
    exit 2
  fi

  echo "Latest upstream api-specs release: $LATEST_RELEASE"
  if [ "$CURRENT_VERSION" != "$LATEST_RELEASE" ]; then
    echo "::error::tools/spec-version.txt ($CURRENT_VERSION) lags latest upstream api-specs-enriched release ($LATEST_RELEASE)"
    exit 1
  else
    echo "✅ Spec version is up to date ($CURRENT_VERSION)"
    exit 0
  fi
else
  echo "::error::gh CLI not available, skipping upstream release comparison. This prevents CI from validating spec freshness."
  exit 2
fi
