---
page_title: "xcsh_smsv2_kvm_runtime landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_kvm_runtime landing."
---

# xcsh_smsv2_kvm_runtime landing

<a id="canonical-d6e6eff6744bfd966ef04ac1a8c4c3c58780b35501827bb7c51f1c87804071a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc57c6152ea787ecce397b39875e1cf23daeb7a1ba48c28d5fe1b51ae83fbed8"></a>

## xcsh_smsv2_kvm_runtime — xcsh_smsv2_kvm_runtime / 6cddbdce6883 / 2

Breadcrumbs:

- xcsh_smsv2_kvm_runtime

Resolves one realized KVM Secure Mesh Site v2 SLO network interface through site UID ownership, the
live registration hostname and device, and an expected MAC address. The name is observed, never
guessed.

<a id="canonical-09d56c4158daacf54499a4198957eb8aa48c7252f7501dce1beb8de478cbb937"></a>

## Prerequisites — xcsh_smsv2_kvm_runtime / 6cddbdce6883 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-a5fb384beccbbf6d5c0de4284e78ba151f5900f36a1f776d5e68b8c236e45e75"></a>

## Minimal configuration — xcsh_smsv2_kvm_runtime / 6cddbdce6883 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-40e0fab2e673f964ec27ab6dfccf345d885f66ed36b9b5f6d67a27325898db1f"></a>

## Root configuration — xcsh_smsv2_kvm_runtime / 6cddbdce6883 / 5

Required root properties: `expected_mac`, `site`. Full root flags and choices appear in the property reference.

<a id="canonical-fe2a74a85874e50f306026ba23b4c4affbbfa41027e77e2e4e5ec1870263de1b"></a>

## Next pages — xcsh_smsv2_kvm_runtime / 6cddbdce6883 / 6

- [Property reference](../guides/data-sources--smsv2_kvm_runtime--reference--group-001.md#canonical-c1611532e0fbc38a6575ad57eb82084c21632bc4f7f049080aa3ea62c86c7928)
- [Examples](../guides/data-sources--smsv2_kvm_runtime--examples--group-001.md#canonical-4ddcfb4c689b259fac263ef3edeedbd4e83d53f481a50b4f3abd514a2d6f3f41)
