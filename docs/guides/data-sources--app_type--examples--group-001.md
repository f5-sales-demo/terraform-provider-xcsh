---
page_title: "xcsh_app_type examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_type examples."
---

# xcsh_app_type examples

<a id="canonical-92963a326d3f64840ee203fb8602b7ecd63f4db1fd0ee3492c5862d8f217a946"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f08f67fdc74e5f219f86445931f4ad6a622db87378d92572c69812e019df1238"></a>

## Examples — Examples / 63e3b9a614d2 / 2

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)
- Examples

<a id="canonical-33d804129a8d22004d2277dd4ef0957336936c0d0fbc3ca3665cb51233f648d0"></a>

## Complete configurations — Examples / 63e3b9a614d2 / 3

- [Data source](data-sources--app_type--examples--group-001.md#canonical-98fa9ac90bc4516f9c17111225e189e602137b49ac15d8fc12a5004f703f111f): valid configuration.

<a id="canonical-9cd364976687596073f102dbb9d1a79247928d42f4c5a5750b6130a0e74dfab9"></a>

## Next pages — Examples / 63e3b9a614d2 / 4

- [Data source](data-sources--app_type--examples--group-001.md#canonical-98fa9ac90bc4516f9c17111225e189e602137b49ac15d8fc12a5004f703f111f)
- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)

<a id="canonical-98fa9ac90bc4516f9c17111225e189e602137b49ac15d8fc12a5004f703f111f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10400b93a51619caa7913e2f2cf08effd2383e5234d69c6c32d319212401f2a1"></a>

## Data source — Data source / a82bd1c707a8 / 2

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)
- [Examples](data-sources--app_type--examples--group-001.md#canonical-92963a326d3f64840ee203fb8602b7ecd63f4db1fd0ee3492c5862d8f217a946)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_type/data-source.tf`; digest `sha256:7219c79a69a1e857679b66c93800c9cdd60d2a0ca9b2ed275f17e432b1f04933`.

```terraform
# AppType Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppType by name
data "xcsh_app_type" "example" {
  name      = "example-app-type"
  namespace = "staging"
}

output "app_type_id" {
  value = data.xcsh_app_type.example.id
}
```

<a id="canonical-f6c9867a60fd59b2c5e0c124a5258915e568f7309105a882c4dda6722c074adc"></a>

## Next pages — Data source / a82bd1c707a8 / 3

- [Examples](data-sources--app_type--examples--group-001.md#canonical-92963a326d3f64840ee203fb8602b7ecd63f4db1fd0ee3492c5862d8f217a946)
- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)
