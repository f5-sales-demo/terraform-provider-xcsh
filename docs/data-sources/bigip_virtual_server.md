---
page_title: "xcsh_bigip_virtual_server"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_virtual_server."
---

# xcsh_bigip_virtual_server

<a id="canonical-1131322003302203-3133232130331303-3111132332133231-2012331201133203-0022310012333001-0003330233211211-0333312201021000-3230332012332020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_bigip_virtual_server

Reads BIG-IP Virtual Server information from F5 Distributed Cloud.

<a id="canonical-0210313202101202-1132323200111113-1003010021301132-0321330330322311-1322210213132130-1022310222220211-1230212212332300-1030003230222021"></a>

### Prerequisites for `xcsh_bigip_virtual_server`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0220322213000110-3311323012120201-3121200232231001-2020020220300133-1023121113213003-3013001021332010-1023310222312331-3113103302032122"></a>

### Minimal configuration for `xcsh_bigip_virtual_server`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BigIPVirtualServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BigIPVirtualServer by name
data "xcsh_bigip_virtual_server" "example" {
  name      = "example-bigip-virtual-server"
  namespace = "staging"
}

output "bigip_virtual_server_id" {
  value = data.xcsh_bigip_virtual_server.example.id
}
```

<a id="canonical-2002301130103223-2132230311030200-1103003303011100-3220022103001023-0200323122312212-3222312333300222-0001231022132332-0302012013321100"></a>

### Root configuration for `xcsh_bigip_virtual_server`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3203100322120320-0230301323221033-3103122112202311-0201221223313022-0033320322101321-0132230022131231-1320113120313220-3212203013300121"></a>

### Explore this collection for `xcsh_bigip_virtual_server`

- [Property reference](../guides/data-sources--bigip_virtual_server--reference--group-001.md#canonical-0330103300101033-0110313212121312-2131323221121012-0122222202113110-1131132021023000-3330131201331123-1312102120012011-0302013001211333)
- [Examples](../guides/data-sources--bigip_virtual_server--examples--group-001.md#canonical-2000033203001131-3130300110310212-0300112012011003-2122112212020211-2211220000232300-3001213332210023-3101102332301333-2331331103233121)
