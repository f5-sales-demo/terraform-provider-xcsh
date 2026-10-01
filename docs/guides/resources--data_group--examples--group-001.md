---
page_title: "xcsh_data_group examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_group examples."
---

# xcsh_data_group examples

<a id="canonical-1d293247dc1bbafe31dca58399cf1cf445f54893e37cb987fa1aae1f27900bb7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88e237385a7ff7d7e7d3a5aa987cc3558f0679e272f4b4551f73aa0c4090aeec"></a>

## Examples — Examples / 8cd9b6ab7b70 / 2

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md#canonical-8fa44c0b348749f6556f32f242093dd90beea888fbbeb6165a302dd0d3a951cd)
- Examples

<a id="canonical-d0d19f503a002eff8caa040e4a792059769387ddb686d3abc8bbff1feb226963"></a>

## Complete configurations — Examples / 8cd9b6ab7b70 / 3

- [Resource](resources--data_group--examples--group-001.md#canonical-a36ed74244fd41ca02b38a37ad37497f65fd54c9b47faac549bf1d861587909f): valid configuration.

<a id="canonical-3520a0d2b410f6cb2136414b6973edf26f3d684df0ad87ec4f4781b99d7f3660"></a>

## Next pages — Examples / 8cd9b6ab7b70 / 4

- [Resource](resources--data_group--examples--group-001.md#canonical-a36ed74244fd41ca02b38a37ad37497f65fd54c9b47faac549bf1d861587909f)
- [xcsh_data_group](../resources/data_group.md#canonical-8fa44c0b348749f6556f32f242093dd90beea888fbbeb6165a302dd0d3a951cd)

<a id="canonical-a36ed74244fd41ca02b38a37ad37497f65fd54c9b47faac549bf1d861587909f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6569ba0b43dcdf945918e7b027fc391187ca440d1a7b519c84b46c4ba878e095"></a>

## Resource — Resource / 41a8bbaaedca / 2

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md#canonical-8fa44c0b348749f6556f32f242093dd90beea888fbbeb6165a302dd0d3a951cd)
- [Examples](resources--data_group--examples--group-001.md#canonical-1d293247dc1bbafe31dca58399cf1cf445f54893e37cb987fa1aae1f27900bb7)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_data_group/resource.tf`; digest `sha256:1bceaa79a1e1b840495f41f1bf5adac4d3dc173e5487643b627a4cc77fc7f690`.

```terraform
# DataGroup Resource Example
# Manages data group in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataGroup configuration
resource "xcsh_data_group" "example" {
  name      = "example-data-group"
  namespace = "staging"
}
```

<a id="canonical-4b5b84f3ebff286040a2cb3bbd3fd7125fb266b14a93e7dad3b092c306104565"></a>

## Next pages — Resource / 41a8bbaaedca / 3

- [Examples](resources--data_group--examples--group-001.md#canonical-1d293247dc1bbafe31dca58399cf1cf445f54893e37cb987fa1aae1f27900bb7)
- [xcsh_data_group](../resources/data_group.md#canonical-8fa44c0b348749f6556f32f242093dd90beea888fbbeb6165a302dd0d3a951cd)
