---
page_title: "xcsh_bot_defense_app_infrastructure examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_defense_app_infrastructure examples."
---

# xcsh_bot_defense_app_infrastructure examples

<a id="canonical-a3e12eaacf713ac0a9bf6f28de4c2b9cf5ec26b752ac8b94c2708f8dd92817dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b17534932e685e75885643f2262214f357a33d8c4dfd390b51ece25f6eddd8ff"></a>

## Examples — Examples / f9043cf77704 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)
- Examples

<a id="canonical-5d93ed27550f7c0c6c5aae7408a8da5e798af2b2e89974493ac95f6bac97089b"></a>

## Complete configurations — Examples / f9043cf77704 / 3

- [Data source](data-sources--bot_defense_app_infrastructure--examples--group-001.md#canonical-f4a81c6beaa5594253e6aae11e3e2608c6d68ab22418f4a12f2de74092c348e6): valid configuration.

<a id="canonical-bc8eab3d39f9e131a029bb57469aef1d72d3b6b343d69279abd2a98cd6d8b587"></a>

## Next pages — Examples / f9043cf77704 / 4

- [Data source](data-sources--bot_defense_app_infrastructure--examples--group-001.md#canonical-f4a81c6beaa5594253e6aae11e3e2608c6d68ab22418f4a12f2de74092c348e6)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)

<a id="canonical-f4a81c6beaa5594253e6aae11e3e2608c6d68ab22418f4a12f2de74092c348e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94c24888721ab1593e5c3574688276382e2c79713e8e95e3d53fc0fab526d117"></a>

## Data source — Data source / d41a5f793e54 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)
- [Examples](data-sources--bot_defense_app_infrastructure--examples--group-001.md#canonical-a3e12eaacf713ac0a9bf6f28de4c2b9cf5ec26b752ac8b94c2708f8dd92817dc)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_defense_app_infrastructure/data-source.tf`; digest `sha256:9d1dcbef65f7cbdf8cc89c39b723d01c2f54dd36e7a2d02f47f90c88068c73e0`.

```terraform
# BotDefenseAppInfrastructure Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotDefenseAppInfrastructure by name
data "xcsh_bot_defense_app_infrastructure" "example" {
  name      = "example-bot-defense-app-infrastructure"
  namespace = "staging"
}

output "bot_defense_app_infrastructure_id" {
  value = data.xcsh_bot_defense_app_infrastructure.example.id
}
```

<a id="canonical-a2dfc74606ad7698f60a4aad78fb750ee37e054c6a4b496a2cd285e6c0cdf43f"></a>

## Next pages — Data source / d41a5f793e54 / 3

- [Examples](data-sources--bot_defense_app_infrastructure--examples--group-001.md#canonical-a3e12eaacf713ac0a9bf6f28de4c2b9cf5ec26b752ac8b94c2708f8dd92817dc)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)
