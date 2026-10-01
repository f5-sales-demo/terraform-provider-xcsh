---
page_title: "xcsh_bot_peer_top_good_bots landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_top_good_bots landing."
---

# xcsh_bot_peer_top_good_bots landing

<a id="canonical-cc077f9c205b2a2bad9388a3f1f0bfb9445ec52b5c9b1a27c681c552089ee8f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-564ba4fe0147816d8b978806fd624e8edcbd7cdad5a9cae5a76a469059d82c08"></a>

## xcsh_bot_peer_top_good_bots — xcsh_bot_peer_top_good_bots / 673cc0a461e0 / 2

Breadcrumbs:

- xcsh_bot_peer_top_good_bots

Bot detection and defense configuration.

<a id="canonical-26683f38e126f375a82f6c71d24e4fb7bda571a9675935758ad475ebd1fbb321"></a>

## Prerequisites — xcsh_bot_peer_top_good_bots / 673cc0a461e0 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-8e62b3e7d37546533e1580e012b17394f8c3aee83eb56d19e2f02984896868d5"></a>

## Minimal configuration — xcsh_bot_peer_top_good_bots / 673cc0a461e0 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-9d0192abae46690dbed376f46bd4ecb08067886d61c89c14c537a1bf56ec8486"></a>

## Root configuration — xcsh_bot_peer_top_good_bots / 673cc0a461e0 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-aa755f81e9b91b4c924a12dd97a29993175e135905e2bd6e30c8f41b5591ca64"></a>

## Next pages — xcsh_bot_peer_top_good_bots / 673cc0a461e0 / 6

- [Property reference](../guides/data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-6ac86678f40f40275bf9c57605bc9b4e11cf7b3003d83b65b2d5f4fdfbda158e)
- [Examples](../guides/data-sources--bot_peer_top_good_bots--examples--group-001.md#canonical-db31b0223ac18b2e200dbc5b695a237d253b717574490767f7c9603adca1fdd9)
