---
page_title: "xcsh_advertise_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_advertise_policy landing."
---

# xcsh_advertise_policy landing

<a id="canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cece5b434f95bfb64b6cf07e59af76375a6866793d02f9425abc26d261c9916f"></a>

## xcsh_advertise_policy — xcsh_advertise_policy / 565db87b456f / 2

Breadcrumbs:

- xcsh_advertise_policy

Manages a Advertise Policy resource in F5 Distributed Cloud for advertise\_policy object controls
how and where a service represented by a given virtual\_host object is advertised to consumers.
configuration.

<a id="canonical-4cc3c23b9ea7ba706a592e201df78c80953e9dba98dfd2ca6dfc4a2e25ed3b9f"></a>

## Prerequisites — xcsh_advertise_policy / 565db87b456f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-770973559abac46afef83e0f52ae1ea0380fdc33f4c1cdf1b399a947381ef4e0"></a>

## Minimal configuration — xcsh_advertise_policy / 565db87b456f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AdvertisePolicy Resource Example
# Manages a Advertise Policy resource in F5 Distributed Cloud for advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AdvertisePolicy configuration
resource "xcsh_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "staging"
}
```

<a id="canonical-673f0876b4f9594da3c7adea3ab37e6e607a95040545d62e0c88bb3133879f3d"></a>

## Root configuration — xcsh_advertise_policy / 565db87b456f / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-6f32fe7de92f7d79b28563dc290caab02489740b420065d9fff987a01eef6daf"></a>

## Next pages — xcsh_advertise_policy / 565db87b456f / 6

- [Property reference](../guides/resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [Examples](../guides/resources--advertise_policy--examples--group-001.md#canonical-6a3b267374c676ae36bae3d8f4ad9b293170f2b9387c9e12c907fa0a47d6f5eb)
- [Import](../guides/resources--advertise_policy--lifecycle--group-001.md#canonical-e8762eed9866c7d9b818c0b610031bc1cb3884723df97984e623ebd057ca215d)
- [Timeouts](../guides/resources--advertise_policy--lifecycle--group-001.md#canonical-718e0bc0254f3269f39143ee0424f73f46eaa95940030d8cc8eeefc71002fe68)
