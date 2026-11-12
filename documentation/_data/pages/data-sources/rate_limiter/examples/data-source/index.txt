---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_rate_limiter."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1051, "body_sha256": "sha256:e9aaf930b239b4d6db563013ce2d44011556768a397cacd14bcbf5017a7c997d", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:rate_limiter:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6445461d09bf17fe28d100ae80d4a2ae205fe7e6d895306172e50a8556fc3fc6", "source_path": "examples/data-sources/xcsh_rate_limiter/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:rate_limiter:example:data-source", "parent_id": "xcsh-docs:data-sources:rate_limiter:examples", "path": "documentation/data-sources/rate_limiter/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-0022322200220222-0122130010010231-1112213222011300-0122322033120111-0001201000200003-2003010003202010-2223332233120323-2000120213332103", "registry_path": "docs/guides/data-sources--rate_limiter--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter/examples/data-source/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Data source for xcsh_rate_limiter.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
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
