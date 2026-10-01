---
page_title: "xcsh_network_policy landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_policy landing."
---

# xcsh_network_policy landing

<a id="canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-331ad5eeb0e456594dc7280c9c4cf801acbf0680605ca4520a8bd113c25cd090"></a>

## xcsh_network_policy — xcsh_network_policy / f0573e00138a / 2

Breadcrumbs:

- xcsh_network_policy

Manages new network policy with configured parameters in specified namespace in F5 Distributed
Cloud.

<a id="canonical-4416d95e68ffd3e5a913485f199dfe05c369af3d3fc7e813051c79f274e3225e"></a>

## Prerequisites — xcsh_network_policy / f0573e00138a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-e4e311ff477fd946ee4be60b5ae4f03d8bdedb887e6fe1d08cc6cbbe9925758d"></a>

## Minimal configuration — xcsh_network_policy / f0573e00138a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicy Resource Example
# Manages new network policy with configured parameters in specified namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicy configuration
resource "xcsh_network_policy" "example" {
  name      = "example-network-policy"
  namespace = "staging"
}
```

<a id="canonical-bc2dc903a2fe6f78c1f404cb2becf06a1058407b35867c2839e0d73bd4458eed"></a>

## Root configuration — xcsh_network_policy / f0573e00138a / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-c60eddc608fae064af4751015d57137939283cfab90d7fa63e0957c99a8158a7"></a>

## Next pages — xcsh_network_policy / f0573e00138a / 6

- [Property reference](../guides/resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [Examples](../guides/resources--network_policy--examples--group-001.md#canonical-ed842b56a95a4e0145e96ee3d036d74fdbe1d9e6ffda6cefc8267078b0b528a9)
- [Import](../guides/resources--network_policy--lifecycle--group-001.md#canonical-238d2d12ec7fab35259efa2a18d886812d265e9dc5a3d8fabe54fd84d3dc112a)
- [Timeouts](../guides/resources--network_policy--lifecycle--group-001.md#canonical-b7dab1a1112d50d39cdc49509ab0a31c4d8be4a4a7974d9dc6cacef06cecd014)
