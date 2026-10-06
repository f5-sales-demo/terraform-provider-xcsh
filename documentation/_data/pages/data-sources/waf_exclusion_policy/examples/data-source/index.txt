---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_waf_exclusion_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1129, "body_sha256": "sha256:90c623158c68ef6ee81292374a3e5c844cb18011e1a3a8dcbb1fccb0e6e79887", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_exclusion_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e7c1c87b69af9a5626274b0968e5e1ac1658731a6d77764810ad0bd8f60e1d5e", "source_path": "examples/data-sources/xcsh_waf_exclusion_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:waf_exclusion_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:waf_exclusion_policy:examples", "path": "documentation/data-sources/waf_exclusion_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-3031113130100123-3132321120330101-0202323213023301-0031311220212222-0031210002221011-2223102101011310-3112010201213110-3131023221022322", "registry_path": "docs/guides/data-sources--waf_exclusion_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_exclusion_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Data source for xcsh_waf_exclusion_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_exclusion_policy/data-source.tf`; digest `sha256:e7c1c87b69af9a5626274b0968e5e1ac1658731a6d77764810ad0bd8f60e1d5e`.

```terraform
# WAFExclusionPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing WAFExclusionPolicy by name
data "xcsh_waf_exclusion_policy" "example" {
  name      = "example-waf-exclusion-policy"
  namespace = "staging"
}

output "waf_exclusion_policy_id" {
  value = data.xcsh_waf_exclusion_policy.example.id
}
```
