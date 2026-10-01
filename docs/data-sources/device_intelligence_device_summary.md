---
page_title: "xcsh_device_intelligence_device_summary landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_device_summary landing."
---

# xcsh_device_intelligence_device_summary landing

<a id="canonical-f8e0600c589deef40fd9b13736eefcb6123534aaac445791790673b2e06e3989"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8c7b167ab83457da48d5789b0e2e484258975e8589986e32127115b15c37f1e"></a>

## xcsh_device_intelligence_device_summary — xcsh_device_intelligence_device_summary / 32e02e00cf94 / 2

Breadcrumbs:

- xcsh_device_intelligence_device_summary

Resource creation operation.

<a id="canonical-3f2509863caedaabcc42cb270c461b05e64dd6b12a7d7e42a7ae66f4e9d425f3"></a>

## Prerequisites — xcsh_device_intelligence_device_summary / 32e02e00cf94 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-52169d7da26d80aeab13922ccd86ed4036693ed61e7bd7c7ccdcee015ab9e471"></a>

## Minimal configuration — xcsh_device_intelligence_device_summary / 32e02e00cf94 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceDeviceSummary DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_device_summary" "example" {
  device_id = "example-value"
  namespace = "example-value"
}

output "device_intelligence_device_summary_result" {
  value = data.xcsh_device_intelligence_device_summary.example
}
```

<a id="canonical-c2a4ef0f52e549841643420661af7a0d637ba33cdb084b0cdc6de118074e6ec2"></a>

## Root configuration — xcsh_device_intelligence_device_summary / 32e02e00cf94 / 5

Required root properties: `device_id`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-9097c592b228fbedf7108dd72539ccf0bfea5a31bbd362392156087c23a42306"></a>

## Next pages — xcsh_device_intelligence_device_summary / 32e02e00cf94 / 6

- [Property reference](../guides/data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-93f159bc76d0b51ddc452c0d614270f715e8fbdad2fe323e8255e24d25ca5d9a)
- [Examples](../guides/data-sources--device_intelligence_device_summary--examples--group-001.md#canonical-f15d9ca0379e4fc7bac628f18fdc6e1e7930d4b49dbab055002e7b17b00a53d7)
