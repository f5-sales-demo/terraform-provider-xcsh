---
page_title: "xcsh_api_crawler"
subcategory: ""
description: "xcsh_api_crawler for xcsh_api_crawler."
xcsh_docs: {"aliases": [], "body_bytes": 1386, "body_sha256": "sha256:35bc384fc4c11763d89131c361fbbfb7b41764e6babecbe8a22ec670b6033f91", "child_ids": ["xcsh-docs:resources:api_crawler:reference", "xcsh-docs:resources:api_crawler:examples", "xcsh-docs:resources:api_crawler:import", "xcsh-docs:resources:api_crawler:timeouts"], "collection_id": "xcsh-docs:resources:api_crawler:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_crawler:fundamentals", "parent_id": null, "path": "documentation/resources/api_crawler/index.md", "provider_name": "api_crawler", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_crawler/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_api_crawler for xcsh_api_crawler.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
