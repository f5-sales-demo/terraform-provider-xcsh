#!/bin/bash
# Copyright (c) 2026 Robin Mordasiewicz. MIT License.

# Pre-commit hook to lint generated files when generators change
# This catches linting errors early without violating the constitution
#
# Workflow:
# 1. Detect if generator tools are being committed
# 2. Temporarily run generators to produce output
# 3. Lint the generated artifacts (Go code AND documentation)
# 4. Restore generated files to original state
# 5. Report linting errors and block commit if needed
#
# This does NOT commit generated files - it only previews what CI/CD will produce
# and validates that it will pass linting.

set -e

# ANSI color codes
RED='\033[0;31m'
YELLOW='\033[1;33m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
BOLD='\033[1m'
NC='\033[0m' # No Color

# Configuration
SPEC_DIR="docs/specifications/api"
TOOLS_DIR="tools"
PROVIDER_DIR="internal/provider"
CLIENT_DIR="internal/client"
DOCS_DIR="docs"
EXAMPLES_DIR="examples"

# Generator tools that produce Go code requiring linting
GO_GENERATORS=(
  "tools/generate-all-schemas.go"
)

# Generator tools that produce documentation/examples requiring linting
DOC_GENERATORS=(
  "tools/transform-docs.go"
  "tools/generate-examples.go"
)

# Track what needs to be linted
LINT_GO=false
LINT_DOCS=false

echo "🔍 Checking for generator tool changes..."

# Get list of staged files
STAGED_FILES=$(git diff --cached --name-only --diff-filter=ACM)

if [ -z "$STAGED_FILES" ]; then
  echo -e "${GREEN}✅ No files staged${NC}"
  exit 0
fi

# Check if any Go generators are being modified
MODIFIED_GO_GENERATORS=()
for generator in "${GO_GENERATORS[@]}"; do
  if echo "$STAGED_FILES" | grep -qE "^${generator}$"; then
    MODIFIED_GO_GENERATORS+=("$generator")
    LINT_GO=true
  fi
done

# Check if any doc generators are being modified
MODIFIED_DOC_GENERATORS=()
for generator in "${DOC_GENERATORS[@]}"; do
  if echo "$STAGED_FILES" | grep -qE "^${generator}$"; then
    MODIFIED_DOC_GENERATORS+=("$generator")
    LINT_DOCS=true
  fi
done

# Also check for spec changes that would trigger regeneration
SPEC_CHANGES=false
if echo "$STAGED_FILES" | grep -qE "^${SPEC_DIR}/"; then
  SPEC_CHANGES=true
  LINT_GO=true
  LINT_DOCS=true
fi

# Check for template changes that affect doc generation
TEMPLATE_CHANGES=false
if echo "$STAGED_FILES" | grep -qE "^templates/"; then
  TEMPLATE_CHANGES=true
  LINT_DOCS=true
fi

# If no generators, specs, or templates modified, nothing to preview
if [ "$LINT_GO" = false ] && [ "$LINT_DOCS" = false ]; then
  echo -e "${GREEN}✅ No generator tools, specs, or templates modified - skipping preview lint${NC}"
  exit 0
fi

# Report what triggered the preview
echo ""
echo -e "${CYAN}╭────────────────────────────────────────────────────────────────────────────╮${NC}"
echo -e "${CYAN}│              GENERATOR CHANGES DETECTED - PREVIEW LINTING                 │${NC}"
echo -e "${CYAN}╰────────────────────────────────────────────────────────────────────────────╯${NC}"
echo ""

if [ ${#MODIFIED_GO_GENERATORS[@]} -gt 0 ]; then
  echo -e "${YELLOW}Modified Go generators:${NC}"
  for gen in "${MODIFIED_GO_GENERATORS[@]}"; do
    echo "  • $gen"
  done
fi

if [ ${#MODIFIED_DOC_GENERATORS[@]} -gt 0 ]; then
  echo -e "${YELLOW}Modified documentation generators:${NC}"
  for gen in "${MODIFIED_DOC_GENERATORS[@]}"; do
    echo "  • $gen"
  done
fi

if [ "$SPEC_CHANGES" = true ]; then
  echo -e "${YELLOW}OpenAPI spec changes detected in:${NC} ${SPEC_DIR}/"
fi

if [ "$TEMPLATE_CHANGES" = true ]; then
  echo -e "${YELLOW}Template changes detected in:${NC} templates/"
fi
echo ""

# Check if specs exist (required for Go generation)
if [ "$LINT_GO" = true ]; then
  if [ ! -d "$SPEC_DIR" ] || ! ls "$SPEC_DIR"/docs-cloud-f5-com.*.ves-swagger.json 1>/dev/null 2>&1; then
    echo -e "${YELLOW}⚠️  No OpenAPI specs found in ${SPEC_DIR}${NC}"
    echo "   Cannot run preview generation without specs."
    echo "   Skipping Go preview lint - CI/CD will handle generation."
    LINT_GO=false
  fi
fi

# Check if golangci-lint is available
if [ "$LINT_GO" = true ]; then
  if ! command -v golangci-lint &>/dev/null; then
    echo -e "${YELLOW}⚠️  golangci-lint not installed${NC}"
    echo "   Install it: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
    echo "   Skipping Go preview lint - CI/CD will run linting."
    LINT_GO=false
  fi
fi

# Check if markdownlint-cli2 is available
if [ "$LINT_DOCS" = true ]; then
  if ! command -v markdownlint-cli2 &>/dev/null; then
    echo -e "${YELLOW}⚠️  markdownlint-cli2 not installed${NC}"
    echo "   Install it: npm install -g markdownlint-cli2"
    echo "   Skipping markdown preview lint - CI/CD will run linting."
    LINT_DOCS=false
  fi
fi

# Check if tfplugindocs is available
if [ "$LINT_DOCS" = true ]; then
  if ! command -v tfplugindocs &>/dev/null; then
    # Try in ~/go/bin
    if [ -x "$HOME/go/bin/tfplugindocs" ]; then
      TFPLUGINDOCS="$HOME/go/bin/tfplugindocs"
    else
      echo -e "${YELLOW}⚠️  tfplugindocs not installed${NC}"
      echo "   Install it: go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest"
      echo "   Skipping doc preview lint - CI/CD will run linting."
      LINT_DOCS=false
    fi
  else
    TFPLUGINDOCS="tfplugindocs"
  fi
fi

# If nothing to lint after checks, exit
if [ "$LINT_GO" = false ] && [ "$LINT_DOCS" = false ]; then
  echo -e "${GREEN}✅ Skipping preview lint (missing tools or specs)${NC}"
  exit 0
fi

echo -e "${BOLD}Running preview generation and lint...${NC}"
echo ""

# Save current state of generated files
TEMP_BACKUP_DIR=$(mktemp -d)
RESTORE_NEEDED=false

# Function to save generated files
save_generated_files() {
  echo "📦 Saving current state of generated files..."

  # Save Go resource files
  if ls ${PROVIDER_DIR}/*_resource.go 1>/dev/null 2>&1; then
    mkdir -p "$TEMP_BACKUP_DIR/provider"
    cp ${PROVIDER_DIR}/*_resource.go "$TEMP_BACKUP_DIR/provider/" 2>/dev/null || true
    RESTORE_NEEDED=true
  fi

  # Save Go data source files
  if ls ${PROVIDER_DIR}/*_data_source.go 1>/dev/null 2>&1; then
    mkdir -p "$TEMP_BACKUP_DIR/provider"
    cp ${PROVIDER_DIR}/*_data_source.go "$TEMP_BACKUP_DIR/provider/" 2>/dev/null || true
    RESTORE_NEEDED=true
  fi

  # Save client types (if they exist)
  if ls ${CLIENT_DIR}/*_types.go 1>/dev/null 2>&1; then
    mkdir -p "$TEMP_BACKUP_DIR/client"
    cp ${CLIENT_DIR}/*_types.go "$TEMP_BACKUP_DIR/client/" 2>/dev/null || true
    RESTORE_NEEDED=true
  fi

  # Save documentation files
  if [ -d "$DOCS_DIR/resources" ]; then
    mkdir -p "$TEMP_BACKUP_DIR/docs"
    cp -r "$DOCS_DIR/resources" "$TEMP_BACKUP_DIR/docs/" 2>/dev/null || true
    RESTORE_NEEDED=true
  fi
  if [ -d "$DOCS_DIR/data-sources" ]; then
    mkdir -p "$TEMP_BACKUP_DIR/docs"
    cp -r "$DOCS_DIR/data-sources" "$TEMP_BACKUP_DIR/docs/" 2>/dev/null || true
    RESTORE_NEEDED=true
  fi
  if [ -f "$DOCS_DIR/index.md" ]; then
    mkdir -p "$TEMP_BACKUP_DIR/docs"
    cp "$DOCS_DIR/index.md" "$TEMP_BACKUP_DIR/docs/" 2>/dev/null || true
    RESTORE_NEEDED=true
  fi

  # Save example files (only generated ones)
  if [ -d "$EXAMPLES_DIR/resources" ]; then
    mkdir -p "$TEMP_BACKUP_DIR/examples"
    cp -r "$EXAMPLES_DIR/resources" "$TEMP_BACKUP_DIR/examples/" 2>/dev/null || true
    RESTORE_NEEDED=true
  fi
  if [ -d "$EXAMPLES_DIR/data-sources" ]; then
    mkdir -p "$TEMP_BACKUP_DIR/examples"
    cp -r "$EXAMPLES_DIR/data-sources" "$TEMP_BACKUP_DIR/examples/" 2>/dev/null || true
    RESTORE_NEEDED=true
  fi
}

# Function to restore generated files (invoked via trap on EXIT)
# shellcheck disable=SC2329
restore_generated_files() {
  if [ "$RESTORE_NEEDED" = true ]; then
    echo ""
    echo "📦 Restoring original generated files..."

    # Remove newly generated Go files first
    rm -f "${PROVIDER_DIR}"/*_resource.go 2>/dev/null || true
    rm -f "${PROVIDER_DIR}"/*_data_source.go 2>/dev/null || true
    rm -f "${CLIENT_DIR}"/*_types.go 2>/dev/null || true

    # Remove newly generated docs
    rm -rf "${DOCS_DIR}/resources" 2>/dev/null || true
    rm -rf "${DOCS_DIR}/data-sources" 2>/dev/null || true
    rm -f "${DOCS_DIR}/index.md" 2>/dev/null || true

    # Remove newly generated examples
    rm -rf "${EXAMPLES_DIR}/resources" 2>/dev/null || true
    rm -rf "${EXAMPLES_DIR}/data-sources" 2>/dev/null || true

    # Restore Go files from backup
    if [ -d "$TEMP_BACKUP_DIR/provider" ]; then
      cp "$TEMP_BACKUP_DIR/provider/"*.go "$PROVIDER_DIR/" 2>/dev/null || true
    fi
    if [ -d "$TEMP_BACKUP_DIR/client" ]; then
      cp "$TEMP_BACKUP_DIR/client/"*.go "$CLIENT_DIR/" 2>/dev/null || true
    fi

    # Restore docs from backup
    if [ -d "$TEMP_BACKUP_DIR/docs/resources" ]; then
      cp -r "$TEMP_BACKUP_DIR/docs/resources" "$DOCS_DIR/" 2>/dev/null || true
    fi
    if [ -d "$TEMP_BACKUP_DIR/docs/data-sources" ]; then
      cp -r "$TEMP_BACKUP_DIR/docs/data-sources" "$DOCS_DIR/" 2>/dev/null || true
    fi
    if [ -f "$TEMP_BACKUP_DIR/docs/index.md" ]; then
      cp "$TEMP_BACKUP_DIR/docs/index.md" "$DOCS_DIR/" 2>/dev/null || true
    fi

    # Restore examples from backup
    if [ -d "$TEMP_BACKUP_DIR/examples/resources" ]; then
      cp -r "$TEMP_BACKUP_DIR/examples/resources" "$EXAMPLES_DIR/" 2>/dev/null || true
    fi
    if [ -d "$TEMP_BACKUP_DIR/examples/data-sources" ]; then
      cp -r "$TEMP_BACKUP_DIR/examples/data-sources" "$EXAMPLES_DIR/" 2>/dev/null || true
    fi

    echo -e "${GREEN}✅ Generated files restored to original state${NC}"
  fi

  # Clean up temp directory
  rm -rf "$TEMP_BACKUP_DIR"
}

# Set up trap to ensure cleanup on exit
trap restore_generated_files EXIT

# Save current state
save_generated_files

# Track overall success
GO_LINT_SUCCESS=true
DOC_LINT_SUCCESS=true
GO_LINT_OUTPUT=""
DOC_LINT_OUTPUT=""

# ==============================================================================
# PHASE 1: Go Code Generation and Linting
# ==============================================================================
if [ "$LINT_GO" = true ]; then
  echo "🔧 Running Go generator: generate-all-schemas.go"
  echo "   (This is a preview only - generated files will be restored)"
  echo ""

  GENERATION_OUTPUT=""

  if ! GENERATION_OUTPUT=$(go run "${TOOLS_DIR}/generate-all-schemas.go" --spec-dir="$SPEC_DIR" 2>&1); then
    echo -e "${RED}╔══════════════════════════════════════════════════════════════════════════════╗${NC}"
    echo -e "${RED}║                     ❌ GO GENERATION FAILED                                   ║${NC}"
    echo -e "${RED}╚══════════════════════════════════════════════════════════════════════════════╝${NC}"
    echo ""
    echo -e "${YELLOW}Generator output:${NC}"
    echo "$GENERATION_OUTPUT"
    echo ""
    echo -e "${CYAN}The Go generator failed to produce output. Fix the generator before committing.${NC}"
    echo -e "${CYAN}File: ${BOLD}${TOOLS_DIR}/generate-all-schemas.go${NC}"
    exit 1
  fi

  echo "   ✅ Go generation completed"
  echo ""

  # Run linting on generated Go files
  echo "🔍 Running golangci-lint on generated Go files..."
  echo ""

  # Lint same paths as CI: ./internal/... .
  if ! GO_LINT_OUTPUT=$(golangci-lint run --timeout=5m ./internal/... . 2>&1); then
    GO_LINT_SUCCESS=false
  fi

  if [ "$GO_LINT_SUCCESS" = true ]; then
    echo -e "   ${GREEN}✅ Go lint passed${NC}"
  else
    echo -e "   ${RED}❌ Go lint failed${NC}"
  fi
  echo ""
fi

# ==============================================================================
# PHASE 2: Documentation Generation and Linting
# ==============================================================================
if [ "$LINT_DOCS" = true ]; then
  echo "🔧 Running documentation generators..."
  echo "   (This is a preview only - generated files will be restored)"
  echo ""

  # Step 1: Generate examples first (needed by tfplugindocs)
  echo "   📝 Generating examples..."
  if ! EXAMPLES_OUTPUT=$(go run "${TOOLS_DIR}/generate-examples.go" 2>&1); then
    echo -e "${RED}╔══════════════════════════════════════════════════════════════════════════════╗${NC}"
    echo -e "${RED}║                  ❌ EXAMPLES GENERATION FAILED                                ║${NC}"
    echo -e "${RED}╚══════════════════════════════════════════════════════════════════════════════╝${NC}"
    echo ""
    echo -e "${YELLOW}Generator output:${NC}"
    echo "$EXAMPLES_OUTPUT"
    echo ""
    echo -e "${CYAN}The examples generator failed. Fix before committing.${NC}"
    echo -e "${CYAN}File: ${BOLD}${TOOLS_DIR}/generate-examples.go${NC}"
    exit 1
  fi
  echo "      ✅ Examples generated"

  # Step 2: Run tfplugindocs
  echo "   📝 Running tfplugindocs..."
  if ! TFPLUGINDOCS_OUTPUT=$($TFPLUGINDOCS generate 2>&1); then
    echo -e "${RED}╔══════════════════════════════════════════════════════════════════════════════╗${NC}"
    echo -e "${RED}║                  ❌ TFPLUGINDOCS GENERATION FAILED                            ║${NC}"
    echo -e "${RED}╚══════════════════════════════════════════════════════════════════════════════╝${NC}"
    echo ""
    echo -e "${YELLOW}Generator output:${NC}"
    echo "$TFPLUGINDOCS_OUTPUT"
    echo ""
    echo -e "${CYAN}tfplugindocs failed. This may be due to schema issues in the provider.${NC}"
    exit 1
  fi
  echo "      ✅ tfplugindocs completed"

  # Step 3: Run transform-docs.go
  echo "   📝 Running transform-docs.go..."
  if ! TRANSFORM_OUTPUT=$(go run "${TOOLS_DIR}/transform-docs.go" 2>&1); then
    echo -e "${RED}╔══════════════════════════════════════════════════════════════════════════════╗${NC}"
    echo -e "${RED}║                  ❌ DOC TRANSFORMATION FAILED                                 ║${NC}"
    echo -e "${RED}╚══════════════════════════════════════════════════════════════════════════════╝${NC}"
    echo ""
    echo -e "${YELLOW}Generator output:${NC}"
    echo "$TRANSFORM_OUTPUT"
    echo ""
    echo -e "${CYAN}The doc transformer failed. Fix before committing.${NC}"
    echo -e "${CYAN}File: ${BOLD}${TOOLS_DIR}/transform-docs.go${NC}"
    exit 1
  fi
  echo "      ✅ Doc transformation completed"
  echo ""

  # Run markdown linting on generated docs
  echo "🔍 Running markdownlint on generated documentation..."
  echo ""

  # Lint docs directory (excluding functions which are manually maintained)
  if ! DOC_LINT_OUTPUT=$(markdownlint-cli2 "docs/resources/*.md" "docs/data-sources/*.md" "docs/index.md" 2>&1); then
    DOC_LINT_SUCCESS=false
  fi

  if [ "$DOC_LINT_SUCCESS" = true ]; then
    echo -e "   ${GREEN}✅ Markdown lint passed${NC}"
  else
    echo -e "   ${RED}❌ Markdown lint failed${NC}"
  fi
  echo ""
fi

# ==============================================================================
# PHASE 3: Report Results
# ==============================================================================

# Check overall results
if [ "$GO_LINT_SUCCESS" = true ] && [ "$DOC_LINT_SUCCESS" = true ]; then
  echo -e "${GREEN}╔══════════════════════════════════════════════════════════════════════════════╗${NC}"
  echo -e "${GREEN}║                  ✅ ALL PREVIEW LINTS PASSED                                  ║${NC}"
  echo -e "${GREEN}╚══════════════════════════════════════════════════════════════════════════════╝${NC}"
  echo ""
  echo "The generated code and documentation will pass linting in CI/CD."
  echo "Generated files have been restored - only your source changes will be committed."
  exit 0
fi

# Report failures
echo -e "${RED}╔══════════════════════════════════════════════════════════════════════════════╗${NC}"
echo -e "${RED}║                  ❌ PREVIEW LINT FAILED                                        ║${NC}"
echo -e "${RED}╚══════════════════════════════════════════════════════════════════════════════╝${NC}"
echo ""

if [ "$GO_LINT_SUCCESS" = false ]; then
  echo -e "${YELLOW}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
  echo -e "${YELLOW}                         GO LINTING ERRORS                                    ${NC}"
  echo -e "${YELLOW}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
  echo ""
  echo "$GO_LINT_OUTPUT"
  echo ""
  echo -e "${CYAN}Fix the Go generator: ${BOLD}${TOOLS_DIR}/generate-all-schemas.go${NC}"
  echo ""
fi

if [ "$DOC_LINT_SUCCESS" = false ]; then
  echo -e "${YELLOW}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
  echo -e "${YELLOW}                      MARKDOWN LINTING ERRORS                                 ${NC}"
  echo -e "${YELLOW}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
  echo ""
  echo "$DOC_LINT_OUTPUT"
  echo ""
  echo -e "${CYAN}Fix the documentation generator: ${BOLD}${TOOLS_DIR}/transform-docs.go${NC}"
  echo ""
fi

echo -e "${CYAN}╭────────────────────────────────────────────────────────────────────────────╮${NC}"
echo -e "${CYAN}│                       HOW TO FIX THIS                                      │${NC}"
echo -e "${CYAN}╰────────────────────────────────────────────────────────────────────────────╯${NC}"
echo ""
echo "  The linting errors above will occur when CI/CD regenerates the files."
echo "  You must fix the generator(s) to produce lint-compliant output."
echo ""
echo "  ${BOLD}Steps:${NC}"
echo "    1. Review the linting errors above"
echo "    2. Fix the appropriate generator tool(s)"
echo "    3. Run this check again: ${BOLD}pre-commit run lint-generated-preview${NC}"
echo ""
echo "  ${BOLD}Note:${NC} Generated files have been restored to their original state."
echo "        Only your generator changes are staged for commit."
echo ""
exit 1
