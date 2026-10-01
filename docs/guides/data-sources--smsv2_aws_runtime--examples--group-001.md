---
page_title: "xcsh_smsv2_aws_runtime examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_aws_runtime examples."
---

# xcsh_smsv2_aws_runtime examples

<a id="canonical-1f1ef6b6795670d3320b1dbfef99fe9e58367a1eb0223c900d0fe48596f3d807"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7845b6edbe5be46f7c947138221ed61a95fe5fab2d86b61d8573c9e8e858ccfd"></a>

## Examples — Examples / 3da1ecfa32fd / 2

Breadcrumbs:

- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md#canonical-585eec57272cb79858684d22159d08a11b9cc9e8b386856fa87be603c746a674)
- Examples

<a id="canonical-07526202d164b9424287ae50f7e9eb344eaa4701bcaf7ca1b3765a738279d416"></a>

## Complete configurations — Examples / 3da1ecfa32fd / 3

- [Data source](data-sources--smsv2_aws_runtime--examples--group-001.md#canonical-6838a886adfd74af24e7b3261474ddaa8fcf082ae95308870dc7359e9071bcf6): valid configuration.

<a id="canonical-323bf798897ce947e98988877d59c88bf28a99967baaf0184c07f3a046838957"></a>

## Next pages — Examples / 3da1ecfa32fd / 4

- [Data source](data-sources--smsv2_aws_runtime--examples--group-001.md#canonical-6838a886adfd74af24e7b3261474ddaa8fcf082ae95308870dc7359e9071bcf6)
- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md#canonical-585eec57272cb79858684d22159d08a11b9cc9e8b386856fa87be603c746a674)

<a id="canonical-6838a886adfd74af24e7b3261474ddaa8fcf082ae95308870dc7359e9071bcf6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43e791c39d36045409da86f58497e1e32928d0bac8cc059a01bf8608388b2620"></a>

## Data source — Data source / 2d79ed5473c8 / 2

Breadcrumbs:

- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md#canonical-585eec57272cb79858684d22159d08a11b9cc9e8b386856fa87be603c746a674)
- [Examples](data-sources--smsv2_aws_runtime--examples--group-001.md#canonical-1f1ef6b6795670d3320b1dbfef99fe9e58367a1eb0223c900d0fe48596f3d807)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_smsv2_aws_runtime/data-source.tf`; digest `sha256:e299ec8ab27b0aebf0736cf9e2b336fdc7524df91d37a6ad7dac1a9e71242dd5`.

```terraform
# Correlate stable logical node keys and AWS-authoritative ENI MAC addresses
# with the SMSv2 interface configuration and runtime health observed by F5 XC.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 6.0.0"
    }
  }
}

data "xcsh_smsv2_aws_runtime" "site" {
  namespace = "system"
  site      = "example-smsv2-site"

  nodes = {
    node_0_slo = {
      node = "node-0"
      role = "slo"
      mac  = "02:00:00:00:00:10"
    }
    node_0_sli = {
      node = "node-0"
      role = "sli"
      mac  = "02:00:00:00:00:11"
    }
  }
}

output "smsv2_interfaces" {
  value = data.xcsh_smsv2_aws_runtime.site.interfaces
}

output "smsv2_healthy" {
  value = data.xcsh_smsv2_aws_runtime.site.healthy
}
```

<a id="canonical-66f027a9f7ffa7bb8545110d919b94528a2ef059f8a1a680784e3a7b0131b9e1"></a>

## Next pages — Data source / 2d79ed5473c8 / 3

- [Examples](data-sources--smsv2_aws_runtime--examples--group-001.md#canonical-1f1ef6b6795670d3320b1dbfef99fe9e58367a1eb0223c900d0fe48596f3d807)
- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md#canonical-585eec57272cb79858684d22159d08a11b9cc9e8b386856fa87be603c746a674)
