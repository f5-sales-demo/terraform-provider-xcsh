---
page_title: "xcsh_bot_infrastructure landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_infrastructure landing."
---

# xcsh_bot_infrastructure landing

<a id="canonical-877bbb7d337b693d557733ad6caa7f00c7f8b096aeb76ab4bf994720688abb64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e19f0f270e56aebac72f6896814998717ce2949108f3ebf7ba122370e4c959f7"></a>

## xcsh_bot_infrastructure — xcsh_bot_infrastructure / 08022d1cd3b3 / 2

Breadcrumbs:

- xcsh_bot_infrastructure

Manages Bot Infrastructure in F5 Distributed Cloud.

<a id="canonical-260351bb88ac3f72a36638f9a78fb119c881cdf799f5a075ba0af3e9185f1bb7"></a>

## Prerequisites — xcsh_bot_infrastructure / 08022d1cd3b3 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1741158fae9d3990ee560b630d665854a899e05d271307ec8d5a596ea9c24ee7"></a>

## Minimal configuration — xcsh_bot_infrastructure / 08022d1cd3b3 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotInfrastructure Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotInfrastructure by name
data "xcsh_bot_infrastructure" "example" {
  name      = "example-bot-infrastructure"
  namespace = "staging"
}

output "bot_infrastructure_id" {
  value = data.xcsh_bot_infrastructure.example.id
}
```

<a id="canonical-5aa05381451fcb10583be87dff2d6a5738cbce3010953bd6acea8f39eb599af2"></a>

## Root configuration — xcsh_bot_infrastructure / 08022d1cd3b3 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1a6b2d16b560a27adafe517d0ed77d084b1bcac90217caf02855df0b28b3cb89"></a>

## Next pages — xcsh_bot_infrastructure / 08022d1cd3b3 / 6

- [Property reference](../guides/data-sources--bot_infrastructure--reference--group-001.md#canonical-e3ce34ec019c1f41a1f0eae669da3664626b6c7209669ae24fa287835acfefee)
- [Examples](../guides/data-sources--bot_infrastructure--examples--group-001.md#canonical-f81d994c8dee00bb9e231d904fced88e7af9e8c8d882c731cb033b9b31919ca3)
