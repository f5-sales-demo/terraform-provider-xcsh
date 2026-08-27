#!/usr/bin/env bash

set -euo pipefail

fail() {
  printf 'error: %s\n' "$*" >&2
  exit 1
}

[ "$#" -eq 2 ] || fail "usage: $0 CURRENT_VERSION INCOMING_VERSION"
: "${GITHUB_OUTPUT:?GITHUB_OUTPUT must name the workflow output file}"

current_major=0
current_minor=0
current_patch=0
incoming_major=0
incoming_minor=0
incoming_patch=0

parse_version() {
  local prefix=$1
  local version=${2#v}

  [[ "$version" =~ ^([0-9]+)\.([0-9]+)\.([0-9]+)$ ]] \
    || fail "version must be vMAJOR.MINOR.PATCH: $2"
  printf -v "${prefix}_major" '%d' "$((10#${BASH_REMATCH[1]}))"
  printf -v "${prefix}_minor" '%d' "$((10#${BASH_REMATCH[2]}))"
  printf -v "${prefix}_patch" '%d' "$((10#${BASH_REMATCH[3]}))"
}

parse_version current "$1"
parse_version incoming "$2"

if ((incoming_major < current_major)) \
  || ((incoming_major == current_major && incoming_minor < current_minor)) \
  || ((incoming_major == current_major && incoming_minor == current_minor && incoming_patch <= current_patch)); then
  fail "incoming version $2 must be newer than current version $1"
fi

if ((incoming_major > current_major)); then
  breaking=true
  pr_title='feat!: update F5 Distributed Cloud OpenAPI specifications'
  breaking_footer="BREAKING CHANGE: updates the provider to API $2 and removes superseded contract fields."
else
  breaking=false
  pr_title='feat: update F5 Distributed Cloud OpenAPI specifications'
  breaking_footer=''
fi

{
  printf 'breaking=%s\n' "$breaking"
  printf 'pr_title=%s\n' "$pr_title"
  printf 'breaking_footer=%s\n' "$breaking_footer"
} >>"$GITHUB_OUTPUT"
