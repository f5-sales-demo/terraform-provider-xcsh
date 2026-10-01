---
page_title: "xcsh_network_policy_rule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_rule landing."
---

# xcsh_network_policy_rule landing

<a id="canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca493fd28fe2cf4c9d8d9651dbfaf28f83ab2f75494f1f795935f3a740901309"></a>

## xcsh_network_policy_rule — xcsh_network_policy_rule / 27bd50252731 / 2

Breadcrumbs:

- xcsh_network_policy_rule

Manages network policy rule with configured parameters in specified namespace in F5 Distributed
Cloud.

<a id="canonical-0a143489c04ec7ada518bc903452a5d9767009930e62bf11bed2004a76fe50e2"></a>

## Prerequisites — xcsh_network_policy_rule / 27bd50252731 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-7e3a5cd73041d7b3b433c525cbb0f5f0926655a7e6e34ad7dc42908ef7264cd1"></a>

## Minimal configuration — xcsh_network_policy_rule / 27bd50252731 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicyRule Resource Example
# Manages network policy rule with configured parameters in specified namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicyRule configuration
resource "xcsh_network_policy_rule" "example" {
  name      = "example-network-policy-rule"
  namespace = "staging"
}
```

<a id="canonical-4d4ebf3813d50ba62c9a727cc5514dbb5c9eaa56f035f1ec3e3b1ed43a0a7607"></a>

## Root configuration — xcsh_network_policy_rule / 27bd50252731 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-87a91f266f2d4cbf494647b4abb80b5353b61c1d6071d774cf1a8b45f6f4346e"></a>

## Next pages — xcsh_network_policy_rule / 27bd50252731 / 6

- [Property reference](../guides/resources--network_policy_rule--reference--group-001.md#canonical-c6cf66c6a64fe4824aab545eb09e94fea44fac2a10f7b8a1bea9efc1b39518f0)
- [Examples](../guides/resources--network_policy_rule--examples--group-001.md#canonical-3e39234cdad1f99cf119f41c27837bd5d429c3fdbdd2062c6c3194ccc84966ad)
- [Import](../guides/resources--network_policy_rule--lifecycle--group-001.md#canonical-017ac2a2da07e714446be716ad390ebf4d106f9f6debaee1144d7b37d8ea33c5)
- [Timeouts](../guides/resources--network_policy_rule--lifecycle--group-001.md#canonical-153204bb246dab33bd07543b7cc3a6a1a868fb8b8feeb2e27bc65abeb833330a)
