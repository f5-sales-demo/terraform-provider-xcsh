---
page_title: "xcsh_address_allocator landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_address_allocator landing."
---

# xcsh_address_allocator landing

<a id="canonical-b37f5189fa71c353983be346c78099ee10d1935bec31c657e73fc56ee3cc9d4b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8ff8988de2dde6f541f5b0852809dacd7afede4a36e6f8eacaf73350a3e0871"></a>

## xcsh_address_allocator — xcsh_address_allocator / 702b6fce20fb / 2

Breadcrumbs:

- xcsh_address_allocator

Manages Address Allocator will create an address allocator object in 'system' namespace of the user
in F5 Distributed Cloud.

<a id="canonical-8dc3b03e86cd8ecef1645fc253209ca8a2110cf236b5d1c5afbc7ae6e3154271"></a>

## Prerequisites — xcsh_address_allocator / 702b6fce20fb / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-70818ce53b30363f71cc64485d38ee811f90731bd820c71be19dcef0c6b9c172"></a>

## Minimal configuration — xcsh_address_allocator / 702b6fce20fb / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-7d9925d6b11d39a8fc60dfea6d1ca4cba3121293fd8d2de8bdef953a92a5895b"></a>

## Root configuration — xcsh_address_allocator / 702b6fce20fb / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-7fdebc0ad9a7c970ec594e2c94d7b71dcbadc76cce018001dd438a3789535362"></a>

## Next pages — xcsh_address_allocator / 702b6fce20fb / 6

- [Property reference](../guides/data-sources--address_allocator--reference--group-001.md#canonical-06d6ef2a4ebffe9fba2c3fad1b3c86d8113e071f5d1178e3a4362816a6103b09)
- [Examples](../guides/data-sources--address_allocator--examples--group-001.md#canonical-761000534cef216afc668fad5da8ef473c794b9038d0e472d3d1e231d57a3618)
