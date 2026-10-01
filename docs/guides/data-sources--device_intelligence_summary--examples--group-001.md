---
page_title: "xcsh_device_intelligence_summary examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_summary examples."
---

# xcsh_device_intelligence_summary examples

<a id="canonical-f97a8f9c74d388e7b38af72d6e8c5345b93515d764c0d104c74440b774b7228c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ca01511d79bc7e8e16793a23d56f33ac9e4a8405c42ad00a60fb8e1e71e4ab6"></a>

## Examples — Examples / be0bed480b75 / 2

Breadcrumbs:

- [xcsh_device_intelligence_summary](../data-sources/device_intelligence_summary.md#canonical-713f059bbd4c71f5f978f47874a8258e6f5fed5f744224aa8b9189615842fff9)
- Examples

<a id="canonical-c52a6416a98c46811052a63e49e938df0050ebe77cacc257385ced9b46840401"></a>

## Complete configurations — Examples / be0bed480b75 / 3

- [Data source](data-sources--device_intelligence_summary--examples--group-001.md#canonical-e24aaff876b122936dae9756d318f2868877a86b05ca2bf172729897721ce159): valid configuration.

<a id="canonical-b9aa28c6ca187f6858e5c82e87637f4ffcfd9f4e92371f364e8e384a0a8613bb"></a>

## Next pages — Examples / be0bed480b75 / 4

- [Data source](data-sources--device_intelligence_summary--examples--group-001.md#canonical-e24aaff876b122936dae9756d318f2868877a86b05ca2bf172729897721ce159)
- [xcsh_device_intelligence_summary](../data-sources/device_intelligence_summary.md#canonical-713f059bbd4c71f5f978f47874a8258e6f5fed5f744224aa8b9189615842fff9)

<a id="canonical-e24aaff876b122936dae9756d318f2868877a86b05ca2bf172729897721ce159"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b10a1423672527f5c847381ef7e5348aedd9ffbb1c861ffc0ee27f43f6f319ce"></a>

## Data source — Data source / 610fe60aa973 / 2

Breadcrumbs:

- [xcsh_device_intelligence_summary](../data-sources/device_intelligence_summary.md#canonical-713f059bbd4c71f5f978f47874a8258e6f5fed5f744224aa8b9189615842fff9)
- [Examples](data-sources--device_intelligence_summary--examples--group-001.md#canonical-f97a8f9c74d388e7b38af72d6e8c5345b93515d764c0d104c74440b774b7228c)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_summary/data-source.tf`; digest `sha256:a3b5af58871245c4f3794bc2ed886a4421185806d4cb6b7c6eb5328aa94877da`.

```terraform
# DeviceIntelligenceSummary DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_summary" "example" {
  namespace = "example-value"
}

output "device_intelligence_summary_result" {
  value = data.xcsh_device_intelligence_summary.example
}
```

<a id="canonical-c551bad0b8e867e0737a5cb3141856bd65c4740871ec0b78a2ae025b757f3f54"></a>

## Next pages — Data source / 610fe60aa973 / 3

- [Examples](data-sources--device_intelligence_summary--examples--group-001.md#canonical-f97a8f9c74d388e7b38af72d6e8c5345b93515d764c0d104c74440b774b7228c)
- [xcsh_device_intelligence_summary](../data-sources/device_intelligence_summary.md#canonical-713f059bbd4c71f5f978f47874a8258e6f5fed5f744224aa8b9189615842fff9)
