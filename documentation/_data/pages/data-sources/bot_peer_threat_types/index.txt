---
page_title: "xcsh_bot_peer_threat_types"
subcategory: ""
description: "xcsh_bot_peer_threat_types for xcsh_bot_peer_threat_types."
xcsh_docs: {"aliases": [], "body_bytes": 1265, "body_sha256": "sha256:01f4d8726afe9d6340300f490b060a0906fcd98a9e01b2a583358287785f6fb0", "child_ids": ["xcsh-docs:data-sources:bot_peer_threat_types:reference", "xcsh-docs:data-sources:bot_peer_threat_types:examples"], "collection_id": "xcsh-docs:data-sources:bot_peer_threat_types:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_threat_types:fundamentals", "parent_id": null, "path": "documentation/data-sources/bot_peer_threat_types/index.md", "provider_name": "bot_peer_threat_types", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_threat_types/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_bot_peer_threat_types for xcsh_bot_peer_threat_types.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bot_peer_threat_types

Breadcrumbs:

- xcsh_bot_peer_threat_types

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotPeerThreatTypes DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_threat_types" "example" {
  namespace = "example-value"
}

output "bot_peer_threat_types_result" {
  value = data.xcsh_bot_peer_threat_types.example
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_threat_types/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_threat_types/examples/)
