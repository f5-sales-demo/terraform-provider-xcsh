---
page_title: "xcsh_dns_load_balancer landing"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_load_balancer landing."
---

# xcsh_dns_load_balancer landing

<a id="canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba37721dbe6bbf714cc455c9dcdc9fdc279f4ec339c25b1f3f6b4fbcc58ae8bd"></a>

## xcsh_dns_load_balancer — xcsh_dns_load_balancer / f6267e53f364 / 2

Breadcrumbs:

- xcsh_dns_load_balancer

Manages DNS Load Balancer in a given namespace. If one already exist it will give a error in F5
Distributed Cloud.

<a id="canonical-e764d65b354e954bcfa73f4a0ccc10edca3556f779dbaea285c8fb1165387165"></a>

## Prerequisites — xcsh_dns_load_balancer / f6267e53f364 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `dns_zone`.

- dns_zone: Parent zone for DNS records

<a id="canonical-e6551273e5967d7496b443c65741bca744fe67e95a74db35e475e508ff171093"></a>

## Minimal configuration — xcsh_dns_load_balancer / f6267e53f364 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSLoadBalancer by name
data "xcsh_dns_load_balancer" "example" {
  name      = "example-dns-load-balancer"
  namespace = "system"
}

output "dns_load_balancer_id" {
  value = data.xcsh_dns_load_balancer.example.id
}
```

<a id="canonical-ce70e0e896bc053928be8d5a3482638e6a0b49f94fe16ebc53d0169661824aa9"></a>

## Root configuration — xcsh_dns_load_balancer / f6267e53f364 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-357ff9e7d4b75ba1ffb18e102d6fec2ffddcf5bce7258a7c1a186af4968d82d4"></a>

## Next pages — xcsh_dns_load_balancer / f6267e53f364 / 6

- [Property reference](../guides/data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [Examples](../guides/data-sources--dns_load_balancer--examples--group-001.md#canonical-53377b934401d143eba0bcda241ab1eabf8ad0e63f31cec3dd8a9ba854484274)
