---
page_title: "xcsh_network_cdn landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_cdn landing."
---

# xcsh_network_cdn landing

<a id="canonical-3332101011301320-0312003011132032-2233133133001013-2021212311222213-1103000302133332-3011321012313330-0002302023203120-2131131001121112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311301030322320-2222030232232212-0102110213110021-2201223313032203-3230001013023323-0101101132132000-1003101302002202-1310331302300230"></a>

## xcsh_network_cdn — xcsh_network_cdn / 033102223101 / 2

Breadcrumbs:

- xcsh_network_cdn

CDN IPv4 networks for origin or network-firewall ingress allowlists. Values are bundled from the
pinned OpenAPI release; this data source performs no network request. Ports and traffic direction
are not encoded in the manifest.

<a id="canonical-1001301322230211-3303232132313030-0331300103233132-1313313123133120-1310210230003200-3300300320311000-3203310103302331-2230312230132202"></a>

## Prerequisites — xcsh_network_cdn / 033102223101 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1012213103121011-1113033131020022-3310203213120331-0332302200132330-1313312303212033-2032130312313101-3032332003112012-1132020001200011"></a>

## Minimal configuration — xcsh_network_cdn / 033102223101 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_cdn" "origin_ingress" {}

output "cdn_https_origin_ingress" {
  value = {
    direction   = "ingress"
    protocol    = "tcp"
    port        = 443
    cidr_blocks = data.xcsh_network_cdn.origin_ingress.cidr_blocks
  }
}
```

<a id="canonical-1323033232031013-1010322210121121-2133210301111100-2031203303331311-2230313012030003-3021311020210211-0111200030110023-0320200010211031"></a>

## Root configuration — xcsh_network_cdn / 033102223101 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-0020133230233221-3300203231103310-2103131120133002-1321022200111200-3031133013012231-3000332022033333-1330321123103312-0133331200302002"></a>

## Next pages — xcsh_network_cdn / 033102223101 / 6

- [Property reference](../guides/data-sources--network_cdn--reference--group-001.md#canonical-3023312220230230-1021020032312223-2101121310120203-1113101030220133-3130221022210110-3202200221031232-3110203322201332-1031200023231013)
- [Examples](../guides/data-sources--network_cdn--examples--group-001.md#canonical-1100110203131112-1333321133001020-1320302012333111-0030313021312021-2112102200121123-2030202111321331-2032230133120311-2023332101312322)
