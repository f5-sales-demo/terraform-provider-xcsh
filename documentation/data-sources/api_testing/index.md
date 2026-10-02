---
page_title: "xcsh_api_testing"
subcategory: ""
description: "Manages a API Testing resource in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["api testing"], "body_bytes": 1292, "body_sha256": "sha256:1142aa523265ed335a6543f867d100b7466dea29d9e4d9d0620576075fe3d31c", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:api_testing:reference", "xcsh-docs:data-sources:api_testing:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_testing:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/api_testing/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230", "registry_path": "docs/data-sources/api_testing.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a API Testing resource in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_testingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_api_testing

Breadcrumbs:

- xcsh_api_testing

Manages a API Testing resource in F5 Distributed Cloud.

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/examples/)
