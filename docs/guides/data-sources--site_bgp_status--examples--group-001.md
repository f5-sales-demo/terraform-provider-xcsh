---
page_title: "xcsh_site_bgp_status examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_bgp_status examples."
---

# xcsh_site_bgp_status examples

<a id="canonical-0202201122133102-2000330333212011-1233100330323223-3310020010003311-2000201110112032-0330000003111300-0311113023220122-2202123321123112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032331202031001-0122223113302201-0010300223302320-2331231103130121-3203131312323032-2023012032131311-3330302122031100-1220212010320110"></a>

## Examples — Examples / 013202220311 / 2

Breadcrumbs:

- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md#canonical-3202131020311001-0122003032002021-0022320312033010-1011001132303230-3130221113000032-1203011221022212-2212133031122301-1222231101030231)
- Examples

<a id="canonical-0330003130223122-0010312213022201-1130233320000231-1021232012030331-2330313030332313-1131200212222101-1211323310001210-0100020032031013"></a>

## Complete configurations — Examples / 013202220311 / 3

- [Data source](data-sources--site_bgp_status--examples--group-001.md#canonical-0202221200310332-2103213103300001-2111012122031330-0303210221320233-2201113001132033-0030000122010313-3332202313202302-1023103132020032): valid configuration.

<a id="canonical-0021333202112200-3233132232131203-0333031001110102-3310100002031110-0213320330300310-1201030011213023-0222100300300013-0131012302303211"></a>

## Next pages — Examples / 013202220311 / 4

- [Data source](data-sources--site_bgp_status--examples--group-001.md#canonical-0202221200310332-2103213103300001-2111012122031330-0303210221320233-2201113001132033-0030000122010313-3332202313202302-1023103132020032)
- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md#canonical-3202131020311001-0122003032002021-0022320312033010-1011001132303230-3130221113000032-1203011221022212-2212133031122301-1222231101030231)

<a id="canonical-0202221200310332-2103213103300001-2111012122031330-0303210221320233-2201113001132033-0030000122010313-3332202313202302-1023103132020032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202022000101213-2313210113321332-0220123202031230-3233211302112110-0231222031323032-1102222121302001-2121213030011233-0323230301333332"></a>

## Data source — Data source / 012333322100 / 2

Breadcrumbs:

- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md#canonical-3202131020311001-0122003032002021-0022320312033010-1011001132303230-3130221113000032-1203011221022212-2212133031122301-1222231101030231)
- [Examples](data-sources--site_bgp_status--examples--group-001.md#canonical-0202201122133102-2000330333212011-1233100330323223-3310020010003311-2000201110112032-0330000003111300-0311113023220122-2202123321123112)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_bgp_status/data-source.tf`; digest `sha256:78479126aafed9697f5378abaad404d355901539446a65e0ee8620a492c5b9cd`.

```terraform
# Wait for MAC-correlated BGP peers and both BGP and simplified route views to
# converge. Peer addresses and expected routes come from authoritative AWS
# TGW Connect resource attributes in a real configuration.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 6.0.0"
    }
  }
}

data "xcsh_site_bgp_status" "site" {
  namespace = "system"
  site      = "example-smsv2-site"

  expected_exported_routes = ["10.40.0.10/32"]

  expected_peers = {
    node_0_slo = {
      node                     = "node-0"
      role                     = "slo"
      mac                      = "02:00:00:00:00:10"
      peer_address             = "169.254.100.1"
      expected_imported_routes = ["10.20.0.0/16"]
    }
    node_0_sli = {
      node                     = "node-0"
      role                     = "sli"
      mac                      = "02:00:00:00:00:11"
      peer_address             = "169.254.101.1"
      expected_imported_routes = ["10.30.0.0/16"]
    }
  }

  timeout_seconds       = 300
  poll_interval_seconds = 10
}

output "bgp_converged" {
  value = data.xcsh_site_bgp_status.site.converged
}

output "bgp_peers" {
  value = data.xcsh_site_bgp_status.site.peers
}
```

<a id="canonical-1022011331331120-3003231213030333-0332010003221203-1331031102002210-2220311022001200-0012113010120310-3232132012102012-0012303210002000"></a>

## Next pages — Data source / 012333322100 / 3

- [Examples](data-sources--site_bgp_status--examples--group-001.md#canonical-0202201122133102-2000330333212011-1233100330323223-3310020010003311-2000201110112032-0330000003111300-0311113023220122-2202123321123112)
- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md#canonical-3202131020311001-0122003032002021-0022320312033010-1011001132303230-3130221113000032-1203011221022212-2212133031122301-1222231101030231)
