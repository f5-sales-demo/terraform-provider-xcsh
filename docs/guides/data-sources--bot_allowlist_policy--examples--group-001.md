---
page_title: "xcsh_bot_allowlist_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_allowlist_policy examples."
---

# xcsh_bot_allowlist_policy examples

<a id="canonical-8f079968fe3841b46d07432afc6540b4864cc86ec358b6de62886703ad90c4b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f309745c6d0c7f1fca0f512ee60d26f716029930bd6cc22ba6c3b3d06e261daa"></a>

## Examples — Examples / 7d63419f7222 / 2

Breadcrumbs:

- [xcsh_bot_allowlist_policy](../data-sources/bot_allowlist_policy.md#canonical-0a54609b088afff008aba1e10b878b96ab1aafdaea5949e64fa726729b39ac67)
- Examples

<a id="canonical-85f57e4c7d9370ff6d56fe0a4474f620dba590ce48803c149b785b576dbda5ac"></a>

## Complete configurations — Examples / 7d63419f7222 / 3

- [Data source](data-sources--bot_allowlist_policy--examples--group-001.md#canonical-2805cf7cd45e9244dd58ac37ae66be6ad9a54d506b44a9d1892dcfde52e36fac): valid configuration.

<a id="canonical-1a636ddfe8ce4f5843fa0560df86666ee07c8f2192df07ec1ed61a7bce513eca"></a>

## Next pages — Examples / 7d63419f7222 / 4

- [Data source](data-sources--bot_allowlist_policy--examples--group-001.md#canonical-2805cf7cd45e9244dd58ac37ae66be6ad9a54d506b44a9d1892dcfde52e36fac)
- [xcsh_bot_allowlist_policy](../data-sources/bot_allowlist_policy.md#canonical-0a54609b088afff008aba1e10b878b96ab1aafdaea5949e64fa726729b39ac67)

<a id="canonical-2805cf7cd45e9244dd58ac37ae66be6ad9a54d506b44a9d1892dcfde52e36fac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ac401ca94f405598bdb0f80aa15d452b2ee9fa076ac68a194d1459fd4885fca"></a>

## Data source — Data source / 8bbd94471ce6 / 2

Breadcrumbs:

- [xcsh_bot_allowlist_policy](../data-sources/bot_allowlist_policy.md#canonical-0a54609b088afff008aba1e10b878b96ab1aafdaea5949e64fa726729b39ac67)
- [Examples](data-sources--bot_allowlist_policy--examples--group-001.md#canonical-8f079968fe3841b46d07432afc6540b4864cc86ec358b6de62886703ad90c4b3)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_allowlist_policy/data-source.tf`; digest `sha256:e79cb351210b30d9b4d2bf6ff1c7b6c44610a7fd5a55e86ba638364a3bd45c0b`.

```terraform
# BotAllowlistPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotAllowlistPolicy by name
data "xcsh_bot_allowlist_policy" "example" {
  name      = "example-bot-allowlist-policy"
  namespace = "staging"
}

output "bot_allowlist_policy_id" {
  value = data.xcsh_bot_allowlist_policy.example.id
}
```

<a id="canonical-8c5de865194d8d51758c29073e8c0530f49cb4ab55e8f9c26c43755c1b68706f"></a>

## Next pages — Data source / 8bbd94471ce6 / 3

- [Examples](data-sources--bot_allowlist_policy--examples--group-001.md#canonical-8f079968fe3841b46d07432afc6540b4864cc86ec358b6de62886703ad90c4b3)
- [xcsh_bot_allowlist_policy](../data-sources/bot_allowlist_policy.md#canonical-0a54609b088afff008aba1e10b878b96ab1aafdaea5949e64fa726729b39ac67)
