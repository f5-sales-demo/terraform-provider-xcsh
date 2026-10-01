---
page_title: "xcsh_cloud_region examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_region examples."
---

# xcsh_cloud_region examples

<a id="canonical-572733779c4c37f3d2f9db8e1bcb5b1a692dd7dc25e9fb0b77bee03110e2a828"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4e722b82342a6d846ea75fa7f29755a21746812cf48d5ac0d4fc26eee499780"></a>

## Examples — Examples / a40c79c8412d / 2

Breadcrumbs:

- [xcsh_cloud_region](../data-sources/cloud_region.md#canonical-767677708e7c109e001971c5f50e0bf0785dfe911b57e8c6a2be2079978b8ef3)
- Examples

<a id="canonical-9d568a22ad1b38cfe7ae2197d3bc218ac28a89ce1b6a64d2ade3fd8225510f09"></a>

## Complete configurations — Examples / a40c79c8412d / 3

- [Data source](data-sources--cloud_region--examples--group-001.md#canonical-d1cde80a1b0f5eadfccc550068213a6dccaac794c7716cde8422fb3b83744434): valid configuration.

<a id="canonical-eda1d58062dff1b6b177497713863188d9114771280f2cb1cd2a8086451030fc"></a>

## Next pages — Examples / a40c79c8412d / 4

- [Data source](data-sources--cloud_region--examples--group-001.md#canonical-d1cde80a1b0f5eadfccc550068213a6dccaac794c7716cde8422fb3b83744434)
- [xcsh_cloud_region](../data-sources/cloud_region.md#canonical-767677708e7c109e001971c5f50e0bf0785dfe911b57e8c6a2be2079978b8ef3)

<a id="canonical-d1cde80a1b0f5eadfccc550068213a6dccaac794c7716cde8422fb3b83744434"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9097054c5c8d589be08aa6beda8758203f5e7904fd41ac47b522cc2483ae4591"></a>

## Data source — Data source / a76cbe4f48ad / 2

Breadcrumbs:

- [xcsh_cloud_region](../data-sources/cloud_region.md#canonical-767677708e7c109e001971c5f50e0bf0785dfe911b57e8c6a2be2079978b8ef3)
- [Examples](data-sources--cloud_region--examples--group-001.md#canonical-572733779c4c37f3d2f9db8e1bcb5b1a692dd7dc25e9fb0b77bee03110e2a828)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_region/data-source.tf`; digest `sha256:78c1e9c6e8676d021e9b37f2f6a06e7aa0e27e7d2ed71f73c7951ef98b928563`.

```terraform
# CloudRegion Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudRegion by name
data "xcsh_cloud_region" "example" {
  name      = "example-cloud-region"
  namespace = "staging"
}

output "cloud_region_id" {
  value = data.xcsh_cloud_region.example.id
}
```

<a id="canonical-8d803f49c71699636af4d0645621603bd6b9befef16dcea01d2a704aa62da81f"></a>

## Next pages — Data source / a76cbe4f48ad / 3

- [Examples](data-sources--cloud_region--examples--group-001.md#canonical-572733779c4c37f3d2f9db8e1bcb5b1a692dd7dc25e9fb0b77bee03110e2a828)
- [xcsh_cloud_region](../data-sources/cloud_region.md#canonical-767677708e7c109e001971c5f50e0bf0785dfe911b57e8c6a2be2079978b8ef3)
