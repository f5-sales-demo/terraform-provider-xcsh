---
page_title: "xcsh_device_intelligence_risk_score_distribution landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_risk_score_distribution landing."
---

# xcsh_device_intelligence_risk_score_distribution landing

<a id="canonical-28a66890f008fdac1b26cd785850ece6401f38637446a0123c4e80cf02ad050e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7cca77d59337cf0869ad7564b8192f31ebf20a452da22a51d4f0154927d7941"></a>

## xcsh_device_intelligence_risk_score_distribution — xcsh_device_intelligence_risk_score_distribution / beaa791e8e4b / 2

Breadcrumbs:

- xcsh_device_intelligence_risk_score_distribution

Resource creation operation.

<a id="canonical-55bbb5e7bce0c01e2523b3236512ea67129d84318b5eb99152de547fa91e7fa3"></a>

## Prerequisites — xcsh_device_intelligence_risk_score_distribution / beaa791e8e4b / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-e0ed4169bedabe4bb3885666df89e5d5a5f6451536f315b8ee09532cf8c93c58"></a>

## Minimal configuration — xcsh_device_intelligence_risk_score_distribution / beaa791e8e4b / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceRiskScoreDistribution DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_risk_score_distribution" "example" {
  namespace = "example-value"
}

output "device_intelligence_risk_score_distribution_result" {
  value = data.xcsh_device_intelligence_risk_score_distribution.example
}
```

<a id="canonical-fac079cfc5e8e5daa58b55852a00c1ecc2124373641b58164a98fce434516439"></a>

## Root configuration — xcsh_device_intelligence_risk_score_distribution / beaa791e8e4b / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-235bce365f94c14c0b7313a6cfdd5a22aee81a7a2e54cbc1e1352d68e5ad025d"></a>

## Next pages — xcsh_device_intelligence_risk_score_distribution / beaa791e8e4b / 6

- [Property reference](../guides/data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-1f6bece68929c1bc8f8d705b9a282c5d577a38a8bc7f44e3d818748cb10109ed)
- [Examples](../guides/data-sources--device_intelligence_risk_score_distribution--examples--group-001.md#canonical-19b438ad97826262ff4db622135397459932ffb52a25c0ab1c3414697615a719)
