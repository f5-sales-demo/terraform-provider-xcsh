---
page_title: "xcsh_lma_region examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_lma_region examples."
---

# xcsh_lma_region examples

<a id="canonical-e1517c3955fe4b14745c199b25cd6429335c8c35e4b10881e5dd6a82c17af7fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09a083514aed2d7ef0649f19d91ef6c79fa8facdd95af719e1a46a1d5e16deb2"></a>

## Examples — Examples / a27288d35a1b / 2

Breadcrumbs:

- [xcsh_lma_region](../data-sources/lma_region.md#canonical-a6178fee5bd6040179d7ca00cf2adb35c615c789980c81e85881be0d32f590d1)
- Examples

<a id="canonical-a69176877cabaf87b8ecb69353ac00627a95655001cd1bdae85d333b2f009fa5"></a>

## Complete configurations — Examples / a27288d35a1b / 3

- [Data source](data-sources--lma_region--examples--group-001.md#canonical-6d142543f719d7b9ec9b100da10141dfab4e3c3282793b834804c86917d24c2d): valid configuration.

<a id="canonical-f37a70355fbb7fea29e0507f0a25206023b96bfd5c050a7e6ee241179801ced0"></a>

## Next pages — Examples / a27288d35a1b / 4

- [Data source](data-sources--lma_region--examples--group-001.md#canonical-6d142543f719d7b9ec9b100da10141dfab4e3c3282793b834804c86917d24c2d)
- [xcsh_lma_region](../data-sources/lma_region.md#canonical-a6178fee5bd6040179d7ca00cf2adb35c615c789980c81e85881be0d32f590d1)

<a id="canonical-6d142543f719d7b9ec9b100da10141dfab4e3c3282793b834804c86917d24c2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-834a4ddd04c0f1e56d25ebc890f11a3ec78ecf2a9fa42db5ddaa264ffa1e453e"></a>

## Data source — Data source / 2b7320d2ac54 / 2

Breadcrumbs:

- [xcsh_lma_region](../data-sources/lma_region.md#canonical-a6178fee5bd6040179d7ca00cf2adb35c615c789980c81e85881be0d32f590d1)
- [Examples](data-sources--lma_region--examples--group-001.md#canonical-e1517c3955fe4b14745c199b25cd6429335c8c35e4b10881e5dd6a82c17af7fc)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_lma_region/data-source.tf`; digest `sha256:6e44cbfbf0065cf8526c29152a263e7eb100be46cd3be2092eca1d815333bf89`.

```terraform
# LmaRegion Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing LmaRegion by name
data "xcsh_lma_region" "example" {
  name      = "example-lma-region"
  namespace = "staging"
}

output "lma_region_id" {
  value = data.xcsh_lma_region.example.id
}
```

<a id="canonical-acb5541988b80cb54b0028685b467247bb1cd084f630af18f2bcb7328406dfc4"></a>

## Next pages — Data source / 2b7320d2ac54 / 3

- [Examples](data-sources--lma_region--examples--group-001.md#canonical-e1517c3955fe4b14745c199b25cd6429335c8c35e4b10881e5dd6a82c17af7fc)
- [xcsh_lma_region](../data-sources/lma_region.md#canonical-a6178fee5bd6040179d7ca00cf2adb35c615c789980c81e85881be0d32f590d1)
