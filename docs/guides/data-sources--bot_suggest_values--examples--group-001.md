---
page_title: "xcsh_bot_suggest_values examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_suggest_values examples."
---

# xcsh_bot_suggest_values examples

<a id="canonical-cb1677c61520b70852bab053740849ce32413907dccd3d431b685980ca92525d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a4ba202e8befad42e1e0a162dbdab3e3d5930e71a57abbc127a8ddbedc09254"></a>

## Examples — Examples / b60b87037a24 / 2

Breadcrumbs:

- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md#canonical-eb5fdcd8e576fd0893458000f2f7792dfce1074e73bded33582900b2e388a473)
- Examples

<a id="canonical-14f8455c64d0894abab29c1b15d43cc4e057d124e251de076189d62943c121b0"></a>

## Complete configurations — Examples / b60b87037a24 / 3

- [Data source](data-sources--bot_suggest_values--examples--group-001.md#canonical-37dbb1444d5746a3c653652dd8d4e4ae811b752d59ed4499c4518a9aae450d00): valid configuration.

<a id="canonical-b9e3344361bbcccd377da13a0f09d06e0f51b58407cbd0c5ab06b4cbf5e38821"></a>

## Next pages — Examples / b60b87037a24 / 4

- [Data source](data-sources--bot_suggest_values--examples--group-001.md#canonical-37dbb1444d5746a3c653652dd8d4e4ae811b752d59ed4499c4518a9aae450d00)
- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md#canonical-eb5fdcd8e576fd0893458000f2f7792dfce1074e73bded33582900b2e388a473)

<a id="canonical-37dbb1444d5746a3c653652dd8d4e4ae811b752d59ed4499c4518a9aae450d00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e11e8f1a38b8120fc90b06b0b57908ea3b0a6929f4dee1156b5531b61254235f"></a>

## Data source — Data source / d684bd1ee2dd / 2

Breadcrumbs:

- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md#canonical-eb5fdcd8e576fd0893458000f2f7792dfce1074e73bded33582900b2e388a473)
- [Examples](data-sources--bot_suggest_values--examples--group-001.md#canonical-cb1677c61520b70852bab053740849ce32413907dccd3d431b685980ca92525d)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_suggest_values/data-source.tf`; digest `sha256:c6c1b3639717b7c716bae9d4f7857fddc676935e02dc229f788549aecb0a97ae`.

```terraform
# BotSuggestValues DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_suggest_values" "example" {
  namespace = "example-value"
}

output "bot_suggest_values_result" {
  value = data.xcsh_bot_suggest_values.example
}
```

<a id="canonical-17f7d51d74170a4beffbc596318768f6c5a3db5f436232b46f29bad624a18157"></a>

## Next pages — Data source / d684bd1ee2dd / 3

- [Examples](data-sources--bot_suggest_values--examples--group-001.md#canonical-cb1677c61520b70852bab053740849ce32413907dccd3d431b685980ca92525d)
- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md#canonical-eb5fdcd8e576fd0893458000f2f7792dfce1074e73bded33582900b2e388a473)
