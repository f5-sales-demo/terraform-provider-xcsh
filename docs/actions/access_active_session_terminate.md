---
page_title: "xcsh_access_active_session_terminate landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_session_terminate landing."
---

# xcsh_access_active_session_terminate landing

<a id="canonical-0233012110120301-2112002303202012-1202012200302211-1131130003102111-0002212030201223-3120132220230001-3123332221233323-3312211113313020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133202300013012-3003000200003323-1113313122010330-3300132013002131-0233202013202013-3223112113033232-3223103011002313-1122010032212231"></a>

## xcsh_access_active_session_terminate — xcsh_access_active_session_terminate / 013311211032 / 2

Breadcrumbs:

- xcsh_access_active_session_terminate

Resource deletion operation.

<a id="canonical-2312000312232202-0003323101332011-1130033100131120-2022203023123202-0031101232022021-3003310230012132-0331222112213020-1220320211132222"></a>

## Prerequisites — xcsh_access_active_session_terminate / 013311211032 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1010030301101113-1111131312002233-2021310230220110-0113232113223023-1300300123332122-1020213013110320-3332133011031020-3231110112021100"></a>

## Minimal configuration — xcsh_access_active_session_terminate / 013311211032 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AccessActiveSessionTerminate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_access_active_session_terminate" "example" {
  config {
    id        = "example-value"
    namespace = "example-value"
  }
}
```

<a id="canonical-1303323230103011-0111010112301200-1113122300211323-1132021333312302-1110012223202110-1103301303210210-0103033232311100-1020020322222110"></a>

## Root configuration — xcsh_access_active_session_terminate / 013311211032 / 5

Required root properties: `id`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3223201003130110-3331032203113113-3200120101101020-1312333210200311-0233120011101112-3030032023110321-1203301133322300-2022202203131221"></a>

## Next pages — xcsh_access_active_session_terminate / 013311211032 / 6

- [Property reference](../guides/actions--access_active_session_terminate--reference--group-001.md#canonical-2012103202003132-2120211331100001-2211110200100202-0002003131013320-2231102211210130-1202110020003233-0223030110331233-3332210231120222)
- [Examples](../guides/actions--access_active_session_terminate--examples--group-001.md#canonical-1133133321103132-3331000032332003-2020202330010032-1212111000232113-3101033331232202-1333332030010121-2032311020210022-3013131203021012)
- [Lifecycle](../guides/actions--access_active_session_terminate--lifecycle--group-001.md#canonical-0021200100200122-2121200111121212-2203321100011030-3331302303110333-2013121000211001-3010130202311122-0110312020333203-1300330200022113)
