---
page_title: "xcsh_waf_threat_campaigns examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_threat_campaigns examples."
---

# xcsh_waf_threat_campaigns examples

<a id="canonical-0221210003110010-0311133122200121-3130132231312010-3310202213023302-0312010220323313-1001113102001320-0112122321211020-0023322002031120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_waf_threat_campaigns](../data-sources/waf_threat_campaigns.md#canonical-1323321121302311-1032111022111323-1120132100221231-1221202233212221-1010312121000210-3221113220023012-3001030210113301-0002023032103021)
- Examples

<a id="canonical-3200023110210231-0112013013301310-1031313310310010-1101001020322312-2023312310323311-2000230012221120-0031200112220030-1013032122103200"></a>

### Complete configurations for `xcsh_waf_threat_campaigns`

- [Data source](data-sources--waf_threat_campaigns--examples--group-001.md#canonical-0103000131000021-1121211022313202-3220322000022023-1100212322033323-3001103323211032-3030233103300230-0312310112231220-0133320133131230): valid configuration.

<a id="canonical-0103000131000021-1121211022313202-3220322000022023-1100212322033323-3001103323211032-3030233103300230-0312310112231220-0133320133131230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_waf_threat_campaigns](../data-sources/waf_threat_campaigns.md#canonical-1323321121302311-1032111022111323-1120132100221231-1221202233212221-1010312121000210-3221113220023012-3001030210113301-0002023032103021)
- [Examples](data-sources--waf_threat_campaigns--examples--group-001.md#canonical-0221210003110010-0311133122200121-3130132231312010-3310202213023302-0312010220323313-1001113102001320-0112122321211020-0023322002031120)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_threat_campaigns/data-source.tf`; digest `sha256:87da0e1071c2bec565e2d7e725d9e896914625d820293d688bc4d4b1d617b7b5`.

```terraform
# WAFThreatCampaigns DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_threat_campaigns" "example" {
}

output "waf_threat_campaigns_result" {
  value = data.xcsh_waf_threat_campaigns.example
}
```
