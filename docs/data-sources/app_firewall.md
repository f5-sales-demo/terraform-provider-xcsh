---
page_title: "xcsh_app_firewall"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_app_firewall."
---

# xcsh_app_firewall

<a id="canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_app_firewall

Reads App Firewall information from F5 Distributed Cloud.

<a id="canonical-2001301102011232-2101232311220033-1330130031031213-1323110313232333-3333112002220332-1000001230101112-1012311123120201-3110302111223213"></a>

### Prerequisites for `xcsh_app_firewall`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `service_policy`.

- service_policy: Fine-grained access control rules

<a id="canonical-1001301300220312-3023010133133312-2023100020300000-0312133021302003-3123230222111320-3100003201012311-0101213320213023-1302003023010203"></a>

### Minimal configuration for `xcsh_app_firewall`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppFirewall Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppFirewall by name
data "xcsh_app_firewall" "example" {
  name      = "example-app-firewall"
  namespace = "staging"
}

output "app_firewall_id" {
  value = data.xcsh_app_firewall.example.id
}
```

<a id="canonical-2032220221012033-1120302311103201-3203033231102333-0301220233302012-1133332002002233-0320131303231132-1330033201031212-1130113001110320"></a>

### Root configuration for `xcsh_app_firewall`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1322302021233133-2032333221220102-3020212120133212-2233101222001330-1130323122220012-0030321030120133-2233010330312113-2210312113211311"></a>

### Explore this collection for `xcsh_app_firewall`

- [Property reference](../guides/data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [Examples](../guides/data-sources--app_firewall--examples--group-001.md#canonical-3131132132112300-2031210311300023-1320230303032133-2101121211031120-0000302131110002-0122121020210203-3302311013131221-3301311322210103)
