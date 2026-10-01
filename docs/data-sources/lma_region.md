---
page_title: "xcsh_lma_region landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_lma_region landing."
---

# xcsh_lma_region landing

<a id="canonical-a6178fee5bd6040179d7ca00cf2adb35c615c789980c81e85881be0d32f590d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c55e7e7cbfc92126b25aa7e0ce3c89b3cc766256e0303d553620386287029fb"></a>

## xcsh_lma_region — xcsh_lma_region / cc7ed4792348 / 2

Breadcrumbs:

- xcsh_lma_region

Manages a Lma Region resource in F5 Distributed Cloud for lma region specification. configuration.
(read-only data source)

<a id="canonical-98057ea21338c0fe1da373a50e7450cd89ce486cf626880696652306ebea8054"></a>

## Prerequisites — xcsh_lma_region / cc7ed4792348 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-cd578262731b53bd240b50d76fe557312b69bb81bbb38c37c87ca3d644c7c150"></a>

## Minimal configuration — xcsh_lma_region / cc7ed4792348 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# LmaRegion Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing LmaRegion by name
data "xcsh_lma_region" "example" {
  name      = "example-lma-region"
  namespace = "staging"
}

output "lma_region_id" {
  value = data.xcsh_lma_region.example.id
}
```

<a id="canonical-44f6f76de5ffc9506c6520016220ee8ea36bd2fde86cdb4f58043a72d4bd3902"></a>

## Root configuration — xcsh_lma_region / cc7ed4792348 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-ea4c4db0328d1a9446e6aa0206341724553472f1a88c11086814eacbf4ea4674"></a>

## Next pages — xcsh_lma_region / cc7ed4792348 / 6

- [Property reference](../guides/data-sources--lma_region--reference--group-001.md#canonical-0b90d404964c736cedf478efd222a6a2158aa0fd49a3a83c8fce45a45fbe181a)
- [Examples](../guides/data-sources--lma_region--examples--group-001.md#canonical-e1517c3955fe4b14745c199b25cd6429335c8c35e4b10881e5dd6a82c17af7fc)
