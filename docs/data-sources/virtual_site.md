---
page_title: "xcsh_virtual_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_virtual_site landing."
---

# xcsh_virtual_site landing

<a id="canonical-1031213012322000-0133010132320022-1031302313301212-0223023031301131-3333302202323221-1032021131132033-3323021133330011-1320213202212000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303030102103020-0223303311333311-1310131310332313-3022022211002313-1203322011020002-0313330121321013-0213100031223333-1200122033312100"></a>

## xcsh_virtual_site — xcsh_virtual_site / 121112122330 / 2

Breadcrumbs:

- xcsh_virtual_site

Manages virtual site object in given namespace in F5 Distributed Cloud.

<a id="canonical-2333112103331303-3110230120103323-3001022313332030-1313011201101103-2110212222001223-3033012013032130-3023221231122232-0030230012330301"></a>

## Prerequisites — xcsh_virtual_site / 121112122330 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-2202202123121131-2300122012103322-0102021131132312-2321200121123331-2323213203223312-0021001221220121-1131120201120313-1001000031233102"></a>

## Minimal configuration — xcsh_virtual_site / 121112122330 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualSite by name
data "xcsh_virtual_site" "example" {
  name      = "example-virtual-site"
  namespace = "staging"
}

output "virtual_site_id" {
  value = data.xcsh_virtual_site.example.id
}
```

<a id="canonical-3310001311000231-0223313320201213-1311322123021031-3103121220020210-1211333201122220-2133302302311003-2333012102111010-1210100330030112"></a>

## Root configuration — xcsh_virtual_site / 121112122330 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0320131223310030-0021223100022101-1331020231131232-1031102302223203-3030101003133320-3233311331013100-3133302123302220-1310130021303230"></a>

## Next pages — xcsh_virtual_site / 121112122330 / 6

- [Property reference](../guides/data-sources--virtual_site--reference--group-001.md#canonical-0001012332120212-3321121210011001-1030302032121112-1211101313330022-1313122033210013-3011100213131302-3231210233123131-0230121213221310)
- [Examples](../guides/data-sources--virtual_site--examples--group-001.md#canonical-1123310110133133-0113110231020120-3023230122332103-2222112230100333-0021123310130212-2011130123011330-1222102221230332-2102220101233112)
