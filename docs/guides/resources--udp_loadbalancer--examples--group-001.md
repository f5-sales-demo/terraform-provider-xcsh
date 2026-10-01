---
page_title: "xcsh_udp_loadbalancer examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer examples."
---

# xcsh_udp_loadbalancer examples

<a id="canonical-52a4f6bdbc3d2127fb3a7f2e0b89dd3a31196ccb7089216ef80a52a3e506b416"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18557a7692206c38f01f14f91b1807549b0c27a889cc16010c404b68aee67da9"></a>

## Examples — Examples / 1a1c5ba7ec66 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- Examples

<a id="canonical-e68410238035227f423d32880231f905e2f2da3b56b067f69ca9a85eec32f2f8"></a>

## Complete configurations — Examples / 1a1c5ba7ec66 / 3

- [Resource](resources--udp_loadbalancer--examples--group-001.md#canonical-71de3b67e8d5e4143c4e53123ff51867c6154cebcb7b31abc4c7fbbe34b37454): valid configuration.

<a id="canonical-57daeee5a0c35601591e49b2f7acd43b0e53a8331fe41b888242179863e46c46"></a>

## Next pages — Examples / 1a1c5ba7ec66 / 4

- [Resource](resources--udp_loadbalancer--examples--group-001.md#canonical-71de3b67e8d5e4143c4e53123ff51867c6154cebcb7b31abc4c7fbbe34b37454)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-71de3b67e8d5e4143c4e53123ff51867c6154cebcb7b31abc4c7fbbe34b37454"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abc7b57845a138bf55ce32cba7702d26ca2883b5b5ebb121906d58aef0c72d56"></a>

## Resource — Resource / 77d6d1836b02 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Examples](resources--udp_loadbalancer--examples--group-001.md#canonical-52a4f6bdbc3d2127fb3a7f2e0b89dd3a31196ccb7089216ef80a52a3e506b416)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_udp_loadbalancer/resource.tf`; digest `sha256:f54b33fab2f01b75ab0f74a179601bfe368063a44f1d292cd4b49a5f0936b681`.

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

<a id="canonical-c86068ab920fe5520f95e34159f0a99f9b22f634d70a782a4aee891cda26be86"></a>

## Next pages — Resource / 77d6d1836b02 / 3

- [Examples](resources--udp_loadbalancer--examples--group-001.md#canonical-52a4f6bdbc3d2127fb3a7f2e0b89dd3a31196ccb7089216ef80a52a3e506b416)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
