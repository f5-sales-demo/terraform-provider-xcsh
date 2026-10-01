---
page_title: "xcsh_infraprotect_synchronize_configuration landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_infraprotect_synchronize_configuration landing."
---

# xcsh_infraprotect_synchronize_configuration landing

<a id="canonical-3003320123032222-1321103312000302-3220231112323333-3132221121231130-0320003131013001-3130000003113133-2111213321103223-2223330322121111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321323003330003-2301302130030102-2321202012313101-2302311112313000-3223233231203122-0230101323001023-1132000022020010-0022133211132303"></a>

## xcsh_infraprotect_synchronize_configuration — xcsh_infraprotect_synchronize_configuration / 202100011303 / 2

Breadcrumbs:

- xcsh_infraprotect_synchronize_configuration

Resource creation operation.

<a id="canonical-3000100233000321-2032200310000321-2230103333221303-1322331230102320-2012113323202302-0222311001331333-1203132100230202-2211000020101110"></a>

## Prerequisites — xcsh_infraprotect_synchronize_configuration / 202100011303 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0313211001233011-2010231223233000-3310003330110221-3130013132102232-0230312000320322-2100333121200101-3122201201031303-2310110131023332"></a>

## Minimal configuration — xcsh_infraprotect_synchronize_configuration / 202100011303 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# InfraprotectSynchronizeConfiguration Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_infraprotect_synchronize_configuration" "example" {
  config {
    namespace = "example-value"
  }
}
```

<a id="canonical-2320201123321213-2310021203031101-0133133221111132-0101220232032001-0102213333320130-3011301030231031-0201002032301300-0303322301002012"></a>

## Root configuration — xcsh_infraprotect_synchronize_configuration / 202100011303 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0220131331211302-1222202303223331-1320022021130233-3210313220032333-0010030011303201-1330322230300011-2122120223101322-3312302300301123"></a>

## Next pages — xcsh_infraprotect_synchronize_configuration / 202100011303 / 6

- [Property reference](../guides/actions--infraprotect_synchronize_configuration--reference--group-001.md#canonical-1002122200200032-2222121022000223-3022022010331111-0210321000112221-1130211020103010-3211233132011332-2332212330122311-3320121113100320)
- [Examples](../guides/actions--infraprotect_synchronize_configuration--examples--group-001.md#canonical-1121022323313230-3323102033132023-1113121222120103-1221110200113103-3211332131310301-3200223020013221-0201323130230231-0113211002201030)
- [Lifecycle](../guides/actions--infraprotect_synchronize_configuration--lifecycle--group-001.md#canonical-3332311111101223-1121130211210301-1233023221332232-0312312001210221-3003100311210321-2022233211020321-2102013301032022-1212310110313321)
