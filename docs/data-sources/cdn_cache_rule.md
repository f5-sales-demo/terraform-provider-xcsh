---
page_title: "xcsh_cdn_cache_rule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_cache_rule landing."
---

# xcsh_cdn_cache_rule landing

<a id="canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331121330323002-0000332222131030-1101000203030102-0211133211001122-0033011112310123-1223311203310200-1212121130011130-1001301321113330"></a>

## xcsh_cdn_cache_rule — xcsh_cdn_cache_rule / 212113313011 / 2

Breadcrumbs:

- xcsh_cdn_cache_rule

Manages a CDN Cache Rule resource in F5 Distributed Cloud for cdn loadbalancer specification.
configuration.

<a id="canonical-0320100012000202-0113300012200010-3120123012232012-2230300101120133-2333310223010122-0113112310133221-1111201200231222-2310311232131320"></a>

## Prerequisites — xcsh_cdn_cache_rule / 212113313011 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2333320112310222-0320222133222321-2221103311332033-0332010203131101-2031133312202032-0022010031211002-0103322332210131-1120010102000123"></a>

## Minimal configuration — xcsh_cdn_cache_rule / 212113313011 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CDNCacheRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNCacheRule by name
data "xcsh_cdn_cache_rule" "example" {
  name      = "example-cdn-cache-rule"
  namespace = "staging"
}

output "cdn_cache_rule_id" {
  value = data.xcsh_cdn_cache_rule.example.id
}
```

<a id="canonical-3122230312131112-0222220201021010-1003101011330322-2100000111103011-0230232120310321-3023333232003211-3220311232020301-0030221203221100"></a>

## Root configuration — xcsh_cdn_cache_rule / 212113313011 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3023231202001220-3303122332131020-2002200222201011-0201213100230112-1002331112001212-1013303332230112-3020122230333131-1101202330301033"></a>

## Next pages — xcsh_cdn_cache_rule / 212113313011 / 6

- [Property reference](../guides/data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [Examples](../guides/data-sources--cdn_cache_rule--examples--group-001.md#canonical-1123232111331111-0131132300333101-3320021113101203-1311110200002032-3130303103303300-0112302233013033-1211133123333011-3211232123213233)
