---
page_title: "xcsh_origin_pool landing"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool landing."
---

# xcsh_origin_pool landing

<a id="canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213233001001222-3112220010332012-3033022333121012-3333111203230022-2022323221232130-0020121210223320-0010133303332033-2211311113032303"></a>

## xcsh_origin_pool — xcsh_origin_pool / 321220220312 / 2

Breadcrumbs:

- xcsh_origin_pool

Manages a Origin Pool resource in F5 Distributed Cloud for defining backend server pools for load
balancer targets.

<a id="canonical-0313220202023010-0000011132210302-2121201322123311-1112022331003111-0121132120332210-0122330130330100-3333003111131011-0102310202311120"></a>

## Prerequisites — xcsh_origin_pool / 321220220312 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `healthcheck`.

- healthcheck: Monitor origin server health

<a id="canonical-1200310330311300-3311023312122111-2232232121013011-1320030113001031-2122132211203132-3320320132021013-1201232130000110-1323202011120001"></a>

## Minimal configuration — xcsh_origin_pool / 321220220312 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# OriginPool Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing OriginPool by name
data "xcsh_origin_pool" "example" {
  name      = "example-origin-pool"
  namespace = "staging"
}

output "origin_pool_id" {
  value = data.xcsh_origin_pool.example.id
}
```

<a id="canonical-1202322122022001-3000120003222222-3212121133132131-3112323313020010-2102200320032321-0221013312120223-1313201130221033-0011100212022022"></a>

## Root configuration — xcsh_origin_pool / 321220220312 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3013313203312113-3202221002000123-0120031121333101-1310113131001120-0200303221303033-2330021231302210-1130100211232300-1331112220021311"></a>

## Next pages — xcsh_origin_pool / 321220220312 / 6

- [Property reference](../guides/data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [Examples](../guides/data-sources--origin_pool--examples--group-001.md#canonical-0212331131013132-1231233202011013-2201121321202120-2001300222023000-2201012103301302-3100013202031002-3021022030203221-0030001203321132)
