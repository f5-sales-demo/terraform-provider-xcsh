---
page_title: "xcsh_api_discovery examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_discovery examples."
---

# xcsh_api_discovery examples

<a id="canonical-93f0346c60bede5e7e23ca2c25083fbc71af112c5b402b831fcc7c6ea64733a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-071bf47b2eaf76d91a2d9938baafbaf70a756196f9cbb37827a6a35a5fccb4ec"></a>

## Examples — Examples / a2e900b409bd / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- Examples

<a id="canonical-96ea0e1b1768f093e2142706a6124ba9a432144dc46b10cd849c0deee4ae3243"></a>

## Complete configurations — Examples / a2e900b409bd / 3

- [Resource](resources--api_discovery--examples--group-001.md#canonical-f2f6e6a6a7c1fc8b39eb974f82fb984040401d3c60d9a0ad28f1033f33a8289c): valid configuration.

<a id="canonical-3ff47a3b17697bf0aeea6539239a3c23a1b18b1ea8887acd916afde6a7705401"></a>

## Next pages — Examples / a2e900b409bd / 4

- [Resource](resources--api_discovery--examples--group-001.md#canonical-f2f6e6a6a7c1fc8b39eb974f82fb984040401d3c60d9a0ad28f1033f33a8289c)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-f2f6e6a6a7c1fc8b39eb974f82fb984040401d3c60d9a0ad28f1033f33a8289c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44872cf3ae9b94663c77defb1ddd82c36cf6b193902dc09c4f25242b7ab624ab"></a>

## Resource — Resource / 94f8e3d56e61 / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Examples](resources--api_discovery--examples--group-001.md#canonical-93f0346c60bede5e7e23ca2c25083fbc71af112c5b402b831fcc7c6ea64733a4)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_api_discovery/resource.tf`; digest `sha256:ff9af826f29fe467445464e3b92dcd212f9666dfb00feabd76aab3e247a3f9e6`.

```terraform
# APIDiscovery Resource Example
# Manages API discovery creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APIDiscovery configuration
resource "xcsh_api_discovery" "example" {
  name      = "example-api-discovery"
  namespace = "staging"
}
```

<a id="canonical-c2c8bf2687bbfbe7da838c4131e764ab3a99065457d8bafebcc832fd92b20684"></a>

## Next pages — Resource / 94f8e3d56e61 / 3

- [Examples](resources--api_discovery--examples--group-001.md#canonical-93f0346c60bede5e7e23ca2c25083fbc71af112c5b402b831fcc7c6ea64733a4)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
