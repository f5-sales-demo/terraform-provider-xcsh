---
page_title: "xcsh_http_loadbalancer landing"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer landing."
---

# xcsh_http_loadbalancer landing

<a id="canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2bfb1b3ba464f0362add41bac67f0905d3f962a2948ff142bf3af2dfd4a97edd"></a>

## xcsh_http_loadbalancer — xcsh_http_loadbalancer / c41a11eb2d60 / 2

Breadcrumbs:

- xcsh_http_loadbalancer

Manages a HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic
with routing and security controls.

<a id="canonical-04ac39c72b5e71824e617dae9d7a7bc1772dba1bab01e2d074e909ec82c74e70"></a>

## Prerequisites — xcsh_http_loadbalancer / c41a11eb2d60 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `origin_pool`.

Optional integrations: `healthcheck`, `app_firewall`, `certificate`, `rate_limiter`, `service_policy`, `bot_defense_policy`.

- origin_pool: Backend servers for traffic distribution

- app_firewall: WAF protection (requires WAAP subscription)

- healthcheck: Monitor backend availability

- certificate: TLS termination for HTTPS

- rate_limiter: Protect against traffic spikes

<a id="canonical-f81ed5e59e0777c9b2b9232f76cb8e89ddb55aa9854a64abec5c6e14fd898987"></a>

## Minimal configuration — xcsh_http_loadbalancer / c41a11eb2d60 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# HTTPLoadBalancer Resource Example
# Manages a HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with routing and security controls.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic HTTPLoadBalancer configuration
resource "xcsh_http_loadbalancer" "example" {
  name      = "example-http-loadbalancer"
  namespace = "staging"

  domains = ["example-value"]
}
```

<a id="canonical-c7c61aac6bdb233d718ecd25038db9b553087246f2d0429d5100e3c8f35e975b"></a>

## Root configuration — xcsh_http_loadbalancer / c41a11eb2d60 / 5

Required root properties: `domains`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-9c745e41d2320a56c73685605e2802da930aaf9db3bb5aafebeaf5d7d6a448d4"></a>

## Next pages — xcsh_http_loadbalancer / c41a11eb2d60 / 6

- [Property reference](../guides/resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [Examples](../guides/resources--http_loadbalancer--examples--group-001.md#canonical-89302db2faa0fec01bb0e0bcfbe8dc56ed072f96ee33b11650c93e01dec7da23)
- [Import](../guides/resources--http_loadbalancer--lifecycle--group-001.md#canonical-404662cb13153781f14bcb346c112626eb5892ffa0d5fc154d46443ea5f41135)
- [Timeouts](../guides/resources--http_loadbalancer--lifecycle--group-001.md#canonical-ae66b6ed8a65cc19135c6b474b0500b1990f74cc503e089c95b147e4c8213b47)
