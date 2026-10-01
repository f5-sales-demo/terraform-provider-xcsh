---
page_title: "xcsh_proxy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy landing."
---

# xcsh_proxy landing

<a id="canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23cc7b5567b977d1d6af814186d0abaf4ccf151de16cf0ef18a8008923fbdfe8"></a>

## xcsh_proxy — xcsh_proxy / 2be16b67fba1 / 2

Breadcrumbs:

- xcsh_proxy

Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification.
configuration.

<a id="canonical-e817e8cff1f33573cc792d100051df4dc20a836663b27f01c84c3f2490ff6bec"></a>

## Prerequisites — xcsh_proxy / 2be16b67fba1 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-c24560618610afa7a232ffa5b4926e5212926ed237d7ca7e3fced6e1bfdc9d5f"></a>

## Minimal configuration — xcsh_proxy / 2be16b67fba1 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Proxy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Proxy by name
data "xcsh_proxy" "example" {
  name      = "example-proxy"
  namespace = "staging"
}

output "proxy_id" {
  value = data.xcsh_proxy.example.id
}
```

<a id="canonical-7061586c451e4b0ace0ff875c62e3b676422fc4bc4eba3a4c55e3b899255bc9c"></a>

## Root configuration — xcsh_proxy / 2be16b67fba1 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-349406f4db3cd8dc326e8023e836f83dc01ea43957a27aa13587259cdf01bc3d"></a>

## Next pages — xcsh_proxy / 2be16b67fba1 / 6

- [Property reference](../guides/data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [Examples](../guides/data-sources--proxy--examples--group-001.md#canonical-d14acff9d989657492459e395475a529ffb9792ff80fe7c7e6c110e70623badf)
