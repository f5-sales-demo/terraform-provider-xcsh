---
page_title: "xcsh_cdn_cache_rule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_cache_rule landing."
---

# xcsh_cdn_cache_rule landing

<a id="canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111220033101313-0102012103003201-2200132210010330-3030131332213220-3313131220023103-2321300301021020-3213230211013032-0213010231011210"></a>

## xcsh_cdn_cache_rule — xcsh_cdn_cache_rule / 111101232210 / 2

Breadcrumbs:

- xcsh_cdn_cache_rule

Manages a CDN Cache Rule resource in F5 Distributed Cloud for cdn loadbalancer specification.
configuration.

<a id="canonical-1300032220112123-3021212133123311-1203203012322030-2012302120111213-2313230213202300-0001033011002300-3013310322023213-0013011022232100"></a>

## Prerequisites — xcsh_cdn_cache_rule / 111101232210 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1122221220032100-2003000300133032-0123012101123210-0303032201030301-3122321131222321-0012022023211100-3231123233200113-0303212003211131"></a>

## Minimal configuration — xcsh_cdn_cache_rule / 111101232210 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CDNCacheRule Resource Example
# Manages a CDN Cache Rule resource in F5 Distributed Cloud for cdn loadbalancer specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNCacheRule configuration
resource "xcsh_cdn_cache_rule" "example" {
  name      = "example-cdn-cache-rule"
  namespace = "staging"
}
```

<a id="canonical-1020120102312203-1100121303230010-1110333103311212-1023100322023032-1133323303320022-1102220100200302-3321313122012021-3101232330101221"></a>

## Root configuration — xcsh_cdn_cache_rule / 111101232210 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0210323023120232-2302002023030220-2030221011332030-0210030003212301-1323302321211231-1031332313332003-3123331333111330-1122332303232132"></a>

## Next pages — xcsh_cdn_cache_rule / 111101232210 / 6

- [Property reference](../guides/resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- [Examples](../guides/resources--cdn_cache_rule--examples--group-001.md#canonical-0212331020322220-2311332101310031-2231102121022101-3011012031230323-2223331103003223-1202213101232312-0332331023100121-0133133331332211)
- [Import](../guides/resources--cdn_cache_rule--lifecycle--group-001.md#canonical-3131031201322010-3112121021233003-3011312020032020-3322133212201312-3002312013333321-3320100231003133-1131332122111201-1131313312031213)
- [Timeouts](../guides/resources--cdn_cache_rule--lifecycle--group-001.md#canonical-0000133011301231-1131323113203102-3312113213123203-0002300021303323-2121223020020010-0110203322300031-2321302120332221-2321030310122322)
