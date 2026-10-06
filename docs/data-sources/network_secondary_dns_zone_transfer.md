---
page_title: "xcsh_network_secondary_dns_zone_transfer"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_secondary_dns_zone_transfer."
---

# xcsh_network_secondary_dns_zone_transfer

<a id="canonical-0211000112320000-3032222023300022-3233300332112210-3202010232211322-0100212322311011-2201301123313213-3323022110111321-3110203223320302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_secondary_dns_zone_transfer

Published Secondary DNS transfer and notify IPv4 addresses. The source does not distinguish their
purposes. Values are bundled from the pinned OpenAPI release; this data source performs no network
request. Ports and traffic direction are not encoded in the manifest.

<a id="canonical-3002013133030011-1001113302033220-3030131201100203-1131000311013221-1020000330232032-2110002200331311-3223101100211311-2000300313023020"></a>

### Prerequisites for `xcsh_network_secondary_dns_zone_transfer`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3032310321331123-1200100303121323-3101102332001032-3103230120220032-2121032330313302-0331323132302203-2011032222220130-0301323032200013"></a>

### Minimal configuration for `xcsh_network_secondary_dns_zone_transfer`

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

data "xcsh_network_secondary_dns_zone_transfer" "authoritative_dns" {}

# The manifest combines transfer and notify sources, so both explicit DNS
# rules use the same published allowlist.
output "secondary_dns_rules" {
  value = [
    {
      direction = "ingress"
      protocol  = "tcp"
      port      = 53
      sources   = data.xcsh_network_secondary_dns_zone_transfer.authoritative_dns.cidr_blocks
    },
    {
      direction = "ingress"
      protocol  = "udp"
      port      = 53
      sources   = data.xcsh_network_secondary_dns_zone_transfer.authoritative_dns.cidr_blocks
    },
  ]
}
```

<a id="canonical-1122222011030323-3013312311310233-2112300223230232-1011123103102331-1233233222303203-2233013232300110-0131300003111203-3031123103313132"></a>

### Root configuration for `xcsh_network_secondary_dns_zone_transfer`

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-2302010001021000-3312101322103121-0230122222123331-0333212000110100-2131232021000111-3323200112123300-0122123003120113-3031233312200321"></a>

### Explore this collection for `xcsh_network_secondary_dns_zone_transfer`

- [Property reference](../guides/data-sources--network_secondary_dns_zone_transfer--reference--group-001.md#canonical-3021101133310013-2330000022200301-3231032123101132-1112201032100203-2213103222230211-0010232310001312-0032012300130210-0333311032021320)
- [Examples](../guides/data-sources--network_secondary_dns_zone_transfer--examples--group-001.md#canonical-2330121313012102-0110302012330233-3333233112001102-1011223330131202-1111002221302211-0201120022000131-0130201201223323-0222223311310233)
