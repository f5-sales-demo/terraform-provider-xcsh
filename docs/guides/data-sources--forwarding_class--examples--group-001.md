---
page_title: "xcsh_forwarding_class examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_forwarding_class examples."
---

# xcsh_forwarding_class examples

<a id="canonical-75575c16923f49aa304e0a71ec8c3ee586191de8d37a0f4e564b37791041c33e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa5e15fca4df4a44e3475c0516bf03129976c493fd84e468e6c4e1631d8a4fe8"></a>

## Examples — Examples / 474e0907e16e / 2

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)
- Examples

<a id="canonical-40e30ac90a4b4fbe0be897a8dd3a0f870bcf2080ebe1431f74aac3942e2f3cd0"></a>

## Complete configurations — Examples / 474e0907e16e / 3

- [Data source](data-sources--forwarding_class--examples--group-001.md#canonical-6571161adc2d5997711278228c4853e1e1558a2b0e07f34cf52f5d258bbd003d): valid configuration.

<a id="canonical-2d752dace3b8ab5918467eea01eb588285e07ccc4a3f17bc803a0b971d96073f"></a>

## Next pages — Examples / 474e0907e16e / 4

- [Data source](data-sources--forwarding_class--examples--group-001.md#canonical-6571161adc2d5997711278228c4853e1e1558a2b0e07f34cf52f5d258bbd003d)
- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)

<a id="canonical-6571161adc2d5997711278228c4853e1e1558a2b0e07f34cf52f5d258bbd003d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0fd88cb7d33c06735a467cd101dfb81b8563ef242ec04e2eaf0ba62dd56e774"></a>

## Data source — Data source / 4b8720e87b09 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)
- [Examples](data-sources--forwarding_class--examples--group-001.md#canonical-75575c16923f49aa304e0a71ec8c3ee586191de8d37a0f4e564b37791041c33e)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_forwarding_class/data-source.tf`; digest `sha256:12747b7c8fc5a62067b4a0cf603f5a2dd940ca02738f214a627abb11e786118d`.

```terraform
# ForwardingClass Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ForwardingClass by name
data "xcsh_forwarding_class" "example" {
  name      = "example-forwarding-class"
  namespace = "staging"
}

output "forwarding_class_id" {
  value = data.xcsh_forwarding_class.example.id
}
```

<a id="canonical-07a4faf3f5536f5c9ad47bed2269be7dcd52fb6e67a797a754879f3132bd08f4"></a>

## Next pages — Data source / 4b8720e87b09 / 3

- [Examples](data-sources--forwarding_class--examples--group-001.md#canonical-75575c16923f49aa304e0a71ec8c3ee586191de8d37a0f4e564b37791041c33e)
- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)
