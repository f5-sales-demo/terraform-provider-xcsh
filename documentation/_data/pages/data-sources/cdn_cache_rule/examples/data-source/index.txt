---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1069, "body_sha256": "sha256:dcf56643141da70fdd553b2d02dc2d5eeb98a6cd72cd99a2c8a145c54c2638b4", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cdn_cache_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fc686ceb8d4fdca2d4b78094b5b3751ef65d5c692bd05f9089c3561d05468ad9", "source_path": "examples/data-sources/xcsh_cdn_cache_rule/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cdn_cache_rule:example:data-source", "parent_id": "xcsh-docs:data-sources:cdn_cache_rule:examples", "path": "documentation/data-sources/cdn_cache_rule/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3131132023323330-2231001231222123-1230022132103031-3223031123010101-3033002301133022-0212323320223202-3122311210222302-1013111031332123", "registry_path": "docs/guides/data-sources--cdn_cache_rule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_cache_rule/examples/data-source/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Data source for xcsh_cdn_cache_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cdn_cache_rule/data-source.tf`; digest `sha256:fc686ceb8d4fdca2d4b78094b5b3751ef65d5c692bd05f9089c3561d05468ad9`.

```terraform
# CDNCacheRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNCacheRule by name
data "xcsh_cdn_cache_rule" "example" {
  name      = "example-cdn-cache-rule"
  namespace = "staging"
}

output "cdn_cache_rule_id" {
  value = data.xcsh_cdn_cache_rule.example.id
}
```
