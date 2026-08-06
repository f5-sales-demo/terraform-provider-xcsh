#!/bin/bash
# Check if tools/spec-version.txt matches the latest release tag from f5-sales-demo/api-specs-enriched
# Issue #1255: Spec version freshness guard

set -e

SPEC_VERSION_FILE="${1:-tools/spec-version.txt}"

if [ ! -f "$SPEC_VERSION_FILE" ]; then
  echo "::error::Spec version file not found at $SPEC_VERSION_FILE"
  exit 1
fi

CURRENT_VERSION=$(tr -d '[:space:]' < "$SPEC_VERSION_FILE")
echo "Current pinned spec version: $CURRENT_VERSION"

if command -v gh >/dev/null 2>&1; then
  LATEST_RELEASE=$(gh api repos/f5-sales-demo/api-specs-enriched/releases --jq '.[0].tag_name' 2>/dev/null || echo "")
  if [ -n "$LATEST_RELEASE" ]; then
    echo "Latest upstream api-specs release: $LATEST_RELEASE"
    if [ "$CURRENT_VERSION" != "$LATEST_RELEASE" ]; then
      echo "::warning::tools/spec-version.txt ($CURRENT_VERSION) lags latest upstream api-specs-enriched release ($LATEST_RELEASE)"
    else
      echo "✅ Spec version is up to date ($CURRENT_VERSION)"
    fi
  else
    echo "Note: Could not query upstream releases from GitHub API, skipping release comparison."
  fi
else
  echo "Note: gh CLI not available, skipping upstream release comparison."
fi

exit 0
