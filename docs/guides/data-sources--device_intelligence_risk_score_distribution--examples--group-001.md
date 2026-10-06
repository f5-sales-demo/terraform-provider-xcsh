---
page_title: "xcsh_device_intelligence_risk_score_distribution examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_risk_score_distribution examples."
---

# xcsh_device_intelligence_risk_score_distribution examples

<a id="canonical-0121231003202231-2113200212021202-3333103123120202-0103110321131011-2121030233332311-0222021130002223-0130031001101221-1312011122130121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-0220221212202100-3300002033312230-0123021230311320-1120110032303212-1000013303201203-1310101222000102-0330103220003033-0002223100110032)
- Examples

<a id="canonical-1130210213330003-2313132302303033-0123112030232001-0300110223201021-3311331320220112-2311130233130110-3111010331313300-1133202223212313"></a>

### Complete configurations for `xcsh_device_intelligence_risk_score_distribution`

- [Data source](data-sources--device_intelligence_risk_score_distribution--examples--group-001.md#canonical-2201101332003111-2103021123300311-3303003233302120-3023332312221122-0232231133000323-2312023311212200-1310301330220333-0131023003231110): valid configuration.

<a id="canonical-2201101332003111-2103021123300311-3303003233302120-3023332312221122-0232231133000323-2312023311212200-1310301330220333-0131023003231110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-0220221212202100-3300002033312230-0123021230311320-1120110032303212-1000013303201203-1310101222000102-0330103220003033-0002223100110032)
- [Examples](data-sources--device_intelligence_risk_score_distribution--examples--group-001.md#canonical-0121231003202231-2113200212021202-3333103123120202-0103110321131011-2121030233332311-0222021130002223-0130031001101221-1312011122130121)
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
