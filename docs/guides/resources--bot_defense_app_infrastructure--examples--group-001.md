---
page_title: "xcsh_bot_defense_app_infrastructure examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_defense_app_infrastructure examples."
---

# xcsh_bot_defense_app_infrastructure examples

<a id="canonical-9f069c3c61ed129432bbf3f92c7f65f9447467862f48adffb246e9dad7f21ef4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc44fb6089b6d0a6aff079aad7d2b66b2dd3dc614e71a79fc8dda91e392e583a"></a>

## Examples — Examples / b3033738c947 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)
- Examples

<a id="canonical-011653e876565354d871485900b58b89fb88dd6b6fd01784c78d869129447682"></a>

## Complete configurations — Examples / b3033738c947 / 3

- [Resource](resources--bot_defense_app_infrastructure--examples--group-001.md#canonical-d468e5128f7312f097c3f2d13b8d1ebe8dbdb1a7dcb7b2e54db16ddc2ebdd698): valid configuration.

<a id="canonical-58371dd1f8ac3926838c2dfb76e668a2742c9551b3b1c1d1947f0016e7153077"></a>

## Next pages — Examples / b3033738c947 / 4

- [Resource](resources--bot_defense_app_infrastructure--examples--group-001.md#canonical-d468e5128f7312f097c3f2d13b8d1ebe8dbdb1a7dcb7b2e54db16ddc2ebdd698)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)

<a id="canonical-d468e5128f7312f097c3f2d13b8d1ebe8dbdb1a7dcb7b2e54db16ddc2ebdd698"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f109283a928f6b7b6ff4882223b7a4d5185aa665b20defbc5e5e42737c2ca42"></a>

## Resource — Resource / c71259849509 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)
- [Examples](resources--bot_defense_app_infrastructure--examples--group-001.md#canonical-9f069c3c61ed129432bbf3f92c7f65f9447467862f48adffb246e9dad7f21ef4)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bot_defense_app_infrastructure/resource.tf`; digest `sha256:3d3368cb414ef6890130a5c64346a138d87252424badf8402f264c1532265f55`.

```terraform
# BotDefenseAppInfrastructure Resource Example
# Manages Bot Defense App Infrastructure in a given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BotDefenseAppInfrastructure configuration
resource "xcsh_bot_defense_app_infrastructure" "example" {
  name      = "example-bot-defense-app-infrastructure"
  namespace = "staging"
}
```

<a id="canonical-f5f0ba926028a3a911b4e79f912c04a9a91c18d93084819afd27ef1242ac684e"></a>

## Next pages — Resource / c71259849509 / 3

- [Examples](resources--bot_defense_app_infrastructure--examples--group-001.md#canonical-9f069c3c61ed129432bbf3f92c7f65f9447467862f48adffb246e9dad7f21ef4)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)
