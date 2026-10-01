---
page_title: "xcsh_securemesh_site landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site landing."
---

# xcsh_securemesh_site landing

<a id="canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320112311100213-3221312103222301-3323130113200330-0013313332201032-0023133122103132-2103110221032211-3300032301203223-2310330233222133"></a>

## xcsh_securemesh_site — xcsh_securemesh_site / 120203013002 / 2

Breadcrumbs:

- xcsh_securemesh_site

Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with
distributed security.

<a id="canonical-1010030203322233-2022031130230122-2333213223333221-1021333122202023-3001001103002011-2201221331022233-3001000320301033-3112313122333302"></a>

## Prerequisites — xcsh_securemesh_site / 120203013002 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0101223110333110-2100030011303032-2121133100000023-3100211201320021-2110222331223221-2232022223002232-2233002022131013-1312011311310200"></a>

## Minimal configuration — xcsh_securemesh_site / 120203013002 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecuremeshSite Resource Example
# Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with distributed security.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecuremeshSite configuration
resource "xcsh_securemesh_site" "example" {
  name      = "example-securemesh-site"
  namespace = "staging"

  volterra_certified_hw = "example-value"
}
```

<a id="canonical-2000220313032200-1302213203211333-0112231001323210-1022232102131310-2313213031222000-0301220001023001-0232001331132302-1033112312012001"></a>

## Root configuration — xcsh_securemesh_site / 120203013002 / 5

Required root properties: `name`, `namespace`, `volterra_certified_hw`. Full root flags and choices appear in the property reference.

<a id="canonical-1012102332303100-1113112332021002-0332002110030313-0033203011133121-3220212313321123-2121102211100231-0201132133123020-3132011323322200"></a>

## Next pages — xcsh_securemesh_site / 120203013002 / 6

- [Property reference](../guides/resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [Examples](../guides/resources--securemesh_site--examples--group-001.md#canonical-3111120021123011-3221120323311113-3112101230310313-2310212131101013-1131333203200301-3002303212200130-3110220100102222-1333003331113232)
- [Import](../guides/resources--securemesh_site--lifecycle--group-001.md#canonical-3122221033201210-1330233003220311-1230032331223330-0321123221100133-1103011020120301-0300103102121102-3202333110231012-1031103103100130)
- [Timeouts](../guides/resources--securemesh_site--lifecycle--group-001.md#canonical-1223112033021212-2212111023211300-2103120133201002-1021301213112131-2210323320220301-1023111213221130-0100112021311110-0002011213223123)
