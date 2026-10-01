---
page_title: "xcsh_origin_pool examples"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool examples."
---

# xcsh_origin_pool examples

<a id="canonical-26f5d1de6dbe2147a167989881c2a2c0a1193c72d01e2342c928c8e90c063e5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34f2b50b5cfcb69070287e04db45d72d429f1a0cc095c22b44ae74337d48d6e8"></a>

## Examples — Examples / 656e87774654 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- Examples

<a id="canonical-a3083a2bde3cd4c300ca8d5f2686f05c72c99173e7df1814639de9637eada4be"></a>

## Complete configurations — Examples / 656e87774654 / 3

- [Data source](data-sources--origin_pool--examples--group-001.md#canonical-da3b8323f3beafe6bca83635fe293bd6d545ba37f7132dcd73bc0ae02266f087): valid configuration.

<a id="canonical-95cb904094a95f31cd9b187dacc5c2c48163c974a2cba18435c3edd9d2096531"></a>

## Next pages — Examples / 656e87774654 / 4

- [Data source](data-sources--origin_pool--examples--group-001.md#canonical-da3b8323f3beafe6bca83635fe293bd6d545ba37f7132dcd73bc0ae02266f087)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-da3b8323f3beafe6bca83635fe293bd6d545ba37f7132dcd73bc0ae02266f087"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc39f271de3c1d92d83c619374eb2993a37070ab08770d7bb1db48ac8abeb352"></a>

## Data source — Data source / f26d91f012e1 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Examples](data-sources--origin_pool--examples--group-001.md#canonical-26f5d1de6dbe2147a167989881c2a2c0a1193c72d01e2342c928c8e90c063e5e)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_origin_pool/data-source.tf`; digest `sha256:cae409ea2c1cde6f7ffac297039f556a77e9684a94017d350d26cb02cd85782a`.

```terraform
# OriginPool Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing OriginPool by name
data "xcsh_origin_pool" "example" {
  name      = "example-origin-pool"
  namespace = "staging"
}

output "origin_pool_id" {
  value = data.xcsh_origin_pool.example.id
}
```

<a id="canonical-b8f26b2079e7398da3a1cb32f8b6716de17db50e3c2efefe9ad2cf88f4e51d84"></a>

## Next pages — Data source / f26d91f012e1 / 3

- [Examples](data-sources--origin_pool--examples--group-001.md#canonical-26f5d1de6dbe2147a167989881c2a2c0a1193c72d01e2342c928c8e90c063e5e)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
