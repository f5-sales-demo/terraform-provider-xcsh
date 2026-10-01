---
page_title: "xcsh_shape_bot_defense_instance examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_shape_bot_defense_instance examples."
---

# xcsh_shape_bot_defense_instance examples

<a id="canonical-50b5f6e1abc26f913446cc3b2915bcec61e66ce812ebcafd5537d8950015d628"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-711f53ca60991ded3da24c7c6e233a32009269fa2b36acd09d8be3d59b24bc83"></a>

## Examples — Examples / 4c13f83feb36 / 2

Breadcrumbs:

- [xcsh_shape_bot_defense_instance](../data-sources/shape_bot_defense_instance.md#canonical-c06d624c324df14fa465e15a9ea25b4fefdb447d6a384de57f971c0b2de209bb)
- Examples

<a id="canonical-48ffd0dbac150d6a507375539f4e1e76d0c4c74d21396b6ab4aadeef350d058e"></a>

## Complete configurations — Examples / 4c13f83feb36 / 3

- [Data source](data-sources--shape_bot_defense_instance--examples--group-001.md#canonical-fa987775553ac5c35489aa5e4fe1486948b17b412ef99c8e8b1e5c91ef5ad04a): valid configuration.

<a id="canonical-d52086ed680a35168676d833557deefce6b3f4a0e02bfae307c800ba0957162f"></a>

## Next pages — Examples / 4c13f83feb36 / 4

- [Data source](data-sources--shape_bot_defense_instance--examples--group-001.md#canonical-fa987775553ac5c35489aa5e4fe1486948b17b412ef99c8e8b1e5c91ef5ad04a)
- [xcsh_shape_bot_defense_instance](../data-sources/shape_bot_defense_instance.md#canonical-c06d624c324df14fa465e15a9ea25b4fefdb447d6a384de57f971c0b2de209bb)

<a id="canonical-fa987775553ac5c35489aa5e4fe1486948b17b412ef99c8e8b1e5c91ef5ad04a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7bf2d897b19267138dfb9c5a6fdea1a18c153694a4f4e8e766b3d8ca855b5640"></a>

## Data source — Data source / 5f6821a37640 / 2

Breadcrumbs:

- [xcsh_shape_bot_defense_instance](../data-sources/shape_bot_defense_instance.md#canonical-c06d624c324df14fa465e15a9ea25b4fefdb447d6a384de57f971c0b2de209bb)
- [Examples](data-sources--shape_bot_defense_instance--examples--group-001.md#canonical-50b5f6e1abc26f913446cc3b2915bcec61e66ce812ebcafd5537d8950015d628)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_shape_bot_defense_instance/data-source.tf`; digest `sha256:559cb2ebbea61576b9b4c66aad8123cbb2b2816c6fc4d4e4939f3b3e386387b6`.

```terraform
# ShapeBotDefenseInstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ShapeBotDefenseInstance by name
data "xcsh_shape_bot_defense_instance" "example" {
  name      = "example-shape-bot-defense-instance"
  namespace = "staging"
}

output "shape_bot_defense_instance_id" {
  value = data.xcsh_shape_bot_defense_instance.example.id
}
```

<a id="canonical-9440623dd01054df4a8478191a0507015e7102186980a451a517f1b414492ad5"></a>

## Next pages — Data source / 5f6821a37640 / 3

- [Examples](data-sources--shape_bot_defense_instance--examples--group-001.md#canonical-50b5f6e1abc26f913446cc3b2915bcec61e66ce812ebcafd5537d8950015d628)
- [xcsh_shape_bot_defense_instance](../data-sources/shape_bot_defense_instance.md#canonical-c06d624c324df14fa465e15a9ea25b4fefdb447d6a384de57f971c0b2de209bb)
