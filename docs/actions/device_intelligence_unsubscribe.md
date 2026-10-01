---
page_title: "xcsh_device_intelligence_unsubscribe landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_unsubscribe landing."
---

# xcsh_device_intelligence_unsubscribe landing

<a id="canonical-d41a0df818a60a762cff58803d2fbe3e5165e9e4cd5c305de18ea9d4a44b3236"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d470f2d1a239346f2896491dc313cc43fae63589afc223e60ece207440f7771"></a>

## xcsh_device_intelligence_unsubscribe — xcsh_device_intelligence_unsubscribe / 44b1dde18a42 / 2

Breadcrumbs:

- xcsh_device_intelligence_unsubscribe

Resource creation operation.

<a id="canonical-cef3c21ed0f928c662e2fd7557f72d91e68eeb455afac0bd9ad127b784034a16"></a>

## Prerequisites — xcsh_device_intelligence_unsubscribe / 44b1dde18a42 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-79ad82c7f23f660f79d198f11d6a5a413a46f4b284b82cc32df28705d5cc7041"></a>

## Minimal configuration — xcsh_device_intelligence_unsubscribe / 44b1dde18a42 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-6a4b70e9cc95135a4a3f17477819065e3dfc57e5f0ddbce81ec2ed6217fa963a"></a>

## Root configuration — xcsh_device_intelligence_unsubscribe / 44b1dde18a42 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-ded4aef0e82dda79a7c92e5e436312bb42bdbb5ac91bea47efbb59cbe97262f5"></a>

## Next pages — xcsh_device_intelligence_unsubscribe / 44b1dde18a42 / 6

- [Property reference](../guides/actions--device_intelligence_unsubscribe--reference--group-001.md#canonical-c4ace8df5e1a91481f8e94bff0231020c0fd855159b38aaacd07a0e7a05a9a59)
- [Examples](../guides/actions--device_intelligence_unsubscribe--examples--group-001.md#canonical-fcb0a411e418156555786e6c3ecce2bd16e8fe899058e9881313dac3d2edf21e)
- [Lifecycle](../guides/actions--device_intelligence_unsubscribe--lifecycle--group-001.md#canonical-cf9ef0a81f17b62f8d3803dd0600994d375ddb7a14e939e78b381957b7c926af)
