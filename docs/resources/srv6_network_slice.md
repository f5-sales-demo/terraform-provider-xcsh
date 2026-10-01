---
page_title: "xcsh_srv6_network_slice landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_srv6_network_slice landing."
---

# xcsh_srv6_network_slice landing

<a id="canonical-431435cc9f00e58094c1dd963b8aa9cec4251ea8dcbf3313b046ca1921402d3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e545bb37ad7778107adbafff3f229d466b78dc992c0cd26e39405ef676e950d"></a>

## xcsh_srv6_network_slice — xcsh_srv6_network_slice / b42b560f8b58 / 2

Breadcrumbs:

- xcsh_srv6_network_slice

Manages srv6\_network\_slice creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-54893bd2ce639a0132ade1ba3da9a75479c12ec6bfa5e17eb28097e5b208bdd4"></a>

## Prerequisites — xcsh_srv6_network_slice / b42b560f8b58 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-47ebac27ac9e0312608f0f27c3fed456c5164f16e81ec9df3ef43dbd77c09878"></a>

## Minimal configuration — xcsh_srv6_network_slice / b42b560f8b58 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Srv6NetworkSlice Resource Example
# Manages srv6_network_slice creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Srv6NetworkSlice configuration
resource "xcsh_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"

  sid_prefixes = ["example-value"]
}
```

<a id="canonical-0a1645ef9f89ae5df1ec2f370f3552d4e4e3514f192fcdbb87dcb4f6e511952c"></a>

## Root configuration — xcsh_srv6_network_slice / b42b560f8b58 / 5

Required root properties: `name`, `sid_prefixes`. Full root flags and choices appear in the property reference.

<a id="canonical-169b11399f55997489f0ffee3171c4520c67f795059416a63b20ec3357bfd839"></a>

## Next pages — xcsh_srv6_network_slice / b42b560f8b58 / 6

- [Property reference](../guides/resources--srv6_network_slice--reference--group-001.md#canonical-a9921f7005bdd8167184c5a6a56fca49fff6fc9a386f4d22e032d049c63fc870)
- [Examples](../guides/resources--srv6_network_slice--examples--group-001.md#canonical-228c5eca67933ce19c2353ffab9baa21c19c683b96c643cd0f23c62c39acc9e1)
- [Import](../guides/resources--srv6_network_slice--lifecycle--group-001.md#canonical-49d2c70de7063a333d4931b3a921f8d08214ee237ad5e3715fbc6fde2768f201)
- [Timeouts](../guides/resources--srv6_network_slice--lifecycle--group-001.md#canonical-6b5170b64d10d3e92f1a2b241f2c94c933794d119256b0a9a5055a34daedf100)
