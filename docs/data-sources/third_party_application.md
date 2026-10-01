---
page_title: "xcsh_third_party_application landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_third_party_application landing."
---

# xcsh_third_party_application landing

<a id="canonical-b8aad6ca3bfaa99ef05f861b2a172f6bea0697789bde7db18a81450cf6385d78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ebdf7f5a2e3429f2d8aebec4311d82ff34aacb4bb2ba2e2d656bf98c4c8c7d8"></a>

## xcsh_third_party_application — xcsh_third_party_application / 1ce731cb7475 / 2

Breadcrumbs:

- xcsh_third_party_application

Manages a Third Party Application resource in F5 Distributed Cloud for third party application
specification. configuration. (read-only data source)

<a id="canonical-81b648de8fb00490b2a0fa52852f7acb2037b0704eb8ae56dd991ca21ca2ce75"></a>

## Prerequisites — xcsh_third_party_application / 1ce731cb7475 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-4d3af696b6a098f34a08627d5096d5cd114f5e9e3d5e3b1c6d2d6ffaec0ab4f3"></a>

## Minimal configuration — xcsh_third_party_application / 1ce731cb7475 / 4

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

<a id="canonical-34398933213bde003d56dd1ce02ff27ae96fc4d0178a5638914ae83490b3cd41"></a>

## Root configuration — xcsh_third_party_application / 1ce731cb7475 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-ccf1f0cdfbdcbba6c84f3a03bddce476a9626b5baebe6c735a2be67b45bf537b"></a>

## Next pages — xcsh_third_party_application / 1ce731cb7475 / 6

- [Property reference](../guides/data-sources--third_party_application--reference--group-001.md#canonical-b1ce831f083af8a3890d676267b14b10ee7b704cd2afe7eba3c1f8a503cdc8df)
- [Examples](../guides/data-sources--third_party_application--examples--group-001.md#canonical-1846a0b2efe35f7b3461c3775766d5787e6856ac42d7df49b74a9afbde5e3993)
