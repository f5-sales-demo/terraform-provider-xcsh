---
page_title: "xcsh_certificate landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate landing."
---

# xcsh_certificate landing

<a id="canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05ba341351069fd8c1e1afd0004450c8e9730b7f105a37aed78402b819e87929"></a>

## xcsh_certificate — xcsh_certificate / c080043eb307 / 2

Breadcrumbs:

- xcsh_certificate

Manages a Certificate resource in F5 Distributed Cloud for certificate. configuration.

<a id="canonical-43747102e2b2ab753901a1aa43fd0f7bb632d40924f2e5f460e361aeff1bd060"></a>

## Prerequisites — xcsh_certificate / c080043eb307 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-f969a549fad594cd494a243684300e8de90411c86420766e117c2f4f5016db8a"></a>

## Minimal configuration — xcsh_certificate / c080043eb307 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Certificate Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Certificate by name
data "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "staging"
}

output "certificate_id" {
  value = data.xcsh_certificate.example.id
}
```

<a id="canonical-75ffafa7e356bcf6651dcca66c0acea379233776915eda4718887e4f90212086"></a>

## Root configuration — xcsh_certificate / c080043eb307 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3750c80b84559280287c0c261868c77e4e19beee77fe47ad71213e80dffb54ff"></a>

## Next pages — xcsh_certificate / c080043eb307 / 6

- [Property reference](../guides/data-sources--certificate--reference--group-001.md#canonical-2dc6dc0f3cf547c9a867ec8d56d2ac439e02f66e2f161031ecb85fc142b90dc0)
- [Examples](../guides/data-sources--certificate--examples--group-001.md#canonical-3edd44319babaa4ce9d9f018bf81a77c5532f1bef794cec91c914fb336d52473)
