---
page_title: "xcsh_bgp landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp landing."
---

# xcsh_bgp landing

<a id="canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18dbcb52813cd2171a56d6113af7963a775465cb6b0eafe92bfa5bfa8ad5536b"></a>

## xcsh_bgp — xcsh_bgp / f3763090c904 / 2

Breadcrumbs:

- xcsh_bgp

Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with
external bgp servers. it is created by users in system namespace. configuration.

<a id="canonical-08c71b1bab78493ed003262ab51fa85b1d9fdb3f79a3f41f27c64b7babefe1a4"></a>

## Prerequisites — xcsh_bgp / f3763090c904 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-6dd2d0bb4b7310ad34468174d2edc835df4739e2df8db96260d1050ccfc1468e"></a>

## Minimal configuration — xcsh_bgp / f3763090c904 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGP by name
data "xcsh_bgp" "example" {
  name      = "example-bgp"
  namespace = "staging"
}

output "bgp_id" {
  value = data.xcsh_bgp.example.id
}
```

<a id="canonical-59126f589a54c1c1c7a5898ff8508144cc02190f6e1ae6182e6d2d7557cea5cf"></a>

## Root configuration — xcsh_bgp / f3763090c904 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-dc65bad84be7da4a91e303a2508c3d65b7b42dd09602890c0d601c779e49a17e"></a>

## Next pages — xcsh_bgp / f3763090c904 / 6

- [Property reference](../guides/data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [Examples](../guides/data-sources--bgp--examples--group-001.md#canonical-05fe067b621ea11533193eb95e9bf3cd816189ec172eb71a345b4eb582563043)
