---
page_title: "xcsh_bot_infrastructure examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_infrastructure examples."
---

# xcsh_bot_infrastructure examples

<a id="canonical-36d6491af31a05beff7c0249315f926bfa716836ea6d7027d1e8b1a7c662ad08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16eaf80903c5c0eb379144d7095e89a502426628267d94b1c006dbf9e38ece7b"></a>

## Examples — Examples / c4eeca49e72f / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-aae65caac0c4aab59621523325e7f8203c0205db6a352330cd7f03cb7cbed28f)
- Examples

<a id="canonical-fc7f849e0230d34d5a9743e23e2d8983d404858b490f11aa5bccbb69c82340e0"></a>

## Complete configurations — Examples / c4eeca49e72f / 3

- [Resource](resources--bot_infrastructure--examples--group-001.md#canonical-aae21dfff9b58b5c7864c73549d6e6699d088efde3bac0434bbd29f03bbdaf1c): valid configuration.

<a id="canonical-78cc81428675b680add267239a9190c28777e1ab397a7172010963424c7bde50"></a>

## Next pages — Examples / c4eeca49e72f / 4

- [Resource](resources--bot_infrastructure--examples--group-001.md#canonical-aae21dfff9b58b5c7864c73549d6e6699d088efde3bac0434bbd29f03bbdaf1c)
- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-aae65caac0c4aab59621523325e7f8203c0205db6a352330cd7f03cb7cbed28f)

<a id="canonical-aae21dfff9b58b5c7864c73549d6e6699d088efde3bac0434bbd29f03bbdaf1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09ef7498f245c7aa00f6adb8e4214651f3822d4d9d056d195f4330cb1a283717"></a>

## Resource — Resource / f88313ad7775 / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-aae65caac0c4aab59621523325e7f8203c0205db6a352330cd7f03cb7cbed28f)
- [Examples](resources--bot_infrastructure--examples--group-001.md#canonical-36d6491af31a05beff7c0249315f926bfa716836ea6d7027d1e8b1a7c662ad08)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bot_infrastructure/resource.tf`; digest `sha256:8bbb63f8f4c5acb98b0adb6309584a0770310129d290ad3dfbcb6667452aa9e4`.

```terraform
# BotInfrastructure Resource Example
# Manages Bot Infrastructure in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BotInfrastructure configuration
resource "xcsh_bot_infrastructure" "example" {
  name      = "example-bot-infrastructure"
  namespace = "staging"
}
```

<a id="canonical-9286f290aa772993a54ce341c4547743db80b4b0bd23b3795a5a56252b5b8788"></a>

## Next pages — Resource / f88313ad7775 / 3

- [Examples](resources--bot_infrastructure--examples--group-001.md#canonical-36d6491af31a05beff7c0249315f926bfa716836ea6d7027d1e8b1a7c662ad08)
- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-aae65caac0c4aab59621523325e7f8203c0205db6a352330cd7f03cb7cbed28f)
