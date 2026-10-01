---
page_title: "xcsh_fast_acl landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl landing."
---

# xcsh_fast_acl landing

<a id="canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020222312322311-2320121021032003-3313220010011032-3311003311232000-1111300223232120-1210202333021123-0130330011022311-3131210311232010"></a>

## xcsh_fast_acl — xcsh_fast_acl / 013323210333 / 2

Breadcrumbs:

- xcsh_fast_acl

Manages object, object contains rules to protect site from denial of service It has
destination\{destination IP, destination port) and references to in F5 Distributed Cloud.

<a id="canonical-1322322230123300-0311200003233221-2310330032233321-3201110102220010-3120320130221113-1201322031031303-3022233213231002-0332310231033322"></a>

## Prerequisites — xcsh_fast_acl / 013323210333 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0133012310232132-1101331031201303-2000330133113133-1222111231312332-1312011100333330-1112033131012310-0132113233332031-3332213211220203"></a>

## Minimal configuration — xcsh_fast_acl / 013323210333 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FastACL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FastACL by name
data "xcsh_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"
}

output "fast_acl_id" {
  value = data.xcsh_fast_acl.example.id
}
```

<a id="canonical-0321011203100233-0303101122312232-0213013321033231-0233101301032003-1330103323230332-1333011110121021-3302210232213001-1003330001203011"></a>

## Root configuration — xcsh_fast_acl / 013323210333 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-2111031200123132-3030030332102133-0322110201300331-3123001000023233-2320220232323331-1113221301131221-0111111200101210-2220111112233232"></a>

## Next pages — xcsh_fast_acl / 013323210333 / 6

- [Property reference](../guides/data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [Examples](../guides/data-sources--fast_acl--examples--group-001.md#canonical-1233100313101303-1021322332001030-0313120000210313-3020023113032003-2211001033322132-2302120033032012-0300033322003332-3021212102003323)
