---
page_title: "xcsh_dns_lb_health_check landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_health_check landing."
---

# xcsh_dns_lb_health_check landing

<a id="canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8435de5bbd6b81e4ef55137df5646c2c0cce1721c52c3232bef976533208f803"></a>

## xcsh_dns_lb_health_check — xcsh_dns_lb_health_check / 6204225b128d / 2

Breadcrumbs:

- xcsh_dns_lb_health_check

Manages DNS Load Balancer Health Check in a given namespace. If one already exist it will give a
error in F5 Distributed Cloud.

<a id="canonical-c1c45d02a6015e9d6fd2a84de2c4d724db048fae98e57a3beb2437ae996e8946"></a>

## Prerequisites — xcsh_dns_lb_health_check / 6204225b128d / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-c461f3e86660feebc3adfdb35090528613e7e488aab908fab199e18248f5c6b1"></a>

## Minimal configuration — xcsh_dns_lb_health_check / 6204225b128d / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSLBHealthCheck Resource Example
# Manages DNS Load Balancer Health Check in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSLBHealthCheck configuration
resource "xcsh_dns_lb_health_check" "example" {
  name      = "example-dns-lb-health-check"
  namespace = "system"
}
```

<a id="canonical-6e70f0ce7f90a206e8d5100f7646bef98824a2da0b3b1b196f09ffb317aef6ce"></a>

## Root configuration — xcsh_dns_lb_health_check / 6204225b128d / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-82f4db848e794d539145d1de25470e0dec0350d344c62d7b99c627e8f866d444"></a>

## Next pages — xcsh_dns_lb_health_check / 6204225b128d / 6

- [Property reference](../guides/resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- [Examples](../guides/resources--dns_lb_health_check--examples--group-001.md#canonical-055fc70af00770f8fbe1dc432055b10902d499b3622da79cb7783c23e528e6a3)
- [Import](../guides/resources--dns_lb_health_check--lifecycle--group-001.md#canonical-ff3546e2a1862ca4dc7c4a12a2f1113035196a76fbe92ebc869fac0d98978f2b)
- [Timeouts](../guides/resources--dns_lb_health_check--lifecycle--group-001.md#canonical-23373e2b7386a6c0e46cdc5a691f2c27ae5fc42c2da65c77fbcfc4d57453d584)
