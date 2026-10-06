---
page_title: "xcsh_smsv2_kvm_runtime examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_kvm_runtime examples."
---

# xcsh_smsv2_kvm_runtime examples

<a id="canonical-1031313033231030-1220212302112133-2230021203323303-3231323231233110-3220033111033310-2001221100231033-0322233111011022-0231123303331001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_smsv2_kvm_runtime](../data-sources/smsv2_kvm_runtime.md#canonical-3112321232333312-1310102333312112-1232330010223001-2220301030033011-2013200023031111-0001200213232313-3011013301302013-2000100013012202)
- Examples

<a id="canonical-1230302010012332-0201123332202023-1220320101100312-3020130321110230-1212021212113222-1313220023120012-2113311300000120-2131132230113332"></a>

### Complete configurations for `xcsh_smsv2_kvm_runtime`

- [Data source](data-sources--smsv2_kvm_runtime--examples--group-001.md#canonical-2120112323021230-0321032311122313-2320331030113323-0303130013123001-1303231022102231-0203013112213213-1200131131101112-2332121310100202): valid configuration.

<a id="canonical-2120112323021230-0321032311122313-2320331030113323-0303130013123001-1303231022102231-0203013112213213-1200131131101112-2332121310100202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_smsv2_kvm_runtime](../data-sources/smsv2_kvm_runtime.md#canonical-3112321232333312-1310102333312112-1232330010223001-2220301030033011-2013200023031111-0001200213232313-3011013301302013-2000100013012202)
- [Examples](data-sources--smsv2_kvm_runtime--examples--group-001.md#canonical-1031313033231030-1220212302112133-2230021203323303-3231323231233110-3220033111033310-2001221100231033-0322233111011022-0231123303331001)
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
