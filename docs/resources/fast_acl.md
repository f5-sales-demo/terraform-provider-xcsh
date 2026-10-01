---
page_title: "xcsh_fast_acl landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl landing."
---

# xcsh_fast_acl landing

<a id="canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2e1f44916aea7eae8097fdcf248aed96d513600009c22cfb221d348891a49f3"></a>

## xcsh_fast_acl — xcsh_fast_acl / df3ac57332a6 / 2

Breadcrumbs:

- xcsh_fast_acl

Manages object, object contains rules to protect site from denial of service It has
destination\{destination IP, destination port) and references to in F5 Distributed Cloud.

<a id="canonical-af171b5cc7a1278b089246e7b6c30fb29aa1610445f5fe3881d30cd476dd8b41"></a>

## Prerequisites — xcsh_fast_acl / df3ac57332a6 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-6083ef0bc20ef1901c803b98ec26e17105b4c3d86a6d1d6546166206f3ccdaf4"></a>

## Minimal configuration — xcsh_fast_acl / df3ac57332a6 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FastACL Resource Example
# Manages object, object contains rules to protect site from denial of service It has destination{destination IP, destination port) and references to in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FastACL configuration
resource "xcsh_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"
}
```

<a id="canonical-875f63ec15ef546e9abc2e486900e1b601eda4212b2ff0047de4e35dc46dfb31"></a>

## Root configuration — xcsh_fast_acl / df3ac57332a6 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-b3dc459c18676d02734eb6db0c390f65691a8f2853b37498394bb7faa42a245a"></a>

## Next pages — xcsh_fast_acl / df3ac57332a6 / 6

- [Property reference](../guides/resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [Examples](../guides/resources--fast_acl--examples--group-001.md#canonical-c6c405700be873f5bc2b525c10a682d3e4871c3edee55da231f7c317c410bf0d)
- [Import](../guides/resources--fast_acl--lifecycle--group-001.md#canonical-e2b522106843439734eade51e69bfcd93e3192dca375798ccd9a9d63f133dc56)
- [Timeouts](../guides/resources--fast_acl--lifecycle--group-001.md#canonical-d5686beae99a6b2c494c7a9ee03f7cb5679cbdcf1dc84dd3fb479538db95e256)
