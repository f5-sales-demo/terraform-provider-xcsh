---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_address_allocator."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1392, "body_sha256": "sha256:e7824cedcd89d96c38e373a69373a3e2783bdd6d70f7a8fa151f75e05d4afec8", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:address_allocator:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0d0c2e8aebf1264b1b9bd98f9f1564545390bccd08c3989c8f49aed2a844faa2", "source_path": "examples/resources/xcsh_address_allocator/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:address_allocator:example:resource", "parent_id": "xcsh-docs:resources:address_allocator:examples", "path": "documentation/resources/address_allocator/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "address_allocator", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2200201023213310-0330310001123331-2033001310110202-2023100221002013-1313323312211120-0002122100113112-1300221101123201-0230210303031011", "registry_path": "docs/guides/resources--address_allocator--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/address_allocator/examples/resource/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Resource for xcsh_address_allocator.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["address_allocatorCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_address_allocator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/address_allocator/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/address_allocator/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_address_allocator/resource.tf`; digest `sha256:0d0c2e8aebf1264b1b9bd98f9f1564545390bccd08c3989c8f49aed2a844faa2`.

```terraform
# AddressAllocator Resource Example
# Manages Address Allocator will create an address allocator object in 'system' namespace of the user in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AddressAllocator configuration
resource "xcsh_address_allocator" "example" {
  name      = "example-address-allocator"
  namespace = "staging"

  address_pool = ["example-value"]
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/address_allocator/examples/)
- [xcsh_address_allocator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/address_allocator/)
