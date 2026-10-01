---
page_title: "xcsh_tcp_loadbalancer landing"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer landing."
---

# xcsh_tcp_loadbalancer landing

<a id="canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6cdcb3caa044d72ca2fa421444bdcf68bc7faf7ddf7eafdc53a0d2c85389cd7"></a>

## xcsh_tcp_loadbalancer — xcsh_tcp_loadbalancer / a6963a294f66 / 2

Breadcrumbs:

- xcsh_tcp_loadbalancer

Manages a TCP Load Balancer resource in F5 Distributed Cloud for load balancing TCP traffic across
origin pools.

<a id="canonical-4322dd661721ad904d6ba737b4771ea72c9103c2bc32c9ef6bd1683a87aa6edb"></a>

## Prerequisites — xcsh_tcp_loadbalancer / a6963a294f66 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `origin_pool`.

Optional integrations: `healthcheck`.

- origin_pool: Backend servers for TCP/UDP traffic

- healthcheck: Monitor origin server health

<a id="canonical-5517a3e68b3573252921facf1abefd14268693f2997d2c5351445444257809a3"></a>

## Minimal configuration — xcsh_tcp_loadbalancer / a6963a294f66 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# TCPLoadBalancer Resource Example
# Manages a TCP Load Balancer resource in F5 Distributed Cloud for load balancing TCP traffic across origin pools.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TCPLoadBalancer configuration
resource "xcsh_tcp_loadbalancer" "example" {
  name      = "example-tcp-loadbalancer"
  namespace = "staging"
}
```

<a id="canonical-eec8a056e1bfec557105b8691fe8f664c47ed1fcb8c3568c394eba040b69c4b0"></a>

## Root configuration — xcsh_tcp_loadbalancer / a6963a294f66 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-a51c44209d71fd229373bcedc85e12686c9c0d3788b4317360c964d82c3a8eb8"></a>

## Next pages — xcsh_tcp_loadbalancer / a6963a294f66 / 6

- [Property reference](../guides/resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [Examples](../guides/resources--tcp_loadbalancer--examples--group-001.md#canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0)
- [Import](../guides/resources--tcp_loadbalancer--lifecycle--group-001.md#canonical-044dcf01f9fc47c41b13129a078444dd9340bbd7f0691fec9285325c81677869)
- [Timeouts](../guides/resources--tcp_loadbalancer--lifecycle--group-001.md#canonical-17b3f0b6d9e50dccbe92c9d4418e3ca18c62555b80225148baf8b177e68fb9eb)
