---
page_title: "xcsh_registration landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration landing."
---

# xcsh_registration landing

<a id="canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-991b03911b429f8f8d3dec0970085ccf53d3e6298f8dfd26484bc5811cf6741a"></a>

## xcsh_registration — xcsh_registration / 6d34c7d60af5 / 2

Breadcrumbs:

- xcsh_registration

Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this
message, never used by users. configuration.

<a id="canonical-7f77860a6e7a820408af4138f3297bc12bdd29f9961562dbd30b28bfa64966e5"></a>

## Prerequisites — xcsh_registration / 6d34c7d60af5 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1d57b31a76e4d1d75632571f5fa65a9311eb19c4d56c255c3d9c8c69522ecfa0"></a>

## Minimal configuration — xcsh_registration / 6d34c7d60af5 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Registration Resource Example
# Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this message, never used by users.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Registration configuration
resource "xcsh_registration" "example" {
  name      = "example-registration"
  namespace = "staging"

  token = "example-value"
}
```

<a id="canonical-e173f76419040e33a08228a24721897bc9bcaf8bd7fc98047cfcb94c1aef8908"></a>

## Root configuration — xcsh_registration / 6d34c7d60af5 / 5

Required root properties: `name`, `namespace`, `token`. Full root flags and choices appear in the property reference.

<a id="canonical-c573f0b04b9fc520f766d6174559a039d31f15322e3dc81d3ff2b7b42e50106d"></a>

## Next pages — xcsh_registration / 6d34c7d60af5 / 6

- [Property reference](../guides/resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [Examples](../guides/resources--registration--examples--group-001.md#canonical-459f2348c9bd2f11d18baccc06df68996e9613672dd6ffd94092fc45ee78b9e3)
- [Import](../guides/resources--registration--lifecycle--group-001.md#canonical-039ec89044403400e7321ea26a166c144c0265deae604680a40701d225dca629)
- [Timeouts](../guides/resources--registration--lifecycle--group-001.md#canonical-cef96a04ebdfbf3bc60414c3c2c0b6895e9976289716696ce688560222a21e5f)
