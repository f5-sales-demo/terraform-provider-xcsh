---
page_title: "xcsh_virtual_network examples"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_virtual_network examples."
---

# xcsh_virtual_network examples

<a id="canonical-aa858708777a299dd1c665fa01fe06408adadb5fbc8ee42c3fe6ede1ba508ff7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2fd3f6aa9abddb8464537aaf5e033e3424f862a1f296cebeafe240f1a84ebed5"></a>

## Examples — Examples / 907e4b4dc3a1 / 2

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)
- Examples

<a id="canonical-4a09a3cb5e0e350e3f1373a1d0885032c0af17bfdbb8f3b1770941b9ccce87f6"></a>

## Complete configurations — Examples / 907e4b4dc3a1 / 3

- [Data source](data-sources--virtual_network--examples--group-001.md#canonical-72e97f01cd4cb926beb19bbf1dc41c3251e45f033333072d3b31adeab24be36c): valid configuration.

<a id="canonical-96e654fed0c15af533a4c12b7081607be437b580233a09e9021e90ff9451b0d7"></a>

## Next pages — Examples / 907e4b4dc3a1 / 4

- [Data source](data-sources--virtual_network--examples--group-001.md#canonical-72e97f01cd4cb926beb19bbf1dc41c3251e45f033333072d3b31adeab24be36c)
- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)

<a id="canonical-72e97f01cd4cb926beb19bbf1dc41c3251e45f033333072d3b31adeab24be36c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06c748bf5a49d8222b311ace0329030ff8152f1249949f6a04035d8609a848d6"></a>

## Data source — Data source / 06e4c9679e51 / 2

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)
- [Examples](data-sources--virtual_network--examples--group-001.md#canonical-aa858708777a299dd1c665fa01fe06408adadb5fbc8ee42c3fe6ede1ba508ff7)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_virtual_network/data-source.tf`; digest `sha256:82f2d94c3d011a8d200fcd015e8a3406aea69e2a77a8f9d9eea9796b4b19e3bd`.

```terraform
# VirtualNetwork Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualNetwork by name
data "xcsh_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"
}

output "virtual_network_id" {
  value = data.xcsh_virtual_network.example.id
}
```

<a id="canonical-805d4b88ebf8ceff7f8cfa92902af375d41252cd041827734ecb9b664d811a4b"></a>

## Next pages — Data source / 06e4c9679e51 / 3

- [Examples](data-sources--virtual_network--examples--group-001.md#canonical-aa858708777a299dd1c665fa01fe06408adadb5fbc8ee42c3fe6ede1ba508ff7)
- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)
