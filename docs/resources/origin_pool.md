---
page_title: "xcsh_origin_pool landing"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool landing."
---

# xcsh_origin_pool landing

<a id="canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6713b8e594dcce3f30e8e8356a23d584df5ef5fdd155c8db29e2b75641d83354"></a>

## xcsh_origin_pool — xcsh_origin_pool / d3ee6f5fb0fd / 2

Breadcrumbs:

- xcsh_origin_pool

Manages a Origin Pool resource in F5 Distributed Cloud for defining backend server pools for load
balancer targets.

<a id="canonical-b5b84d8c2c31f3f8e48ac6a496ad23ca348a181cc313b0703dcd9e9d74b1c520"></a>

## Prerequisites — xcsh_origin_pool / d3ee6f5fb0fd / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `healthcheck`.

- healthcheck: Monitor origin server health

<a id="canonical-0af529ac10c27aaaeea524943adc5b88b18f221ef89f003305d64e636e79fd30"></a>

## Minimal configuration — xcsh_origin_pool / d3ee6f5fb0fd / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# OriginPool Resource Example
# Manages a Origin Pool resource in F5 Distributed Cloud for defining backend server pools for load balancer targets.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic OriginPool configuration
resource "xcsh_origin_pool" "example" {
  name      = "example-origin-pool"
  namespace = "staging"
}
```

<a id="canonical-ca8720f7fd83465e69b7f0d9f52ac5eee3ebe8ce49b0d0b918a6653cefa79529"></a>

## Root configuration — xcsh_origin_pool / d3ee6f5fb0fd / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-6947af9465ae343175915decfecf3efaf0c38d00277d94f4e3aac3e825ff5819"></a>

## Next pages — xcsh_origin_pool / d3ee6f5fb0fd / 6

- [Property reference](../guides/resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [Examples](../guides/resources--origin_pool--examples--group-001.md#canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c)
- [Import](../guides/resources--origin_pool--lifecycle--group-001.md#canonical-1dbef138a46a53f27737d16a2eaecfed87ed7fb4677e2d9b12a53ab38972e0ef)
- [Timeouts](../guides/resources--origin_pool--lifecycle--group-001.md#canonical-26dcc6b21288679d1f8d9ca7ea7abf0f366e33ad7e01a1af53ba0cf91763e463)
