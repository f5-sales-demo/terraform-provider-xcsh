---
page_title: "xcsh_lma_region landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_lma_region landing."
---

# xcsh_lma_region landing

<a id="canonical-2212011320333232-1123311200100001-1321311330220000-3033022231230311-3012011130132021-2120003020013220-1120200123320031-0302331121003101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230111132133213-3023333021020102-1223021122221332-0030320330202123-0330301312120211-1232000300033111-1103120200032012-0220130002213323"></a>

## xcsh_lma_region — xcsh_lma_region / 132102031020 / 2

Breadcrumbs:

- xcsh_lma_region

Manages a Lma Region resource in F5 Distributed Cloud for lma region specification. configuration.
(read-only data source)

<a id="canonical-2120001113322202-0103032030003332-0131220313032211-0032131011003031-2021303210201230-3312021220200012-2112121102030012-3223322220001110"></a>

## Prerequisites — xcsh_lma_region / 132102031020 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3031111320021202-1303012311032331-0210002311003113-1233321111130301-0223122123232001-2323230320300313-3020133022033112-1010301330011100"></a>

## Minimal configuration — xcsh_lma_region / 132102031020 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# LmaRegion Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing LmaRegion by name
data "xcsh_lma_region" "example" {
  name      = "example-lma-region"
  namespace = "staging"
}

output "lma_region_id" {
  value = data.xcsh_lma_region.example.id
}
```

<a id="canonical-1010331233131231-3211333330211100-1230121102000001-1202020032322032-2203122331023331-3220123031231033-1120001003221302-3110233103210002"></a>

## Root configuration — xcsh_lma_region / 132102031020 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3222103010312300-0302203101222110-1012321222220002-0012031001130210-1111031013023301-2220203001010020-1220011032223023-3310322210121310"></a>

## Next pages — xcsh_lma_region / 132102031020 / 6

- [Property reference](../guides/data-sources--lma_region--reference--group-001.md#canonical-0023210031100010-2112103013031230-3231331013203233-3102020222122202-0111202222003331-1021220322200330-2033303210112210-1133233201200122)
- [Examples](../guides/data-sources--lma_region--examples--group-001.md#canonical-3201110113300321-1111333210230110-1310113001212123-0211303112100221-0303113020300311-3210230100202001-3211313112222002-3001132233133330)
