---
page_title: "xcsh_address_allocator landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_address_allocator landing."
---

# xcsh_address_allocator landing

<a id="canonical-a90eda9d0678adc17d442e30a4391e1edf799c2a7419e020376896c3e91a3aba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37696de01c36c1a28466a3e93086a67f6f6dd47db054f887e1cba15ba15000f0"></a>

## xcsh_address_allocator — xcsh_address_allocator / 21a0bd8fcd61 / 2

Breadcrumbs:

- xcsh_address_allocator

Manages Address Allocator will create an address allocator object in 'system' namespace of the user
in F5 Distributed Cloud.

<a id="canonical-dcb2631afd864373ff038ed947e752c2bd28fab8bda9707831dff9fbb328be85"></a>

## Prerequisites — xcsh_address_allocator / 21a0bd8fcd61 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-84e64d115b843330ca2b5de47298b36576af636928feb8474bb1ce68c3ae2625"></a>

## Minimal configuration — xcsh_address_allocator / 21a0bd8fcd61 / 4

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

<a id="canonical-4b7f696125b68712c676cb962790092945ccd6547956557cddc1d9e33c39a354"></a>

## Root configuration — xcsh_address_allocator / 21a0bd8fcd61 / 5

Required root properties: `address_pool`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-99ef6f8f753c0f8449598a8be31de26fba9a5eb6a712a9d704cfc64b2bae297a"></a>

## Next pages — xcsh_address_allocator / 21a0bd8fcd61 / 6

- [Property reference](../guides/resources--address_allocator--reference--group-001.md#canonical-6c53b0269d6a197e4c6beae947a7d28d842ea2701871f0c541ede9d544757f27)
- [Examples](../guides/resources--address_allocator--examples--group-001.md#canonical-9abf27a8eeb302922bcfca6b62f937985c07ee9f303d8f0174b830c873cb1952)
- [Import](../guides/resources--address_allocator--lifecycle--group-001.md#canonical-7e1fa737f98bb2ea13e24c2c1e1453a19863d6210d8db84bae3d54b88865c7fe)
- [Timeouts](../guides/resources--address_allocator--lifecycle--group-001.md#canonical-60eb9705c00e730ea1ecfe9263165fab29f22f76a32e5eb5f5caa8596af75662)
