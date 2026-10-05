---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1303, "body_sha256": "sha256:9884a2df95f48dec748229b3b4cba0fc79dc6168f2d88410b92a4367d34dd274", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cdn_cache_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fc686ceb8d4fdca2d4b78094b5b3751ef65d5c692bd05f9089c3561d05468ad9", "source_path": "examples/data-sources/xcsh_cdn_cache_rule/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cdn_cache_rule:example:data-source", "parent_id": "xcsh-docs:data-sources:cdn_cache_rule:examples", "path": "documentation/data-sources/cdn_cache_rule/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3131132023323330-2231001231222123-1230022132103031-3223031123010101-3033002301133022-0212323320223202-3122311210222302-1013111031332123", "registry_path": "docs/guides/data-sources--cdn_cache_rule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_cache_rule/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_cdn_cache_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/examples/)
- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
