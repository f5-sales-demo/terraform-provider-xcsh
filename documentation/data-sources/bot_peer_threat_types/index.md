---
page_title: "xcsh_bot_peer_threat_types"
subcategory: ""
description: "Reads Bot Peer Threat Types information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["bot peer threat types"], "body_bytes": 1316, "body_sha256": "sha256:955dcde9d9d71a2c46af54ff05f8ff74dacae0d5701b1d6697f52a338e0f9449", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_peer_threat_types:reference", "xcsh-docs:data-sources:bot_peer_threat_types:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_threat_types:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_threat_types:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/bot_peer_threat_types/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_threat_types", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-2013010020311013-0321123032202111-3000312123121121-3301311011300221-1213110110313232-0030311231210211-3033301232310333-0010002333121001", "registry_path": "docs/data-sources/bot_peer_threat_types.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_threat_types/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Reads Bot Peer Threat Types information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": [], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bot_peer_threat_types

Breadcrumbs:

- xcsh_bot_peer_threat_types

Reads Bot Peer Threat Types information from F5 Distributed Cloud.

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_threat_types/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_threat_types/examples/)
