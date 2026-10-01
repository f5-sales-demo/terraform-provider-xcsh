---
page_title: "xcsh_origin_pool landing"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool landing."
---

# xcsh_origin_pool landing

<a id="canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213010323203211-2110313030320333-0300322032200311-1222020331112010-3133113233113331-3101111130203123-0221320223131112-1001312003031110"></a>

## xcsh_origin_pool — xcsh_origin_pool / 113323003331 / 2

Breadcrumbs:

- xcsh_origin_pool

Manages a Origin Pool resource in F5 Distributed Cloud for defining backend server pools for load
balancer targets.

<a id="canonical-2311232010312030-0230030133033320-3210202230122210-2112223102033022-0310202201200130-3003010323001300-0331303121322131-1310230130110200"></a>

## Prerequisites — xcsh_origin_pool / 113323003331 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `healthcheck`.

- healthcheck: Monitor origin server health

<a id="canonical-0022331102212230-0100300213222222-3232221102102110-0322313011232020-2301203302020132-3320213300000303-0011311210321203-1232132133310300"></a>

## Minimal configuration — xcsh_origin_pool / 113323003331 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# OriginPool Resource Example
# Manages a Origin Pool resource in F5 Distributed Cloud for defining backend server pools for load balancer targets.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic OriginPool configuration
resource "xcsh_origin_pool" "example" {
  name      = "example-origin-pool"
  namespace = "staging"
}
```

<a id="canonical-3022201302003313-3331200310121132-1221231333003121-3311022230113232-3203322332203032-1021230031002321-0120221212110330-3233221321110221"></a>

## Root configuration — xcsh_origin_pool / 113323003331 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1221101322332110-1211223203100301-1311210111313230-3332303303323322-3300300320310000-0213133121103310-3203222230033220-0211333311200121"></a>

## Next pages — xcsh_origin_pool / 113323003331 / 6

- [Property reference](../guides/resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [Examples](../guides/resources--origin_pool--examples--group-001.md#canonical-0000311122100023-3302033101223232-1212303232002022-2202323221002012-2312220110133231-0212131003330311-3211202222332020-3322333102120030)
- [Import](../guides/resources--origin_pool--lifecycle--group-001.md#canonical-0131233233010320-2210122211033302-1313031331011222-0232223230333231-2013323113332310-1213133202312123-0102221103222303-2021130232003233)
- [Timeouts](../guides/resources--origin_pool--lifecycle--group-001.md#canonical-0212313030122302-0102202012132131-0133203121302213-3222132223330033-0312123203032231-1332000122012233-1103232200303321-0113120332101203)
