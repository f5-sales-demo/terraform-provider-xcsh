---
page_title: "xcsh_srv6_network_slice landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_srv6_network_slice landing."
---

# xcsh_srv6_network_slice landing

<a id="canonical-1003011003113030-2133000032112000-2110300131312112-0323202222213032-3010021101322220-3130233303030103-2300101230220121-0201100002310332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132111011232303-1322311313132001-0013223123223333-3303330202213110-1212231320313021-2102300030310212-3203211000113233-1213123221110031"></a>

## xcsh_srv6_network_slice — xcsh_srv6_network_slice / 003320231120 / 2

Breadcrumbs:

- xcsh_srv6_network_slice

Manages srv6\_network\_slice creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-1110202103233102-3032120321220001-0302223132012322-0331222122131110-1321300102323012-2333221132011332-2302200021133211-2302002023313110"></a>

## Prerequisites — xcsh_srv6_network_slice / 003320231120 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1013322322300213-2230213200030102-1200203300330213-3003333231101112-3011011210330112-3220013230213133-0332331003312331-1313300021201320"></a>

## Minimal configuration — xcsh_srv6_network_slice / 003320231120 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Srv6NetworkSlice Resource Example
# Manages srv6_network_slice creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Srv6NetworkSlice configuration
resource "xcsh_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"

  sid_prefixes = ["example-value"]
}
```

<a id="canonical-0022011210113233-2133202122321131-3301323002330313-0033031111023110-3210320311011033-0121023330312323-2013313023103312-3211010121110230"></a>

## Root configuration — xcsh_srv6_network_slice / 003320231120 / 5

Required root properties: `name`, `sid_prefixes`. Full root flags and choices appear in the property reference.

<a id="canonical-0112212301010321-2133111121211310-2021330033333232-0301130130101102-0030121333132111-0011211001122212-0323020032300303-1113233331200321"></a>

## Next pages — xcsh_srv6_network_slice / 003320231120 / 6

- [Property reference](../guides/resources--srv6_network_slice--reference--group-001.md#canonical-2221210201331300-0011233131200112-1301201030112212-2211123330221021-3333331233302122-0320123310310202-3200030231001021-3012033330201300)
- [Examples](../guides/resources--srv6_network_slice--examples--group-001.md#canonical-0202203011323022-1213210303303201-2130020311033333-2223212322220201-3001213012200323-2112301210033031-0033020330120230-0321223030213201)
- [Import](../guides/resources--srv6_network_slice--lifecycle--group-001.md#canonical-1021310230130031-3213001203220303-0331102103012303-2221020133203100-2002011032320203-1322311132031301-1133233012333132-0213122033020001)
- [Timeouts](../guides/resources--srv6_network_slice--lifecycle--group-001.md#canonical-1223110113002312-1031010031033221-0233012202230210-0133023021103021-0303132110310101-2102111223002221-2211001111220310-3122323133010000)
