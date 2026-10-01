---
page_title: "xcsh_cdn_purge_command landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_purge_command landing."
---

# xcsh_cdn_purge_command landing

<a id="canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9145f3bee58c4f8d22f11ad706a24d553f0793af0bc9a04926d683e889531b25"></a>

## xcsh_cdn_purge_command — xcsh_cdn_purge_command / 47d69e90e57c / 2

Breadcrumbs:

- xcsh_cdn_purge_command

Manages a CDN Purge Command resource in F5 Distributed Cloud for cdn purge command specification.
configuration.

<a id="canonical-c7f895449af87586f83354e8e99a113c73a0550c32804fc9ceed23d574caddf5"></a>

## Prerequisites — xcsh_cdn_purge_command / 47d69e90e57c / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-652788b2c5a674dab633964fcbb68e6e436961dc73fee8f7464ced87111313eb"></a>

## Minimal configuration — xcsh_cdn_purge_command / 47d69e90e57c / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CDNPurgeCommand Resource Example
# Manages a CDN Purge Command resource in F5 Distributed Cloud for cdn purge command specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNPurgeCommand configuration
resource "xcsh_cdn_purge_command" "example" {
  name      = "example-cdn-purge-command"
  namespace = "staging"
}
```

<a id="canonical-86367379db918deb73e16cba167fcd3fb7ccfc920fe4fd4b3541b478f6b762fb"></a>

## Root configuration — xcsh_cdn_purge_command / 47d69e90e57c / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-6eef6cdb13c2cc1a660590a86656de60a7b1a33a5f86012af8693c65bdc57c16"></a>

## Next pages — xcsh_cdn_purge_command / 47d69e90e57c / 6

- [Property reference](../guides/resources--cdn_purge_command--reference--group-001.md#canonical-0d7ff38d62b870c2da94865b1a7e4f05b5f53d124a92ef6bcc79c52b5a816d32)
- [Examples](../guides/resources--cdn_purge_command--examples--group-001.md#canonical-67e622085485d9bee851ad3e118817e9c995ea98ce48604db20bdd0c540bacfe)
- [Import](../guides/resources--cdn_purge_command--lifecycle--group-001.md#canonical-f3319a8679df95d813cc6a0822f851b187726880c8aaa536048996eaf6c32c79)
- [Timeouts](../guides/resources--cdn_purge_command--lifecycle--group-001.md#canonical-13b19cc2b9cf89506b7e83acfcfd7eb70c50670975026bf9f0b60bfef63dfd10)
