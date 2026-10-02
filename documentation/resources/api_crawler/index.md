---
page_title: "xcsh_api_crawler"
subcategory: ""
description: "Manages a API Crawler resource in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["api crawler"], "body_bytes": 1485, "body_sha256": "sha256:e82cbd2de415cdc5385d73756f8ce822c851135fe4f327e9797c2de4cd030c19", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_crawler:reference", "xcsh-docs:resources:api_crawler:examples", "xcsh-docs:resources:api_crawler:import", "xcsh-docs:resources:api_crawler:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_crawler:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_crawler:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/api_crawler/index.md", "product": "distributed-cloud", "provider_name": "api_crawler", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3133010202200222-3213011120000312-1321322311223200-2311232132101333-2211031002320321-3221011221222302-0233012202301010-3201312221013200", "registry_path": "docs/resources/api_crawler.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_crawler/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a API Crawler resource in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_api_crawler

Breadcrumbs:

- xcsh_api_crawler

Manages a API Crawler resource in F5 Distributed Cloud.

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/lifecycle/timeouts/)
