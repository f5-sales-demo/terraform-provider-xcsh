#!/bin/bash
# Check if tools/spec-version.txt matches the latest release tag from f5-sales-demo/api-specs-enriched
# Issue #1255: Spec version freshness guard

set -e

SPEC_VERSION_FILE="${1:-tools/spec-version.txt}"
PENDING_DELIVERY_FILE="${2:-tools/spec-pending-delivery.json}"

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
    if [ "${ALLOW_PENDING_DELIVERY:-false}" = "true" ]; then
      [ -f "$PENDING_DELIVERY_FILE" ] || {
        echo "::error::Pending-delivery recovery was requested without a pending delivery"
        exit 1
      }
      jq -e --arg tag "$CURRENT_VERSION" '
        type == "object" and
        (keys | sort) == ["delivery_id", "release_tag", "target_commit", "version"] and
        (.delivery_id | test("^[0-9a-f]{64}$")) and
        .release_tag == $tag and
        .release_tag == ("v" + .version) and
        (.target_commit | test("^[0-9a-f]{40}$"))
      ' "$PENDING_DELIVERY_FILE" >/dev/null || {
        echo "::error::Pending delivery cannot authorize stale-pin recovery"
        exit 1
      }
      CANONICAL=$(jq -cnS \
        --arg commit "$(jq -r '.target_commit' "$PENDING_DELIVERY_FILE")" \
        --arg event_type enriched-specs-updated \
        --arg source f5-sales-demo/api-specs-enriched \
        --arg tag "$(jq -r '.release_tag' "$PENDING_DELIVERY_FILE")" \
        --arg target f5-sales-demo/terraform-provider-xcsh \
        --arg version "$(jq -r '.version' "$PENDING_DELIVERY_FILE")" \
        '{commit:$commit,event_type:$event_type,source:$source,tag:$tag,target:$target,version:$version}')
      EXPECTED_ID=$(printf '%s' "$CANONICAL" | shasum -a 256 | awk '{print $1}')
      [ "$(jq -r '.delivery_id' "$PENDING_DELIVERY_FILE")" = "$EXPECTED_ID" ] || {
        echo "::error::Pending delivery ID is not canonical"
        exit 1
      }
      echo "::notice::Allowing workflow-only recovery for exact pending delivery ${EXPECTED_ID}"
      exit 0
    fi
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
