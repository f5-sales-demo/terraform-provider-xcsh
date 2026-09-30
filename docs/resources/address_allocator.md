---
page_title: "xcsh_address_allocator"
subcategory: ""
description: "xcsh_address_allocator for xcsh_address_allocator."
xcsh_docs: {"aliases": [], "body_bytes": 1447, "body_sha256": "sha256:066184e96fc2fd7cb329896cc63fefab2f80a9fd5c6149cd3b5c693bc48143d7", "canonical_id": "xcsh-docs:resources:address_allocator:fundamentals", "child_ids": ["xcsh-docs:resources:address_allocator:reference", "xcsh-docs:resources:address_allocator:examples", "xcsh-docs:resources:address_allocator:import", "xcsh-docs:resources:address_allocator:timeouts"], "collection_id": "xcsh-docs:resources:address_allocator:collection", "completeness": "complete", "id": "xcsh-docs:resources:address_allocator:fundamentals", "parent_id": null, "path": "docs/resources/address_allocator.md", "provider_name": "address_allocator", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/address_allocator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_address_allocator for xcsh_address_allocator.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["address_allocatorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

- [Property reference](../guides/resources--address_allocator--reference.md)
- [Examples](../guides/resources--address_allocator--examples.md)
- [Import](../guides/resources--address_allocator--import.md)
- [Timeouts](../guides/resources--address_allocator--timeouts.md)
