---
page_title: "xcsh_network_firewall examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_firewall examples."
---

# xcsh_network_firewall examples

<a id="canonical-4d393ee01a4205be78f085b1fd174118f811c4d3aef74db20b437356a1fa0aa1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63bbd51d16545bb89d871b82e314a40c932aa5ea3bef68f2a26376f6b4a5a5bb"></a>

## Examples — Examples / 206c6dc58f77 / 2

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
- Examples

<a id="canonical-7a28adc50b9937d083d6484c0dcb928ae8c411db98526092360418ac447d8030"></a>

## Complete configurations — Examples / 206c6dc58f77 / 3

- [Resource](resources--network_firewall--examples--group-001.md#canonical-dd0b3d4731aa926856edc6eac927ad33e29377afcff9f012a340cab210f7ea28): valid configuration.

<a id="canonical-eadb546b2a58e6d2559346261152759131b792fab205c291df04db2942279d9d"></a>

## Next pages — Examples / 206c6dc58f77 / 4

- [Resource](resources--network_firewall--examples--group-001.md#canonical-dd0b3d4731aa926856edc6eac927ad33e29377afcff9f012a340cab210f7ea28)
- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)

<a id="canonical-dd0b3d4731aa926856edc6eac927ad33e29377afcff9f012a340cab210f7ea28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a5bb2bf19e05ca374d9fad9759b587405e86aee0d081e93a262be7bd73dfad7"></a>

## Resource — Resource / 452aa0ce22ea / 2

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
- [Examples](resources--network_firewall--examples--group-001.md#canonical-4d393ee01a4205be78f085b1fd174118f811c4d3aef74db20b437356a1fa0aa1)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_firewall/resource.tf`; digest `sha256:7eeae23d1d7b0c0adbae62b0e59aafa91676ffc023ed6aa05396e73c69627d84`.

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

<a id="canonical-a3d6bb4966973f31aeb5db18a2348243189adb3fe04d18b8a783edaa4d695d6b"></a>

## Next pages — Resource / 452aa0ce22ea / 3

- [Examples](resources--network_firewall--examples--group-001.md#canonical-4d393ee01a4205be78f085b1fd174118f811c4d3aef74db20b437356a1fa0aa1)
- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
