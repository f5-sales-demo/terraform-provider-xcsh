---
page_title: "xcsh_smsv2_kvm_runtime_interface examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_kvm_runtime_interface examples."
---

# xcsh_smsv2_kvm_runtime_interface examples

<a id="canonical-3011002130312120-1033333003130030-2223021023111030-1210133003313201-0020030302311203-3333330322030123-2222303322223031-0332113320101331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_smsv2_kvm_runtime_interface](../resources/smsv2_kvm_runtime_interface.md#canonical-0233003001122211-1020213201010101-2002322101121221-3331312002100322-3220301113102331-2301333011333001-3210222301200020-1102130003212303)
- Examples

<a id="canonical-1100121300233022-2032310000333120-0232111222013212-0010220212332213-2322012111011223-2023001120132330-3120021022203210-0313130110013230"></a>

### Complete configurations for `xcsh_smsv2_kvm_runtime_interface`

- [Resource](resources--smsv2_kvm_runtime_interface--examples--group-001.md#canonical-2101230030301211-1101032323033100-1030311030320223-1330001120102223-0113331110210010-0323230232320002-0022332030232231-3220223233211313): valid configuration.

<a id="canonical-2101230030301211-1101032323033100-1030311030320223-1330001120102223-0113331110210010-0323230232320002-0022332030232231-3220223233211313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_smsv2_kvm_runtime_interface](../resources/smsv2_kvm_runtime_interface.md#canonical-0233003001122211-1020213201010101-2002322101121221-3331312002100322-3220301113102331-2301333011333001-3210222301200020-1102130003212303)
- [Examples](resources--smsv2_kvm_runtime_interface--examples--group-001.md#canonical-3011002130312120-1033333003130030-2223021023111030-1210133003313201-0020030302311203-3333330322030123-2222303322223031-0332113320101331)
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
