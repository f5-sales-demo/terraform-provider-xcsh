---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_peer_status."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1239, "body_sha256": "sha256:8f0c48a1c2a5c84e7f795259366ff8c9bc13d573e034853b761d016e0cd5d452", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_status:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:97b1bb9bbcf7e35e2e9068fc82fcedf34184ffea0125e8cf3c1304beb4255fcc", "source_path": "examples/data-sources/xcsh_bot_peer_status/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_peer_status:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_peer_status:examples", "path": "documentation/data-sources/bot_peer_status/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_status", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0221003332320320-1313110133232103-2211003222031301-0331101231032103-3030101102000201-0300120000321230-2311001010331233-1221013233132003", "registry_path": "docs/guides/data-sources--bot_peer_status--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_status/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_bot_peer_status.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_status/examples/)
- [xcsh_bot_peer_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_status/)
