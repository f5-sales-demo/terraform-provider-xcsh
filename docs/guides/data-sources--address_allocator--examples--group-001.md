---
page_title: "xcsh_address_allocator examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_address_allocator examples."
---

# xcsh_address_allocator examples

<a id="canonical-761000534cef216afc668fad5da8ef473c794b9038d0e472d3d1e231d57a3618"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f20ea7babc1239b695264edcda8e76c221573a04aea10537e8d547da1ce67a88"></a>

## Examples — Examples / 0f0f393284c2 / 2

Breadcrumbs:

- [xcsh_address_allocator](../data-sources/address_allocator.md#canonical-b37f5189fa71c353983be346c78099ee10d1935bec31c657e73fc56ee3cc9d4b)
- Examples

<a id="canonical-89078e71135bdcf2a72b4c06bfea91fcfe984cd4930b46ee4628f9722c992bb8"></a>

## Complete configurations — Examples / 0f0f393284c2 / 3

- [Data source](data-sources--address_allocator--examples--group-001.md#canonical-e4707975080aae2f33b0b4af39d15eb1d2732629bebb3d0254203934eec6ae81): valid configuration.

<a id="canonical-f1eb9d48245a89691ea56b43d35a05e92cfb018c76ce05737fc9556cbadf26fe"></a>

## Next pages — Examples / 0f0f393284c2 / 4

- [Data source](data-sources--address_allocator--examples--group-001.md#canonical-e4707975080aae2f33b0b4af39d15eb1d2732629bebb3d0254203934eec6ae81)
- [xcsh_address_allocator](../data-sources/address_allocator.md#canonical-b37f5189fa71c353983be346c78099ee10d1935bec31c657e73fc56ee3cc9d4b)

<a id="canonical-e4707975080aae2f33b0b4af39d15eb1d2732629bebb3d0254203934eec6ae81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b82ad682f3760a1227a9f878bac00b1756e810f6cc500393d98c6cc15176e30"></a>

## Data source — Data source / 7e11b861253e / 2

Breadcrumbs:

- [xcsh_address_allocator](../data-sources/address_allocator.md#canonical-b37f5189fa71c353983be346c78099ee10d1935bec31c657e73fc56ee3cc9d4b)
- [Examples](data-sources--address_allocator--examples--group-001.md#canonical-761000534cef216afc668fad5da8ef473c794b9038d0e472d3d1e231d57a3618)
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

<a id="canonical-00291add9730bd1faab3523ce1b70f381224863742320b53a6f58357af5589be"></a>

## Next pages — Data source / 7e11b861253e / 3

- [Examples](data-sources--address_allocator--examples--group-001.md#canonical-761000534cef216afc668fad5da8ef473c794b9038d0e472d3d1e231d57a3618)
- [xcsh_address_allocator](../data-sources/address_allocator.md#canonical-b37f5189fa71c353983be346c78099ee10d1935bec31c657e73fc56ee3cc9d4b)
