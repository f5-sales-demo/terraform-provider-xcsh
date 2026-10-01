---
page_title: "xcsh_virtual_network examples"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_virtual_network examples."
---

# xcsh_virtual_network examples

<a id="canonical-1136a78de868fce47dd66a1ef294d0b019c86f62b789f3df9e913d18ecb8cca1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20e2c56673a1466d4c4c2bfd9eaa9de7c5732b00b0c8398205ee0b3c44cd9f7b"></a>

## Examples — Examples / 1e466ebb4dc9 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)
- Examples

<a id="canonical-6dc420c1292fde6d7be7ecc03f89465dd375bc4ad37ed09106894c4bac55b7f2"></a>

## Complete configurations — Examples / 1e466ebb4dc9 / 3

- [Resource](resources--virtual_network--examples--group-001.md#canonical-7a6d8dacde52e71832b105ca3dd4c873c291fa29fe6e7237d858fc428831902e): valid configuration.

<a id="canonical-7ebf084a85be0df1269cd5bb239dd1317e342bfcf20688b62b23912ea7145a92"></a>

## Next pages — Examples / 1e466ebb4dc9 / 4

- [Resource](resources--virtual_network--examples--group-001.md#canonical-7a6d8dacde52e71832b105ca3dd4c873c291fa29fe6e7237d858fc428831902e)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)

<a id="canonical-7a6d8dacde52e71832b105ca3dd4c873c291fa29fe6e7237d858fc428831902e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17426555440c7181bb3a8061e4ff01157c8e88c41c7a5f602674647cfd5752f1"></a>

## Resource — Resource / 423457a038b6 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)
- [Examples](resources--virtual_network--examples--group-001.md#canonical-1136a78de868fce47dd66a1ef294d0b019c86f62b789f3df9e913d18ecb8cca1)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_virtual_network/resource.tf`; digest `sha256:24ff8ad3e56ba4ece18f3beaaabdcc55805f2e25c74acd17556111d54902f473`.

```terraform
# VirtualNetwork Resource Example
# Manages virtual network in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualNetwork configuration
resource "xcsh_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"
}
```

<a id="canonical-f6a100793542230fa90a3bba3198ece06f0cd0c4313baed6ec2fb8195052547a"></a>

## Next pages — Resource / 423457a038b6 / 3

- [Examples](resources--virtual_network--examples--group-001.md#canonical-1136a78de868fce47dd66a1ef294d0b019c86f62b789f3df9e913d18ecb8cca1)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)
