---
page_title: "xcsh_namespace landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_namespace landing."
---

# xcsh_namespace landing

<a id="canonical-5d4d79984896dc184c2d678bc72c561a372fa7660188fc55cde1c40e575c98b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ada6c8c475fffe940e73cb147c55e7895c7b567b5539fdb7f71a096bfa11bec7"></a>

## xcsh_namespace — xcsh_namespace / e5720604b29a / 2

Breadcrumbs:

- xcsh_namespace

Manages new namespace. Name of the object is name of the name space in F5 Distributed Cloud.

<a id="canonical-10417fa16bdc05ace616ca764839587d5dd51f02c0b99a17fda9b5ead7d43d4b"></a>

## Prerequisites — xcsh_namespace / e5720604b29a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-b461753d214bc2fec11961259e0f32ead2103ca5f13667603aa0bd07603247b5"></a>

## Minimal configuration — xcsh_namespace / e5720604b29a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Namespace Resource Example
# Manages new namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Credentials are supplied externally.
provider "xcsh" {}

# Basic Namespace configuration
resource "xcsh_namespace" "this" {
  name = "example-namespace"
}
```

<a id="canonical-fd949e3acc376e5ab01099671308b86a32cb8749947bfc9bc6dfbc9aca6cbfeb"></a>

## Root configuration — xcsh_namespace / e5720604b29a / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-0a61fb67e57d8303ca31a83b82d34cb5d2aa4988c9fc2a9cedd0669a51e4e088"></a>

## Next pages — xcsh_namespace / e5720604b29a / 6

- [Property reference](../guides/resources--namespace--reference--group-001.md#canonical-09c5d1f37f639be8d72ab53f1082e4c1d53428b0ed5d288fca41f55d74dce8ac)
- [Examples](../guides/resources--namespace--examples--group-001.md#canonical-1b8bc34f331033b05402414f3e116c5e5532af0c28d25550b4e20bc0b278be2c)
- [Import](../guides/resources--namespace--lifecycle--group-001.md#canonical-380303388909b8d633df53f36f3e50f7426a3e15bce59894a553c6c5c2465039)
- [Timeouts](../guides/resources--namespace--lifecycle--group-001.md#canonical-2ba63e0380c2dbdbee97f5e55661eaff6ba003c4e9094ce413e306162ac09d51)
