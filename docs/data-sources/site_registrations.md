---
page_title: "xcsh_site_registrations landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations landing."
---

# xcsh_site_registrations landing

<a id="canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-622bc159617d17103b5c3facc2bc08d054c595149eee637973e32858e1878d57"></a>

## xcsh_site_registrations — xcsh_site_registrations / 661e81cdc3fa / 2

Breadcrumbs:

- xcsh_site_registrations

List Customer Edge registrations.

<a id="canonical-7e5c9134549aafb9baa5623fde1917603e008b381d7e21e01a4d127109e1db34"></a>

## Prerequisites — xcsh_site_registrations / 661e81cdc3fa / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-ba69836d0b810cdb6b429c7d4cd3bfc121632269f8823ff17391f4875aa348d0"></a>

## Minimal configuration — xcsh_site_registrations / 661e81cdc3fa / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteRegistrations DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations" "example" {
  namespace = "example-value"
}

output "site_registrations_result" {
  value = data.xcsh_site_registrations.example
}
```

<a id="canonical-b7d0b469c9b98558285c9d10aae1dfcae0cd5900900fe98d516bca97e0308a79"></a>

## Root configuration — xcsh_site_registrations / 661e81cdc3fa / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2202711dd710f8d6a018129790b4dbf8f70320f878b9dd10671740c9da606285"></a>

## Next pages — xcsh_site_registrations / 661e81cdc3fa / 6

- [Property reference](../guides/data-sources--site_registrations--reference--group-001.md#canonical-e9ce9f37e9d5fb7e5425df57d652e74015a83817a0e9796c5622ad59128f54cc)
- [Examples](../guides/data-sources--site_registrations--examples--group-001.md#canonical-ce2131793ed98a34945fac766d85be28b3f9df7b9ed07d195dbeb37477f644f7)
