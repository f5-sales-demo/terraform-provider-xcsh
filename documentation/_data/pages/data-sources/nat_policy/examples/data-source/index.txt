---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1253, "body_sha256": "sha256:b0d28ff4d1fa2fb1099a2108f6f2ff3ba0a6cb6277d5b7532d10df0d9d963159", "child_ids": [], "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d086b598bab70cc64257bc21568a65534a135b8fdb72d6f6952421f88e8adafe", "source_path": "examples/data-sources/xcsh_nat_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:nat_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:nat_policy:examples", "path": "documentation/data-sources/nat_policy/examples/data-source/index.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/examples/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
