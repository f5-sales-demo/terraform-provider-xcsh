---
page_title: "xcsh_third_party_application landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_third_party_application landing."
---

# xcsh_third_party_application landing

<a id="canonical-2320222231123022-0323332222212132-3300113320120123-0222011302331223-3222001221131320-2123313213312301-2022200110110030-3312032011311320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032233133133311-2202320310022133-0231202232233230-1003010131200233-3303102222302310-2323022322023202-3112111223332120-3010302030133120"></a>

## xcsh_third_party_application — xcsh_third_party_application / 302313101311 / 2

Breadcrumbs:

- xcsh_third_party_application

Manages a Third Party Application resource in F5 Distributed Cloud for third party application
specification. configuration. (read-only data source)

<a id="canonical-2001231210203132-2033230000102100-2302220033221102-2011023313223023-0200031323001300-1032232022321112-3131212101302202-0130220230321311"></a>

## Prerequisites — xcsh_third_party_application / 302313101311 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1031032233122112-2312220021203303-1022002012021331-1100211231113031-0101103311322132-0331113203230130-1231023112333322-3230002223103303"></a>

## Minimal configuration — xcsh_third_party_application / 302313101311 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ThirdPartyApplication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ThirdPartyApplication by name
data "xcsh_third_party_application" "example" {
  name      = "example-third-party-application"
  namespace = "staging"
}

output "third_party_application_id" {
  value = data.xcsh_third_party_application.example.id
}
```

<a id="canonical-0310032120210303-0201032331320000-0331111231310130-3200023333021322-3221123330103100-0113202211120320-2101102232200310-2100230330311001"></a>

## Root configuration — xcsh_third_party_application / 302313101311 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3030330133003031-3323313023232212-3020103303220003-2331313032101312-2221120212231123-2232233212301303-1122022332121323-1011233311031323"></a>

## Next pages — xcsh_third_party_application / 302313101311 / 6

- [Property reference](../guides/data-sources--third_party_application--reference--group-001.md#canonical-2301303220030133-0020032233202203-2021003112131202-1213230110230100-3232132313001030-3102223332133223-2203300133202211-0003303130203133)
- [Examples](../guides/data-sources--third_party_application--examples--group-001.md#canonical-0120101222002302-3233320311331323-0310120130031313-1113121231111320-1332122011122230-1002311331331021-2313102221223323-3132113203212103)
