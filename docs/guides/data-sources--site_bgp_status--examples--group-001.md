---
page_title: "xcsh_site_bgp_status examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_bgp_status examples."
---

# xcsh_site_bgp_status examples

<a id="canonical-0202201122133102-2000330333212011-1233100330323223-3310020010003311-2000201110112032-0330000003111300-0311113023220122-2202123321123112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md#canonical-3202131020311001-0122003032002021-0022320312033010-1011001132303230-3130221113000032-1203011221022212-2212133031122301-1222231101030231)
- Examples

<a id="canonical-3032331202031001-0122223113302201-0010300223302320-2331231103130121-3203131312323032-2023012032131311-3330302122031100-1220212010320110"></a>

### Complete configurations for `xcsh_site_bgp_status`

- [Data source](data-sources--site_bgp_status--examples--group-001.md#canonical-0202221200310332-2103213103300001-2111012122031330-0303210221320233-2201113001132033-0030000122010313-3332202313202302-1023103132020032): valid configuration.

<a id="canonical-0202221200310332-2103213103300001-2111012122031330-0303210221320233-2201113001132033-0030000122010313-3332202313202302-1023103132020032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

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
