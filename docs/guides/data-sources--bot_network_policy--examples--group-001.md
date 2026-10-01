---
page_title: "xcsh_bot_network_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_network_policy examples."
---

# xcsh_bot_network_policy examples

<a id="canonical-022ba32b36308c5407c9dc966cfe8e3dcde2eef6d9858ccd4c90ceb17cfa7059"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb4f919a916859c09493d3c4829fea6ebb57bcb874924a8502ed397e9e7a5505"></a>

## Examples — Examples / 12091dbf87b4 / 2

Breadcrumbs:

- [xcsh_bot_network_policy](../data-sources/bot_network_policy.md#canonical-1db4aff0b95edbb987adc1f675e48bc99801201bf045f697bf744647c8e90762)
- Examples

<a id="canonical-b75e66102ac7fd3b630affaf30916019008123f4bb31867bf0d06a15e81b49dc"></a>

## Complete configurations — Examples / 12091dbf87b4 / 3

- [Data source](data-sources--bot_network_policy--examples--group-001.md#canonical-4d4f2ac539eed470fdb9a33c4be449dd483ee229360f77e006439d2568fe67dd): valid configuration.

<a id="canonical-ec6a83a1394603195fc48d26b21a54a5c9fb3704b93838021024d8227d485d3d"></a>

## Next pages — Examples / 12091dbf87b4 / 4

- [Data source](data-sources--bot_network_policy--examples--group-001.md#canonical-4d4f2ac539eed470fdb9a33c4be449dd483ee229360f77e006439d2568fe67dd)
- [xcsh_bot_network_policy](../data-sources/bot_network_policy.md#canonical-1db4aff0b95edbb987adc1f675e48bc99801201bf045f697bf744647c8e90762)

<a id="canonical-4d4f2ac539eed470fdb9a33c4be449dd483ee229360f77e006439d2568fe67dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3137801b917e4c16daa72c6c00b2a6a201ac21f7c3ebaa027bd7b3e340309429"></a>

## Data source — Data source / 92c8740e8239 / 2

Breadcrumbs:

- [xcsh_bot_network_policy](../data-sources/bot_network_policy.md#canonical-1db4aff0b95edbb987adc1f675e48bc99801201bf045f697bf744647c8e90762)
- [Examples](data-sources--bot_network_policy--examples--group-001.md#canonical-022ba32b36308c5407c9dc966cfe8e3dcde2eef6d9858ccd4c90ceb17cfa7059)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_network_policy/data-source.tf`; digest `sha256:b2aec065924f60efb4e75a54937e4cdc96ee0d4e38bbfe5651736f1aa7a5ffe9`.

```terraform
# BotNetworkPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotNetworkPolicy by name
data "xcsh_bot_network_policy" "example" {
  name      = "example-bot-network-policy"
  namespace = "staging"
}

output "bot_network_policy_id" {
  value = data.xcsh_bot_network_policy.example.id
}
```

<a id="canonical-3ae496ef801012189dc870530c8956c842f137adf41abbb69f0fc6697c75837c"></a>

## Next pages — Data source / 92c8740e8239 / 3

- [Examples](data-sources--bot_network_policy--examples--group-001.md#canonical-022ba32b36308c5407c9dc966cfe8e3dcde2eef6d9858ccd4c90ceb17cfa7059)
- [xcsh_bot_network_policy](../data-sources/bot_network_policy.md#canonical-1db4aff0b95edbb987adc1f675e48bc99801201bf045f697bf744647c8e90762)
