---
page_title: "xcsh_bgp_routing_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_routing_policy landing."
---

# xcsh_bgp_routing_policy landing

<a id="canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04bdb217961d202f3553b90dfec6e76ff0b4ee59bbdf37fe79d2af7b2ffb3976"></a>

## xcsh_bgp_routing_policy — xcsh_bgp_routing_policy / 44365971262c / 2

Breadcrumbs:

- xcsh_bgp_routing_policy

Manages a BGP Routing Policy resource in F5 Distributed Cloud for bgp routing policy is a list of
rules containing match criteria and action to be applied. these rules help control routes which are
imported or exported to bgp peers. configuration.

<a id="canonical-5f6b67c6866f7bf07868d33de32961fc4ee1ea78786d8bb4d80c44ae28d5e773"></a>

## Prerequisites — xcsh_bgp_routing_policy / 44365971262c / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-f0acaa1f0112fc7371b4aa265e68b326ed8613fb186f63e842f56609b93571df"></a>

## Minimal configuration — xcsh_bgp_routing_policy / 44365971262c / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGPRoutingPolicy Resource Example
# Manages a BGP Routing Policy resource in F5 Distributed Cloud for bgp routing policy is a list of rules containing match criteria and action to be applied.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BGPRoutingPolicy configuration
resource "xcsh_bgp_routing_policy" "example" {
  name      = "example-bgp-routing-policy"
  namespace = "staging"
}
```

<a id="canonical-bf923c00321d60e4f09f9d8da4813e68d62d13cbb5620989154b2b4a2b630f6b"></a>

## Root configuration — xcsh_bgp_routing_policy / 44365971262c / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-021ddb19fe11248deb0880cf490a06ac27a94e721ed2a914992f3701d972a4b3"></a>

## Next pages — xcsh_bgp_routing_policy / 44365971262c / 6

- [Property reference](../guides/resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- [Examples](../guides/resources--bgp_routing_policy--examples--group-001.md#canonical-fdbe4961f86b407d0bce43e723f625ebf170cc095b75a73eb63e9c5b06d8e27e)
- [Import](../guides/resources--bgp_routing_policy--lifecycle--group-001.md#canonical-b2f9ec3ffd9acc7a047dc67ff0281d3c0724379f46ced7f487adae92a12e8fda)
- [Timeouts](../guides/resources--bgp_routing_policy--lifecycle--group-001.md#canonical-c592839e358ee317aef5197bc9cacf91cb685e2d3189e737879ab44366916b75)
