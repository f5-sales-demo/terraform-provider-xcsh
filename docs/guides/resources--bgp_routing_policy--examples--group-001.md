---
page_title: "xcsh_bgp_routing_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_routing_policy examples."
---

# xcsh_bgp_routing_policy examples

<a id="canonical-fdbe4961f86b407d0bce43e723f625ebf170cc095b75a73eb63e9c5b06d8e27e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c35de8f09c3d444cf871972b93e6cebac525af1b8044d6705ba9033e126d65e9"></a>

## Examples — Examples / aa23c0f1f59e / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- Examples

<a id="canonical-b2ee78cf85d4c631aac473a7f75a6f0477e8f51f8983fcc471069f302b241d98"></a>

## Complete configurations — Examples / aa23c0f1f59e / 3

- [Resource](resources--bgp_routing_policy--examples--group-001.md#canonical-fd9a0cc5f862db8d59a6e51794dd84f0c05b91053a695901ba162de513c4cf89): valid configuration.

<a id="canonical-d630de7f882da6542b2fe572ac45fe0432c91fb9c4a3ac32393664350b6b0747"></a>

## Next pages — Examples / aa23c0f1f59e / 4

- [Resource](resources--bgp_routing_policy--examples--group-001.md#canonical-fd9a0cc5f862db8d59a6e51794dd84f0c05b91053a695901ba162de513c4cf89)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)

<a id="canonical-fd9a0cc5f862db8d59a6e51794dd84f0c05b91053a695901ba162de513c4cf89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b36cd5ab38d32cfd44f824c801283a9771238970691209c18f14112806358292"></a>

## Resource — Resource / 405060aadd9e / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- [Examples](resources--bgp_routing_policy--examples--group-001.md#canonical-fdbe4961f86b407d0bce43e723f625ebf170cc095b75a73eb63e9c5b06d8e27e)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bgp_routing_policy/resource.tf`; digest `sha256:5114b3e2f71f278003899287dd9fce44709b024db2b4ac16ae2e23054bd84151`.

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

<a id="canonical-5458aeb4a93fa680410f4eae1e1511d94743bdb92f43048e336cb0b0375bbb42"></a>

## Next pages — Resource / 405060aadd9e / 3

- [Examples](resources--bgp_routing_policy--examples--group-001.md#canonical-fdbe4961f86b407d0bce43e723f625ebf170cc095b75a73eb63e9c5b06d8e27e)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
