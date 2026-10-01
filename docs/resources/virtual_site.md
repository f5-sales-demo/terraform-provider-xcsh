---
page_title: "xcsh_virtual_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_virtual_site landing."
---

# xcsh_virtual_site landing

<a id="canonical-0021200012330233-0110101013310110-3130112023213001-0301030013233201-0302311302331132-3122030310203321-2022313333210003-2333313130301011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310311032032131-2233301220332202-1220330212312322-3310331313312220-0203323113320221-1303213033021123-2221132120322321-0221030110230301"></a>

## xcsh_virtual_site — xcsh_virtual_site / 032000212311 / 2

Breadcrumbs:

- xcsh_virtual_site

Manages virtual site object in given namespace in F5 Distributed Cloud.

<a id="canonical-1333230212032200-3202011000201120-3101211323213000-2322202230230330-1013130021211010-2330301112003203-0113231132103130-1233023232110210"></a>

## Prerequisites — xcsh_virtual_site / 032000212311 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-3231011201031320-2031100121231100-2320121000022211-3102023223301200-0310313031231020-0112230103130033-0203200001311002-2332101330300320"></a>

## Minimal configuration — xcsh_virtual_site / 032000212311 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualSite Resource Example
# Manages virtual site object in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualSite configuration
resource "xcsh_virtual_site" "example" {
  name      = "example-virtual-site"
  namespace = "staging"
}
```

<a id="canonical-0011012221103032-1013303100302322-2122020131130020-2233022333032303-2030111333102213-0323200033213032-3210201312310322-2012002320120101"></a>

## Root configuration — xcsh_virtual_site / 032000212311 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3000231310311120-1102023013221323-0313331000320311-0101132100012023-3010330010210032-3312121021011313-2202303003123330-0230031310121123"></a>

## Next pages — xcsh_virtual_site / 032000212311 / 6

- [Property reference](../guides/resources--virtual_site--reference--group-001.md#canonical-3012320130112032-2032311131023201-2230001222123031-0023020210201003-3230331331010002-2310130201122230-2211332320033010-2101122120231312)
- [Examples](../guides/resources--virtual_site--examples--group-001.md#canonical-2012211220121122-0020110113010220-0003310321031112-3022120221101003-0102001113103033-3213021033123003-3333023021232031-2133232021100201)
- [Import](../guides/resources--virtual_site--lifecycle--group-001.md#canonical-2203103200003320-0013322321202333-1222011210030202-0100100121022331-0211122210202331-0002010031312231-2303212101221333-0021010102001201)
- [Timeouts](../guides/resources--virtual_site--lifecycle--group-001.md#canonical-1312112231301022-1010302223000312-2320130133100200-1233100132013300-1030220122310030-3223321110303221-0121302221211021-1203221232222231)
