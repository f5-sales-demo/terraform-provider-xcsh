---
page_title: "xcsh_bot_endpoint_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_endpoint_policy examples."
---

# xcsh_bot_endpoint_policy examples

<a id="canonical-f3fc9b8121eec2ea862c71afcbe8dda2f435d39b3bb58aa7280c10002e0e5d34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ea857b1d7606c2c83003edb2fa0725045f6bfc3d52d4520d971200de3785484"></a>

## Examples — Examples / eebbbe11dd93 / 2

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-38154a6190bec8d1fc40fab4f7821bafd82f9fad21eccdc4504f325c1f4bafed)
- Examples

<a id="canonical-c4481076ac953e4882fa617dc7f52fa4fd2c0d4c161a99fe257429b898758ddc"></a>

## Complete configurations — Examples / eebbbe11dd93 / 3

- [Data source](data-sources--bot_endpoint_policy--examples--group-001.md#canonical-87f8a0c65be07eb385ddc517d4afd5272bc0c12110e9cd5d7994cc892ecd1c5e): valid configuration.

<a id="canonical-03e4390b85ed4246bba617b2b8e8418ed0ce39423185e5cc96a37817745aa141"></a>

## Next pages — Examples / eebbbe11dd93 / 4

- [Data source](data-sources--bot_endpoint_policy--examples--group-001.md#canonical-87f8a0c65be07eb385ddc517d4afd5272bc0c12110e9cd5d7994cc892ecd1c5e)
- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-38154a6190bec8d1fc40fab4f7821bafd82f9fad21eccdc4504f325c1f4bafed)

<a id="canonical-87f8a0c65be07eb385ddc517d4afd5272bc0c12110e9cd5d7994cc892ecd1c5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c18cee63fa84abda841b811d891fe05860058d1e9bb183cb77000f6b5da814ca"></a>

## Data source — Data source / ee10ba285f29 / 2

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-38154a6190bec8d1fc40fab4f7821bafd82f9fad21eccdc4504f325c1f4bafed)
- [Examples](data-sources--bot_endpoint_policy--examples--group-001.md#canonical-f3fc9b8121eec2ea862c71afcbe8dda2f435d39b3bb58aa7280c10002e0e5d34)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_endpoint_policy/data-source.tf`; digest `sha256:742dea18fab3d10fddee899cf0137a64f4bc6d260ffbdc975152403793a768eb`.

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

<a id="canonical-19a1f62437c534645728e70a2526d7b61edca0410eace7b7f46e59cb05edf9db"></a>

## Next pages — Data source / ee10ba285f29 / 3

- [Examples](data-sources--bot_endpoint_policy--examples--group-001.md#canonical-f3fc9b8121eec2ea862c71afcbe8dda2f435d39b3bb58aa7280c10002e0e5d34)
- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-38154a6190bec8d1fc40fab4f7821bafd82f9fad21eccdc4504f325c1f4bafed)
