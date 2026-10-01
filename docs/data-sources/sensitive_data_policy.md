---
page_title: "xcsh_sensitive_data_policy landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_sensitive_data_policy landing."
---

# xcsh_sensitive_data_policy landing

<a id="canonical-2321213032331000-2132213131031233-2100233222312121-0113013100321030-1011013210020203-2003033231120211-3313210223013012-0130113122220200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300212100123023-3210220133321202-1220230121332313-2203030113010001-1210120130312302-1100131332330103-1033112201011120-0012100220320223"></a>

## xcsh_sensitive_data_policy — xcsh_sensitive_data_policy / 012000121333 / 2

Breadcrumbs:

- xcsh_sensitive_data_policy

Manages sensitive\_data\_policy creates a new object in the storage backend for metadata.namespace
in F5 Distributed Cloud.

<a id="canonical-0112001111200001-2110003331121023-3013303033011132-0331033022231003-1103130201033113-2002212113231201-0103102312233132-2001213331021203"></a>

## Prerequisites — xcsh_sensitive_data_policy / 012000121333 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-3202301103221232-1212100332121120-1231301321323301-2333000232003112-3322110233300013-2213112021231033-1120110101111323-1021102223233100"></a>

## Minimal configuration — xcsh_sensitive_data_policy / 012000121333 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SensitiveDataPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SensitiveDataPolicy by name
data "xcsh_sensitive_data_policy" "example" {
  name      = "example-sensitive-data-policy"
  namespace = "staging"
}

output "sensitive_data_policy_id" {
  value = data.xcsh_sensitive_data_policy.example.id
}
```

<a id="canonical-1302330222113321-2233021220300122-1332202210103121-1220011022132330-0103322200023302-1120220101101222-1010112300332103-2323030330232120"></a>

## Root configuration — xcsh_sensitive_data_policy / 012000121333 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3231332323133220-0010012230002000-3213020220123220-0233301101223002-2303202113312020-1201323030100303-0031330230032312-2301000020031333"></a>

## Next pages — xcsh_sensitive_data_policy / 012000121333 / 6

- [Property reference](../guides/data-sources--sensitive_data_policy--reference--group-001.md#canonical-2012220221230301-2110311100333312-1310211131030013-1101130223100303-2002300211120322-3233031313000120-3123212231312123-0033113022332033)
- [Examples](../guides/data-sources--sensitive_data_policy--examples--group-001.md#canonical-3130112210311310-3220231221203112-2212010312211031-2031330001130100-3203332022210322-0021312303111210-0231223120331200-2020130310211003)
