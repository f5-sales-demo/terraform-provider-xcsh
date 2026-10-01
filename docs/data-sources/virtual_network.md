---
page_title: "xcsh_virtual_network landing"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_virtual_network landing."
---

# xcsh_virtual_network landing

<a id="canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a1cbe4ffaf33181399f1a918b4aa64c3cfd8d3e2a2ed2014ce286f34d43be14"></a>

## xcsh_virtual_network — xcsh_virtual_network / e753cd0e252d / 2

Breadcrumbs:

- xcsh_virtual_network

Manages virtual network in given namespace in F5 Distributed Cloud.

<a id="canonical-a4483751c196ada9bfbf52c22516f5ac1df57dd03ac88a7fb617636bedb0f8c8"></a>

## Prerequisites — xcsh_virtual_network / e753cd0e252d / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `network_connector`.

- network_connector: Connect to external networks

<a id="canonical-3275144567b94ec0c00e333172c4e86c78c56f83e560ec71963540c61c6e5ed6"></a>

## Minimal configuration — xcsh_virtual_network / e753cd0e252d / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-7db938fc4743c6f114e08a74a4292a721eebe43d93628a305be0eec22ecff589"></a>

## Root configuration — xcsh_virtual_network / e753cd0e252d / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-2060129f3714adfb8591db24962711acae76ba987b40d595bcde7fbae4715c5d"></a>

## Next pages — xcsh_virtual_network / e753cd0e252d / 6

- [Property reference](../guides/data-sources--virtual_network--reference--group-001.md#canonical-9c449ec5c2f1fe76d44c1c939fe391b4c4e3c20e4e3572d7fc09ce3420ccce57)
- [Examples](../guides/data-sources--virtual_network--examples--group-001.md#canonical-aa858708777a299dd1c665fa01fe06408adadb5fbc8ee42c3fe6ede1ba508ff7)
