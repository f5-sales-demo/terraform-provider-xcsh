---
page_title: "xcsh_bot_peer_status landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_status landing."
---

# xcsh_bot_peer_status landing

<a id="canonical-1213323001223033-2102310131020031-2130002023320333-3233002313233120-0102113233101212-0211233322133022-3213211320322133-3330331011313210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110302101223120-2222032210122020-3213111031122022-0233010311310221-3330130031313011-0100211320122110-1332123122000011-3203032320030032"></a>

## xcsh_bot_peer_status — xcsh_bot_peer_status / 300200330210 / 2

Breadcrumbs:

- xcsh_bot_peer_status

Resource creation operation.

<a id="canonical-2100301003232001-1112313202202022-1133011220030020-0112313120132021-0023301222030112-3223331100312233-1310020331133100-3322333111312203"></a>

## Prerequisites — xcsh_bot_peer_status / 300200330210 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3323110333013023-3010232002223010-3212132212022211-3030102010010331-2301023302120100-0113011003301300-0013233210212020-2023031300213320"></a>

## Minimal configuration — xcsh_bot_peer_status / 300200330210 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotPeerStatus DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_status" "example" {
  namespace = "example-value"
}

output "bot_peer_status_result" {
  value = data.xcsh_bot_peer_status.example
}
```

<a id="canonical-0331110323233131-2302101020210130-3102311202311130-2230110103013331-0012011110122332-3230220303210311-2132120312302213-3111321233220302"></a>

## Root configuration — xcsh_bot_peer_status / 300200330210 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1030321022320000-3000222300221332-1311333030332321-1120333203023103-1321123131013220-1330231002002113-1323212102230330-1010202001030331"></a>

## Next pages — xcsh_bot_peer_status / 300200330210 / 6

- [Property reference](../guides/data-sources--bot_peer_status--reference--group-001.md#canonical-2233112331322323-3212110322020321-3202221110020001-3200220303102201-3002320020113100-1010310321312301-2132001113233301-3312110023102211)
- [Examples](../guides/data-sources--bot_peer_status--examples--group-001.md#canonical-1013303221323110-1301112223102022-3233003023303032-3303121320330131-1103031223202233-0201233102030332-0223000332130021-1011032031033213)
