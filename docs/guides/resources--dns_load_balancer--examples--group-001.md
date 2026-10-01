---
page_title: "xcsh_dns_load_balancer examples"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_load_balancer examples."
---

# xcsh_dns_load_balancer examples

<a id="canonical-83e1b17aff65a9e2ed9945ff957d160b0ea3e62f4fd6603efb706dea85eabc21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82c4e9e24bfdf978f3187ee242f32e8a6e038c3e6b605e31de0dd62abcae2a88"></a>

## Examples — Examples / a72b3e27290c / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- Examples

<a id="canonical-e5899f11f1357b2dc47dd535f497ae80de409fe3e1a4ac927a4a86d37d3c6142"></a>

## Complete configurations — Examples / a72b3e27290c / 3

- [Resource](resources--dns_load_balancer--examples--group-001.md#canonical-dada4625b1bf75166b0723c3f9ffa0252fdf9733f240aa2291f01dd0307727c1): valid configuration.

<a id="canonical-33e7a8ed9938e3c380fb7717fabd8b910ededd7aefa93b6885e36fc477d2a2a3"></a>

## Next pages — Examples / a72b3e27290c / 4

- [Resource](resources--dns_load_balancer--examples--group-001.md#canonical-dada4625b1bf75166b0723c3f9ffa0252fdf9733f240aa2291f01dd0307727c1)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-dada4625b1bf75166b0723c3f9ffa0252fdf9733f240aa2291f01dd0307727c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8d5c8f750e84b7141833a739feb332208f5f50b01dadc1702353ccdff03fb27"></a>

## Resource — Resource / 73e421c707ea / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Examples](resources--dns_load_balancer--examples--group-001.md#canonical-83e1b17aff65a9e2ed9945ff957d160b0ea3e62f4fd6603efb706dea85eabc21)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_load_balancer/resource.tf`; digest `sha256:103fabb6495b3c7befdb807ff9ebc6d11ed6c79f3f4a4083e4f2815709c3a01c`.

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

<a id="canonical-97fad1cbaa007814d825df7187765f02d6663bcc703d670ba7b5942d8ef4f179"></a>

## Next pages — Resource / 73e421c707ea / 3

- [Examples](resources--dns_load_balancer--examples--group-001.md#canonical-83e1b17aff65a9e2ed9945ff957d160b0ea3e62f4fd6603efb706dea85eabc21)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
