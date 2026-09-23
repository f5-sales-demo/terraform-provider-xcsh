#!/usr/bin/env bash

set -euo pipefail

# Schema export compiles the unusually large generated provider. Bound both
# package and compiler parallelism so constrained CI runners cannot trade
# determinism for a kernel OOM kill; callers may override these limits when a
# larger dedicated builder is available.
export GOFLAGS="${GOFLAGS:--p=1}"
export GOMAXPROCS="${GOMAXPROCS:-2}"
export GOMEMLIMIT="${GOMEMLIMIT:-1200MiB}"

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

echo "::group::Validate exact release-surface examples"
expected_example_dirs="$temporary_root/expected-example-dirs.txt"
actual_example_dirs="$temporary_root/actual-example-dirs.txt"
jq -r '
  (.resources[] | "examples/resources/xcsh_" + .),
  (.data_sources[] | "examples/data-sources/xcsh_" + .),
  (.actions[] | "examples/actions/xcsh_" + .)
' smsv2-release-surface.json | LC_ALL=C sort >"$expected_example_dirs"
find examples/resources examples/data-sources examples/actions \
  -mindepth 1 -maxdepth 1 -type d -print | LC_ALL=C sort >"$actual_example_dirs"
if ! diff -u "$expected_example_dirs" "$actual_example_dirs"; then
  fail "example directories do not exactly match smsv2-release-surface.json"
fi

validated_canonical_examples=0
while IFS= read -r directory; do
  case $directory in
  examples/resources/*) canonical_file="$directory/resource.tf" ;;
  examples/data-sources/*) canonical_file="$directory/data-source.tf" ;;
  examples/actions/*) canonical_file="$directory/action.tf" ;;
  *) fail "unexpected example directory ${directory}" ;;
  esac
  [ -f "$canonical_file" ] || fail "missing canonical example ${canonical_file}"
  validate_example "$canonical_file"
  validated_canonical_examples=$((validated_canonical_examples + 1))
done <"$expected_example_dirs"
expected_canonical_examples=$(wc -l <"$expected_example_dirs" | tr -d '[:space:]')
[ "$validated_canonical_examples" -eq "$expected_canonical_examples" ] ||
  fail "validated ${validated_canonical_examples} canonical examples, want ${expected_canonical_examples}"
echo "Validated ${validated_canonical_examples} exact release-surface examples"
echo "::endgroup::"

if [ "$validate_examples_only" = true ]; then
  exit 0
fi

echo "::group::Generate and transform provider documentation"
tfplugindocs generate --provider-name xcsh
go run tools/transform-docs.go
# The exhaustive documentation assertions inspect the transformed files. Run
# them after the first transformation pass; running them against tfplugindocs'
# raw nested-schema output rejects a valid fresh generation before the
# transformer has had a chance to add stable anchors and flattened paths.
go test tools/transform-docs.go tools/transform-docs_test.go

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

expected_doc_files="$temporary_root/expected-doc-files.txt"
actual_doc_files="$temporary_root/actual-doc-files.txt"
jq -r '
  (.resources[] | "docs/resources/" + . + ".md"),
  (.data_sources[] | "docs/data-sources/" + . + ".md"),
  (.actions[] | "docs/actions/" + . + ".md")
' smsv2-release-surface.json | LC_ALL=C sort >"$expected_doc_files"
find docs/resources docs/data-sources docs/actions \
  -maxdepth 1 -type f -name '*.md' ! -name index.md -print | LC_ALL=C sort >"$actual_doc_files"
if ! diff -u "$expected_doc_files" "$actual_doc_files"; then
  fail "generated documentation does not exactly match smsv2-release-surface.json"
fi
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
schema_resources="$temporary_root/schema-resources.txt"
schema_data_sources="$temporary_root/schema-data-sources.txt"
schema_actions="$temporary_root/schema-actions.txt"
expected_resources="$temporary_root/expected-resources.txt"
expected_data_sources="$temporary_root/expected-data-sources.txt"
expected_actions="$temporary_root/expected-actions.txt"
jq -r '.resources[] | "xcsh_" + .' smsv2-release-surface.json | LC_ALL=C sort >"$expected_resources"
jq -r '.data_sources[] | "xcsh_" + .' smsv2-release-surface.json | LC_ALL=C sort >"$expected_data_sources"
jq -r '.actions[] | "xcsh_" + .' smsv2-release-surface.json | LC_ALL=C sort >"$expected_actions"
jq -r '.provider_schemas["registry.terraform.io/f5-sales-demo/xcsh"].resource_schemas | keys[]' "$schema_output" | LC_ALL=C sort >"$schema_resources"
jq -r '.provider_schemas["registry.terraform.io/f5-sales-demo/xcsh"].data_source_schemas | keys[]' "$schema_output" | LC_ALL=C sort >"$schema_data_sources"
jq -r '.provider_schemas["registry.terraform.io/f5-sales-demo/xcsh"].action_schemas | keys[]' "$schema_output" | LC_ALL=C sort >"$schema_actions"
diff -u "$expected_resources" "$schema_resources" || fail "installed resource schema does not match the release surface"
diff -u "$expected_data_sources" "$schema_data_sources" || fail "installed data-source schema does not match the release surface"
diff -u "$expected_actions" "$schema_actions" || fail "installed action schema does not match the release surface"
jq -e '
  .provider_schemas["registry.terraform.io/f5-sales-demo/xcsh"] as $provider |
  (($provider.functions // {}) | length) == 0
' "$schema_output" >/dev/null || fail "installed provider unexpectedly exposes functions"
echo "Exported provider schema: $(wc -c <"$schema_output" | tr -d '[:space:]') bytes"
echo "::endgroup::"

echo "::group::Generate machine-readable documentation indexes"
go run tools/generate-llms-txt.go
npx --yes @biomejs/biome@2.5.6 format --write docs/terraform-llms-index.json
echo "::endgroup::"
