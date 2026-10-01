---
page_title: "xcsh_bot_network_policy"
subcategory: ""
description: "xcsh_bot_network_policy for xcsh_bot_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1348, "body_sha256": "sha256:3c0b295516cdaacc43faf8071477d65dbcd5e13ecffd31f3dbb0b474b3c22cc7", "canonical_id": "xcsh-docs:data-sources:bot_network_policy:fundamentals", "child_ids": ["xcsh-docs:data-sources:bot_network_policy:reference", "xcsh-docs:data-sources:bot_network_policy:examples"], "collection_id": "xcsh-docs:data-sources:bot_network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_network_policy:fundamentals", "parent_id": null, "path": "docs/data-sources/bot_network_policy.md", "provider_name": "bot_network_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_network_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_bot_network_policy for xcsh_bot_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bot_network_policy

Breadcrumbs:

- xcsh_bot_network_policy

Manages a Bot Network Policy resource in F5 Distributed Cloud for get bot network policy.
configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotNetworkPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotNetworkPolicy by name
data "xcsh_bot_network_policy" "example" {
  name      = "example-bot-network-policy"
  namespace = "staging"
}

output "bot_network_policy_id" {
  value = data.xcsh_bot_network_policy.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--bot_network_policy--reference.md)
- [Examples](../guides/data-sources--bot_network_policy--examples.md)
