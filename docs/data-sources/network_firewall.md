---
page_title: "xcsh_network_firewall landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_firewall landing."
---

# xcsh_network_firewall landing

<a id="canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1448b139ff834b5b7f267ab6ac3821649bb474a8fe3f0e5774116f328e99a3f"></a>

## xcsh_network_firewall — xcsh_network_firewall / 9238a8acb185 / 2

Breadcrumbs:

- xcsh_network_firewall

Manages a Network Firewall resource in F5 Distributed Cloud for network firewall is created by users
in system namespace. configuration.

<a id="canonical-02e43d91f675a3492e7f03f98cc87f675786c4e21fc65954e9be5ec088ea42f0"></a>

## Prerequisites — xcsh_network_firewall / 9238a8acb185 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-57b6d24bd375e22524746b52a17bbe694e02f061c92d7f8982626dc7f764b6e1"></a>

## Minimal configuration — xcsh_network_firewall / 9238a8acb185 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkFirewall Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkFirewall by name
data "xcsh_network_firewall" "example" {
  name      = "example-network-firewall"
  namespace = "system"
}

output "network_firewall_id" {
  value = data.xcsh_network_firewall.example.id
}
```

<a id="canonical-a578e5f1dadccbe7755130e8bdae62b980bb1db1d1703717873e44111b893483"></a>

## Root configuration — xcsh_network_firewall / 9238a8acb185 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1ba21f9b254f3dde7038d491278f5da728496c7c9befd1d17e17b43c95a51437"></a>

## Next pages — xcsh_network_firewall / 9238a8acb185 / 6

- [Property reference](../guides/data-sources--network_firewall--reference--group-001.md#canonical-3788cba45b5e266537cf6289bb39431b1592d6bdb952ff53cafe1b93ac1711d3)
- [Examples](../guides/data-sources--network_firewall--examples--group-001.md#canonical-66ef622e203ba224649c17f5e45fd3d55d87a3a0afb4f33e4afc95a5d0d63fdd)
