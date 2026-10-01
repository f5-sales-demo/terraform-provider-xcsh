---
page_title: "xcsh_bot_peer_top_good_bots landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_top_good_bots landing."
---

# xcsh_bot_peer_top_good_bots landing

<a id="canonical-3030001313332130-0200112302220223-2231210320202203-3301330023332321-1010113230110223-1130212301220213-3012200130111102-0020213232203302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112102322103332-0001101320011231-2023211320200012-3331120210322032-3130233113303122-3111222130223211-2213122210122100-1121312002300020"></a>

## xcsh_bot_peer_top_good_bots — xcsh_bot_peer_top_good_bots / 221012013200 / 2

Breadcrumbs:

- xcsh_bot_peer_top_good_bots

Bot detection and defense configuration.

<a id="canonical-0212122003330320-3201021233031311-2220023312301301-3102103210332313-2331221113012221-1213112103111311-2022311013113223-3101332323030201"></a>

## Prerequisites — xcsh_bot_peer_top_good_bots / 221012013200 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2032120223033213-3103131110121103-0332011120003200-0102230113032110-3320300322323220-0332231112310121-3202330002212010-2021122012203111"></a>

## Minimal configuration — xcsh_bot_peer_top_good_bots / 221012013200 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotPeerTopGoodBots DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_top_good_bots" "example" {
  namespace = "example-value"
}

output "bot_peer_top_good_bots_result" {
  value = data.xcsh_bot_peer_top_good_bots.example
}
```

<a id="canonical-2131000121022223-2232101212210031-2332310313123310-1223311032302300-2000121320201231-1201302021300110-3011031322012333-1112323020102012"></a>

## Root configuration — xcsh_bot_peer_top_good_bots / 221012013200 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2222131111332001-3221232101231030-2102102201023131-2113220221212103-0113113201031121-0011320223311232-0300302033100123-1111210130221210"></a>

## Next pages — xcsh_bot_peer_top_good_bots / 221012013200 / 6

- [Property reference](../guides/data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-1222302012121320-3310003310000213-1123332130111312-0011233021231032-0101303313230300-0003312003231211-2302311133103331-3323312201112032)
- [Examples](../guides/data-sources--bot_peer_top_good_bots--examples--group-001.md#canonical-3123030123000202-0322300120230232-0200003123301123-1221112202031331-0211032313011311-1310102100131213-3313302112000322-3130220133313121)
