---
page_title: "xcsh_virtual_network landing"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_virtual_network landing."
---

# xcsh_virtual_network landing

<a id="canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-faba0277830c910bb581c60bf3f42c01fbae52f0b3b53e2ab158852ed55975cb"></a>

## xcsh_virtual_network — xcsh_virtual_network / ed4c72be6b34 / 2

Breadcrumbs:

- xcsh_virtual_network

Manages virtual network in given namespace in F5 Distributed Cloud.

<a id="canonical-b5b4c787ac98041a865dcc8b11f35fdbc56307c2bbfbe6463bd6cfda966b8683"></a>

## Prerequisites — xcsh_virtual_network / ed4c72be6b34 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `network_connector`.

- network_connector: Connect to external networks

<a id="canonical-438f3ad15b38a48ef5b8901462103ea5ebfec7984060f986879109fbea485b5e"></a>

## Minimal configuration — xcsh_virtual_network / ed4c72be6b34 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-0ac5fff443668b3a23032ff355ae53da138a3934a1c03964935b92113f7b92db"></a>

## Root configuration — xcsh_virtual_network / ed4c72be6b34 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-58729813761853841f7bf72622c5e7511ea47b7c7dd4716de44995fc004b1ffb"></a>

## Next pages — xcsh_virtual_network / ed4c72be6b34 / 6

- [Property reference](../guides/resources--virtual_network--reference--group-001.md#canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf)
- [Examples](../guides/resources--virtual_network--examples--group-001.md#canonical-1136a78de868fce47dd66a1ef294d0b019c86f62b789f3df9e913d18ecb8cca1)
- [Import](../guides/resources--virtual_network--lifecycle--group-001.md#canonical-b77c9c94c6bab3e06ab6c7ae8d3916bb9df610d4412aec376cb1a98908fdd0d4)
- [Timeouts](../guides/resources--virtual_network--lifecycle--group-001.md#canonical-15d1f867c89c5630f849d88c02050fab6153880f566ef63a6babb8d7c56f904c)
