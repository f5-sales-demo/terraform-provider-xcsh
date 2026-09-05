# Makefile for terraform-provider-xcsh
# Automated build, test, and code generation

BINARY_NAME=terraform-provider-xcsh
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

# Directories
TOOLS_DIR=tools
PROVIDER_DIR=internal/provider
CLIENT_DIR=internal/client
DOCS_DIR=docs
SPEC_DIR?=docs/specifications/api

# API spec source
ENRICHED_REPO?=f5-sales-demo/api-specs-enriched

# Go commands
GO=go
GOFMT=gofmt
GOLINT=golangci-lint
GOLANGCI_LINT_VERSION?=$(shell tr -d '[:space:]' < .golangci-version)
TERRAFORM_VERSION?=$(shell tr -d '[:space:]' < .terraform-version)

.PHONY: all build test lint fmt clean clean-generated regenerate generate docs validate-examples normalize-minimum-configs normalize-network-examples llms-txt install help download-specs sweep sweep-dry-run testacc testacc-mock testacc-real testacc-staging testacc-all test-report test-comprehensive test-comprehensive-mock test-comprehensive-real test-pr-subset uat

# Default target
all: generate build lint test docs

# Help
help:
	@echo "terraform-provider-xcsh Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make              - Generate, build, lint, test, and generate docs"
	@echo "  make build        - Build the provider binary"
	@echo "  make test         - Run tests"
	@echo "  make lint         - Run linters"
	@echo "  make fmt          - Format Go code"
	@echo "  make generate     - Generate resources from OpenAPI specs"
	@echo "  make download-specs - Download latest F5 XC API specs (requires gh CLI)"
	@echo "  make docs         - Generate Terraform documentation"
	@echo "  make clean        - Remove build artifacts"
	@echo "  make install      - Install provider locally"
	@echo ""
	@echo "Acceptance Testing (Categorized):"
	@echo "  make testacc      - Run all acceptance tests (requires F5XC credentials)"
	@echo "  make testacc-real - Run REAL API tests only (TestAcc* prefix)"
	@echo "  make testacc-mock - Run MOCK API tests only (TestMock* prefix)"
	@echo "  make testacc-all     - Run both real and mock tests with report"
	@echo "  make testacc-staging - Run curated staging tests (15 representative resources)"
	@echo "  make test-report     - Generate test report from last test run"
	@echo ""
	@echo "Comprehensive Testing (CI/CD):"
	@echo "  make test-comprehensive      - Full test suite with professional reports"
	@echo "  make test-comprehensive-mock - Mock tests only (parallel, fast)"
	@echo "  make test-comprehensive-real - Real API tests only (sequential, rate-limited)"
	@echo "  make test-pr-subset          - PR validation (mock tests only)"
	@echo ""
	@echo "Test Categories:"
	@echo "  REAL_API (TestAcc*) - Tests against real F5 XC API endpoints"
	@echo "  MOCK_API (TestMock*) - Tests against local mock server"
	@echo "  UNIT (Test*) - Unit tests without external dependencies"
	@echo ""
	@echo "Test Resource Cleanup:"
	@echo "  make sweep        - Clean up ALL orphaned test resources (prefix-based)"
	@echo "                      WARNING: Deletes any resource with tf-acc-test-* or tf-test-* prefix"
	@echo "                      Use only when no other users are running tests on the same tenant"
	@echo "  make sweep-resource RESOURCE=xcsh_namespace - Sweep specific resource type"
	@echo ""
	@echo "  For SAFE multi-user cleanup, use CleanupTracked() in your test code:"
	@echo "    defer acctest.CleanupTracked()  // Only deletes resources THIS test created"
	@echo ""
	@echo "Environment Variables:"
	@echo "  TF_ACC=1           - Enable real acceptance tests"
	@echo "  XCSH_MOCK_MODE=1   - Enable mock server tests"
	@echo "  SPEC_DIR           - Directory containing OpenAPI specs (default: docs/specifications/api)"
	@echo "  XCSH_SPEC_DIR      - Alternative env var for spec directory"
	@echo ""
	@echo "For real acceptance tests, set one of:"
	@echo "  XCSH_API_URL + XCSH_P12_FILE + XCSH_P12_PASSWORD (P12 auth)"
	@echo "  XCSH_API_URL + XCSH_CERT + XCSH_KEY (PEM auth)"
	@echo "  XCSH_API_URL + XCSH_API_TOKEN (Token auth)"

# Build the provider
build:
	@echo "Building $(BINARY_NAME)..."
	$(GO) build -ldflags="-X main.version=$(VERSION)" -o $(BINARY_NAME) .

# Run tests
test:
	@echo "Running tests..."
	$(GO) test -v -race ./internal/... ./tools/...

# Run linters
lint:
	@echo "Running linters..."
	$(GOLINT) run --timeout=5m ./internal/... . ./tools/...

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) -s -w .

# Download the exact pinned F5 XC API specs from the enriched repository.
download-specs:
	@echo "Downloading pinned F5 XC API specs..."
	@mkdir -p $(SPEC_DIR)
	@RELEASE_TAG=$$(tr -d '[:space:]' < tools/spec-version.txt); \
		echo "Version: $$RELEASE_TAG"; \
		STAGING_ROOT=$$(mktemp -d); \
		trap 'rm -rf "$$STAGING_ROOT"' EXIT; \
		SPEC_ZIP="$$STAGING_ROOT/specs.zip"; \
		BUNDLE_NAME="f5xc-api-specs-$$RELEASE_TAG.zip"; \
		ASSET_ID=$$(gh api repos/$(ENRICHED_REPO)/releases/tags/$$RELEASE_TAG --jq \
			"[.assets[] | select(.name==\"$$BUNDLE_NAME\")] | if length == 1 then .[0].id else empty end"); \
		[ -n "$$ASSET_ID" ] || { echo "Missing required release asset: $$BUNDLE_NAME" >&2; exit 1; }; \
		gh api repos/$(ENRICHED_REPO)/releases/assets/$$ASSET_ID \
			-H "Accept: application/octet-stream" > "$$SPEC_ZIP"; \
		EXPECTED_SHA=$$(jq -r --arg name "$$BUNDLE_NAME" '.assets[$$name]' tools/spec-release.json); \
		ACTUAL_SHA="sha256:$$(sha256sum "$$SPEC_ZIP" | awk '{print $$1}')"; \
		[ "$$ACTUAL_SHA" = "$$EXPECTED_SHA" ] || { echo "Release digest mismatch: $$BUNDLE_NAME" >&2; exit 1; }; \
		rm -rf $(SPEC_DIR)/*; \
		unzip -o "$$SPEC_ZIP" -d $(SPEC_DIR); \
		for asset in \
			api-catalog.json \
			concurrency_contracts.json \
			index.json \
			minimal-export-defaults.json \
			openapi.json \
			smsv2-contract-manifest.json \
			smsv2-contract.json \
			smsv2-evidence-receipt.json \
			smsv2_parity_manifest.json \
			upstream-contract-removals.json; do \
			AID=$$(gh api repos/$(ENRICHED_REPO)/releases/tags/$$RELEASE_TAG --jq \
				"[.assets[] | select(.name==\"$$asset\")] | if length == 1 then .[0].id else empty end"); \
			[ -n "$$AID" ] || { echo "Missing required release asset: $$asset" >&2; exit 1; }; \
			gh api repos/$(ENRICHED_REPO)/releases/assets/$$AID \
				-H "Accept: application/octet-stream" > $(SPEC_DIR)/$$asset; \
			EXPECTED_SHA=$$(jq -r --arg name "$$asset" '.assets[$$name]' tools/spec-release.json); \
			ACTUAL_SHA="sha256:$$(sha256sum $(SPEC_DIR)/$$asset | awk '{print $$1}')"; \
			[ "$$ACTUAL_SHA" = "$$EXPECTED_SHA" ] || { echo "Release digest mismatch: $$asset" >&2; exit 1; }; \
		done; \
		echo "Specs downloaded to $(SPEC_DIR)"

# Generate resources from OpenAPI specs
generate: generate-schemas
	@echo "Generation complete"

# Derive the import-default suppression data file from discovered API defaults
# (tools/api-defaults.json). Run after 'make discover' (or the discover-defaults
# workflow) to auto-populate server-default oneof members suppressed on import.
# Consumed by the code generator; see issue #1006.
emit-import-suppressions:
	@echo "Deriving import-default suppressions from $(TOOLS_DIR)/api-defaults.json..."
	@$(GO) run $(TOOLS_DIR)/emit-import-suppressions.go

generate-schemas:
	@echo "Generating schemas from OpenAPI specs..."
	@if [ -d "$(SPEC_DIR)" ] && [ -f "$(SPEC_DIR)/index.json" ] && [ -d "$(SPEC_DIR)/domains" ]; then \
		$(GO) run $(TOOLS_DIR)/generate-all-schemas.go --spec-dir=$(SPEC_DIR); \
	else \
		echo "No v2 OpenAPI specs found in $(SPEC_DIR). Skipping generation."; \
		echo "Run 'make download-specs' first, or set SPEC_DIR to a directory with index.json + domains/"; \
	fi
	@# Examples are generated schema-driven from the committed provider schema (spec-free),
	@# so this runs after the provider is generated and needs no specs. generate-examples
	@# also prunes orphan example dirs and runs terraform fmt.
	@echo "Generating Terraform examples from provider schema..."
	$(GO) run $(TOOLS_DIR)/generate-examples.go

# Generate Terraform documentation
docs:
	@echo "Generating Terraform documentation..."
	GOTOOLCHAIN=auto $(GO) install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.25.0
	PATH="$$(go env GOPATH)/bin:$$PATH" scripts/generate-provider-docs.sh

# Validate generated Terraform examples.
# Examples are schema-driven (generated from the TerraformAttribute tree), so EVERY example
# must terraform-validate. Any invalid example fails the build — no warnings-only escape hatch.
validate-examples:
	@echo "Regenerating and validating every provider example..."
	GOTOOLCHAIN=auto $(GO) install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.25.0
	PATH="$$(go env GOPATH)/bin:$$PATH" scripts/generate-provider-docs.sh

# Normalize embedded minimum-configuration examples at their source.
normalize-minimum-configs:
	$(GO) run $(TOOLS_DIR)/normalize-minimum-configs.go

# Normalize hand-maintained network fixtures to RFC 5737 documentation ranges.
normalize-network-examples:
	$(GO) run $(TOOLS_DIR)/normalize-network-examples.go

# Generate llms.txt hierarchy (L0 + L1 category + L2 per-resource)
llms-txt:
	@echo "Generating llms.txt hierarchy..."
	$(GO) run $(TOOLS_DIR)/generate-llms-txt.go
	@# json.MarshalIndent output is not biome-formatted; format the index so the
	@# committed file matches the biome-check gate (pre-commit + super-linter).
	@if command -v biome >/dev/null 2>&1; then \
		biome format --write docs/terraform-llms-index.json >/dev/null && echo "Formatted docs/terraform-llms-index.json with biome"; \
	else \
		echo "WARNING: biome not found — run 'biome format --write docs/terraform-llms-index.json' before committing"; \
	fi
	@echo "llms.txt hierarchy generation complete"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	rm -rf dist/

# Clean all generated files (for full regeneration)
clean-generated:
	@echo "Cleaning generated files..."
	rm -f $(PROVIDER_DIR)/*_resource.go
	rm -f $(PROVIDER_DIR)/*_data_source.go
	rm -f $(PROVIDER_DIR)/provider.go
	rm -f $(CLIENT_DIR)/*_types.go
	@echo "Generated files cleaned. Run 'make generate' to regenerate."

# Full clean rebuild from specs
regenerate: clean-generated generate
	@echo "Full regeneration complete"

# Install provider locally for testing
install: build
	@echo "Installing provider locally..."
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/f5-sales-demo/xcsh/$(VERSION)/$(GOOS)_$(GOARCH)
	cp $(BINARY_NAME) ~/.terraform.d/plugins/registry.terraform.io/f5-sales-demo/xcsh/$(VERSION)/$(GOOS)_$(GOARCH)/

# Acceptance testing and cleanup
testacc:
	@echo "Running acceptance tests..."
	TF_ACC=1 $(GO) test -v -timeout 120m ./internal/provider/...

# Sweep test resources - clean up orphaned resources from failed tests
# WARNING: This will delete ALL resources matching test prefixes, including
# resources created by other users. For safe multi-user cleanup, use
# CleanupTracked() in your test code instead.
#
# Usage: make sweep
# Environment variables required:
#   - XCSH_API_URL: F5 XC API URL
#   - XCSH_P12_FILE and XCSH_P12_PASSWORD (for P12 auth)
#   - OR XCSH_CERT and XCSH_KEY (for PEM auth)
#   - OR XCSH_API_TOKEN (for token auth)
sweep:
	@echo "⚠️  WARNING: Prefix-based sweep - will delete ALL test resources!"
	@echo "Sweeping resources with prefix 'tf-acc-test-' or 'tf-test-'..."
	@echo "This may delete resources created by other users on the same tenant."
	@echo ""
	@echo "For SAFE multi-user cleanup, use CleanupTracked() in your tests."
	@echo ""
	TF_ACC=1 $(GO) test ./internal/acctest -v -sweep=all -timeout 30m

# Sweep specific resource type
# Usage: make sweep-resource RESOURCE=xcsh_namespace
sweep-resource:
	@if [ -z "$(RESOURCE)" ]; then \
		echo "Error: RESOURCE variable not set"; \
		echo "Usage: make sweep-resource RESOURCE=xcsh_namespace"; \
		exit 1; \
	fi
	@echo "Sweeping $(RESOURCE) resources..."
	TF_ACC=1 $(GO) test ./internal/acctest -v -sweep=$(RESOURCE) -timeout 30m

# CI targets
.PHONY: ci ci-lint ci-test ci-build ci-generate

ci: ci-generate ci-build ci-lint ci-test

ci-generate:
	@echo "CI: Generating schemas (if specs available)..."
	@if [ -d "$(SPEC_DIR)" ] && [ -f "$(SPEC_DIR)/index.json" ] && [ -d "$(SPEC_DIR)/domains" ]; then \
		$(GO) run $(TOOLS_DIR)/generate-all-schemas.go --spec-dir=$(SPEC_DIR); \
	fi

ci-build:
	@echo "CI: Building..."
	$(GO) build -v .

ci-lint:
	@echo "CI: Linting..."
	$(GO) install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
	$(GOLINT) run --timeout=5m ./internal/... . ./tools/...

ci-test:
	@echo "CI: Testing..."
	$(GO) test -v -race ./internal/... ./tools/...

# Release preparation
.PHONY: release-prep
release-prep: fmt lint test docs
	@echo "Release preparation complete"
	@echo "Ensure all changes are committed before tagging"

# Verify no uncommitted generated changes
.PHONY: verify-generate
# Covers examples/ and docs/ as well as the Go output: those are generated too, and
# leaving them out is why 143 examples and 91 docs drifted unnoticed (#1397).
verify-generate: generate
	@echo "Verifying no uncommitted changes from generation..."
	@if [ -n "$$(git status --porcelain $(PROVIDER_DIR) $(CLIENT_DIR) examples docs tools/smsv2-parity-matrix.json)" ]; then \
		echo "Error: Generated files have uncommitted changes"; \
		git status --porcelain $(PROVIDER_DIR) $(CLIENT_DIR) examples docs tools/smsv2-parity-matrix.json; \
		exit 1; \
	fi
	@echo "All generated files are up to date"

# =============================================================================
# Categorized Acceptance Tests
# =============================================================================

# Run REAL API tests only (TestAcc* prefix)
# These tests require F5XC credentials and run against the real API
testacc-real:
	@echo "Running REAL API acceptance tests (TestAcc*)..."
	@echo "Category: REAL_API - Tests against real F5 XC API endpoints"
	@echo ""
	TF_ACC=1 $(GO) test -v -timeout 120m ./internal/provider/... -run "^TestAcc" 2>&1 | tee .test-output-real.txt
	@echo ""
	@echo "Test output saved to .test-output-real.txt"

# Run curated staging acceptance tests (representative subset across all domains)
# Sequential execution to avoid rate limiting. Requires XCSH_API_URL and XCSH_API_TOKEN.
# Covers 9 verified resources: Namespace, Healthcheck, OriginPool, AppFirewall,
# VirtualSite, AlertPolicy, AlertReceiver, ServicePolicy, GlobalLogReceiver
testacc-staging:
	@echo "Running curated staging acceptance tests..."
	@echo "Target: $${XCSH_API_URL}"
	@echo ""
	TF_ACC=1 $(GO) test -v -timeout 60m -count=1 -parallel=1 \
		./internal/provider/... \
		-run "^TestAcc(Namespace|Healthcheck|OriginPool|AppFirewall|VirtualSite|AlertPolicy|AlertReceiver|ServicePolicy|GlobalLogReceiver)Resource_basic$$" \
		2>&1 | tee .test-output-staging.txt
	@echo ""
	@echo "Test output saved to .test-output-staging.txt"

# Run MOCK API tests only (TestMock* prefix)
# These tests use the mock server and don't require real credentials
testacc-mock:
	@echo "Running MOCK API acceptance tests (TestMock*)..."
	@echo "Category: MOCK_API - Tests against local mock server"
	@echo ""
	TF_ACC=1 XCSH_MOCK_MODE=1 $(GO) test -v -timeout 30m ./internal/provider/... -run "^TestMock" 2>&1 | tee .test-output-mock.txt
	@echo ""
	@echo "Test output saved to .test-output-mock.txt"

# Run both real and mock tests with JSON output and generate report
testacc-all:
	@echo "Running ALL acceptance tests (Real + Mock) with categorized report..."
	@echo ""
	@echo "========================================================================"
	@echo "PHASE 1: MOCK API TESTS (no credentials required)"
	@echo "========================================================================"
	TF_ACC=1 XCSH_MOCK_MODE=1 $(GO) test -json -timeout 30m ./internal/provider/... -run "^TestMock" 2>&1 | tee .test-json-mock.txt | $(GO) run $(TOOLS_DIR)/test-report/main.go || true
	@echo ""
	@echo "========================================================================"
	@echo "PHASE 2: REAL API TESTS (requires credentials)"
	@echo "========================================================================"
	@if [ -n "$$XCSH_API_URL" ]; then \
		TF_ACC=1 $(GO) test -json -timeout 120m ./internal/provider/... -run "^TestAcc" 2>&1 | tee .test-json-real.txt | $(GO) run $(TOOLS_DIR)/test-report/main.go || true; \
	else \
		echo "⚠️  Skipping real API tests: XCSH_API_URL not set"; \
	fi
	@echo ""
	@echo "========================================================================"
	@echo "COMBINED REPORT"
	@echo "========================================================================"
	@cat .test-json-mock.txt .test-json-real.txt 2>/dev/null | $(GO) run $(TOOLS_DIR)/test-report/main.go || echo "No test data to report"

# Generate a test report from JSON test output
# Usage: go test -json ./... > test-output.json && make test-report
test-report:
	@echo "Generating test report..."
	@if [ -f ".test-json-mock.txt" ] || [ -f ".test-json-real.txt" ]; then \
		cat .test-json-mock.txt .test-json-real.txt 2>/dev/null | $(GO) run $(TOOLS_DIR)/test-report/main.go; \
	else \
		echo "No test output files found. Run tests with:"; \
		echo "  make testacc-all"; \
		echo "Or manually:"; \
		echo "  go test -json ./internal/provider/... | go run tools/test-report/main.go"; \
	fi

# Generate markdown test report
test-report-md:
	@echo "Generating markdown test report..."
	@cat .test-json-mock.txt .test-json-real.txt 2>/dev/null | $(GO) run $(TOOLS_DIR)/test-report/main.go -format=markdown -output=test-report.md
	@echo "Report saved to test-report.md"

# Clean test output files
clean-test-output:
	@echo "Cleaning test output files..."
	rm -f .test-output-*.txt .test-json-*.txt test-report.md test-report.json

# =============================================================================
# Comprehensive Testing (CI/CD Ready)
# =============================================================================
# These targets use the comprehensive test runner script that produces
# professional reports in multiple formats (Text, JSON, Markdown, JUnit XML).
#
# Key differences from testacc-* targets:
#   - Mock tests run in PARALLEL (no rate limiting - local tests)
#   - Real API tests run SEQUENTIAL with rate limiting (API protection)
#   - Generates JUnit XML for GitHub Actions test UI
#   - Detects transient errors (rate limit, timeout, connection)
#   - Categorizes skip reasons
#   - Tracks slowest tests for optimization

# Run full comprehensive test suite (mock + real)
# Mock tests run parallel, real API tests run sequential with rate limiting
test-comprehensive:
	@echo "Running comprehensive test suite..."
	./scripts/run-comprehensive-tests.sh --mode full

# Run mock tests only - PARALLEL (fast, no rate limiting needed)
test-comprehensive-mock:
	@echo "Running comprehensive mock tests (parallel)..."
	./scripts/run-comprehensive-tests.sh --mode mock-only

# Run real API tests only - SEQUENTIAL with rate limiting
# Requires: XCSH_API_URL, XCSH_P12_FILE, XCSH_P12_PASSWORD (or XCSH_API_TOKEN)
test-comprehensive-real:
	@echo "Running comprehensive real API tests (sequential)..."
	./scripts/run-comprehensive-tests.sh --mode real-only

# Run PR subset tests - mock tests only for PR validation
test-pr-subset:
	@echo "Running PR subset tests (mock only)..."
	./scripts/run-comprehensive-tests.sh --mode pr-subset

# Clean comprehensive test reports
clean-test-reports:
	@echo "Cleaning test reports..."
	rm -rf test-reports/

# Run UAT test harness for AI-generated Terraform plans
uat:
	@echo "Running UAT test harness..."
	tools/uat/run-uat.sh
