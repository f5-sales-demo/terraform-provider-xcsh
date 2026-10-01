---
page_title: "xcsh_site_registrations_by_site landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations_by_site landing."
---

# xcsh_site_registrations_by_site landing

<a id="canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4363e85678f6ff5d7bc7d8a8d5765897dd0d27ed7ba04d8d6da9cf75465e55e"></a>

## xcsh_site_registrations_by_site — xcsh_site_registrations_by_site / 9197600ad2bd / 2

Breadcrumbs:

- xcsh_site_registrations_by_site

List registrations for a Customer Edge site.

<a id="canonical-82a8020cabbf9f39f794735a4fe5f2246594bb50fcef6a7d766a5c23f2acaeae"></a>

## Prerequisites — xcsh_site_registrations_by_site / 9197600ad2bd / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-395ab5b578314dfed1b5472752af1d0db21f27d78d4c9838599c98aec9343736"></a>

## Minimal configuration — xcsh_site_registrations_by_site / 9197600ad2bd / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteRegistrationsBySite DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations_by_site" "example" {
  site_name = "example-value"
}

output "site_registrations_by_site_result" {
  value = data.xcsh_site_registrations_by_site.example
}
```

<a id="canonical-4cd36fb3812b2da67173c6dde3df7adf47e22fcb17304588ba0634004b93b65c"></a>

## Root configuration — xcsh_site_registrations_by_site / 9197600ad2bd / 5

Required root properties: `site_name`. Full root flags and choices appear in the property reference.

<a id="canonical-826b80d9685ad853d6a4bb45b7b44bfed156e398ef8f1f48d7f7d90f7ae8f5bb"></a>

## Next pages — xcsh_site_registrations_by_site / 9197600ad2bd / 6

- [Property reference](../guides/data-sources--site_registrations_by_site--reference--group-001.md#canonical-1248603df02ae3e61f32ee8b6fee3170e79477a7407bcfde66af4d7c062772c0)
- [Examples](../guides/data-sources--site_registrations_by_site--examples--group-001.md#canonical-9dfe40fa0db7c48dfa71d784e8447fe909be11dbb249be1956696f9f35cef12f)
