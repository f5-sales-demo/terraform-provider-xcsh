---
page_title: "xcsh_http_loadbalancer landing"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer landing."
---

# xcsh_http_loadbalancer landing

<a id="canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46e0b4d089461c7c5277cd695b38dda7eba54cc9a71c543666d0ae3b42e9172c"></a>

## xcsh_http_loadbalancer — xcsh_http_loadbalancer / 4da6e2b99b16 / 2

Breadcrumbs:

- xcsh_http_loadbalancer

Manages a HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic
with routing and security controls.

<a id="canonical-89da182d8c7ad828db61cc72d918128390d29f2329dd909d165d3f717e1892f1"></a>

## Prerequisites — xcsh_http_loadbalancer / 4da6e2b99b16 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `origin_pool`.

Optional integrations: `healthcheck`, `app_firewall`, `certificate`, `rate_limiter`, `service_policy`, `bot_defense_policy`.

- origin_pool: Backend servers for traffic distribution

- app_firewall: WAF protection (requires WAAP subscription)

- healthcheck: Monitor backend availability

- certificate: TLS termination for HTTPS

- rate_limiter: Protect against traffic spikes

<a id="canonical-0d317c4b25f4b943ee7607c4561f97ce755aa20a4750ef4a75db40b730a6e373"></a>

## Minimal configuration — xcsh_http_loadbalancer / 4da6e2b99b16 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# HTTPLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing HTTPLoadBalancer by name
data "xcsh_http_loadbalancer" "example" {
  name      = "example-http-loadbalancer"
  namespace = "staging"
}

output "http_loadbalancer_id" {
  value = data.xcsh_http_loadbalancer.example.id
}
```

<a id="canonical-865d5c8fe91660beaf0504f85fa0b0cd79eb3a9d9c8e227d5055244adabd195f"></a>

## Root configuration — xcsh_http_loadbalancer / 4da6e2b99b16 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3a70b113dc44c1da562faf50a3f16de29e5fc6b49bac386719d7e16ee3f2ce1c"></a>

## Next pages — xcsh_http_loadbalancer / 4da6e2b99b16 / 6

- [Property reference](../guides/data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [Examples](../guides/data-sources--http_loadbalancer--examples--group-001.md#canonical-864faaea9ccf4e93d6c60b5343dae9fe756da6f5fa5779fee9023c494681a6d4)
