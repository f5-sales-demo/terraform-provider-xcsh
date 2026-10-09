---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_api_crawler."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1002, "body_sha256": "sha256:83371d2ee4fcf4eca732854ad5355fe1eddf870a2f74e43fde603ba9b69c46ed", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_crawler:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a73dd5a1c2c2dd7d614c3510b02a2ee9ecb351afd8ffe9c697a5142ace64fc30", "source_path": "examples/resources/xcsh_api_crawler/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:api_crawler:example:resource", "parent_id": "xcsh-docs:resources:api_crawler:examples", "path": "documentation/resources/api_crawler/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "api_crawler", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1333320203321222-1031031111030302-2110210221023332-2020100021013313-2022001012210210-2112211123110133-3211311312211331-3102211222213323", "registry_path": "docs/guides/resources--api_crawler--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_crawler/examples/resource/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Resource for xcsh_api_crawler.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
