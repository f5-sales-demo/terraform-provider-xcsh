---
page_title: "xcsh_dns_load_balancer landing"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_load_balancer landing."
---

# xcsh_dns_load_balancer landing

<a id="canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2853be8249c89f7d71a36c2d9c01baca45ac3055f3bddf554e37e9dd00f9388f"></a>

## xcsh_dns_load_balancer — xcsh_dns_load_balancer / 51f900ef6841 / 2

Breadcrumbs:

- xcsh_dns_load_balancer

Manages DNS Load Balancer in a given namespace. If one already exist it will give a error in F5
Distributed Cloud.

<a id="canonical-7025ddb915451b2d6d05e178cde191c0379681fe55abd0619c14b2e254f60642"></a>

## Prerequisites — xcsh_dns_load_balancer / 51f900ef6841 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `dns_zone`.

- dns_zone: Parent zone for DNS records

<a id="canonical-16db2a78c88ca51a5d995d4835577898a4fa6ae7ab92ccb07449164217580b09"></a>

## Minimal configuration — xcsh_dns_load_balancer / 51f900ef6841 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSLoadBalancer Resource Example
# Manages DNS Load Balancer in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSLoadBalancer configuration
resource "xcsh_dns_load_balancer" "example" {
  name      = "example-dns-load-balancer"
  namespace = "system"
}
```

<a id="canonical-1c1efa0f752afa359424011d081f080e60607619d4b6b1fa44defcdd98eb792b"></a>

## Root configuration — xcsh_dns_load_balancer / 51f900ef6841 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-861a00a0cb79b10091ec97fda64c9d5c68bae1d2fa255dcf7fe7a08cf2bbf44e"></a>

## Next pages — xcsh_dns_load_balancer / 51f900ef6841 / 6

- [Property reference](../guides/resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [Examples](../guides/resources--dns_load_balancer--examples--group-001.md#canonical-83e1b17aff65a9e2ed9945ff957d160b0ea3e62f4fd6603efb706dea85eabc21)
- [Import](../guides/resources--dns_load_balancer--lifecycle--group-001.md#canonical-1c9044ea645e38eddb76684426b45cc8354e794f52cb6e7d8dc67edb7b738314)
- [Timeouts](../guides/resources--dns_load_balancer--lifecycle--group-001.md#canonical-8bfb142fed7a9c4bf590de06eeb9feacfe35c24b13ab118dcdc8f758cc06ae63)
