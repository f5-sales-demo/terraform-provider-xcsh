---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_rate_limiter."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1279, "body_sha256": "sha256:9ec9f8f2e4f707ef0e3586832f68d0bf62cca6af743227f54fb4d832e4a9a3ab", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:rate_limiter:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6445461d09bf17fe28d100ae80d4a2ae205fe7e6d895306172e50a8556fc3fc6", "source_path": "examples/data-sources/xcsh_rate_limiter/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:rate_limiter:example:data-source", "parent_id": "xcsh-docs:data-sources:rate_limiter:examples", "path": "documentation/data-sources/rate_limiter/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0022322200220222-0122130010010231-1112213222011300-0122322033120111-0001201000200003-2003010003202010-2223332233120323-2000120213332103", "registry_path": "docs/guides/data-sources--rate_limiter--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_rate_limiter.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/examples/)
- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/)
