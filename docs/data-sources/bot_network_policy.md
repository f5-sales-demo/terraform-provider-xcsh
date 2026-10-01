---
page_title: "xcsh_bot_network_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_network_policy landing."
---

# xcsh_bot_network_policy landing

<a id="canonical-1db4aff0b95edbb987adc1f675e48bc99801201bf045f697bf744647c8e90762"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b10aa8452d199962284840e7aa639891fc5f786dfb3b4abb20d823088bdbf387"></a>

## xcsh_bot_network_policy — xcsh_bot_network_policy / 66105fac6288 / 2

Breadcrumbs:

- xcsh_bot_network_policy

Manages a Bot Network Policy resource in F5 Distributed Cloud for get bot network policy.
configuration. (read-only data source)

<a id="canonical-094a7c765fc9d618747429c3169494d76179029364938efc50ecbf6ff8699bb5"></a>

## Prerequisites — xcsh_bot_network_policy / 66105fac6288 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-e8770ee286835d68b8f4e4f1db290f2be0e581f7ce2bdb056f1e435afda46418"></a>

## Minimal configuration — xcsh_bot_network_policy / 66105fac6288 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotNetworkPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotNetworkPolicy by name
data "xcsh_bot_network_policy" "example" {
  name      = "example-bot-network-policy"
  namespace = "staging"
}

output "bot_network_policy_id" {
  value = data.xcsh_bot_network_policy.example.id
}
```

<a id="canonical-b428ce77f118a1926acc4d27136819b76f971848965c46bb06b9d687b458a259"></a>

## Root configuration — xcsh_bot_network_policy / 66105fac6288 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-5d0d00cca3de54e360935ec326a96c2cb45d3773168c87980b8f048e99ef843c"></a>

## Next pages — xcsh_bot_network_policy / 66105fac6288 / 6

- [Property reference](../guides/data-sources--bot_network_policy--reference--group-001.md#canonical-40a1ba93758d581a74bd7ac4b3bf87ac8b1586a0088bdc96a1f1e0fa5b76916c)
- [Examples](../guides/data-sources--bot_network_policy--examples--group-001.md#canonical-022ba32b36308c5407c9dc966cfe8e3dcde2eef6d9858ccd4c90ceb17cfa7059)
