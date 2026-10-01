---
page_title: "xcsh_bgp landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp landing."
---

# xcsh_bgp landing

<a id="canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120312330231102-2001033031020113-0122111231120101-0322331321120322-1313111012113023-1223003222333221-0223332211233322-2022311111031223"></a>

## xcsh_bgp — xcsh_bgp / 210030210010 / 2

Breadcrumbs:

- xcsh_bgp

Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with
external bgp servers. it is created by users in system namespace. configuration.

<a id="canonical-0020301301230123-2223132010210332-3100000302120222-2311013322201123-0131213331230333-1321220333100133-0213301210231323-2223323332012210"></a>

## Prerequisites — xcsh_bgp / 210030210010 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1231310231002323-1023130301002231-0310101220011310-3102323130200311-3133101303213202-3133203123211202-1200310100110030-3033300110122032"></a>

## Minimal configuration — xcsh_bgp / 210030210010 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGP by name
data "xcsh_bgp" "example" {
  name      = "example-bgp"
  namespace = "staging"
}

output "bgp_id" {
  value = data.xcsh_bgp.example.id
}
```

<a id="canonical-1121010212331120-2122111030013001-3013221120212033-3320110020011010-3030000201210033-1232012232120120-0232123102311311-1113303222113033"></a>

## Root configuration — xcsh_bgp / 210030210010 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3130121123223120-1023321331221022-2101320300032202-1100203003311211-2313231002313100-2112000220210030-0031120001301313-2132102122011332"></a>

## Next pages — xcsh_bgp / 210030210010 / 6

- [Property reference](../guides/data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [Examples](../guides/data-sources--bgp--examples--group-001.md#canonical-0011333200121323-1202013222010111-0303012103322321-1132212333033031-2001120120213230-0113023223130122-0310112310322311-2002111203001003)
