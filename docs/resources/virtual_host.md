---
page_title: "xcsh_virtual_host landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host landing."
---

# xcsh_virtual_host landing

<a id="canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08bae49741579226b8de02ebe189923a9e82fe8701affc9a30e36f9c383aab1a"></a>

## xcsh_virtual_host — xcsh_virtual_host / 2b79f807777d / 2

Breadcrumbs:

- xcsh_virtual_host

Manages virtual host in a given namespace in F5 Distributed Cloud.

<a id="canonical-e58e256fc530ce819b044eeff433d27569b6b7a2356c7d9cc34961b1fdca0fc4"></a>

## Prerequisites — xcsh_virtual_host / 2b79f807777d / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-b6233299dd5294cac474df008746ba2fb41ad2d3905521e229b5f8d40293c126"></a>

## Minimal configuration — xcsh_virtual_host / 2b79f807777d / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualHost Resource Example
# Manages virtual host in a given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualHost configuration
resource "xcsh_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "staging"
}
```

<a id="canonical-1368bf05ba6062814432e9bc723bfd53265b8beab11919869bb6967a35f35f7e"></a>

## Root configuration — xcsh_virtual_host / 2b79f807777d / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-31064daf90086da4a46eea24b0a503ab253b7044f8b31fbace2706483bf50ebd"></a>

## Next pages — xcsh_virtual_host / 2b79f807777d / 6

- [Property reference](../guides/resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [Examples](../guides/resources--virtual_host--examples--group-001.md#canonical-6478e4ef8c92ee0e5eeebc8e11c97ab12ee08c898ee8603da8ef3f11720a52eb)
- [Import](../guides/resources--virtual_host--lifecycle--group-001.md#canonical-c46b530bf177650e4368bfadcda25ab5045d3743d483cbcd4b2b1333f0adb531)
- [Timeouts](../guides/resources--virtual_host--lifecycle--group-001.md#canonical-9c1c11b4fe1638102f14ce07ac3a0ecf892e9ee349802099ce01b7a934a5d344)
