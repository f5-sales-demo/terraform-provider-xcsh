---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1099, "body_sha256": "sha256:02437630eba2a63c09e97d17077566fd61f1dc9f969b6f1bb2dcbe0e52614bd0", "canonical_id": "xcsh-docs:data-sources:network_policy:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6c10eea9b1a5304339370438b5c012d3317ee2ecb95c539f946182ee126b8c01", "source_path": "examples/data-sources/xcsh_network_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:network_policy:examples", "path": "docs/guides/data-sources--network_policy--example--data-source.md", "provider_name": "network_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md)
- [Examples](data-sources--network_policy--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_policy/data-source.tf`; digest `sha256:6c10eea9b1a5304339370438b5c012d3317ee2ecb95c539f946182ee126b8c01`.

```terraform
# NetworkPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicy by name
data "xcsh_network_policy" "example" {
  name      = "example-network-policy"
  namespace = "staging"
}

output "network_policy_id" {
  value = data.xcsh_network_policy.example.id
}
```

## Next pages

- [Examples](data-sources--network_policy--examples.md)
- [xcsh_network_policy](../data-sources/network_policy.md)
