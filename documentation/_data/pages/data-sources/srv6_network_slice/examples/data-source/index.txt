---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_srv6_network_slice."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1108, "body_sha256": "sha256:2bdac929e8748959718d2d5a507f19ca0d568fae38a5f7b2bfb22f9fe8246c39", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:srv6_network_slice:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5f3b6cfc9cb152dc2c604228803d1e4acdb80bbc9c71bb42003bba14ed6bbdd1", "source_path": "examples/data-sources/xcsh_srv6_network_slice/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:srv6_network_slice:example:data-source", "parent_id": "xcsh-docs:data-sources:srv6_network_slice:examples", "path": "documentation/data-sources/srv6_network_slice/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "srv6_network_slice", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3202233013313201-3100332002310112-3103033200121200-3031101001110300-1210103233130211-1122300322000211-2132011222121230-1212022312211011", "registry_path": "docs/guides/data-sources--srv6_network_slice--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/srv6_network_slice/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_srv6_network_slice.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["srv6_network_sliceCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_srv6_network_slice](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_srv6_network_slice/data-source.tf`; digest `sha256:5f3b6cfc9cb152dc2c604228803d1e4acdb80bbc9c71bb42003bba14ed6bbdd1`.

```terraform
# Srv6NetworkSlice Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Srv6NetworkSlice by name
data "xcsh_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"
}

output "srv6_network_slice_id" {
  value = data.xcsh_srv6_network_slice.example.id
}
```
