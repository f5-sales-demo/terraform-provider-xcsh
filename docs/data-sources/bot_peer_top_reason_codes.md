---
page_title: "xcsh_bot_peer_top_reason_codes landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_top_reason_codes landing."
---

# xcsh_bot_peer_top_reason_codes landing

<a id="canonical-07a07206937eae30b859c6d2292d9376c48db6e56deabe8cee1ccc1272a3593f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38a95e341090a97bd120099e442810b4fe7a45f7e097b1ace2de0be11b8ad1f4"></a>

## xcsh_bot_peer_top_reason_codes — xcsh_bot_peer_top_reason_codes / ac009aea9eaa / 2

Breadcrumbs:

- xcsh_bot_peer_top_reason_codes

Resource creation operation.

<a id="canonical-a27a1de813b8be708d802780adf6c140dce75a0ceaaefe865f8f4fcb31ee08fe"></a>

## Prerequisites — xcsh_bot_peer_top_reason_codes / ac009aea9eaa / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-c653057ab67fa9ae9f6f4272c4c210f983d3638bdd6a8c090bf9e968963835a8"></a>

## Minimal configuration — xcsh_bot_peer_top_reason_codes / ac009aea9eaa / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-7e21854e8d920bb40fff8d186ec700fcbad37bdfffb30896839ddd5bb5306da7"></a>

## Root configuration — xcsh_bot_peer_top_reason_codes / ac009aea9eaa / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-d76594cc9e1c2232aaf417cff93bb6afaf653a7129269270a9dfd62f17b6f802"></a>

## Next pages — xcsh_bot_peer_top_reason_codes / ac009aea9eaa / 6

- [Property reference](../guides/data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-1458b8c00d28b21f8f369757866bee112100ec9dea2320a099dc3c9ccabea02c)
- [Examples](../guides/data-sources--bot_peer_top_reason_codes--examples--group-001.md#canonical-38eccde791314d913c5228630d1e49363f56ad2c39bcfda06812987212794fe8)
