---
page_title: "xcsh_bot_peer_traffic_overview examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_traffic_overview examples."
---

# xcsh_bot_peer_traffic_overview examples

<a id="canonical-0322022003103201-2112122103103021-0101101112323221-3332001032101210-2033223120303102-2200003130331002-1020101003213103-2310223133303303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bot_peer_traffic_overview](../data-sources/bot_peer_traffic_overview.md#canonical-1123001103021232-1031112101301311-0233312222211013-0202321023223122-3222120111310211-0001130132131101-3010132202021212-0221102121013113)
- Examples

<a id="canonical-3003021221332303-3032020313013101-1332113311201303-1132112220210211-1310000312123032-3000021112223002-1021012210320103-0130201101023233"></a>

### Complete configurations for `xcsh_bot_peer_traffic_overview`

- [Data source](data-sources--bot_peer_traffic_overview--examples--group-001.md#canonical-2021013333003013-2022132231103003-2110023121003113-0331223002011203-3131000012302013-2022221320002110-2200322122122320-0112033001031231): valid configuration.

<a id="canonical-2021013333003013-2022132231103003-2110023121003113-0331223002011203-3131000012302013-2022221320002110-2200322122122320-0112033001031231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bot_peer_traffic_overview](../data-sources/bot_peer_traffic_overview.md#canonical-1123001103021232-1031112101301311-0233312222211013-0202321023223122-3222120111310211-0001130132131101-3010132202021212-0221102121013113)
- [Examples](data-sources--bot_peer_traffic_overview--examples--group-001.md#canonical-0322022003103201-2112122103103021-0101101112323221-3332001032101210-2033223120303102-2200003130331002-1020101003213103-2310223133303303)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_traffic_overview/data-source.tf`; digest `sha256:02ae4e984486512606c527da0d434568ac636c951e2a90ff31f351d9dd36be11`.

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
