---
page_title: "xcsh_address_allocator examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_address_allocator examples."
---

# xcsh_address_allocator examples

<a id="canonical-9abf27a8eeb302922bcfca6b62f937985c07ee9f303d8f0174b830c873cb1952"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b073a226426efe86379f71b56f001f8709da576494a9f71ef02b8452df6cc863"></a>

## Examples — Examples / 5d4ba36cb141 / 2

Breadcrumbs:

- [xcsh_address_allocator](../resources/address_allocator.md#canonical-a90eda9d0678adc17d442e30a4391e1edf799c2a7419e020376896c3e91a3aba)
- Examples

<a id="canonical-38755d46a97442d8a661db63e09061e4afe1fb9bf53ea35fb5d4a437b6e6059f"></a>

## Complete configurations — Examples / 5d4ba36cb141 / 3

- [Resource](resources--address_allocator--examples--group-001.md#canonical-a084b9f43cd016fd8f0745228b42908777ef6958026905d670a516e12c933345): valid configuration.

<a id="canonical-b50108d32038d6df6dcb86e3343325d1f1768af3fc14c53736c334e66f8d45e6"></a>

## Next pages — Examples / 5d4ba36cb141 / 4

- [Resource](resources--address_allocator--examples--group-001.md#canonical-a084b9f43cd016fd8f0745228b42908777ef6958026905d670a516e12c933345)
- [xcsh_address_allocator](../resources/address_allocator.md#canonical-a90eda9d0678adc17d442e30a4391e1edf799c2a7419e020376896c3e91a3aba)

<a id="canonical-a084b9f43cd016fd8f0745228b42908777ef6958026905d670a516e12c933345"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9efc601364377e5abca337131a8280da77ff9f87932eb4e8e50959fb5452444"></a>

## Resource — Resource / 5504769cc71b / 2

Breadcrumbs:

- [xcsh_address_allocator](../resources/address_allocator.md#canonical-a90eda9d0678adc17d442e30a4391e1edf799c2a7419e020376896c3e91a3aba)
- [Examples](resources--address_allocator--examples--group-001.md#canonical-9abf27a8eeb302922bcfca6b62f937985c07ee9f303d8f0174b830c873cb1952)
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

<a id="canonical-861d9e41491c2957fa124ef52ff0af3018003f3419700be079a07d02f9c226ea"></a>

## Next pages — Resource / 5504769cc71b / 3

- [Examples](resources--address_allocator--examples--group-001.md#canonical-9abf27a8eeb302922bcfca6b62f937985c07ee9f303d8f0174b830c873cb1952)
- [xcsh_address_allocator](../resources/address_allocator.md#canonical-a90eda9d0678adc17d442e30a4391e1edf799c2a7419e020376896c3e91a3aba)
