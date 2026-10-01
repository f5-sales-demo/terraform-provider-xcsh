---
page_title: "xcsh_network_connector examples"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_network_connector examples."
---

# xcsh_network_connector examples

<a id="canonical-da5d1aadf602caf35fa642617b8652781dcdfdb9c4b01054e2e6ee1d216efbff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c1d12a9e7b834829285ca53d3d472d030c0064c68b3347bf916ff72f2c1265e"></a>

## Examples — Examples / 827942b9d907 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- Examples

<a id="canonical-e1cf0cb0687320a7a0bf167a5e0efe36fef89a2718935e84f115e457f5420112"></a>

## Complete configurations — Examples / 827942b9d907 / 3

- [Resource](resources--network_connector--examples--group-001.md#canonical-4c91713cd5651d325c973b6ecc97e5352e571d184fb418d79f0487e7f8109309): valid configuration.

<a id="canonical-08772fc31412274c1c735b043b99dc10c534e897341f7db0020e68d648c87e56"></a>

## Next pages — Examples / 827942b9d907 / 4

- [Resource](resources--network_connector--examples--group-001.md#canonical-4c91713cd5651d325c973b6ecc97e5352e571d184fb418d79f0487e7f8109309)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-4c91713cd5651d325c973b6ecc97e5352e571d184fb418d79f0487e7f8109309"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f734af1a9bf089b0ef49f58bba1aeed58414d051092813d3a9f789f8512dff85"></a>

## Resource — Resource / 09534b58cb0c / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Examples](resources--network_connector--examples--group-001.md#canonical-da5d1aadf602caf35fa642617b8652781dcdfdb9c4b01054e2e6ee1d216efbff)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_connector/resource.tf`; digest `sha256:79950822068b0bbf291a51dad0cff745ae6eacc14d286990d529310031a03379`.

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

<a id="canonical-0ea70a4ed06aee85200de62dd1348b21c7120066c1c5a4d6adc7d5195469b268"></a>

## Next pages — Resource / 09534b58cb0c / 3

- [Examples](resources--network_connector--examples--group-001.md#canonical-da5d1aadf602caf35fa642617b8652781dcdfdb9c4b01054e2e6ee1d216efbff)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
