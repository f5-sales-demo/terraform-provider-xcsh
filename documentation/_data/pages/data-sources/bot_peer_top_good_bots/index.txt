---
page_title: "xcsh_bot_peer_top_good_bots"
subcategory: ""
description: "Bot detection and defense configuration."
xcsh_docs: {"aliases": ["bot peer top good bots"], "body_bytes": 1297, "body_sha256": "sha256:b7c091163e5289bc1b97718f6b2d63179bb3e42e2ebdf72d93fd7f1c308cfa52", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_peer_top_good_bots:reference", "xcsh-docs:data-sources:bot_peer_top_good_bots:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_top_good_bots:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/bot_peer_top_good_bots/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_top_good_bots", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3030001313332130-0200112302220223-2231210320202203-3301330023332321-1010113230110223-1130212301220213-3012200130111102-0020213232203302", "registry_path": "docs/data-sources/bot_peer_top_good_bots.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_top_good_bots/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Bot detection and defense configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/examples/)
