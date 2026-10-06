---
page_title: "xcsh_waf_threat_campaigns"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_threat_campaigns."
---

# xcsh_waf_threat_campaigns

<a id="canonical-1323321121302311-1032111022111323-1120132100221231-1221202233212221-1010312121000210-3221113220023012-3001030210113301-0002023032103021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_waf_threat_campaigns

Reads WAF Threat Campaigns information from F5 Distributed Cloud.

<a id="canonical-3333101333101303-3302233113031323-0112000121223333-2100101223030220-2110122132122110-0100101011000300-2120203333313331-2233023213030110"></a>

### Prerequisites for `xcsh_waf_threat_campaigns`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2301133010020232-2201103313313202-3100122312233320-3130310220010302-0030001330202303-0022211230202301-0223023323101020-0310200221232210"></a>

### Minimal configuration for `xcsh_waf_threat_campaigns`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-3001203121320003-3102301011103230-0210123213021003-0131333123130311-1103303312021200-3102013121100202-3022013103200032-2120133221111200"></a>

### Root configuration for `xcsh_waf_threat_campaigns`

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-1313003021021011-3102022303302001-3120220021010000-3313112012012133-1001202310220303-0201230113222113-1000032103200001-2001132113021322"></a>

### Explore this collection for `xcsh_waf_threat_campaigns`

- [Property reference](../guides/data-sources--waf_threat_campaigns--reference--group-001.md#canonical-2312320213131010-2313332320110100-1222123311130230-1202203132203210-2221111330220221-2123233132102221-2022012303331212-0311010310111022)
- [Examples](../guides/data-sources--waf_threat_campaigns--examples--group-001.md#canonical-0221210003110010-0311133122200121-3130132231312010-3310202213023302-0312010220323313-1001113102001320-0112122321211020-0023322002031120)
