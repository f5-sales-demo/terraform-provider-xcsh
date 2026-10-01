---
page_title: "xcsh_network_firewall landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_firewall landing."
---

# xcsh_network_firewall landing

<a id="canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-494c216c554ecf9290286fabd44e0a2f4230a0d692f0f38bb5f751f258f447f6"></a>

## xcsh_network_firewall — xcsh_network_firewall / c473d1d90e00 / 2

Breadcrumbs:

- xcsh_network_firewall

Manages a Network Firewall resource in F5 Distributed Cloud for network firewall is created by users
in system namespace. configuration.

<a id="canonical-aeef1e1004afd8e378a7993aedc2d3f2445a9d3c504cdf707f1cb68f1eeb12ef"></a>

## Prerequisites — xcsh_network_firewall / c473d1d90e00 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-4c2e4e24f0fb036e6e429f5755245af049025dbae4ebc77cf89ef00ea2e731f6"></a>

## Minimal configuration — xcsh_network_firewall / c473d1d90e00 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkFirewall Resource Example
# Manages a Network Firewall resource in F5 Distributed Cloud for network firewall is created by users in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkFirewall configuration
resource "xcsh_network_firewall" "example" {
  name      = "example-network-firewall"
  namespace = "system"
}
```

<a id="canonical-4a24fa562de7d0f3c14f71ac0ca88f9b24463888db6cce91d01571af3678f6e7"></a>

## Root configuration — xcsh_network_firewall / c473d1d90e00 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-7e0e21884ca44d21e78a905879478a693705b24b2dce97df39d5fe17569482f0"></a>

## Next pages — xcsh_network_firewall / c473d1d90e00 / 6

- [Property reference](../guides/resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- [Examples](../guides/resources--network_firewall--examples--group-001.md#canonical-4d393ee01a4205be78f085b1fd174118f811c4d3aef74db20b437356a1fa0aa1)
- [Import](../guides/resources--network_firewall--lifecycle--group-001.md#canonical-3e3abc62914d512e29685590aa5d0358e6c9aab5b040a2af0eadbef2c68de363)
- [Timeouts](../guides/resources--network_firewall--lifecycle--group-001.md#canonical-673f7ac9a575f532f891576211a2bb959c6bd0a34c8acc76b0aabfcde516ebac)
