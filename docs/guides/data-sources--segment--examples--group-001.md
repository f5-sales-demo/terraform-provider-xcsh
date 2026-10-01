---
page_title: "xcsh_segment examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_segment examples."
---

# xcsh_segment examples

<a id="canonical-248c349c9eaccfb1eb432c30ddd926ca75a9811b28d5ce7c70003f7f597195f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b617fc10e77218b760d61e2720c5dae2e789b91b19979cc4064e91f5d49f5bea"></a>

## Examples — Examples / 8cdf7d1c8bee / 2

Breadcrumbs:

- [xcsh_segment](../data-sources/segment.md#canonical-7a8b795fee80e2e6cc7579e56e701793aec6fb3c8a0be25bcde59beed1d43650)
- Examples

<a id="canonical-788351726ebd3c3b4da354f6104d2ec8eef7ea09ca4f0e1a33d489cd3bdcee26"></a>

## Complete configurations — Examples / 8cdf7d1c8bee / 3

- [Data source](data-sources--segment--examples--group-001.md#canonical-ae12e263a935790b1d99457403a79e0d8b5c063835ff7fef46143012c9a005aa): valid configuration.

<a id="canonical-b2d05acd504614688e88886445616d530b8fd39b72f8182d82fd831f667e0e39"></a>

## Next pages — Examples / 8cdf7d1c8bee / 4

- [Data source](data-sources--segment--examples--group-001.md#canonical-ae12e263a935790b1d99457403a79e0d8b5c063835ff7fef46143012c9a005aa)
- [xcsh_segment](../data-sources/segment.md#canonical-7a8b795fee80e2e6cc7579e56e701793aec6fb3c8a0be25bcde59beed1d43650)

<a id="canonical-ae12e263a935790b1d99457403a79e0d8b5c063835ff7fef46143012c9a005aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-875d42166c79c958ceebf2ee3bb10667cb5bc89d240a597e033575c110baab6d"></a>

## Data source — Data source / c17d852684e8 / 2

Breadcrumbs:

- [xcsh_segment](../data-sources/segment.md#canonical-7a8b795fee80e2e6cc7579e56e701793aec6fb3c8a0be25bcde59beed1d43650)
- [Examples](data-sources--segment--examples--group-001.md#canonical-248c349c9eaccfb1eb432c30ddd926ca75a9811b28d5ce7c70003f7f597195f5)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_segment/data-source.tf`; digest `sha256:2e9bfc7c69e50200df48c3eca01ce203db4e0516ed11b6bf987fdf485d490ac4`.

```terraform
# Segment Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Segment by name
data "xcsh_segment" "example" {
  name      = "example-segment"
  namespace = "system"
}

output "segment_id" {
  value = data.xcsh_segment.example.id
}
```

<a id="canonical-f49b63459fe5944d1786f5515f101c156e7c61e4dd4e61446b3a8bb2e577d6b7"></a>

## Next pages — Data source / c17d852684e8 / 3

- [Examples](data-sources--segment--examples--group-001.md#canonical-248c349c9eaccfb1eb432c30ddd926ca75a9811b28d5ce7c70003f7f597195f5)
- [xcsh_segment](../data-sources/segment.md#canonical-7a8b795fee80e2e6cc7579e56e701793aec6fb3c8a0be25bcde59beed1d43650)
