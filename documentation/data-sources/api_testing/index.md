---
page_title: "xcsh_api_testing"
subcategory: ""
description: "Reads API Testing information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["api testing"], "body_bytes": 1306, "body_sha256": "sha256:40fda73b9fcfc661f3c712d2e2016a591e075af2a1457986b7989231eced7170", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:api_testing:reference", "xcsh-docs:data-sources:api_testing:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_testing:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/api_testing/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230", "registry_path": "docs/data-sources/api_testing.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Reads API Testing information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["api_testingCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_api_testing

Breadcrumbs:

- xcsh_api_testing

Reads API Testing information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APITesting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APITesting by name
data "xcsh_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"
}

output "api_testing_id" {
  value = data.xcsh_api_testing.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/examples/)
