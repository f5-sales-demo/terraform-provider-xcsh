---
page_title: "xcsh_dns_proxy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy landing."
---

# xcsh_dns_proxy landing

<a id="canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-247a321d150d3a979515cb730e595462537e0322e239ce7d7022774eb39ba0cb"></a>

## xcsh_dns_proxy — xcsh_dns_proxy / e1ca4a125a1f / 2

Breadcrumbs:

- xcsh_dns_proxy

Manages DNS Proxy in a given namespace. If one already exists it will give an error in F5
Distributed Cloud.

<a id="canonical-d652f05696f042990c9c1396c30fc12c9957eb9eb053e2dc7c3564ac29e4e0e2"></a>

## Prerequisites — xcsh_dns_proxy / e1ca4a125a1f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-b950f8b3f1745465ee54109bb0a45483530a6c032f83361fa66db9524fd195f2"></a>

## Minimal configuration — xcsh_dns_proxy / e1ca4a125a1f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSProxy Resource Example
# Manages DNS Proxy in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSProxy configuration
resource "xcsh_dns_proxy" "example" {
  name      = "example-dns-proxy"
  namespace = "system"
}
```

<a id="canonical-fd2fbce358591138ed00af8164b94861a42273801506abddcfa728bceaf805d3"></a>

## Root configuration — xcsh_dns_proxy / e1ca4a125a1f / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-f2cd1b441d82b76b5e5bbaf407ca6fbe9af0ae4e53fe0d386b9723d82b613c94"></a>

## Next pages — xcsh_dns_proxy / e1ca4a125a1f / 6

- [Property reference](../guides/resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [Examples](../guides/resources--dns_proxy--examples--group-001.md#canonical-9fa90dae695ab8998220060b8a3ee60cd4dfcfd96133704fb891d295262d40e5)
- [Import](../guides/resources--dns_proxy--lifecycle--group-001.md#canonical-f3bfd76d635a929d47f389f9c3f5813e9e54c2e1012f9cb23782add538ad5a9d)
- [Timeouts](../guides/resources--dns_proxy--lifecycle--group-001.md#canonical-4706b3992f5d0f4ff8e2cf6d938d28bb6e371b0a0aa517f31c086296d6e5a520)
