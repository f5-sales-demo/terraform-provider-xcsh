---
page_title: "xcsh_bgp landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp landing."
---

# xcsh_bgp landing

<a id="canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65f0c5ac6d6a3853466b00b51a5ca46e4124d37d45bf9bf51f320571e0faddfb"></a>

## xcsh_bgp — xcsh_bgp / a7858fe9b415 / 2

Breadcrumbs:

- xcsh_bgp

Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with
external bgp servers. it is created by users in system namespace. configuration.

<a id="canonical-0c8a25398fcd4dd61e930cd2b1666c8c66b77fcb28533d2ff2f967dac7e95f53"></a>

## Prerequisites — xcsh_bgp / a7858fe9b415 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-ed2017665ecbfc8d8b8e0c1bb207737edb1d5bdf9cb5d8fcba2972ec076e75d6"></a>

## Minimal configuration — xcsh_bgp / a7858fe9b415 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGP Resource Example
# Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with external bgp servers.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BGP configuration
resource "xcsh_bgp" "example" {
  name      = "example-bgp"
  namespace = "staging"
}
```

<a id="canonical-67afb1afccc316dfba93d4ded4df18d860146ce60b3d06b65f2495e791ae83ef"></a>

## Root configuration — xcsh_bgp / a7858fe9b415 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-686b85466a8894113c3f4a222b737990c21bb7a0be71edbb6d1288d2cefcd4e6"></a>

## Next pages — xcsh_bgp / a7858fe9b415 / 6

- [Property reference](../guides/resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [Examples](../guides/resources--bgp--examples--group-001.md#canonical-0b5fe495434f4ae93528ca40a2c382856380185f9580f1f16c5b072a7d6a2bca)
- [Import](../guides/resources--bgp--lifecycle--group-001.md#canonical-1078d0e498358790f7493c65620697323658a43741ab40880d6be54d64c31acc)
- [Timeouts](../guides/resources--bgp--lifecycle--group-001.md#canonical-cf0f94b8f0be554446522de77d0c4d87aa513342e488240aa49c6eae45d697f2)
