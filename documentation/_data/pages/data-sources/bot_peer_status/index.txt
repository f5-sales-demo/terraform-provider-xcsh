---
page_title: "xcsh_bot_peer_status"
subcategory: ""
description: "xcsh_bot_peer_status for xcsh_bot_peer_status."
xcsh_docs: {"aliases": [], "body_bytes": 1119, "body_sha256": "sha256:0cc88b9e822b7adffb1a0b27ad0812ad608df050b8e85515d777032c1e6a6c38", "child_ids": ["xcsh-docs:data-sources:bot_peer_status:reference", "xcsh-docs:data-sources:bot_peer_status:examples"], "collection_id": "xcsh-docs:data-sources:bot_peer_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_status:fundamentals", "parent_id": null, "path": "documentation/data-sources/bot_peer_status/index.md", "provider_name": "bot_peer_status", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_status/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_bot_peer_status for xcsh_bot_peer_status.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_status/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_status/examples/)
