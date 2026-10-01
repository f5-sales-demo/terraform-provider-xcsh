---
page_title: "xcsh_bot_endpoint_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_endpoint_policy landing."
---

# xcsh_bot_endpoint_policy landing

<a id="canonical-38154a6190bec8d1fc40fab4f7821bafd82f9fad21eccdc4504f325c1f4bafed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a081323c1b57b750d178d4f4cc2e5abd89319e2f16275c10b070eb59048f7c1c"></a>

## xcsh_bot_endpoint_policy — xcsh_bot_endpoint_policy / 82e9bb8a1208 / 2

Breadcrumbs:

- xcsh_bot_endpoint_policy

Manages a Bot Endpoint Policy resource in F5 Distributed Cloud for get bot endpoint policy.
configuration. (read-only data source)

<a id="canonical-00f0049ab1d5737fe492f1a94e175f168af93e46fdae2adf01e224aec73d9ada"></a>

## Prerequisites — xcsh_bot_endpoint_policy / 82e9bb8a1208 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-be5031fd931c53cf48b20f33590eb0a9b6e95b52a2bb229b8f0c362a36298210"></a>

## Minimal configuration — xcsh_bot_endpoint_policy / 82e9bb8a1208 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotEndpointPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotEndpointPolicy by name
data "xcsh_bot_endpoint_policy" "example" {
  name      = "example-bot-endpoint-policy"
  namespace = "staging"
}

output "bot_endpoint_policy_id" {
  value = data.xcsh_bot_endpoint_policy.example.id
}
```

<a id="canonical-ea3f8cecb958876b9f7ba4e18ae0f8fa9b586199316cd0153dfaae14b3f5045c"></a>

## Root configuration — xcsh_bot_endpoint_policy / 82e9bb8a1208 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-53862e96d2e955c82f93d661f1dcbca3fcbb4b5ab52ee11884e6c92b8309d74d"></a>

## Next pages — xcsh_bot_endpoint_policy / 82e9bb8a1208 / 6

- [Property reference](../guides/data-sources--bot_endpoint_policy--reference--group-001.md#canonical-a19d06c6acb78fe89eceb8f207e8e8348e5eecfbb5171261bb11f2427e65d205)
- [Examples](../guides/data-sources--bot_endpoint_policy--examples--group-001.md#canonical-f3fc9b8121eec2ea862c71afcbe8dda2f435d39b3bb58aa7280c10002e0e5d34)
