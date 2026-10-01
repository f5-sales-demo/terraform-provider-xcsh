---
page_title: "xcsh_bot_peer_status examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_status examples."
---

# xcsh_bot_peer_status examples

<a id="canonical-47ce9ed4715ab48aef0cbccef3678f1d5336b8af21bd233e2b03e7094538d3e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e9daf10ee8510a8e768d6558354148ff6e4303f0e6678a8839185268c94748b"></a>

## Examples — Examples / c04b5836fdc7 / 2

Breadcrumbs:

- [xcsh_bot_peer_status](../data-sources/bot_peer_status.md#canonical-67ec1acf92d1d20d9c08be3fef0b7bd8125ef46625bfa7cae7978e9ffcf45de4)
- Examples

<a id="canonical-c3c9644426dd8d43b642a7cadaeded7ed35b824649149f03e83ae628d208a3b7"></a>

## Complete configurations — Examples / c04b5836fdc7 / 3

- [Data source](data-sources--bot_peer_status--examples--group-001.md#canonical-290fee387751fb93a50ea3713d46d393cc45202130600e6cb5044f6f691ef783): valid configuration.

<a id="canonical-3765cfbad6436cc2bbfb3cc935f3e87d615cda5d49945e61297386d86bd0872c"></a>

## Next pages — Examples / c04b5836fdc7 / 4

- [Data source](data-sources--bot_peer_status--examples--group-001.md#canonical-290fee387751fb93a50ea3713d46d393cc45202130600e6cb5044f6f691ef783)
- [xcsh_bot_peer_status](../data-sources/bot_peer_status.md#canonical-67ec1acf92d1d20d9c08be3fef0b7bd8125ef46625bfa7cae7978e9ffcf45de4)

<a id="canonical-290fee387751fb93a50ea3713d46d393cc45202130600e6cb5044f6f691ef783"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-326803db04f76bfd207360e18363ce47ef87b8465038859a29f65ee6d4c07333"></a>

## Data source — Data source / 7978d48a1735 / 2

Breadcrumbs:

- [xcsh_bot_peer_status](../data-sources/bot_peer_status.md#canonical-67ec1acf92d1d20d9c08be3fef0b7bd8125ef46625bfa7cae7978e9ffcf45de4)
- [Examples](data-sources--bot_peer_status--examples--group-001.md#canonical-47ce9ed4715ab48aef0cbccef3678f1d5336b8af21bd233e2b03e7094538d3e7)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_status/data-source.tf`; digest `sha256:97b1bb9bbcf7e35e2e9068fc82fcedf34184ffea0125e8cf3c1304beb4255fcc`.

```terraform
# BotPeerStatus DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_status" "example" {
  namespace = "example-value"
}

output "bot_peer_status_result" {
  value = data.xcsh_bot_peer_status.example
}
```

<a id="canonical-70241b5e8f1244e0f7efc347067de48ae94f329b55d01cab217c60b4d0459c06"></a>

## Next pages — Data source / 7978d48a1735 / 3

- [Examples](data-sources--bot_peer_status--examples--group-001.md#canonical-47ce9ed4715ab48aef0cbccef3678f1d5336b8af21bd233e2b03e7094538d3e7)
- [xcsh_bot_peer_status](../data-sources/bot_peer_status.md#canonical-67ec1acf92d1d20d9c08be3fef0b7bd8125ef46625bfa7cae7978e9ffcf45de4)
