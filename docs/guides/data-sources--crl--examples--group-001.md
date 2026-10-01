---
page_title: "xcsh_crl examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_crl examples."
---

# xcsh_crl examples

<a id="canonical-b3f67ac2969dbe28fd8ecd6d912bf4b2163c31694d52b729046d2acbf3da56ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-972b441b1c2d914c686fedee6eb32ec78fed7b6e084420093398bd7273a603e2"></a>

## Examples — Examples / fa5b0acce0aa / 2

Breadcrumbs:

- [xcsh_crl](../data-sources/crl.md#canonical-73f9576382ba605e0c52a90b24f3c47e46f2fa2fe6a252eff356e5dcb2c0e943)
- Examples

<a id="canonical-f8e8850d2faca1e905082727f4fae75d71a25d74e3fc7ae2d8f309c274f8f147"></a>

## Complete configurations — Examples / fa5b0acce0aa / 3

- [Data source](data-sources--crl--examples--group-001.md#canonical-c866f335da3fac7fb6960d058dc00b1b97cd02ef74868304923246e13677687a): valid configuration.

<a id="canonical-a52a6a567a09be4921a2d84465365f2db617320677c9a9398f85224e38ba0607"></a>

## Next pages — Examples / fa5b0acce0aa / 4

- [Data source](data-sources--crl--examples--group-001.md#canonical-c866f335da3fac7fb6960d058dc00b1b97cd02ef74868304923246e13677687a)
- [xcsh_crl](../data-sources/crl.md#canonical-73f9576382ba605e0c52a90b24f3c47e46f2fa2fe6a252eff356e5dcb2c0e943)

<a id="canonical-c866f335da3fac7fb6960d058dc00b1b97cd02ef74868304923246e13677687a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f79e812c039d7e5e990edf383c170ee7f3213e1549eb1fbc8157c88febdf639"></a>

## Data source — Data source / c12919622dd0 / 2

Breadcrumbs:

- [xcsh_crl](../data-sources/crl.md#canonical-73f9576382ba605e0c52a90b24f3c47e46f2fa2fe6a252eff356e5dcb2c0e943)
- [Examples](data-sources--crl--examples--group-001.md#canonical-b3f67ac2969dbe28fd8ecd6d912bf4b2163c31694d52b729046d2acbf3da56ec)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_crl/data-source.tf`; digest `sha256:7250b8110197f6a15add9c7d7bb8ba0193edd8faf34e61501e055b382c218d8a`.

```terraform
# CRL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CRL by name
data "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"
}

output "crl_id" {
  value = data.xcsh_crl.example.id
}
```

<a id="canonical-7952785259672d420a554e8145f4f68c2150c9e147e5a24b6546dda27ea3db2c"></a>

## Next pages — Data source / c12919622dd0 / 3

- [Examples](data-sources--crl--examples--group-001.md#canonical-b3f67ac2969dbe28fd8ecd6d912bf4b2163c31694d52b729046d2acbf3da56ec)
- [xcsh_crl](../data-sources/crl.md#canonical-73f9576382ba605e0c52a90b24f3c47e46f2fa2fe6a252eff356e5dcb2c0e943)
