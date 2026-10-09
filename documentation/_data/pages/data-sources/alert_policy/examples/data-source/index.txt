---
page_title: "Data source"
subcategory: "Monitoring"
description: "Data source for xcsh_alert_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1051, "body_sha256": "sha256:d34191cf554d54ca43be8fe806a7400a23052487e163d0c1ce427139ccbca220", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c6152f124e4da7e8b1a6f94adbb961ca4985954776a9f55f7df6b34b8d892657", "source_path": "examples/data-sources/xcsh_alert_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:alert_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:alert_policy:examples", "path": "documentation/data-sources/alert_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2230333020132001-0302320213001132-2013223221312333-0323213112021313-3102021333300103-2003123031312221-2233311201030302-2121032010321323", "registry_path": "docs/guides/data-sources--alert_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_alert_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["alert_policyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_alert_policy/data-source.tf`; digest `sha256:c6152f124e4da7e8b1a6f94adbb961ca4985954776a9f55f7df6b34b8d892657`.

```terraform
# AlertPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertPolicy by name
data "xcsh_alert_policy" "example" {
  name      = "example-alert-policy"
  namespace = "staging"
}

output "alert_policy_id" {
  value = data.xcsh_alert_policy.example.id
}
```
