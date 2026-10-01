---
page_title: "xcsh_route landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route landing."
---

# xcsh_route landing

<a id="canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2413735406552bf337f0c56fe47d9c3dcc4e492007b9b9c0e598150b1d89c4bf"></a>

## xcsh_route — xcsh_route / 87eaa604d9ad / 2

Breadcrumbs:

- xcsh_route

Manages route object in a given namespace. Route object is list of route rules. Each rule has match
condition to match incoming requests and actions to take on matching requests in F5 Distributed
Cloud.

<a id="canonical-a543ff0ea95bb0537963f2365981abb10924ab06032ca563865252a7925a5d27"></a>

## Prerequisites — xcsh_route / 87eaa604d9ad / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-d793bfbdbc58c33a5bb846445b4eb13ca585431d1c996d007e4c9c0128d7a391"></a>

## Minimal configuration — xcsh_route / 87eaa604d9ad / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Route Resource Example
# Manages route object in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Route configuration
resource "xcsh_route" "example" {
  name      = "example-route"
  namespace = "staging"
}
```

<a id="canonical-9e53fe87305b357e4cdb2d52badb11442cbbc8c0d5894557cbf603abc1bb0e74"></a>

## Root configuration — xcsh_route / 87eaa604d9ad / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-8b2b6527f65e0ff84182f38ae2d1ad6d33c42553d37971c8a484179be82a0c98"></a>

## Next pages — xcsh_route / 87eaa604d9ad / 6

- [Property reference](../guides/resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [Examples](../guides/resources--route--examples--group-001.md#canonical-6aded232f88506504f0ff82eecce407706fca028a2976781dc43e66282e9d3c3)
- [Import](../guides/resources--route--lifecycle--group-001.md#canonical-418303b3d2b1bb887f69b7a33104bb194baaa5f45c8ea5fb6f18f7344e587431)
- [Timeouts](../guides/resources--route--lifecycle--group-001.md#canonical-f3991ef03c10951752c0411c3d95b493e64685b6e6af4a2a11bb43b7e3d06283)
