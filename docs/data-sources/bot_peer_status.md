---
page_title: "xcsh_bot_peer_status"
subcategory: ""
description: "xcsh_bot_peer_status for xcsh_bot_peer_status."
xcsh_docs: {"aliases": [], "body_bytes": 1133, "body_sha256": "sha256:05f2cd721c2e6ad6a13704e5b52d6f6fb31e7128f5ea8c8467610d0ba0b2a435", "canonical_id": "xcsh-docs:data-sources:bot_peer_status:fundamentals", "child_ids": ["xcsh-docs:data-sources:bot_peer_status:reference", "xcsh-docs:data-sources:bot_peer_status:examples"], "collection_id": "xcsh-docs:data-sources:bot_peer_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_status:fundamentals", "parent_id": null, "path": "docs/data-sources/bot_peer_status.md", "provider_name": "bot_peer_status", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_status/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_bot_peer_status for xcsh_bot_peer_status.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bot_peer_status

Breadcrumbs:

- xcsh_bot_peer_status

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotPeerStatus DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_status" "example" {
  namespace = "example-value"
}

output "bot_peer_status_result" {
  value = data.xcsh_bot_peer_status.example
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--bot_peer_status--reference.md)
- [Examples](../guides/data-sources--bot_peer_status--examples.md)
