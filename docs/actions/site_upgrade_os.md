---
page_title: "xcsh_site_upgrade_os landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_upgrade_os landing."
---

# xcsh_site_upgrade_os landing

<a id="canonical-094cbfe1e9c98c0b0352113f3cdf16dc1917d92802e806ed418e980550d32637"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e89a78dbfce7a837e1598b76a82e9b65a56ce80c8f8dfc3764ba533d5b8c06e6"></a>

## xcsh_site_upgrade_os — xcsh_site_upgrade_os / da0d9a8ee468 / 2

Breadcrumbs:

- xcsh_site_upgrade_os

Request an in-place site operating-system upgrade.

<a id="canonical-1275a20e948835c62daf9afe009c5ee0d0d5cdac339e7ce486cbbaeedb272859"></a>

## Prerequisites — xcsh_site_upgrade_os / da0d9a8ee468 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-80daae7aebb03a9a2b556b32f52354323798edcc3136a458ae88e350cd984673"></a>

## Minimal configuration — xcsh_site_upgrade_os / da0d9a8ee468 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteUpgradeOS Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_site_upgrade_os" "example" {
  config {
    site       = "example-value"
    os_version = "example-value"
  }
}
```

<a id="canonical-c542b20dbc4711e665bfcb2df79433ba2016c2eddd1f8fbe104018c0b82298b9"></a>

## Root configuration — xcsh_site_upgrade_os / da0d9a8ee468 / 5

Required root properties: `os_version`, `site`. Full root flags and choices appear in the property reference.

<a id="canonical-fb96cd661c858331b40096f26ce23bd64daea01270179a33e0e208cf47eef715"></a>

## Next pages — xcsh_site_upgrade_os / da0d9a8ee468 / 6

- [Property reference](../guides/actions--site_upgrade_os--reference--group-001.md#canonical-2fc5839fe4af09ff725b9e707b65161924e9564b90248891f32c46b1397b4921)
- [Examples](../guides/actions--site_upgrade_os--examples--group-001.md#canonical-01f1a7047468e631893088da94970716a9806f8b80faf254913c119178c591a2)
- [Lifecycle](../guides/actions--site_upgrade_os--lifecycle--group-001.md#canonical-028279c17021882b70ebec1f61e0cc1a398d0aff25c0456483a83acbae703e02)
