---
page_title: "xcsh_discovery landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery landing."
---

# xcsh_discovery landing

<a id="canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230322100133032-1032130310311133-3202223023121120-3301222012231102-2111222021323222-3300301112111102-1122221302020013-1012200223312133"></a>

## xcsh_discovery — xcsh_discovery / 331130132013 / 2

Breadcrumbs:

- xcsh_discovery

Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site
or virtual site in system namespace. configuration.

<a id="canonical-1233303032000021-0130323301112023-2300132023313122-0301001120210122-0023002220023000-1123322022330303-3201332222233321-0222311222310203"></a>

## Prerequisites — xcsh_discovery / 331130132013 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3202002000012000-3222223202231312-3221002101102032-2213113323131111-3323133220022222-2100003100103131-1312202020332120-0002131323100230"></a>

## Minimal configuration — xcsh_discovery / 331130132013 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Discovery Resource Example
# Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site or virtual site in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Discovery configuration
resource "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}
```

<a id="canonical-1222020120212131-2330002003111201-1033113133102221-3120121233303232-1212302202133223-3200222020131232-1321031123130132-0313131332302121"></a>

## Root configuration — xcsh_discovery / 331130132013 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3310022110111201-0100012222002231-1020002222133003-1320223303230032-3111302333201131-0233310011011201-3211331331003231-3113001010232330"></a>

## Next pages — xcsh_discovery / 331130132013 / 6

- [Property reference](../guides/resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [Examples](../guides/resources--discovery--examples--group-001.md#canonical-0302301231230113-2120313302220013-1021221011331111-2001300103031013-3110221312211320-0012310022011122-0031200230200333-0110131331330233)
- [Import](../guides/resources--discovery--lifecycle--group-001.md#canonical-1313311020210021-3020033220313300-1322001302302330-2131100313112221-0110231330303020-2121121002302101-2003332221010300-3030100033333011)
- [Timeouts](../guides/resources--discovery--lifecycle--group-001.md#canonical-3320213313210301-2322320001013033-0321323121023112-2330203111032133-2001303103233203-0221013112330320-3030121031001013-2122031130000130)
