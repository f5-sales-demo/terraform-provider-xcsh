---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_nat_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1031, "body_sha256": "sha256:b2110e1345441dbb083294441d154ddeb379ab11dbd5d48906825c0936001876", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d086b598bab70cc64257bc21568a65534a135b8fdb72d6f6952421f88e8adafe", "source_path": "examples/data-sources/xcsh_nat_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:nat_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:nat_policy:examples", "path": "documentation/data-sources/nat_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2100301132120231-2110131320221302-3302122201310033-1132133200033301-3300120310310021-1200333213323021-1231103221312301-1032111132023320", "registry_path": "docs/guides/data-sources--nat_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Data source for xcsh_nat_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["nat_policyCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nat_policy/data-source.tf`; digest `sha256:d086b598bab70cc64257bc21568a65534a135b8fdb72d6f6952421f88e8adafe`.

```terraform
# NATPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NATPolicy by name
data "xcsh_nat_policy" "example" {
  name      = "example-nat-policy"
  namespace = "staging"
}

output "nat_policy_id" {
  value = data.xcsh_nat_policy.example.id
}
```
