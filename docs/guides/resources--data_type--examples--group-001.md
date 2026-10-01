---
page_title: "xcsh_data_type examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_type examples."
---

# xcsh_data_type examples

<a id="canonical-f6fcb5b6928c9a403046c371b1e2041d0b724dd72cc231fccb0851e9f5346050"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7d4a784dbb37926fa400d05db9147f008cb945c357ab3a29fdab5b33348ad71"></a>

## Examples — Examples / a74bbe8a581d / 2

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
- Examples

<a id="canonical-61ee5c5d6b6f949d4b9fb3530572086579441f27918033b957e3891550ac6c66"></a>

## Complete configurations — Examples / a74bbe8a581d / 3

- [Resource](resources--data_type--examples--group-001.md#canonical-db1290047050e502085e4ee3fb0cf725eb7a5c8a454ee05fd11469809b81bbea): valid configuration.

<a id="canonical-24dccdb563aeda5818cc93e0f7f95a570540d4d9e74282e4953c4a2ff8c62fa1"></a>

## Next pages — Examples / a74bbe8a581d / 4

- [Resource](resources--data_type--examples--group-001.md#canonical-db1290047050e502085e4ee3fb0cf725eb7a5c8a454ee05fd11469809b81bbea)
- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)

<a id="canonical-db1290047050e502085e4ee3fb0cf725eb7a5c8a454ee05fd11469809b81bbea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16e9313c9fad2c543fa312cd76fccbfdbfa7d32de4ba974518fbb73f25beb1e8"></a>

## Resource — Resource / 4433128b238c / 2

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
- [Examples](resources--data_type--examples--group-001.md#canonical-f6fcb5b6928c9a403046c371b1e2041d0b724dd72cc231fccb0851e9f5346050)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_data_type/resource.tf`; digest `sha256:292681f3b6c18ec2ffbe2be1a0e17beb9c57b44e1777befb42567a6c3a8f12ee`.

```terraform
# DataType Resource Example
# Manages data_type creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataType configuration
resource "xcsh_data_type" "example" {
  name      = "example-data-type"
  namespace = "staging"
}
```

<a id="canonical-c3162b3dd1e0d132cc3604512b873257bd57256f23cf53ba9103e598d83903cd"></a>

## Next pages — Resource / 4433128b238c / 3

- [Examples](resources--data_type--examples--group-001.md#canonical-f6fcb5b6928c9a403046c371b1e2041d0b724dd72cc231fccb0851e9f5346050)
- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
