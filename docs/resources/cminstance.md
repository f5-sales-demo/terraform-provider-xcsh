---
page_title: "xcsh_cminstance landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cminstance landing."
---

# xcsh_cminstance landing

<a id="canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6045abac7ef7a231d0e5962a9b9c3d2a32bcaeeb9ec1c491a38d1442e30800fb"></a>

## xcsh_cminstance — xcsh_cminstance / 94e1c5389bf6 / 2

Breadcrumbs:

- xcsh_cminstance

Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed
Cloud.

<a id="canonical-c932a88d36fa415edca6227e25eb859337a1cfef075ba01bbe1407b30110da80"></a>

## Prerequisites — xcsh_cminstance / 94e1c5389bf6 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-6a27a8b54c7a3f80c80bef20e55dcfaf6b755b6f86f4d8780a3d62981ef6e56c"></a>

## Minimal configuration — xcsh_cminstance / 94e1c5389bf6 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Cminstance Resource Example
# Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Cminstance configuration
resource "xcsh_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "staging"

  port     = 1
  username = "example-value"
}
```

<a id="canonical-76c1c0da06d772b683f20d73621f79bd95bed3456910c33d21c16cd8d3ba936f"></a>

## Root configuration — xcsh_cminstance / 94e1c5389bf6 / 5

Required root properties: `name`, `namespace`, `port`, `username`. Full root flags and choices appear in the property reference.

<a id="canonical-6dfb788468128ce512f3b49e469e0ac011b5451c5407c5938e17a056505e7d7e"></a>

## Next pages — xcsh_cminstance / 94e1c5389bf6 / 6

- [Property reference](../guides/resources--cminstance--reference--group-001.md#canonical-1a3303bad1df452da182bfce55ae7d71e5cba24da48067fcf379e36045f7be59)
- [Examples](../guides/resources--cminstance--examples--group-001.md#canonical-65ac32d0660495d526e5fd75556d24254c3381e30194324ea5c2a80aebbd6132)
- [Import](../guides/resources--cminstance--lifecycle--group-001.md#canonical-9be4b50ce32bb76c686326dd4c2966143ba3d18eabbe35e9382f9f34ca901891)
- [Timeouts](../guides/resources--cminstance--lifecycle--group-001.md#canonical-90f38ea18892824f2ee3f56a70070beececdda85f32436767baee54764d243bd)
