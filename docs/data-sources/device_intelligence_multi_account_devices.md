---
page_title: "xcsh_device_intelligence_multi_account_devices landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_multi_account_devices landing."
---

# xcsh_device_intelligence_multi_account_devices landing

<a id="canonical-d7032758bccd44c0b77f242e2c10eff4929329ac7188194b5b4d1630cd87bd40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97b3fe33c442be09e55a3280f7315af5b63c1790bb1322087994be785d1e0958"></a>

## xcsh_device_intelligence_multi_account_devices — xcsh_device_intelligence_multi_account_devices / 68089e80e65e / 2

Breadcrumbs:

- xcsh_device_intelligence_multi_account_devices

Resource creation operation.

<a id="canonical-ea05a9ee21c718c1c73926109601a10604c91d210cc6c5999d2734fd772f9678"></a>

## Prerequisites — xcsh_device_intelligence_multi_account_devices / 68089e80e65e / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2059222f6d30b20a9ac4999bebc55bde58535569c23f80dfb27abe8e89969772"></a>

## Minimal configuration — xcsh_device_intelligence_multi_account_devices / 68089e80e65e / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceMultiAccountDevices DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_multi_account_devices" "example" {
  namespace = "example-value"
}

output "device_intelligence_multi_account_devices_result" {
  value = data.xcsh_device_intelligence_multi_account_devices.example
}
```

<a id="canonical-9be7ebaed0486558fc49aea3785f6a1f0aa3591354547fd275b8f11c867968f4"></a>

## Root configuration — xcsh_device_intelligence_multi_account_devices / 68089e80e65e / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-76223d3f81521486e79fd03a9b6332e79860c8cce6a5cb82ffb3cfb29d8cc463"></a>

## Next pages — xcsh_device_intelligence_multi_account_devices / 68089e80e65e / 6

- [Property reference](../guides/data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-0297391df16a3a9dd9f80b68cd4db055feaadb98c0563c5575f355edfa4a9064)
- [Examples](../guides/data-sources--device_intelligence_multi_account_devices--examples--group-001.md#canonical-93853d1839c4e0042bb32f26d8621ff01e93335761b40166a77432b3e30de778)
