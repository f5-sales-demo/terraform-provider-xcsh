---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_peer_threat_types."
xcsh_docs: {"aliases": [], "body_bytes": 1205, "body_sha256": "sha256:545b0c973c992ae43db7097a63e720bb3faec9094e581b5208249accb492248e", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_peer_threat_types:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:340fefb3e7c9dd8e5383ef09986b56ee68d03ed73a4bcf9a8c9ec927ce63eafa", "source_path": "examples/data-sources/xcsh_bot_peer_threat_types/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_peer_threat_types:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_peer_threat_types:examples", "path": "documentation/data-sources/bot_peer_threat_types/examples/data-source/index.md", "provider_name": "bot_peer_threat_types", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_threat_types/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bot_peer_threat_types.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
