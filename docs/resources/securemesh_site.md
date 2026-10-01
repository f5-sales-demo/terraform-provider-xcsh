---
page_title: "xcsh_securemesh_site landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site landing."
---

# xcsh_securemesh_site landing

<a id="canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b85b5427e9d93ab1fb71783c07dfe84e0b7da4de935293a5f03b18ebb4f2fa9f"></a>

## xcsh_securemesh_site — xcsh_securemesh_site / 8c8d7e6231c2 / 2

Breadcrumbs:

- xcsh_securemesh_site

Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with
distributed security.

<a id="canonical-44323eaf8a35cb1abf9ebfe949fda88bc1053085a1a7d2afc1038c4fd6ddaff2"></a>

## Prerequisites — xcsh_securemesh_site / 8c8d7e6231c2 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-11ad4fd490305cce997d000bd0961e0994abdae9ae2ab0aeaf08a74776175d20"></a>

## Minimal configuration — xcsh_securemesh_site / 8c8d7e6231c2 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecuremeshSite Resource Example
# Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with distributed security.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecuremeshSite configuration
resource "xcsh_securemesh_site" "example" {
  name      = "example-securemesh-site"
  namespace = "staging"

  volterra_certified_hw = "example-value"
}
```

<a id="canonical-80a373a0729e397f16b41ee44ab92774b79cda8031a012c12e07d7b24f5b6181"></a>

## Root configuration — xcsh_securemesh_site / 8c8d7e6231c2 / 5

Required root properties: `name`, `namespace`, `volterra_certified_hw`. Full root flags and choices appear in the property reference.

<a id="canonical-464becd0575be2423e0943370f8c57d9e89b7e5b994a542d2179f6c8de17bea0"></a>

## Next pages — xcsh_securemesh_site / 8c8d7e6231c2 / 6

- [Property reference](../guides/resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [Examples](../guides/resources--securemesh_site--examples--group-001.md#canonical-d56096c5e963bd57d646cd37b499d4475dfe3831c2ce681cd4a104aa7f0fd5ee)
- [Import](../guides/resources--securemesh_site--lifecycle--group-001.md#canonical-daa4f8647cbc3a356c3bdafc396e941f53148631304d2652e2fd4b464d4d341c)
- [Timeouts](../guides/resources--securemesh_site--lifecycle--group-001.md#canonical-6b58f266a654b9709361f84249c6759da4ef8a314b567a5c10589d5402167adb)
