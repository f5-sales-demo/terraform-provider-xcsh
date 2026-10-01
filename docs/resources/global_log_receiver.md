---
page_title: "xcsh_global_log_receiver landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver landing."
---

# xcsh_global_log_receiver landing

<a id="canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b3d5eaf45b7113e2ea91b6b564da149ea04a7f2e6a8b42f529cf9109b3fb65b"></a>

## xcsh_global_log_receiver — xcsh_global_log_receiver / acfd4e516ac0 / 2

Breadcrumbs:

- xcsh_global_log_receiver

Manages new Global Log Receiver object in F5 Distributed Cloud.

<a id="canonical-b5be1c88739c5b102b32beeeb9a790d5c7c232dceb839dce1d5bc3515d83f6bd"></a>

## Prerequisites — xcsh_global_log_receiver / acfd4e516ac0 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-ada2358e62f26fe628fb7e70302710f5bbfdf1de1f6a32366a51bb7f9a1cd6c2"></a>

## Minimal configuration — xcsh_global_log_receiver / acfd4e516ac0 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# GlobalLogReceiver Resource Example
# Manages new Global Log Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic GlobalLogReceiver configuration
resource "xcsh_global_log_receiver" "example" {
  name      = "example-global-log-receiver"
  namespace = "staging"
}
```

<a id="canonical-5d480a62de965ea7dc31d2619efd86f918acb776af2398225c4fd3f046583309"></a>

## Root configuration — xcsh_global_log_receiver / acfd4e516ac0 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-c3e87aad10003ffc677e4c4a4ea0e8577a5268df53c9b83ef99292432692b6a4"></a>

## Next pages — xcsh_global_log_receiver / acfd4e516ac0 / 6

- [Property reference](../guides/resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [Examples](../guides/resources--global_log_receiver--examples--group-001.md#canonical-5c26fc0d1d2807236ae02af52707999e5e595a8aae44687b256637a6e9951aaf)
- [Import](../guides/resources--global_log_receiver--lifecycle--group-001.md#canonical-fd636c89e15c31c2420003286cdeffc263c9d5d51c7c91c1858f492c8a0eedaa)
- [Timeouts](../guides/resources--global_log_receiver--lifecycle--group-001.md#canonical-26a15a7a91bf182fca614bc5847ced9fcdc552d9e414a370ace5079ce25917d8)
