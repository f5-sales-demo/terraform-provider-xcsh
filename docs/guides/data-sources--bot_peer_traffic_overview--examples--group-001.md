---
page_title: "xcsh_bot_peer_traffic_overview examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_traffic_overview examples."
---

# xcsh_bot_peer_traffic_overview examples

<a id="canonical-3a2834e1966934c911456ee9fe04e4648fad8cd2a00dcf42484439d3b4adfcf3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3269fb3ce2371d17e5f58735e5a8925740366cec0256ac2491a4e131c8512ef"></a>

## Examples — Examples / fc25020cff20 / 2

Breadcrumbs:

- [xcsh_bot_peer_traffic_overview](../data-sources/bot_peer_traffic_overview.md#canonical-5b05326e4d591c752fdaa94722e4badaea615d250171e751c47a2266294991d7)
- Examples

<a id="canonical-6458dd9d2093be2e8e8838d1fa67135f1b8ddde78059e275a92a15e620a84985"></a>

## Complete configurations — Examples / fc25020cff20 / 3

- [Data source](data-sources--bot_peer_traffic_overview--examples--group-001.md#canonical-891ff0c78a7ad4c3942d90d73dac2163dd006c878aa78094a0e9a6b8163c136d): valid configuration.

<a id="canonical-d7aa5777a5918b51720e1bd3dd70c430bef5099533cb89110790428b070b4825"></a>

## Next pages — Examples / fc25020cff20 / 4

- [Data source](data-sources--bot_peer_traffic_overview--examples--group-001.md#canonical-891ff0c78a7ad4c3942d90d73dac2163dd006c878aa78094a0e9a6b8163c136d)
- [xcsh_bot_peer_traffic_overview](../data-sources/bot_peer_traffic_overview.md#canonical-5b05326e4d591c752fdaa94722e4badaea615d250171e751c47a2266294991d7)

<a id="canonical-891ff0c78a7ad4c3942d90d73dac2163dd006c878aa78094a0e9a6b8163c136d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8b9742dad18ff79a77c91b050eca226f9cd7b0c7ff5dd096aa2ef1f0bbeb67c"></a>

## Data source — Data source / 36da8473f95b / 2

Breadcrumbs:

- [xcsh_bot_peer_traffic_overview](../data-sources/bot_peer_traffic_overview.md#canonical-5b05326e4d591c752fdaa94722e4badaea615d250171e751c47a2266294991d7)
- [Examples](data-sources--bot_peer_traffic_overview--examples--group-001.md#canonical-3a2834e1966934c911456ee9fe04e4648fad8cd2a00dcf42484439d3b4adfcf3)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_traffic_overview/data-source.tf`; digest `sha256:02ae4e984486512606c527da0d434568ac636c951e2a90ff31f351d9dd36be11`.

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

<a id="canonical-9f729b85437e37daaaec76692594c1120bf63f5254947c0041885e97f92b2361"></a>

## Next pages — Data source / 36da8473f95b / 3

- [Examples](data-sources--bot_peer_traffic_overview--examples--group-001.md#canonical-3a2834e1966934c911456ee9fe04e4648fad8cd2a00dcf42484439d3b4adfcf3)
- [xcsh_bot_peer_traffic_overview](../data-sources/bot_peer_traffic_overview.md#canonical-5b05326e4d591c752fdaa94722e4badaea615d250171e751c47a2266294991d7)
