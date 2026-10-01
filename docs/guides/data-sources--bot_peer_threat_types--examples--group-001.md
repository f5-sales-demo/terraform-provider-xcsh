---
page_title: "xcsh_bot_peer_threat_types examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_threat_types examples."
---

# xcsh_bot_peer_threat_types examples

<a id="canonical-537b48ea58dc8f1705e50d5b7fa96b05d81cf486cc6fc4e35a410e40a432474d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9417b6e3d28013b0d8383f409cf78cf880c880ca9cd45d8c8efc2a7403969cdf"></a>

## Examples — Examples / 0d72afea012c / 2

Breadcrumbs:

- [xcsh_bot_peer_threat_types](../data-sources/bot_peer_threat_types.md#canonical-87108d47396ce895c0d9b659f1d45c2967514dee0cd6d925cfc6ed3f040bf641)
- Examples

<a id="canonical-35bc319b93024d5db9e7727e55611a10c085ca7472fd00569ab9c02357e50973"></a>

## Complete configurations — Examples / 0d72afea012c / 3

- [Data source](data-sources--bot_peer_threat_types--examples--group-001.md#canonical-f3c921687d39ac31f2f17dc9d8d98fa5142c3f68835f786485762e0bb333881a): valid configuration.

<a id="canonical-c278e7acc61f9ac6e1351b3ab579dd2076dde06451413097a156c7007ebc1b27"></a>

## Next pages — Examples / 0d72afea012c / 4

- [Data source](data-sources--bot_peer_threat_types--examples--group-001.md#canonical-f3c921687d39ac31f2f17dc9d8d98fa5142c3f68835f786485762e0bb333881a)
- [xcsh_bot_peer_threat_types](../data-sources/bot_peer_threat_types.md#canonical-87108d47396ce895c0d9b659f1d45c2967514dee0cd6d925cfc6ed3f040bf641)

<a id="canonical-f3c921687d39ac31f2f17dc9d8d98fa5142c3f68835f786485762e0bb333881a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc64f448c50138aa10f5ffeab07bc5faaccfb7dd91365e8535e6bcae1a4e79e8"></a>

## Data source — Data source / a2dcffdf1389 / 2

Breadcrumbs:

- [xcsh_bot_peer_threat_types](../data-sources/bot_peer_threat_types.md#canonical-87108d47396ce895c0d9b659f1d45c2967514dee0cd6d925cfc6ed3f040bf641)
- [Examples](data-sources--bot_peer_threat_types--examples--group-001.md#canonical-537b48ea58dc8f1705e50d5b7fa96b05d81cf486cc6fc4e35a410e40a432474d)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_threat_types/data-source.tf`; digest `sha256:340fefb3e7c9dd8e5383ef09986b56ee68d03ed73a4bcf9a8c9ec927ce63eafa`.

```terraform
# BotPeerThreatTypes DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_threat_types" "example" {
  namespace = "example-value"
}

output "bot_peer_threat_types_result" {
  value = data.xcsh_bot_peer_threat_types.example
}
```

<a id="canonical-9901dcff0caf2f2ebbc6e01ed0d3f7e901de0cfa304264dd4a0caea781693a25"></a>

## Next pages — Data source / a2dcffdf1389 / 3

- [Examples](data-sources--bot_peer_threat_types--examples--group-001.md#canonical-537b48ea58dc8f1705e50d5b7fa96b05d81cf486cc6fc4e35a410e40a432474d)
- [xcsh_bot_peer_threat_types](../data-sources/bot_peer_threat_types.md#canonical-87108d47396ce895c0d9b659f1d45c2967514dee0cd6d925cfc6ed3f040bf641)
