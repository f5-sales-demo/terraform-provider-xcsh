---
page_title: "xcsh_site examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site examples."
---

# xcsh_site examples

<a id="canonical-5dcb0b1f7736d456c2ca9b82b4d4af76c43b08bfe5a9a9e4f3f086ed99c8371a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a9c3a743384a45e8a259ecf753172d5cbdd2104f0b32688dcbe3633ff344784"></a>

## Examples — Examples / 02920ab38d89 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- Examples

<a id="canonical-8906929a6ee0b04843120cc1ea48c72c5ad007f3aeb2d08e06cc595898e2d4f1"></a>

## Complete configurations — Examples / 02920ab38d89 / 3

- [Data source](data-sources--site--examples--group-001.md#canonical-d6db56fed405231480b9ae29d2d4259314844583d410f1cadcbc0016da722b92): valid configuration.

<a id="canonical-d2d7a059b321f9de91ff966f1a8a095d84de8df15d93d3433b4f631a67d8d86a"></a>

## Next pages — Examples / 02920ab38d89 / 4

- [Data source](data-sources--site--examples--group-001.md#canonical-d6db56fed405231480b9ae29d2d4259314844583d410f1cadcbc0016da722b92)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-d6db56fed405231480b9ae29d2d4259314844583d410f1cadcbc0016da722b92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c01138a8b945c3fd355a73e700cde0f510a963b44e2074d2072ed62a49ed7e28"></a>

## Data source — Data source / ff4ec68cebda / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Examples](data-sources--site--examples--group-001.md#canonical-5dcb0b1f7736d456c2ca9b82b4d4af76c43b08bfe5a9a9e4f3f086ed99c8371a)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site/data-source.tf`; digest `sha256:3e384d82af36bf37fe3ff1c335c8824183b28374dc3c5fb9a56016a0094aa348`.

```terraform
# Site Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Site by name
data "xcsh_site" "example" {
  name      = "example-site"
  namespace = "staging"
}

output "site_id" {
  value = data.xcsh_site.example.id
}
```

<a id="canonical-880918abc07d833bfd0aefa18e1511091f27a5baf517d561ff9d1d19b311a087"></a>

## Next pages — Data source / ff4ec68cebda / 3

- [Examples](data-sources--site--examples--group-001.md#canonical-5dcb0b1f7736d456c2ca9b82b4d4af76c43b08bfe5a9a9e4f3f086ed99c8371a)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
