---
page_title: "xcsh_certified_hardware"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_certified_hardware."
---

# xcsh_certified_hardware

<a id="canonical-3203100033102110-1200003111311001-2130133202021111-2000112331003330-0223333332201310-0221311220013120-2000203133203030-3221223311122122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_certified_hardware

Reads Certified Hardware information from F5 Distributed Cloud.

<a id="canonical-0213003231210213-3001212022010311-2103010212122301-2023301203100320-1103130222120310-2220211003331203-0320110321120020-3123311231012112"></a>

### Prerequisites for `xcsh_certified_hardware`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0320301303133231-0100000012201200-3312132101112132-2301201303120030-0312301030302111-3012233101300303-1003213112230010-1212313202001010"></a>

### Minimal configuration for `xcsh_certified_hardware`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CertifiedHardware Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CertifiedHardware by name
data "xcsh_certified_hardware" "example" {
  name      = "example-certified-hardware"
  namespace = "staging"
}

output "certified_hardware_id" {
  value = data.xcsh_certified_hardware.example.id
}
```

<a id="canonical-2301210310302023-2030030033232111-3220023200301220-3213123100311232-2112030222131001-2133210311310211-1311332001130023-1133032020300220"></a>

### Root configuration for `xcsh_certified_hardware`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1223212011133022-1131020001020113-2223112001212200-2332111213031032-2131111130033021-2101003103023112-1122203301113223-3213012023102030"></a>

### Explore this collection for `xcsh_certified_hardware`

- [Property reference](../guides/data-sources--certified_hardware--reference--group-001.md#canonical-0302130231202120-2203110233320123-1120332022311120-3011032230213311-2103323001103203-0212122123333022-1013300331312111-0101311012103201)
- [Examples](../guides/data-sources--certified_hardware--examples--group-001.md#canonical-3211013330020133-2220002032011211-2213232320113221-0321132313332230-1131300330100223-1100020132102113-1120013311233100-2101331022200313)
