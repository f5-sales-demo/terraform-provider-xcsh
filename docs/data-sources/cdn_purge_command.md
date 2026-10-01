---
page_title: "xcsh_cdn_purge_command landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_purge_command landing."
---

# xcsh_cdn_purge_command landing

<a id="canonical-6df34c3c75523044e1456b37c5cc1e89af059047c2c92713d5005e82a8faf6a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-228618987d190490319777bbe9029891d61a8d77759445205a6bdd4c3c49f5f5"></a>

## xcsh_cdn_purge_command — xcsh_cdn_purge_command / c2b3ee328c49 / 2

Breadcrumbs:

- xcsh_cdn_purge_command

Manages a CDN Purge Command resource in F5 Distributed Cloud for cdn purge command specification.
configuration.

<a id="canonical-d86e6077b82452cc1fc04cc96f47a941eef43484503cb8c381b2dbb36556cdaf"></a>

## Prerequisites — xcsh_cdn_purge_command / c2b3ee328c49 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2a05ab6242216462f7533ec9151b5b23c617f2fdac8a3ce5e027e89db692474d"></a>

## Minimal configuration — xcsh_cdn_purge_command / c2b3ee328c49 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CDNPurgeCommand Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNPurgeCommand by name
data "xcsh_cdn_purge_command" "example" {
  name      = "example-cdn-purge-command"
  namespace = "staging"
}

output "cdn_purge_command_id" {
  value = data.xcsh_cdn_purge_command.example.id
}
```

<a id="canonical-ab5d2ef38ced56e457170fb1093eea6709cd3e4a7503a7512b4ee56befa3d5d0"></a>

## Root configuration — xcsh_cdn_purge_command / c2b3ee328c49 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-b9a7bf04b08635b6940041366e59ac315f6cf2b7ad094afe467a2b2193a794de"></a>

## Next pages — xcsh_cdn_purge_command / c2b3ee328c49 / 6

- [Property reference](../guides/data-sources--cdn_purge_command--reference--group-001.md#canonical-bfb3a1d5ef969fc07c8a01794300d9a760cd621bff3a35dfc4dfb8ef43f1ec0e)
- [Examples](../guides/data-sources--cdn_purge_command--examples--group-001.md#canonical-9a67f832bd5c4a7e9da92f7e6628c30730c414d355cb64c2770d045a240f87ee)
