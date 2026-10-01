---
page_title: "xcsh_bot_peer_status landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_status landing."
---

# xcsh_bot_peer_status landing

<a id="canonical-67ec1acf92d1d20d9c08be3fef0b7bd8125ef46625bfa7cae7978e9ffcf45de4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14c91ad8aa3a4688e754d68a2f135d29fc70ddc5109786947e6da005e33b830e"></a>

## xcsh_bot_peer_status — xcsh_bot_peer_status / f30e1dc20f24 / 2

Breadcrumbs:

- xcsh_bot_peer_status

Resource creation operation.

<a id="canonical-90c43b8156de288a5f16830816dd87890bc6a316ebf50daf7423d7d0fafd5da3"></a>

## Prerequisites — xcsh_bot_peer_status / f30e1dc20f24 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-fb53f1cbc4b82ac4e67a62a5cc48413db12f261017143c7007be49888b3709f8"></a>

## Minimal configuration — xcsh_bot_peer_status / f30e1dc20f24 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-3d53bbddb244891cd2d62d5cac5131fd061546beeca339359e636ca7d5e6fa32"></a>

## Root configuration — xcsh_bot_peer_status / f30e1dc20f24 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-4ce4ae00c0ab0a7e75fccfb958fe32d3796dd1e87cb420977b992b3c4488133d"></a>

## Next pages — xcsh_bot_peer_status / f30e1dc20f24 / 6

- [Property reference](../guides/data-sources--bot_peer_status--reference--group-001.md#canonical-af5bdebbe653a239e2a54201e0a334a1c2e085d044d39db19e057bf1f650b4a5)
- [Examples](../guides/data-sources--bot_peer_status--examples--group-001.md#canonical-47ce9ed4715ab48aef0cbccef3678f1d5336b8af21bd233e2b03e7094538d3e7)
