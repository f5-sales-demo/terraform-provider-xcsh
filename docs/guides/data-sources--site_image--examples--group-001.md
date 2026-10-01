---
page_title: "xcsh_site_image examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_image examples."
---

# xcsh_site_image examples

<a id="canonical-466f5d8a3bd4437c6e0cb2fd0469ed916df5aee1038a8f0caa3bece7f63a73c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e429d72dab85543f81c0b46db9be73b6aaa608bbab850bd84726d07b1faaa09"></a>

## Examples — Examples / 99f6d61b422d / 2

Breadcrumbs:

- [xcsh_site_image](../data-sources/site_image.md#canonical-ddab21760c47eb7a0955908fcfd208fc71d1b2d519a90c01d28e9ccb3de38892)
- Examples

<a id="canonical-b57e2ac61abd3da7460227154327299708f593e3a7a03899eb3fb50c661503ac"></a>

## Complete configurations — Examples / 99f6d61b422d / 3

- [Data source](data-sources--site_image--examples--group-001.md#canonical-5ea4cf94f9a48467558282af80ce35128a10e02fd551ba046cb438ee1c49013c): valid configuration.

<a id="canonical-f6aee2a612533b7314c822d989387c4a91605370c8433a10dbbc09db934526e9"></a>

## Next pages — Examples / 99f6d61b422d / 4

- [Data source](data-sources--site_image--examples--group-001.md#canonical-5ea4cf94f9a48467558282af80ce35128a10e02fd551ba046cb438ee1c49013c)
- [xcsh_site_image](../data-sources/site_image.md#canonical-ddab21760c47eb7a0955908fcfd208fc71d1b2d519a90c01d28e9ccb3de38892)

<a id="canonical-5ea4cf94f9a48467558282af80ce35128a10e02fd551ba046cb438ee1c49013c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4ba7ecd09d8e01bdc0f86baaa1a282d4f7da7d31e8a0e83305f3c92f91cf08f"></a>

## Data source — Data source / eeae1ca3222a / 2

Breadcrumbs:

- [xcsh_site_image](../data-sources/site_image.md#canonical-ddab21760c47eb7a0955908fcfd208fc71d1b2d519a90c01d28e9ccb3de38892)
- [Examples](data-sources--site_image--examples--group-001.md#canonical-466f5d8a3bd4437c6e0cb2fd0469ed916df5aee1038a8f0caa3bece7f63a73c3)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_image/data-source.tf`; digest `sha256:f749f6969ede8da33793449ff8e0b7f889a9d0fc592de94c4d8d386d5af43af8`.

```terraform
# SiteImage DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_image" "example" {
  site_name = "example-value"
}

output "site_image_result" {
  value     = data.xcsh_site_image.example
  sensitive = true
}
```

<a id="canonical-3cd836a3d9bab01bc11f6b7844b938727ea2afed53c9a14ae07403a5a85c090b"></a>

## Next pages — Data source / eeae1ca3222a / 3

- [Examples](data-sources--site_image--examples--group-001.md#canonical-466f5d8a3bd4437c6e0cb2fd0469ed916df5aee1038a8f0caa3bece7f63a73c3)
- [xcsh_site_image](../data-sources/site_image.md#canonical-ddab21760c47eb7a0955908fcfd208fc71d1b2d519a90c01d28e9ccb3de38892)
