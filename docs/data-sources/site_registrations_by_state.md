---
page_title: "xcsh_site_registrations_by_state landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations_by_state landing."
---

# xcsh_site_registrations_by_state landing

<a id="canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37671c9061087e4ccc244773df41c7cfb664a373126b16c8e1ac5ce161c703ac"></a>

## xcsh_site_registrations_by_state — xcsh_site_registrations_by_state / 66a53c338971 / 2

Breadcrumbs:

- xcsh_site_registrations_by_state

List Customer Edge registrations by state.

<a id="canonical-3356251a0518421f0fd22228b7ed313658fd859ccf67f78a81b32454f34947b2"></a>

## Prerequisites — xcsh_site_registrations_by_state / 66a53c338971 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-9d37c76a6c19ab0b019e0d2e38d9bbfa0696a710abc0555237834aa70843b341"></a>

## Minimal configuration — xcsh_site_registrations_by_state / 66a53c338971 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteRegistrationsByState DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations_by_state" "example" {
  state = "NOTSET"
}

output "site_registrations_by_state_result" {
  value = data.xcsh_site_registrations_by_state.example
}
```

<a id="canonical-2eb16714a43e2e8e7977dcbb9cde685feeb1d0397ce8d4333a232cfa89e050c2"></a>

## Root configuration — xcsh_site_registrations_by_state / 66a53c338971 / 5

Required root properties: `state`. Full root flags and choices appear in the property reference.

<a id="canonical-874c4150059f0287eb894a48ba5d417d9c9d92a85a0ff04d8fdcef1e289036b9"></a>

## Next pages — xcsh_site_registrations_by_state / 66a53c338971 / 6

- [Property reference](../guides/data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [Examples](../guides/data-sources--site_registrations_by_state--examples--group-001.md#canonical-f3844d2c83ae807351e728918ca3aadf7ff8666e92b43c036423d51b385ded73)
