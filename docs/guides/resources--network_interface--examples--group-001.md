---
page_title: "xcsh_network_interface examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface examples."
---

# xcsh_network_interface examples

<a id="canonical-deb23a5ecdf3dfc02a1fdea55794fc92d32fd586e58dcc71f144a1a5b39f9e4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25ba53ef8968981f13fd4ea4481d44fec73b7c2b1fbc2ca12d803e30305f6c00"></a>

## Examples — Examples / 45d4ae324763 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- Examples

<a id="canonical-7b2f0d0da8b9c800184d1c02d893d855d9e0e16d309f04636ba31c82ed3e0ecc"></a>

## Complete configurations — Examples / 45d4ae324763 / 3

- [Resource](resources--network_interface--examples--group-001.md#canonical-83c3e3453485b2d18170899c1f41c2ed1a03618c270f0e80f7a7af6e33810389): valid configuration.

<a id="canonical-51a1c47a792e0d8ff9ac931913dd7fda10e4944a6d9b393b141529f3bce90bf3"></a>

## Next pages — Examples / 45d4ae324763 / 4

- [Resource](resources--network_interface--examples--group-001.md#canonical-83c3e3453485b2d18170899c1f41c2ed1a03618c270f0e80f7a7af6e33810389)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-83c3e3453485b2d18170899c1f41c2ed1a03618c270f0e80f7a7af6e33810389"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6cd16f9b3be38da8bc8ab85d88b9d16bacdcb3e91008cc74f3bcbff42e039e54"></a>

## Resource — Resource / 221fa5db33ef / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Examples](resources--network_interface--examples--group-001.md#canonical-deb23a5ecdf3dfc02a1fdea55794fc92d32fd586e58dcc71f144a1a5b39f9e4d)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_interface/resource.tf`; digest `sha256:7844384dc87eb9b41829fe1bcb33636a61a1469e9b6c183c89c2bb605fd224c5`.

```terraform
# NetworkInterface Resource Example
# Manages a Network Interface resource in F5 Distributed Cloud for network interface represents configuration of a network device.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkInterface configuration
resource "xcsh_network_interface" "example" {
  name      = "example-network-interface"
  namespace = "staging"
}
```

<a id="canonical-fa15f91d6685d0d729501961038ad939a2787e94afe9674fea95afd2ef048fd4"></a>

## Next pages — Resource / 221fa5db33ef / 3

- [Examples](resources--network_interface--examples--group-001.md#canonical-deb23a5ecdf3dfc02a1fdea55794fc92d32fd586e58dcc71f144a1a5b39f9e4d)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
