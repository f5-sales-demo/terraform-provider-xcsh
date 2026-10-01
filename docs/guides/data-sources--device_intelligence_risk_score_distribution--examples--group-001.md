---
page_title: "xcsh_device_intelligence_risk_score_distribution examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_risk_score_distribution examples."
---

# xcsh_device_intelligence_risk_score_distribution examples

<a id="canonical-19b438ad97826262ff4db622135397459932ffb52a25c0ab1c3414697615a719"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c927f03b77b2ccf1b58cb813052b849f5f78a16b572f714d513ddf05f8ab9b7"></a>

## Examples — Examples / 0a2817c2da20 / 2

Breadcrumbs:

- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-28a66890f008fdac1b26cd785850ece6401f38637446a0123c4e80cf02ad050e)
- Examples

<a id="canonical-fb0879da018c0166935c79971b2d37dc220c764faa03a9a817f05c3971307786"></a>

## Complete configurations — Examples / 0a2817c2da20 / 3

- [Data source](data-sources--device_intelligence_risk_score_distribution--examples--group-001.md#canonical-a147e0d59325bc35f30efc98cbfb6a5a2eb5f03bb62f59a074c7ca3f1d2c3b54): valid configuration.

<a id="canonical-dfb55a08f0331ec67b7670845e6027e481a102bf44df18fd5c8f50d59fa82c77"></a>

## Next pages — Examples / 0a2817c2da20 / 4

- [Data source](data-sources--device_intelligence_risk_score_distribution--examples--group-001.md#canonical-a147e0d59325bc35f30efc98cbfb6a5a2eb5f03bb62f59a074c7ca3f1d2c3b54)
- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-28a66890f008fdac1b26cd785850ece6401f38637446a0123c4e80cf02ad050e)

<a id="canonical-a147e0d59325bc35f30efc98cbfb6a5a2eb5f03bb62f59a074c7ca3f1d2c3b54"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-548e10aec99fecb7e509afebf55ecd6e35add66ffcfee7bf9ccd1a0ae967aeb0"></a>

## Data source — Data source / 13f654698aea / 2

Breadcrumbs:

- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-28a66890f008fdac1b26cd785850ece6401f38637446a0123c4e80cf02ad050e)
- [Examples](data-sources--device_intelligence_risk_score_distribution--examples--group-001.md#canonical-19b438ad97826262ff4db622135397459932ffb52a25c0ab1c3414697615a719)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_risk_score_distribution/data-source.tf`; digest `sha256:c10972aeff2a32c44eb299757977d3dc58259ac93f292c8fdb4d0bb8272ceb96`.

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

<a id="canonical-585e74bdc0c1f8de8478c888ac79e279baab83f0213893eb6a3413927240cfda"></a>

## Next pages — Data source / 13f654698aea / 3

- [Examples](data-sources--device_intelligence_risk_score_distribution--examples--group-001.md#canonical-19b438ad97826262ff4db622135397459932ffb52a25c0ab1c3414697615a719)
- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-28a66890f008fdac1b26cd785850ece6401f38637446a0123c4e80cf02ad050e)
