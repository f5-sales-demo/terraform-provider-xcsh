---
page_title: "xcsh_bot_allowlist_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_allowlist_policy landing."
---

# xcsh_bot_allowlist_policy landing

<a id="canonical-0a54609b088afff008aba1e10b878b96ab1aafdaea5949e64fa726729b39ac67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9ea3c511d39d427f2e08310c97adcf89b93a123b8d429de14a9513d57180234"></a>

## xcsh_bot_allowlist_policy — xcsh_bot_allowlist_policy / 2f428033d69a / 2

Breadcrumbs:

- xcsh_bot_allowlist_policy

Manages a Bot Allowlist Policy resource in F5 Distributed Cloud for get bot allowlist policy.
configuration. (read-only data source)

<a id="canonical-9b124f1da7101408f4fd25ae1a52e204d0c4b3c59172ac4a3618a2ae59abe270"></a>

## Prerequisites — xcsh_bot_allowlist_policy / 2f428033d69a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-b944498c23a9cbe2d5126842f2871793381ec8959493400379b6c6b3af24e0fb"></a>

## Minimal configuration — xcsh_bot_allowlist_policy / 2f428033d69a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-79be690454003466e09061d0aaa3677a25a79f6defbe672a4de823bdfb7e6810"></a>

## Root configuration — xcsh_bot_allowlist_policy / 2f428033d69a / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-b31c7ddeb754a37a598a8dcff07eb6d01cce67239d018e2042eb9cf5c49cdaa5"></a>

## Next pages — xcsh_bot_allowlist_policy / 2f428033d69a / 6

- [Property reference](../guides/data-sources--bot_allowlist_policy--reference--group-001.md#canonical-45b4719f99aa352d1d2ba9a7378ea2084fe7c7e9cc9db8dfe979a0e9fb00b69a)
- [Examples](../guides/data-sources--bot_allowlist_policy--examples--group-001.md#canonical-8f079968fe3841b46d07432afc6540b4864cc86ec358b6de62886703ad90c4b3)
