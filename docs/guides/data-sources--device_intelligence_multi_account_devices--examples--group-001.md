---
page_title: "xcsh_device_intelligence_multi_account_devices examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_multi_account_devices examples."
---

# xcsh_device_intelligence_multi_account_devices examples

<a id="canonical-93853d1839c4e0042bb32f26d8621ff01e93335761b40166a77432b3e30de778"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-216c344557bb5231e96eaaebba2a3b8a5f51dd7bb613813ab91b041dfe51660d"></a>

## Examples — Examples / 12162aa2b137 / 2

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-d7032758bccd44c0b77f242e2c10eff4929329ac7188194b5b4d1630cd87bd40)
- Examples

<a id="canonical-02abc5ad95168e9e2882f0e683e3249c8944b6312c814a822352a4e316595d6c"></a>

## Complete configurations — Examples / 12162aa2b137 / 3

- [Data source](data-sources--device_intelligence_multi_account_devices--examples--group-001.md#canonical-7054299f1000aa86e84372eb995052a56227ac5aa39e64b958d55f0f1ae096f6): valid configuration.

<a id="canonical-190199432cb16149e13e52aacfe113c5427f7662ff7565b14b922b6c99774318"></a>

## Next pages — Examples / 12162aa2b137 / 4

- [Data source](data-sources--device_intelligence_multi_account_devices--examples--group-001.md#canonical-7054299f1000aa86e84372eb995052a56227ac5aa39e64b958d55f0f1ae096f6)
- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-d7032758bccd44c0b77f242e2c10eff4929329ac7188194b5b4d1630cd87bd40)

<a id="canonical-7054299f1000aa86e84372eb995052a56227ac5aa39e64b958d55f0f1ae096f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1907e59467ea6c12b97e1f38d189538b706bd920da3f9c2190f6a9e4c89c7524"></a>

## Data source — Data source / 042344add4b3 / 2

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-d7032758bccd44c0b77f242e2c10eff4929329ac7188194b5b4d1630cd87bd40)
- [Examples](data-sources--device_intelligence_multi_account_devices--examples--group-001.md#canonical-93853d1839c4e0042bb32f26d8621ff01e93335761b40166a77432b3e30de778)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_multi_account_devices/data-source.tf`; digest `sha256:83db70480833f20bc19383610323bcf199f0707dd687a03af490d0098255cfcb`.

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

<a id="canonical-7afa75b8eee401700d15d94c0677131abd9643e861d0b7c8dd0ee174e702f6fc"></a>

## Next pages — Data source / 042344add4b3 / 3

- [Examples](data-sources--device_intelligence_multi_account_devices--examples--group-001.md#canonical-93853d1839c4e0042bb32f26d8621ff01e93335761b40166a77432b3e30de778)
- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-d7032758bccd44c0b77f242e2c10eff4929329ac7188194b5b4d1630cd87bd40)
