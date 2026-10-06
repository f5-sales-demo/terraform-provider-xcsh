---
page_title: "xcsh_site_upgrade_status examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_upgrade_status examples."
---

# xcsh_site_upgrade_status examples

<a id="canonical-3223111323110230-0123231002213330-0202201232313001-1031031203203312-0001221101110320-3222113013223302-2012232221110223-2002010112100100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_site_upgrade_status](../data-sources/site_upgrade_status.md#canonical-0311103030013201-2023013222321300-1123321212232123-1232102210103331-2001110222022222-0103301221312320-3101213121122310-1101333312101011)
- Examples

<a id="canonical-1323310113011022-2322201033301201-1310100120001331-1200103312112130-2112331233203111-1203202203212302-2332300033001211-1201203100211210"></a>

### Complete configurations for `xcsh_site_upgrade_status`

- [Data source](data-sources--site_upgrade_status--examples--group-001.md#canonical-2300222103311212-2122102320121120-1002232212000123-0112301033311311-3211020331002101-0123231331013330-0322323313202022-3000302011202323): valid configuration.

<a id="canonical-2300222103311212-2122102320121120-1002232212000123-0112301033311311-3211020331002101-0123231331013330-0322323313202022-3000302011202323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_site_upgrade_status](../data-sources/site_upgrade_status.md#canonical-0311103030013201-2023013222321300-1123321212232123-1232102210103331-2001110222022222-0103301221312320-3101213121122310-1101333312101011)
- [Examples](data-sources--site_upgrade_status--examples--group-001.md#canonical-3223111323110230-0123231002213330-0202201232313001-1031031203203312-0001221101110320-3222113013223302-2012232221110223-2002010112100100)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_upgrade_status/data-source.tf`; digest `sha256:3c8a578ab685be3268d5b0dab58926aa035dcddae8cb2d3ebaa0f6a17186a683`.

```terraform
# Observe upgrade eligibility or wait for supplied software and OS targets to
# be installed with the site back ONLINE.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 7.3.0"
    }
  }
}

data "xcsh_site_upgrade_status" "site" {
  site = "example-smsv2-site"

  expected_software_version = "crt-20260201-0179"
  expected_os_version       = "9.2026.17"
  wait                      = true
  timeout_seconds           = 7200
  poll_interval_seconds     = 30
}

output "upgrade_converged" {
  value = data.xcsh_site_upgrade_status.site.target_converged
}
```
