---
page_title: "xcsh_access_active_sessions_terminate landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_sessions_terminate landing."
---

# xcsh_access_active_sessions_terminate landing

<a id="canonical-2232333201313313-0211033213132122-1112020213103320-3212121022102212-1220232111200021-2230200120313012-2130013111233211-1310221022123111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023102322110100-1110213201303111-0031000012323111-2131030122321001-2302320101201223-2102013220103322-3302131211231312-0202220023320012"></a>

## xcsh_access_active_sessions_terminate — xcsh_access_active_sessions_terminate / 232022311130 / 2

Breadcrumbs:

- xcsh_access_active_sessions_terminate

Resource creation operation.

<a id="canonical-3011321122032301-3200310031002303-0212222002310303-1103333112103003-0332130020123000-1230211030020221-1323223103202222-3132312331310321"></a>

## Prerequisites — xcsh_access_active_sessions_terminate / 232022311130 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3200132322120020-3321012310211133-0031032012110200-0023233310321110-1112013211002210-3211310332100210-2122001321031230-0020110100113030"></a>

## Minimal configuration — xcsh_access_active_sessions_terminate / 232022311130 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AccessActiveSessionsTerminate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_access_active_sessions_terminate" "example" {
  config {
    namespace = "example-value"
  }
}
```

<a id="canonical-1302123321203012-0322123313101200-2222110300230123-1030233121032013-1030110022322001-2130232332320202-3230010012232231-1120312121231302"></a>

## Root configuration — xcsh_access_active_sessions_terminate / 232022311130 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3332131022013012-3111322323132033-0033203013103100-3231220331310223-1130111023000110-2233331012232101-0230101020313012-3201011210303013"></a>

## Next pages — xcsh_access_active_sessions_terminate / 232022311130 / 6

- [Property reference](../guides/actions--access_active_sessions_terminate--reference--group-001.md#canonical-1011310031122232-1101123230020033-3121112303102021-3303203200231201-2030111000030313-0022320120213221-2333100310201032-0210322102310330)
- [Examples](../guides/actions--access_active_sessions_terminate--examples--group-001.md#canonical-1233211303213321-3101030230133003-1313322330111221-1220211022000320-3122330010001302-0000122213122000-2213020333311111-1103212212202312)
- [Lifecycle](../guides/actions--access_active_sessions_terminate--lifecycle--group-001.md#canonical-3312120011223001-0313020321201221-3033001230232021-0211132333013010-0323313100022300-1112023303322203-3133212221022112-2213321013203221)
