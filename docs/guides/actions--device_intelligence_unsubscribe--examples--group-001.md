---
page_title: "xcsh_device_intelligence_unsubscribe examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_unsubscribe examples."
---

# xcsh_device_intelligence_unsubscribe examples

<a id="canonical-fcb0a411e418156555786e6c3ecce2bd16e8fe899058e9881313dac3d2edf21e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8021758923e20d59cc937654dc9305dfd3e2cdaa9a8b870f60524fb89c544e35"></a>

## Examples — Examples / a49a4c1531f1 / 2

Breadcrumbs:

- [xcsh_device_intelligence_unsubscribe](../actions/device_intelligence_unsubscribe.md#canonical-d41a0df818a60a762cff58803d2fbe3e5165e9e4cd5c305de18ea9d4a44b3236)
- Examples

<a id="canonical-b7c0734f2050bd6b1a953e630d4930c800f767b14bda37967de247d3b0faf634"></a>

## Complete configurations — Examples / a49a4c1531f1 / 3

- [Action](actions--device_intelligence_unsubscribe--examples--group-001.md#canonical-f5543964db99ad02eee7010debf25b6161f0368edf354e7b70aabec0e1e7fe26): valid configuration.

<a id="canonical-23943fae91f6e71e0b4402f4ca5fa2fe20c433d4015fa4e36710161190137a73"></a>

## Next pages — Examples / a49a4c1531f1 / 4

- [Action](actions--device_intelligence_unsubscribe--examples--group-001.md#canonical-f5543964db99ad02eee7010debf25b6161f0368edf354e7b70aabec0e1e7fe26)
- [xcsh_device_intelligence_unsubscribe](../actions/device_intelligence_unsubscribe.md#canonical-d41a0df818a60a762cff58803d2fbe3e5165e9e4cd5c305de18ea9d4a44b3236)

<a id="canonical-f5543964db99ad02eee7010debf25b6161f0368edf354e7b70aabec0e1e7fe26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1957256026b0548e44b970614528d59af37b49ee4fd9f10b223477f09ce61972"></a>

## Action — Action / ba6a1666ed06 / 2

Breadcrumbs:

- [xcsh_device_intelligence_unsubscribe](../actions/device_intelligence_unsubscribe.md#canonical-d41a0df818a60a762cff58803d2fbe3e5165e9e4cd5c305de18ea9d4a44b3236)
- [Examples](actions--device_intelligence_unsubscribe--examples--group-001.md#canonical-fcb0a411e418156555786e6c3ecce2bd16e8fe899058e9881313dac3d2edf21e)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_device_intelligence_unsubscribe/action.tf`; digest `sha256:391d15fe0fb026a460e8e78fbc03f16af6138839116ddca4fd7a4862c6cad9eb`.

```terraform
# DeviceIntelligenceUnsubscribe Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_device_intelligence_unsubscribe" "example" {
  config {
  }
}
```

<a id="canonical-757d71c4d4c9551a2ea996dc1254cfb44716b9d3aea3a8cc94310f53195201ed"></a>

## Next pages — Action / ba6a1666ed06 / 3

- [Examples](actions--device_intelligence_unsubscribe--examples--group-001.md#canonical-fcb0a411e418156555786e6c3ecce2bd16e8fe899058e9881313dac3d2edf21e)
- [xcsh_device_intelligence_unsubscribe](../actions/device_intelligence_unsubscribe.md#canonical-d41a0df818a60a762cff58803d2fbe3e5165e9e4cd5c305de18ea9d4a44b3236)
