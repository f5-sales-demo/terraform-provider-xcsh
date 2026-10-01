---
page_title: "xcsh_bot_peer_traffic_overview landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_traffic_overview landing."
---

# xcsh_bot_peer_traffic_overview landing

<a id="canonical-5b05326e4d591c752fdaa94722e4badaea615d250171e751c47a2266294991d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8eb701007e9aae83cc6d1d59152398c8cbbbd27ea3a5df8565db3062fa10b34f"></a>

## xcsh_bot_peer_traffic_overview — xcsh_bot_peer_traffic_overview / 53ec1b08503f / 2

Breadcrumbs:

- xcsh_bot_peer_traffic_overview

Resource creation operation.

<a id="canonical-922511d3b7474a8f8d3d6eb43d3b248d3c4723109365e26f98bc6327a2dfbde8"></a>

## Prerequisites — xcsh_bot_peer_traffic_overview / 53ec1b08503f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-310d5001ffd8e29cb3977e35834ee16f5690516341333f228e54de9dca982fc9"></a>

## Minimal configuration — xcsh_bot_peer_traffic_overview / 53ec1b08503f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotPeerTrafficOverview DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_traffic_overview" "example" {
  namespace = "example-value"
}

output "bot_peer_traffic_overview_result" {
  value = data.xcsh_bot_peer_traffic_overview.example
}
```

<a id="canonical-d69d7e495018be8ec12d6e0f5943b29c8d9b54ef2bf1e52423b041cc9770e4c0"></a>

## Root configuration — xcsh_bot_peer_traffic_overview / 53ec1b08503f / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-eb067f042cb61796616151e46150c83c7cf7853dd84ed70eea203f0514bbca50"></a>

## Next pages — xcsh_bot_peer_traffic_overview / 53ec1b08503f / 6

- [Property reference](../guides/data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-53367a814368233d59c0d60053b9993843031752de8473de41d0df93b435773a)
- [Examples](../guides/data-sources--bot_peer_traffic_overview--examples--group-001.md#canonical-3a2834e1966934c911456ee9fe04e4648fad8cd2a00dcf42484439d3b4adfcf3)
