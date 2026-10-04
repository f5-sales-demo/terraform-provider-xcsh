---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_virtual_host."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1279, "body_sha256": "sha256:d254da02daa4a8bd7b204ac4c32fdd5005954ed9b300e991dd9aaa2ea08e2a42", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:23ea7a52b77c427f7e6b79f3895a5a23c6caa53b8b9e855b6dcaac10831f3a72", "source_path": "examples/data-sources/xcsh_virtual_host/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:virtual_host:example:data-source", "parent_id": "xcsh-docs:data-sources:virtual_host:examples", "path": "documentation/data-sources/virtual_host/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3310330001110212-1232132132312200-0232230121322300-3003102003212202-3022012012310032-2323033202222022-2133023021331231-1003112013212303", "registry_path": "docs/guides/data-sources--virtual_host--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/examples/data-source/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Data source for xcsh_virtual_host.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_virtual_host/data-source.tf`; digest `sha256:23ea7a52b77c427f7e6b79f3895a5a23c6caa53b8b9e855b6dcaac10831f3a72`.

```terraform
# VirtualHost Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualHost by name
data "xcsh_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "staging"
}

output "virtual_host_id" {
  value = data.xcsh_virtual_host.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/examples/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
