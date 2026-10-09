---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_peer_status."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1002, "body_sha256": "sha256:1a682d46b172daec658558f54c420f158085b755b6b8fc268ca9d19ce753b6b1", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_status:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:97b1bb9bbcf7e35e2e9068fc82fcedf34184ffea0125e8cf3c1304beb4255fcc", "source_path": "examples/data-sources/xcsh_bot_peer_status/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_peer_status:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_peer_status:examples", "path": "documentation/data-sources/bot_peer_status/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_status", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0221003332320320-1313110133232103-2211003222031301-0331101231032103-3030101102000201-0300120000321230-2311001010331233-1221013233132003", "registry_path": "docs/guides/data-sources--bot_peer_status--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_status/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_bot_peer_status.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_peer_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_status/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_status/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_status/data-source.tf`; digest `sha256:97b1bb9bbcf7e35e2e9068fc82fcedf34184ffea0125e8cf3c1304beb4255fcc`.

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
