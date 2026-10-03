---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_api_crawler."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1266, "body_sha256": "sha256:01670e350fd5d3c2f93a963d82d9aeda0b4a9cae95db3f789099d2b555de1d90", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_crawler:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3bf584a1f7a1b8c51ad4462cae33ab926ce0db3d2feff9eb685959704de3f510", "source_path": "examples/data-sources/xcsh_api_crawler/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:api_crawler:example:data-source", "parent_id": "xcsh-docs:data-sources:api_crawler:examples", "path": "documentation/data-sources/api_crawler/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "api_crawler", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3232313232030132-1313101131032321-3321122323312312-2121023201302333-1131210120221023-0201010031003033-3221100000321310-3302203000200300", "registry_path": "docs/guides/data-sources--api_crawler--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_crawler/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_api_crawler.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_api_crawler/data-source.tf`; digest `sha256:3bf584a1f7a1b8c51ad4462cae33ab926ce0db3d2feff9eb685959704de3f510`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/examples/)
- [xcsh_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/)
