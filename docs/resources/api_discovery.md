---
page_title: "xcsh_api_discovery landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_discovery landing."
---

# xcsh_api_discovery landing

<a id="canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03400f7f027ae9a70eca0fa98fa5fc0e06cd4db117954d4613195900112e58a5"></a>

## xcsh_api_discovery — xcsh_api_discovery / f56936e36999 / 2

Breadcrumbs:

- xcsh_api_discovery

Manages API discovery creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-e280e62362ceba01c7896d1d56ba70b85ca77ca18bea198ffc15ab8a724b2c1a"></a>

## Prerequisites — xcsh_api_discovery / f56936e36999 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-90f94afce14ad453fce18dbced02610a0373c02eda3e8f6462502735a02f7a6f"></a>

## Minimal configuration — xcsh_api_discovery / f56936e36999 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APIDiscovery Resource Example
# Manages API discovery creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APIDiscovery configuration
resource "xcsh_api_discovery" "example" {
  name      = "example-api-discovery"
  namespace = "staging"
}
```

<a id="canonical-3349fdef81cead2c2baa7b10c842411fd818836acaec16ff61c1a4c1345d9f68"></a>

## Root configuration — xcsh_api_discovery / f56936e36999 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-a7a470df8194f2f246deb345d86d0a13e45a81a2624b9546d01a2e5f68f32f47"></a>

## Next pages — xcsh_api_discovery / f56936e36999 / 6

- [Property reference](../guides/resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [Examples](../guides/resources--api_discovery--examples--group-001.md#canonical-93f0346c60bede5e7e23ca2c25083fbc71af112c5b402b831fcc7c6ea64733a4)
- [Import](../guides/resources--api_discovery--lifecycle--group-001.md#canonical-75d7cc9e6fedb06de3d4ad194cc8a9b7d5bd9d541f44dd5efa44042fff6fc79d)
- [Timeouts](../guides/resources--api_discovery--lifecycle--group-001.md#canonical-0709184a68e01a860b032b1c5564ee0183d91f635046e59c85a8cc806e41e4ad)
