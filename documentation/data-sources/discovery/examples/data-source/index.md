---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_discovery."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1242, "body_sha256": "sha256:599b7d9678d8ff604ffe405310d651f88567d7fc4ae6e0a9f1c886515e46fb7f", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:de2200d3e23389756392f1e95f9b7544c16ad111ecd7a837e1de4f6caed18861", "source_path": "examples/data-sources/xcsh_discovery/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:discovery:example:data-source", "parent_id": "xcsh-docs:data-sources:discovery:examples", "path": "documentation/data-sources/discovery/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1121032231312020-2323313120122112-3033101332110102-2122233330232211-2330102211332322-3020033233123213-2121023211303312-0120030000011231", "registry_path": "docs/guides/data-sources--discovery--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_discovery/data-source.tf`; digest `sha256:de2200d3e23389756392f1e95f9b7544c16ad111ecd7a837e1de4f6caed18861`.

```terraform
# Discovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Discovery by name
data "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}

output "discovery_id" {
  value = data.xcsh_discovery.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/examples/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
