---
page_title: "xcsh_bot_peer_top_good_bots"
subcategory: ""
description: "Bot detection and defense configuration."
xcsh_docs: {"aliases": ["bot peer top good bots"], "body_bytes": 1284, "body_sha256": "sha256:1bd9d1a2bd80ce85fd0a1118bd99109caf95c13e4bd695272bab9a7c66b81bec", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_peer_top_good_bots:reference", "xcsh-docs:data-sources:bot_peer_top_good_bots:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_top_good_bots:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/bot_peer_top_good_bots/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_top_good_bots", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3030001313332130-0200112302220223-2231210320202203-3301330023332321-1010113230110223-1130212301220213-3012200130111102-0020213232203302", "registry_path": "docs/data-sources/bot_peer_top_good_bots.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_top_good_bots/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Bot detection and defense configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bot_peer_top_good_bots

Breadcrumbs:

- xcsh_bot_peer_top_good_bots

Bot detection and defense configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotPeerTopGoodBots DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_top_good_bots" "example" {
  namespace = "example-value"
}

output "bot_peer_top_good_bots_result" {
  value = data.xcsh_bot_peer_top_good_bots.example
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/examples/)
