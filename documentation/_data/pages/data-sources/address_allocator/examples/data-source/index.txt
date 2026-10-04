---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_address_allocator."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1344, "body_sha256": "sha256:5637bf2b003f076addf7fc46662ac562d3d141a48878ce03272fda3546cf91ee", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:address_allocator:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8199b921632de8cfafabbbf9803b9c84ef8413a91ef1f5c0757142ce1c4bea0d", "source_path": "examples/data-sources/xcsh_address_allocator/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:address_allocator:example:data-source", "parent_id": "xcsh-docs:data-sources:address_allocator:examples", "path": "documentation/data-sources/address_allocator/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "address_allocator", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3210130013211311-0020002222320233-0303230023102233-0321310111322301-3102130302120221-2332232303310002-1110020003210310-3232301222322001", "registry_path": "docs/guides/data-sources--address_allocator--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/address_allocator/examples/data-source/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Data source for xcsh_address_allocator.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["address_allocatorCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_address_allocator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_address_allocator/data-source.tf`; digest `sha256:8199b921632de8cfafabbbf9803b9c84ef8413a91ef1f5c0757142ce1c4bea0d`.

```terraform
# AddressAllocator Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AddressAllocator by name
data "xcsh_address_allocator" "example" {
  name      = "example-address-allocator"
  namespace = "staging"
}

output "address_allocator_id" {
  value = data.xcsh_address_allocator.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/examples/)
- [xcsh_address_allocator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/)
