---
page_title: "xcsh_segment_connection landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_segment_connection landing."
---

# xcsh_segment_connection landing

<a id="canonical-0c16447dbb90eb025faf624280439f7e29eddfddfb4f25b086f9782baa1798bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6da0fae50efbd2f57ad4dc447851fe15187a25b177a66d90c43f9389fb71213"></a>

## xcsh_segment_connection — xcsh_segment_connection / ab7c2a6a9503 / 2

Breadcrumbs:

- xcsh_segment_connection

Manages a Segment Connection resource in F5 Distributed Cloud for segment connector specification.
configuration. (read-only data source)

<a id="canonical-1633da37043634c1656a397a0e7035cfe96b4c5bab70b2ce0ac8c3f28e713e3a"></a>

## Prerequisites — xcsh_segment_connection / ab7c2a6a9503 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-8ed8aa6ea6c5041713194d800620a54b446b15b9b0623f71ab4217d73c733187"></a>

## Minimal configuration — xcsh_segment_connection / ab7c2a6a9503 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-1f33b2dd9dd344bbcf5e6ddd23f268232d474dd64f5be2bae9fe52a1889fc5af"></a>

## Root configuration — xcsh_segment_connection / ab7c2a6a9503 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-cad4eaed6f92c9095265e90a2802d410578c3d8fe59f675a20869fad4236e4e3"></a>

## Next pages — xcsh_segment_connection / ab7c2a6a9503 / 6

- [Property reference](../guides/data-sources--segment_connection--reference--group-001.md#canonical-039339d9bb77663d5c6ad587cbb365f2ec5c0bac87336e75f19b8d462289de67)
- [Examples](../guides/data-sources--segment_connection--examples--group-001.md#canonical-24ec7a5c619a8fc6d498f9870b47000d9fb9103887ff32960750e2644c0b7bdf)
