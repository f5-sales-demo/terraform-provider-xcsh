---
page_title: "xcsh_bot_defense_app_infrastructure landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_defense_app_infrastructure landing."
---

# xcsh_bot_defense_app_infrastructure landing

<a id="canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecfc7a0c347ddc18516ea35039540f65e407df189081041384dc664da23a9536"></a>

## xcsh_bot_defense_app_infrastructure — xcsh_bot_defense_app_infrastructure / 34908bd73bcb / 2

Breadcrumbs:

- xcsh_bot_defense_app_infrastructure

Manages Bot Defense App Infrastructure in a given namespace in F5 Distributed Cloud.

<a id="canonical-a6e1efd505c0be5758706a43af38284a1530eefe209b0cf9561efa59a5e73779"></a>

## Prerequisites — xcsh_bot_defense_app_infrastructure / 34908bd73bcb / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-e3f7251b95501a5703c46da368fcd09743c675efd8bc137402a060a2f9f0a7a0"></a>

## Minimal configuration — xcsh_bot_defense_app_infrastructure / 34908bd73bcb / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-af1534bd5af7c53b03ce7965c71e4e9e5561dfb3d2c560f65f26f06d83e64e34"></a>

## Root configuration — xcsh_bot_defense_app_infrastructure / 34908bd73bcb / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1a36223b272ce17a592bf827336ce282201cee90e961fa5d66ed41d591179f71"></a>

## Next pages — xcsh_bot_defense_app_infrastructure / 34908bd73bcb / 6

- [Property reference](../guides/resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ab45a834092534caee16f8d6620b6949464c5f1b4a9001ce716fc0b8c3b7aec4)
- [Examples](../guides/resources--bot_defense_app_infrastructure--examples--group-001.md#canonical-9f069c3c61ed129432bbf3f92c7f65f9447467862f48adffb246e9dad7f21ef4)
- [Import](../guides/resources--bot_defense_app_infrastructure--lifecycle--group-001.md#canonical-cdf3f61f578246a0ebe8cbeb466d80aa0ce1dc47ddcd239217a5c3dcc3989680)
- [Timeouts](../guides/resources--bot_defense_app_infrastructure--lifecycle--group-001.md#canonical-ad627479e2ec2f13af7959e424facb55240235a4b16e9ef77d4547a1c93df710)
