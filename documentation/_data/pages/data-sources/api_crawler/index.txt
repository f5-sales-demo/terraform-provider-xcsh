---
page_title: "xcsh_api_crawler"
subcategory: ""
description: "xcsh_api_crawler for xcsh_api_crawler."
xcsh_docs: {"aliases": [], "body_bytes": 1292, "body_sha256": "sha256:57cd6e4b6405f9e4b17b90e8f7efcf000f53a26e2f3bdeb1292e23aecd445020", "child_ids": ["xcsh-docs:data-sources:api_crawler:reference", "xcsh-docs:data-sources:api_crawler:examples"], "collection_id": "xcsh-docs:data-sources:api_crawler:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_crawler:fundamentals", "parent_id": null, "path": "documentation/data-sources/api_crawler/index.md", "provider_name": "api_crawler", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_crawler/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_api_crawler for xcsh_api_crawler.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
# APICrawler Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APICrawler by name
data "xcsh_api_crawler" "example" {
  name      = "example-api-crawler"
  namespace = "staging"
}

output "api_crawler_id" {
  value = data.xcsh_api_crawler.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/examples/)
