---
page_title: "xcsh_securemesh_site examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site examples."
---

# xcsh_securemesh_site examples

<a id="canonical-d56096c5e963bd57d646cd37b499d4475dfe3831c2ce681cd4a104aa7f0fd5ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79cf7574cbad963bb8f2c7e9dcdf56cb2bf4eab1b8037a4c8998f66e1e9d9a38"></a>

## Examples — Examples / 2c67a4f02cda / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- Examples

<a id="canonical-8e5fcbe0df7205632d1ba73f303c4aff6fd8313f4491956bea5523f595cc31bd"></a>

## Complete configurations — Examples / 2c67a4f02cda / 3

- [Resource](resources--securemesh_site--examples--group-001.md#canonical-212a816b79d70f0cb121e5c00c058f336080099f3223cb69d82d8a1a0a3de040): valid configuration.

<a id="canonical-d021287f60b5cc65f298a1f78a42a7c0177c8334b3dc93d5a4bad034c0fc0a7a"></a>

## Next pages — Examples / 2c67a4f02cda / 4

- [Resource](resources--securemesh_site--examples--group-001.md#canonical-212a816b79d70f0cb121e5c00c058f336080099f3223cb69d82d8a1a0a3de040)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-212a816b79d70f0cb121e5c00c058f336080099f3223cb69d82d8a1a0a3de040"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f26a48f9fb845b4df3d6b5460e61152e6a797e1d04404f9d055f9f3c666916eb"></a>

## Resource — Resource / effbb2bbaf03 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Examples](resources--securemesh_site--examples--group-001.md#canonical-d56096c5e963bd57d646cd37b499d4475dfe3831c2ce681cd4a104aa7f0fd5ee)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_securemesh_site/resource.tf`; digest `sha256:d12d5ac4d45f23102fc4cf022413c9bdb3fd5f59873115079a2e3a8fe4e2c28a`.

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

<a id="canonical-d23340bd8620ccbb3a8edb31f58808255fc52c8af874b908f0c483ba17e9eccc"></a>

## Next pages — Resource / effbb2bbaf03 / 3

- [Examples](resources--securemesh_site--examples--group-001.md#canonical-d56096c5e963bd57d646cd37b499d4475dfe3831c2ce681cd4a104aa7f0fd5ee)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
