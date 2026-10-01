---
page_title: "xcsh_segment_connection examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_segment_connection examples."
---

# xcsh_segment_connection examples

<a id="canonical-24ec7a5c619a8fc6d498f9870b47000d9fb9103887ff32960750e2644c0b7bdf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c6d063665e50adcdfdf5dd755897d2d99ba64e20a871d7335a6d45c2833f51d"></a>

## Examples — Examples / 2a24f407d138 / 2

Breadcrumbs:

- [xcsh_segment_connection](../data-sources/segment_connection.md#canonical-0c16447dbb90eb025faf624280439f7e29eddfddfb4f25b086f9782baa1798bb)
- Examples

<a id="canonical-e70906c38fd399c15435cc09ce4d31eaacde8d82a7f6bbb4498f3aaeced98445"></a>

## Complete configurations — Examples / 2a24f407d138 / 3

- [Data source](data-sources--segment_connection--examples--group-001.md#canonical-07e232b149c444c345cbb65ad80c6e372a95c6ba3a9230cb0eb2af65a7c5f5c3): valid configuration.

<a id="canonical-2812542c9232bc02427e28bb4e9e755895d5381e67161e32a39c7b68b9c58793"></a>

## Next pages — Examples / 2a24f407d138 / 4

- [Data source](data-sources--segment_connection--examples--group-001.md#canonical-07e232b149c444c345cbb65ad80c6e372a95c6ba3a9230cb0eb2af65a7c5f5c3)
- [xcsh_segment_connection](../data-sources/segment_connection.md#canonical-0c16447dbb90eb025faf624280439f7e29eddfddfb4f25b086f9782baa1798bb)

<a id="canonical-07e232b149c444c345cbb65ad80c6e372a95c6ba3a9230cb0eb2af65a7c5f5c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb8c48cab4323e88f00d8891f939a62298f9cce9b6158f4d9b5329515efba1d3"></a>

## Data source — Data source / ac77b8b1f527 / 2

Breadcrumbs:

- [xcsh_segment_connection](../data-sources/segment_connection.md#canonical-0c16447dbb90eb025faf624280439f7e29eddfddfb4f25b086f9782baa1798bb)
- [Examples](data-sources--segment_connection--examples--group-001.md#canonical-24ec7a5c619a8fc6d498f9870b47000d9fb9103887ff32960750e2644c0b7bdf)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_segment_connection/data-source.tf`; digest `sha256:d15a7947bb7cd84484433c713ef71c626c3d8bc137b6e6bda822fbc845d28228`.

```terraform
# SegmentConnection Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SegmentConnection by name
data "xcsh_segment_connection" "example" {
  name      = "example-segment-connection"
  namespace = "staging"
}

output "segment_connection_id" {
  value = data.xcsh_segment_connection.example.id
}
```

<a id="canonical-513ad14d2830a8f1b46f14e741dc7594a4091e0222d971e2e5198b2d7f515a4b"></a>

## Next pages — Data source / ac77b8b1f527 / 3

- [Examples](data-sources--segment_connection--examples--group-001.md#canonical-24ec7a5c619a8fc6d498f9870b47000d9fb9103887ff32960750e2644c0b7bdf)
- [xcsh_segment_connection](../data-sources/segment_connection.md#canonical-0c16447dbb90eb025faf624280439f7e29eddfddfb4f25b086f9782baa1798bb)
