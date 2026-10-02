---
page_title: "xcsh_address_allocator"
subcategory: ""
description: "Manages Address Allocator will create an address allocator object in 'system' namespace of the user in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["address allocator"], "body_bytes": 1735, "body_sha256": "sha256:beed13a9ca14415e8dc37d5c41874630508c6dee4d14575c86f8b1de34de60f0", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:address_allocator:reference", "xcsh-docs:resources:address_allocator:examples", "xcsh-docs:resources:address_allocator:import", "xcsh-docs:resources:address_allocator:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:address_allocator:collection", "completeness": "complete", "id": "xcsh-docs:resources:address_allocator:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/address_allocator/index.md", "product": "distributed-cloud", "provider_name": "address_allocator", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2221003231222131-0012132022313001-1331101002320300-2210032101320132-3133132121300222-1310012132000200-0313122021123003-3221012203222322", "registry_path": "docs/resources/address_allocator.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/address_allocator/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages Address Allocator will create an address allocator object in 'system' namespace of the user in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["address_allocatorCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_address_allocator

Breadcrumbs:

- xcsh_address_allocator

Manages Address Allocator will create an address allocator object in 'system' namespace of the user
in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `address_pool`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/address_allocator/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/address_allocator/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/address_allocator/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/address_allocator/lifecycle/timeouts/)
