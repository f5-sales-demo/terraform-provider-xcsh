---
page_title: "xcsh_virtual_site examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_virtual_site examples."
---

# xcsh_virtual_site examples

<a id="canonical-5bd147df1752d218cbb1af93aa5ac43f096f47268571b17c6a4a9b3e92a11bd6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-783982428a369f9fec4765216c9440f0394ff3fc805af8f8ed7fe57b4b3a9615"></a>

## Examples — Examples / 26573980a49c / 2

Breadcrumbs:

- [xcsh_virtual_site](../data-sources/virtual_site.md#canonical-4d9c6e801f11ee0a4dcb7c662b2cdc5dffca2ee94e25d78ffb25ff05789e2980)
- Examples

<a id="canonical-c7b5b1aed99f5dd7ec96b9772f844b5aa938af53b9385148c67597b3c05cd29b"></a>

## Complete configurations — Examples / 26573980a49c / 3

- [Data source](data-sources--virtual_site--examples--group-001.md#canonical-407c42c431b2c264f6b8a225844a6fa0102143c58b9461c36bd5933994b53f56): valid configuration.

<a id="canonical-849f5154ef08285d299bf8303b762c19aca1382fc861d408ec6ae43ddd5c3a1f"></a>

## Next pages — Examples / 26573980a49c / 4

- [Data source](data-sources--virtual_site--examples--group-001.md#canonical-407c42c431b2c264f6b8a225844a6fa0102143c58b9461c36bd5933994b53f56)
- [xcsh_virtual_site](../data-sources/virtual_site.md#canonical-4d9c6e801f11ee0a4dcb7c662b2cdc5dffca2ee94e25d78ffb25ff05789e2980)

<a id="canonical-407c42c431b2c264f6b8a225844a6fa0102143c58b9461c36bd5933994b53f56"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2668789eb19fc42f882816cc85a7baf5b577d7e3430eb3a63e68e5518c0cd5e3"></a>

## Data source — Data source / 2a2625df640f / 2

Breadcrumbs:

- [xcsh_virtual_site](../data-sources/virtual_site.md#canonical-4d9c6e801f11ee0a4dcb7c662b2cdc5dffca2ee94e25d78ffb25ff05789e2980)
- [Examples](data-sources--virtual_site--examples--group-001.md#canonical-5bd147df1752d218cbb1af93aa5ac43f096f47268571b17c6a4a9b3e92a11bd6)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_virtual_site/data-source.tf`; digest `sha256:319f1c9eed60a6b55a583796c0636dc04effd666a48047cf3b9134b1863d1477`.

```terraform
# VirtualSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualSite by name
data "xcsh_virtual_site" "example" {
  name      = "example-virtual-site"
  namespace = "staging"
}

output "virtual_site_id" {
  value = data.xcsh_virtual_site.example.id
}
```

<a id="canonical-99a72146a16dfbcb7134237958515cf38c7f532afbfeea7de826dd19407b75df"></a>

## Next pages — Data source / 2a2625df640f / 3

- [Examples](data-sources--virtual_site--examples--group-001.md#canonical-5bd147df1752d218cbb1af93aa5ac43f096f47268571b17c6a4a9b3e92a11bd6)
- [xcsh_virtual_site](../data-sources/virtual_site.md#canonical-4d9c6e801f11ee0a4dcb7c662b2cdc5dffca2ee94e25d78ffb25ff05789e2980)
