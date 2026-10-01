---
page_title: "xcsh_cdn_cache_rule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_cache_rule landing."
---

# xcsh_cdn_cache_rule landing

<a id="canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15a0f477121930e1a07a413ccc77e9e8f77682d3b9c31248e7b251ce2712d164"></a>

## xcsh_cdn_cache_rule — xcsh_cdn_cache_rule / da3a49551ba4 / 2

Breadcrumbs:

- xcsh_cdn_cache_rule

Manages a CDN Cache Rule resource in F5 Distributed Cloud for cdn loadbalancer specification.
configuration.

<a id="canonical-703a859bc999f6f5638c6e8c86c98567b7b278b0013c50b0c7d3a2e70714ab90"></a>

## Prerequisites — xcsh_cdn_cache_rule / da3a49551ba4 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-5aa68390830307ce1b1916e4333a1331dae5dab90628b950ed6ef8173398395d"></a>

## Minimal configuration — xcsh_cdn_cache_rule / da3a49551ba4 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CDNCacheRule Resource Example
# Manages a CDN Cache Rule resource in F5 Distributed Cloud for cdn loadbalancer specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNCacheRule configuration
resource "xcsh_cdn_cache_rule" "example" {
  name      = "example-cdn-cache-rule"
  namespace = "staging"
}
```

<a id="canonical-48612da350673b0454fd3d664b43a2ce5fef3e0a52a10832f9dda189d1bbc469"></a>

## Root configuration — xcsh_cdn_cache_rule / da3a49551ba4 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-24ecb62eb208b3288ca45f8c243039b17bcb996d4dfb7f83dbf7f57c5afb3b9e"></a>

## Next pages — xcsh_cdn_cache_rule / da3a49551ba4 / 6

- [Property reference](../guides/resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [Examples](../guides/resources--cdn_cache_rule--examples--group-001.md#canonical-26f48ea8b5f91d0dad499291c518db3babf530eb629d1bb63ef4b4191f7fdfa5)
- [Import](../guides/resources--cdn_cache_rule--lifecycle--group-001.md#canonical-dd361e84d6649bc3c5d88388fa7e6876c2d87ff9f842d0df5df9a5615ddf6367)
- [Timeouts](../guides/resources--cdn_cache_rule--lifecycle--group-001.md#canonical-007c5c6d5ded78d2f65e76e302c09cfb99ac8204148fac0db9c98fa9b93346ba)
