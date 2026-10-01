---
page_title: "xcsh_dns_lb_health_check examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_health_check examples."
---

# xcsh_dns_lb_health_check examples

<a id="canonical-055fc70af00770f8fbe1dc432055b10902d499b3622da79cb7783c23e528e6a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47b69560e6f1f09b06b0ff14b60db9a3af5da54e772a1ed3e751f9d5ffca7cbd"></a>

## Examples — Examples / 7b2ba997b1fd / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
- Examples

<a id="canonical-a37824d1e998780c8d635984051dd7137e8a303e3f7be80bd57a1580c658361d"></a>

## Complete configurations — Examples / 7b2ba997b1fd / 3

- [Resource](resources--dns_lb_health_check--examples--group-001.md#canonical-6693da9567d88ab4225cb1f602720e6f660e0a5ddd75f950c7352f9b9902eb61): valid configuration.

<a id="canonical-e178ecd74c06cbdc7e1698601994b8244d4df0103f9092aa462a7294745f6507"></a>

## Next pages — Examples / 7b2ba997b1fd / 4

- [Resource](resources--dns_lb_health_check--examples--group-001.md#canonical-6693da9567d88ab4225cb1f602720e6f660e0a5ddd75f950c7352f9b9902eb61)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)

<a id="canonical-6693da9567d88ab4225cb1f602720e6f660e0a5ddd75f950c7352f9b9902eb61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1faab17706c8090c5c8e4c134fd5b57400b001f56a4f015f6dc53842fd2ed1ee"></a>

## Resource — Resource / c84008439ee8 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
- [Examples](resources--dns_lb_health_check--examples--group-001.md#canonical-055fc70af00770f8fbe1dc432055b10902d499b3622da79cb7783c23e528e6a3)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_lb_health_check/resource.tf`; digest `sha256:219de26c8a2c69aaf219ac20d2e05053708abed48480ac00f40c1cb16da5893e`.

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

<a id="canonical-043d9e8bb4b62d3c3d999a50a5351b1e6a80c83570d88d41f8abad247d84a01d"></a>

## Next pages — Resource / c84008439ee8 / 3

- [Examples](resources--dns_lb_health_check--examples--group-001.md#canonical-055fc70af00770f8fbe1dc432055b10902d499b3622da79cb7783c23e528e6a3)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
