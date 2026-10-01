---
page_title: "xcsh_filter_set landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_filter_set landing."
---

# xcsh_filter_set landing

<a id="canonical-3033322202313333-3021122111030012-3023212212330123-0300102111110132-3312003112301221-1222013003030021-1210202002210210-0310110313101001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023231212121230-1302320111033120-2002332002210322-1112020011000030-2330323232000010-1001110103002021-3213202303013213-2100110222201210"></a>

## xcsh_filter_set — xcsh_filter_set / 110030120332 / 2

Breadcrumbs:

- xcsh_filter_set

Manages specification in F5 Distributed Cloud.

<a id="canonical-1230001003220020-3032133133210211-2221001322110003-1012110021313233-3110223021013223-0101021212033210-0222321302313013-2120003021332031"></a>

## Prerequisites — xcsh_filter_set / 110030120332 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2223100301303023-0330213023110303-1323031131201021-1222310110312211-0131103002022311-2230222301311232-2013333122023110-0232221013133222"></a>

## Minimal configuration — xcsh_filter_set / 110030120332 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FilterSet Resource Example
# Manages specification in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FilterSet configuration
resource "xcsh_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "staging"

  context_key = "example-value"
}
```

<a id="canonical-3123002301323302-0300220331033020-0233223213203312-3110110323110030-3032221021201301-3131201032130101-1023301013210030-0201203101110013"></a>

## Root configuration — xcsh_filter_set / 110030120332 / 5

Required root properties: `context_key`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1213212022230102-0031223021202021-0231002011120102-1122222231211001-1300001021001331-0120201203322032-0132310211002331-3310030000223201"></a>

## Next pages — xcsh_filter_set / 110030120332 / 6

- [Property reference](../guides/resources--filter_set--reference--group-001.md#canonical-1023200121033312-1222030021033022-2031210010033023-1200303020100031-1303323030030032-3011012223301222-1332302023110212-1031302002022311)
- [Examples](../guides/resources--filter_set--examples--group-001.md#canonical-3222211130023131-3030002002322101-1200030222123000-1200233111210312-2220111122231232-1122302111201322-1223301011102233-3121310201130203)
- [Import](../guides/resources--filter_set--lifecycle--group-001.md#canonical-3010011210203102-2031200311101231-2000222222323212-3302211012122000-3002113302232030-1223210112200021-2233310221322022-1301000203202301)
- [Timeouts](../guides/resources--filter_set--lifecycle--group-001.md#canonical-0300231211200110-1220203033103123-3032302320331103-1023311021103302-3010212110001111-3310011212132303-0100220102110133-3003300330332120)
