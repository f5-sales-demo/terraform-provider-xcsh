---
page_title: "xcsh_origin_pool landing"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool landing."
---

# xcsh_origin_pool landing

<a id="canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7bc106ad6a04f86cf2bf646ff563b0a8aee9b9c08664af8047f3f8fa5d573b3"></a>

## xcsh_origin_pool — xcsh_origin_pool / 726ce2e68a36 / 2

Breadcrumbs:

- xcsh_origin_pool

Manages a Origin Pool resource in F5 Distributed Cloud for defining backend server pools for load
balancer targets.

<a id="canonical-37a222c40015e9329987a6f5562bd0d519798fa41af1cf10ff0d574512d22d58"></a>

## Prerequisites — xcsh_origin_pool / 726ce2e68a36 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `healthcheck`.

- healthcheck: Monitor origin server health

<a id="canonical-60d3cd70f52f6695aeb991c57831704d9a7a58def8e1e24761b9c0147b885601"></a>

## Minimal configuration — xcsh_origin_pool / 726ce2e68a36 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# OriginPool Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing OriginPool by name
data "xcsh_origin_pool" "example" {
  name      = "example-origin-pool"
  namespace = "staging"
}

output "origin_pool_id" {
  value = data.xcsh_origin_pool.example.id
}
```

<a id="canonical-62e9a281c0603aaae665f79dd6ef7204928383b9291f662b7785ca4f0542628a"></a>

## Root configuration — xcsh_origin_pool / 726ce2e68a36 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-c7de3d97e2a4201b18359fd1745dd05820ce9ccfbc26dca45c425bb07d5a8275"></a>

## Next pages — xcsh_origin_pool / 726ce2e68a36 / 6

- [Property reference](../guides/data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [Examples](../guides/data-sources--origin_pool--examples--group-001.md#canonical-26f5d1de6dbe2147a167989881c2a2c0a1193c72d01e2342c928c8e90c063e5e)
