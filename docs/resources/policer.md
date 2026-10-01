---
page_title: "xcsh_policer landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policer landing."
---

# xcsh_policer landing

<a id="canonical-2122311132312033-0121013211002130-0311121323112230-0010131112301322-1312303210110212-3032113221002232-1112102011022312-1111003020123113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200023133211110-2220212101310300-2010003312311003-1213012031200030-3312102021132101-1211021120102332-0113202301300133-2231112021321322"></a>

## xcsh_policer — xcsh_policer / 011320130322 / 2

Breadcrumbs:

- xcsh_policer

Manages new policer with traffic rate limits in F5 Distributed Cloud.

<a id="canonical-0330323203000312-3123203213202311-1123221032320023-1103100023103201-3223111000322331-3111001023112021-0213223022030100-2033122301330101"></a>

## Prerequisites — xcsh_policer / 011320130322 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1203330311200003-3322020012032130-1130233013211300-0132102212110021-3031301222111302-3232323221020303-3021033020330333-0111213022110130"></a>

## Minimal configuration — xcsh_policer / 011320130322 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Policer Resource Example
# Manages new policer with traffic rate limits in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Policer configuration
resource "xcsh_policer" "example" {
  name      = "example-policer"
  namespace = "staging"

  burst_size                 = 1
  committed_information_rate = 1
}
```

<a id="canonical-0230111103120202-1120021203323203-1003101020220322-0222032212232103-0002332230011130-1033220303032133-3012331032031001-0220002102222322"></a>

## Root configuration — xcsh_policer / 011320130322 / 5

Required root properties: `burst_size`, `committed_information_rate`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1130022130311123-3311220222100320-3322200130320202-1212013031113303-3031120122330211-0210122113031333-2332133131123312-3112122003210100"></a>

## Next pages — xcsh_policer / 011320130322 / 6

- [Property reference](../guides/resources--policer--reference--group-001.md#canonical-0211330131013020-2310303133111013-2220323023232330-2223132030123020-2011122031201112-2310200001103220-1221101020202030-3311103200011300)
- [Examples](../guides/resources--policer--examples--group-001.md#canonical-1311202001001303-0121132302210321-3212302221232023-2331311030121133-2003102010101310-2221203033211131-3300210210103331-2100210312012321)
- [Import](../guides/resources--policer--lifecycle--group-001.md#canonical-1203212110220131-0233233101122320-0213000303323231-0320300032332300-1100131122030312-3033202030310222-0302021222001310-2322311101201010)
- [Timeouts](../guides/resources--policer--lifecycle--group-001.md#canonical-3022231033111130-2123221132203121-2331112202031112-2311031330310120-0330323002203313-2111123200132110-3001232302213013-1232033322113021)
