---
page_title: "xcsh_bot_endpoint_policy"
subcategory: ""
description: "xcsh_bot_endpoint_policy for xcsh_bot_endpoint_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1360, "body_sha256": "sha256:e5ea34da2317d8108c6ada954c931bfb595e469a5c76571bc271fc88fb6bf387", "canonical_id": "xcsh-docs:data-sources:bot_endpoint_policy:fundamentals", "child_ids": ["xcsh-docs:data-sources:bot_endpoint_policy:reference", "xcsh-docs:data-sources:bot_endpoint_policy:examples"], "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:fundamentals", "parent_id": null, "path": "docs/data-sources/bot_endpoint_policy.md", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_bot_endpoint_policy for xcsh_bot_endpoint_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bot_endpoint_policy

Breadcrumbs:

- xcsh_bot_endpoint_policy

Manages a Bot Endpoint Policy resource in F5 Distributed Cloud for get bot endpoint policy.
configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotEndpointPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotEndpointPolicy by name
data "xcsh_bot_endpoint_policy" "example" {
  name      = "example-bot-endpoint-policy"
  namespace = "staging"
}

output "bot_endpoint_policy_id" {
  value = data.xcsh_bot_endpoint_policy.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--bot_endpoint_policy--reference.md)
- [Examples](../guides/data-sources--bot_endpoint_policy--examples.md)
