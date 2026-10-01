---
page_title: "xcsh_bot_peer_top_good_bots examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_top_good_bots examples."
---

# xcsh_bot_peer_top_good_bots examples

<a id="canonical-db31b0223ac18b2e200dbc5b695a237d253b717574490767f7c9603adca1fdd9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75df3b538b83bf2bd98075df7da17270b2538755f2f08b8872cc1a80098f28a0"></a>

## Examples — Examples / ba32fd8b7def / 2

Breadcrumbs:

- [xcsh_bot_peer_top_good_bots](../data-sources/bot_peer_top_good_bots.md#canonical-cc077f9c205b2a2bad9388a3f1f0bfb9445ec52b5c9b1a27c681c552089ee8f2)
- Examples

<a id="canonical-582af2efc3210418e050b0c0a44a338ece473f6297e2a271cfbe35816c9e76a0"></a>

## Complete configurations — Examples / ba32fd8b7def / 3

- [Data source](data-sources--bot_peer_top_good_bots--examples--group-001.md#canonical-46ed82e9617dfcf0615c07d798f85aa8de2e7a66a2c5c355b0bb88290756f7cd): valid configuration.

<a id="canonical-0955cc27374db876ef294644a43d1d782adc61513a7ddc69d197647e92c8ee0f"></a>

## Next pages — Examples / ba32fd8b7def / 4

- [Data source](data-sources--bot_peer_top_good_bots--examples--group-001.md#canonical-46ed82e9617dfcf0615c07d798f85aa8de2e7a66a2c5c355b0bb88290756f7cd)
- [xcsh_bot_peer_top_good_bots](../data-sources/bot_peer_top_good_bots.md#canonical-cc077f9c205b2a2bad9388a3f1f0bfb9445ec52b5c9b1a27c681c552089ee8f2)

<a id="canonical-46ed82e9617dfcf0615c07d798f85aa8de2e7a66a2c5c355b0bb88290756f7cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a98b9c2e5f73a3e91ac99f6ded174c21670a83bd3d7bc058e4b66f680b6ef45"></a>

## Data source — Data source / f5581fa02d14 / 2

Breadcrumbs:

- [xcsh_bot_peer_top_good_bots](../data-sources/bot_peer_top_good_bots.md#canonical-cc077f9c205b2a2bad9388a3f1f0bfb9445ec52b5c9b1a27c681c552089ee8f2)
- [Examples](data-sources--bot_peer_top_good_bots--examples--group-001.md#canonical-db31b0223ac18b2e200dbc5b695a237d253b717574490767f7c9603adca1fdd9)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_top_good_bots/data-source.tf`; digest `sha256:cbd5508dce9276d9af60eef5a236b141ce290117e051aad973c941688a149920`.

```terraform
# BotPeerTopGoodBots DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_top_good_bots" "example" {
  namespace = "example-value"
}

output "bot_peer_top_good_bots_result" {
  value = data.xcsh_bot_peer_top_good_bots.example
}
```

<a id="canonical-e73742fba145e50de3a643354e6c1960d7611819969fff10544cbf43996d7ff5"></a>

## Next pages — Data source / f5581fa02d14 / 3

- [Examples](data-sources--bot_peer_top_good_bots--examples--group-001.md#canonical-db31b0223ac18b2e200dbc5b695a237d253b717574490767f7c9603adca1fdd9)
- [xcsh_bot_peer_top_good_bots](../data-sources/bot_peer_top_good_bots.md#canonical-cc077f9c205b2a2bad9388a3f1f0bfb9445ec52b5c9b1a27c681c552089ee8f2)
