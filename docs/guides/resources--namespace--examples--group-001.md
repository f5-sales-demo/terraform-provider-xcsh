---
page_title: "xcsh_namespace examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_namespace examples."
---

# xcsh_namespace examples

<a id="canonical-1b8bc34f331033b05402414f3e116c5e5532af0c28d25550b4e20bc0b278be2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-934aaa6dfb4ee37a2d06779bf1de60b5780ea858ace9bb75dcac27922e5aa266"></a>

## Examples — Examples / 28a96535397d / 2

Breadcrumbs:

- [xcsh_namespace](../resources/namespace.md#canonical-5d4d79984896dc184c2d678bc72c561a372fa7660188fc55cde1c40e575c98b4)
- Examples

<a id="canonical-2e5fceb1eec79dfd7c27497cc09a14d59d4e958ce2fbb1fcd84d6971add46720"></a>

## Complete configurations — Examples / 28a96535397d / 3

- [Resource](resources--namespace--examples--group-001.md#canonical-40e2ec23c3fd40d074c7449bcda5951a5a5a41fab88a8edb03f0b44f80bfe9f0): valid configuration.

<a id="canonical-7bcd5dca5c56e8e52b9989f2c7fc27bd605164671a4c32ef71674faf931b9795"></a>

## Next pages — Examples / 28a96535397d / 4

- [Resource](resources--namespace--examples--group-001.md#canonical-40e2ec23c3fd40d074c7449bcda5951a5a5a41fab88a8edb03f0b44f80bfe9f0)
- [xcsh_namespace](../resources/namespace.md#canonical-5d4d79984896dc184c2d678bc72c561a372fa7660188fc55cde1c40e575c98b4)

<a id="canonical-40e2ec23c3fd40d074c7449bcda5951a5a5a41fab88a8edb03f0b44f80bfe9f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cfe5f285c93fbb8f908a8dfe530d03f65aed0620c64b56795fc88a94d36b034"></a>

## Resource — Resource / 07f0a4be0866 / 2

Breadcrumbs:

- [xcsh_namespace](../resources/namespace.md#canonical-5d4d79984896dc184c2d678bc72c561a372fa7660188fc55cde1c40e575c98b4)
- [Examples](resources--namespace--examples--group-001.md#canonical-1b8bc34f331033b05402414f3e116c5e5532af0c28d25550b4e20bc0b278be2c)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_namespace/resource.tf`; digest `sha256:95871f5bf448147bdd129adaa3690bae9af907a5a988941728cc5c87a1b9e514`.

```terraform
# Namespace Resource Example
# Manages new namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Credentials are supplied externally.
provider "xcsh" {}

# Basic Namespace configuration
resource "xcsh_namespace" "this" {
  name = "example-namespace"
}
```

<a id="canonical-bbc3448681f660386202bfae916261da9a3a6884a3b33cd4bdb697bc9635f4ae"></a>

## Next pages — Resource / 07f0a4be0866 / 3

- [Examples](resources--namespace--examples--group-001.md#canonical-1b8bc34f331033b05402414f3e116c5e5532af0c28d25550b4e20bc0b278be2c)
- [xcsh_namespace](../resources/namespace.md#canonical-5d4d79984896dc184c2d678bc72c561a372fa7660188fc55cde1c40e575c98b4)
