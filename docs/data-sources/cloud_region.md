---
page_title: "xcsh_cloud_region landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_region landing."
---

# xcsh_cloud_region landing

<a id="canonical-1312131213131300-2032133001002132-0000012113013011-3311003200233300-1320113133322101-0123111332203012-2202233202001321-2113202320323303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020030110223333-3121310013311311-0230033311031202-3332031231022111-2301131203033113-3130223301012130-1211002302112313-0213203211331013"></a>

## xcsh_cloud_region — xcsh_cloud_region / 000000123031 / 2

Breadcrumbs:

- xcsh_cloud_region

Manages a Cloud Region resource in F5 Distributed Cloud for cloud re specification. configuration.
(read-only data source)

<a id="canonical-0301201322123032-3013121231211122-0001013123130303-1310320321202133-3020000112230001-3230202320223321-0331310203103303-2020120332110232"></a>

## Prerequisites — xcsh_cloud_region / 000000123031 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3322231323333310-2220212221120313-0212302120120000-1110100203231302-3122232232013233-0001310122031312-2022323112030210-2323103131100110"></a>

## Minimal configuration — xcsh_cloud_region / 000000123031 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudRegion Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudRegion by name
data "xcsh_cloud_region" "example" {
  name      = "example-cloud-region"
  namespace = "staging"
}

output "cloud_region_id" {
  value = data.xcsh_cloud_region.example.id
}
```

<a id="canonical-1103202112002113-2330102300103222-1200300111211320-1231033212003230-3202113203310320-1332313013212031-3102201130210200-0330011230013320"></a>

## Root configuration — xcsh_cloud_region / 000000123031 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0313111101120012-2233011122200130-1130331002233312-3013031023001023-1120033010022013-1022112201201210-2123312130203301-0220030323333013"></a>

## Next pages — xcsh_cloud_region / 000000123031 / 6

- [Property reference](../guides/data-sources--cloud_region--reference--group-001.md#canonical-1110122010012032-2131203101013123-0321120130133102-2023302312022110-2201320321102001-2021000331031320-3100331122100231-2101010311313120)
- [Examples](../guides/data-sources--cloud_region--examples--group-001.md#canonical-1113021303031313-2130103003133303-3102332131232032-0123302311230122-1221023131133130-0211322133230023-1313233232000301-0100320222200220)
