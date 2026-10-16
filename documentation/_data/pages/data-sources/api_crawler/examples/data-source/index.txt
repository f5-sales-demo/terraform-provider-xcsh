---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_api_crawler."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1041, "body_sha256": "sha256:b1824b671d8d5264536807d437ce9cccc7f5c5263e487429cdaf08310c23abd9", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_crawler:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3bf584a1f7a1b8c51ad4462cae33ab926ce0db3d2feff9eb685959704de3f510", "source_path": "examples/data-sources/xcsh_api_crawler/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:api_crawler:example:data-source", "parent_id": "xcsh-docs:data-sources:api_crawler:examples", "path": "documentation/data-sources/api_crawler/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "api_crawler", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3232313232030132-1313101131032321-3321122323312312-2121023201302333-1131210120221023-0201010031003033-3221100000321310-3302203000200300", "registry_path": "docs/guides/data-sources--api_crawler--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_crawler/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Data source for xcsh_api_crawler.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
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
