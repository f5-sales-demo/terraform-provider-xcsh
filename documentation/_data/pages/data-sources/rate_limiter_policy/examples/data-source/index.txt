---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1368, "body_sha256": "sha256:68ed9f5fdd326bf3b8b14346db3b423a6f67170a9cd4eb786d732f94462749ff", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:rate_limiter_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:631ebdd1844abfdc3c403666a18c881b536cd51ce00cfb0a5d316458e5a24bba", "source_path": "examples/data-sources/xcsh_rate_limiter_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:rate_limiter_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:rate_limiter_policy:examples", "path": "documentation/data-sources/rate_limiter_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3232132001302223-0202101231322303-2121003002023330-2032000102311231-3221010313312332-1103122100323211-1232301303201101-1303112221111113", "registry_path": "docs/guides/data-sources--rate_limiter_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_rate_limiter_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_rate_limiter_policy/data-source.tf`; digest `sha256:631ebdd1844abfdc3c403666a18c881b536cd51ce00cfb0a5d316458e5a24bba`.

```terraform
# RateLimiterPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing RateLimiterPolicy by name
data "xcsh_rate_limiter_policy" "example" {
  name      = "example-rate-limiter-policy"
  namespace = "staging"
}

output "rate_limiter_policy_id" {
  value = data.xcsh_rate_limiter_policy.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/examples/)
- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/)
