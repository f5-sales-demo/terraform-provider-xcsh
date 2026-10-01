---
page_title: "xcsh_device_intelligence_subscribe examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_subscribe examples."
---

# xcsh_device_intelligence_subscribe examples

<a id="canonical-22922a3a08a3d66db2080ecb870b0035d995256d7b533c74ba7b10ac90d24f96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a11b51c287eb34deff19423b7bd6d5d98aab519f8ca7b54560e27f7cb202c035"></a>

## Examples — Examples / 6f47f120f9ce / 2

Breadcrumbs:

- [xcsh_device_intelligence_subscribe](../actions/device_intelligence_subscribe.md#canonical-944ff0ac7f5f6c03cd1603eb46573f11db26913e12d02841ae3bb74567fdd738)
- Examples

<a id="canonical-065bb8389c849454806eaf5c3c55d315ed4aed9efe037b666c475cb5405ae558"></a>

## Complete configurations — Examples / 6f47f120f9ce / 3

- [Action](actions--device_intelligence_subscribe--examples--group-001.md#canonical-59cb0512c93b78c8fefb5157453450e0a96ca0737becc642e91b39a0c75555d4): valid configuration.

<a id="canonical-27f09e91a818824242047c59dd83b3d33109a5548b52df22ddd0b7e9c0af5c9c"></a>

## Next pages — Examples / 6f47f120f9ce / 4

- [Action](actions--device_intelligence_subscribe--examples--group-001.md#canonical-59cb0512c93b78c8fefb5157453450e0a96ca0737becc642e91b39a0c75555d4)
- [xcsh_device_intelligence_subscribe](../actions/device_intelligence_subscribe.md#canonical-944ff0ac7f5f6c03cd1603eb46573f11db26913e12d02841ae3bb74567fdd738)

<a id="canonical-59cb0512c93b78c8fefb5157453450e0a96ca0737becc642e91b39a0c75555d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e30cc4f77785fc90922508bfcb517f46defa619c532ff704b6da37a74adc37e0"></a>

## Action — Action / d69caab75796 / 2

Breadcrumbs:

- [xcsh_device_intelligence_subscribe](../actions/device_intelligence_subscribe.md#canonical-944ff0ac7f5f6c03cd1603eb46573f11db26913e12d02841ae3bb74567fdd738)
- [Examples](actions--device_intelligence_subscribe--examples--group-001.md#canonical-22922a3a08a3d66db2080ecb870b0035d995256d7b533c74ba7b10ac90d24f96)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_device_intelligence_subscribe/action.tf`; digest `sha256:f3174d51768147466c8e2454428347852090bb677186e716ca6cee429d892fbc`.

```terraform
# DeviceIntelligenceSubscribe Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_device_intelligence_subscribe" "example" {
  config {
  }
}
```

<a id="canonical-7bd795982df814b33533c90ab908d6f7bc33a097d15663384285b57ba94ef116"></a>

## Next pages — Action / d69caab75796 / 3

- [Examples](actions--device_intelligence_subscribe--examples--group-001.md#canonical-22922a3a08a3d66db2080ecb870b0035d995256d7b533c74ba7b10ac90d24f96)
- [xcsh_device_intelligence_subscribe](../actions/device_intelligence_subscribe.md#canonical-944ff0ac7f5f6c03cd1603eb46573f11db26913e12d02841ae3bb74567fdd738)
