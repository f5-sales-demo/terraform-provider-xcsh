---
page_title: "xcsh_smsv2_kvm_runtime examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_kvm_runtime examples."
---

# xcsh_smsv2_kvm_runtime examples

<a id="canonical-4ddcfb4c689b259fac263ef3edeedbd4e83d53f481a50b4f3abd514a2d6f3f41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6cc841be216fe88b68e11436c873952c662665ea77a0b60697d700189d7ac5fe"></a>

## Examples — Examples / 250b0a0de8c4 / 2

Breadcrumbs:

- [xcsh_smsv2_kvm_runtime](../data-sources/smsv2_kvm_runtime.md#canonical-d6e6eff6744bfd966ef04ac1a8c4c3c58780b35501827bb7c51f1c87804071a2)
- Examples

<a id="canonical-3e86eee5913150e944525828f6bd653968167dc719224df177a608e12b01a396"></a>

## Complete configurations — Examples / 250b0a0de8c4 / 3

- [Data source](data-sources--smsv2_kvm_runtime--examples--group-001.md#canonical-985bb26c393b56b7b8f4c5fb337076c173b4a4ad231d69e76075d456be674422): valid configuration.

<a id="canonical-90567e8bf23d55e8340336471494da2af22c3c71dbcc11233c726bb9b22f975c"></a>

## Next pages — Examples / 250b0a0de8c4 / 4

- [Data source](data-sources--smsv2_kvm_runtime--examples--group-001.md#canonical-985bb26c393b56b7b8f4c5fb337076c173b4a4ad231d69e76075d456be674422)
- [xcsh_smsv2_kvm_runtime](../data-sources/smsv2_kvm_runtime.md#canonical-d6e6eff6744bfd966ef04ac1a8c4c3c58780b35501827bb7c51f1c87804071a2)

<a id="canonical-985bb26c393b56b7b8f4c5fb337076c173b4a4ad231d69e76075d456be674422"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81e181b0bda9370535c90fe2226055b64676f682abdca66a2b4efe5cdea8c6b9"></a>

## Data source — Data source / 835adb892a0b / 2

Breadcrumbs:

- [xcsh_smsv2_kvm_runtime](../data-sources/smsv2_kvm_runtime.md#canonical-d6e6eff6744bfd966ef04ac1a8c4c3c58780b35501827bb7c51f1c87804071a2)
- [Examples](data-sources--smsv2_kvm_runtime--examples--group-001.md#canonical-4ddcfb4c689b259fac263ef3edeedbd4e83d53f481a50b4f3abd514a2d6f3f41)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_smsv2_kvm_runtime/data-source.tf`; digest `sha256:59afa846179726aea892d691cd1d6b9fa1a3b1c0e460e92807612ddcb56e2ba6`.

```terraform
# Resolve the exact XC interface object created for one registered KVM CE.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 9.5.2"
    }
  }
}

data "xcsh_smsv2_kvm_runtime" "ce" {
  namespace    = "system"
  site         = "example-kvm-smsv2-site"
  expected_mac = "52:54:00:10:00:11"
}

output "kvm_interface_name" {
  value = data.xcsh_smsv2_kvm_runtime.ce.interface_name
}

output "kvm_registration_device" {
  value = data.xcsh_smsv2_kvm_runtime.ce.device
}
```

<a id="canonical-bb167df39f13f77b9a91628ea787588ed81c17f2937d2206e76eed7af4c09020"></a>

## Next pages — Data source / 835adb892a0b / 3

- [Examples](data-sources--smsv2_kvm_runtime--examples--group-001.md#canonical-4ddcfb4c689b259fac263ef3edeedbd4e83d53f481a50b4f3abd514a2d6f3f41)
- [xcsh_smsv2_kvm_runtime](../data-sources/smsv2_kvm_runtime.md#canonical-d6e6eff6744bfd966ef04ac1a8c4c3c58780b35501827bb7c51f1c87804071a2)
