---
page_title: "xcsh_network_firewall"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_firewall."
---

# xcsh_network_firewall

<a id="canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_firewall

Manages a Network Firewall resource in F5 Distributed Cloud for network firewall is created by users
in system namespace. configuration.

<a id="canonical-1021103002011230-1111103230332102-2100022012332223-3110103200220233-1002030022003112-2102330033032023-2311331311013302-1120331010133312"></a>

### Prerequisites for `xcsh_network_firewall`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-2232323301320100-0010223331203203-1320221321210322-3231300231033302-1010112221310330-1100103031331300-1333013023122033-0132322301023233"></a>

### Minimal configuration for `xcsh_network_firewall`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkFirewall Resource Example
# Manages a Network Firewall resource in F5 Distributed Cloud for network firewall is created by users in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkFirewall configuration
resource "xcsh_network_firewall" "example" {
  name      = "example-network-firewall"
  namespace = "system"
}
```

<a id="canonical-1030023210320210-3300332300031232-1232100221331113-1111021011223300-1021000211312322-3210322330131330-3320213233000032-2202321303013312"></a>

### Root configuration for `xcsh_network_firewall`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1022021033221112-0231321331003303-3001103313012230-0030222020332123-0210101203202020-3123123030322101-3100011113012233-0312132033123213"></a>

### Explore this collection for `xcsh_network_firewall`

- [Property reference](../guides/resources--network_firewall--reference--group-001.md#canonical-1131321320323100-0133323001330022-1211003101102003-2000211000122312-2031221212010020-2323201312033210-3331113002030021-0010133133330010)
- [Examples](../guides/resources--network_firewall--examples--group-001.md#canonical-1031032103323200-0122100200112332-1320330020112301-3331011310010120-3320010130103103-2232331310312302-0023100313031112-2201332200222201)
- [Import](../guides/resources--network_firewall--lifecycle--group-001.md#canonical-0332032223301202-2101103111010232-0221122011112100-2222113100031120-3212302122222311-2300100022022233-0032223123323302-3012203132031203)
- [Timeouts](../guides/resources--network_firewall--lifecycle--group-001.md#canonical-1213033313223021-2211131133110302-3320210111131202-0101220223232111-2130122331002203-1030202230301312-2300222223333031-3211011232232230)
