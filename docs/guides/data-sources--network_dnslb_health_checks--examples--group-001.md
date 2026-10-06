---
page_title: "xcsh_network_dnslb_health_checks examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_dnslb_health_checks examples."
---

# xcsh_network_dnslb_health_checks examples

<a id="canonical-0303310301112000-3100222121100210-2121303122313213-2313222313312303-0311232030113220-1220132333221233-3321201233332121-3112130311120012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_dnslb_health_checks](../data-sources/network_dnslb_health_checks.md#canonical-1333320111300032-2212003230021231-1213001111023002-3323113032220220-3102221212223011-3012000302133013-0331320230013312-1210133113122312)
- Examples

<a id="canonical-1231023302303213-0033231222130021-3331220022001303-0002302123332313-2010203220212221-3200110222112301-3012313312230302-3300210031112020"></a>

### Complete configurations for `xcsh_network_dnslb_health_checks`

- [Data source](data-sources--network_dnslb_health_checks--examples--group-001.md#canonical-0321111320013123-0223332231031100-3000220203002123-2232221023212333-2331320120300320-0322302331131021-1022131010202012-1111322330310001): valid configuration.

<a id="canonical-0321111320013123-0223332231031100-3000220203002123-2232221023212333-2331320120300320-0322302331131021-1022131010202012-1111322330310001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_dnslb_health_checks](../data-sources/network_dnslb_health_checks.md#canonical-1333320111300032-2212003230021231-1213001111023002-3323113032220220-3102221212223011-3012000302133013-0331320230013312-1210133113122312)
- [Examples](data-sources--network_dnslb_health_checks--examples--group-001.md#canonical-0303310301112000-3100222121100210-2121303122313213-2313222313312303-0311232030113220-1220132333221233-3321201233332121-3112130311120012)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_dnslb_health_checks/data-source.tf`; digest `sha256:806558a3420657577ad88d195487a5583b29bf785d14c8b410da5c5fe1315a79`.

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
