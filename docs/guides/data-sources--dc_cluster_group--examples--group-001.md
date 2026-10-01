---
page_title: "xcsh_dc_cluster_group examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dc_cluster_group examples."
---

# xcsh_dc_cluster_group examples

<a id="canonical-4a9ea2aeea2342c033a4b640f81a52986acdf9e373a1343b192be86253d100c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4e6f35858c8c8ffe455ef1c5c5c91ca6af7e88861fe8746d3f05daadc313e9f"></a>

## Examples — Examples / 57665d6a6bce / 2

Breadcrumbs:

- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-a0240865dc3b4d8fc779d80a3fc0926a687cf9901210cdea3e46d6b423dc41b7)
- Examples

<a id="canonical-43a27ed6576e6a303b785c9e6f190a2b3bd046d1db076034ed09ed4f6058a186"></a>

## Complete configurations — Examples / 57665d6a6bce / 3

- [Data source](data-sources--dc_cluster_group--examples--group-001.md#canonical-afdc5204791f94d6224f5db291489b66d522e0cec6e58c31a1d40575f2892102): valid configuration.

<a id="canonical-ba4bc49195ef13531751862f2bde4453359d03c0ca440d1478f2c1bc5cbcf91c"></a>

## Next pages — Examples / 57665d6a6bce / 4

- [Data source](data-sources--dc_cluster_group--examples--group-001.md#canonical-afdc5204791f94d6224f5db291489b66d522e0cec6e58c31a1d40575f2892102)
- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-a0240865dc3b4d8fc779d80a3fc0926a687cf9901210cdea3e46d6b423dc41b7)

<a id="canonical-afdc5204791f94d6224f5db291489b66d522e0cec6e58c31a1d40575f2892102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d827082460977b4f033d20d3aa25490c9b0a85396dcb90bb47c7bf82b09229d"></a>

## Data source — Data source / 449c0335942f / 2

Breadcrumbs:

- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-a0240865dc3b4d8fc779d80a3fc0926a687cf9901210cdea3e46d6b423dc41b7)
- [Examples](data-sources--dc_cluster_group--examples--group-001.md#canonical-4a9ea2aeea2342c033a4b640f81a52986acdf9e373a1343b192be86253d100c9)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dc_cluster_group/data-source.tf`; digest `sha256:82b040ec9e5742587a1d4d2feb953ac78e9dca3a538c42aff66a3020a25fe4c5`.

```terraform
# DcClusterGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DcClusterGroup by name
data "xcsh_dc_cluster_group" "example" {
  name      = "example-dc-cluster-group"
  namespace = "system"
}

output "dc_cluster_group_id" {
  value = data.xcsh_dc_cluster_group.example.id
}
```

<a id="canonical-eb5338a9cddd2090e95341bcb91b08c919e8fa6abde57dbb20e210fc60f9e1b7"></a>

## Next pages — Data source / 449c0335942f / 3

- [Examples](data-sources--dc_cluster_group--examples--group-001.md#canonical-4a9ea2aeea2342c033a4b640f81a52986acdf9e373a1343b192be86253d100c9)
- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-a0240865dc3b4d8fc779d80a3fc0926a687cf9901210cdea3e46d6b423dc41b7)
