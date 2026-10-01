---
page_title: "xcsh_bot_peer_top_reason_codes examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_top_reason_codes examples."
---

# xcsh_bot_peer_top_reason_codes examples

<a id="canonical-38eccde791314d913c5228630d1e49363f56ad2c39bcfda06812987212794fe8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-596eb0e19d5051f760eeaa3e6f0075ef515948e57dfade952d58cf14a1d9801e"></a>

## Examples — Examples / f940a67e67cc / 2

Breadcrumbs:

- [xcsh_bot_peer_top_reason_codes](../data-sources/bot_peer_top_reason_codes.md#canonical-07a07206937eae30b859c6d2292d9376c48db6e56deabe8cee1ccc1272a3593f)
- Examples

<a id="canonical-2d37df06a7835adfc95da78f6c13ef20ac8e3b07467d7d42771737a1b54ccbfc"></a>

## Complete configurations — Examples / f940a67e67cc / 3

- [Data source](data-sources--bot_peer_top_reason_codes--examples--group-001.md#canonical-ac7d5b3a622cff4ce09c62caf27cd0abd43c656f11ae57ce746425c58bfeb831): valid configuration.

<a id="canonical-d6640b30924243f231b8a37c87b0301c0095ddfb9867755739d7e3bf500afc01"></a>

## Next pages — Examples / f940a67e67cc / 4

- [Data source](data-sources--bot_peer_top_reason_codes--examples--group-001.md#canonical-ac7d5b3a622cff4ce09c62caf27cd0abd43c656f11ae57ce746425c58bfeb831)
- [xcsh_bot_peer_top_reason_codes](../data-sources/bot_peer_top_reason_codes.md#canonical-07a07206937eae30b859c6d2292d9376c48db6e56deabe8cee1ccc1272a3593f)

<a id="canonical-ac7d5b3a622cff4ce09c62caf27cd0abd43c656f11ae57ce746425c58bfeb831"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ce0f39ff2b062e8b0d86903fae547615c3d55862a619d97d67aa90665e61cf3"></a>

## Data source — Data source / be82ce046ce8 / 2

Breadcrumbs:

- [xcsh_bot_peer_top_reason_codes](../data-sources/bot_peer_top_reason_codes.md#canonical-07a07206937eae30b859c6d2292d9376c48db6e56deabe8cee1ccc1272a3593f)
- [Examples](data-sources--bot_peer_top_reason_codes--examples--group-001.md#canonical-38eccde791314d913c5228630d1e49363f56ad2c39bcfda06812987212794fe8)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_top_reason_codes/data-source.tf`; digest `sha256:088c7908202451a1670b0688576af84d4f3a2cb203efd005dc0803bb8dc3d516`.

```terraform
# BotPeerTopReasonCodes DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_top_reason_codes" "example" {
  namespace = "example-value"
}

output "bot_peer_top_reason_codes_result" {
  value = data.xcsh_bot_peer_top_reason_codes.example
}
```

<a id="canonical-f1edb926db12963d5faef90ab0c32332e3f33c0eefb42ecc6139f318aa790e98"></a>

## Next pages — Data source / be82ce046ce8 / 3

- [Examples](data-sources--bot_peer_top_reason_codes--examples--group-001.md#canonical-38eccde791314d913c5228630d1e49363f56ad2c39bcfda06812987212794fe8)
- [xcsh_bot_peer_top_reason_codes](../data-sources/bot_peer_top_reason_codes.md#canonical-07a07206937eae30b859c6d2292d9376c48db6e56deabe8cee1ccc1272a3593f)
