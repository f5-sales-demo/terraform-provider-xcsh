---
page_title: "xcsh_device_intelligence_risk_score_distribution"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_risk_score_distribution."
---

# xcsh_device_intelligence_risk_score_distribution

<a id="canonical-0220221212202100-3300002033312230-0123021230311320-1120110032303212-1000013303201203-1310101222000102-0330103220003033-0002223100110032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_device_intelligence_risk_score_distribution

Reads Device Intelligence Risk Score Distribution information from F5 Distributed Cloud.

<a id="canonical-3113303022131331-1121030313303300-2012212231131112-1023200121023303-0132233302002210-1102312202022211-0131103300011110-2102133113211001"></a>

### Prerequisites for `xcsh_device_intelligence_risk_score_distribution`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1111232323113213-2330320030000132-0211020323030203-1211010232221213-0102213120100301-2023113223212101-1102313211101333-2221013213332203"></a>

### Minimal configuration for `xcsh_device_intelligence_risk_score_distribution`

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

<a id="canonical-3200323110011221-2332312223321023-2303202011121212-3133202132113111-2211331210110111-0312330301112320-3232002111030230-3320302103301120"></a>

### Root configuration for `xcsh_device_intelligence_risk_score_distribution`

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3322300013213033-3011322032113122-2211202311112011-0222000030013230-3002010210031303-1210012311200112-1022212033303210-0310110112100321"></a>

### Explore this collection for `xcsh_device_intelligence_risk_score_distribution`

- [Property reference](../guides/data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-0133122332303212-2021022130012330-2033203113001123-2122022002301131-1113132203202220-2330133310103203-3120012013102030-2301000100213231)
- [Examples](../guides/data-sources--device_intelligence_risk_score_distribution--examples--group-001.md#canonical-0121231003202231-2113200212021202-3333103123120202-0103110321131011-2121030233332311-0222021130002223-0130031001101221-1312011122130121)
