---
page_title: "xcsh_bot_infrastructure landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_infrastructure landing."
---

# xcsh_bot_infrastructure landing

<a id="canonical-aae65caac0c4aab59621523325e7f8203c0205db6a352330cd7f03cb7cbed28f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8573968cf36d517af703dd80d7e68090de4fe83abb14a3136ee63ba0753c52d"></a>

## xcsh_bot_infrastructure — xcsh_bot_infrastructure / 35b96318ed3e / 2

Breadcrumbs:

- xcsh_bot_infrastructure

Manages Bot Infrastructure in F5 Distributed Cloud.

<a id="canonical-74a9fbd822076e5930fcf09edbcdf0d6975aeecd64efdc0563ebbd3a7cf78e80"></a>

## Prerequisites — xcsh_bot_infrastructure / 35b96318ed3e / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1e70e0b4342b890eb8a67992f26ad45c30f11a0785848e66f2abb543ebec8922"></a>

## Minimal configuration — xcsh_bot_infrastructure / 35b96318ed3e / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-6478d376f632c37a8a46f6ef38f738981390b1bab41aa44edb768af3ce94ef37"></a>

## Root configuration — xcsh_bot_infrastructure / 35b96318ed3e / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-e4e9b5feb3bfd11fb5b0360560d9af551c193ceb5236e515e4a7073ade0bec14"></a>

## Next pages — xcsh_bot_infrastructure / 35b96318ed3e / 6

- [Property reference](../guides/resources--bot_infrastructure--reference--group-001.md#canonical-e561a7e41a50a82e651979a59c0ecb0c612a900a06a77b3600482ca3210b9251)
- [Examples](../guides/resources--bot_infrastructure--examples--group-001.md#canonical-36d6491af31a05beff7c0249315f926bfa716836ea6d7027d1e8b1a7c662ad08)
- [Import](../guides/resources--bot_infrastructure--lifecycle--group-001.md#canonical-322446d62d6e6bf6916696a98d8289b344ddfcd15baa2b2655f8e85ae2ed09d6)
- [Timeouts](../guides/resources--bot_infrastructure--lifecycle--group-001.md#canonical-d3aff639de2e4c3c392ecb7a596842f46c956d36c27e5c84298a371ee4d57843)
