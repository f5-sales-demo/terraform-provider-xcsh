---
page_title: "xcsh_cdn_loadbalancer landing"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer landing."
---

# xcsh_cdn_loadbalancer landing

<a id="canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6bfe32952ba62d41ba8255ee34cf24df204a9dd5399b7cf82b80f86cc9eab1e9"></a>

## xcsh_cdn_loadbalancer — xcsh_cdn_loadbalancer / fd165e066e20 / 2

Breadcrumbs:

- xcsh_cdn_loadbalancer

Manages a CDN Load Balancer resource in F5 Distributed Cloud for content delivery and edge caching
with load balancing.

<a id="canonical-3e9cb90bde82b444754188848a2c84563ed405b45e5d8df1cb8adfedb9d1ee72"></a>

## Prerequisites — xcsh_cdn_loadbalancer / fd165e066e20 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cdn_origin_pool`.

- cdn_origin_pool: Origin servers for CDN content

<a id="canonical-de6d723cd98b100b3b7defd7ca8adf4e9329978b6f8666ad48c5778ca6973430"></a>

## Minimal configuration — xcsh_cdn_loadbalancer / fd165e066e20 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CDNLoadBalancer Resource Example
# Manages a CDN Load Balancer resource in F5 Distributed Cloud for content delivery and edge caching with load balancing.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNLoadBalancer configuration
resource "xcsh_cdn_loadbalancer" "example" {
  name      = "example-cdn-loadbalancer"
  namespace = "staging"

  domains = ["example-value"]
}
```

<a id="canonical-5a289f763423ad0643b08a59900b8fc57e2591f97a22f3e483b7899e218dfff5"></a>

## Root configuration — xcsh_cdn_loadbalancer / fd165e066e20 / 5

Required root properties: `domains`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-c1f50eadf2481dab5e7cd75443c66af76cb57ef53968e13640e59a5f1d66bae0"></a>

## Next pages — xcsh_cdn_loadbalancer / fd165e066e20 / 6

- [Property reference](../guides/resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [Examples](../guides/resources--cdn_loadbalancer--examples--group-001.md#canonical-abfd3355bd1d9030ad4bf5366d5a54f3431a0ea24b3a1f50fa5f3c9767812356)
- [Import](../guides/resources--cdn_loadbalancer--lifecycle--group-001.md#canonical-178f877581617c9362a2e90d6490255c6a0b53b5a0ad51f0566589ba25997b7d)
- [Timeouts](../guides/resources--cdn_loadbalancer--lifecycle--group-001.md#canonical-1269b9cb0aeeacc69f35c45ea16292c267f394d5f1eaa8d4ebc1cf93f2cee792)
