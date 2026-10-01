---
page_title: "xcsh_tcp_loadbalancer landing"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer landing."
---

# xcsh_tcp_loadbalancer landing

<a id="canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f99fb61322f8227c7523c31bd19def0bea7647d3482247260d23f46613afed9e"></a>

## xcsh_tcp_loadbalancer — xcsh_tcp_loadbalancer / 4d38c9e2d713 / 2

Breadcrumbs:

- xcsh_tcp_loadbalancer

Manages a TCP Load Balancer resource in F5 Distributed Cloud for load balancing TCP traffic across
origin pools.

<a id="canonical-471ed10cd0a4e5b88fc0d1b7d68db1db14f40b059d0d9ff874a89d14c9848a79"></a>

## Prerequisites — xcsh_tcp_loadbalancer / 4d38c9e2d713 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `origin_pool`.

Optional integrations: `healthcheck`.

- origin_pool: Backend servers for TCP/UDP traffic

- healthcheck: Monitor origin server health

<a id="canonical-57cd3d624f7d4184d0c8295ab2adbe7ab1568a59436912565f1713ea23efa436"></a>

## Minimal configuration — xcsh_tcp_loadbalancer / 4d38c9e2d713 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# TCPLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TCPLoadBalancer by name
data "xcsh_tcp_loadbalancer" "example" {
  name      = "example-tcp-loadbalancer"
  namespace = "staging"
}

output "tcp_loadbalancer_id" {
  value = data.xcsh_tcp_loadbalancer.example.id
}
```

<a id="canonical-9f520aa524f1e3fab0578d9111b4b1de77c51e26ed5013d22ef1a3820a593ad7"></a>

## Root configuration — xcsh_tcp_loadbalancer / 4d38c9e2d713 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-79c498e71fa9d58ae821d99c4e95d120305241bf0d7340c1f857776a32578d46"></a>

## Next pages — xcsh_tcp_loadbalancer / 4d38c9e2d713 / 6

- [Property reference](../guides/data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [Examples](../guides/data-sources--tcp_loadbalancer--examples--group-001.md#canonical-65b66562cddaab262db2cae975842adac58ad22257dcba28c2897932fad6d3c7)
