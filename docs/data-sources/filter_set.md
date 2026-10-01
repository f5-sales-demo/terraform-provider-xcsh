---
page_title: "xcsh_filter_set landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_filter_set landing."
---

# xcsh_filter_set landing

<a id="canonical-3220320303320131-0131310102101111-3102330020220020-3222001011311133-0032232323232000-2210333331011111-1110011031200321-2111213033230001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320221001323123-3212131320203223-0113030111133222-0101313113121232-3311032203003212-1032231233320213-3031023021113100-3103012121011112"></a>

## xcsh_filter_set — xcsh_filter_set / 132032021100 / 2

Breadcrumbs:

- xcsh_filter_set

Manages specification in F5 Distributed Cloud.

<a id="canonical-2223230302022013-2200133320022111-1101121110303111-1332312130102031-1032333030212203-1022330211110312-1033320333303021-1221331313300033"></a>

## Prerequisites — xcsh_filter_set / 132032021100 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3131111220233203-3012023031101020-2021001133321301-3212013311113200-2221322322213123-2313303103222222-1332132130013200-1202010023020020"></a>

## Minimal configuration — xcsh_filter_set / 132032021100 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FilterSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FilterSet by name
data "xcsh_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "staging"
}

output "filter_set_id" {
  value = data.xcsh_filter_set.example.id
}
```

<a id="canonical-0032323212031020-2001013033022202-0303233132211202-2323300210201211-3332311202230000-1101310011123301-1201112300020101-0102103331220003"></a>

## Root configuration — xcsh_filter_set / 132032021100 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0022221200211021-3230300311110011-2132222232321331-0130300111232230-1210202313312131-1220011131032032-3231001133030112-0310023200320323"></a>

## Next pages — xcsh_filter_set / 132032021100 / 6

- [Property reference](../guides/data-sources--filter_set--reference--group-001.md#canonical-3103220200110331-2212303313301103-2220311222333203-2213110301121201-1003211001200312-3213023223222201-1232232300201201-0121121213112121)
- [Examples](../guides/data-sources--filter_set--examples--group-001.md#canonical-1221330213020113-0002233103033230-1202120022111310-1110203311220112-1101333133103123-2101032321203030-1310100203212111-0210320120222021)
