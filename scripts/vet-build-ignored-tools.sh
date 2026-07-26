#!/usr/bin/env bash
# Copyright (c) 2026 Robin Mordasiewicz. MIT License.
#
# Type-check the generator entry points under tools/ that carry `//go:build
# ignore`.
#
# Those files are hand-written code that consumes generated symbols, so they rot
# the same way internal/acctest did in #1351 — but the build tag hides them from
# `go build ./...`, `go vet ./...` and `go test ./...` alike, so nothing type-
# checks them until on-merge actually runs one and the regeneration cascade dies
# mid-flight.
#
# They cannot be vetted as a package: they are all `package main`, so a single
# `go vet ./tools/...` (even with -tags ignore) collapses them into one package
# full of redeclared symbols. Naming a file explicitly overrides its build tag,
# so each is vetted on its own.

set -euo pipefail

cd "$(dirname "$0")/.."

failed=0
checked=0

for file in tools/*.go; do
  [ -e "$file" ] || continue
  # Only the build-ignored entry points; anything without the tag is already
  # covered by the ordinary `go vet ./...`.
  if ! grep -q '^//go:build ignore' "$file"; then
    continue
  fi
  checked=$((checked + 1))
  echo "[vet] $file"
  if ! go vet "$file"; then
    failed=1
  fi
done

if [ "$checked" -eq 0 ]; then
  echo "ERROR: no build-ignored tools found under tools/ — this check must never silently pass over an empty set" >&2
  exit 1
fi

if [ "$failed" -ne 0 ]; then
  echo "ERROR: $checked build-ignored generator(s) checked, at least one failed go vet" >&2
  exit 1
fi

echo "OK: $checked build-ignored generator(s) type-check cleanly"
