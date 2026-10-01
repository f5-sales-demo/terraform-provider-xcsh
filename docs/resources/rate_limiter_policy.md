---
page_title: "xcsh_rate_limiter_policy landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter_policy landing."
---

# xcsh_rate_limiter_policy landing

<a id="canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-066f9f04ab41fdf94a771923c35b09f053c0699b5dde2cad30a192379d59babd"></a>

## xcsh_rate_limiter_policy — xcsh_rate_limiter_policy / c0b66b505e5b / 2

Breadcrumbs:

- xcsh_rate_limiter_policy

Manages a Rate Limiter Policy resource in F5 Distributed Cloud for rate limiter policy create
specification. configuration.

<a id="canonical-81ac1a536b822fc5dc7a3975ab0545ab13db82d4c307c6507f82959f7df161ca"></a>

## Prerequisites — xcsh_rate_limiter_policy / c0b66b505e5b / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-b7dac72dfd72d202cfa5aa1716472c617d5d70d92d28f712e682c1574e0f3677"></a>

## Minimal configuration — xcsh_rate_limiter_policy / c0b66b505e5b / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# RateLimiterPolicy Resource Example
# Manages a Rate Limiter Policy resource in F5 Distributed Cloud for rate limiter policy create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic RateLimiterPolicy configuration
resource "xcsh_rate_limiter_policy" "example" {
  name      = "example-rate-limiter-policy"
  namespace = "staging"
}
```

<a id="canonical-b9b55c4853b796c873fe832387714114ab52d754738911d527ceba08f11f0608"></a>

## Root configuration — xcsh_rate_limiter_policy / c0b66b505e5b / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-8efbfa99d148e1251820753da72e69d306c76b86378a528d9586537696c484f7"></a>

## Next pages — xcsh_rate_limiter_policy / c0b66b505e5b / 6

- [Property reference](../guides/resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [Examples](../guides/resources--rate_limiter_policy--examples--group-001.md#canonical-2741be1581f1f9d73872c6a2366d1a89cb9fa023b50acd8e9f1053d77c69673b)
- [Import](../guides/resources--rate_limiter_policy--lifecycle--group-001.md#canonical-ca33a0a19ffcc56ee6a120b6441073b417b696d1d599ad4ab7d963f12bb6d426)
- [Timeouts](../guides/resources--rate_limiter_policy--lifecycle--group-001.md#canonical-c4af72ed78de66e325de99bbc007587a8f18cf22f0b9a9df5b5e7e4fca236ef2)
