---
page_title: "xcsh_waf_threat_campaigns examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_threat_campaigns examples."
---

# xcsh_waf_threat_campaigns examples

<a id="canonical-29903504357da819dc7add84f48a72f236128ef7415d2078166b99480be82358"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e02d492d161c7c744ddf4d0451048eb68bdb4ef580b06a580d816a0c4739a4e0"></a>

## Examples — Examples / 7e160d7fa449 / 2

Breadcrumbs:

- [xcsh_waf_threat_campaigns](../data-sources/waf_threat_campaigns.md#canonical-7be59cb54e54a57b58790a6d698af9a944d99024e95e82c6c13245f1022ce4c9)
- Examples

<a id="canonical-969ca3c9e1e0cb4fc9fed04dd4d1bfe8141b1a6cdc66f989dedefe9c8330ff3e"></a>

## Complete configurations — Examples / 7e160d7fa449 / 3

- [Data source](data-sources--waf_threat_campaigns--examples--group-001.md#canonical-1301d0095994ade2e8e8028b509ba3fbc14fb94eccbd3c2c36d16b681fe1f76c): valid configuration.

<a id="canonical-f64215927adc2f4cec0b5de9cb7ec8c26c959a5d127522f7bbdcd2dd032f12e2"></a>

## Next pages — Examples / 7e160d7fa449 / 4

- [Data source](data-sources--waf_threat_campaigns--examples--group-001.md#canonical-1301d0095994ade2e8e8028b509ba3fbc14fb94eccbd3c2c36d16b681fe1f76c)
- [xcsh_waf_threat_campaigns](../data-sources/waf_threat_campaigns.md#canonical-7be59cb54e54a57b58790a6d698af9a944d99024e95e82c6c13245f1022ce4c9)

<a id="canonical-1301d0095994ade2e8e8028b509ba3fbc14fb94eccbd3c2c36d16b681fe1f76c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5281863a31a95d246d0354ff9dffe2e65ff74c0c72dfc40624057c0010ff716d"></a>

## Data source — Data source / 21cc693a268d / 2

Breadcrumbs:

- [xcsh_waf_threat_campaigns](../data-sources/waf_threat_campaigns.md#canonical-7be59cb54e54a57b58790a6d698af9a944d99024e95e82c6c13245f1022ce4c9)
- [Examples](data-sources--waf_threat_campaigns--examples--group-001.md#canonical-29903504357da819dc7add84f48a72f236128ef7415d2078166b99480be82358)
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

<a id="canonical-2915eda99342b622ace7529c0148c854ce52446a6cec8ef6241fbc96d4971711"></a>

## Next pages — Data source / 21cc693a268d / 3

- [Examples](data-sources--waf_threat_campaigns--examples--group-001.md#canonical-29903504357da819dc7add84f48a72f236128ef7415d2078166b99480be82358)
- [xcsh_waf_threat_campaigns](../data-sources/waf_threat_campaigns.md#canonical-7be59cb54e54a57b58790a6d698af9a944d99024e95e82c6c13245f1022ce4c9)
