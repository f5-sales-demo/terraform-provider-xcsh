---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_api_crawler."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1221, "body_sha256": "sha256:4a31cac80ceccb351dee65dacd82aa886aea5e78a73b7bf8c48fa0b40c4d676a", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_crawler:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a73dd5a1c2c2dd7d614c3510b02a2ee9ecb351afd8ffe9c697a5142ace64fc30", "source_path": "examples/resources/xcsh_api_crawler/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:api_crawler:example:resource", "parent_id": "xcsh-docs:resources:api_crawler:examples", "path": "documentation/resources/api_crawler/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "api_crawler", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1333320203321222-1031031111030302-2110210221023332-2020100021013313-2022001012210210-2112211123110133-3211311312211331-3102211222213323", "registry_path": "docs/guides/resources--api_crawler--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_crawler/examples/resource/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Resource for xcsh_api_crawler.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_api_crawler/resource.tf`; digest `sha256:a73dd5a1c2c2dd7d614c3510b02a2ee9ecb351afd8ffe9c697a5142ace64fc30`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/examples/)
- [xcsh_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/)
