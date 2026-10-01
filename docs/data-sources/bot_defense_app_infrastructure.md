---
page_title: "xcsh_bot_defense_app_infrastructure landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_defense_app_infrastructure landing."
---

# xcsh_bot_defense_app_infrastructure landing

<a id="canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c15e3997c8cd6a3187fdb55ad53269a25ab337ecb7fe833fc6cdfb97f44772b7"></a>

## xcsh_bot_defense_app_infrastructure — xcsh_bot_defense_app_infrastructure / 750608e3fc94 / 2

Breadcrumbs:

- xcsh_bot_defense_app_infrastructure

Manages Bot Defense App Infrastructure in a given namespace in F5 Distributed Cloud.

<a id="canonical-76b2a2e34fc49c482f309a5fce156d380d8231100658b472ecd2d3bf58b18c3a"></a>

## Prerequisites — xcsh_bot_defense_app_infrastructure / 750608e3fc94 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-31c4f94e169b25f06f6c6dfc0dce5ad6364c82cd24d77755008e2ee4c72733bd"></a>

## Minimal configuration — xcsh_bot_defense_app_infrastructure / 750608e3fc94 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-9e4937f0e0a6b46e1082eeaf9e76037d532b9e1baf19231b08cbc594ecf94e9c"></a>

## Root configuration — xcsh_bot_defense_app_infrastructure / 750608e3fc94 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3b8c6a34b502d0215cb6c94243936635e5e3bea9757c109c6bdb461bb528f7da"></a>

## Next pages — xcsh_bot_defense_app_infrastructure / 750608e3fc94 / 6

- [Property reference](../guides/data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-9ed1ea7dcbecfc8b88df10bacaa7a6abc8c9e7cef807d0f3b9cbbe1690640bda)
- [Examples](../guides/data-sources--bot_defense_app_infrastructure--examples--group-001.md#canonical-a3e12eaacf713ac0a9bf6f28de4c2b9cf5ec26b752ac8b94c2708f8dd92817dc)
