---
page_title: "xcsh_srv6_network_slice landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_srv6_network_slice landing."
---

# xcsh_srv6_network_slice landing

<a id="canonical-01eb91def3bc993a7ab14854fcd32915db33e69c6c189dc0599a4ae5db267c6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-782baef6572ae7e35811ae1c67168e31c7d2cc015b365659acdbcf6622838ee7"></a>

## xcsh_srv6_network_slice — xcsh_srv6_network_slice / 7057f7b40947 / 2

Breadcrumbs:

- xcsh_srv6_network_slice

Manages srv6\_network\_slice creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-b0836d7705045e7434ed76fbb9cc6517b3423081f0e2b73fc4bf9f127be8e080"></a>

## Prerequisites — xcsh_srv6_network_slice / 7057f7b40947 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3c81f6e5004778850ba21a6b833a6ff11de6535288732305398732f0661f6862"></a>

## Minimal configuration — xcsh_srv6_network_slice / 7057f7b40947 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Srv6NetworkSlice Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Srv6NetworkSlice by name
data "xcsh_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"
}

output "srv6_network_slice_id" {
  value = data.xcsh_srv6_network_slice.example.id
}
```

<a id="canonical-d04fe8eedef32340506833ce5784294d2538c7d5d2bcda754f0652a6ffed879a"></a>

## Root configuration — xcsh_srv6_network_slice / 7057f7b40947 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-5141db31c468f7013b2d9f56a8e420628bca4b6c5c4505463da1760ccd2fb51d"></a>

## Next pages — xcsh_srv6_network_slice / 7057f7b40947 / 6

- [Property reference](../guides/data-sources--srv6_network_slice--reference--group-001.md#canonical-ceae08b891b45a658e6ab9b9d513560c18aeedd96150bf1affbd3770714da82d)
- [Examples](../guides/data-sources--srv6_network_slice--examples--group-001.md#canonical-b2b705b70a59b78ad13e1fbdde0f0af4ac8ab79472c63dc8818b8fb18de8c11a)
