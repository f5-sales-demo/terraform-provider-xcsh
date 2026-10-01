---
page_title: "xcsh_site_cloud_init examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_cloud_init examples."
---

# xcsh_site_cloud_init examples

<a id="canonical-2af1fe52fef4663a7a40baa6d63bcf1d5b04630a18ac1e64cc5f4883fd0f6ca4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60af9d18a96ebec6f3a60e35efaca9a9066c3a7a6f6cca81314df6389e385e3b"></a>

## Examples — Examples / 6e991a1ba79d / 2

Breadcrumbs:

- [xcsh_site_cloud_init](../data-sources/site_cloud_init.md#canonical-bb83253dcef707bb0c00385589b2ba967da3ad61756414aa9015e45e3d75e738)
- Examples

<a id="canonical-fc1f9dcca92c00b6e5dbf7fb9c9caf64cb6e5837893defedd9f27f20d7aef1e6"></a>

## Complete configurations — Examples / 6e991a1ba79d / 3

- [Data source](data-sources--site_cloud_init--examples--group-001.md#canonical-cb35aadfca6195d4b9ff17ccb00d1a1985cf2d1b78246f3de5055d79649a6b7c): valid configuration.

<a id="canonical-69515a7381e0888996bf7e9c8ff8ffcf1a638be2fa48157ac672a40152e945fa"></a>

## Next pages — Examples / 6e991a1ba79d / 4

- [Data source](data-sources--site_cloud_init--examples--group-001.md#canonical-cb35aadfca6195d4b9ff17ccb00d1a1985cf2d1b78246f3de5055d79649a6b7c)
- [xcsh_site_cloud_init](../data-sources/site_cloud_init.md#canonical-bb83253dcef707bb0c00385589b2ba967da3ad61756414aa9015e45e3d75e738)

<a id="canonical-cb35aadfca6195d4b9ff17ccb00d1a1985cf2d1b78246f3de5055d79649a6b7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4df88570f8b1fc81b495b01ea0e6e33aea54bacb8f246b1e444d80561afc69e"></a>

## Data source — Data source / 076bdfcec1aa / 2

Breadcrumbs:

- [xcsh_site_cloud_init](../data-sources/site_cloud_init.md#canonical-bb83253dcef707bb0c00385589b2ba967da3ad61756414aa9015e45e3d75e738)
- [Examples](data-sources--site_cloud_init--examples--group-001.md#canonical-2af1fe52fef4663a7a40baa6d63bcf1d5b04630a18ac1e64cc5f4883fd0f6ca4)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_cloud_init/data-source.tf`; digest `sha256:edce0cae6fa8845014088889818debc08fd6df550bae0c288c669538441379ae`.

```terraform
# SiteCloudInit DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_cloud_init" "example" {
  provider_ref = "example-value"
  site_name    = "example-value"
}

output "site_cloud_init_result" {
  value     = data.xcsh_site_cloud_init.example
  sensitive = true
}
```

<a id="canonical-1c3b8f1ddd927d4ea00cf282097fdbb0698af97168f5258e0ab81c274a8be6b9"></a>

## Next pages — Data source / 076bdfcec1aa / 3

- [Examples](data-sources--site_cloud_init--examples--group-001.md#canonical-2af1fe52fef4663a7a40baa6d63bcf1d5b04630a18ac1e64cc5f4883fd0f6ca4)
- [xcsh_site_cloud_init](../data-sources/site_cloud_init.md#canonical-bb83253dcef707bb0c00385589b2ba967da3ad61756414aa9015e45e3d75e738)
