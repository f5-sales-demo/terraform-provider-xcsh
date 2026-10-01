---
page_title: "xcsh_smsv2_kvm_runtime_interface landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_kvm_runtime_interface landing."
---

# xcsh_smsv2_kvm_runtime_interface landing

<a id="canonical-2f0c16a5489e111182e91669fdd8243ae8c574bdb1fc5fc1e4ab1808527039b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7419d295b36832e85f1fc5f8bccb3f33336ac9af25f2c04ac1b1488429c1ac4"></a>

## xcsh_smsv2_kvm_runtime_interface — xcsh_smsv2_kvm_runtime_interface / e39e9ebd75d3 / 2

Breadcrumbs:

- xcsh_smsv2_kvm_runtime_interface

Adopts one existing XC-owned KVM Secure Mesh Site v2 SLI child and manages only its DHCP/static IPv4
mode. It never creates or deletes the runtime child.

<a id="canonical-8f56999d56af45026aac106c5dbcba6e1936477ae0946f91c99d38490a73c4e0"></a>

## Prerequisites — xcsh_smsv2_kvm_runtime_interface / e39e9ebd75d3 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-6690c5fd96291476327509e59ac45f1c1af3521a44e3a616cc617a68a3a086ec"></a>

## Minimal configuration — xcsh_smsv2_kvm_runtime_interface / e39e9ebd75d3 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-4bc06b86d6c81bc4f4784a27c0e268cc4820e2a7e0ce2f9b81a10fcc2b3dc70e"></a>

## Root configuration — xcsh_smsv2_kvm_runtime_interface / e39e9ebd75d3 / 5

Required root properties: `expected_mac`, `ipv4_cidr`, `site`. Full root flags and choices appear in the property reference.

<a id="canonical-2ac613ae08d44ebbcf618040c6266b030eacc9969a600d3023cf716b62bda003"></a>

## Next pages — xcsh_smsv2_kvm_runtime_interface / e39e9ebd75d3 / 6

- [Property reference](../guides/resources--smsv2_kvm_runtime_interface--reference--group-001.md#canonical-c66a286cbeb4e0f17e9c67b0ec0c11e7d79620ab55229d34ee040b173be79841)
- [Examples](../guides/resources--smsv2_kvm_runtime_interface--examples--group-001.md#canonical-c509cd984ffc370cab24b54c647c3de108332d63fff3a31baacfaacd3e5f847d)
