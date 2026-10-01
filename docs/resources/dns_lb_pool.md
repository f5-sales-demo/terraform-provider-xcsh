---
page_title: "xcsh_dns_lb_pool landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_pool landing."
---

# xcsh_dns_lb_pool landing

<a id="canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8440b6f529ed73a6297550d178f111dbe31456a1796f9ef3b72369127a5c3738"></a>

## xcsh_dns_lb_pool — xcsh_dns_lb_pool / 1e975d067e95 / 2

Breadcrumbs:

- xcsh_dns_lb_pool

Manages DNS Load Balancer Pool in a given namespace. If one already exist it will give a error in F5
Distributed Cloud.

<a id="canonical-016df376c42d6cb2139672b1ad3272408b461b7ccc9a133894a11ba3cc918cd4"></a>

## Prerequisites — xcsh_dns_lb_pool / 1e975d067e95 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-28f9b66fde69bcd3679369cbb63ef4cb425bf17d4c4c1126538d979158e796e7"></a>

## Minimal configuration — xcsh_dns_lb_pool / 1e975d067e95 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSLBPool Resource Example
# Manages DNS Load Balancer Pool in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSLBPool configuration
resource "xcsh_dns_lb_pool" "example" {
  name      = "example-dns-lb-pool"
  namespace = "system"
}
```

<a id="canonical-e46bdbbf802b04f8038af228c86b358e7f32bbbcdcfd02369be42bf2a852c70a"></a>

## Root configuration — xcsh_dns_lb_pool / 1e975d067e95 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-e04665f6fe9e6a0ce226d4221b4e05aa0e78cf611d5191c2054c4bd819ad17d9"></a>

## Next pages — xcsh_dns_lb_pool / 1e975d067e95 / 6

- [Property reference](../guides/resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [Examples](../guides/resources--dns_lb_pool--examples--group-001.md#canonical-c9d179e71810ad3d046c54203f75de777a65f82b4e362f3fcb56fc10b8a6fcc0)
- [Import](../guides/resources--dns_lb_pool--lifecycle--group-001.md#canonical-422811f1450679f47c551f868d9fbc36708b742963fa95d6a12847e7432ac153)
- [Timeouts](../guides/resources--dns_lb_pool--lifecycle--group-001.md#canonical-a6da6af97ce085b8e29bb920e35e05fd69771cc396acba71fb9a9290ae7648ea)
