---
page_title: "xcsh_ike2 landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike2 landing."
---

# xcsh_ike2 landing

<a id="canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003201101010333-3121330232222003-1310102210030000-0331331023101321-1322020203303312-1030210220230011-0200202300133232-3300303011011202"></a>

## xcsh_ike2 — xcsh_ike2 / 321302123233 / 2

Breadcrumbs:

- xcsh_ike2

Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification. configuration.

<a id="canonical-1102013311210323-2302022210103121-1320100022311312-2213302330011010-3101223333103301-0220033213302230-0313030313112202-0313321231000132"></a>

## Prerequisites — xcsh_ike2 / 321302123233 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3202003220213010-1033303001212203-0123232000212300-2101020232211123-0200002202131011-3333022001213023-3220201220131212-1202000132023012"></a>

## Minimal configuration — xcsh_ike2 / 321302123233 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Ike2 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Ike2 by name
data "xcsh_ike2" "example" {
  name      = "example-ike2"
  namespace = "staging"
}

output "ike2_id" {
  value = data.xcsh_ike2.example.id
}
```

<a id="canonical-2022222101020000-2311201332113201-1113110213123113-1011301311300132-3213303333213021-3222200113203023-1310233000223231-3133320023132020"></a>

## Root configuration — xcsh_ike2 / 321302123233 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1031121020123131-2000322111211031-0110220032212101-1013112130220003-3200222220031202-3113233003003130-2003223113312311-1020110012030302"></a>

## Next pages — xcsh_ike2 / 321302123233 / 6

- [Property reference](../guides/data-sources--ike2--reference--group-001.md#canonical-2210220022220321-0032112331013302-1322232000331110-0332122122320233-1212003130332202-2130200012212213-0001200123130221-1102033222321323)
- [Examples](../guides/data-sources--ike2--examples--group-001.md#canonical-1130101033033222-2212233311300131-1210202130100120-3023323030232133-0213103222002211-1303133001221200-2221113230210131-3112213100212123)
