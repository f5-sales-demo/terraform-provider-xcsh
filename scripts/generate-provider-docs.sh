#!/usr/bin/env bash

set -euo pipefail

repository_root=$(git rev-parse --show-toplevel)
cd "$repository_root"

fail() {
  echo "documentation generation failed: $*" >&2
  exit 1
}

validate_examples_only=false
case ${1:-} in
"") ;;
--validate-examples-only) validate_examples_only=true ;;
*) fail "unknown argument: $1" ;;
esac

required_commands=(go jq terraform)
if [ "$validate_examples_only" = false ]; then
  required_commands+=(tfplugindocs npx)
fi
for command in "${required_commands[@]}"; do
  command -v "$command" >/dev/null 2>&1 || fail "required command is unavailable: $command"
done
if [ "$validate_examples_only" = false ]; then
  [ -f docs/specifications/api/index.json ] ||
    fail "verified API specifications are missing; download the pinned release first"
  [ -d docs/specifications/api/domains ] ||
    fail "verified API specification domains are missing"
  [ -n "$(find docs/specifications/api/domains -type f -name '*.json' -print -quit)" ] ||
    fail "verified API specification domains are empty"
fi

terraform_version=$(tr -d '[:space:]' <.terraform-version)
[[ "$terraform_version" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]] ||
  fail ".terraform-version must contain an exact semantic version"
installed_terraform_version=$(terraform version -json | jq -er '.terraform_version')
[ "$installed_terraform_version" = "$terraform_version" ] ||
  fail "Terraform ${terraform_version} is required, found ${installed_terraform_version}"

temporary_root=$(mktemp -d)
trap 'rm -rf "$temporary_root"' EXIT

if [ "$validate_examples_only" = false ]; then
  echo "::group::Calculate minimum Terraform version"
  go run tools/calculate-terraform-version.go --update-templates
  echo "::endgroup::"

  echo "::group::Generate Terraform examples"
  go run tools/generate-examples.go
  go run tools/generate-test-examples.go
  echo "::endgroup::"
fi

# Build the checked-out provider once and expose it through a local filesystem
# mirror. Every validation case therefore uses these exact source bytes without
# consulting the Terraform Registry or whichever provider version is newest.
provider_version=99.0.0
provider_os=$(go env GOOS)
provider_arch=$(go env GOARCH)
provider_mirror="$temporary_root/provider-mirror"
provider_package_dir="${provider_mirror}/registry.terraform.io/f5-sales-demo/xcsh/${provider_version}/${provider_os}_${provider_arch}"
mkdir -p "$provider_package_dir"
go build -trimpath \
  -o "${provider_package_dir}/terraform-provider-xcsh_v${provider_version}" .

terraform_cli_config="$temporary_root/terraform.rc"
printf '%s\n' \
  'provider_installation {' \
  '  filesystem_mirror {' \
  "    path    = \"${provider_mirror}\"" \
  '    include = ["registry.terraform.io/f5-sales-demo/xcsh"]' \
  '  }' \
  '  direct {' \
  '    exclude = ["registry.terraform.io/f5-sales-demo/xcsh"]' \
  '  }' \
  '}' >"$terraform_cli_config"

plugin_cache_dir="$temporary_root/plugin-cache"
mkdir -p "$plugin_cache_dir"

validate_example() {
  local source_file=$1
  local case_dir case_name init_log selected_version validate_log version_log
  case_dir=$(mktemp -d "${temporary_root}/example.XXXXXX")
  case_name=${source_file#examples/}
  init_log="$case_dir/init.log"
  validate_log="$case_dir/validate.log"
  version_log="$case_dir/version.json"
  install -m 0600 "$source_file" "$case_dir/main.tf"
  if ! TF_CLI_CONFIG_FILE="$terraform_cli_config" TF_PLUGIN_CACHE_DIR="$plugin_cache_dir" \
    terraform -chdir="$case_dir" init -backend=false -input=false -no-color \
    >"$init_log" 2>&1; then
    cat "$init_log" >&2
    fail "terraform init rejected generated example ${case_name}"
  fi
  if ! TF_CLI_CONFIG_FILE="$terraform_cli_config" TF_PLUGIN_CACHE_DIR="$plugin_cache_dir" \
    terraform -chdir="$case_dir" version -json >"$version_log" 2>&1; then
    cat "$version_log" >&2
    fail "terraform could not report the provider selected for generated example ${case_name}"
  fi
  selected_version=$(jq -er \
    '.provider_selections["registry.terraform.io/f5-sales-demo/xcsh"]' \
    "$version_log") ||
    fail "generated example ${case_name} did not select f5-sales-demo/xcsh"
  [ "$selected_version" = "$provider_version" ] ||
    fail "generated example ${case_name} selected xcsh ${selected_version}, want local ${provider_version}"
  if ! TF_CLI_CONFIG_FILE="$terraform_cli_config" TF_PLUGIN_CACHE_DIR="$plugin_cache_dir" \
    terraform -chdir="$case_dir" validate -no-color >"$validate_log" 2>&1; then
    cat "$validate_log" >&2
    fail "terraform validate rejected generated example ${case_name}"
  fi
}

echo "::group::Validate every canonical generated example"
# The expected total is derived from the generated tree, not written down here.
# Every generated example directory owes exactly one canonical file, so the
# directory count IS the expectation. A literal total is only correct until the
# next spec release changes how many types the provider ships — this check was
# pinned at 279 from enriched spec v2.1.207 and went stale the moment the
# provider moved to 277, reporting a spec bump as a generation failure.
expected_canonical_examples=0
while IFS= read -r directory; do
  case $directory in
  examples/resources/*) canonical_file="$directory/resource.tf" ;;
  *) canonical_file="$directory/data-source.tf" ;;
  esac
  [ -f "$canonical_file" ] ||
    fail "generated example directory ${directory} has no canonical example file"
  expected_canonical_examples=$((expected_canonical_examples + 1))
done < <(find examples/resources examples/data-sources \
  -mindepth 1 -maxdepth 1 -type d -print | LC_ALL=C sort)

# Vacuity floor. The provider ships well over two hundred resources and data
# sources; a collapse to a handful means the generator or this walk broke, and a
# coverage check that has lost sight of the tree must fail rather than pass
# everything that remains.
minimum_canonical_examples=200
[ "$expected_canonical_examples" -ge "$minimum_canonical_examples" ] ||
  fail "found only ${expected_canonical_examples} generated example directories, want at least ${minimum_canonical_examples}; the generated tree is incomplete"

validated_canonical_examples=0
while IFS= read -r example; do
  validate_example "$example"
  validated_canonical_examples=$((validated_canonical_examples + 1))
done < <({
  find examples/resources -mindepth 2 -maxdepth 2 -type f -name resource.tf -print
  find examples/data-sources -mindepth 2 -maxdepth 2 -type f -name data-source.tf -print
} | LC_ALL=C sort)
[ "$validated_canonical_examples" -eq "$expected_canonical_examples" ] ||
  fail "validated ${validated_canonical_examples} canonical examples, want ${expected_canonical_examples} (one per generated example directory)"
echo "Validated ${validated_canonical_examples} canonical generated examples"
echo "::endgroup::"

echo "::group::Validate every named acceptance-derived example"
# Derived the same way as the canonical count above: the named examples are
# extracted from the acceptance suite, so their number changes whenever that
# suite does. Requiring each listed directory to contribute at least one example
# catches under-collection without a literal total that goes stale silently.
validated_named_examples=0
named_example_directories=(
  examples/resources/xcsh_http_loadbalancer
  examples/resources/xcsh_tcp_loadbalancer
  examples/resources/xcsh_healthcheck
  examples/resources/xcsh_app_firewall
  examples/resources/xcsh_origin_pool
  examples/resources/xcsh_rate_limiter
  examples/resources/xcsh_service_policy
  examples/resources/xcsh_user_identification
  examples/resources/xcsh_malicious_user_mitigation
)
for directory in "${named_example_directories[@]}"; do
  [ -d "$directory" ] ||
    fail "named example directory ${directory} does not exist"
  named_in_directory=$(find "$directory" -maxdepth 1 -type f -name '*.tf' \
    ! -name resource.tf -print | wc -l | tr -d '[:space:]')
  [ "$named_in_directory" -gt 0 ] ||
    fail "named example directory ${directory} contributed no examples"
done

while IFS= read -r example; do
  validate_example "$example"
  validated_named_examples=$((validated_named_examples + 1))
done < <(find "${named_example_directories[@]}" \
  -maxdepth 1 -type f -name '*.tf' ! -name resource.tf -print | LC_ALL=C sort)

# Vacuity floor: the acceptance suite yields several examples per listed
# resource. Dropping to a handful means extraction broke, and a check that has
# lost sight of its inputs must fail rather than pass the remainder.
minimum_named_examples=40
[ "$validated_named_examples" -ge "$minimum_named_examples" ] ||
  fail "validated only ${validated_named_examples} named examples, want at least ${minimum_named_examples}; acceptance-derived extraction is incomplete"
echo "Validated ${validated_named_examples} named generated examples"
echo "::endgroup::"

if [ "$validate_examples_only" = true ]; then
  exit 0
fi

echo "::group::Generate and transform provider documentation"
tfplugindocs generate --provider-name xcsh
go test tools/transform-docs.go tools/transform-docs_test.go
go run tools/transform-docs.go

# A transformer that converges only after a second run makes generated output
# depend on repository history. Snapshot the first pass, run it again, and fail
# on any byte-level drift.
first_pass_docs="$temporary_root/first-pass-docs"
mkdir -p "$first_pass_docs"
(tar -cf - docs) |
  (cd "$first_pass_docs" && tar -xf -)
go run tools/transform-docs.go
idempotence_diff="$temporary_root/transform-idempotence.diff"
if ! diff -qr "$first_pass_docs/docs" docs >"$idempotence_diff"; then
  sed -n '1,200p' "$idempotence_diff" >&2
  fail "documentation transformer changed its own first-pass output"
fi
echo "Verified documentation transformer idempotence"
echo "::endgroup::"

# Exporting the schema is an executable contract test for the exact provider
# binary. Use the same local mirror as the examples so init and schema export
# cannot silently fall back to a previously published provider.
schema_case="$temporary_root/schema"
mkdir -p "$schema_case"
printf '%s\n' \
  'terraform {' \
  '  required_providers {' \
  '    xcsh = {' \
  '      source  = "f5-sales-demo/xcsh"' \
  "      version = \"= ${provider_version}\"" \
  '    }' \
  '  }' \
  '}' >"$schema_case/main.tf"
echo "::group::Export provider schema"
TF_CLI_CONFIG_FILE="$terraform_cli_config" TF_PLUGIN_CACHE_DIR="$plugin_cache_dir" \
  terraform -chdir="$schema_case" init -backend=false -input=false -no-color
schema_output="$temporary_root/terraform-schema.json"
TF_CLI_CONFIG_FILE="$terraform_cli_config" TF_PLUGIN_CACHE_DIR="$plugin_cache_dir" \
  terraform -chdir="$schema_case" providers schema -json >"$schema_output"
jq -e '
  type == "object" and
  (.format_version | type == "string") and
  (.provider_schemas["registry.terraform.io/f5-sales-demo/xcsh"] | type == "object")
' "$schema_output" >/dev/null || fail "Terraform returned an incomplete provider schema"
echo "Exported provider schema: $(wc -c <"$schema_output" | tr -d '[:space:]') bytes"
echo "::endgroup::"

echo "::group::Generate machine-readable documentation indexes"
go run tools/generate-llms-txt.go
npx --yes @biomejs/biome@2.5.6 format --write docs/terraform-llms-index.json
echo "::endgroup::"
