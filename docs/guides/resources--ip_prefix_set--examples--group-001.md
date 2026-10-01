---
page_title: "xcsh_ip_prefix_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ip_prefix_set examples."
---

# xcsh_ip_prefix_set examples

<a id="canonical-663015d63e3a93116ac51fd9a7e4beaa6cb87aac18bd9ffbdc01f28309a6c55e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11e05b2fbb62878752e0dd335198fec9a63463d2a5fc484bfc22a15e8c6f4486"></a>

## Examples — Examples / 90fe2d9ce1e8 / 2

Breadcrumbs:

- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-5a49cde9dc0a3a39ff028cd9a8304b00e480fb0bbedc342df01389fd1e3f30ae)
- Examples

<a id="canonical-34bc205e12d0965ec47e9105cd7e677bb23d1c57b9e5c5763616efe31c04d47b"></a>

## Complete configurations — Examples / 90fe2d9ce1e8 / 3

- [Resource](resources--ip_prefix_set--examples--group-001.md#canonical-3f448c83e320ea39c841066d8d7727f12c37b3956a440a3e53118f75dc1f777b): valid configuration.

<a id="canonical-5db8af32ea8e6838c83ff4b0fb2eb24ab1420d0bd406b983a1b4ba593f926f8b"></a>

## Next pages — Examples / 90fe2d9ce1e8 / 4

- [Resource](resources--ip_prefix_set--examples--group-001.md#canonical-3f448c83e320ea39c841066d8d7727f12c37b3956a440a3e53118f75dc1f777b)
- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-5a49cde9dc0a3a39ff028cd9a8304b00e480fb0bbedc342df01389fd1e3f30ae)

<a id="canonical-3f448c83e320ea39c841066d8d7727f12c37b3956a440a3e53118f75dc1f777b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e827cae74715a4bd808ca6e310c9b13e82d2cac736cdd581af9a77a662ceafb"></a>

## Resource — Resource / 69901996e497 / 2

Breadcrumbs:

- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-5a49cde9dc0a3a39ff028cd9a8304b00e480fb0bbedc342df01389fd1e3f30ae)
- [Examples](resources--ip_prefix_set--examples--group-001.md#canonical-663015d63e3a93116ac51fd9a7e4beaa6cb87aac18bd9ffbdc01f28309a6c55e)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ip_prefix_set/resource.tf`; digest `sha256:00f8d9b11740968b5c7c6ae54585108409034d987dc62393dd9ee54d031d3c3a`.

```terraform
# IPPrefixSet Resource Example
# Manages ip_prefix_set creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic IPPrefixSet configuration
resource "xcsh_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "staging"
}
```

<a id="canonical-0277dadc2f7de71757d2bfd449fda90b855f4f2eb2f1c9efa641c96fff1c1956"></a>

## Next pages — Resource / 69901996e497 / 3

- [Examples](resources--ip_prefix_set--examples--group-001.md#canonical-663015d63e3a93116ac51fd9a7e4beaa6cb87aac18bd9ffbdc01f28309a6c55e)
- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-5a49cde9dc0a3a39ff028cd9a8304b00e480fb0bbedc342df01389fd1e3f30ae)
