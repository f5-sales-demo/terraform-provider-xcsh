---
page_title: "xcsh_network_customer_edge_defaults"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_customer_edge_defaults."
---

# xcsh_network_customer_edge_defaults

<a id="canonical-2030202230323022-0201110111033011-3112302323000311-1231130201223133-0123022101003012-2310121012022212-0120311221210122-2230332213331100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_customer_edge_defaults

Default DNS and NTP destinations for Customer Edge firewall rules. Values are bundled from the
pinned OpenAPI release; this data source performs no network request. Ports and traffic direction
are not encoded in the manifest.

<a id="canonical-2321231200300301-1333003313331000-0023021311200330-2133021010003131-0111303030222032-0310313313033021-2303202012201133-0001232130212122"></a>

### Prerequisites for `xcsh_network_customer_edge_defaults`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1301323210202133-1103302302033231-3330300301301020-3011100330213332-0012000202123131-0003223302220200-2121301003133320-3103030322322033"></a>

### Minimal configuration for `xcsh_network_customer_edge_defaults`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_customer_edge_defaults" "system_services" {}

output "customer_edge_default_egress" {
  value = {
    dns = {
      direction    = "egress"
      protocols    = ["udp", "tcp"]
      port         = 53
      destinations = data.xcsh_network_customer_edge_defaults.system_services.dns_servers
    }
    ntp = {
      direction    = "egress"
      protocols    = ["udp"]
      port         = 123
      destinations = data.xcsh_network_customer_edge_defaults.system_services.ntp_servers
    }
  }
}
```

<a id="canonical-0213112021003012-1301322030222000-1110031221113332-3010112220100321-3123320222322220-2012311213301310-2303310001322231-0302131202332302"></a>

### Root configuration for `xcsh_network_customer_edge_defaults`

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-2123233300300300-3122312011202331-2012213133123203-3003323133103120-3310021312221211-3031310121332123-1321311223201213-1321213330330011"></a>

### Explore this collection for `xcsh_network_customer_edge_defaults`

- [Property reference](../guides/data-sources--network_customer_edge_defaults--reference--group-001.md#canonical-2131013332023213-0010232033113021-0202033330201320-2220210231031321-0211222320311011-1100301313023322-0211102303121020-0101122212031210)
- [Examples](../guides/data-sources--network_customer_edge_defaults--examples--group-001.md#canonical-3122311111020020-0121221101322000-1332133023312001-2223023020210131-0201032210033320-1030130221333023-3011100312033110-3033021330133130)
