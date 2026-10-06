---
page_title: "xcsh_site_upgrade_os"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_upgrade_os."
---

# xcsh_site_upgrade_os

<a id="canonical-0021103023333201-3221302120300023-0003110201010333-0330313301123130-0121011331210220-0002322000123231-1001203221200011-1100310302120313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_site_upgrade_os

Request an in-place site operating-system upgrade.

<a id="canonical-3220212213203123-3330321322200313-3201112120231312-2220023221231211-2211123032200030-2033203133300313-1210232211030331-1123203000123212"></a>

### Prerequisites for `xcsh_site_upgrade_os`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0102131122020032-2110202003113012-0231223321223332-0000213011323200-3100311130312230-0303213213303210-2012302323223232-3123021302201121"></a>

### Minimal configuration for `xcsh_site_upgrade_os`

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

<a id="canonical-2000312222321322-3223230003222122-0223111112230302-3311020311100302-0313212032313030-0301031222101120-2232202032031100-3031212010121303"></a>

### Root configuration for `xcsh_site_upgrade_os`

Required root properties: `os_version`, `site`. Full root flags and choices appear in the property reference.

<a id="canonical-3011100223020031-2330101301013212-1211233330230231-3313211003032322-0200011230023231-3131013320332332-0100100001203000-2320020221202321"></a>

### Explore this collection for `xcsh_site_upgrade_os`

- [Property reference](../guides/actions--site_upgrade_os--reference--group-001.md#canonical-0233301120032133-3210223300213333-1302112321321300-1323121101120121-0210322111121023-2100021020202101-3303023010122301-0321132310210201)
- [Examples](../guides/actions--site_upgrade_os--examples--group-001.md#canonical-0001330122130010-1310122032120301-2021030020203122-2110211300130112-2221200012332023-2000332233021110-2101033001012101-1320301121012202)
- [Lifecycle](../guides/actions--site_upgrade_os--lifecycle--group-001.md#canonical-0002200213213001-1300020120200223-1300322332300133-1201320030300122-0321203100223333-0211300010111210-2003222003223023-2232130003320002)
