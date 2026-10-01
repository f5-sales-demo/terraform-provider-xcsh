---
page_title: "xcsh_segment landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_segment landing."
---

# xcsh_segment landing

<a id="canonical-115b348cc2da26907c0d3a2200fee1c45353481468d739f4370fd6bf8ca9a410"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b87d34d642c93f734ccd643d108ec1c1b3a7a59acf4351f8e76823614deaf6f"></a>

## xcsh_segment — xcsh_segment / fbb84e0c6676 / 2

Breadcrumbs:

- xcsh_segment

Manages a Segment resource in F5 Distributed Cloud for segment. configuration.

<a id="canonical-5c6bef26a3dd8a4b6fc01fcf64d822fe8b267856d1105d07ce42fa18cfdf47a6"></a>

## Prerequisites — xcsh_segment / fbb84e0c6676 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-36af67194e07fa7355c5ff2c1f6dd8f74725c7cbf6721143e24ce93ea8a50311"></a>

## Minimal configuration — xcsh_segment / fbb84e0c6676 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Segment Resource Example
# Manages a Segment resource in F5 Distributed Cloud for segment.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Segment configuration
resource "xcsh_segment" "example" {
  name      = "example-segment"
  namespace = "system"
}
```

<a id="canonical-3ee7b26f82bd808da5ee4f56ce1d3a7ace1bed4b10dc80b9fe8df9614218ccde"></a>

## Root configuration — xcsh_segment / fbb84e0c6676 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-66facd3ba25d7df72ee9e3805390b5f6a5e085ee007cd11acc7241d9457628f5"></a>

## Next pages — xcsh_segment / fbb84e0c6676 / 6

- [Property reference](../guides/resources--segment--reference--group-001.md#canonical-3fd2ddf918ccce02c13c05b12c882f76272305dd050891bce28f013df6b7f704)
- [Examples](../guides/resources--segment--examples--group-001.md#canonical-0418839d33d311e131d6c8096eb96c74b36716c6c6f8f95cbcf0cdbf37919cff)
- [Import](../guides/resources--segment--lifecycle--group-001.md#canonical-41092275632d3f0289dab80de7bc6fb98dec83c9ace3e6df0234a4718e97b503)
- [Timeouts](../guides/resources--segment--lifecycle--group-001.md#canonical-22318fba13a82551dbeb74f91235b2cd9207b08cbdd54a072a3325d7ab102927)
