---
page_title: "xcsh_udp_loadbalancer landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer landing."
---

# xcsh_udp_loadbalancer landing

<a id="canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef815a58a94937babf7009a299f4bdce932446f03fb88098d60a97752ddd5026"></a>

## xcsh_udp_loadbalancer — xcsh_udp_loadbalancer / c3ba271f093f / 2

Breadcrumbs:

- xcsh_udp_loadbalancer

Manages a UDP Load Balancer resource in F5 Distributed Cloud for load balancing UDP traffic across
origin pools.

<a id="canonical-64ffd8636c2bafc0ac24acadb2a3deadee1d9c98ced505de2b71b7f370e0cfee"></a>

## Prerequisites — xcsh_udp_loadbalancer / c3ba271f093f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-4a957503cd4d98880eed7e08961ec73b90d6e9cf730d5d2ed6af93457bf2d178"></a>

## Minimal configuration — xcsh_udp_loadbalancer / c3ba271f093f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UDPLoadBalancer Resource Example
# Manages a UDP Load Balancer resource in F5 Distributed Cloud for load balancing UDP traffic across origin pools.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UDPLoadBalancer configuration
resource "xcsh_udp_loadbalancer" "example" {
  name      = "example-udp-loadbalancer"
  namespace = "staging"
}
```

<a id="canonical-fa751e54d81e95234ce422a3b130086e88b9cb253bd14da201a03ba99d310a1f"></a>

## Root configuration — xcsh_udp_loadbalancer / c3ba271f093f / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1a47e95f3be13990f1c7945e1fa07616fe13db2e3488c623765f9e5a1e22e928"></a>

## Next pages — xcsh_udp_loadbalancer / c3ba271f093f / 6

- [Property reference](../guides/resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [Examples](../guides/resources--udp_loadbalancer--examples--group-001.md#canonical-52a4f6bdbc3d2127fb3a7f2e0b89dd3a31196ccb7089216ef80a52a3e506b416)
- [Import](../guides/resources--udp_loadbalancer--lifecycle--group-001.md#canonical-7a9a0a828f13ad0c69ad6dffe543bed7d504bdc4b77f94482251ab97d2bba8bc)
- [Timeouts](../guides/resources--udp_loadbalancer--lifecycle--group-001.md#canonical-99a56a9c00fd60e99225b6a094034d0afe83500882da47184ff23fe46ac791db)
