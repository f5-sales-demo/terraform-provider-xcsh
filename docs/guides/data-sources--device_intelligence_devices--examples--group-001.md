---
page_title: "xcsh_device_intelligence_devices examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_devices examples."
---

# xcsh_device_intelligence_devices examples

<a id="canonical-e29fe453efdf882c0fe3e41b13b48c36e410f492a713d71dc15c59df971c1c92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2af2edd0d2ded59e49ccc22332dd4a2fdfc1cc79641f88b5fd51f903011250ec"></a>

## Examples — Examples / 58d22a918a1e / 2

Breadcrumbs:

- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)
- Examples

<a id="canonical-c28ee7c9cbb83c4e4bd0dc97e820c244edb3c3218e1e533be3c506eea13ef6b4"></a>

## Complete configurations — Examples / 58d22a918a1e / 3

- [Data source](data-sources--device_intelligence_devices--examples--group-001.md#canonical-bba07230bdcac5cf9f95e5ce131f8b3e284cd452215a5fed4656caec64e00f2f): valid configuration.

<a id="canonical-93f1939829929f040bd64c3e32add0ad21757b1519839c30e00dde7a5d2c0874"></a>

## Next pages — Examples / 58d22a918a1e / 4

- [Data source](data-sources--device_intelligence_devices--examples--group-001.md#canonical-bba07230bdcac5cf9f95e5ce131f8b3e284cd452215a5fed4656caec64e00f2f)
- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)

<a id="canonical-bba07230bdcac5cf9f95e5ce131f8b3e284cd452215a5fed4656caec64e00f2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c46639778cd8297cd8a1ad87ddb42291437537df9680465595aebccf3158aeb"></a>

## Data source — Data source / f5ed97bd2e8d / 2

Breadcrumbs:

- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)
- [Examples](data-sources--device_intelligence_devices--examples--group-001.md#canonical-e29fe453efdf882c0fe3e41b13b48c36e410f492a713d71dc15c59df971c1c92)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_devices/data-source.tf`; digest `sha256:59c1a13d7be54a7427ce8e310798e9737dbdf20c08bef5ed916ec4310ce79e92`.

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

<a id="canonical-25bc7ee71185612a5fba288f2819d655bcb4f2978deaf2f35bafa61c3bf938ba"></a>

## Next pages — Data source / f5ed97bd2e8d / 3

- [Examples](data-sources--device_intelligence_devices--examples--group-001.md#canonical-e29fe453efdf882c0fe3e41b13b48c36e410f492a713d71dc15c59df971c1c92)
- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)
