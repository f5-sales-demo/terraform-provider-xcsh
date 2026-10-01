---
page_title: "xcsh_aws_tgw_site examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site examples."
---

# xcsh_aws_tgw_site examples

<a id="canonical-df6730039bad354a53ad7c0e0d66872ca27b44704e247e143cbf5dce3628592f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c1ef3dffded5769c00027a596627d1bb194debf3ffa7e68d2e10104ea80ba07"></a>

## Examples — Examples / 6ff46d81a376 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- Examples

<a id="canonical-afec1b9706c30a417ebec5c8cc95a6e14b403749e2f909fd3b1dc0ed55bca046"></a>

## Complete configurations — Examples / 6ff46d81a376 / 3

- [Data source](data-sources--aws_tgw_site--examples--group-001.md#canonical-84b22fde4e7b5f331e943d9d769888929be1b84e830881b435a50b7d3b85dbb2): valid configuration.

<a id="canonical-89e07d0b8b7b29d7eec5c6425193e063ab8ff79e7a8a2ebc10e55d7462e86c62"></a>

## Next pages — Examples / 6ff46d81a376 / 4

- [Data source](data-sources--aws_tgw_site--examples--group-001.md#canonical-84b22fde4e7b5f331e943d9d769888929be1b84e830881b435a50b7d3b85dbb2)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-84b22fde4e7b5f331e943d9d769888929be1b84e830881b435a50b7d3b85dbb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5f236ff61a103a1d4d9e391cfc8810dad79bbfa8e14f6458a19a190e095c55c"></a>

## Data source — Data source / f3a11dedf6b2 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Examples](data-sources--aws_tgw_site--examples--group-001.md#canonical-df6730039bad354a53ad7c0e0d66872ca27b44704e247e143cbf5dce3628592f)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_aws_tgw_site/data-source.tf`; digest `sha256:aef0f2c4d7539fcdc86d3bd364e4b9a22f4422d4d0595364ea8d26b8fc2a6cf7`.

```terraform
# AWSTGWSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AWSTGWSite by name
data "xcsh_aws_tgw_site" "example" {
  name      = "example-aws-tgw-site"
  namespace = "staging"
}

output "aws_tgw_site_id" {
  value = data.xcsh_aws_tgw_site.example.id
}
```

<a id="canonical-f2006a0eb6136e7dc048d0604404208d74e019c1706278f3132bf4feb6593081"></a>

## Next pages — Data source / f3a11dedf6b2 / 3

- [Examples](data-sources--aws_tgw_site--examples--group-001.md#canonical-df6730039bad354a53ad7c0e0d66872ca27b44704e247e143cbf5dce3628592f)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
