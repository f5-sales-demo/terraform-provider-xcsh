---
page_title: "xcsh_bgp_routing_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_routing_policy landing."
---

# xcsh_bgp_routing_policy landing

<a id="canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1f4ec795f7fb4f2afc2aba2563caafcfe1ddf4b1279ecf2e5d0b9d598c12705"></a>

## xcsh_bgp_routing_policy — xcsh_bgp_routing_policy / aa67c075a44a / 2

Breadcrumbs:

- xcsh_bgp_routing_policy

Manages a BGP Routing Policy resource in F5 Distributed Cloud for bgp routing policy is a list of
rules containing match criteria and action to be applied. these rules help control routes which are
imported or exported to bgp peers. configuration.

<a id="canonical-25787579a4f93ca6ce6dcc8379c803793de1fa8cbf9a38c0fb6a236ae433f5bc"></a>

## Prerequisites — xcsh_bgp_routing_policy / aa67c075a44a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2fe8ac201b6d39f9a0080be4cebfd15e319f027a06119adfec8f8398f4b81857"></a>

## Minimal configuration — xcsh_bgp_routing_policy / aa67c075a44a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGPRoutingPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGPRoutingPolicy by name
data "xcsh_bgp_routing_policy" "example" {
  name      = "example-bgp-routing-policy"
  namespace = "staging"
}

output "bgp_routing_policy_id" {
  value = data.xcsh_bgp_routing_policy.example.id
}
```

<a id="canonical-bdc5c82107692c5f9abfac7c38273800851ac249aa0f98a7447793fe3e825b7e"></a>

## Root configuration — xcsh_bgp_routing_policy / aa67c075a44a / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-d3c8c7c59069e37259402df70395bf34da24afc9a53258ad6fed6b29091f2f34"></a>

## Next pages — xcsh_bgp_routing_policy / aa67c075a44a / 6

- [Property reference](../guides/data-sources--bgp_routing_policy--reference--group-001.md#canonical-eccae3ab92e4281f4d8846e4b097e7e7a2471f1c57d3a01c64e029bb315272fb)
- [Examples](../guides/data-sources--bgp_routing_policy--examples--group-001.md#canonical-081e2a3754d1a66d63c5257d16e6901391a00cf262454a13ddc7c0ead7261e26)
