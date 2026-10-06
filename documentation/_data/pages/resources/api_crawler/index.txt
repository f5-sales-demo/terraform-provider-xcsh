---
page_title: "xcsh_api_crawler"
subcategory: ""
description: "Manages an API Crawler resource in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["api crawler"], "body_bytes": 1499, "body_sha256": "sha256:d5da8a2159afe5ac7d1a33a1f0eed2b3795619f02b30a43253b6226592717f47", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_crawler:reference", "xcsh-docs:resources:api_crawler:examples", "xcsh-docs:resources:api_crawler:import", "xcsh-docs:resources:api_crawler:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_crawler:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_crawler:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/api_crawler/index.md", "product": "distributed-cloud", "provider_name": "api_crawler", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3133010202200222-3213011120000312-1321322311223200-2311232132101333-2211031002320321-3221011221222302-0233012202301010-3201312221013200", "registry_path": "docs/resources/api_crawler.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_crawler/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Manages an API Crawler resource in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_api_crawler

Breadcrumbs:

- xcsh_api_crawler

Manages an API Crawler resource in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APICrawler Resource Example
# Manages a API Crawler resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APICrawler configuration
resource "xcsh_api_crawler" "example" {
  name      = "example-api-crawler"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/lifecycle/timeouts/)
