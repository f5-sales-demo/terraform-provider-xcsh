---
page_title: "Data source"
subcategory: "Networking"
description: "Data source for xcsh_virtual_network."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1080, "body_sha256": "sha256:d8b6737b74cf8351d4e3cfe0b3b64b2b0dc7599ac2e3379cc219d7e4f3ff33f7", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:virtual_network:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:82f2d94c3d011a8d200fcd015e8a3406aea69e2a77a8f9d9eea9796b4b19e3bd", "source_path": "examples/data-sources/xcsh_virtual_network/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:virtual_network:example:data-source", "parent_id": "xcsh-docs:data-sources:virtual_network:examples", "path": "documentation/data-sources/virtual_network/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "virtual_network", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1302322113330001-3031103023210212-2332230121232333-0131301001300302-1101321011330003-0303030300130231-0323030122313222-2302102332031230", "registry_path": "docs/guides/data-sources--virtual_network--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_network/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_virtual_network.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_virtual_network/data-source.tf`; digest `sha256:82f2d94c3d011a8d200fcd015e8a3406aea69e2a77a8f9d9eea9796b4b19e3bd`.

```terraform
# VirtualNetwork Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualNetwork by name
data "xcsh_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"
}

output "virtual_network_id" {
  value = data.xcsh_virtual_network.example.id
}
```
