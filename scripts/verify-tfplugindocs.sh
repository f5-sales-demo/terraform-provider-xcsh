#!/usr/bin/env bash
set -euo pipefail

binary=$(command -v tfplugindocs)
metadata=$(go version -m "$binary") || {
  printf "%s\n" "$metadata" >&2
  exit 1
}
printf '%s\n' "$metadata" >&2
version=$(awk '$1 == "mod" && $2 == "github.com/hashicorp/terraform-plugin-docs" {print $3}' <<<"$metadata")
case "$version" in
v0.25.0) ;;
v0.25.0+dirty)
  # The checksum-pinned upstream release ZIP contains this exact modified-build
  # marker. Accept only its measured binary, never an arbitrary dirty build.
  checksum=$(sha256sum "$binary" | awk '{print $1}')
  test "$checksum" = 7e3e9e12d913576b3c1cb4e3ac15e68b15648452dfa2a28fc11fc21bdc2fb3eb || {
    echo 'tfplugindocs modified build differs from the official release binary' >&2
    exit 1
  }
  ;;
*)
  echo 'tfplugindocs module version is not the verified 0.25.0 release' >&2
  exit 1
  ;;
esac
printf 'v0.25.0\n'
