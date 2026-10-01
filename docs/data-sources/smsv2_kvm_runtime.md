---
page_title: "xcsh_smsv2_kvm_runtime landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_kvm_runtime landing."
---

# xcsh_smsv2_kvm_runtime landing

<a id="canonical-3112321232333312-1310102333312112-1232330010223001-2220301030033011-2013200023031111-0001200213232313-3011013301302013-2000100013012202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030111330120111-0232221320133230-3032032113230321-2013113201303302-0331223223132201-2322102030022031-1133320123110122-3220033323323120"></a>

## xcsh_smsv2_kvm_runtime — xcsh_smsv2_kvm_runtime / 303212202003 / 2

Breadcrumbs:

- xcsh_smsv2_kvm_runtime

Resolves one realized KVM Secure Mesh Site v2 SLO network interface through site UID ownership, the
live registration hostname and device, and an expected MAC address. The name is observed, never
guessed.

<a id="canonical-0021311112301001-1120312222303311-1010212122100121-2021111332232022-2210203013021102-3313110001313032-0123322320313210-1320302323210313"></a>

## Prerequisites — xcsh_smsv2_kvm_runtime / 303212202003 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2211332303201023-3230302323331231-1130003132100220-1032132023220111-0133112100003303-1222013313131231-1132122023203002-0312321011321311"></a>

## Minimal configuration — xcsh_smsv2_kvm_runtime / 303212202003 / 4

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

<a id="canonical-1000320033222302-3212130333211210-3230021322231231-3330303303101131-2020113312123231-0312232123113312-3112132202130302-1120212031230133"></a>

## Root configuration — xcsh_smsv2_kvm_runtime / 303212202003 / 5

Required root properties: `expected_mac`, `site`. Full root flags and choices appear in the property reference.

<a id="canonical-3332022213102220-1120131032110033-0300120002122322-0203231030102233-3323233322100100-0213321313320232-1032113230012013-0002120331320123"></a>

## Next pages — xcsh_smsv2_kvm_runtime / 303212202003 / 6

- [Property reference](../guides/data-sources--smsv2_kvm_runtime--reference--group-001.md#canonical-3001120101110302-3200332330032022-1211131122311113-3223200200201030-0201120302233010-3313330010210020-0022220332221202-3020123013210220)
- [Examples](../guides/data-sources--smsv2_kvm_runtime--examples--group-001.md#canonical-1031313033231030-1220212302112133-2230021203323303-3231323231233110-3220033111033310-2001221100231033-0322233111011022-0231123303331001)
