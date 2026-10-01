---
page_title: "xcsh_bot_peer_threat_types landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_threat_types landing."
---

# xcsh_bot_peer_threat_types landing

<a id="canonical-87108d47396ce895c0d9b659f1d45c2967514dee0cd6d925cfc6ed3f040bf641"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d2822af9c8ed756398de666c82495ab54a7952294e34110e5560d1433992135"></a>

## xcsh_bot_peer_threat_types — xcsh_bot_peer_threat_types / 52f0d239e3f6 / 2

Breadcrumbs:

- xcsh_bot_peer_threat_types

Resource creation operation.

<a id="canonical-b26d9d8ab1067cb53966be72b5e29f10edb873d5bdfa69d9f9eb0079ee4b5948"></a>

## Prerequisites — xcsh_bot_peer_threat_types / 52f0d239e3f6 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-e0e9a27e28f37df1b7732cfae314a6c0cf75d843c6fe5063f4c846ebc833a1a6"></a>

## Minimal configuration — xcsh_bot_peer_threat_types / 52f0d239e3f6 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-dc13780b81bae8362dcdad4a582bbbca65ac26be06d26a260f7b8bde8ae6efcb"></a>

## Root configuration — xcsh_bot_peer_threat_types / 52f0d239e3f6 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-13e46c2e3578c53a525fbe28b91886587fc5c4f9d7b069328b5913ae6e165bc3"></a>

## Next pages — xcsh_bot_peer_threat_types / 52f0d239e3f6 / 6

- [Property reference](../guides/data-sources--bot_peer_threat_types--reference--group-001.md#canonical-fe41abbeffdeb22d2d0f00d523f6ba4a5c0487f7ff396036ada01cc5a9aac87e)
- [Examples](../guides/data-sources--bot_peer_threat_types--examples--group-001.md#canonical-537b48ea58dc8f1705e50d5b7fa96b05d81cf486cc6fc4e35a410e40a432474d)
