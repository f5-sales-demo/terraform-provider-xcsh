---
page_title: "xcsh_smsv2_kvm_runtime_interface landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_kvm_runtime_interface landing."
---

# xcsh_smsv2_kvm_runtime_interface landing

<a id="canonical-0233003001122211-1020213201010101-2002322101121221-3331312002100322-3220301113102331-2301333011333001-3210222301200020-1102130003212303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213100121310221-1123031220030232-2011330133301133-2023303023033303-0303031222302122-3302113302300010-2230012301102020-1002213001223010"></a>

## xcsh_smsv2_kvm_runtime_interface — xcsh_smsv2_kvm_runtime_interface / 233113113103 / 2

Breadcrumbs:

- xcsh_smsv2_kvm_runtime_interface

Adopts one existing XC-owned KVM Secure Mesh Site v2 SLI child and manages only its DHCP/static IPv4
mode. It never creates or deletes the runtime child.

<a id="canonical-2033111221212131-1112223310110002-1222223001001230-1131233023221232-0121031210131322-3200211012332101-3021213103201021-0022130330103200"></a>

## Prerequisites — xcsh_smsv2_kvm_runtime_interface / 233113113103 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1212210030113331-2112022101101312-0302131100213211-2122301011330130-0122330311020122-1010320322120112-3030120113221220-2203220020123230"></a>

## Minimal configuration — xcsh_smsv2_kvm_runtime_interface / 233113113103 / 4

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

<a id="canonical-1023300012232012-3112302001233010-3310132010220213-3000320212203030-1020020032022213-3200303202332123-2001220100333030-0223033130130032"></a>

## Root configuration — xcsh_smsv2_kvm_runtime_interface / 233113113103 / 5

Required root properties: `expected_mac`, `ipv4_cidr`, `site`. Full root flags and choices appear in the property reference.

<a id="canonical-0222301201032232-0020311010322323-3033120120001000-3012021212230003-0032223030212112-2122120000310300-0203303313011223-1202233122000003"></a>

## Next pages — xcsh_smsv2_kvm_runtime_interface / 233113113103 / 6

- [Property reference](../guides/resources--smsv2_kvm_runtime_interface--reference--group-001.md#canonical-3012122202201230-2332231032003301-1332213012132300-3230003001013213-3113211202002223-1111020221310310-3232001000230113-0323321321201001)
- [Examples](../guides/resources--smsv2_kvm_runtime_interface--examples--group-001.md#canonical-3011002130312120-1033333003130030-2223021023111030-1210133003313201-0020030302311203-3333330322030123-2222303322223031-0332113320101331)
