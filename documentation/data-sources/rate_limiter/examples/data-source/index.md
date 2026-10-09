---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_rate_limiter."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1051, "body_sha256": "sha256:e9aaf930b239b4d6db563013ce2d44011556768a397cacd14bcbf5017a7c997d", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:rate_limiter:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6445461d09bf17fe28d100ae80d4a2ae205fe7e6d895306172e50a8556fc3fc6", "source_path": "examples/data-sources/xcsh_rate_limiter/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:rate_limiter:example:data-source", "parent_id": "xcsh-docs:data-sources:rate_limiter:examples", "path": "documentation/data-sources/rate_limiter/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0022322200220222-0122130010010231-1112213222011300-0122322033120111-0001201000200003-2003010003202010-2223332233120323-2000120213332103", "registry_path": "docs/guides/data-sources--rate_limiter--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_rate_limiter.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_rate_limiter/data-source.tf`; digest `sha256:6445461d09bf17fe28d100ae80d4a2ae205fe7e6d895306172e50a8556fc3fc6`.

```terraform
# RateLimiter Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing RateLimiter by name
data "xcsh_rate_limiter" "example" {
  name      = "example-rate-limiter"
  namespace = "staging"
}

output "rate_limiter_id" {
  value = data.xcsh_rate_limiter.example.id
}
```
