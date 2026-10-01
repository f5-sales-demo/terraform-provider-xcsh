---
page_title: "xcsh_certified_hardware landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_certified_hardware landing."
---

# xcsh_certified_hardware landing

<a id="canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-270ed927c198a135931266b18bc634385372a634a8943f6338539608dbd6d196"></a>

## xcsh_certified_hardware — xcsh_certified_hardware / 2761b7d47baa / 2

Breadcrumbs:

- xcsh_certified_hardware

Manages a Certified Hardware resource in F5 Distributed Cloud for get certified hardware object.
configuration. (read-only data source)

<a id="canonical-38c737ed10006860f679159eb187360c36c4cc95c6bd1c33439d6b0466de2044"></a>

## Prerequisites — xcsh_certified_hardware / 2761b7d47baa / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-b1934c8b8c30fb95e82e0c68e76d0d6e9632a7419f935d2575f8170b5f388c28"></a>

## Minimal configuration — xcsh_certified_hardware / 2761b7d47baa / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CertifiedHardware Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CertifiedHardware by name
data "xcsh_certified_hardware" "example" {
  name      = "example-certified-hardware"
  namespace = "staging"
}

output "certified_hardware_id" {
  value = data.xcsh_certified_hardware.example.id
}
```

<a id="canonical-6b9857ca5d201217ab5819a0be56734e9d55c3c9910d32d65a8f15ebe718b48c"></a>

## Root configuration — xcsh_certified_hardware / 2761b7d47baa / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0975cf9c938d0ad49baa852ca38082d59a6503b4d5b05505194838ae0bdf4319"></a>

## Next pages — xcsh_certified_hardware / 2761b7d47baa / 6

- [Property reference](../guides/data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- [Examples](../guides/data-sources--certified_hardware--examples--group-001.md#canonical-e51fc21fa808e165a7bb85e9397b7fac5dc3c42b5021e497581f5bd091f4a837)
