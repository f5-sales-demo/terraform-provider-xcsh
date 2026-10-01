---
page_title: "xcsh_srv6_network_slice examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_srv6_network_slice examples."
---

# xcsh_srv6_network_slice examples

<a id="canonical-b2b705b70a59b78ad13e1fbdde0f0af4ac8ab79472c63dc8818b8fb18de8c11a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70c30f2557da1d56a4f6ced013f5a294fe8bf7f224aac76961c99f41b4ebed86"></a>

## Examples — Examples / de8b55c9cf85 / 2

Breadcrumbs:

- [xcsh_srv6_network_slice](../data-sources/srv6_network_slice.md#canonical-01eb91def3bc993a7ab14854fcd32915db33e69c6c189dc0599a4ae5db267c6b)
- Examples

<a id="canonical-745a3999d9676cd62ee744dbbcc6d4249a6a4e52c74da98c0ef34b4a34fc7ac9"></a>

## Complete configurations — Examples / de8b55c9cf85 / 3

- [Data source](data-sources--srv6_network_slice--examples--group-001.md#canonical-e2bc7de1d0f82d16d33e0660cd441530644ef7255ac3a0259e16a66c662b6945): valid configuration.

<a id="canonical-ae1be97667c0c779ce40a33ebc011975e792cbe03f9f1b0ab67d7cad736f2c46"></a>

## Next pages — Examples / de8b55c9cf85 / 4

- [Data source](data-sources--srv6_network_slice--examples--group-001.md#canonical-e2bc7de1d0f82d16d33e0660cd441530644ef7255ac3a0259e16a66c662b6945)
- [xcsh_srv6_network_slice](../data-sources/srv6_network_slice.md#canonical-01eb91def3bc993a7ab14854fcd32915db33e69c6c189dc0599a4ae5db267c6b)

<a id="canonical-e2bc7de1d0f82d16d33e0660cd441530644ef7255ac3a0259e16a66c662b6945"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97d0a5e420d43104edca4c10cccce7f30cc832498f15baa59dbe52852d3bf669"></a>

## Data source — Data source / f0852b9d078d / 2

Breadcrumbs:

- [xcsh_srv6_network_slice](../data-sources/srv6_network_slice.md#canonical-01eb91def3bc993a7ab14854fcd32915db33e69c6c189dc0599a4ae5db267c6b)
- [Examples](data-sources--srv6_network_slice--examples--group-001.md#canonical-b2b705b70a59b78ad13e1fbdde0f0af4ac8ab79472c63dc8818b8fb18de8c11a)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_srv6_network_slice/data-source.tf`; digest `sha256:5f3b6cfc9cb152dc2c604228803d1e4acdb80bbc9c71bb42003bba14ed6bbdd1`.

```terraform
# Srv6NetworkSlice Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Srv6NetworkSlice by name
data "xcsh_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"
}

output "srv6_network_slice_id" {
  value = data.xcsh_srv6_network_slice.example.id
}
```

<a id="canonical-e1f3cd3ae8cd373b0002687e457f86ba9cb373c95b1d4fcc3506f915a00d89a2"></a>

## Next pages — Data source / f0852b9d078d / 3

- [Examples](data-sources--srv6_network_slice--examples--group-001.md#canonical-b2b705b70a59b78ad13e1fbdde0f0af4ac8ab79472c63dc8818b8fb18de8c11a)
- [xcsh_srv6_network_slice](../data-sources/srv6_network_slice.md#canonical-01eb91def3bc993a7ab14854fcd32915db33e69c6c189dc0599a4ae5db267c6b)
