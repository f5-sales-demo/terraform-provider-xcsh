---
page_title: "xcsh_bot_peer_traffic_overview"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_traffic_overview."
---

# xcsh_bot_peer_traffic_overview

<a id="canonical-1123001103021232-1031112101301311-0233312222211013-0202321023223122-3222120111310211-0001130132131101-3010132202021212-0221102121013113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_bot_peer_traffic_overview

Reads Bot Peer Traffic Overview information from F5 Distributed Cloud.

<a id="canonical-2032231300010000-1332212222322003-3030123101311121-0111020321203020-3023232331021332-2203221131332011-1211312303001202-3322010023031033"></a>

### Prerequisites for `xcsh_bot_peer_traffic_overview`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2102021101013103-2313101310222033-2031033112322310-0331032302102031-0330101302030100-2103121132021233-2120233012030213-2202313323313220"></a>

### Minimal configuration for `xcsh_bot_peer_traffic_overview`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotPeerTrafficOverview DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_traffic_overview" "example" {
  namespace = "example-value"
}

output "bot_peer_traffic_overview_result" {
  value = data.xcsh_bot_peer_traffic_overview.example
}
```

<a id="canonical-0301003111000001-3333312032022130-2303211313320311-2003103232011233-1112210011011203-1001030303330202-2032111031322131-3022212002333021"></a>

### Root configuration for `xcsh_bot_peer_traffic_overview`

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3112213113321021-1100012023322032-3001023112320033-1121100323022130-2031212311103233-0223330132110210-0203230010013030-2113130032103000"></a>

### Explore this collection for `xcsh_bot_peer_traffic_overview`

- [Property reference](../guides/data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-1103031213222001-1003122002030331-1121300031120000-1103232121210320-1003000301131102-3132201013033132-1001310031332103-2310031113130322)
- [Examples](../guides/data-sources--bot_peer_traffic_overview--examples--group-001.md#canonical-0322022003103201-2112122103103021-0101101112323221-3332001032101210-2033223120303102-2200003130331002-1020101003213103-2310223133303303)
