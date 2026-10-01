---
page_title: "xcsh_device_intelligence_subscribe landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_subscribe landing."
---

# xcsh_device_intelligence_subscribe landing

<a id="canonical-944ff0ac7f5f6c03cd1603eb46573f11db26913e12d02841ae3bb74567fdd738"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae3667b0901796b0ee08f4fc280930527e908971f644c3b66425b77912b3b6b3"></a>

## xcsh_device_intelligence_subscribe — xcsh_device_intelligence_subscribe / b813fbdac41a / 2

Breadcrumbs:

- xcsh_device_intelligence_subscribe

Resource creation operation.

<a id="canonical-87595dae3e6149ba3dddd28215b33e57b5fa804c4430fb6b21a6b940f775038a"></a>

## Prerequisites — xcsh_device_intelligence_subscribe / b813fbdac41a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-d72fce567f4f8527f00b450bef0bdeda478844886de883a4b53623ec1d6a95ca"></a>

## Minimal configuration — xcsh_device_intelligence_subscribe / b813fbdac41a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-2d2b93f533e46883fcd6e7b0d827d64348dcc13882f602bec328f26daaea84bc"></a>

## Root configuration — xcsh_device_intelligence_subscribe / b813fbdac41a / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-90be7a19db24ba3bf1a7ec23e4758e4c2b8b2005e8af745b85501ea7aa5b8e11"></a>

## Next pages — xcsh_device_intelligence_subscribe / b813fbdac41a / 6

- [Property reference](../guides/actions--device_intelligence_subscribe--reference--group-001.md#canonical-38e3b274ec38ed412333731177cdb118519fbe0caad9627a3cb4ff17db35e79c)
- [Examples](../guides/actions--device_intelligence_subscribe--examples--group-001.md#canonical-22922a3a08a3d66db2080ecb870b0035d995256d7b533c74ba7b10ac90d24f96)
- [Lifecycle](../guides/actions--device_intelligence_subscribe--lifecycle--group-001.md#canonical-f46fe04986b83f463c8fc3d1ddc1a00be4392ed6b1f94e666e29aac75fe313cd)
