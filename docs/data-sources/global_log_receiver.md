---
page_title: "xcsh_global_log_receiver landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver landing."
---

# xcsh_global_log_receiver landing

<a id="canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303333012233131-0021333133031010-0213330300003332-1120021002030203-3003121111323323-3001211213313332-0211133103321232-0230122013333101"></a>

## xcsh_global_log_receiver — xcsh_global_log_receiver / 222030311003 / 2

Breadcrumbs:

- xcsh_global_log_receiver

Manages new Global Log Receiver object in F5 Distributed Cloud.

<a id="canonical-3330123103313001-2213000213220222-3000322220123230-2011201330031120-1330200222200133-2213323030101022-0021122112000330-0230012112332000"></a>

## Prerequisites — xcsh_global_log_receiver / 222030311003 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2020113101230000-3211121133121133-3102020313233012-2023113123232223-3030002230330201-0102321113321203-0320122020123220-2331031002122120"></a>

## Minimal configuration — xcsh_global_log_receiver / 222030311003 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# GlobalLogReceiver Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing GlobalLogReceiver by name
data "xcsh_global_log_receiver" "example" {
  name      = "example-global-log-receiver"
  namespace = "staging"
}

output "global_log_receiver_id" {
  value = data.xcsh_global_log_receiver.example.id
}
```

<a id="canonical-1102302301330321-3131333230211322-1332021230123322-2320313203013011-0120222202210201-3103122323320322-1201222021100022-2131231003311101"></a>

## Root configuration — xcsh_global_log_receiver / 222030311003 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0021031101022210-1020112132021212-2003111202312222-3323303200302132-0211332211311311-1010123003311113-2310010021210322-1211011300112122"></a>

## Next pages — xcsh_global_log_receiver / 222030311003 / 6

- [Property reference](../guides/data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [Examples](../guides/data-sources--global_log_receiver--examples--group-001.md#canonical-3023311232200010-0133300122022312-3102212013022001-1130103102302322-3213013111200203-3211331221323312-0213002213312012-2002030133101203)
