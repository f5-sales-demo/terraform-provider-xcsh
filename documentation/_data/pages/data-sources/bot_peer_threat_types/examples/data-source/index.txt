---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_peer_threat_types."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1304, "body_sha256": "sha256:821b65f3fbf51c09c060ccf17f3f78b3e7b936715fc3350fcf5eca24a1d669f0", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_threat_types:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:340fefb3e7c9dd8e5383ef09986b56ee68d03ed73a4bcf9a8c9ec927ce63eafa", "source_path": "examples/data-sources/xcsh_bot_peer_threat_types/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_peer_threat_types:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_peer_threat_types:examples", "path": "documentation/data-sources/bot_peer_threat_types/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_threat_types", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3303302102011220-1331032122300301-3302330113313021-3120312120332211-0110023003331220-2003113313201210-2011131202320023-2303030320200122", "registry_path": "docs/guides/data-sources--bot_peer_threat_types--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_threat_types/examples/data-source/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Data source for xcsh_bot_peer_threat_types.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": [], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_peer_threat_types](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_threat_types/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_threat_types/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_threat_types/data-source.tf`; digest `sha256:340fefb3e7c9dd8e5383ef09986b56ee68d03ed73a4bcf9a8c9ec927ce63eafa`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_threat_types/examples/)
- [xcsh_bot_peer_threat_types](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_threat_types/)
