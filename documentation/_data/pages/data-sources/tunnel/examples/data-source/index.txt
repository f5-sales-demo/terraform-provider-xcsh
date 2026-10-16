---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_tunnel."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 993, "body_sha256": "sha256:3a9d24b1399711866a0ebec110fb9ed4ba7d8b963a0d2ca68283ed5055824dbb", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6bb09ba6599023728c287794941b0e304024407337d302d5ca4e83c7bdd1f81d", "source_path": "examples/data-sources/xcsh_tunnel/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:tunnel:example:data-source", "parent_id": "xcsh-docs:data-sources:tunnel:examples", "path": "documentation/data-sources/tunnel/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-2310011211200203-2221333001022301-1121322111133330-0000313213120301-0020230122311012-3011230220120212-3323130220033033-1320130313311222", "registry_path": "docs/guides/data-sources--tunnel--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Data source for xcsh_tunnel.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["tunnelCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
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
