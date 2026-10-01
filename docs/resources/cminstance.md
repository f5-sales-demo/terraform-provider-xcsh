---
page_title: "xcsh_cminstance landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cminstance landing."
---

# xcsh_cminstance landing

<a id="canonical-3030301003002223-1121312121232112-1232303010012000-0213330232003222-0030121133101123-1130200110322121-0013030301203132-0203031312323211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200101122232230-1332331322020301-3100321121120222-2123213003310222-0302233022323223-2132300130102101-2203203101101002-3203002000003323"></a>

## xcsh_cminstance — xcsh_cminstance / 032021233312 / 2

Breadcrumbs:

- xcsh_cminstance

Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed
Cloud.

<a id="canonical-3021030222202031-0312332210011132-3130221202021332-0211322320112103-0313220130333233-0013112322000123-2332011000132303-0001010031222000"></a>

## Prerequisites — xcsh_cminstance / 032021233312 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1222021322202311-1030132203332000-3020002332330200-3211113130332233-1223131111231233-2012331031201320-0022033112022120-0132331232111230"></a>

## Minimal configuration — xcsh_cminstance / 032021233312 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Cminstance Resource Example
# Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Cminstance configuration
resource "xcsh_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "staging"

  port     = 1
  username = "example-value"
}
```

<a id="canonical-1312300130003122-0012311313022312-2003330200311303-1202013313212331-2111233231031011-1221010030030331-0201300112303120-3103232221031233"></a>

## Root configuration — xcsh_cminstance / 032021233312 / 5

Required root properties: `name`, `namespace`, `port`, `username`. Full root flags and choices appear in the property reference.

<a id="canonical-1231332313202010-1220010220303211-0102330323102132-1012213200223000-0101231110110130-1110001330112103-2032011322001112-1100113213311332"></a>

## Next pages — xcsh_cminstance / 032021233312 / 6

- [Property reference](../guides/resources--cminstance--reference--group-001.md#canonical-0122030300032322-3101313310110231-2201200223333032-1111223213311301-3211302322021031-2210200012133330-3303132132031200-1011331323321121)
- [Examples](../guides/resources--cminstance--examples--group-001.md#canonical-1211223003023100-1212001021113111-0212321133311311-1111123102100211-1030030320013203-0001211003021032-2211300222200022-3223233112010302)
- [Import](../guides/resources--cminstance--lifecycle--group-001.md#canonical-2123321023110030-3203022323131230-1220120302123131-1030022112120110-0323220331012032-2223233203113221-0320023321330310-3022210001202101)
- [Timeouts](../guides/resources--cminstance--lifecycle--group-001.md#canonical-2100330320322201-2020210220021033-0232320333111222-1300001300233232-3032303131222011-3303021003121312-1323223232111013-1210310210032331)
