---
page_title: "xcsh_smsv2_kvm_runtime_interface examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_kvm_runtime_interface examples."
---

# xcsh_smsv2_kvm_runtime_interface examples

<a id="canonical-c509cd984ffc370cab24b54c647c3de108332d63fff3a31baacfaacd3e5f847d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50670bca8ed00fd82e56a1e604a26fa7ba19516b8b0587bcd824a8e4377141ec"></a>

## Examples — Examples / 5fa2d49a5b05 / 2

Breadcrumbs:

- [xcsh_smsv2_kvm_runtime_interface](../resources/smsv2_kvm_runtime_interface.md#canonical-2f0c16a5489e111182e91669fdd8243ae8c574bdb1fc5fc1e4ab1808527039b3)
- Examples

<a id="canonical-ca219b431377fa0c0cfac3ce7b5a8c3f3c7a43873095a53aadc49e2b7ad6203b"></a>

## Complete configurations — Examples / 5fa2d49a5b05 / 3

- [Resource](resources--smsv2_kvm_runtime_interface--examples--group-001.md#canonical-91b0cc65513bb3d04cd4ce2b7c0584ab17f549043bb2ee020af8cbade8aef977): valid configuration.

<a id="canonical-a8cd5033ae78948985e006d28d9503733bf59264877ce007c85a8460cfdcf896"></a>

## Next pages — Examples / 5fa2d49a5b05 / 4

- [Resource](resources--smsv2_kvm_runtime_interface--examples--group-001.md#canonical-91b0cc65513bb3d04cd4ce2b7c0584ab17f549043bb2ee020af8cbade8aef977)
- [xcsh_smsv2_kvm_runtime_interface](../resources/smsv2_kvm_runtime_interface.md#canonical-2f0c16a5489e111182e91669fdd8243ae8c574bdb1fc5fc1e4ab1808527039b3)

<a id="canonical-91b0cc65513bb3d04cd4ce2b7c0584ab17f549043bb2ee020af8cbade8aef977"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c36241a1e1a627c8da67226a970ca1e684b818e6373d44560562ee2071a7f5aa"></a>

## Resource — Resource / 1ad4cdd0a99e / 2

Breadcrumbs:

- [xcsh_smsv2_kvm_runtime_interface](../resources/smsv2_kvm_runtime_interface.md#canonical-2f0c16a5489e111182e91669fdd8243ae8c574bdb1fc5fc1e4ab1808527039b3)
- [Examples](resources--smsv2_kvm_runtime_interface--examples--group-001.md#canonical-c509cd984ffc370cab24b54c647c3de108332d63fff3a31baacfaacd3e5f847d)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_smsv2_kvm_runtime_interface/resource.tf`; digest `sha256:3203711bd7a88fd8f1838fcf238c8f13492e522a7f5e46972dc68cf409895eb4`.

```terraform
# Adopt the exact XC-owned SLI child discovered after a KVM CE registers.

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 10.0.0"
    }
  }
}

resource "xcsh_smsv2_kvm_runtime_interface" "sli" {
  namespace    = "system"
  site         = "onprem-example-kvm"
  expected_mac = "52:54:00:20:00:11"
  ipv4_cidr    = "10.201.0.11/24"
}
```

<a id="canonical-11d797c441dacc78ccd0df27f78e325f90c2751097a40bd14bb06e2c0a0ec927"></a>

## Next pages — Resource / 1ad4cdd0a99e / 3

- [Examples](resources--smsv2_kvm_runtime_interface--examples--group-001.md#canonical-c509cd984ffc370cab24b54c647c3de108332d63fff3a31baacfaacd3e5f847d)
- [xcsh_smsv2_kvm_runtime_interface](../resources/smsv2_kvm_runtime_interface.md#canonical-2f0c16a5489e111182e91669fdd8243ae8c574bdb1fc5fc1e4ab1808527039b3)
