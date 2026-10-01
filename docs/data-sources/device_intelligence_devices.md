---
page_title: "xcsh_device_intelligence_devices landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_devices landing."
---

# xcsh_device_intelligence_devices landing

<a id="canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6c4d946fac43b80207e80af9ef9bf47a758568c762ae16352a27cd2b60b3d61"></a>

## xcsh_device_intelligence_devices — xcsh_device_intelligence_devices / d22e5138e991 / 2

Breadcrumbs:

- xcsh_device_intelligence_devices

Resource creation operation.

<a id="canonical-c9cdca98965b72e695b1236ee387c0aed0d5dcb3f7d800c24fcca5a2841d2c34"></a>

## Prerequisites — xcsh_device_intelligence_devices / d22e5138e991 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-660dce212786df4c00134ff0de33ae73d4230b0812f2767d3f2d708166ef92a1"></a>

## Minimal configuration — xcsh_device_intelligence_devices / d22e5138e991 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceDevices DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_devices" "example" {
  namespace = "example-value"
}

output "device_intelligence_devices_result" {
  value = data.xcsh_device_intelligence_devices.example
}
```

<a id="canonical-cce0a46f7ba4fe95421dbbf2282a064842aed4aaa5d1cc8ec9f4897308c13ef0"></a>

## Root configuration — xcsh_device_intelligence_devices / d22e5138e991 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-686cb426dc3805aaa372abbd9b024bc4f51037a245b05b8e70b0f003fe904bc5"></a>

## Next pages — xcsh_device_intelligence_devices / d22e5138e991 / 6

- [Property reference](../guides/data-sources--device_intelligence_devices--reference--group-001.md#canonical-44a62deff463fb871382d9c68671f83b580d2862aa9293e18cb445f06d1be873)
- [Examples](../guides/data-sources--device_intelligence_devices--examples--group-001.md#canonical-e29fe453efdf882c0fe3e41b13b48c36e410f492a713d71dc15c59df971c1c92)
