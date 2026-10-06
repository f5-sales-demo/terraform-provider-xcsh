---
page_title: "xcsh_network_firewall"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_firewall."
---

# xcsh_network_firewall

<a id="canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_firewall

Reads a Network Firewall created in the system namespace.

<a id="canonical-2301101020230103-2133332003102311-2313330212132223-1222300320020112-1021232310131022-2033320333003211-1313100101123303-0220322121220333"></a>

### Prerequisites for `xcsh_network_firewall`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-0002321003312101-3312131122031021-0232133300033321-2030302013331213-1113201230103202-0133301211211110-3221233211323000-2020322210023300"></a>

### Minimal configuration for `xcsh_network_firewall`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkFirewall Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkFirewall by name
data "xcsh_network_firewall" "example" {
  name      = "example-network-firewall"
  namespace = "system"
}

output "network_firewall_id" {
  value = data.xcsh_network_firewall.example.id
}
```

<a id="canonical-1113231231021023-3103131132020211-0210131012231102-2201132323321221-1032000233001201-3021023113332021-2002120212313013-3313121023123201"></a>

### Root configuration for `xcsh_network_firewall`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-2211132032113301-3122313030233213-1311110103003220-2331223212022321-2000232301312301-3101130003130113-2013033210100101-0123202103102003"></a>

### Explore this collection for `xcsh_network_firewall`

- [Property reference](../guides/data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- [Examples](../guides/data-sources--network_firewall--examples--group-001.md#canonical-1212323312020232-0200032322020210-1210213001133311-3210113331033111-1131201322032200-2233231033030332-1022333021112211-3100311203333131)
