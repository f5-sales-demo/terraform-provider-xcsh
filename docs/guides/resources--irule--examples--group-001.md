---
page_title: "xcsh_irule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_irule examples."
---

# xcsh_irule examples

<a id="canonical-127fa3038056d5fabadf3cc4ce5a2f01422cfdae2d9180f183fce32902d60d6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd3a8b22e6d6e17a3d2d8d4e5a4893cb014e0b2e3460191faf2b7e307794921a"></a>

## Examples — Examples / fdbbb60f8f33 / 2

Breadcrumbs:

- [xcsh_irule](../resources/irule.md#canonical-85fd6f4d086f69670bdf77b16e6524c8e8a0dcea34c8d6b1e850fef55d090d23)
- Examples

<a id="canonical-a90f8e30740dfe09f93fc653213947f1e420c7e54677946e5aad415a376917b0"></a>

## Complete configurations — Examples / fdbbb60f8f33 / 3

- [Resource](resources--irule--examples--group-001.md#canonical-193f4ea197ba5062a60acb3b4b8a3ca31d8bae380a90533e8573f2e9d7e3763d): valid configuration.

<a id="canonical-bfaf632b69dd3e16f0ba5c4be4715298caa366c39c1150dc4723511c8cbd81cb"></a>

## Next pages — Examples / fdbbb60f8f33 / 4

- [Resource](resources--irule--examples--group-001.md#canonical-193f4ea197ba5062a60acb3b4b8a3ca31d8bae380a90533e8573f2e9d7e3763d)
- [xcsh_irule](../resources/irule.md#canonical-85fd6f4d086f69670bdf77b16e6524c8e8a0dcea34c8d6b1e850fef55d090d23)

<a id="canonical-193f4ea197ba5062a60acb3b4b8a3ca31d8bae380a90533e8573f2e9d7e3763d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dac567f7063eaf7ce784e93a3556b81b5cc82ea51d6f146c91790f176e1009f0"></a>

## Resource — Resource / 75f6358a0df2 / 2

Breadcrumbs:

- [xcsh_irule](../resources/irule.md#canonical-85fd6f4d086f69670bdf77b16e6524c8e8a0dcea34c8d6b1e850fef55d090d23)
- [Examples](resources--irule--examples--group-001.md#canonical-127fa3038056d5fabadf3cc4ce5a2f01422cfdae2d9180f183fce32902d60d6d)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_irule/resource.tf`; digest `sha256:42cb09305e24144e92e6cea8fc9300e3823e0b03a09b749cb4b4a6120a2dca4b`.

```terraform
# Irule Resource Example
# Manages iRule in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Irule configuration
resource "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"

  description_spec = "example-value"
  irule            = "example-value"
}
```

<a id="canonical-5e4997ee40719982909d07ee82ed8e628abb89b22e6c8b99fc244613b6f8f6f6"></a>

## Next pages — Resource / 75f6358a0df2 / 3

- [Examples](resources--irule--examples--group-001.md#canonical-127fa3038056d5fabadf3cc4ce5a2f01422cfdae2d9180f183fce32902d60d6d)
- [xcsh_irule](../resources/irule.md#canonical-85fd6f4d086f69670bdf77b16e6524c8e8a0dcea34c8d6b1e850fef55d090d23)
