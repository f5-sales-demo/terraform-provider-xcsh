---
page_title: "xcsh_network_connector landing"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_network_connector landing."
---

# xcsh_network_connector landing

<a id="canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c872e68322ce1b859e1facfb4266c2c171e9b8cc15d3accda8784bd4bff90e4d"></a>

## xcsh_network_connector — xcsh_network_connector / 21cee089d396 / 2

Breadcrumbs:

- xcsh_network_connector

Manages a Network Connector resource in F5 Distributed Cloud for network connector is created by
users in system namespace. configuration.

<a id="canonical-d4f719351158e916e905cbd659c72e428ced335d7842dc12ba3d1540b02e5e07"></a>

## Prerequisites — xcsh_network_connector / 21cee089d396 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `virtual_network`.

- virtual_network: Network to connect

<a id="canonical-b5df403e1b02801dd3c4817a7d4c26cc422a9ee1a8327a386c0d3f29ddab5962"></a>

## Minimal configuration — xcsh_network_connector / 21cee089d396 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkConnector Resource Example
# Manages a Network Connector resource in F5 Distributed Cloud for network connector is created by users in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkConnector configuration
resource "xcsh_network_connector" "example" {
  name      = "example-network-connector"
  namespace = "staging"
}
```

<a id="canonical-dbbe2425b661ee799ecc65f99c2646470567ae2e6cdf1e3e635b90626764d4f2"></a>

## Root configuration — xcsh_network_connector / 21cee089d396 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0bb130aabb7f75cce809a4d8a3e80f2e61c144438c1d00779bab3f0e39472e48"></a>

## Next pages — xcsh_network_connector / 21cee089d396 / 6

- [Property reference](../guides/resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [Examples](../guides/resources--network_connector--examples--group-001.md#canonical-da5d1aadf602caf35fa642617b8652781dcdfdb9c4b01054e2e6ee1d216efbff)
- [Import](../guides/resources--network_connector--lifecycle--group-001.md#canonical-99266101de64536830d8356e1a4262545ec691de5a88830c5007df5b7700e727)
- [Timeouts](../guides/resources--network_connector--lifecycle--group-001.md#canonical-c6eaac6ffbfdc7cfe36e41f8d49435aa9468b5a55d119910f009f70a8135e916)
