---
page_title: "xcsh_site_bgp_status"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_bgp_status."
---

# xcsh_site_bgp_status

<a id="canonical-3202131020311001-0122003032002021-0022320312033010-1011001132303230-3130221113000032-1203011221022212-2212133031122301-1222231101030231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_site_bgp_status

Polls authoritative F5 XC BGP and route observations until they agree with MAC-bound AWS
expectations.

<a id="canonical-1133001211100003-3311201020100011-1310222311122123-2023301101230020-1033012022302010-3020321331313233-2310121213132312-0310331230002323"></a>

### Prerequisites for `xcsh_site_bgp_status`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2210302230233220-2112030120330211-0301032313213002-2203032202300321-3231033031101330-2312001123223113-2022032123021300-2131113330201213"></a>

### Minimal configuration for `xcsh_site_bgp_status`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-0032310110003103-2010222030312030-1233132330333131-1123110210213102-1233033212330111-2202233312023313-2223223311121233-1121333131003230"></a>

### Root configuration for `xcsh_site_bgp_status`

Required root properties: `expected_exported_routes`, `expected_peers`, `namespace`, `site`. Full root flags and choices appear in the property reference.

<a id="canonical-0133133011312130-3133203203233032-0121112222311323-3203213320032301-2111200322331311-0221210031011132-0101020320231333-3111002013310001"></a>

### Explore this collection for `xcsh_site_bgp_status`

- [Property reference](../guides/data-sources--site_bgp_status--reference--group-001.md#canonical-2303311322002032-0001001012023033-2023213203123302-2111022211132330-1321230133100322-0231121310002212-0031102223011210-2123231130230213)
- [Examples](../guides/data-sources--site_bgp_status--examples--group-001.md#canonical-0202201122133102-2000330333212011-1233100330323223-3310020010003311-2000201110112032-0330000003111300-0311113023220122-2202123321123112)
