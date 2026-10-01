---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 1175, "body_sha256": "sha256:010f9c1041ffce340f93745c104a1da696a78c2a1b60ed9bf348aca1ac283c02", "canonical_id": "xcsh-docs:data-sources:policy_based_routing:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d902f95ac0d5d170c71a1f393d0718d60b4c9a9b4856200aec4f7a2fbf5ad0fc", "source_path": "examples/data-sources/xcsh_policy_based_routing/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:policy_based_routing:example:data-source", "parent_id": "xcsh-docs:data-sources:policy_based_routing:examples", "path": "docs/guides/data-sources--policy_based_routing--example--data-source.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md)
- [Examples](data-sources--policy_based_routing--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_policy_based_routing/data-source.tf`; digest `sha256:d902f95ac0d5d170c71a1f393d0718d60b4c9a9b4856200aec4f7a2fbf5ad0fc`.

```terraform
# PolicyBasedRouting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing PolicyBasedRouting by name
data "xcsh_policy_based_routing" "example" {
  name      = "example-policy-based-routing"
  namespace = "staging"
}

output "policy_based_routing_id" {
  value = data.xcsh_policy_based_routing.example.id
}
```

## Next pages

- [Examples](data-sources--policy_based_routing--examples.md)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md)
