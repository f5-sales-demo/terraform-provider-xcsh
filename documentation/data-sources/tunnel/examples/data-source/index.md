---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_tunnel."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1203, "body_sha256": "sha256:f956d6706727cfc8bed9f49bbd1e6fdb1254343e7395fb63707911c25c2c4f20", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6bb09ba6599023728c287794941b0e304024407337d302d5ca4e83c7bdd1f81d", "source_path": "examples/data-sources/xcsh_tunnel/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:tunnel:example:data-source", "parent_id": "xcsh-docs:data-sources:tunnel:examples", "path": "documentation/data-sources/tunnel/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2310011211200203-2221333001022301-1121322111133330-0000313213120301-0020230122311012-3011230220120212-3323130220033033-1320130313311222", "registry_path": "docs/guides/data-sources--tunnel--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_tunnel.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tunnelCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_tunnel/data-source.tf`; digest `sha256:6bb09ba6599023728c287794941b0e304024407337d302d5ca4e83c7bdd1f81d`.

```terraform
# Tunnel Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Tunnel by name
data "xcsh_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "staging"
}

output "tunnel_id" {
  value = data.xcsh_tunnel.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/examples/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
