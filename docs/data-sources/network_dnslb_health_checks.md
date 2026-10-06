---
page_title: "xcsh_network_dnslb_health_checks"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_dnslb_health_checks."
---

# xcsh_network_dnslb_health_checks

<a id="canonical-1333320111300032-2212003230021231-1213001111023002-3323113032220220-3102221212223011-3012000302133013-0331320230013312-1210133113122312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_dnslb_health_checks

DNS Load Balancer health-check probe IPv4 addresses. Values are bundled from the pinned OpenAPI
release; this data source performs no network request. Ports and traffic direction are not encoded
in the manifest.

<a id="canonical-0110310233112032-3102103212300022-3232132332030231-2011020023302032-1003020123322322-0002001013303330-3000100110203103-1010202210202023"></a>

### Prerequisites for `xcsh_network_dnslb_health_checks`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3320203000033231-0003112100222212-1321020020200300-2200020210110223-3322323220230202-3132323032020231-3223212223021321-1231223333031133"></a>

### Minimal configuration for `xcsh_network_dnslb_health_checks`

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

data "xcsh_network_dnslb_health_checks" "https_probe" {}

# Match this explicit ingress port to the monitored endpoint.
output "https_health_check_ingress" {
  value = {
    direction   = "ingress"
    protocol    = "tcp"
    port        = 443
    cidr_blocks = data.xcsh_network_dnslb_health_checks.https_probe.cidr_blocks
  }
}
```

<a id="canonical-0112213312001313-1301312123302002-2001202213331212-1213112203220312-0202212010000210-0313330231133323-2201001222301021-1300021020332020"></a>

### Root configuration for `xcsh_network_dnslb_health_checks`

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-2210133113000000-0313023302032300-2021132233330320-2212001000112101-3031010320011320-0111100321100310-3110023232113232-1121100112131332"></a>

### Explore this collection for `xcsh_network_dnslb_health_checks`

- [Property reference](../guides/data-sources--network_dnslb_health_checks--reference--group-001.md#canonical-3223033123011120-0230303232120302-3331202330201322-3220130233113211-1303113102133001-0312222023330311-3123201131202200-2232113203112322)
- [Examples](../guides/data-sources--network_dnslb_health_checks--examples--group-001.md#canonical-0303310301112000-3100222121100210-2121303122313213-2313222313312303-0311232030113220-1220132333221233-3321201233332121-3112130311120012)
